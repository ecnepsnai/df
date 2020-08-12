/*
Package df provides a way to getting disk usage and capacity information for a provided
file path. It currently only supports UNIX-like operating systems.
*/
package df

import (
	"fmt"
)

// Stats describes disk usage and capacity statistics
type Stats struct {
	// The number of bytes available
	AvailableB uint64
	// The number of bytes consumed
	UsedB uint64
	// The maximum number of bytes that can be used
	CapacityB uint64
	// The percentage of bytes that has been used of the total capacity
	UsedPerc float64
	// The percentage of bytes that has not been used from the total capacity.
	FreePerc float64
}

func (t Stats) String() string {
	return fmt.Sprintf("capacity_b=%d available_b=%d used_b=%d used_p=%f free_p=%f", t.CapacityB, t.AvailableB, t.UsedB, t.UsedPerc, t.FreePerc)
}
