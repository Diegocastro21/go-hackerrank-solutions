import (
	"fmt"
	"time"
)
func timeConversion(s string) string {
    // Write your code here
    layout := "03:04:05PM"
    t, err := time.Parse(layout, s)

    if err != nil {
        return fmt.Sprintln(err)
    }
    
    return t.Format("15:04:05")
}