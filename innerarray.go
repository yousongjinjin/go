package main

import "fmt"

func main(){
	slice2 := []string{"a","b","c","d","e"}
	slice3 := make([]string, 3)
	fmt.Println(slice3)
	slice2[2] = "Queen"
	slice2 = append(slice2, "akmu")
	fmt.Println(slice3)
	fmt.Println("%x\n", &slice3[2])
	fmt.Println(slice2)
	fmt.Println("%x\n", &slice2[2])
}
