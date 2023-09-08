package main

import "log"

func makeMarkdownTable(columns []string, rows []string, table [][]string) string {
	n := len(rows)
	m := len(columns)
	if len(table) != n {
		log.Panicln("len(table) != n")
	}
	result := "||"
	for _, column := range columns {
		result += column + "|"
	}
	result += "\n|"
	for i := 0; i < m+1; i++ {
		result += "--|"
	}
	result += "\n"
	for index, eachRow := range table {
		result += "|" + rows[index] + "|"
		if len(eachRow) != m {
			log.Panicln("len(eachRow) != m")
		}
		for _, elem := range eachRow {
			result += elem + "|"
		}
		result += "\n"
	}
	return result
}

func MakeMarkdownForNoun(noun Noun) string {
	if !noun.IsFilled() {
		log.Panicln("noun declensions are not filled")
	}
	columns := []string{"単数", "複数"}
	rows := []string{
		"主格",
		"属格",
		"与格",
		"対格",
		"奪格",
		"呼格",
	}
	table := make([][]string, 6)
	for i := 0; i < 6; i++ {
		table[i] = make([]string, 2)
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 6; j++ {
			table[j][i] = *noun.Declensions[i][j]
		}
	}
	header := `---
title: 名詞 ` + *noun.Declensions[0][0] + `
icon: ../latin.ico
---
`
	return header + makeMarkdownTable(columns, rows, table)
}
