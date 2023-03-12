// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: copyfrom.go

package db

import (
	"context"
)

// iteratorForCreateCandles implements pgx.CopyFromSource.
type iteratorForCreateCandles struct {
	rows                 []CreateCandlesParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateCandles) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateCandles) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Uid,
		r.rows[0].Date,
		r.rows[0].Open,
		r.rows[0].Close,
		r.rows[0].High,
		r.rows[0].Low,
		r.rows[0].Volume,
	}, nil
}

func (r iteratorForCreateCandles) Err() error {
	return nil
}

func (q *Queries) CreateCandles(ctx context.Context, arg []CreateCandlesParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"calculate", "candles"}, []string{"uid", "date", "open", "close", "high", "low", "volume"}, &iteratorForCreateCandles{rows: arg})
}
