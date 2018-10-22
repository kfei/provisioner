package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestTicker(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	var ctx, cancel = context.WithCancel(context.Background())
	var eg, _ = errgroup.WithContext(ctx)

	eg.Go(func() error {
		return runTicker(ctx)
	})

	eg.Go(func() error {
		time.Sleep(time.Second * 3)
		// Available coupons should be 30 after 3 seconds
		require.InDelta(30, available, 15.0)
		cancel()
		return nil
	})

	require.NoError(eg.Wait())
}
