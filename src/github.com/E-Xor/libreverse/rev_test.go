package libreverse

import "testing"

func TestRev(t *testing.T) {
  cases := []struct {
    in, want string
  } {
    {"Hello", "olleH"},
    {"Hello, мир", "рим ,olleH"},
    {"", ""},
  }

  for _, c := range cases {
    got := Rev(c.in)
    if got != c.want {
      t.Errorf("Rev(%q) == %q, want %q", c.in, got, c.want)
    }
  }
}