package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// ChooseInfo Reads info list and choosing one pair
func ChooseInfo() (string, string, string, string) {

	type PairsList struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	}

	logoType := "info"
	infoFilePath := "./info.list"
	infoList := make([]PairsList, 0)

	infoFile, err := os.Open(infoFilePath)
	if err != nil {
		log.Fatal(err)
	}

	yamlByte, err := ioutil.ReadAll(infoFile)
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(yamlByte, &infoList); err != nil {
		log.Fatal(err)
	}

	infoNumber := random(0, len(infoList))
	listElement := infoList[infoNumber]

	// fmt.Println(logoType, logoType, listElement.Name, listElement.URL)
	// fmt.Printf("INFO CHOSEN RANDOM NUMBER: %v\n", infoNumber)
	// fmt.Printf("INFO MAP SIZE: %v\n", len(infoList))

	return logoType, logoType, listElement.Name, listElement.URL
}
