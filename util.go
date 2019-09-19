package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"
)

func tmplToString(tmplString string, data interface{}) (string, error) {
	tmpl, parseErr := template.New("tmpl").Parse(tmplString)
	if parseErr != nil {
		return "", parseErr
	}
	buffer := new(bytes.Buffer)
	executeErr := tmpl.Execute(buffer, data)
	return string(buffer.Bytes()), executeErr
}

func fetch(method string, path string, data map[string]interface{}) ([]byte, error) {
	url := cfg.Coolq.API + path
	dataByte, _ := json.Marshal(data)
	reader := bytes.NewReader(dataByte)
	client := &http.Client{}
	request, _ := http.NewRequest(method, url, reader)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+cfg.Coolq.Token)
	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}
	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return body, err
}

func nextTime(unix int64) time.Time {
	now := time.Unix(unix, 0)
	var newHour, newMinute int
	if now.Minute() >= 30 {
		newHour = now.Hour() + 1
		newMinute = 0
	} else {
		newHour = now.Hour()
		newMinute = 30
	}
	var next time.Time
	next = time.Date(now.Year(), now.Month(), now.Day(), newHour, newMinute, 0, 0, time.Local)
	if newHour == 6 {
		next = next.Add(time.Hour * 18)
	}
	return next
}
