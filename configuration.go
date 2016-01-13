package main

import (
  "encoding/json"
  "os"
  "fmt"
)

type configuration struct {
  Database map[string]string
}

func constructApplicationConf() configuration {
  file, _ := os.Open("configuration.json")
  decoder := json.NewDecoder(file)
  confHash := configuration{}
  err := decoder.Decode(&confHash)

  if err != nil {
    panic(fmt.Sprintf("Cannot read configuration file, reason: %s!", err))
  }
  connectString := "sslmode=disable"
  keysHash := map[string]string{
    "NameKey" : "dbname",
    "UserNameKey" : "user",
    "HostKey" : "host",
  }

  for key, value := range keysHash {
    envValue := os.Getenv(confHash.Database[key])
    if len(envValue) > 0 {
      connectString += (" " + value + "=" + envValue)
    }
  }

  confHash.Database["connectString"] =  connectString
  return confHash
}
// ApplicationConf - Appplication configuration for board
var ApplicationConf = constructApplicationConf()
