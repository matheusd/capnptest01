package capnptest01

import (
	"errors"
	"fmt"
	"slices"

	capnp "capnproto.org/go/capnp/v3"
)

// ManualSingleSegArena is an Arena implementation that stores message data
// in a continguous slice.  Allocation is performed by first allocating a
// new slice and copying existing data. SingleSegment arena does not fail
// unless the caller attempts to access another segment.
type ManualSingleSegArena []byte

// ManualSingleSegment constructs a ManualSingleSegArena from b.  b MAY be nil.
// Callers MAY use b to populate the segment for reading, or to reserve
// memory of a specific size.
func ManualSingleSegment(b []byte) *ManualSingleSegArena {
	return (*ManualSingleSegArena)(&b)
}

func (ssa ManualSingleSegArena) NumSegments() int64 {
	return 1
}

func (ssa ManualSingleSegArena) Data(id capnp.SegmentID) ([]byte, error) {
	if id != 0 {
		return nil, fmt.Errorf("segment %d requested in single segment arena", id)
	}
	return ssa, nil
}

func (ssa *ManualSingleSegArena) Allocate(sz capnp.Size, segs map[capnp.SegmentID]*capnp.Segment) (capnp.SegmentID, []byte, error) {
	// wordSize is the number of bytes in a Cap'n Proto word.
	const wordSize capnp.Size = 8

	data := []byte(*ssa)
	if segs[0] != nil {
		data = segs[0].Data()
	}
	if len(data)%int(wordSize) != 0 {
		return 0, nil, errors.New("segment size is not a multiple of word size")
	}

	// Pad to the wordsize.
	n := capnp.Size(wordSize - 1)
	sz = (sz + n) &^ n

	*ssa = slices.Grow(*ssa, int(sz))
	return 0, *ssa, nil
}

func (ssa ManualSingleSegArena) String() string {
	return fmt.Sprintf("single-segment arena [len=%d cap=%d]",
		len(ssa), cap(ssa))
}

func (ssa *ManualSingleSegArena) Release() {
	if *ssa != nil {
		*ssa = []byte(*ssa)[:0]
	}
}
