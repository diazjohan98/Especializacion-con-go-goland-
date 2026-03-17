package main

import (
	"fmt"
	"sync"
)

func main() {
	sharedCounter := 0

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10000; i++ {
			sharedCounter++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			sharedCounter++
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Expected Counter: 20000")
	fmt.Println("Actual Final Counter:", sharedCounter)
}

/*
Race Condition Explanation

1. What is a Race Condition?
A race condition is a critical flaw in concurrent programming that occurs when two or more threads (or Goroutines) access a shared piece of memory at the same time, and at least one of them is trying to modify that memory. Because the operating system's scheduler handles the execution order unpredictably, the final outcome of the program depends on the exact, random timing of these Goroutines, leading to corrupted data or unpredictable behavior.

2. How it occurs in the provided code:
In the provided Go program, we have a single variable called sharedCounter initialized to 0. We launch two separate Goroutines that both execute a loop 10,000 times. Inside the loop, both Goroutines execute the statement sharedCounter++.

While sharedCounter++ looks like a single, atomic operation in the source code, at the CPU hardware level, it is actually three distinct steps:

Read: The CPU reads the current value of sharedCounter from memory.

Modify: The CPU adds 1 to that value.

Write: The CPU writes the new value back to memory.

3. The Collision (Why it fails):
Because there is no synchronization (like a Mutex or a Channel) protecting the variable, Goroutine 1 and Goroutine 2 can read the memory at the exact same millisecond.
For example, if the counter is at 50:

Goroutine 1 reads 50.

Goroutine 2 also reads 50 before Goroutine 1 has a chance to write.

Goroutine 1 adds 1 and writes 51.

Goroutine 2 also adds 1 to its scanned value and overwrites the memory with 51.

Even though two additions were processed, the counter only went up by 1 instead of 2. One operation completely overwrote and erased the other. This interleaving happens thousands of times during the loop, which is why the final printed number is almost always unpredictably lower than the expected 20,000.

*/
