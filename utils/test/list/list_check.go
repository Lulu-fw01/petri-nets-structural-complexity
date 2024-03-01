package list

import (
	"slices"
	"testing"
)

func CheckStringList(t *testing.T, expectedValues []string, actualValues []string) {
	for _, el := range expectedValues {
		if !slices.Contains(actualValues, el) {
			t.Fatalf("There are not al elements in list. Can't find %s", el)
		}
	}
}
