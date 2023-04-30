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
	"strconv"
	"sync"
	"time"

	mystruct "firewalla/mystructs"
)

// create global counter and waitgroup
var counter int
var wg sync.WaitGroup

// read config.json with url and token
func readJsonConfig() (string, string) {

	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	// read the token variable from the config.json file
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

	//client := &http.Client{}

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

func makePostRequest(url string, token string, startTs float64, endTs float64) []byte {

	postData := map[string]float64{
		"start": startTs,
		"end":   endTs,
	}
	postBody, err := json.Marshal(postData)

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

	client := &http.Client{}
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

	// Create file
	f, err := os.Create("devices.json")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	body := makeGetRequest(client, url+"device/list", token)
	devices := mystruct.FirewallaDevices{}
	json.Unmarshal([]byte(body), &devices)

	wg = sync.WaitGroup{}

	for _, device := range devices {

		wg.Add(1)
		go getDeviceDetail(client, url, token, device.Gid, device.Mac, f)

	}

	wg.Wait()

	fmt.Println("* devices saved to devices.json (", counter, ")")

}

// get device detail
func getDeviceDetail(client *http.Client, url string, token string, gid string, mac string, f *os.File) {

	body := makeGetRequest(client, url+"device/"+gid+"/"+mac, token)

	deviceDetail := mystruct.FirewallaDeviceDetail{}
	json.Unmarshal([]byte(body), &deviceDetail)

	// Write all fields as a JSON string to file
	jsonString, _ := json.Marshal(deviceDetail)
	f.Write([]byte(string(jsonString) + "\n"))

	counter++

	if counter%5 == 0 {

		fmt.Println("- ", counter, " devices completed")

	}

	wg.Done()

}

// get alarms
func getAlarms(client *http.Client, url, token string) {

	// Reset global counter
	counter = 0

	// Create file
	file, err := os.Create("alarms.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := makeGetRequest(client, url+"alarm/list", token)
	alarms := mystruct.FirewallaAlarms{}
	json.Unmarshal([]byte(body), &alarms)

	fmt.Println("* alarms started")

	for _, alarm := range alarms {

		aid := alarm.Aid
		gid := alarm.Gid

		wg.Add(1)

		go getAlarmDetail(client, url, token, aid, gid, file)

	}

	wg.Wait()

	fmt.Println("* alarms saved to alarms.json (", counter, ")")
}

// get alarm detail
func getAlarmDetail(client *http.Client, url string, token string, aid string, gid string, f *os.File) {

	body := makeGetRequest(client, url+"alarm/"+gid+"/"+aid, token)

	alarmDetail := mystruct.FirewallaAlarmDetail{}
	json.Unmarshal([]byte(body), &alarmDetail)

	// Write all fields as a JSON string to file
	jsonString, _ := json.Marshal(alarmDetail)
	f.Write([]byte(string(jsonString) + "\n"))

	counter++

	if counter%5 == 0 {
		fmt.Println("- ", counter, "alarms completed")
	}

	wg.Done()
}

// get flow logs
func getFlowLogs(client *http.Client, url string, token string, hours int64) {

	// Reset global counter
	counter = 0

	// Create file
	file, err := os.Create("flowlogs.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// get flow logs from last X hours
	startTs := float64(time.Now().Unix() - hours*60*60)
	endTs := float64(time.Now().Unix())

	fmt.Println("* flowlogs started - get last " + strconv.Itoa(int(hours)) + " hours")

	for {

		// get flow logs
		body := makePostRequest(
			url+"flows/query",
			token,
			startTs,
			endTs,
		)

		// convert to struct
		flowlogs := mystruct.FirewallaFlowlogDetail{}
		json.Unmarshal([]byte(body), &flowlogs)

		// create new minTs
		minTs := endTs

		// loop through flowlogs
		for _, flowlog := range flowlogs {

			// check for lowest timestamp
			if flowlog.Ts < minTs {
				minTs = flowlog.Ts
			}

			// write all fields as a JSON string to file
			jsonString, _ := json.Marshal(flowlog)

			// add date string to JSON
			jsonDateString := `{"date":"` + time.Unix(int64(flowlog.Ts), 0).Format("2006-01-02 15:04:05") + `",`
			jsonString = []byte(jsonDateString + string(jsonString[1:]))

			// write to file
			file.Write([]byte(string(jsonString) + "\n"))

			// increment counter
			counter++

		}

		// calculate percentage done based on startTs and endTs
		diff := endTs - startTs
		percentage := 100 - int((diff/float64(hours*60*60))*100)

		fmt.Println("- flows", percentage, "% completed")

		// set new endTs
		endTs = minTs

		// break if startTs is greater than endTs
		// this means we have reached the end of the flowlogs
		if len(flowlogs) != 200 || startTs > endTs {
			break
		}
	}

	fmt.Println("* flowlogs saved to flowlogs.json (", counter, ")")
}

// main function
func main() {

	client := &http.Client{}

	// read config.json file
	url, token := readJsonConfig()

	// parse arguments
	modeFlag := flag.String("mode", "", "Mode to run the script (devices, alarms, or flowlogs)")
	hoursFlag := flag.Int("hours", 1, "Flowlog hours to retrieve")

	// Define help flag
	helpFlag := flag.Bool("help", false, "Display help menu")

	flag.Parse()

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
		getFlowLogs(client, url, token, int64(*hoursFlag))
	}
}
