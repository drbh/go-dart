package main

// import (
// 	"fmt"
// )

// {
// 	"name": "david",
// 	"id": "341412",
// 	"summary": [
// 		{
// 			"major": "mkt",
// 			"comp":  90,
// 			"auditID": "afasfjk",
// 			"activityID": "fdsfas"
// 			"lastUpdate": 120214
// 		}
// 	]
// }

type SutdentSummary struct {
	Name      string
	Id        string
	Summaries []ProgressSummary
}

type ProgressSummary struct {
	Major      string
	StudentKey string
	Comp       int
	Total      int
	AuditID    string
	TaskID     string
	ActivityID string
	LastUpdate string
}
