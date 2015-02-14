// @begin_cut
package utils
// @end_cut
var sc scanner.Scanner

func ReadInt() (Int, error) {
  var res Int
  tok := sc.Scan()
  if tok != scanner.EOF {
    fmt.Sscan(sc.TokenText(), &res)
    return res, nil
  } else {
    return -1, io.EOF
  }
}

func ReadString() (string, error) {
  var res string
  tok := sc.Scan()
  if tok != scanner.EOF {
    res = sc.TokenText()
    return res, nil
  } else {
    return "", nil
  }
}

