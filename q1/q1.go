package q1

//Dado um array de números inteiros "nums" e um número inteiro "target", retorne os índices dos dois números cuja soma
//seja igual a "target".
//
//Você pode assumir que cada entrada terá exatamente uma solução e não poderá usar o mesmo elemento duas vezes.
//
//Você pode retornar a resposta em qualquer ordem.

func TwoSum(nums []int, target int) []int {
	result := []int{}

	for index, valor := range nums {
		for index_2, valor_2 := range nums {
			if valor+valor_2 == target {
				result = append(result, index, index_2)
				break
			}
		}
		if len(result) != 0 {
			break
		}

	}
	return result
}
