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
	var videosFilePath = "./videos.list"
	var videosList = make([]PairsList, 0)

	videosFile, err := os.Open(videosFilePath)
	if err != nil {
		log.Fatal(err)
	}

	yamlByte, err := ioutil.ReadAll(videosFile)
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(yamlByte, &videosList); err != nil {
		log.Fatal(err)
	}

	videoNumber := random(0, len(videosList)-1)
	listElement := videosList[videoNumber]

	// fmt.Printf("VIDEO MAP SIZE: %v\n", len(videosList))
	// fmt.Println(logoType, logoType, listElement.Name, listElement.URL)

	return logoType, logoType, listElement.Name, listElement.URL
}
