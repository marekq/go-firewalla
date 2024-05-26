package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	mystruct "firewalla/mystructs"
)

// define files and folders
var (
	configFile     string
	devicesFile    string
	alarmsFile     string
	flowlogsFolder string
)

var timeLayout = "2006-01-02_15-04"

// init file and folder paths
func init() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFilePath)

	configFile = filepath.Join(currentDir, "config.json")
	devicesFile = filepath.Join(currentDir, "devices.json")
	alarmsFile = filepath.Join(currentDir, "alarms.json")
	flowlogsFolder = filepath.Join(currentDir, "flowlogs/")

}

// create global counter and waitgroup
var counter int
var wg sync.WaitGroup

// read config.json with url and token
func readJsonConfig() (string, string) {

	jsonFile, err := os.Open(configFile)

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	// read the token variable from the config file
	var token mystruct.JsonToken
	json.Unmarshal(byteValue, &token)

	defer jsonFile.Close()

	// return url and token
	return token.Url, "Token " + token.Token

}

// make get request to api, return body
func makeGetRequest(client *http.Client, url string, token string) []byte {

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	if err != nil {
		log.Fatal("Error creating request ", err)
	}

	// set headers
	req.Header.Set("Authorization", token)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request ", err, req)
	}

	if resp.StatusCode != 200 {
		log.Fatal("Status code error:", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return body

}

// get devices
func getDevices(client *http.Client, url string, token string, limit int64) {

	// Reset global counter
	counter = 0

	fmt.Println("* devices started")
	wg = sync.WaitGroup{}

	// Create file
	file, err := os.Create(devicesFile)
	if err != nil {
		log.Fatal("Error creating file ", err)
	}
	defer file.Close()

	body := makeGetRequest(client, url+"devices", token)
	devices := mystruct.FirewallaDevices{}
	json.Unmarshal([]byte(body), &devices)

	for _, device := range devices {

		fmt.Println(" * ", device)
		counter++

		// write deviceDetail to file
		json.NewEncoder(file).Encode(device)

	}

	fmt.Println("* completed -", counter, "devices saved to devices.json")
}

// get alarms
func getAlarms(client *http.Client, baseUrl, token string, limit int64, hours int64) {

	fmt.Println("* alarms started")

	// Reset global counter
	counter = 0

	// Create file
	file, err := os.Create(alarmsFile)
	if err != nil {
		log.Fatal("Error creating file ", err)
	}
	defer file.Close()

	groupBy := "ts,status,alarmType,summary,box,device,deviceIPAddress,category,domain,destination"
	encodedGroupBy := "groupBy=" + url.QueryEscape(groupBy)

	fullUrl := baseUrl + "alarms" + "?" + encodedGroupBy + "&limit=" + strconv.FormatInt(limit, 10)

	body := makeGetRequest(client, fullUrl, token)
	alarms := mystruct.FirewallaAlarms{}
	json.Unmarshal([]byte(body), &alarms)

	fmt.Println("* alarms started")
	wg = sync.WaitGroup{}

	fmt.Println("* found", strconv.Itoa(alarms.Count), "alarms")

	for _, alarm := range alarms.Results {

		// get alarm detail
		wg.Add(1)
		go getAlarmDetail(client, baseUrl, token, strconv.Itoa(alarm.Aid), alarm.Gid, file)

	}

	wg.Wait()

	fmt.Println("* completed -", counter, "alarms saved to alarms.json")
}

// get alarm detail
func getAlarmDetail(client *http.Client, baseUrl string, token string, aid string, gid string, file *os.File) {

	fullUrl := baseUrl + "alarms/" + gid + "/" + aid
	body := makeGetRequest(client, fullUrl, token)

	alarmDetail := mystruct.FirewallaAlarmDetail{}
	json.Unmarshal([]byte(body), &alarmDetail)
	fmt.Println(" % ", alarmDetail)

	// convert float64 to int
	ts := int64(alarmDetail.Ts)
	datestr := time.Unix(ts, 0).Format(time.RFC3339)
	alarmDetail.Date = datestr

	counter++

	if counter%5 == 0 {
		fmt.Println("-", counter, "alarms completed")
	}

	// write alarmDetail to file
	json.NewEncoder(file).Encode(alarmDetail)

	wg.Done()
}

// get newest timestamp from flowlogs folder
// if no timestamp found, return default timestamp
func getFlowlogsFirstTs(hours int64) time.Time {

	// set default timestamp
	newestTimestamp := time.Now().Add(-time.Duration(hours) * time.Hour)

	// get newest timestamp from flowlogs folder
	err := filepath.Walk(flowlogsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// get timestamp from csv filename
		if strings.Contains(info.Name(), ".csv") {

			// get timestamp from filename by searching for 10 digit timestamp
			re := regexp.MustCompile(`_(\d{10})\.csv$`)
			matches := re.FindStringSubmatch(info.Name())

			// if timestamp found
			if len(matches) > 0 {

				ts, err := strconv.ParseInt(matches[1], 10, 64)
				if err != nil {
					log.Fatal("Error converting timestamp ", err)
				}

				timeTs := time.Unix(ts, 0)

				// check if newer than newestTimestamp
				if timeTs.After(newestTimestamp) {
					newestTimestamp = timeTs
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("! error walking through the directory: ", err)

	}

	if newestTimestamp.IsZero() {
		fmt.Println("! no flowlog folder found, using ", hours, " hours")

		// return X hours ago
		return time.Now().Add(-time.Duration(hours) * time.Hour)

	}

	// return newestTimestamp
	return newestTimestamp
}

// get flowlogs
func getFlowLogs(client *http.Client, baseUrl string, token string, limitFlag int64, hoursFlag int64, debugFlag bool) {

	fmt.Println("* flowlogs started - ", hoursFlag, "hours", "limit", limitFlag, "debug", debugFlag)

	// reset global counter
	counter = 0
	wg = sync.WaitGroup{}

	// create firstTs and lastTs
	firstTs := getFlowlogsFirstTs(hoursFlag).Unix()
	lastTs := time.Now().Unix()

	// round down the firstTs to the nearest full hour (i.e. 15:30 -> 15:00)
	firstTs -= firstTs % 3600

	// Loop through all hours
	for currentTs := firstTs; currentTs < lastTs; currentTs += 3600 {

		// start and end time
		startTime := time.Unix(currentTs, 0)
		endTime := time.Unix(currentTs+3600, 0)

		// date path: dd/mm/yyyy/
		datePath := fmt.Sprintf("/%d/%02d/%02d/", startTime.Year(), startTime.Month(), startTime.Day())
		err := os.MkdirAll(flowlogsFolder+datePath, 0755)
		if err != nil {
			fmt.Println("Error creating folders:", err)
			return
		}

		// file name: flowlogs_YYYYMMDDTHHMM_to_YYYYMMDDTHHMM.csv
		logFileName := fmt.Sprintf("%s_%d.csv", startTime.Format(timeLayout), currentTs)

		// full path: flowlogs/2020/01/01/20200101T0000.csv
		logFilePath := flowlogsFolder + datePath + logFileName

		// call getFlowLogsDetails for each hour
		wg.Add(1)
		go getFlowLogsDetail(
			client,
			baseUrl,
			token,
			logFilePath,
			startTime.Unix(),
			endTime.Unix(),
			firstTs,
			limitFlag,
			hoursFlag,
			debugFlag,
			counter,
		)

	}

	wg.Wait()

}

// get flow logs from specific hour (startTs to endTs)
func getFlowLogsDetail(
	client *http.Client,
	baseUrl string,
	token string,
	logFilePath string,
	startTs int64,
	endTs int64,
	firstTs int64,
	limit int64,
	hoursFlag int64,
	debugFlag bool,
	counter int,
) {

	// create file
	file, err := os.Create(logFilePath)
	if err != nil {
		log.Fatal("Cannot create file ", err)
	}
	defer file.Close()

	groupBy := "ts,status,box,source,sourceIP,sport,device,network,destination,destinationIP,dport,domain,protocol,category,region,direction,blockType,upload,download,total,count"
	encodedGroupBy := "groupBy=" + url.QueryEscape(groupBy)

	cursor := ""

	// loop through flow logs
	for {

		fullUrl := baseUrl + "flows?query=ts%3A" + strconv.FormatInt(startTs, 10) + "-" + strconv.FormatInt(endTs, 10) + "&sortBy=ts:desc&limit=" + strconv.Itoa(int(limit)) + "&" + encodedGroupBy + "&cursor=" + url.QueryEscape(cursor)

		// get flow logs for specific hour (startTs to endTs)
		body := makeGetRequest(client, fullUrl, token)

		// convert to struct
		flowlogs := mystruct.FirewallaFlowlog{}
		json.Unmarshal([]byte(body), &flowlogs)

		// create new minTs
		minTs := endTs

		// loop through flowlog responses line by line
		for _, flowlog := range flowlogs.Results {

			if debugFlag {
				fmt.Println(" % ", flowlog)
			}
			counter += 1

			// check for lowest timestamp found
			if int64(flowlog.Ts) < minTs {
				minTs = int64(flowlog.Ts)
			}

			// get date
			datestr := time.Unix(int64(flowlog.Ts), 0).Format(time.RFC3339)
			flowlog.Date = datestr

			// write flowlog to file
			json.NewEncoder(file).Encode(flowlog)

		}

		counter += flowlogs.Count

		fmt.Println("- retrieved " + strconv.Itoa(counter) + " - " + time.Unix(minTs, 0).Format("2006-01-02 15:04:05") + " in " + logFilePath)

		// set new endTs
		endTs = minTs

		// set new cursor
		cursor = flowlogs.NextCursor

		// break if startTs is greater than endTs
		// this means we have reached the end of the flowlogs
		if cursor == "" {
			break
		}
	}

	wg.Done()
}

// main function
func main() {

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	// read config.json file
	baseUrl, token := readJsonConfig()

	// parse arguments
	modeFlag := flag.String("mode", "flowlogs", "Mode to run the script (devices, alarms, flowlogs) (default flowlogs)")
	hoursFlag := flag.Int64("hours", 24, "(Optional) Flowlog hours to retrieve (default 24)")
	limitFlag := flag.Int64("limit", 200, "(Optional) Results limit (default 200)")
	debugFlag := flag.Bool("debug", false, "Debug mode")

	// Define help flag
	helpFlag := flag.Bool("help", false, "Display help menu")
	flag.Parse()

	// store error message
	errorMsg := "Usage: ./firewalla -mode [devices|alarms|flowlogs] -hours [number] -limit [number] -debug"

	if *helpFlag {
		fmt.Println(errorMsg)
		flag.PrintDefaults()
		return
	}

	if *modeFlag == "devices" || *modeFlag == "d" {

		getDevices(client, baseUrl, token, *limitFlag)

	} else if *modeFlag == "alarms" || *modeFlag == "a" {

		getAlarms(client, baseUrl, token, *limitFlag, *hoursFlag)

	} else if *modeFlag == "flowlogs" || *modeFlag == "f" {

		getFlowLogs(client, baseUrl, token, *hoursFlag, *limitFlag, *debugFlag)

	} else {

		fmt.Println(errorMsg)
		flag.PrintDefaults()
		return

	}
}
