package meander

import (
	"fmt"
	"strings"
)

type Cost int

const (
	_ Cost = iota
	Cost1
	Cost2
	Cost3
	Cost4
	Cost5
)

var costStrings = map[string]Cost{
	"$":     Cost1,
	"$$":    Cost2,
	"$$$":   Cost3,
	"$$$$":  Cost4,
	"$$$$$": Cost5,
}

func (c Cost) String() string {
	for k, v := range costStrings {
		if c == v {
			return k
		}
	}

	return "不正な値です"
}

func ParseCost(s string) Cost {
	return costStrings[s]
}

type CostRange struct {
	From Cost
	To   Cost
}

func (cr *CostRange) String() string {
	return fmt.Sprintf("%s...%s", cr.From, cr.To)
}

func ParseCostRange(s string) *CostRange {
	seg := strings.Split(s, "...")
	return &CostRange{ParseCost(seg[0]), ParseCost(seg[1])}
}
