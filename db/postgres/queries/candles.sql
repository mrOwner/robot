-- name: CreateCandles :copyfrom
INSERT INTO calculate.candles (
  uid,
  date,
  open,
  close,
  high,
  low,
  volume
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
);
