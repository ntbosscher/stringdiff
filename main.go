package stringdiff

import (
	"strings"
)

type DiffPart struct {
	Before string
	After  string
}

func (d *DiffPart) String() string {
	if d.Before == d.After {
		return "> " + d.Before
	}

	return "- " + d.Before + "\n+ " + d.After
}

type Diff []*DiffPart

func (d Diff) String() string {
	var strs []string

	for _, item := range d {
		strs = append(strs, item.String())
	}

	return strings.Join(strs, "\n")
}

func New(a string, b string) Diff {
	aWords := strings.Split(a, " ")
	bWords := strings.Split(b, " ")

	ops := calculate(aWords, bWords)
	return simplify(ops)
}

func simplify(input []op) []*DiffPart {
	var out []*DiffPart

	before := []string{}
	after := []string{}

	for i, item := range input {
		if i > 0 {
			rowIsSame := item.Kind == same
			lastIsSame := input[i-1].Kind == same

			if rowIsSame != lastIsSame {
				out = append(out, &DiffPart{
					Before: strings.Join(before, " "),
					After:  strings.Join(after, " "),
				})

				before = nil
				after = nil
			}
		}

		switch item.Kind {
		case insert:
			after = append(after, item.Value)
		case remove:
			before = append(before, item.Value)
		case same:
			after = append(after, item.Value)
			before = append(before, item.Value)
		}
	}

	if len(before) > 0 || len(after) > 0 {
		out = append(out, &DiffPart{
			Before: strings.Join(before, " "),
			After:  strings.Join(after, " "),
		})
	}

	return out
}

type opType int

const (
	_ opType = iota
	insert
	remove
	same
)

type op struct {
	Value string
	Kind  opType
}

type rowStruct struct {
	a     []string
	b     []string
	value []op
	cost  int
}

func calculate(inputA []string, inputB []string) []op {

	searchEdge := []*rowStruct{{
		a: inputA,
		b: inputB,
	}}

	solutions := []*rowStruct{}

	for len(searchEdge) > 0 {
		nextSearchEdge := []*rowStruct{}

		for _, row := range searchEdge {
			if len(row.a) == 0 && len(row.b) == 0 {
				solutions = append(solutions, row)
				continue
			}

			// next word is identical, take it (greedy)
			if len(row.a) > 0 && len(row.b) > 0 && row.a[0] == row.b[0] {
				row.value = append(row.value, op{Value: row.a[0], Kind: same})
				row.a = row.a[1:]
				row.b = row.b[1:]
				nextSearchEdge = append(nextSearchEdge, row)
				continue
			}

			if len(row.a) == 0 {
				row.value = append(row.value, op{Value: row.b[0], Kind: insert})
				row.b = row.b[1:]
				row.cost++
				nextSearchEdge = append(nextSearchEdge, row)
				continue
			}

			if len(row.b) == 0 {
				row.value = append(row.value, op{Value: row.a[0], Kind: remove})
				row.a = row.a[1:]
				row.cost++
				nextSearchEdge = append(nextSearchEdge, row)
				continue
			}

			rowB := &rowStruct{
				a:     append([]string{}, row.a...),
				b:     append([]string{}, row.b...),
				value: append([]op{}, row.value...),
			}

			row.cost++
			row.value = append(row.value, op{Value: row.a[0], Kind: remove})
			row.a = row.a[1:]
			nextSearchEdge = append(nextSearchEdge, row)

			rowB.cost++
			rowB.value = append(rowB.value, op{Value: row.b[0], Kind: insert})
			rowB.b = rowB.b[1:]
			nextSearchEdge = append(nextSearchEdge, rowB)
		}

		searchEdge = nextSearchEdge

		if len(solutions) > 0 {
			minCost := solutions[0]
			for _, sol := range solutions {
				if sol.cost < minCost.cost {
					minCost = sol
				}
			}

			return minCost.value
		}
	}

	panic("no solutions")
}
