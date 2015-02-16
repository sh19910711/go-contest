package min

import "../../header"
import "testing"

type Int header.Int

func TestMinInt(t *testing.T) {
  if MinInt(1, 2) != 1 {
    t.Fail()
  }
}

func TestMinRune(t *testing.T) {
  if MinRune('1', '2') != '1' {
    t.Fail()
  }
}

