// @begin_cut
package utils
// @end_cut
func Stdin() io.Reader {
  return os.Stdin
}

func Strin(s string) io.Reader {
  return strings.NewReader(s)
}
