package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// TaskMap struct which contains
// an array of requirements
type TaskMap struct {
	Name         string        `json:"name"`
	Requirements []Requirement `json:"requirements"`
}

type Requirement struct {
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Group      string      `json:"group"`
	Satisfiers []Satisfier `json:"satisfiers"`
}

type Satisfier struct {
	Uuid  string   `json:"uuid"`
	Count int      `json:"count"`
	Core  []string `json:"core"`
}

func readTaskJSON(fname string) TaskMap {
	// Open our jsonFile
	jsonFile, err := os.Open(fname)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our task struct
	var task TaskMap

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &task)

	return task
}
