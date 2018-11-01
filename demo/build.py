import glob, json, requests

counter = 0
major_map_name_id_map = {}

# loop over all majors
for filepath in glob.iglob("json-degrees/*.json"):
	mmap = json.loads(open(filepath).read())
	add_student_url = "http://localhost:3000/api/task-map"
	r = requests.post(add_student_url, json=mmap)
	map_key = r.json()["key"]
	major_map_name_id_map[mmap["name"]] = map_key

# print json.dumps(major_map_name_id_map)

student_body = [
{
	"Name": "Alice",
	"Majors": ["Art History"]
},
{
	"Name": "Bob",
	"Majors": ["ElecEngineering", "Construction"]
},
{
	"Name": "Carl",
	"Majors": ["CompSci", "ElecEngineering"]
},
{
	"Name": "Doug",
	"Majors": ["Marketing"]
},
{
	"Name": "Elizabeth",
	"Majors": ["Art History"]
},
{
	"Name": "Frankie",
	"Majors": ["ElecEngineering", "Construction"]
},
{
	"Name": "Grace",
	"Majors": ["Art History", "CompSci", "ElecEngineering", "Construction", "Marketing"]
},
{
	"Name": "Hilda",
	"Majors": ["Art History", "CompSci", "ElecEngineering", "Construction", "Marketing"]
},
]

for student in student_body:
	name = student['Name']
	counter = counter + 1

	# make a transript for the student
	actvity_log = { "name": "Transcript", "completed":{ "ROOT":True } }
	add_transcript_url = "http://localhost:3000/api/activity-log"
	r = requests.post(add_transcript_url, json=actvity_log)
	activity_key = r.json()["key"]	

	# loop over names of majors to get id
	major_id_list = []
	for major in student['Majors']:
		major_id_list += [major_map_name_id_map[major]]

	# build student
	student_payload = {
		"id" : str(counter),
		"name": name,
		"taskMapList": major_id_list,
		"activityLogList": [activity_key]
	}
	# print student_payload
	add_student_url = "http://localhost:3000/api/enroll-student"
	r = requests.post(add_student_url, json=student_payload)
	activity_key = r.json()["student-key"]
	print name, activity_key
print "Done"


