package df

import "fmt"

func FromPath(inPath string) (*Stats, error) {
	return nil, fmt.Errorf("unsupported platform")
}

func MustFromPath(inPath string) Stats {
	stats, err := FromPath(inPath)
	if err != nil {
		panic(err)
	}
	return *stats
}
