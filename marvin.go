// +build !amd64 noasm

package marvin32

import "math/bits"

func Sum32(seed uint64, data []byte) uint32 {

	lo := uint32(seed)
	hi := uint32(seed >> 32)

	for len(data) >= 8 {
		k1 := uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24

		lo += k1
		hi ^= lo
		lo = bits.RotateLeft32(lo, 20) + hi
		hi = bits.RotateLeft32(hi, 9) ^ lo
		lo = bits.RotateLeft32(lo, 27) + hi
		hi = bits.RotateLeft32(hi, 19)

		k1 = uint32(data[4]) | uint32(data[5])<<8 | uint32(data[6])<<16 | uint32(data[7])<<24

		lo += k1
		hi ^= lo
		lo = bits.RotateLeft32(lo, 20) + hi
		hi = bits.RotateLeft32(hi, 9) ^ lo
		lo = bits.RotateLeft32(lo, 27) + hi
		hi = bits.RotateLeft32(hi, 19)

		data = data[8:]

	}

	if len(data) >= 4 {
		k1 := uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24

		lo += k1
		hi ^= lo
		lo = bits.RotateLeft32(lo, 20) + hi
		hi = bits.RotateLeft32(hi, 9) ^ lo
		lo = bits.RotateLeft32(lo, 27) + hi
		hi = bits.RotateLeft32(hi, 19)

		data = data[4:]
	}

	/* pad the final 0-3 bytes with 0x80 */
	final := uint32(0x80)

	switch len(data) {

	case 3:
		final = (final << 8) | uint32(data[2])
		fallthrough
	case 2:
		final = (final << 8) | uint32(data[1])
		fallthrough
	case 1:
		final = (final << 8) | uint32(data[0])
	}

	lo += final
	hi ^= lo
	lo = bits.RotateLeft32(lo, 20) + hi
	hi = bits.RotateLeft32(hi, 9) ^ lo
	lo = bits.RotateLeft32(lo, 27) + hi
	hi = bits.RotateLeft32(hi, 19)

	lo += 0
	hi ^= lo
	lo = bits.RotateLeft32(lo, 20) + hi
	hi = bits.RotateLeft32(hi, 9) ^ lo
	lo = bits.RotateLeft32(lo, 27) + hi
	hi = bits.RotateLeft32(hi, 19)

	return lo ^ hi
}
