package libreverse // folder name, not filename

func Rev(s string) string {
  r := []rune(s)
  l := len(r)
  for i,j := 0, l-1; i < l/2; i, j = i+1, j-1 {
    r[i], r[j] = r[j], r[i]
  }

  return string(r)
}