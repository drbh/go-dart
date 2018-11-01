package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	// "github.com/davecgh/go-spew/spew"
	"log"
	"os"
	"path/filepath"
)

func location() string {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	exPath = exPath + "/bolt.db"
	return exPath
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb(path string) {
	var err error
	bc.boltDB, err = bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

}

// READ FROM DB, RETURN STRINGS

func ReadTaskMap(keyString string) string {
	var result string

	key := []byte(keyString)
	client.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(task_maps)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", task_maps)
		}

		val := bucket.Get(key)

		result = string(val)
		return nil
	})

	return result
}

func ReadActivityLog(keyString string) string {
	var result string

	key := []byte(keyString)
	client.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(activity_logs)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", activity_logs)
		}

		val := bucket.Get(key)
		result = string(val)
		return nil
	})

	return result
}

func ReadStudent(keyString string) string {
	var result string

	key := []byte(keyString)
	client.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(enrolled_students)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", enrolled_students)
		}

		val := bucket.Get(key)
		result = string(val)
		return nil
	})

	return result
}

func ReadTaskAudit(keyString string) string {

	var result string

	key := []byte(keyString)
	client.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(task_audits)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", task_audits)
		}

		val := bucket.Get(key)
		result = string(val)
		return nil
	})

	return result
}

// WRITING BYTES TO DB

func WriteTaskMap(jsonData []byte) []byte {

	key := []byte(randSeq(10))
	value := jsonData

	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(task_maps)
		if err != nil {
			return err
		}

		bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	return key
}

func WriteActivityLog(jsonData []byte) []byte {

	key := []byte(randSeq(10))
	value := jsonData

	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(activity_logs)
		if err != nil {
			return err
		}

		bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	return key
}

func WriteTaskAudit(jsonData []byte) []byte {

	key := []byte(randSeq(10))
	value := jsonData

	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(task_audits)
		if err != nil {
			return err
		}

		bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	return key
}
func WriteWhatsLeftAudit(jsonData []byte) []byte {

	key := []byte(randSeq(10))
	value := jsonData

	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(whats_left_audits)
		if err != nil {
			return err
		}

		bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	return key
}

func WriteStudent(jsonData []byte) []byte {

	key := []byte(randSeq(10))
	value := jsonData

	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(enrolled_students)
		if err != nil {
			return err
		}

		bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	return key
}

// GETTING KEYS

func GetActivityKeys() []string {
	var result []string
	client.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(activity_logs)

		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			result = append(result, string(k))
		}
		return nil
	})

	return result
}

func GetTaskMapKeys() []string {
	var result []string
	client.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(task_maps)

		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			result = append(result, string(k))
		}
		return nil
	})

	return result
}

func GetTaskAuditKeys() []string {
	var result []string
	client.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(task_audits)

		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			result = append(result, string(k))
		}
		return nil
	})

	return result
}
func GetStudentKeys() []string {

	var result []string
	client.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(enrolled_students)
		if b == nil {
			return nil
		}

		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			result = append(result, string(k))
		}
		return nil
	})

	return result
}

// ADD OBJECT TO DB, MAKES OBJECT BYTES

func AddActivityLog(act ActivityLog) []byte {

	jsonData, err := json.Marshal(act)

	if err != nil {
		fmt.Println(err)
	}
	key := WriteActivityLog(jsonData)
	return key
}

func AddStudent(stu Student) []byte {

	jsonData, err := json.Marshal(stu)

	if err != nil {
		fmt.Println(err)
	}
	key := WriteStudent(jsonData)
	return key
}

func AddTaskMap(task TaskMap) []byte {

	jsonData, err := json.Marshal(task)

	if err != nil {
		fmt.Println(err)
	}
	key := WriteTaskMap(jsonData)
	return key
}

func AddTaskAudit(task TaskAudit) []byte {

	jsonData, err := json.Marshal(task)
	// fmt.Println(jsonData)

	if err != nil {
		fmt.Println(err)
	}
	key := WriteTaskAudit(jsonData)
	return key
}

func AddWhatsLeftAudit(wl WhatsLeft) []byte {

	jsonData, err := json.Marshal(wl)
	// fmt.Println(jsonData)

	if err != nil {
		fmt.Println(err)
	}
	key := WriteWhatsLeftAudit(jsonData)
	return key
}

// READ FROM DB, MAKE STRING INTO STRUCT

func GetStudent(key string) Student {
	student := ReadStudent(key)
	var stu Student

	json.Unmarshal([]byte(student), &stu)
	return stu
}

func GetTaskMap(key string) TaskMap {
	if tmCache[string(key)] == nil {
		taskMap := ReadTaskMap(key)
		var task TaskMap
		json.Unmarshal([]byte(taskMap), &task)
		tmCache[key] = &task
		return task
	}
	return *tmCache[key]
}

func GetActivityLog(key string) ActivityLog {

	activityLog := ReadActivityLog(key)
	fmt.Println(activityLog)
	var activity ActivityLog
	json.Unmarshal([]byte(activityLog), &activity)
	return activity
}

func GetTaskAudit(key string) TaskAudit {

	myTaskAudit := ReadTaskAudit(key)
	var taskAudit TaskAudit
	json.Unmarshal([]byte(myTaskAudit), &taskAudit)

	return taskAudit
}

// READ FROM DB, MAKE STRING INTO STRUCT

func UpdateActivityLog(key string, jsonData []byte) []byte {
	fmt.Println("UpdateActivityLog")
	value := jsonData
	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(activity_logs)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", activity_logs)
		}
		bucket.Put([]byte(key), value)
		return nil
	})
	return []byte(key)
}

func UpdateStudent(key string, jsonData []byte) []byte {
	fmt.Println("UpdateStudent")
	value := jsonData
	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(enrolled_students)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", enrolled_students)
		}
		bucket.Put([]byte(key), value)
		return nil
	})
	return []byte(key)
}

// activity log ID
// class
// AddCoreToActivityLog( "TCUAXHXKQF", "CON 101" )
func AddCoreToActivityLog(activityLogKey string, className string) []byte {
	fmt.Println("AddCoreToActivityLog")
	al := GetActivityLog(activityLogKey)
	al.Completed[className] = true
	jsond, _ := json.Marshal(al)
	key := UpdateActivityLog(activityLogKey, jsond)
	return key
}

// student ID
// class
// AddCoreToStudentActivityLog( "MRAJWWHTHC", "CON 101" )
func AddCoreToStudentActivityLog(studentKey string, className string) []byte {
	fmt.Println("AddCoreToStudentActivityLog")
	// fmt.Println(studentKey)
	student := GetStudent(studentKey)

	student.ShouldUpdateAuditFlag = true
	activityLogKey := student.ActivityLogList[0]
	key := AddCoreToActivityLog(activityLogKey, className)

	studentBytes, _ := json.Marshal(student)
	UpdateStudent(studentKey, studentBytes)
	return key
}

// student ID
// major ID
// AddCoreToStudentActivityLog( "MRAJWWHTHC", "CON 101" )
func AddMajorToStudent(studentKey string, majorID string) []byte {
	fmt.Println("AddMajorToStudent")
	// fmt.Println(studentKey)
	student := GetStudent(studentKey)

	student.ShouldUpdateAuditFlag = true
	student.TaskMapList = append(student.TaskMapList, majorID)
	// activityLogKey := student.ActivityLogList[0]
	// key := AddCoreToActivityLog(activityLogKey, className)

	studentBytes, _ := json.Marshal(student)
	UpdateStudent(studentKey, studentBytes)
	return []byte(studentKey)
}

// WRITING BYTES TO DB

func WriteStudentSummary(key string, jsonData []byte) []byte {
	mykey := []byte(key)
	value := jsonData
	client.boltDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(student_summary)
		if err != nil {
			return err
		}

		bucket.Put(mykey, value)
		if err != nil {
			return err
		}
		return nil
	})
	return mykey
}

func ReadStudentSummary(keyString string) string {
	var result string

	key := []byte(keyString)
	client.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(student_summary)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", student_summary)
		}

		val := bucket.Get(key)
		result = string(val)
		return nil
	})

	return result
}

func GetStudentSummary(key string) SutdentSummary {
	student := ReadStudentSummary(key)
	var stuSum SutdentSummary
	json.Unmarshal([]byte(student), &stuSum)
	return stuSum
}

func UpsertNewStudentSummary(key string, currentProgress ProgressSummary) {
	student := GetStudentSummary(key)

	// If no summary than we neeed to make one
	// fmt.Println(currentProgress)
	// fmt.Println(student)

	// if len(student.Summaries) == 0 {
	// 	fmt.Println("First Student Summary", currentProgress.LastUpdate)

	// }

	// // else {
	// // 	// if there is a summary we need to update any old entries
	// fmt.Println(len(student.Summaries))

	// // fmt.Println("\nAdd New Progress Summary to, Student Summary")
	// // student.Summaries = []ProgressSummary{currentProgress} // append(student.Summaries, currentProgress)
	// spew.Dump(currentProgress)
	justAppend := true
	for i := 0; i < len(student.Summaries); i++ {
		stusumProg := student.Summaries[i]
		// if they are the same major and its newer
		if (stusumProg.LastUpdate) < currentProgress.LastUpdate && stusumProg.Major == currentProgress.Major {
			// we update!
			fmt.Println("\nUpdate, Student Summary", currentProgress.Major, currentProgress.LastUpdate)
			// student.Summaries = append(student.Summaries, currentProgress)
			// should replace this stusumProg with currentProgress
			student.Summaries[i] = currentProgress
			justAppend = false
			// spew.Dump(student)
		}
	}

	if justAppend {
		student.Summaries = append(student.Summaries, currentProgress)
	}
	// // }
	// fmt.Println("\n Writing, Student Summary")
	// // spew.Dump(student)

	jsond, _ := json.Marshal(student)
	WriteStudentSummary(key, jsond)
	return
}
