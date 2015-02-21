package header // @cut
import "fmt"
import "io"
import "os"
import "strings"

func Println(args ...interface{}) {
  fmt.Println(args...)
}

func Stdin() io.Reader {
  return os.Stdin
}

func Strin(s string) io.Reader {
  return strings.NewReader(s)
}

