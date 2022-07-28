package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
goroutine ler arası haberleşmeyi sağlar.

<- ile kullanılır.

threadların haberleşmesidir.

kanallar goroutine ler arası gezdirip elde edilir.

*/

// func sum(s []int, c chan int){
// 	sum := 0
// 	for _,v := range s{
// 		sum += v
// 	}
// 	c <- sum //chanel'a sum degeri yazılır.
// }

// func main(){
// 	s:= []int{1,2,3,4,5,6}

// 	c:= make(chan int)

// 	go sum(s[:len(s)/2], c)

// 	x := <-c //chanel'dan deger alınır.

// 	fmt.Println(x)
// }


func main(){
	runtime.GOMAXPROCS(8)
	ch := make(chan string)

	go xFunc(ch)
	go print(ch)
	time.Sleep(100 * time.Millisecond)

}

func xFunc(ch chan string){
	for l := byte('a'); l<= byte('z'); l++{
		ch <- string(l)
	}
}

//islemin yapilmasi

func printer(ch chan string){
	for{
		fmt.Println(<-ch)
	}
}

//islemin yazdirilmasi