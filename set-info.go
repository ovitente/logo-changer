package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// ChooseInfo Reads elements list and choosing one pair
func ChooseInfo() (string, string, string, string) {

	type PairsList struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	}

	logoType := "info"
	elementsFilePath := "./list-info"
	elementsList := make([]PairsList, 0)

	elementsFile, err := os.Open(elementsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	yamlByte, err := ioutil.ReadAll(elementsFile)
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(yamlByte, &elementsList); err != nil {
		log.Fatal(err)
	}

	elementsNumber := random(0, len(elementsList))
	listElement := elementsList[elementsNumber]

	// fmt.Println(logoType, logoType, listElement.Name, listElement.URL)
	// fmt.Printf("INFO CHOSEN RANDOM NUMBER: %v\n", elementsNumber)
	// fmt.Printf("INFO MAP SIZE: %v\n", len(elementsList))

	return logoType, logoType, listElement.Name, listElement.URL
}
