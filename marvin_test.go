package marvin32

import (
	"encoding/hex"
	"strconv"
	"testing"
)

func TestDotNetRuntime(t *testing.T) {
	// Tests from https://github.com/dotnet/runtime/blob/master/src/libraries/Common/tests/Tests/System/MarvinTests.cs

	const Seed1 uint64 = 0x4FB61A001BDBCC
	const Seed2 uint64 = 0x804FB61A001BDBCC
	const Seed3 uint64 = 0x804FB61A801BDBCC

	const TestDataString0Byte = ""
	const TestDataString1Byte = "af"
	const TestDataString2Byte = "e70f"
	const TestDataString3Byte = "37f495"
	const TestDataString4Byte = "8642dc59"
	const TestDataString5Byte = "153fb79826"
	const TestDataString6Byte = "0932e6246c47"
	const TestDataString7Byte = "ab427ea8d10fc7"

	var tests = []struct {
		seed uint64
		in   string
		out  uint64 // hi << 32 | lo
	}{
		{Seed1, TestDataString0Byte, 0x30ED35C100CD3C7D},
		{Seed1, TestDataString1Byte, 0x48E73FC77D75DDC1},
		{Seed1, TestDataString2Byte, 0xB5F6E1FC485DBFF8},
		{Seed1, TestDataString3Byte, 0xF0B07C789B8CF7E8},
		{Seed1, TestDataString4Byte, 0x7008F2E87E9CF556},
		{Seed1, TestDataString5Byte, 0xE6C08C6DA2AFA997},
		{Seed1, TestDataString6Byte, 0x6F04BF1A5EA24060},
		{Seed1, TestDataString7Byte, 0xE11847E4F0678C41},

		{Seed2, TestDataString0Byte, 0x10A9D5D3996FD65D},
		{Seed2, TestDataString1Byte, 0x68201F91960EBF91},
		{Seed2, TestDataString2Byte, 0x64B581631F6AB378},
		{Seed2, TestDataString3Byte, 0xE1F2DFA6E5131408},
		{Seed2, TestDataString4Byte, 0x36289D9654FB49F6},
		{Seed2, TestDataString5Byte, 0xA06114B13464DBD},
		{Seed2, TestDataString6Byte, 0xD6DD5E40AD1BC2ED},
		{Seed2, TestDataString7Byte, 0xE203987DBA252FB3},

		{Seed3, "00", 0xA37FB0DA2ECAE06C},
		{Seed3, "FF", 0xFECEF370701AE054},
		{Seed3, "00FF", 0xA638E75700048880},
		{Seed3, "FF00", 0xBDFB46D969730E2A},
		{Seed3, "FF00FF", 0x9D8577C0FE0D30BF},
		{Seed3, "00FF00", 0x4F9FBDDE15099497},
		{Seed3, "00FF00FF", 0x24EAA279D9A529CA},
		{Seed3, "FF00FF00", 0xD3BEC7726B057943},
		{Seed3, "FF00FF00FF", 0x920B62BBCA3E0B72},
		{Seed3, "00FF00FF00", 0x1D7DDF9DFDF3C1BF},
		{Seed3, "00FF00FF00FF", 0xEC21276A17E821A5},
		{Seed3, "FF00FF00FF00", 0x6911A53CA8C12254},
		{Seed3, "FF00FF00FF00FF", 0xFDFD187B1D3CE784},
		{Seed3, "00FF00FF00FF00", 0x71876F2EFB1B0EE8},
	}

	for _, tt := range tests {
		inp, _ := hex.DecodeString(tt.in)
		want := uint32(tt.out>>32) ^ uint32(tt.out)
		got := Sum32(tt.seed, inp)
		if got != want {
			t.Errorf("Sum32(%x, %q=%x, want %x", tt.seed, tt.in, got, want)
		}
	}
}

var res32 uint32

// 256-bytes random string
var buf = make([]byte, 8192)

var sizes = []int{8, 16, 32, 40, 60, 64, 72, 80, 100, 150, 200, 250, 512, 1024, 8192}

func BenchmarkHash(b *testing.B) {
	var r uint32

	for _, n := range sizes {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			b.SetBytes(int64(n))
			for i := 0; i < b.N; i++ {
				// record the result to prevent the compiler eliminating the function call
				r = Sum32(0, buf[:n])
			}
			// store the result to a package level variable so the compiler cannot eliminate the Benchmark itself
			res32 = r
		})
	}

}
