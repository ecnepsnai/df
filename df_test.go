package df_test

import (
	"testing"

	"github.com/ecnepsnai/df"
)

func TestGet(t *testing.T) {
	stats, err := df.FromPath("/")
	if err != nil {
		t.Fatalf("Error getting stats of root partition: %s", err.Error())
	}

	if stats.CapacityB <= 0 {
		t.Fatalf("0 capacity for root partition?")
	}
}
