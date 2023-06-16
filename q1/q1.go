package q1

//Dado um array de números inteiros "nums" e um número inteiro "target", retorne os índices dos dois números cuja soma
//seja igual a "target".
//
//Você pode assumir que cada entrada terá exatamente uma solução e não poderá usar o mesmo elemento duas vezes.
//
//Você pode retornar a resposta em qualquer ordem.

func TwoSum(nums []int, target int) []int {
	var res = []int{}

	for n1 := range nums {
		for n2 := range nums {
			if nums[n1]+nums[n2] == target {
				res = append(res, n1)
				res = append(res, n2)
			}
		}
	}

	return res
}
