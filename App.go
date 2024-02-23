package main

import "fmt"

const ebulicaoK = 373

func main() {
	ebulicaoC := ebulicaoK - 273
	fmt.Printf("O ponto de ebulição em K é: %vK.\n"+
		"E o ponto de ebulição em C é: %vC.\n", ebulicaoK, ebulicaoC)
}
