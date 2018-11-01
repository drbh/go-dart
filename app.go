// Real Time Degree Auditing.
package main

// Test.
import (
	"fmt"
	"gopkg.in/olahol/melody.v1"
	"log"
	"os"
	"path/filepath"
)

const (
	// OptionsFirst pass this setting if options should be passed as first
	// args.
	str string = "to"
)

// This is a test stucture
type Test struct {
	Yes string `json:"name"`
}

var tmCache = map[string]*TaskMap{}
var client BoltClient
var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var integers = []rune("123456789")
var task_maps = []byte("task_maps")
var activity_logs = []byte("activity_logs")
var task_audits = []byte("task_audits")
var whats_left_audits = []byte("whats_left_audits")
var enrolled_students = []byte("enrolled_students")
var student_summary = []byte("student_summary")
var databasePath = string(location())

var globalSocketHandler *melody.Melody

// go-swagger examples.
//
// The purpose of this application is to provide some
// use cases describing how to generate docs for your API
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("\n")
	databasePath := dir + "/bolt.db"
	client.OpenBoltDb(databasePath)

	StartRealTimeWatcher()
	fmt.Println("Start Server \n\n\n\n")
	StartServer()

}
