// +build !amd64 noasm

package marvin32

type marvin struct {
	lo, hi uint32
}

// rotate x left by r bits
func rotl32(x uint32, r uint32) uint32 {
	return (x << r) | (x >> (32 - r))
}

func (m *marvin) update(v uint32) {
	m.lo += v
	m.hi ^= m.lo
	m.lo = rotl32(m.lo, 20) + m.hi
	m.hi = rotl32(m.hi, 9) ^ m.lo
	m.lo = rotl32(m.lo, 27) + m.hi
	m.hi = rotl32(m.hi, 19)
}

func Sum32(seed uint64, data []byte) uint32 {
	var m marvin

	m.lo = uint32(seed)
	m.hi = uint32(seed >> 32)

	for len(data) >= 4 {
		k1 := uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
		m.update(k1)
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

	m.update(final)
	m.update(0)

	return m.lo ^ m.hi
}
