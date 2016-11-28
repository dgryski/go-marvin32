package marvin32

import (
	"testing"
)

// This table was generated from reference C implementation
var goldenMarvin = []struct {
	out uint32
	in  string
}{
	{0xf7f2c954, ""},
	{0xd46e71f7, "a"},
	{0xb40c651c, "ab"},
	{0x5b3bc23d, "abc"},
	{0x6b15e57b, "abcd"},
	{0x601e6ea8, "abcde"},
	{0xfc18bd2c, "abcdef"},
	{0x79b01bfb, "abcdefg"},
	{0x54793238, "abcdefgh"},
	{0xebf98191, "abcdefghi"},
	{0x68a8001d, "abcdefghij"},
	{0x659105c1, "Discard medicine more than two years old."},
	{0xb98b31d, "He who has a shady past knows that nice guys finish last."},
	{0xbae17c9a, "I wouldn't marry him with a ten foot pole."},
	{0x9a299f69, "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave"},
	{0xb463d704, "The days of the digital watch are numbered.  -Tom Stoppard"},
	{0xe6059c5f, "Nepal premier won't resign."},
	{0xbdd4f772, "For every action there is an equal and opposite government program."},
	{0x12af7ede, "His money is twice tainted: 'taint yours and 'taint mine."},
	{0x1e9cae8, "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977"},
	{0xcb683e33, "It's a tiny change to the code and not completely disgusting. - Bob Manchek"},
	{0x2074fbfa, "size:  a.out:  bad magic"},
	{0x52abb615, "The major problem is with sendmail.  -Mark Horton"},
	{0x5a509711, "Give me a rock, paper and scissors and I will move the world.  CCFestoon"},
	{0xf97f5273, "If the enemy is within range, then so are you."},
	{0x494c0cb, "It's well we cannot hear the screams/That we create in others' dreams."},
	{0x7150a3c0, "You remind me of a TV show, but that's all right: I watch it anyway."},
	{0xc5f56430, "C is as portable as Stonehedge!!"},
	{0x712bcf01, "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley"},
	{0xedd44de6, "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule"},
	{0xd9440105, "How can you write a big system without C++?  -Paul Glick"},
}

func TestMarvin(t *testing.T) {
	for _, g := range goldenMarvin {
		sum := Sum32(0x5D70D359C498B3F8, []byte(g.in))
		if sum != g.out {
			t.Errorf("Sum32(%s) = 0x%x want 0x%x", g.in, sum, g.out)
		}
	}
}
