package task_2

import (
	"fmt"
)

/*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，
另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
/*func main() {
	printChan()
	time.Sleep(1 * time.Second)
}*/

func printChan() {

	var ch = make(chan int)

	go func(c chan<- int) {
		for i := 1; i < 11; i++ {
			ch <- i
			fmt.Printf("发送: %d\n", i)
		}
		close(ch)
	}(ch)

	go func(d <-chan int) {

		for v := range ch {
			fmt.Printf("接收到: %d\n", v)
		}
	}(ch)

}
