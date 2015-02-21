package utils // @cut
import "fmt" // @cut
import "io" // @cut
func Read(a ...interface{}) bool {
  _, err := fmt.Fscan(in, a...)
  return err != io.EOF
}
