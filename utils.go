package cache

import "math/bits"

// nextPowerOf2 returns the next power of 2 greater than or equal to n
func nextPowerOf2(n int) int {
	if n <= 1 {
		return 1
	}
	return 1 << bits.Len(uint(n-1))
}

func defaultPathExtractor(key string) string {
	// Default implementation assumes the key contains the URL path
	// For hashed keys, this won't work well - users should provide custom extractor
	return ""
}
