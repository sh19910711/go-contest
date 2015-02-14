// @begin_cut
package utils
// @end_cut

func MinRune(a rune, b rune) rune {
  if a < b {
    return a
  } else {
    return b
  }
}

func MinInt(a Int, b Int) Int {
  if a < b {
    return a
  } else {
    return b
  }
}
