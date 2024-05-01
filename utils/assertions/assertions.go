package assertions

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 0.000001

// todo refactor or use convey.
func AssertMetric(t *testing.T, expectedValue, actualValue float64) {
	if diff := math.Abs(expectedValue - actualValue); diff <= float64EqualityThreshold {
		return
	}
	t.Fatalf("Wrong metric, expected %f, actual: %f", expectedValue, actualValue)
}

func IsCorrect(expectedValue, actualValue float64) bool {
	if expectedValue == actualValue {
		return true
	}
	const tolerance = 1e-12
	d := math.Abs(expectedValue - actualValue)
	if actualValue == 0 {
		return d < tolerance
	}
	return (d / math.Abs(actualValue)) < tolerance
}
