package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	// "golang.org/x/net/websocket"
	"github.com/gin-contrib/cors"
	"gopkg.in/olahol/melody.v1"
	"net/http"
	"path/filepath"
	// "strconv"
	"time"
)

type AuditRequest struct {
	ActivityLogID string `json:"activityLogIDs"`
	TaskMapID     string `json:"taskMapIDs"`
}

type UpdateRequest struct {
	MyType  string `json:"type,omitempty"`
	Who     string `json:"who,omitempty"`
	Major   string `json:"major,omitempty"`
	Class   string `json:"class,omitempty"`
	Catyear string `json:"catyr,omitempty"`
}

type AdHocRequest struct {
	Act  ActivityLog `json:"activityLog"`
	Task TaskMap     `json:"taskMap"`
}

// func EchoLengthServer(ws *websocket.Conn) {
// 	var msg string

// 	for {
// 		websocket.Message.Receive(ws, &msg)
// 		fmt.Println("Got Message", msg)
// 		length := len(msg)
// 		if err := websocket.Message.Send(ws, strconv.FormatInt(int64(length), 10)); err != nil {
// 			fmt.Println("can't send message length")
// 			break
// 		}
// 	}
// }

func StartServer() {
	fmt.Println("TEST")
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	router.Use(cors.Default())
	m := melody.New()
	globalSocketHandler = m

	// Serve frontend static files
	// publicPath, _ := filepath.Abs("./views")
	publicPath, _ := filepath.Abs("./build")
	fmt.Println(publicPath)
	// s := static.Serve("/ui", static.LocalFile(publicPath, true))
	router.Use(static.Serve("/ui", static.LocalFile(publicPath, true)))
	router.Use(static.Serve("/", static.LocalFile(publicPath, true)))
	// router.NotFound(static.Serve("/ui", static.LocalFile(publicPath, true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		// CreateOrder swagger:route POST /orders orders createOrder
		//
		// Handler to create an order.
		//
		// Responses:
		//        200: orderResponse
		//        422: validationError
		api.POST("/enroll-student", func(c *gin.Context) {

			var stu Student
			stu.ShouldUpdateAuditFlag = true
			fmt.Printf("Added Update Flag to")
			c.BindJSON(&stu)
			fmt.Printf("Student to store: %v\n", stu)

			key := AddStudent(stu)
			c.JSON(http.StatusOK, gin.H{
				"student-key": string(key),
			})

			// Push UI update
			m.Broadcast([]byte(""))
		})

		// CreateOrder swagger:route POST /task-map orders createOrder
		//
		// Handler to create an order.
		//
		// Responses:
		//        200: orderResponse
		//        422: validationError
		api.POST("/task-map", func(c *gin.Context) {

			var url TaskMap
			c.BindJSON(&url)
			fmt.Printf("Task Map to store: %v\n", url)

			key := AddTaskMap(url)
			c.JSON(http.StatusOK, gin.H{
				"key": string(key),
			})
		})

		// CreateOrder swagger:route POST /activity-log orders createOrder
		//
		// Handler to create an order.
		//
		// Responses:
		//        200: orderResponse
		//        422: validationError
		api.POST("/activity-log", func(c *gin.Context) {
			var url ActivityLog
			c.BindJSON(&url)
			fmt.Printf("Task Map to store: %v\n", url)

			key := AddActivityLog(url)
			c.JSON(http.StatusOK, gin.H{
				"key": string(key),
			})
		})

		api.POST("/ad-hoc", func(c *gin.Context) {
			start := time.Now()
			var ahreq AdHocRequest
			var audit *TaskAudit

			c.BindJSON(&ahreq)
			fmt.Println(ahreq.Act)
			fmt.Println(ahreq.Task)

			audit, _ = RunAudit(ahreq.Task, ahreq.Act)
			fmt.Println(audit)
			timeTrack(start, "Ad Hoc Ran\t\t\t\t")

			// jsond, _ := json.Marshal(*audit)

			c.JSON(http.StatusOK, gin.H{
				"data": audit,
			})
		})

		api.POST("/process", func(c *gin.Context) {

			var adreq AuditRequest
			c.BindJSON(&adreq)
			var audit *TaskAudit

			beg := time.Now()

			task := GetTaskMap(adreq.TaskMapID)
			activityLog := GetActivityLog(adreq.ActivityLogID)
			// timeTrack(start, "Fetched/Marshaled\t\t\t")

			start := time.Now()
			audit, _ = RunAudit(task, activityLog)
			timeTrack(start, "Processed Time\t\t\t\t")

			start = time.Now()
			key := AddTaskAudit(*audit)
			timeTrack(start, "Written Time\t\t\t\t")
			timeTrack(beg, "Total Time  \t\t\t\t")

			c.JSON(http.StatusOK, gin.H{
				"data":  key,
				"audit": audit,
			})
		})

		api.GET("/acts", func(c *gin.Context) {
			acts := GetActivityKeys()
			c.JSON(http.StatusOK, gin.H{
				"data": acts,
			})
		})

		api.GET("/students", func(c *gin.Context) {
			students := GetStudentKeys()
			c.JSON(http.StatusOK, gin.H{
				"data": students,
			})
		})

		api.GET("/students-expanded", func(c *gin.Context) {
			students := GetStudentKeys()

			var payload []map[string]string

			for i := 0; i < len(students); i++ {

				entry := map[string]string{}

				student := GetStudent(students[i])

				// entry["name"] = make(map[string]string)

				entry["name"] = student.Name
				entry["id"] = students[i]
				fmt.Println(entry)

				payload = append(payload, entry)

			}
			// {
			//   name: "Doug",
			//   id: "CDFCAA47-4E2E-D20C-DAFD-4D912FBCC0D1",
			// }
			c.JSON(http.StatusOK, gin.H{
				"data": payload,
			})
		})

		api.GET("/maps", func(c *gin.Context) {
			maps := GetTaskMapKeys()
			c.JSON(http.StatusOK, gin.H{
				"data": maps,
			})
		})

		api.GET("/map/:name", func(c *gin.Context) {
			// maps := GetTaskMapKeys()
			name := c.Param("name")
			taskMap := GetTaskMap(name)

			c.JSON(http.StatusOK, gin.H{
				"data": taskMap,
			})
		})

		api.GET("/student/:name", func(c *gin.Context) {
			// maps := GetTaskMapKeys()
			name := c.Param("name")
			student := GetStudent(name)

			c.JSON(http.StatusOK, gin.H{
				"data": student,
			})
		})

		api.GET("/student-expanded/:name", func(c *gin.Context) {
			// maps := GetTaskMapKeys()
			name := c.Param("name")
			student := GetStudent(name)

			var auditList []TaskAudit
			for i := 0; i < len(student.TaskAuditList); i++ {
				audit := GetTaskAudit(student.TaskAuditList[i])
				// fmt.Println(audit)
				auditList = append(auditList, audit)
			}
			var activityList []ActivityLog
			for i := 0; i < len(student.ActivityLogList); i++ {
				activity := GetActivityLog(student.ActivityLogList[i])
				// fmt.Println(audit)
				activityList = append(activityList, activity)
			}

			c.JSON(http.StatusOK, gin.H{
				"data":         student,
				"audits":       auditList,
				"activityList": activityList,
			})
		})

		api.GET("/act/:name", func(c *gin.Context) {
			// maps := GetTaskMapKeys()
			name := c.Param("name")
			fmt.Println(name)
			actlog := GetActivityLog(name)

			c.JSON(http.StatusOK, gin.H{
				"data": actlog,
			})
		})

		api.GET("/audit/:name", func(c *gin.Context) {
			// maps := GetTaskMapKeys()
			name := c.Param("name")
			taskAudit := GetTaskAudit(name)

			c.JSON(http.StatusOK, gin.H{
				"data": taskAudit,
			})
		})

		api.GET("/student-summary/:name", func(c *gin.Context) {
			// maps := GetTaskMapKeys()
			name := c.Param("name")
			studentSummary := GetStudentSummary(name)

			c.JSON(http.StatusOK, gin.H{
				"data": studentSummary,
			})
		})

		api.GET("/ws", func(c *gin.Context) {
			m.HandleRequest(c.Writer, c.Request)
		})

		api.POST("/auto-build-student", func(c *gin.Context) {
			var req UpdateRequest
			c.BindJSON(&req)

			// fmt.Println()
			// KRBEMFDZDC

			key := AddCoreToStudentActivityLog(req.Who, req.Class)
			// key := AddCoreToActivityLog("TCUAXHXKQF", req.Class)
			c.JSON(http.StatusOK, gin.H{
				"data": key,
			})
		})

		api.POST("/update", func(c *gin.Context) {
			var req UpdateRequest
			c.BindJSON(&req)
			key := []byte("")
			fmt.Println(req)
			// fmt.Print("Write ", req.Type, " as ")

			switch req.MyType {

			case "add-class-activity":

				fmt.Println("Add Class to Student")
				key = AddCoreToStudentActivityLog(req.Who, req.Class)

			case "enroll-in-major":
				fmt.Println("Add Major to Student")
				AddMajorToStudent(req.Who, req.Major)

			case "add-blank-student":

				var url ActivityLog
				key := AddActivityLog(url)

				var stu Student
				stu.ShouldUpdateAuditFlag = true
				stu.ActivityLogList = []string{string(key)}
				studentkey := AddStudent(stu)
				fmt.Println(studentkey)

			}

			//push update to UI
			m.Broadcast([]byte(""))

			// KRBEMFDZDC

			// key := AddCoreToActivityLog("TCUAXHXKQF", req.Class)
			c.JSON(http.StatusOK, gin.H{
				"data": key,
			})
		})

		m.HandleMessage(func(s *melody.Session, msg []byte) {
			m.Broadcast(msg)
		})

		wing := func() {
			// fmt.Println("did")
			updateUI := false
			if updateUI {
				m.Broadcast([]byte("test"))
			}

		}
		schedule(wing, 3000*time.Millisecond)
		// stop <- true

	}

	// Start and run the server
	router.Run()
}
