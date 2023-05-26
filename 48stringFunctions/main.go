package main

import (
	"fmt"
	s "strings"
)

func main() {
	var p = fmt.Println

	p("Contains:      ", s.Contains("test", "es"))
	p("Cout:          ", s.Count("test", "t"))
	p("HasPrefix:     ", s.HasPrefix("test", "te"))
	p("HasSufix:      ", s.HasSuffix("test", "st"))
	p("Index:         ", s.Index("test", "s"))
	p("Join:          ", s.Join([]string{"a", "b", "c"}, "-"))
	p("Repeate:       ", s.Repeat("a", 10))
	p("Replace:       ", s.Replace("foo", "o", "0", 1))
	p("Replace:       ", s.Replace("foo", "o", "0", -1))
	p("Split:         ", s.Split("a-b-c-d", "-"))
	p("ToUpper:       ", s.ToUpper("test"))
	p("ToLower:       ", s.ToLower("TEST"))
}
