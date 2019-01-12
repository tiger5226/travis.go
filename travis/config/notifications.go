package config

type Notifications struct {
	Email    Email  `json:"email"`
	FlowDock string `json:"flowdock"`
	IRC      IRC    `json:"irc"`
	Slack    Slack  `json:"slack"`
}

type HipChat struct {
	roomNotification
	Format         string `json:"format"`
	Notify         bool   `json:"notify"`
	OnPullRequests bool   `json:"on_pull_requests"`
}

type Campfire struct {
	roomNotification
}

type IRC struct {
	baseNotification
	Channels   []string `json:"channels"`
	ChannelKey string   `json:"channel_key"`
	NickName   string   `json:"nick"`
	Password   string   `json:"password"`
	SkipJoin   bool     `json:"skip_join"`
	Template   []string `json:"template"`
	UseNotice  bool     `json:"use_notice"`
}

type Email struct {
	baseNotification
	Recipients []string `json:"recipients"`
}

type Slack struct {
	roomNotification
}

const (
	ALWAYS = "always"
	NEVER  = "never"
	CHANGE = "change"
)

type roomNotification struct {
	baseNotification
	Rooms    []string `json:"rooms"`
	Template []string `json:"template"`
}

type baseNotification struct {
	OnSuccess string `json:"on_success"`
	OnFailure string `json:"on_failure"`
}
