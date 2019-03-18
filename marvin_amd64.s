// Code generated by command: go run asm.go -out marvin_amd64.s. DO NOT EDIT.

#include "textflag.h"

// func Sum32(seed uint64, data []byte) uint32
TEXT ·Sum32(SB), NOSPLIT, $0-36
	MOVQ seed+0(FP), CX
	MOVL CX, AX
	SHRQ $0x20, CX
	MOVL CX, CX
	MOVQ data_base+8(FP), BX
	MOVQ data_len+16(FP), BP
	CMPQ BP, $0x04
	JL   loop_end

loop_begin:
	MOVL (BX), DX
	ADDL DX, AX
	XORL AX, CX
	ROLL $0x14, AX
	ADDL CX, AX
	ROLL $0x09, CX
	XORL AX, CX
	ROLL $0x1b, AX
	ADDL CX, AX
	ROLL $0x13, CX
	ADDQ $0x04, BX
	SUBQ $0x04, BP
	CMPQ BP, $0x00000004
	JGE  loop_begin

loop_end:
	MOVL $0x00000080, DX
	CMPQ BP, $0x00
	JE   after
	CMPQ BP, $0x03
	JE   sw3
	CMPQ BP, $0x02
	JE   sw2
	CMPQ BP, $0x01
	JE   sw1

sw3:
	SHLL    $0x08, DX
	MOVBLZX 2(BX), BP
	ORL     BP, DX

sw2:
	SHLL    $0x08, DX
	MOVBLZX 1(BX), BP
	ORL     BP, DX

sw1:
	SHLL    $0x08, DX
	MOVBLZX (BX), BP
	ORL     BP, DX

after:
	ADDL DX, AX
	XORL AX, CX
	ROLL $0x14, AX
	ADDL CX, AX
	ROLL $0x09, CX
	XORL AX, CX
	ROLL $0x1b, AX
	ADDL CX, AX
	ROLL $0x13, CX
	XORL DX, DX
	ADDL DX, AX
	XORL AX, CX
	ROLL $0x14, AX
	ADDL CX, AX
	ROLL $0x09, CX
	XORL AX, CX
	ROLL $0x1b, AX
	ADDL CX, AX
	ROLL $0x13, CX
	XORL CX, AX
	MOVL AX, ret+32(FP)
	RET
