package travis

import (
	"time"

	"github.com/tiger5226/travis.go/travis/config"
	"github.com/tiger5226/travis.go/travis/matrix"
)

// https://docs.travis-ci.com/user/notifications/#Webhooks-Delivery-Format

const (
	statusSuccess = 0
)

// DeployCheckFn describes the acceptable function signature DeployCheckHook. This should set on application startup if
// used.
type DeployCheckFn func(w Webhook) bool

// MatchFn describes the acceptable function signature for the IsMatchHook.  This should set on application startup if
// used.
type MatchFn func(w Webhook) bool

// ShouldDeployFn executes when non-nil. The default checks that it is on the master branch, it was successful, and that
// it is not a pull request.
var ShouldDeployFn DeployCheckFn

// IsMatchFn executes when non-nil. The default only checks the branch, repo, and owner. This is important to make sure
// a 3rd party is not sending a travis webhook notification to an api that leverages this. Always make sure it is for you!
// If you have checks in addition to the basic defaults, you should use this.
var IsMatchFn MatchFn

// Webhook is the data structure of the payload travis sends for webhook notifications.
type Webhook struct {
	AuthorName        string        `json:"author_name"`
	AuthorEmail       string        `json:"author_email"`
	BaseCommit        string        `json:"base_commit"`
	Branch            string        `json:"branch"`
	BuildURL          string        `json:"build_url"`
	Commit            string        `json:"commit"`
	CommitID          int           `json:"commit_id"`
	CompareURL        string        `json:"compare_url"`
	CommittedAt       time.Time     `json:"committed_at"`
	CommitterName     string        `json:"committer_name"`
	CommitterEmail    string        `json:"committer_email"`
	Config            config.Config `json:"config"`
	Duration          int           `json:"duration"`
	FinishedAt        time.Time     `json:"finished_at"`
	HeadCommit        string        `json:"head_commit"`
	ID                int           `json:"id"`
	Message           string        `json:"message"`
	Matrix            matrix.Matrix `json:"matrix"`
	Number            string        `json:"number"`
	PullRequest       bool          `json:"pull_request"`
	PullRequestNumber int           `json:"pull_request_number"`
	PullRequestTitle  string        `json:"pull_request_title"`
	Tag               string        `json:"tag"`
	Type              string        `json:"type"`
	Repository        Repository    `json:"repository"`
	Result            int           `json:"result"`
	ResultMessage     string        `json:"result_message"`
	StartedAt         time.Time     `json:"started_at"`
	State             string        `json:"state"`
	Status            int           `json:"status"`         // status and result are the same
	StatusMessage     string        `json:"status_message"` // status_message and result_message are the same
}

// IsMatch make sure the webhook is for you...
func (w Webhook) IsMatch(branch string, repo string, owner string) bool {

	return (IsMatchFn == nil || IsMatchFn(w)) &&
		w.Branch == branch &&
		w.Repository.Name == repo &&
		w.Repository.OwnerName == owner
}

func (w Webhook) ShouldDeploy() bool {
	if ShouldDeployFn == nil {
		return ShouldDeployFn(w)
	}
	// when travis builds a pull request, Branch is the target branch, not the origin branch
	// source: https://docs.travis-ci.com/user/environment-variables/#Default-Environment-Variables
	return w.Status == statusSuccess && w.Branch == "master" && !w.PullRequest
}

func (w Webhook) DeploySummary() string {
	return w.Commit[:8] + ": " + w.Message
}
