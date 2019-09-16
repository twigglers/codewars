package kata

import (
  "strconv"
  "strings"
)

func CountBits(input uint) int {
  binary_string := strconv.FormatInt(int64(input), 2)
  no_zeros := strings.Replace(binary_string, "0", "", -1)

  return len(no_zeros)
}