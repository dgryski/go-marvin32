// +build ignore

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
)

func update(lo, hi, v Op) {

	if r, ok := v.(Register); ok {
		ADDL(r, lo)
	}
	XORL(lo, hi)

	ROLL(Imm(20), lo)
	ADDL(hi, lo)

	ROLL(Imm(9), hi)
	XORL(lo, hi)

	ROLL(Imm(27), lo)
	ADDL(hi, lo)

	ROLL(Imm(19), hi)
}

func main() {
	Package("github.com/dgryski/go-marvin32")
	TEXT("Sum32", NOSPLIT, "func(seed uint64, data []byte) uint32")
	Doc("Sum32 computes the Marvin32 hash of the key")

	reg_lo := GP32()
	reg_hi := GP32()
	reg_k1 := GP32()

	t := GP64()
	Load(Param("seed"), t)

	MOVL(t.As32(), reg_lo)
	SHRQ(Imm(32), t)
	MOVL(t.As32(), reg_hi)

	reg_d := Load(Param("data").Base(), GP64())
	reg_d_len := Load(Param("data").Len(), GP64())

	loop_begin := "loop_begin"
	loop_end := "loop_end"

	CMPQ(reg_d_len, Imm(4))
	JL(LabelRef(loop_end))
	Label(loop_begin)
	MOVL(Mem{Base: reg_d}, reg_k1)
	update(reg_lo, reg_hi, reg_k1)
	ADDQ(Imm(4), reg_d)
	SUBQ(Imm(4), reg_d_len)
	CMPQ(reg_d_len, Imm(4))
	JGE(LabelRef(loop_begin))
	Label(loop_end)

	MOVL(U32(0x80), reg_k1)

	// no support for jump tables
	after := "after"
	sw3 := "sw3"
	sw2 := "sw2"
	sw1 := "sw1"

	CMPQ(reg_d_len, Imm(0))
	JE(LabelRef(after))

	CMPQ(reg_d_len, Imm(0x03))
	JE(LabelRef(sw3))

	CMPQ(reg_d_len, Imm(0x02))
	JE(LabelRef(sw2))

	CMPQ(reg_d_len, Imm(0x01))
	JE(LabelRef(sw1))

	reg_b := GP32()

	Label(sw3)
	SHLL(Imm(8), reg_k1)
	MOVBLZX(Mem{Base: reg_d, Disp: 2}, reg_b)
	ORL(reg_b, reg_k1)

	Label(sw2)
	SHLL(Imm(8), reg_k1)
	MOVBLZX(Mem{Base: reg_d, Disp: 1}, reg_b)
	ORL(reg_b, reg_k1)

	Label(sw1)
	SHLL(Imm(8), reg_k1)
	MOVBLZX(Mem{Base: reg_d, Disp: 0}, reg_b)
	ORL(reg_b, reg_k1)

	Label(after)

	update(reg_lo, reg_hi, reg_k1)
	XORL(reg_k1, reg_k1)
	update(reg_lo, reg_hi, Imm(0))

	XORL(reg_hi, reg_lo)

	Store(reg_lo, ReturnIndex(0))
	RET()

	Generate()
}
