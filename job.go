package cron

import "context"

type Job struct {
	Spec Spec
	Fn   func(ctx context.Context)
}

func NewCronJob(spec Spec, fn func(context.Context)) *Job {
	return &Job{
		Spec: spec,
		Fn:   fn,
	}
}
