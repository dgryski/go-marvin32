//go:build amd64 && !noasm && !purego

package marvin32

//go:generate go run -C asm asm.go -out ../marvin_amd64.s
//go:noescape

func Sum32(seed uint64, data []byte) uint32
