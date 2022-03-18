package arrayAndSlices

func Sum(n [5]int) int {
  sum := 0
  for _, num := range n {
    sum += num
  }
  return sum
}
