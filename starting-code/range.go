package main

import "fmt"

func main(){
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		fmt.Println("index:", i)
		fmt.Println("number:", num)
	}

	kvs := map[string]string{"a": "apple", "b": "string", "c": "string123"}
	for k, v := range kvs{
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go"{
		fmt.Println(i, c)
	}
}
