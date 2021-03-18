package main

import "fmt"

func main() {
	arr:=make([]int,5,10)
	fmt.Printf("%#v",arr)
	arr = append(arr,1)
	fmt.Printf("%#v",arr)
}