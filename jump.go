// Package jump implements Google's Jump Consistent Hash
/*

From the paper "A Fast, Minimal Memory, Consistent Hash Algorithm" by John Lamping, Eric Veach (2014).

http://arxiv.org/abs/1406.2294
*/
package jump

// Hash consistently chooses a hash bucket number in the range [0, numBuckets) for the given key. numBuckets must be >= 1.
func Hash(key uint64, numBuckets int) int32 {
	if buckets <= 0 {
		buckets = 1
	}
	var b int64 = -1
	var j int64

	for j < int64(numBuckets) {
		b = j
		key = key* uint64(2862933555777941757) + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}

	return int32(b)
}
