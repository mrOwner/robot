package collectcandles

import (
	"bufio"
	"context"
	"errors"
	"os"
	"strings"
	"time"

	db "github.com/mrOwner/robot/db/postgres/sqlc"
	cg "github.com/mrOwner/robot/pkg/candles_grabber"
	"github.com/mrOwner/robot/util"
	"golang.org/x/sync/errgroup"
)

type Collector struct {
	path    string
	grabber cg.Reader
	querier db.Querier
	cancel  context.CancelFunc
}

func New(path string, grabber cg.Reader, dbtx db.DBTX) *Collector {
	return &Collector{
		path:    path,
		grabber: grabber,
		querier: db.New(dbtx),
	}
}

func (c *Collector) Start(ctx context.Context) (err error) {
	ctx, c.cancel = context.WithCancel(ctx)
	defer c.cancel()

	file, err := os.Open(c.path)
	if err != nil {
		return err
	}
	defer util.JoinErrs(err, file.Close)

	ch := make(chan []cg.Candle)
	eg, gtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return c.read(gtx, file, ch)
	})
	eg.Go(func() error {
		return c.save(gtx, ch)
	})

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

func (c *Collector) Stop(_ context.Context) error {
	c.cancel()

	return nil
}

func (c *Collector) read(ctx context.Context, file *os.File, ch chan []cg.Candle) error {
	defer close(ch)

	bf := bufio.NewScanner(file)
	for bf.Scan() {
		figi := strings.TrimSpace(bf.Text())
		for y := time.Now().Year(); ; y-- {
			candles, err := c.grabber.Read(ctx, figi, y)
			if err != nil {
				if errors.Is(err, cg.ErrDataNotFound) {
					break
				}
				if errors.Is(err, cg.ErrRateLimit) {
					time.Sleep(5 * time.Second)
					y++
					continue
				}
				return err
			}

			ch <- candles
		}
	}
	if bf.Err() != nil {
		return bf.Err()
	}

	return nil
}

func (c *Collector) save(ctx context.Context, ch chan []cg.Candle) error {
	for candles := range ch {
		arg := make([]db.CreateCandlesParams, len(candles))
		for i, candle := range candles {
			arg[i] = db.CreateCandlesParams{
				Uid:    candle.UID.String(),
				Date:   candle.Date,
				Open:   candle.Open,
				Close:  candle.Close,
				High:   candle.High,
				Low:    candle.Low,
				Volume: int64(candle.Volume),
			}
		}
		_, err := c.querier.CreateCandles(ctx, arg)
		if err != nil {
			return err
		}
	}

	return nil
}
