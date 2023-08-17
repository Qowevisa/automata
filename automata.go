package automata

import (
	"errors"
	"fmt"
)

type Automata struct {
	Cells        map[string]*Cell
	CellsNumber  uint
	ParsedTokens []string
}

type Rule struct {
	From  string
	Value string
	To    string
}

func CreateRule(from string, value string, to string) *Rule {
	return &Rule{
		From:  from,
		Value: value,
		To:    to,
	}
}

func CreateAutomata(rules ...*Rule) (*Automata, error) {
	var auto Automata
	auto.Cells = make(map[string]*Cell)
	for _, rule := range rules {
		_, ok := auto.Cells[rule.From]
		if !ok {
			auto.Cells[rule.From] = getCell(rule.From)
			auto.CellsNumber++
		}
		_, ok = auto.Cells[rule.To]
		if !ok {
			auto.Cells[rule.To] = getCell(rule.To)
			auto.CellsNumber++
		}
		from := auto.Cells[rule.From]
		to := auto.Cells[rule.To]
		if !from.canMove(rule.Value) {
			from.AddEdge(to, rule.Value)
		}
	}
	return &auto, nil
}

type Cell struct {
	Name  string
	Edges []Edge
}

type Edge struct {
	From  *Cell
	Value string
	To    *Cell
}

func getCell(name string) *Cell {
	return &Cell{
		Name: name,
	}
}

func (c *Cell) AddEdge(to *Cell, val string) {
	c.Edges = append(c.Edges, Edge{
		From:  c,
		To:    to,
		Value: val,
	})
}

func (c *Cell) canMove(val string) bool {
	for _, edge := range c.Edges {
		if edge.Value == val {
			return true
		}
	}
	return false
}

func (c *Cell) canMoveWithID(val string) (bool, int) {
	for i, edge := range c.Edges {
		if edge.Value == val {
			return true, i
		}
	}
	return false, 0
}

func (c *Cell) Move(val string) (*Cell, error) {
	can, id := c.canMoveWithID(val)
	if can {
		return c.Edges[id].To, nil
	}
	text := fmt.Sprintf("Can't jump from %s with %s value", c.Name, val)
	return nil, errors.New(text)
}
