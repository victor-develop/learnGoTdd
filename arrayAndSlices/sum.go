package arrayAndSlices

func Sum(n []int) int {
  sum := 0
  for _, num := range n {
    sum += num
  }
  return sum
}

func SumAll(numbersToSum ...[]int) []int {
  lengthOfNumbers := len(numbersToSum)
  sums := make([]int, lengthOfNumbers)

  for i, numbers := range numbersToSum {
    sums[i] = Sum(numbers)
  }
  return sums
}
