package statistics

import (
	"errors"
)

func Media(list []int) (int, error) {
	if len(list) == 0 {
		return 0, errors.New("the slice is empty")
	}
	acc := 0
	for _, v := range list {
		acc += v
	}
	return acc / len(list), nil

}

func Mediana(list []int) (int, error) {
	if len(list) == 0 {
		return 0, errors.New("the slice is empty")
	}
	//se é par, retorna a média dos dois elementos centrais
	if len(list)%2 == 0 {
		center := list[(len(list) / 2) : (len(list)/2)+1]
		return Media(center)
	}

	return Media(list[(len(list) / 2):list[(len(list)/2)]])
}

func Moda(list []int) (int, int, error) {
	if len(list) == 0 {
		return 0, 0, errors.New("the slice is empty")
	}
	mapa := make(map[int]int, 0)
	for _, v := range list {
		mapa[v] += 1
	}

	chave, valor := 0, 0
	for c, v := range mapa {
		if mapa[c] > valor {
			valor = v
			chave = c
		}
	}

	return chave, valor, nil
}
