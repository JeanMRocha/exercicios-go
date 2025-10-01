package exemplo

import "fmt"

func Nota(b bool) {
	fmt.Printf("Nota: ")
	if b {
		fmt.Print("Aprovado")
	} else {
		fmt.Print("Reprovado")
	}
}
