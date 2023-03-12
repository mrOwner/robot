package candlesgrabber

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mrOwner/robot/util"
)

const DateFormat = time.RFC3339 // "2006-01-02T15:04:05Z"

var (
	ErrRateLimit    = errors.New("rate limit exceed")
	ErrInvalidToken = errors.New("invalid token")
	ErrDataNotFound = errors.New("data not found")
	ErrUnknown      = errors.New("unknown error")
)

// TinkoffApiReader is a reader of historical candles from the Tinkoff API.
type TinkoffApiReader struct {
	Token string // Access token.
	URL   string // URL candles.
}

// NewTinkoffApiReader create a new tinkoff reader.
func NewTinkoffApiReader(token, URL string) Reader {
	return &TinkoffApiReader{
		Token: token,
		URL:   URL,
	}
}

// Read read historical candles by a FIGI and a year.
func (r *TinkoffApiReader) Read(ctx context.Context, figi string, year int) ([]Candle, error) {
	return r.read(ctx, figi, year)
}

func (r *TinkoffApiReader) read(ctx context.Context, figi string, year int) (result []Candle, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", r.URL, nil)
	if err != nil {
		return nil, err

	}

	bearer := "Bearer " + r.Token
	req.Header.Set("Authorization", bearer)

	q := req.URL.Query()
	q.Add("figi", figi)
	q.Add("year", strconv.Itoa(year))
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer util.JoinErrs(err, resp.Body.Close)

	switch resp.StatusCode {
	case http.StatusOK:
		// OK
	case http.StatusTooManyRequests:
		return nil, ErrRateLimit
	case http.StatusUnauthorized, http.StatusInternalServerError:
		return nil, ErrInvalidToken
	case http.StatusNotFound:
		return nil, ErrDataNotFound
	default:
		return nil, ErrUnknown
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, err
	}

	for _, file := range zipReader.File {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			data, err := readZipFile(file)
			if err != nil {
				return nil, err
			}
			candles, err := readCSVFile(data)
			if err != nil {
				return nil, err
			}
			result = append(result, candles...)
		}
	}

	return result, nil
}

func readZipFile(file *zip.File) ([]byte, error) {
	of, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer of.Close()

	return io.ReadAll(of)
}

func readCSVFile(bytes []byte) ([]Candle, error) {
	rows := strings.Split(string(bytes), "\n")
	result := make([]Candle, 0, len(rows))

	for _, row := range rows {
		if row == "" {
			continue
		}

		col := strings.Split(row, ";")

		uid, err := uuid.Parse(col[0])
		if err != nil {
			return nil, err
		}

		date, err := time.Parse(DateFormat, col[1])
		if err != nil {
			return nil, err
		}
		open, err := strconv.ParseFloat(col[2], 32)
		if err != nil {
			return nil, err
		}
		close, err := strconv.ParseFloat(col[3], 32)
		if err != nil {
			return nil, err
		}
		high, err := strconv.ParseFloat(col[4], 32)
		if err != nil {
			return nil, err
		}
		low, err := strconv.ParseFloat(col[5], 32)
		if err != nil {
			return nil, err
		}
		vol, err := strconv.Atoi(col[6])
		if err != nil {
			return nil, err
		}

		result = append(result, Candle{
			UID:    uid,
			Date:   date,
			Open:   float32(open),
			Close:  float32(close),
			High:   float32(high),
			Low:    float32(low),
			Volume: vol,
		})
	}

	return result, nil
}
