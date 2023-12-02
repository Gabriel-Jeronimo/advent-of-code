package trebuchet_test

import (
	"day-1/trebuchet"

	"testing"
)

func TestGetSum(t *testing.T) {
	expectedValue := 142
	sum, err := trebuchet.GetSum("../data/test-1")

	if err != nil || sum == -1 {
		t.Fatalf("Something went wrong, error: %v", err)
	}

	if sum != expectedValue {
		t.Fatalf("Expected value %d, instead received %d", 142, sum)
	}
}
