package decimal

import (
	"encoding/binary"
)

func (d *Decimal) Size() int {
	if d == nil {
		return 0
	}
	return 1
}

// MarshalBinary implements gogo proto custom MarshalTo .
func (d Decimal) MarshalTo() (data []byte, err error) {
	// Write the exponent first since it's a fixed size
	v1 := make([]byte, 4)
	binary.BigEndian.PutUint32(v1, uint32(d.exp))

	// Add the value
	var v2 []byte
	if v2, err = d.value.GobEncode(); err != nil {
		return
	}

	// Return the byte array
	data = append(v1, v2...)
	return
}
