package utils // @cut
import "fmt" // @cut
import "io" // @cut

func ReadInt(n *Int) bool {
  _, err := fmt.Fscan(in, n)
  return err != io.EOF
}

