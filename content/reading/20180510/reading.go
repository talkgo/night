package main

import "unsafe"

type hmap struct {
	// Note: the format of the Hmap is encoded in ../../cmd/internal/gc/reflect.go and
	// ../reflect/type.go. Don't change this structure without also changing that code!
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	// extra *mapextra // optional fields
}

// Switch to indexShortStr when IndexByte produces too many false positives.
// Too many means more that 1 error per 8 characters.
// Allow some errors in the beginning.
func main() {
	// ... 省略上下文，为了不报提示，所以直接注释掉
	// if fails > (i+16)/8 {
	// 	r := indexShortStr(s[i:], substr)
	// 	if r >= 0 {
	// 		return r + i
	// 	}
	// 	return -1
	// }
}
