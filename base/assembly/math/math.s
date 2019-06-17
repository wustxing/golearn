#include "textflag.h"

TEXT ·Get(SB), NOSPLIT, $0-8
    MOVQ ·a(SB), AX
    MOVQ AX, ret+0(FP)
    RET


TEXT ·Add(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    MOVQ a+8(FP),BX
    ADDQ BX,AX
    MOVQ AX, ret+16(FP)
    RET

TEXT ·Sub(SB), NOSPLIT, $0-24
     MOVQ a+0(FP), AX
     MOVQ a+8(FP),BX
     SUBQ BX,AX
     MOVQ AX, ret+16(FP)
     RET

TEXT ·Mul(SB), NOSPLIT, $0-24
     MOVQ a+0(FP), AX
     MOVQ a+8(FP),BX
     IMULQ BX,AX
     MOVQ AX, ret+16(FP)
     RET




