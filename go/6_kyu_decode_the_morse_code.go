package kata
import (
  "bytes"
  "strings"
)

func DecodeMorse(morseCode string) string {
  morseCode = strings.TrimSpace(morseCode)
  letters := strings.Split(morseCode, " ")
  var phrase bytes.Buffer

  i := 0
  for i < len(letters) - 1 {
    if (letters[i]=="" && letters[i+1]=="") {
      phrase.WriteString(" ")
      i = i + 1
    } else {
      phrase.WriteString(MORSE_CODE[letters[i]])
    }
    i = i+1
  }
  phrase.WriteString(MORSE_CODE[letters[i]])

  return phrase.String()
}
