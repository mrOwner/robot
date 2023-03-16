package collectcandles

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	db "github.com/mrOwner/robot/db/postgres/sqlc"
	cg "github.com/mrOwner/robot/pkg/candles_grabber"
	"github.com/stretchr/testify/require"
)

func TestCopier(t *testing.T) {
	got := cg.Candles{
		{
			UID:    uuid.New(),
			Date:   time.Now(),
			Open:   1,
			Close:  2,
			High:   3,
			Low:    4,
			Volume: 5,
		},
	}
	want := []db.CreateCandlesParams{}

	err := copier.CopyWithOption(&want, got, copierOpt())

	require.NoError(t, err)
	require.Len(t, want, len(got))
	opts := cmp.Options{
		cmp.FilterValues(func(x, y interface{}) bool {
			_, okX1 := x.(cg.Candles)
			_, okX2 := x.([]db.CreateCandlesParams)
			_, okY1 := y.(cg.Candles)
			_, okY2 := y.([]db.CreateCandlesParams)
			return (okX1 || okX2) && (okY1 || okY2)
		}, cmp.Transformer("", func(a interface{}) []db.CreateCandlesParams {
			candles, ok := a.(cg.Candles)
			if !ok {
				return a.([]db.CreateCandlesParams)
			}
			arg := make([]db.CreateCandlesParams, len(candles))
			for i, candle := range candles {
				arg[i] = db.CreateCandlesParams{
					UID:    candle.UID.String(),
					Date:   candle.Date,
					Open:   candle.Open,
					Close:  candle.Close,
					High:   candle.High,
					Low:    candle.Low,
					Volume: int64(candle.Volume),
				}
			}
			return arg
		})),
	}
	require.Empty(t, cmp.Diff(want, got, opts))
}
