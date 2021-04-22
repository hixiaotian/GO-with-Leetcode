package main

import "fmt"

func main(){
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for k := 7; k <= 9; k ++{
		fmt.Println(k)
	}

	for {
		fmt.Println("loop")
		break
	}
}
