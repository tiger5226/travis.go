package matrix

import "github.com/tiger5226/travis.go/travis/config"

type Matrix map[string]interface{}

type matrix []Element

type Element struct {
	ID           int           `json:"id"`
	RepositoryID int           `json:"repository_id"`
	ParentID     int           `json:"parent_id"`
	Number       string        `json:"number"`
	State        string        `json:"state"`
	Config       config.Config `json:"config"`
	Status       int           `json:"status"`
	Result       int           `json:"result"`
	Commit       string        `json:"commit"`
	Branch       string        `json:"branch"`
	Message      string        `json:"message"`
}
