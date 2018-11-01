package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ActivityLog struct {
	Name      string          `json:"name"`
	Completed map[string]bool `json:"completed"`
}

// type Instance struct {
// 	Ucid  string `json:"ucid"`
// 	Grade int    `json:"grade"`
// }

func readPerson(fname string) ActivityLog {
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

	// we initialize our person array
	var person ActivityLog

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &person)
	// fmt.Println(person)
	return person
}
