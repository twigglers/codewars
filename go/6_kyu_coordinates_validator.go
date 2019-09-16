package kata

import (
  "strings"
  "strconv"
)

func IsValidCoordinates(coordinates string) bool {
  slice := strings.Split(coordinates, ",")
  if len(slice) != 2 || strings.Contains(coordinates, "e") {
    return false
  }

  if strings.HasPrefix(slice[1], " ") {
    slice[1] = slice[1][1:]
  } else {
    return false
  }

  var float_lat, float_long float64
  var err_lat, err_long error
  float_lat, err_lat = strconv.ParseFloat(slice[0], 64)
  float_long, err_long = strconv.ParseFloat(slice[1], 64)

  if err_lat != nil || err_long != nil || float_lat <= -90 || float_lat >= 90 || float_long <= -180 || float_long >= 180 {
    return false
  }

  /*
  if float_lat, err_lat := strconv.ParseFloat(slice[0], 64); err_lat != nil || float_lat < -90 || float_lat > 90 {
    return false
  }
  if float_long, err_long := strconv.ParseFloat(slice[1], 64); err_long != nil || float_long < -180 || float_long > 180 {
    return false
  }
  */

  return true
}