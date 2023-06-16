package q5

import "strings"

//Uma frase é um palíndromo se, após converter todas as letras maiúsculas em letras minúsculas e remover todos os
//caracteres não alfanuméricos, ela for lida da mesma forma da esquerda para a direita e vice-versa. Caracteres
//alfanuméricos incluem letras e números.
//
//Dada uma string "s", retorne verdadeiro se for um palíndromo e falso caso contrário.

func IsPalindrome(s string) bool {
	strings.ReplaceAll(s, " ", "")
	strings.ReplaceAll(s, ",", "")
	strings.ReplaceAll(s, ":", "")
	sm := strings.ToLower(s)

	var s2 string

	for l := len(sm) - 1; l >= 0; l-- {
		s2 += string(sm[l])
	}

	if sm == s2 {
		return true
	} else {
		return false
	}
}
