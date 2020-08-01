package main
import "fmt"
/**
7. 整数反转
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 */
func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	y := 0
	for x!=0 {
		y = y*10 + x%10
		if !( -(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}