package FunctionsForParser

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

type Logs struct {
	Date string `json:"date"`
	User string `json:"user"`
	Command string `json:"command"`
	TypeOurLog string `json:"type"`
}

func WriterInJson(logItems ...string) {
	if len(logItems) != 4 {
		fmt.Println("Неверное количество аргументов для logItems")
		return
	}

	log := []Logs{
		{
			Date: logItems[0],
			User: logItems[1],
			Command: logItems[2],
			TypeOurLog: logItems[3],
		},
	}

	fileName := "New-logs.json"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {

		jsonLog, err := json.MarshalIndent(log, "", " ")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = ioutil.WriteFile(fileName, jsonLog, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("File created")
		return
	}

	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("FILE ISN'T WORKING", err)
		return
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("He can't read file", err)
		return
	}

	var existingLogs []Logs
	err = json.Unmarshal(data, &existingLogs)
	if err != nil {
		fmt.Println("JsonMap", err)
		return
	}

	existingLogs = append(existingLogs, log...)

	newData, err := json.MarshalIndent(existingLogs, "", " ")
	if err != nil {
	fmt.Println(err)
	return
	}

	err = ioutil.WriteFile(fileName, newData, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}