package main

type Commit struct {
	// Requirement string
	// Statisifier string
	Activity string
	// Person      string
}

type SatisfactionAudit struct {
	Complete      bool
	Name          string
	Count         int
	ExpectedCount int
	Commits       []Commit
}

type RequirementAudit struct {
	HasComplete bool
	Name        string
	SatCommits  []SatisfactionAudit
}
type TaskAudit struct {
	Name       string
	Person     string
	ReqCommits []RequirementAudit
}

type WhatsLeft struct {
	Name       string
	Person     string
	ReqCommits []RequirementAudit
}

func RunAudit(task TaskMap, activitylog ActivityLog) (*TaskAudit, *WhatsLeft) {
	counter := 0

	taskAudit := new(TaskAudit)
	whatsLeft := new(WhatsLeft)

	taskAudit.Name = task.Name
	// This loops over all of the requirements in the task (i)
	for i := 0; i < len(task.Requirements); i++ {

		reqAudit := new(RequirementAudit)
		reqAudit.HasComplete = false
		// list of OR statements
		// This loops through all of the possible satisfiers (j)
		for j := 0; j < len(task.Requirements[i].Satisfiers); j++ {
			hitCount := 0
			audit := new(SatisfactionAudit)
			commit := new(Commit)
			// Now we want to check the requirements and care about the true

			// list of AND statements
			// This loop goes through all of the AND statment assets
			for l := 0; l < len(task.Requirements[i].Satisfiers[j].Core); l++ {

				thecore := task.Requirements[i].Satisfiers[j].Core[l]
				// fmt.Println(counter)
				if activitylog.Completed[thecore] {

					// fmt.Println(counter)
					hitCount++
					// commit.Requirement = task.Requirements[i].Name
					commit.Activity = thecore
					// commit.Person = activitylog.Name

					// fmt.Println(*commit)
					audit.Commits = append(audit.Commits, *commit)
				}

			}
			audit.Count = hitCount
			audit.ExpectedCount = task.Requirements[i].Satisfiers[j].Count
			audit.Name = task.Requirements[i].Satisfiers[j].Uuid

			if audit.Count >= audit.ExpectedCount {
				audit.Complete = true
				reqAudit.HasComplete = true
			}

			if len(audit.Commits) > 0 {
				reqAudit.SatCommits = append(reqAudit.SatCommits, *audit)

			}

			counter++
		}
		reqAudit.Name = task.Requirements[i].Name
		if reqAudit.HasComplete == false {
			whatsLeft.ReqCommits = append(whatsLeft.ReqCommits, *reqAudit)
		}
		// fmt.Println(*reqAudit)
		taskAudit.Person = activitylog.Name
		// taskAudit.AuditID = activitylog.
		taskAudit.ReqCommits = append(taskAudit.ReqCommits, *reqAudit)
	}
	return taskAudit, whatsLeft
}
