package main

import (
	"crypto/rand"
	"fmt"
	"log"
	// "math/rand"
	"time"
)

func pseudo_uuid() (uuid string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}

func randSeq(n int) string {
	// _ := n
	// b := make([]rune, n)
	// for i := range b {
	// 	b[i] = letters[rand.Intn(len(letters))]
	// }
	return pseudo_uuid() //string(b)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	if elapsed == elapsed {

	}
	log.Printf("%s took %s", name, elapsed)
}

// audit no new
func AuditProcess(ltid []byte, laids []byte) *TaskAudit {
	task := GetTaskMap(string(ltid))
	activityLog := GetActivityLog(string(laids))
	audit, _ := RunAudit(task, activityLog)
	return audit
}
