// Package word : 금칙어 처리 패키지.
package word

import (
	"strings"
	"unicode"
)

// FilterNode ...
type FilterNode struct {
	value   rune
	end     bool
	mapNext TNodeMap
}

// NewFilterNode ...
func NewFilterNode(r rune) *FilterNode {
	return &FilterNode{
		value:   r,
		end:     false,
		mapNext: make(TNodeMap),
	}
}

// TNodeMap ...
type TNodeMap map[rune]*FilterNode

// WordFilter ...
type WordFilter struct {
	root   *FilterNode
	size   int
	ignore map[rune]bool
}

// NewWordFilter ...
func NewWordFilter() *WordFilter {
	return &WordFilter{
		root:   NewFilterNode(0),
		ignore: make(map[rune]bool),
	}
}

func (w *WordFilter) addFilterWord(s string) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return
	}
	s = strings.ToUpper(s)

	now := w.root
	for _, r := range s {
		if node, ok := now.mapNext[r]; ok {
			now = node
		} else {
			node = NewFilterNode(r)
			now.mapNext[r] = node
			now = node
		}
	}
	now.end = true

	w.size++
}

func (w *WordFilter) addIgnoreChar(str string) {
	for _, r := range str {
		w.ignore[r] = true
	}
}

// IsInclude ...
func (w *WordFilter) IsInclude(s string) bool {
	str := []rune(s)
	size := len(str)
	for i := 0; i < size; i++ {
		now := w.root
		for j := i; j < size; j++ {
			ch := unicode.ToUpper(str[j])

			// pass ignore ch
			if _, ok := w.ignore[ch]; ok {
				continue
			}

			if node, ok := now.mapNext[ch]; ok {
				now = node
				if now.end {
					return true
				}
			} else {
				break
			}
		}
	}
	return false
}

// ToFilter ...
func (w *WordFilter) ToFilter(s string) string {
	inStr := []rune(s)
	size := len(inStr)

	outStr := inStr
	var start, end int

	for i := 0; i < size; i++ {
		start = -1
		end = -1

		now := w.root
		for j := i; j < size; j++ {
			ch := unicode.ToUpper(inStr[j])

			// pass ignore ch
			if _, ok := w.ignore[ch]; ok {
				continue
			}

			if node, ok := now.mapNext[ch]; ok {
				if start == -1 {
					start = i
				}

				now = node
				if now.end {
					end = j
				}
			} else {
				break
			}
		}

		// replace *
		if start != -1 && end != -1 {
			for t := start; t <= end; t++ {
				outStr[t] = '*'
			}
		}
	}
	return string(outStr)
}
