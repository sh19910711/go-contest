package max

import "testing"

type Int int64

func TestMaxInt(t *testing.T) {
  if MaxInt(Int(1), Int(2)) != Int(2) {
    t.Fail()
  }
}

func TestMaxRune(t *testing.T) {
  if MaxRune('1', '2') != '2' {
    t.Fail()
  }
}

