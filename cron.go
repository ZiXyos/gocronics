package cron

import (
	"context"
	"log/slog"

	rcron "github.com/robfig/cron"
)

type Handler struct {
	logger *slog.Logger
	cron   *rcron.Cron

	jobs []*Job
}

type HandlerOption func(*Handler)

func WithLogger(logger *slog.Logger) HandlerOption {
	return func(c *Handler) {
		c.logger = logger
	}
}

func WithCronFunction(jobs ...*Job) HandlerOption {
	return func(c *Handler) {
		c.jobs = jobs
	}
}

func NewCronHandler(options ...HandlerOption) *Handler {
	cron := rcron.New()

	handler := &Handler{
		cron: cron,
	}

	for _, option := range options {
		option(handler)
	}

	return handler
}

func (ch *Handler) Run(ctx context.Context) error {
	for _, job := range ch.jobs {
		if err := ch.cron.AddFunc(job.Spec.toCron(), func() { job.Fn(ctx) }); err != nil {
			ch.logger.WarnContext(ctx, "failed to register job", "error", err)
			continue
		}
	}

	ch.cron.Start()
	defer ch.cron.Stop()

	<-ctx.Done()
	return nil
}

func (ch *Handler) Stop(ctx context.Context) error {
	<-ctx.Done()
	ch.cron.Stop()

	return nil
}

func (ch *Handler) Name() string {
	return "cron-job-handler"
}
