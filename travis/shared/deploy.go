package shared

type Deploy struct {
	App string `json:"app"`
	On  On     `json:"true"`
}

type On struct {
	Branch []string `json:"branch"`
	ApiKey struct {
		Secure string `json:"secure"`
	} `json:"api_key"`
	Provider    string `json:"provider"`
	SkipCleanup bool   `json:"skip_cleanup"`
}
