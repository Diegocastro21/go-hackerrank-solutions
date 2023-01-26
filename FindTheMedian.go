import (
	"fmt"
	"sort"
)

func findMedian(arr []int32) int32{

	sort.Slice(arr, func(i, j int) bool{
		return arr[i] < arr[j]
   	})

	var mediana int32
	mitad := len(arr) / 2

	if len(arr) % 2 == 0 {
		mediana = ( arr[mitad - 1] + arr[mitad] ) / 2
	}else {
		mediana = arr[mitad]
	}
	return mediana
}
