package kata

func InAscOrder(numbers []int) bool {
  for i := 0; i <= len(numbers) - 2; i++ {
    if numbers[i] > numbers[i+1] {
      return false
    }
  }

  return true
}