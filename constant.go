package cron

// Spec type represent the cron job spec.
type Spec string

const (
	// Monthly runs the job once a month.
	Monthly Spec = "@monthly"

	// Weekly runs the job once a week.
	Weekly Spec = "@weekly"

	// Daily runs the job once a day.
	Daily Spec = "@daily"

	// Hourly runs the job once an hour.
	Hourly Spec = "@hourly"
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
