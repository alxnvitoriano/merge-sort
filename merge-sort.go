package main

import (
	"fmt"
)

// Função original de MergeSort mantida
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr //base para elemento ja ordenado
    }
    mid := len(arr) / 2
    left := mergeSort(arr[:mid]) //vai ordenar a parte da esquerda
    right := mergeSort(arr[mid:])//vai ordenar a parte da direita
    
    return merge(left, right)
}

// juntar as duas partes ordenadas em uma para mostrar o valor completo
func merge(left, right []int) []int {
    result :=[]int{}
    i, j := 0, 0

    //comparar cada parte do elemento
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }
    //apresenta os dois resultados
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)
    return result
}

func main() {
    // Teste do algoritmo de ordenação MergeSort
    arr := []int{5, 2, 8, 1, 9}
    fmt.Println("=== Algoritmo MergeSort ===")
    fmt.Println("Array antes da ordenação:", arr)
    sorted := mergeSort(arr)
    fmt.Println("Array depois da ordenação:", sorted)
    fmt.Println()
    
    // Teste do algoritmo de troco mínimo
    TesteTrocoMinimo()
    
    // Teste do algoritmo de atividades compatíveis
    TesteAtividadesCompativeis()
}