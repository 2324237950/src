package main

import "fmt"

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", HelloServer)
// 	_ = http.ListenAndServe(":8080", nil)
// }

// func HelloServer(w http.ResponseWriter, r *http.Request) {
// 	_, err := w.Write([]byte(`hello world`))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Print(s[len(s)/2:])
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}
