// +build windows,amd64

#include "textflag.h"
#include "funcdata.h" 

// func asmCopy(hSrc syscall.Handle, hDst syscall.Handle, n *int) error
TEXT ·asmCopy(SB), NOSPLIT, $0-32 
    MOVQ  hSrc+0(FP), SI      // src Handle
    MOVQ  hDst+8(FP), DI     // dst Handle
    MOVQ  $buf+16(FP), DX    // buffer
    MOVQ  $0, CX
    MOVQ  CX, n+16(FP)       // written bytes
    
loop: 
    MOVL  $0, AX           // read syscall
    SYSCALL
    
    CMOVQ AX, CX            // bytes read
    CMPQ  CX, $0            // eof?
    JE   done
    
    MOVL  $1, AX            // write syscall
    SYSCALL
    
    ADDQ  CX, n+16(FP)      // update written bytes
    CMPQ  CX, DX            // buffer full?
    JNE  loop               // continue
    
done:
    MOVQ  n+16(FP), AX      // written bytes
    MOVQ  $0, err+24(FP)    // nil error
ret