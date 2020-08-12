package df

import "syscall"

// FromPath gather disk usage and capacity stats for the given file path
func FromPath(inPath string) (*Stats, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(inPath, &stat); err != nil {
		return nil, err
	}

	blockSize := uint64(stat.Bsize)
	capacityB := stat.Blocks * blockSize
	availableB := stat.Bavail * blockSize
	usedB := capacityB - availableB
	usedPerc := (float64(usedB) / float64(capacityB)) * 100.0
	freePerc := (float64(availableB) / float64(capacityB)) * 100.0

	return &Stats{
		AvailableB: availableB,
		UsedB:      usedB,
		CapacityB:  capacityB,
		UsedPerc:   usedPerc,
		FreePerc:   freePerc,
	}, nil
}

// MustFromPath gather disk usage and capacity stats for the given file path. Panics on error.
func MustFromPath(inPath string) Stats {
	stats, err := FromPath(inPath)
	if err != nil {
		panic(err)
	}
	return *stats
}
