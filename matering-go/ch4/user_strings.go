package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))
	f("%s\n", s.Title("tHis wiLL be A title!"))

	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))

	//To Upper: HELLO THERE
	//To Lower: hello there
	//THis WiLL Be A Title!
	//EqualFold: true
	//EqualFold: false

	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))
	//Prefix: true
	//Prefix: false
	//Suffix: true
	//Suffix: false

	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))
	//Index: 2
	//Index: -1

	f("Count: %v\n", s.Count("Mihalis", "i"))
	f("Count: %v\n", s.Count("Mihalis", "I"))
	f("Count: %v\n", s.Count("hahaha", "hah"))
	//Count: 2
	//Count: 0
	//Count: 1

	f("Repeat: %s\n", s.Repeat("ab", 5))
	//Repeat: ababababab

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t"))
	//TrimSpace: This is a line.
	//TrimLeft: This is a line.
	//TrimRight:      This is a        line.

	f("Compare: %v\n", s.Compare("aab", "aac"))
	f("Compare: %v\n", s.Compare("aab", "aab"))
	f("Compare: %v\n", s.Compare("aab", "aad"))
	//Compare: -1
	//Compare: 0
	//Compare: 1

	f("Fields: %v\n", s.Fields("This is a string!")) // 공백 문자 기준으로 split
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))

	f("%s\n", s.Split("abcd efg", ""))
	//Fields: [This is a string!]
	//Fields: [Thisis a string!]
	//[a b c d   e f g]

	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))
	//_a_b_c_d_ _e_f_g_
	//_a_b_c_d efg
	//_a_bcd efg

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))
	//Join: Line 1+++Line 2+++Line 3

	f("SplitAfter", s.SplitAfter("123++432++", "++"))
	// SplitAfter%!(EXTRA []string=[123++ 432++ ])

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
	// TrimFunc: abc ABC
}
