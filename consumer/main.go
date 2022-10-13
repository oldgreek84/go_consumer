package main

import (
	"consumer/mypackage"
	"fmt"
)

func main() {
	res := mypackage.GetTopics()
	fmt.Printf("\n%T: ", res)
	mypackage.Consume(res)

}
