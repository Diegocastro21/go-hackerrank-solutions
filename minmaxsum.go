import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
    
    // sort the arr values
    sort.Slice(arr, func(i, j int) bool{
         return arr[i] < arr[j]
    })
    
    // assigned the lowest value
    var (
        min = arr[0]
        max = arr[0]
    )
    var sum int64
    sum = 0;
    
    for _, value := range arr {
        sum += int64(value)
        if max < value {
            max = value
        } else if min > value {
            min = value
        }
    }
    fmt.Printf("%d %d", (sum - int64(max)), (sum - int64(min)))

}