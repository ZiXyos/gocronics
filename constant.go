package cron

type Spec string

const (
	Monthly Spec = "@monthly"
	Weekly  Spec = "@weekly"
	Daily   Spec = "@daily"
	Hourly  Spec = "@hourly"
)

func (s Spec) toCron() string {
	switch s {
	case Monthly:
		return "0 0 1 * *"
	case Weekly:
		return "0 0 * * 0"
	case Daily:
		return "0 0 * * *"
	case Hourly:
		return "0 * * * *"
	default:
		return string(s)
	}
}

