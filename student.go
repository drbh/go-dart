package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// TaskMap struct which contains
// an array of requirements
type Student struct {
	StudentId             string   `json:"id"`
	Name                  string   `json:"name"`
	TaskMapList           []string `json:"taskMapList"`
	ActivityLogList       []string `json:"activityLogList"`
	TaskAuditList         []string `json:"taskAuditList"`
	WhatsLeftList         []string `json:"whatsLeftList"`
	Metadata              StudentMetaData
	AuditConfig           StudentAuditConfig
	ShouldUpdateAuditFlag bool
	NumReportsRan         int `json:"numReportsRan,omitempty"`
}

type StudentMetaData struct {
	ReportRan         int
	NumMaps           int
	NumActs           int
	MainMajorComplete string
	MainMajor         string
	SIS               string
	Advisor           string
	Contact           string
}

type StudentAuditConfig struct {
	AutoRun            bool
	ExcludeMaps        []string
	MostRecentActivity bool
}

func TestStudentLoop() {
	meta := StudentMetaData{}

	config := StudentAuditConfig{
		true,
		[]string{},
		true,
	}

	student := Student{
		"75e7c8",
		"David Holtz",
		[]string{
			"KRBEMFDZDC",
			"BZRJXAWNWE",
		},
		[]string{
			"LDNJOBCSNV",
		},
		[]string{},
		[]string{},
		meta,
		config,
		true,
		0,
	}

	jsonData, _ := json.Marshal(student)
	fmt.Println(string(jsonData))

}

func GetNRandom(n int, list []string) []string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var tskLst []string
	for i := 0; i < n; i++ {
		indx := r1.Intn(len(list))
		tskLst = append(tskLst, list[indx])
	}
	// fmt.Println(tskLst)
	return tskLst
}

func MakeStudent(idf string, name string, taskList []string, actList []string) Student {
	tskLst := GetNRandom(10, taskList)
	stdLst := GetNRandom(1, actList)

	meta := StudentMetaData{}

	config := StudentAuditConfig{
		true,
		[]string{},
		true,
	}

	student := Student{
		idf,
		name,
		tskLst,
		stdLst,
		[]string{},
		[]string{},
		meta,
		config,
		true,
		0,
	}

	// fmt.Println("\n\n")
	// spew.Dump(student)
	return student
}

func MakeStudents() {
	// client.OpenBoltDb()

	students := GetStudentKeys()
	fmt.Println("\nNumber of students:\t", len(students), "\n")
	// for i := 0; i < len(students); i++ {
	// 	student := GetStudent(students[i])
	// 	fmt.Println(student.StudentId)
	// }

	tasks := GetTaskMapKeys()
	acts := GetActivityKeys()
	for i := 0; i < 1; i++ {

		s1 := rand.NewSource(time.Now().UnixNano())

		b := make([]rune, 6)
		for i := range b {
			r1 := rand.New(s1)
			b[i] = letters[r1.Intn(len(letters))]
		}
		idf := string(b)

		stu := MakeStudent(idf, randSeq(8), tasks, acts)
		// key :=
		AddStudent(stu)
		// spew.Dump(stu)
		// spew.Dump(key)
		// fmt.Println(string(key))
	}

}
