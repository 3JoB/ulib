// +build windows,amd64

#include "textflag.h"
#include "funcdata.h" 

// func asmCopy(hSrc syscall.Handle, hDst syscall.Handle, n *int) error
TEXT ·asmCopy(SB), NOSPLIT, $0-24
    MOVQ  hSrc+0(FP), SI
    MOVQ  hDst+8(FP), DI
    MOVQ  $buf+16(FP), DX
    MOVQ  $0, CX
    MOVQ  CX, n+16(FP)

loop: 
    MOVL  $0, AX
    SYSCALL
    
    CMOVQ AX, CX
    CMPQ  CX, $0
    JE  done
    
    MOVL  $1, AX
    SYSCALL
    
    ADDQ  CX, n+16(FP)
    CMPQ  CX, DX
    JNE  loop

done:
    MOVQ  n+16(FP), AX
    MOVQ  $0, err+24(FP)
    RET