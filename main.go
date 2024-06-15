package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"Parser/FunctionsForParser"
	"github.com/spf13/viper"
)


func Parser(log []string, LastSuperUser string)(string) {
	count := 0
	LogString := ""
	for count < len(log) {
		LogString = LogString + " " + log[count]
		count ++
	}
	if strings.Contains(LogString, viper.GetString("KeyForLog_1")) && strings.Contains(LogString, viper.GetString("KeyForLog_2")) {
		if log[5] == viper.GetString("KeyForLog_3"){
			FunctionsForParser.WriterInJson(FunctionsForParser.TimeConstructor(FunctionsForParser.NewTime(log[0])), log[10]+"(localhost)", 
                        log[2] + log[3] + log[4] + " " + log[5], viper.GetString("TypeLog_1"))
			LastSuperUser = log[10]
		} else if log[5] == viper.GetString("KeyForLog_4") {
			FunctionsForParser.WriterInJson(FunctionsForParser.TimeConstructor(FunctionsForParser.NewTime(log[0])), LastSuperUser+"(localhost)", 
                        log[2] + log[3] + log[4]+ " " + log[5], viper.GetString("TypeLog_2"))
		}
	} else if strings.Contains(LogString, viper.GetString("KeyForLog_5")) {
		FunctionsForParser.WriterInJson(FunctionsForParser.TimeConstructor(FunctionsForParser.NewTime(log[0])), log[3]+"(localhost)", log[11], viper.GetString("TypeLog_3"))
	} else if strings.Contains(LogString, viper.GetString("KeyForLog_6")) {
		if strings.Contains(LogString, viper.GetString("KeyForLog_3")) {
			FunctionsForParser.WriterInJson(FunctionsForParser.TimeConstructor(FunctionsForParser.NewTime(log[0])), log[8]+"(localhost)", log[3] + log[4] + " " + log[5], viper.GetString("TypeLog_4"))
		} else if strings.Contains(LogString, viper.GetString("KeyForLog_4")) {
			FunctionsForParser.WriterInJson(FunctionsForParser.TimeConstructor(FunctionsForParser.NewTime(log[0])), log[8]+"(localhost)", log[3] + log[4] + " " + log[5], viper.GetString("TypeLog_5"))
		}
	} else if strings.Contains(LogString, viper.GetString("KeyForLog_7")) && strings.Contains(LogString, viper.GetString("KeyForLog_8")) {
		FunctionsForParser.WriterInJson(FunctionsForParser.TimeConstructor(FunctionsForParser.NewTime(log[0])), log[8]+"(localhost)", log[2]+log[3]+" "+log[4]+" "+log[5], viper.GetString("TypeLog_6"))
	}
	return LastSuperUser
}

func  main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error sdgsg %s", err)
	}
	value := viper.GetString("key")
	fmt.Println(value)
	var LogTime[14]int
	for {
		file, err := os.Open("/var/log/auth.log")
		if err != nil{
			fmt.Println(err)
			log.Fatal(err)
			os.Exit(1)
		}
		
		defer file.Close()

		scanner := bufio.NewScanner(file)

		LastSuperUser := ""

		for scanner.Scan() {
			data := strings.Fields(scanner.Text())
			if data[0][10] != 0 {
				var OurTime [7]int
				key := false
				LogTime[0], LogTime[1], LogTime[2], LogTime[3], LogTime[4], LogTime[5], LogTime[6] = FunctionsForParser.NewTime(data[0])
				OurTime, key = FunctionsForParser.WhatTimeItIsNow(LogTime)
				LogTime[7], LogTime[8], LogTime[9],
				LogTime[10], LogTime[11], LogTime[12], LogTime[13] = OurTime[0], OurTime[1], OurTime[2],
				OurTime[3], OurTime[4], OurTime[5], OurTime[6]
				if key {
					LastSuperUser = Parser(data, LastSuperUser)
				}
			}
		}

		if err := scanner.Err(); err!= nil {
			log.Fatal(err)
		}

	}
}
