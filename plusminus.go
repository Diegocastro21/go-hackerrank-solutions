package plus

func plusMinus(arr []int32) {
    // Write your code here
    var positive, negative, zero int32
    length := len(arr)
    for _ , value := range arr {
        if value > 0 {
            positive++
        }else if value < 0 {
            negative++
        }else {
            zero++
        }
    }

    // casting int32 to float64 
    var (
        propPositive = float64(positive) / float64(length)
        propNegative = float64(negative) / float64(length)
        propZero = float64(zero) / float64(length)
    ) 
    
    fmt.Printf("%0.6f\n%0.6f\n%0.6f\n", propPositive, propNegative, propZero)

}