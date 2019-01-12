package config

import "github.com/tiger5226/travis.go/travis/shared"

type Config map[string]interface{}

type config struct {
	AddOns   shared.AddOns `json:"addons"`
	Branches struct {
		Only []string `json:"only"`
	} `json:"branches"`
	Cache struct {
		Directories []string `json"directories"`
	} `json:"cache"`
	Distribution string        `json:"dist"`
	GlobalEnv    []string      `json:"global_env"`
	Go           []string      `json:"go"`
	Group        string        `json:"group"`
	Language     string        `json:"language"`
	Install      []string      `json:"install"`
	OS           string        `json:"os"`
	Script       []string      `json:"script"`
	Sudo         bool          `json:"sudo"`
	Type         string        `json:"type"`
	Deploy       shared.Deploy `json:"deploy"`
}
