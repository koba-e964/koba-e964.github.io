package main

import (
	"testing"

	"github.com/bayashi/actually"
)

func TestMakeMarkdownTable(t *testing.T) {
	columns := []string{"a", "b", "c"}
	rows := []string{"1", "2"}
	table := [][]string{
		{"a1", "b1", "c1"},
		{"a2", "b2", "c2"},
	}
	expected := `||a|b|c|
|--|--|--|--|
|1|a1|b1|c1|
|2|a2|b2|c2|
`
	got := makeMarkdownTable(columns, rows, table)
	actually.Expect(expected).Got(got).Same(t)
}
