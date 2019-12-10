package kdb

import "github.com/0990/golearn/go-patterns/builder/car"

type Builder struct {
	table   string
	conn    *Connection
	grammar *Grammar

	bindings map[string][]interface{}
	columns  []string
	wheres   []where
}

type where struct {
	typ      string
	column   interface{}
	operator string
	value    interface{}
	glue     string
}

func newBuilder(conn *Connection, grammar *Grammar) *Builder {
	b := new(Builder)
	b.conn = conn
	b.grammar = grammar
	b.bindings = make(map[string][]interface{})
	return b
}

func (b *Builder) Table(table string) *Builder {
	b.table = table
	return b
}

func (b *Builder) Select(columns ...string) *Builder {
	if len(columns) == 0 {
		columns = append(columns, "*")
	}
	b.columns = columns
	return b
}

func (b *Builder) Where(column interface{}, args ...interface{}) *Builder {
	if len(args) == 0 {
		return b.WhereIsNull(column)
	}

	w := new(where)
	w.column = column
	w.glue = "and"
	w.typ = "basic"

	switch len(args) {
	case 1:
		w.operator = "="
		w.value = args[0]
	case 2:
		w.operator = args[0].(string)
		w.value = args[1]
	case 3:
		w.operator = args[0].(string)
		w.value = args[1]
		w.glue = args[2].(string)
	case 4:
		w.operator = args[0].(string)
		w.value = args[1]
		w.glue = args[2].(string)
		w.typ = args[3].(string)
	}

	b.addBinding("where", []interface{}{w.value})
	b.wheres = append(b.wheres, *w)
	return b
}

func (b *Builder) WhereisNull(column interface{}) *Builder {
	w := new(where)
	w.column = column
	w.glue = "and"
	w.typ = "null"
	w.operator = "is"
	w.value = "null"
	b.wheres = append(b.wheres, *w)
	return b
}

func (b *Builder) Get(columns ...string) *Rows {
	if len(columns) > 0 {
		b.Select(columns...)
	}
	return b.runSelect()
}

func (b *Builder) addBinding(typ string, value []interface{}) {
	if _, ok := b.bindings[typ]; ok {
		b.bindings[typ] = append(b.bindings[typ], value...)
	} else {
		b.bindings[typ] = value
	}
}

func (b *Builder) getBindings() (bindings []interface{}) {
	bindings = make([]interface{}, 0)

	if v, ok := b.bindings["where"]; ok {
		bindings = append(bindings, v...)
	}
	return
}

func(b *Builder)runSelect()*Rows{
	return b.conn.Select(b.toSQL(),b.getBindings())
}

func(b *Builder)toSQL()string{
	ruetnr b.grammar.comp
}
