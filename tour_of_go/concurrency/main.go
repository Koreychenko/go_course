package main

func sum(numbers []int, c chan int) {
	s := 0

	for _, number := range numbers {
		s += number
	}

	c <- s
}

func getSum(numbers []int) int {
	c := make(chan int)

	go sum(numbers[:len(numbers)/2], c)
	go sum(numbers[len(numbers)/2:], c)

	x, y := <-c, <-c

	return x + y
}

func main() {
	x := getSum([]int{7, 2, 8, -9, 4, 0})

	print(x)
}
