package main

import "fmt"

func square(n int) int{
	return n * n
}
func sum(numbers ...int) int{
	r := 0
	for _ , number := range numbers {
		r = r + number
	}
	return r
}

func main(){
	//var inhaArray [7] string //array
<<<<<<< HEAD
//	var inhaSlice [] string //slice
//	inhaSlice = make([]string, 7)
	inhaSlice := []string{"i","n","h","a" }
        nums := []int{1, 2, 3, 4, 5}
	mySlice := inhaSlice[:]
	fmt.Println(mySlice)
//	inhaSlice[0] = "i"
//	inhaSlice[1] = "n"
//	inhaSlice[2] = "h"
//	inhaSlice[3] = "a"

=======
	var inhaSlice [] string //slice
	inhaSlice = make([]string, 7)
//	inhaSlice := []string{"i","n","h","a" }
        nums := []int{1, 2, -3, 4, 5}
	inhaSlice[0] = "i"
	inhaSlice[1] = "n"
	inhaSlice[2] = "h"
	inhaSlice[3] = "a"
	fmt.Println(square(7))
>>>>>>> 22d4592252d33f28e8354f4019d64806d3225132
	for i := 0 ; i < len(inhaSlice); i++ {
		fmt.Println(inhaSlice[i])
	}
	fmt.Println(sum(1,5,3))
	fmt.Println(sum(3,1))
	fmt.Println(sum(nums...))

//	fmt.Println(1)
}
