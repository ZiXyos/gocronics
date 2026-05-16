package cron

import "context"

// Job type represent the CronJob structures.
type Job struct {
	Spec Spec
	Fn   func(ctx context.Context)
}

// NewJob create a new job function and it's cron spec.
func NewJob(spec Spec, fn func(context.Context)) *Job {
	return &Job{
		Spec: spec,
		Fn:   fn,
	}
}
