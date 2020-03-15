package kdb

import (
	"fmt"
	"strings"
)

type Grammar struct {
}

func NewGrammar() *Grammar {
	return new(Grammar)
}

func (g *Grammar) compileSelect(b *Builder) string {
	if len(b.columns) == 0 {
		b.columns = []string{"*"}
	}

	return fmt.Sprintf("select %s", strings.TrimSpace(strings.Join(g.compileComponents(b), " ")))
}

func (g *Grammar) compileComponents(b *Builder) []string {
	sql := make([]string, 0)
	if len(b.columns) > 0 {
		sql = append(sql, g.compileColumns(b))
	}

	if b.table != "" {
		sql = append(sql, g.compileFrom(b))
	}

	if len(b.wheres) > 0 {
		whereSql := g.compileWheres(b)
		if whereSql != "" {
			sql = append(sql, whereSql)
		}
	}
	return sql
}
func (g *Grammar) compileFrom(b *Builder) string {
	return fmt.Sprintf("from %s", g.wrapTable(b.table))
}

func (g *Grammar) compileColumns(b *Builder) string {
	return g.wrapColumn(b.columns...)
}

func (g *Grammar) compileWheres(b *Builder) string {
	var sql string
	for k, w := range b.wheres {
		if k == 0 {
			w.glue = ""
		}
		switch w.typ {
		case "basic":
			sql = fmt.Sprintf("%s %s %s %s %s", strings.TrimSpace(sql), w.glue, g.wrapColumn(w.column.(string)), w.operator, "?")
		case "null":
			sql = fmt.Sprintf("%s %s %s %s %s", strings.TrimSpace(sql), w.glue, g.wrapColumn(w.column.(string)), w.operator, w.value)
		}
	}
	return fmt.Sprintf("where %s", strings.TrimSpace(sql))
}

func (g *Grammar) wrapTable(table string) string {
	return fmt.Sprintf("%s%s", "", table)
}
func (g *Grammar) wrapColumn(columns ...string) string {
	for i, column := range columns {
		segments := strings.Split(column, ".")
		if len(segments) > 1 {
			segments[0] = g.wrapTable(segments[0])
			if segments[1] != "*" && !strings.Contains(segments[0], "->") {
				segments[1] = fmt.Sprintf("`%s`", segments[1])
			}
		} else {
			if segments[1] != "*" && !strings.Contains(segments[0], "->") {
				segments[1] = fmt.Sprintf("`%s`", segments[0])
			}
		}
		column = strings.Join(segments, ".")
		columns[i] = column
	}
	return fmt.Sprintf("%s", strings.Join(columns, ","))
}
