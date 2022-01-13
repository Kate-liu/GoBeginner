package main

import "testing"

func BenchmarkDirectCall(b *testing.B) {
	c := Cat{Name: "draven"}
	for n := 0; n < b.N; n++ {
		// MOVQ	AX, (SP)
		// MOVQ	$6, 8(SP)
		// CALL	"".Cat.Quack(SB)
		c.Quack()
	}
}

func BenchmarkDynamicDispatch(b *testing.B) {
	c := Duck(Cat{Name: "draven"})
	for n := 0; n < b.N; n++ {
		// MOVQ	16(SP), AX
		// MOVQ	24(SP), CX
		// MOVQ	AX, "".d+32(SP)
		// MOVQ	CX, "".d+40(SP)
		// MOVQ	"".d+32(SP), AX
		// MOVQ	24(AX), AX
		// MOVQ	"".d+40(SP), CX
		// MOVQ	CX, (SP)
		// CALL	AX
		c.Quack()
	}
}
