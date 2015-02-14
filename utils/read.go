// @begin_cut
package utils
// @end_cut
var sc scanner.Scanner

func ReadInt(n *Int) bool {
  tok := sc.Scan()
  if tok != scanner.EOF {
    fmt.Sscan(sc.TokenText(), n)
    return true
  } else {
    return false
  }
}

func ReadString(s *string) bool {
  tok := sc.Scan()
  if tok != scanner.EOF {
    *s = sc.TokenText()
    return true
  } else {
    return false
  }
}

