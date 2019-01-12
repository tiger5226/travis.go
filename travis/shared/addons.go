package shared

type AddOns struct {
	Apt Apt `json:"apt"`
}

type Apt struct {
	Packages []string `json:"packages"`
}
