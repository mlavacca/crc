package actions

type Action int

const (
	Start Action = iota
	Stop
	Restart
	Enable
	Disable
	Status
	DaemonReload
)

func (action Action) String() string {
	actions := [...]string{
		"start",
		"stop",
		"restart",
		"enable",
		"disable",
		"status",
		"daemon-reload",
	}

	if int(action) >= 0 && int(action) < len(actions) {
		return actions[action]
	}

	return ""
}
