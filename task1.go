package main
import "fmt"

func subsum(a []int) (res [3]int) {
	res = [3]int{}
	var pre int
	for i := range a {
		pre = a[i]
		for j:= i+1; j < len(a); j++ {
			pre =  pre + a[j]
			if res[0] <  pre {
				res[0] = pre
				res[1] = i
				res[2] = j
			}
		}
	}
	return 
}



func main() {
       /*var x []int*/
	var n int
	fmt.Scan(&n)
	x := make([]int,n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	fmt.Println(x)
	res :=  subsum(x)
	fmt.Println("sum :", res[0], "i:", res[1],"j:", res[2])
}

