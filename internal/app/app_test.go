package app

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mock "github.com/mrOwner/robot/mocks"
	"github.com/mrOwner/robot/util"
	"github.com/stretchr/testify/require"
)

func TestApp(t *testing.T) {
	n := 100

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	lifecyclemock := mock.NewMockLifeCycle(ctrl)
	lifecyclemock.EXPECT().Start(gomock.Any()).Times(n)

	app := New()

	for i := 0; i < n/2; i++ {
		app.AddSequential(util.RandomString(10), lifecyclemock)
		app.AddSynchronous(util.RandomString(10), lifecyclemock)
	}

	err := app.Start(ctx)
	require.NoError(t, err)

	err = app.Stop(ctx)
	require.NoError(t, err)

	time.Sleep(time.Second)
}
