package main

import (
	"fmt"
	"runtime"
)

// ===========Go Routine Without WaitGroup===========

func doTask1() {
	fmt.Println("Task 1 done...")
}

func doTask2() {
	fmt.Println("Task 2 done...")
}

func doTask3() {
	fmt.Println("Task 3 done...")
}

func main() {
	fmt.Println("Before runtime.NumGoroutine()", runtime.NumGoroutine())

	go doTask1()
	go doTask2()
	go doTask3()

	fmt.Println("After runtime.NumGoroutine()", runtime.NumGoroutine())
}

// ===========Go Routine With WaitGroup===========
/**
var wg sync.WaitGroup

func doTask1() {
	fmt.Println("Task 1 done...")
	wg.Done()
}

func doTask2() {
	fmt.Println("Task 2 done...")
	wg.Done()
}

func doTask3() {
	fmt.Println("Task 3 done...")
	wg.Done()
}

func main() {
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())

	fmt.Println("Start game...")
	wg.Add(3)

	go doTask1()
	fmt.Println("Start doTask1")
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())

	go doTask2()
	fmt.Println("Start doTask2")
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())

	go doTask3()
	fmt.Println("Start doTask3")
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("End game...")
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())
}
**/

// ===========Go Routine With 2 WaitGroup===========
/**


var wg sync.WaitGroup

func doTask1() {
	fmt.Println("Task 1 done...")
	wg.Done()
}

func doTask2() {
	fmt.Println("Task 2 done...")
	wg.Done()
}

func doTask3() {
	fmt.Println("Task 3 done...")
	wg.Done()
}

func doTask4() {
	fmt.Println("Task 4 done...")
	wg.Done()
}

func doTask5() {
	fmt.Println("Task 5 done...")
	wg.Done()
}

func doTask6() {
	fmt.Println("Task 6 done...")
	wg.Done()
}

func main() {
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())

	fmt.Println("Start Group 1...")
	wg.Add(3)

	go doTask1()
	fmt.Println("Start doTask1")
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())

	go doTask2()
	fmt.Println("Start doTask2")
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())

	go doTask3()
	fmt.Println("Start doTask3")
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("End Group 1...")
	fmt.Println("Group 1 runtime.NumGoroutine()", runtime.NumGoroutine())

	// Start Group 2

	wg.Add(3)

	fmt.Println("=====================================================")
	fmt.Println("Start Group 2...")
	go doTask4()
	fmt.Println("Start doTask4")
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())

	go doTask5()
	fmt.Println("Start doTask5")
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())

	go doTask6()
	fmt.Println("Start doTask6")
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("End Group 2...")
	fmt.Println("Group 2 runtime.NumGoroutine()", runtime.NumGoroutine())
}
**/
