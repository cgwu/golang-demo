package main

import (
	"bytes"
	"fmt"
)

/*
《The Go Programming Language 2015.11》P184
Sets in Go are usu ally implemente d as a map[T]bool, where T is the element typ e. A set represent
ed by a map is ver y flexible but, for cer tain pro blems, a speci alized represent ation may
outperform it. For example, in domains such as dataflow analysis where set elements are small
non-negative integers, sets have many elements, and set operations like union and intersection
are common, a bit vector is ideal.
A bit vec tor uses a slice of unsigned integer values or ‘‘words,’’ each bit of which represents a
possible element of the set. The set contains i if the i-th bit is set. The following program
demon strates a simple bit vector type with three methods:
*/

type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println("x:", x.String())

	y.Add(9)
	y.Add(42)
	y.Add(156)
	fmt.Println("y:", y.String())

	x.UnionWith(&y)
	fmt.Println("x union y:", x.String())
	fmt.Println("y:", y.String())
	fmt.Println(x.Has(9), x.Has(123))
}
