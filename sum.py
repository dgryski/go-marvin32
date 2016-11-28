import peachpy.x86_64

seed = Argument(uint64_t)
data_base = Argument(ptr())
data_len = Argument(int64_t)
data_cap = Argument(int64_t)

def update(lo, hi, v):
    ADD(lo, v)
    XOR(hi, lo)

    ROL(lo, 20)
    ADD(lo, hi)

    ROL(hi, 9)
    XOR(hi, lo)

    ROL(lo, 27)
    ADD(lo, hi)

    ROL(hi, 19)

with Function("Sum32", (seed, data_base, data_len, data_cap), uint32_t, target=uarch.default) as function:

    reg_lo = GeneralPurposeRegister32()
    reg_hi = GeneralPurposeRegister32()
    reg_d = GeneralPurposeRegister64()
    reg_d_len = GeneralPurposeRegister64()

    reg_k1 = GeneralPurposeRegister32()

    LOAD.ARGUMENT(reg_d, seed)
    MOV(reg_lo, reg_d.as_dword)
    SHR(reg_d, 32)
    MOV(reg_hi, reg_d.as_dword)

    LOAD.ARGUMENT(reg_d, data_base)
    LOAD.ARGUMENT(reg_d_len, data_len)

    loop = Loop()
    CMP(reg_d_len, 4)
    JL(loop.end)
    with loop:
        MOV(reg_k1, dword[reg_d])
        update(reg_lo, reg_hi, reg_k1)
        ADD(reg_d, 4)
        SUB(reg_d_len, 4)
        CMP(reg_d_len, 4)
        JGE(loop.begin)

    MOV(reg_k1, 0x80)

    # no support for jump tables
    after = Label("after")
    sw3 = Label("sw3")
    sw2 = Label("sw2")
    sw1 = Label("sw1")

    CMP(reg_d_len, 0x00)
    JE(after)

    CMP(reg_d_len, 0x03)
    JE(sw3)

    CMP(reg_d_len, 0x02)
    JE(sw2)

    CMP(reg_d_len, 0x01)
    JE(sw1)

    reg_b = GeneralPurposeRegister32()

    LABEL(sw3)
    SHL(reg_k1, 8)
    MOVZX(reg_b, byte[reg_d+2])
    OR(reg_k1, reg_b)

    LABEL(sw2)
    SHL(reg_k1, 8)
    MOVZX(reg_b, byte[reg_d+1])
    OR(reg_k1, reg_b)

    LABEL(sw1)
    SHL(reg_k1, 8)
    MOVZX(reg_b, byte[reg_d])
    OR(reg_k1, reg_b)

    LABEL(after)

    update(reg_lo, reg_hi, reg_k1)
    update(reg_lo, reg_hi, 0)

    XOR(reg_lo, reg_hi)

    RETURN(reg_lo)
