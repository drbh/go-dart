package main

import (
	"encoding/json"
	"fmt"
	// "reflect"
	"github.com/boltdb/bolt"
	// "log"
	// "strconv"
	"sync"
	"time"
)

func schedule(what func(), delay time.Duration) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

type RealTimeWatcher struct {
	StartTime  time.Time
	Pluse      time.Duration
	IsWatching chan bool
}

func StartRealTimeWatcher() {

	rt := RealTimeWatcher{
		time.Now(),
		(1 * 500),
		make(chan bool),
	}
	fmt.Println(client.boltDB)
	ping := func() {
		beg := time.Now()
		didProcess := false
		students := GetStudentKeys()

		// loop over all students
		// fmt.Println("Num Students: ", students)

		// slice := []string{"a", "b", "c", "d", "e"}
		// sliceLength := len(slice)

		// for i := 0; i < len(students); i++ {

		var wg sync.WaitGroup
		wg.Add(len(students))
		// fmt.Println("Running for loop…")

		for i := 0; i < len(students); i++ {
			go func(i int) {
				defer wg.Done()
				// msg := fmt.Sprintf("%4d", time.Now().UTC().Unix()) + ": Heartbeat"
				// globalSocketHandler.Broadcast([]byte(msg))
				// val := students[i]
				// fmt.Printf("i: %v, val: %v\n", i, val)

				student := GetStudent(students[i])

				// if student has update flag
				if student.ShouldUpdateAuditFlag {
					// fmt.Println("Should Update!")
					// fmt.Println(student.Name)
					fmt.Println("Got:", students[i])

					studentkey := students[i]
					// check all thier enrolled tasks
					for i := 0; i < len(student.TaskMapList); i++ {
						tmap := student.TaskMapList[i]
						taskMap := GetTaskMap(tmap)

						// run against each activity
						for j := 0; j < len(student.ActivityLogList); j++ {
							alog := student.ActivityLogList[j]
							activityLog := GetActivityLog(alog)
							audit, wl := RunAudit(taskMap, activityLog)
							fmt.Println("In Activity", studentkey)

							var summary ProgressSummary

							wkey := AddWhatsLeftAudit(*wl)
							key := AddTaskAudit(*audit)
							student.TaskAuditList = append(student.TaskAuditList, string(key))
							student.WhatsLeftList = append(student.WhatsLeftList, string(wkey))
							// fmt.Println("Taysk Audit Key: ", key)

							summary.Major = taskMap.Name
							summary.StudentKey = studentkey
							summary.Comp = len(audit.ReqCommits) - len(wl.ReqCommits)
							summary.Total = len(taskMap.Requirements)
							summary.AuditID = string(key)
							summary.TaskID = tmap
							summary.ActivityID = alog
							summary.LastUpdate = fmt.Sprintf("%4d", time.Now().UTC().Unix())
							// fmt.Println(len(audit.Commits))
							UpsertNewStudentSummary(studentkey, summary)

						}

					}

					student.ShouldUpdateAuditFlag = false

					client.boltDB.Update(func(tx *bolt.Tx) error {

						b, err := tx.CreateBucketIfNotExists([]byte("enrolled_students"))
						if err != nil {
							return err
						}

						// b := tx.Bucket()
						// if b == nil {
						// 	return fmt.Errorf("Bucket %s not found!", "enrolled_students")
						// }
						// v := b.Get([]byte(students[i]))

						jsonData, err := json.Marshal(student)

						if err != nil {
							fmt.Println(err)
						}
						err = b.Put([]byte(students[i]), jsonData)
						return err
					})
					didProcess = true
					globalSocketHandler.Broadcast([]byte(studentkey))
				}
			}(i)

		}

		wg.Wait()
		// fmt.Println("Finished for loop")

		if didProcess {

			timeTrack(beg, "Daemon Total Run Time  \t\t\t")
		}

	}

	stop := schedule(ping, rt.Pluse*time.Millisecond)
	rt.IsWatching = stop
	// time.Sleep(2500 * time.Millisecond)
	// stop <- true
	// time.Sleep(2500 * time.Millisecond)

	// fmt.Println(rt)

	fmt.Println("Done")
}
