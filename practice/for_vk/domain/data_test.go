package domain

import (
	"context"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestGetDataIncrementsMetrics(t *testing.T) {
	beforeCalls := testutil.ToFloat64(DataCallsTotal)

	_ = GetData(context.Background())

	afterCalls := testutil.ToFloat64(DataCallsTotal)

	if afterCalls-beforeCalls != 1 {
		t.Fatalf("expected data_calls_total to increase by 1, got delta=%v", afterCalls-beforeCalls)
	}
}
