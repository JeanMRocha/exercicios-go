package bug

import "bgfmt"

func Dia(correta bool) {
	if correta {
		fmt.Printf("De dia! ☀️")
	} else {
		fmt.Printf("De noite 😢")
	}
}
