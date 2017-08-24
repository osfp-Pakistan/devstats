package gha2db

import (
	"time"
)

// AnyArray - holds array of interface{} - just a shortcut
type AnyArray []interface{}

// Event - full GHA (GitHub Archive) event structure
type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Org       *Org      `json:"org"`
	Payload   Payload   `json:"payload"`
}

// Payload - GHA Payload structure
type Payload struct {
	PushID       *int      `json:"push_id"`
	Size         *int      `json:"size"`
	Ref          *string   `json:"ref"`
	Head         *string   `json:"head"`
	Before       *string   `json:"before"`
	Action       *string   `json:"action"`
	RefType      *string   `json:"ref_type"`
	MasterBranch *string   `json:"master_branch"`
	Description  *string   `json:"description"`
	Number       *int      `json:"number"`
	Forkee       *Forkee   `json:"forkee"`
	Release      *Release  `json:"release"`
	Member       *Actor    `json:"member"`
	Issue        *Issue    `json:"issue"`
	Comment      *Comment  `json:"comment"`
	Commits      *[]Commit `json:"commits"`
	Pages        *[]Page   `json:"pages"`
}

// Repo - GHA Repo structure
type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Actor - GHA Actor structure
type Actor struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}

// Org - GHA Org structure
type Org struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}

// Forkee - GHA Forkee structure
type Forkee struct {
	ID int `json:"id"`
}

// Release - GHA Release structure
type Release struct {
	ID int `json:"id"`
}

// Issue - GHA Issue structure
type Issue struct {
	ID int `json:"id"`
}

// Comment - GHA Comment structure
type Comment struct {
	ID                  int       `json:"id"`
	Body                string    `json:"body"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	User                Actor     `json:"user"`
	CommitID            *string   `json:"commit_id"`
	OriginalCommitID    *string   `json:"original_commit_id"`
	DiffHunk            *string   `json:"diff_hunk"`
	Position            *int      `json:"position"`
	OriginalPosition    *int      `json:"original_position"`
	Path                *string   `json:"path"`
	PullRequestReviewID *int      `json:"pull_request_review_id"`
	Line                *int      `json:"line"`
}

// Commit - GHA Commit structure
type Commit struct {
	SHA      string `json:"sha"`
	Author   Author `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
}

// Author - GHA Commit Author structure
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Page - GHA Page structure
type Page struct {
	SHA    string `json:"sha"`
	Action string `json:"action"`
	Title  string `json:"title"`
}

// OrgIDOrNil - return Org ID from pointer or nil
func OrgIDOrNil(orgPtr *Org) interface{} {
	if orgPtr == nil {
		return nil
	}
	return orgPtr.ID
}

// IssueIDOrNil - return Issue ID from pointer or nil
func IssueIDOrNil(issuePtr *Issue) interface{} {
	if issuePtr == nil {
		return nil
	}
	return issuePtr.ID
}

// CommentIDOrNil - return Comment ID from pointer or nil
func CommentIDOrNil(commPtr *Comment) interface{} {
	if commPtr == nil {
		return nil
	}
	return commPtr.ID
}

// ForkeeIDOrNil - return Forkee ID from pointer or nil
func ForkeeIDOrNil(forkPtr *Forkee) interface{} {
	if forkPtr == nil {
		return nil
	}
	return forkPtr.ID
}

// ActorIDOrNil - return Actor ID from pointer or nil
func ActorIDOrNil(actPtr *Actor) interface{} {
	if actPtr == nil {
		return nil
	}
	return actPtr.ID
}

// ReleaseIDOrNil - return Release ID from pointer or nil
func ReleaseIDOrNil(relPtr *Release) interface{} {
	if relPtr == nil {
		return nil
	}
	return relPtr.ID
}