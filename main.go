package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	mystruct "firewalla/mystructs"
)

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

// make request to api, return body
func makeRequest(url, token string) []byte {

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	if err != nil {
		log.Fatal(err)
	}

	// set headers
	req.Header.Set("Authorization", token)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

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

// get devices
func getDevices(url, token string) {

	fmt.Println("* devices started")

	// Create file
	f, err := os.Create("devices.json")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	body := makeRequest(url+"device/list", token)
	devices := mystruct.FirewallaDevices{}
	json.Unmarshal([]byte(body), &devices)

	wg = sync.WaitGroup{}

	for _, device := range devices {

		wg.Add(1)
		go getDeviceDetail(url, token, device.Gid, device.Mac, f)

	}

	wg.Wait()

	fmt.Println("- devices completed", len(devices))

}

// get device detail
func getDeviceDetail(url string, token string, gid string, mac string, f *os.File) {

	body := makeRequest(url+"device/"+gid+"/"+mac, token)

	deviceDetail := mystruct.FirewallaDeviceDetail{}
	json.Unmarshal([]byte(body), &deviceDetail)

	// Write all fields as a JSON string to file
	jsonString, _ := json.Marshal(deviceDetail)
	f.Write([]byte(string(jsonString) + "\n"))

	counter++

	if counter%5 == 0 {

		fmt.Println("* devices completed", counter)

	}

	wg.Done()

}

// get alarms
func getAlarms(url, token string) {

	// Create file
	file, err := os.Create("alarms.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := makeRequest(url+"alarm/list", token)
	alarms := mystruct.FirewallaAlarms{}
	json.Unmarshal([]byte(body), &alarms)

	fmt.Println("* alarms started")

	for _, alarm := range alarms {

		aid := alarm.Aid
		gid := alarm.Gid

		wg.Add(1)

		go getAlarmDetail(url, token, aid, gid, file)

	}

	wg.Wait()

	fmt.Println("- alarms completed:", counter)
}

// get alarm detail
func getAlarmDetail(url string, token string, aid string, gid string, f *os.File) {

	body := makeRequest(url+"alarm/"+gid+"/"+aid, token)

	alarmDetail := mystruct.FirewallaAlarmDetail{}
	json.Unmarshal([]byte(body), &alarmDetail)

	// Write all fields as a JSON string to file
	jsonString, _ := json.Marshal(alarmDetail)
	f.Write([]byte(string(jsonString) + "\n"))

	counter++

	if counter%5 == 0 {
		fmt.Println("* alarms completed:", counter)
	}

	wg.Done()
}

// main function
func main() {

	url, token := readJsonConfig()

	// parse arguments
	args := os.Args[1:]

	if len(args) > 0 {
		if args[0] == "devices" || args[0] == "d" {
			getDevices(url, token)

		} else if args[0] == "alarms" || args[0] == "a" {
			getAlarms(url, token)

		} else if args[0] == "all" {
			getDevices(url, token)
			getAlarms(url, token)

		} else {
			fmt.Println("! please specify devices, alarms or all on the command line, quitting")

		}

	} else {
		fmt.Println("! please specify devices, alarms or all on the command line, quitting")
	}
}
