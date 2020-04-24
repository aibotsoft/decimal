package decimal

import (
	"encoding/binary"
	"math/big"
)

func (d *Decimal) Size() int {
	if d == nil {
		return 0
	}
	return 4
}

// MarshalBinary implements gogo proto custom MarshalTo .
func (d Decimal) Marshal() (data []byte, err error) {
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

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (d *Decimal) Unmarshal(data []byte) error {
	// Extract the exponent
	d.exp = int32(binary.BigEndian.Uint32(data[:4]))

	// Extract the value
	d.value = new(big.Int)
	return d.value.GobDecode(data[4:])
}

//type custom interface {
//	Marshal() ([]byte, error)
//	Unmarshal(data []byte) error
//	Size() int
//}
