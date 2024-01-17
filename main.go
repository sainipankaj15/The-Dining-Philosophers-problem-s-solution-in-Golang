// The Dining Philosophers problem
package main

import (
	"fmt"
	"sync"
	"time"
)

// Defining constant
var totalRoundOfEatingPerPerson int = 2

// Defining variables
var persons = []string{"Pankaj", "Rakesh", "Mahesh"}
var wg sync.WaitGroup
var sleepTime time.Duration = 1 * time.Second
var eatingTime time.Duration = 3 * time.Second
var thinkingTime time.Duration = 1 * time.Second
var orderMutex sync.Mutex
var whoFinishedFirst = []string{}

func diningProblem(person string, leftHand, rightHand *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(person, "is seated")
	time.Sleep(sleepTime)

	// Total Rounds per Person
	for i := 1; i <= totalRoundOfEatingPerPerson; i++ {
		fmt.Println(person, "is Hungry")
		time.Sleep(sleepTime)

		// Have to lock both fork
		// Lock left fork
		leftHand.Lock()
		fmt.Println(person, "Picked Left fork")
		// Lock right fork
		rightHand.Lock()
		fmt.Println(person, "Picked Right fork")

		// Now he will eat
		fmt.Println(person, "is eating as he has both forks")
		time.Sleep(eatingTime)

		// He will think for sometime
		fmt.Println(person, "is thinking")
		time.Sleep(thinkingTime)

		// Have to unlock both fork
		// UnLock left fork
		leftHand.Unlock()
		fmt.Println(person, "put down Left fork")
		// UnLock right fork
		rightHand.Unlock()
		fmt.Println(person, "put down Right fork")

		time.Sleep(sleepTime)
	}

	fmt.Println(person, "is done with eating and he is ready to left")
	time.Sleep(sleepTime)

	fmt.Println(person, "left the table")

	orderMutex.Lock()
	whoFinishedFirst = append(whoFinishedFirst, person)
	orderMutex.Unlock()

}

func main() {
	fmt.Println("Dining Philosopher Problem solution using Mutex and Goroutine")

	// Very first we will make one left hand fork mutex to lock the fork
	leftFork := &sync.Mutex{}

	// spawn one goroutine for each person
	for _, person := range persons {

		rightFork := &sync.Mutex{}

		wg.Add(1)
		// calling a goroutine for every person/philiospher
		go diningProblem(person, leftFork, rightFork, &wg)

		// now rightFork will become leftFork for the next person so
		leftFork = rightFork

	}

	wg.Wait()

	fmt.Println("Left is Empty Now.")
	fmt.Println("Order of Person who finished first")

	for _, name := range whoFinishedFirst {
		fmt.Println(name)
	}
}
