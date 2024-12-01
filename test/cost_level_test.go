package cost_level_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/meander"
)

func TestCostValues(t *testing.T) {
	//Arrange
	costs := []int{
		int(meander.Cost1),
		int(meander.Cost2),
		int(meander.Cost3),
		int(meander.Cost4),
		int(meander.Cost5),
	}
	expec := []int{1, 2, 3, 4, 5}
	assert := assert.New(t)

	//Assert
	for i := 0; i < len(costs); i++ {
		assert.Equal(expec[i], costs[i])
	}
}

func TestParseCost(t *testing.T) {
	//Arrange
	args := []string{"$", "$$", "$$$", "$$$$", "$$$$$"}
	wants := []int{int(meander.Cost1), int(meander.Cost2), int(meander.Cost3), int(meander.Cost4), int(meander.Cost5)}
	gots := []int{}
	assert := assert.New(t)

	//Act
	for _, arg := range args {
		gots = append(gots, int(meander.ParseCost(arg)))
	}

	//Assert
	for i := 0; i < len(args); i++ {
		assert.Equal(wants[i], gots[i])
	}
}

func TestParseCostRange(t *testing.T) {
	//Arrange
	arg := "$$...$$$"
	want := &meander.CostRange{
		From: meander.ParseCost("$$"),
		To:   meander.ParseCost("$$$"),
	}

	//Act
	got := meander.ParseCostRange(arg)

	//Assert
	assert.Equal(t, want, got)
}
