package candlesgrabber

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Reader is an interface for managing kind of readers.
type Reader interface {
	// Read read historical candles by a FIGI and a year.
	Read(ctx context.Context, figi string, year int) ([]Candle, error)
}

// Candle is the information about one candle.
type Candle struct {
	UID    uuid.UUID
	Date   time.Time
	Open   float32
	Close  float32
	High   float32
	Low    float32
	Volume int
}
