# Welcome to go-dart

![Prototype UI](https://github.com/drbh/go-dart/blob/master/screenshot.png)

The fastest way to audit students degrees

This is the Golang implementation of 🎯 D.A.R.T (Degree Auditing Real Time)

A suiting name for the application. It does exactly that. It audits degrees in real time!

### How this is different
This project aids a core process at any university - the degree auditing process.

This is an open source program written in `Golang`, that can audit 100,000+ students with 10+ degrees each in realtime and provide insights into a universities student body degree completion.

Ed-Tech applications generally focus on retention, ease of access, or UI for faculty to better aid students.

**This is not that**

This is an application focused on performance, in order to allow schools to have options for core systems like degree auditing applications.

Enterprise degree auditing software are closed-source, difficult for analysis integration and do not support modern/future computer science efforts around machine learning, microservices, concurrency, and parallel computing. 

DART aims to fill the need for a fast degree auditing system that is light-weight, flexible and provide access to data.

Additionally DART can provide prescriptive student outreach and deep insights into trends across all your student's educational journeys/

## I just want to:

1) Run the application (Heroku app going to be deployed soon)
2) Install the app
3) Review architecture
4) Read API docs

## 1. Demo App

Sorry, this is not up and running yet. First, an example API will be available, then a UI app will be deployed.

## 2. How to install

Make sure you have go installed, also your gonna need a few go libraries that your terminal will tell you about when you first try to build the app.

```
git clone https://github.com/drbh/go-dart.git
cd go-dart
```

`go get` the go packages that are the dependencies. 

*deps*
```
 go get github.com/gin-gonic/contrib/static
 go get github.com/gin-gonic/gin
 go get github.com/gin-contrib/cors
 go get gopkg.in/olahol/melody.v1
```

okay now you can build the app without issue

```
go build
./go-dart
```

## 3. Basic Architecture

There are a few important items in DART

1. Student
2. Task Map
3. Activity Log
4. Task Audit
5. What Left
6. Student Summary

The whole framework is student-centric

Student is an item that holds IDs for all of the other items, image coming soon.

![Basic Framework](https://github.com/drbh/go-dart/blob/master/datamodel.png)


## 4. API docs
`GET    /api/`  
Base URL, this does nothing

`POST   /api/enroll-student`  
You can enroll your students in the program

`POST   /api/task-map`  
You can add a task-map or as some schools call it - a major map

`POST   /api/activity-log`  
You can add an activity-log or a JSON blob of student history

`POST   /api/ad-hoc`  
I don't think this works right now. But this is an endpoint to run ad-hoc degree audits (doesn't save the task-map, activity-log or audit to the DB)

`POST   /api/process`  


`GET    /api/acts`  
get all activity log ids

`GET    /api/students`  
get all students ids

`GET    /api/students-expanded`  
get a lot of student info back. all audits, all ids to audits and task map and activity logs

`GET    /api/maps`  
get all task maps

`GET    /api/map/:name`  
get a single map back

`GET    /api/student/:name`  
get a single student info back (no expanded audits)

`GET    /api/student-expanded/:name`  
get a single student back expanded (shows all their degree run information)

`GET    /api/act/:name`  
get single actvity log back by id

`GET    /api/audit/:name`  
get single audit back by id

`GET    /api/student-summary/:name`  
get a single student summary back (shows all of their current degree standings)

`GET    /api/ws`  
WebSocket endpoint to listen for updates (for a UI)

`POST   /api/auto-build-student`  


`POST   /api/update`  
update a students activity (like when a class completes!)