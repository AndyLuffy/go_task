package main

import (
	"fmt"
	"sync"
	"time"
)

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
/*func main() {

	tasks := []Task{
		{ID: 0, Function: func() { time.Sleep(1 * time.Second) }},
		{ID: 1, Function: func() { time.Sleep(2 * time.Second) }},
		{ID: 2, Function: func() { time.Sleep(500 * time.Millisecond) }},
	}

	results := Scheduler(tasks)

	for _, result := range results {
		fmt.Printf("任务%d: 开始时间 %v, 结束时间 %v, 耗时 %v\n",
			result.TaskID,
			result.StartTime.Format("15:04:05.000"),
			result.EndTime.Format("15:04:05.000"),
			result.Duration)
	}

}*/

func goroutine() {
	go print1()
	go print2()
	time.Sleep(2 * time.Second)
}

func print1() {

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

func print2() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

type Task struct {
	ID       int
	Function func()
}

type TaskResult struct {
	TaskID    int
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
}

func Scheduler(tasks []Task) []TaskResult {
	var wg sync.WaitGroup
	results := make([]TaskResult, len(tasks))
	resultChan := make(chan TaskResult, len(tasks))

	for _, task := range tasks {
		wg.Add(1)
		go func(id int, f func()) {
			defer wg.Done()
			start := time.Now()
			f()
			end := time.Now()
			resultChan <- TaskResult{
				TaskID:    id,
				StartTime: start,
				EndTime:   end,
				Duration:  end.Sub(start),
			}
		}(task.ID, task.Function)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		results[result.TaskID] = result
	}

	return results
}
