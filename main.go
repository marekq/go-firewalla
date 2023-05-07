package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	mystruct "firewalla/mystructs"
)

// define files and folders
var configFile = "config.json"
var devicesFile = "devices.json"
var alarmsFile = "alarms.json"
var flowlogsFolder = "flowlogs/"

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
		log.Fatal(err)
	}

	// set headers
	req.Header.Set("Authorization", token)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("Status code error:", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return body

}

// make post request to api, return body
func makePostRequest(url string, token string, startTs float64, endTs float64, client *http.Client) []byte {

	postData := map[string]float64{
		"start": startTs,
		"end":   endTs,
	}

	postBody, err := json.MarshalIndent(postData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatal(err, req)
	}

	// set headers
	req.Header.Set("Authorization", token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err, resp)
	}

	if resp.StatusCode != 200 {
		log.Fatal("Status code error:", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()

	// convert to []byte
	respBody, _ := io.ReadAll(resp.Body)

	return respBody
}

// get devices
func getDevices(client *http.Client, url string, token string) {

	// Reset global counter
	counter = 0

	fmt.Println("* devices started")
	wg = sync.WaitGroup{}

	// Create file
	file, err := os.Create(devicesFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := makeGetRequest(client, url+"device/list", token)
	devices := mystruct.FirewallaDevices{}
	json.Unmarshal([]byte(body), &devices)

	for _, device := range devices {

		wg.Add(1)
		go getDeviceDetail(client, url, token, device.Gid, device.Mac, file)

	}

	wg.Wait()

	fmt.Println("* completed -", counter, "devices saved to devices.json")
}

// get device detail
func getDeviceDetail(client *http.Client, url string, token string, gid string, mac string, file *os.File) {

	body := makeGetRequest(client, url+"device/"+gid+"/"+mac, token)

	deviceDetail := mystruct.FirewallaDeviceDetail{}
	json.Unmarshal([]byte(body), &deviceDetail)

	// convert lastActive float64 to datestring str
	ts, err := strconv.ParseFloat(deviceDetail.LastActive, 64)
	if err != nil {
		log.Fatal(err)
	}
	datestr := time.Unix(int64(ts), 0).Format("2006-01-02 15:04:05")
	deviceDetail.Date = datestr

	counter++

	if counter%5 == 0 {
		fmt.Println("-", counter, "devices completed")
	}

	// write deviceDetail to file
	json.NewEncoder(file).Encode(deviceDetail)

	wg.Done()
}

// get alarms
func getAlarms(client *http.Client, url, token string) {

	// Reset global counter
	counter = 0

	// Create file
	file, err := os.Create(alarmsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := makeGetRequest(client, url+"alarm/list", token)
	alarms := mystruct.FirewallaAlarms{}
	json.Unmarshal([]byte(body), &alarms)

	fmt.Println("* alarms started")
	wg = sync.WaitGroup{}

	for _, alarm := range alarms {

		wg.Add(1)
		go getAlarmDetail(client, url, token, alarm.Aid, alarm.Gid, file)

	}

	wg.Wait()

	fmt.Println("* completed -", counter, "alarms saved to alarms.json")
}

// get alarm detail
func getAlarmDetail(client *http.Client, url string, token string, aid string, gid string, file *os.File) {

	body := makeGetRequest(client, url+"alarm/"+gid+"/"+aid, token)

	alarmDetail := mystruct.FirewallaAlarmDetail{}
	json.Unmarshal([]byte(body), &alarmDetail)

	// convert float64 to int
	ts := int64(alarmDetail.Timestamp)
	datestr := time.Unix(ts, 0).Format("2006-01-02 15:04:05")
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
					log.Fatal(err)
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
		fmt.Println("! error walking through the directory:", err)

	}

	if newestTimestamp.IsZero() {
		fmt.Println("! no flowlog folder found, using ", hours, " hours")

		// return hours ago
		return time.Now().Add(-time.Duration(hours) * time.Hour)

	}

	fmt.Println("* flowlogs starting with timestamp:", newestTimestamp.Format("2006-01-02 15:04"))

	// return newestTimestamp
	return newestTimestamp
}

// get flowlogs
func getFlowLogs(client *http.Client, url string, token string, hours int64) {

	// reset global counter
	counter = 0
	wg = sync.WaitGroup{}

	// create firstTs and lastTs
	firstTs := getFlowlogsFirstTs(hours).Unix()
	lastTs := time.Now().Unix()

	// round down the firstTs to the nearest full hour (i.e. 15:30 -> 15:00)
	firstTs -= firstTs % 3600

	// Loop through all hours
	for currentTs := firstTs; currentTs < lastTs; currentTs += 3600 {

		// start and end time
		startTime := time.Unix(currentTs, 0)
		endTime := time.Unix(currentTs+3600, 0)

		// date path: dd/mm/yyyy/
		datePath := fmt.Sprintf("%d/%02d/%02d/", startTime.Year(), startTime.Month(), startTime.Day())
		err := os.MkdirAll(flowlogsFolder+datePath, 0755)
		if err != nil {
			fmt.Println("Error creating folders:", err)
			return
		}

		// file name: flowlogs_YYYYMMDDTHHMM_to_YYYYMMDDTHHMM.csv
		logFileName := fmt.Sprintf("%s_%d.csv", startTime.Format("2006-01-02_15-04"), currentTs)

		// full path: flowlogs/2020/01/01/20200101T0000.csv
		logFilePath := flowlogsFolder + datePath + logFileName

		// call getFlowLogsDetails for each hour
		wg.Add(1)
		go getFlowLogsDetail(client, url, token, logFilePath, float64(startTime.Unix()), float64(endTime.Unix()), float64(firstTs))

	}

	wg.Wait()

	fmt.Println("* completed", counter, "flowlogs, saved to flowlogs/ folder")

}

// get flow logs from specific hour (startTs to endTs)
func getFlowLogsDetail(client *http.Client, url string, token string, logFilePath string, startTs float64, endTs float64, firstTs float64) {

	flowLogCounter := 0
	startDateString := time.Unix(int64(startTs), 0).Format("2006-01-02 15:04")

	// create file
	file, err := os.Create(logFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// loop through flow logs
	for {

		// get flow logs for specific hour (startTs to endTs)
		body := makePostRequest(
			url+"flows/query",
			token,
			startTs,
			endTs,
			client,
		)

		// convert to struct
		flowlogs := []mystruct.FirewallaFlowlogDetail{}
		json.Unmarshal([]byte(body), &flowlogs)

		// create new minTs
		minTs := endTs

		// loop through flowlog responses line by line
		for _, flowlog := range flowlogs {

			flowLogCounter += 1

			// check for lowest timestamp found
			if flowlog.Ts < minTs {
				minTs = flowlog.Ts
			}

			// get date
			datestr := time.Unix(int64(flowlog.Ts), 0).Format("2006-01-02 15:04:05")
			flowlog.Date = datestr

			// write flowlog to file
			json.NewEncoder(file).Encode(flowlog)

			// increment counter
			counter++

		}

		// set new endTs
		endTs = minTs

		// break if startTs is greater than endTs
		// this means we have reached the end of the flowlogs
		if len(flowlogs) != 200 || startTs > endTs {
			break
		}
	}

	fmt.Printf("- %d flowlogs for %s\n", flowLogCounter, startDateString)
	wg.Done()
}

// main function
func main() {

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	// read config.json file
	url, token := readJsonConfig()

	// parse arguments
	modeFlag := flag.String("mode", "", "Mode to run the script (devices, alarms, or flowlogs)")
	hoursFlag := flag.Int64("hours", 24, "Flowlog hours to retrieve")

	// Define help flag
	helpFlag := flag.Bool("help", false, "Display help menu")
	flag.Parse()

	// store error message
	errorMsg := "Usage: ./firewalla -mode [devices|alarms|flowlogs -hours [number]]"

	if *helpFlag {
		fmt.Println(errorMsg)
		flag.PrintDefaults()
		return
	}

	if *modeFlag == "devices" || *modeFlag == "d" {

		getDevices(client, url, token)

	} else if *modeFlag == "alarms" || *modeFlag == "a" {

		getAlarms(client, url, token)

	} else if *modeFlag == "flowlogs" || *modeFlag == "f" {

		getFlowLogs(client, url, token, *hoursFlag)

	}
}
