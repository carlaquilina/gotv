package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"gotv.com/models"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Print(cfg)
	//get config
	//login and set object to make requests
	//make a request to get all movies
	//check when movies will expire
	//return html (use templates) to show list of movies expiry in ascending order
}

func getConfig() (models.Config, error) {
	f, err := os.Open("config.yml")
	if err != nil {
		return models.Config{}, err
	}
	defer f.Close()

	var cfg models.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return models.Config{}, err
	}
	return cfg, nil
}
