package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	mystruct "firewalla/mystructs"
)

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
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", resp.Body)
	body, _ := io.ReadAll(resp.Body)

	return body

}

// get devices
func getDevices(url, token string) {

	// Create file
	f, err := os.Create("devices.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	body := makeRequest(url+"device/list", token)
	devices := mystruct.FirewallaDevices{}
	json.Unmarshal([]byte(body), &devices)

	for _, device := range devices {

		timestamp, err := strconv.ParseFloat(device.LastActive, 64)
		if err != nil {
			fmt.Println("Error parsing string:", err)
		}

		sec := int64(timestamp)
		t := time.Unix(sec, int64((timestamp)))

		formatted := t.Format("2006-01-02 15:04")
		fmt.Println("DEVICE", formatted, device.Name, device.IP, device.MacVendor)

		// Write all fields as a JSON string to file
		jsonString, _ := json.Marshal(device)
		f.Write([]byte(formatted + " " + string(jsonString) + "\n"))
	}

}

// get alarms
func getAlarms(url, token string) {

	// Create file
	f, err := os.Create("alarms.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	body := makeRequest(url+"alarm/list", token)
	alarms := mystruct.FirewallaAlarms{}
	json.Unmarshal([]byte(body), &alarms)

	fmt.Println(alarms)

	for _, alarm := range alarms {

		sec := int64(alarm.Timestamp)
		t := time.Unix(sec, int64((alarm.Timestamp)))

		formatted := t.Format("2006-01-02 15:04")
		fmt.Println("ALARM", formatted, alarm.Device, alarm.Type, alarm.Message)

		// Write all fields as a JSON string to file
		jsonString, _ := json.Marshal(alarm)
		f.Write([]byte(formatted + " " + string(jsonString) + "\n"))

	}
}

// main function
func main() {

	url, token := readJsonConfig()

	getDevices(url, token)
	getAlarms(url, token)

}
