# gocronics

A lightweight Go library for scheduling cron jobs with a clean, context-aware API built on top of [`robfig/cron`](https://github.com/robfig/cron).

## Installation

```bash
go get github.com/zixyos/gocronics
```

## Usage

```go
import cron "github.com/zixyos/gocronics"

handler := cron.NewCronHandler(
    cron.WithLogger(slog.Default()),
    cron.WithCronFunction(
        cron.NewCronJob(cron.Daily, func(ctx context.Context) {
            // runs every day at midnight
        }),
        cron.NewCronJob(cron.Hourly, func(ctx context.Context) {
            // runs every hour
        }),
        cron.NewCronJob("30 9 * * 1-5", func(ctx context.Context) {
            // custom expression: weekdays at 09:30
        }),
    ),
)

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

if err := handler.Run(ctx); err != nil {
    log.Fatal(err)
}
```

## Predefined schedules

| Constant  | Runs at                  |
|-----------|--------------------------|
| `Monthly` | Midnight on the 1st      |
| `Weekly`  | Midnight every Sunday    |
| `Daily`   | Midnight every day       |
| `Hourly`  | Start of every hour      |

Any standard cron expression (`"* * * * *"`) is also accepted as a `Spec`.

## API

### `NewCronHandler(options ...HandlerOption) *Handler`

Creates a new handler. Options:

- `WithLogger(logger *slog.Logger)` — attaches a structured logger (used to warn on job registration failures).
- `WithCronFunction(jobs ...*Job)` — registers one or more jobs.

### `NewCronJob(spec Spec, fn func(context.Context)) *Job`

Creates a job from a schedule spec and a context-aware function.

### `(*Handler).Run(ctx context.Context) error`

Starts the scheduler and blocks until `ctx` is cancelled.

### `(*Handler).Stop(ctx context.Context) error`

Stops the scheduler when `ctx` is cancelled.

## License

MIT
