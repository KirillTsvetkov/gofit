package models

type GoalStatus int

const (
	InProgress GoalStatus = iota
	Failed
	Success
)

func (s GoalStatus) String() string {
	return [...]string{"In Progress", "Failed", "Success"}[s]
}

func StatusFromString(s string) GoalStatus {
	switch s {
	case "In Progress":
		return InProgress
	case "Failed":
		return Failed
	case "Success":
		return Success
	default:
		return InProgress
	}
}
