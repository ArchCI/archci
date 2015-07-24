package gitlabutil

import (
	"time"
)

type GitlabPushHook struct {
	Before string `json:"before"`
	After string `json:"after"`
	Ref string `json:"ref"`
	UserID int `json:"user_id"`
	UserName string `json:"user_name"`
	ProjectID int `json:"project_id"`
	Repository struct {
			   Name string `json:"name"`
			   URL string `json:"url"`
			   Description string `json:"description"`
			   Homepage string `json:"homepage"`
		   } `json:"repository"`
	Commits []struct {
		ID string `json:"id"`
		Message string `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL string `json:"url"`
		Author struct {
			   Name string `json:"name"`
			   Email string `json:"email"`
		   } `json:"author"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
}