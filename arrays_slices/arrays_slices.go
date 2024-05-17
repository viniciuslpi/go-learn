package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string

	array1[0] = "Pos 1"

	array2 := [5]string{"pos 1", "pos 2", "pos 3", "pos 4", "pos 5"}

	array3 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(array1, array2, array3)

	slice := []int{10, 321, 12, 2, 45, 6}

	fmt.Println(slice)

}
