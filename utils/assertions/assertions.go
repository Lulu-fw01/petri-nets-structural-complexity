package assertions

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 0.000001

func AssertMetric(t *testing.T, expectedValue, actualValue float64) {
	if diff := math.Abs(expectedValue - actualValue); diff <= float64EqualityThreshold {
		return
	}
	t.Fatalf("Wrong metric, expected %f, actual: %f", expectedValue, actualValue)
}
