package utils

import "../header"
import "fmt"
import "testing"

type Int header.Int

func Println(a ...interface{}) {
  fmt.Println(a)
}

func TestReadInt(t *testing.T) {
  in = header.Strin("123")
  var n Int
  ReadInt(&n)
  if n != 123 {
    t.Fail()
  }
}

func AssertEqualIntSlice(a []Int, b []Int) bool {
  if len(a) != len(b) {
    return false
  }
  for i := range a {
    if a[i] != b[i] {
      return false
    }
  }
  return true
}

func TestReadIntSlice(t *testing.T) {
  in = header.Strin("5 1 2 3 4 5")
  var n Int
  ReadInt(&n)
  if n != 5 {
    t.Fail()
  }
  a := make([]Int, n)
  for i := Int(0); i < n; i += 1 {
    ReadInt(&a[i])
  }
  expected := []Int{1, 2, 3, 4, 5}
  if ! AssertEqualIntSlice(a, expected) {
    t.Fail()
  }
}

func TestReadIntSliceMultiLines(t *testing.T) {
  input := `5
1
2
3
4
5
`
  in = header.Strin(input)
  var n Int
  ReadInt(&n)
  if n != 5 {
    t.Fail()
  }
  a := make([]Int, n)
  for i := Int(0); i < n; i += 1 {
    ReadInt(&a[i])
  }
  expected := []Int{1, 2, 3, 4, 5}
  if ! AssertEqualIntSlice(a, expected) {
    t.Fail()
  }
}

