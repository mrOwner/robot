package app

import (
	"context"
	"time"
	"unicode/utf8"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

const (
	format = "%*s.....................%s"
	keyDur = "duration"
)

//go:generate mockgen  --build_flags=--mod=mod -destination=../../mocks/app.go -package=mock github.com/mrOwner/robot/internal/app LifeCycle
type LifeCycle interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type task struct {
	name string
	LifeCycle
}

type App struct {
	seq    []task
	sync   []task
	cancel context.CancelFunc
}

func New() *App {
	return &App{}
}

func (a *App) AddSequential(name string, lc LifeCycle) {
	a.seq = append(a.seq, task{
		name:      name,
		LifeCycle: lc,
	})
}

func (a *App) AddSynchronous(name string, lc LifeCycle) {
	a.sync = append(a.sync, task{
		name:      name,
		LifeCycle: lc,
	})
}

func (a *App) Start(ctx context.Context) error {
	ctx, a.cancel = context.WithCancel(ctx)
	defer a.cancel()

	maxchar := a.maxLengthName()

	for _, task := range a.seq {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			start := time.Now()
			log.Info().Msgf(format, maxchar, task.name, "starting.")
			if err := task.Start(ctx); err != nil {
				return err
			}
			log.Info().Dur(keyDur, time.Since(start)).Msgf(format, maxchar, task.name, "done.")
		}
	}

	eg, gtx := errgroup.WithContext(ctx)
	for _, task := range a.sync {
		task := task
		eg.Go(func() error {
			start := time.Now()
			defer func() {
				log.Info().Dur(keyDur, time.Since(start)).Msgf(format, maxchar, task.name, "done.")
			}()
			log.Info().Msgf(format, maxchar, task.name, "starting.")

			return task.Start(gtx)
		})
	}

	done := make(chan error)
	defer close(done)
	go func() {
		done <- eg.Wait()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}

func (a *App) Stop(_ context.Context) error {
	a.cancel()

	return nil
}

func (a *App) maxLengthName() int {
	var max, n int
	for _, task := range append(a.seq, a.sync...) {
		n = utf8.RuneCountInString(task.name)
		if max < n {
			max = n
		}
	}

	return max
}
