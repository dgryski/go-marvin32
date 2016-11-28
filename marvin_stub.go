// +build amd64

package marvin32

//go:generate python -m peachpy.x86_64 sum.py -S -o marvin_amd64.s -mabi=goasm
//go:noescape

func Sum32(seed uint64, data []byte) uint32
