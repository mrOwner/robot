package candlesgrabber_test

import (
	"archive/zip"
	"bytes"
	"context"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
	grb "github.com/mrOwner/robot/pkg/candles_grabber"
	"github.com/mrOwner/robot/util"
	"github.com/stretchr/testify/require"
)

func TestTinkoffApi(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	candles := RandomCandles()
	body, err := randomZIParchive(candles)
	if err != nil {
		t.Fatal(err)
	}

	figi := util.RandomFIGI()
	year := util.RandomYear()

	url := "https://foo.bar"
	query := map[string]string{
		"figi": figi,
		"year": strconv.Itoa(year),
	}

	httpmock.RegisterResponderWithQuery("GET", url, query, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	reader := grb.NewTinkoffApiReader("", url)
	resp, err := reader.Read(ctx, figi, year)
	require.NoError(t, err)
	require.Equal(t, httpmock.GetTotalCallCount(), 1)
	require.NotEmpty(t, resp)
	require.Len(t, resp, len(candles))

	opts := cmp.Options{
		cmp.Comparer(func(x, y time.Time) bool {
			a := x.Truncate(time.Second)
			b := y.Truncate(time.Second)
			return a.Equal(b)
		}),
	}
	require.Empty(t, cmp.Diff(candles, resp, opts))
}

func RandomCandles() []grb.Candle {
	n := util.RandomInt(1, 15)
	candles := make([]grb.Candle, n)
	for i := range candles {
		candles[i] = grb.Candle{
			UID:    uuid.New(),
			Date:   time.Now(),
			Open:   util.RandomFloat32(0.1, 9999999),
			Close:  util.RandomFloat32(0.1, 9999999),
			High:   util.RandomFloat32(0.1, 9999999),
			Low:    util.RandomFloat32(0.1, 9999999),
			Volume: util.RandomInt(0, 999999999),
		}
	}

	return candles
}

func randomZIParchive(candles []grb.Candle) ([]byte, error) {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)

	for len(candles) > 0 {
		n := util.RandomInt(1, len(candles))

		file, err := zw.Create(util.RandomString(9) + ".csv")
		if err != nil {
			return nil, err
		}

		_, err = file.Write(toCSV(candles[:n]))
		if err != nil {
			return nil, err
		}

		candles = candles[n:]
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func toCSV(candles []grb.Candle) []byte {
	bb := bytes.Buffer{}
	for _, c := range candles {
		bb.WriteString(c.UID.String())
		bb.WriteRune(';')
		bb.WriteString(c.Date.Format(time.RFC3339))
		bb.WriteRune(';')
		bb.WriteString(strconv.FormatFloat(float64(c.Open), 'f', -1, 64))
		bb.WriteRune(';')
		bb.WriteString(strconv.FormatFloat(float64(c.Close), 'f', -1, 64))
		bb.WriteRune(';')
		bb.WriteString(strconv.FormatFloat(float64(c.High), 'f', -1, 64))
		bb.WriteRune(';')
		bb.WriteString(strconv.FormatFloat(float64(c.Low), 'f', -1, 64))
		bb.WriteRune(';')
		bb.WriteString(strconv.Itoa(c.Volume))
		bb.WriteRune('\n')
	}

	return bb.Bytes()
}
