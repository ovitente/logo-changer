package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// ChooseVideo Reads video list and choosing one pair
func ChooseVideo() (string, string, string, string) {

	type PairsList struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	}

	logoType := "videos"
	var elementsFilePath = "./list-videos"
	var elementsList = make([]PairsList, 0)

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

	elementNumber := random(0, len(elementsList)-1)
	listElement := elementsList[elementNumber]

	// fmt.Printf("VIDEO MAP SIZE: %v\n", len(elementsList))
	// fmt.Println(logoType, logoType, listElement.Name, listElement.URL)

	return logoType, logoType, listElement.Name, listElement.URL
}
