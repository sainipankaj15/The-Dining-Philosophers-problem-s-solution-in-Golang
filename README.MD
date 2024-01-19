# Dining Philosophers Problem Solution in Go using Mutex

The Dining Philosophers problem is a classic synchronization and concurrency problem in computer science. The scenario involves five philosophers sitting around a dining table, spending their lives thinking and eating. Each philosopher has a bowl of spaghetti in front of them, and they can only perform two actions: think or eat.

## Constraints

1. To eat, a philosopher needs two forks—one in the left hand and one in the right.
2. There is one fork between each pair of adjacent philosophers.
3. A philosopher can only use the forks on their immediate left and right.

## Challenge

Design a concurrent algorithm that allows each philosopher to alternate between thinking and eating, preventing deadlock and avoiding contention for resources (forks) that could lead to other issues.

## Solution

The solution involves careful coordination of the philosophers' actions to ensure they can take turns eating without conflicting with each other. Synchronization mechanisms like semaphores or mutexes are used to manage access to the forks, ensuring that philosophers can acquire the necessary resources without causing conflicts or getting stuck in a deadlock.


## Features

- **Mutex Synchronization**: The solution employs Go's `sync.Mutex` to manage access to the forks, preventing conflicts among philosophers and avoiding potential deadlocks.


- **Console Output**: The program provides console output to visualize the actions of each philosopher, including thinking, grabbing forks, eating, and putting forks down.

## Usage

To run the program, simply execute the Go file. Adjust the number of philosophers or customize the timing parameters based on specific requirements.

```bash
go run main.go
