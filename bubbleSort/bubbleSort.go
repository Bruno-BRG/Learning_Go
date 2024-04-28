package main

import "fmt"

// Função que ordena um array de inteiros utilizando o algoritmo Bubble Sort
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ { //quantidade de vezes que o array será percorrido
		for j := 0; j < len(arr)-i-1; j++ { //percorre o array e compara os elementos adjacentes
			if arr[j] > arr[j+1] { //se o elemento atual for maior que o próximo
				arr[j], arr[j+1] = arr[j+1], arr[j] //troca os elementos de posição
			}
		}
	}
	return arr
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}          //array de exemplo
	fmt.Println("Unsorted array is: ", arr)           //array desordenado
	fmt.Println("Sorted array is: ", bubbleSort(arr)) //array ordenado
}
