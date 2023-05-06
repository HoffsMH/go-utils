package kata

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(input []int) []int {
  return quickSort(input, 0, len(input)-1);
}

func quickSort(input []int, low, high int) []int {
  fmt.Println("in quickSort high and low is", high, " ", low)
  if high > low {
    partition(input, low, high)
    // quickSort(input, high, pivot + 1)
    // quickSort(input, pivot - 1, low)
  }
  return input
}

func partition(input []int, low, high int) int {
  // random new pivot
  rand.Seed(time.Now().UnixNano())
  pivot_index := rand.Intn(high - low + 1)
  fmt.Println("new pivot_index", pivot_index, high, low)
  pivot := input[pivot_index]
  i := low - 1

  // swap high and pivot
  input[high], input[pivot_index] = input[pivot_index], input[high]

  // loop through high low
  for j := low; j < high; j++ {
    if (input[j] < pivot) {
      input[i], input[j] = input[j], input[i]
      i++
    }
  }
  fmt.Println("here is input while in partition", input)

  return pivot
}
