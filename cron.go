// package cron represent the cron job service handler package.
package cron

import (
	"context"
	"log/slog"

	rcron "github.com/robfig/cron"
)

// Handler type represent the cron job structure.
type Handler struct {
	logger *slog.Logger
	cron   *rcron.Cron

	jobs []*Job
}

// HandlerOption type represent the function type for cron handler service.
type HandlerOption func(*Handler)

// WithLogger add a logger to the service handler.
func WithLogger(logger *slog.Logger) HandlerOption {
	return func(c *Handler) {
		c.logger = logger
	}
}

// WithCronFunction push the defined jobs in the handler.
func WithCronFunction(jobs ...*Job) HandlerOption {
	return func(c *Handler) {
		c.jobs = jobs
	}
}

// NewCronHandler create a new cron job service handler.
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

// Run register and execute the jobs.
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

// Stop gracefully shutdown jobs.
// TODO: implement gracefull shutdown.
func (ch *Handler) Stop(ctx context.Context) error {
	<-ctx.Done()
	ch.cron.Stop()

	return nil
}

// Name return the service name.
func (ch *Handler) Name() string {
	return "cron-job-handler"
}
