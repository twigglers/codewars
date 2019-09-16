package kata

import (
  "strconv"
  "math/big"
)

func IsPrime(n int) bool {
  return big.NewInt(int64(n)).ProbablyPrime(0)
}

func ReverseInt(n int) int {
  reversed_int := 0
  for n > 0 {
    remainder := n % 10
    reversed_int *= 10
    reversed_int += remainder
    n /= 10
  }
  return reversed_int
}

func BasicValidation(i int) bool {
  first_char := string(strconv.Itoa(i)[0])
  first_digit, _ := strconv.Atoi(first_char)
  if first_digit%2 == 0  || first_digit == 5 {
    return false
  }
  return true
}

func BackwardsPrime(start int, stop int) []int {
  var slice []int
  for i := start; i <= stop; i++ {
    reversed_i := ReverseInt(i)
    if i != reversed_i && IsPrime(i) && IsPrime(reversed_i){
      slice = append(slice, i)
    }
    /*
    if BasicValidation(i) && i != reversed_i && IsPrime(i) && IsPrime(reversed_i){
      slice = append(slice, i)
    }
    */
  }
  return slice
}