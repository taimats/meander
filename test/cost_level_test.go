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

func TestCostRangeString(t *testing.T) {
	//Arange
	tests := map[string]struct {
		from    string
		to      string
		want    string
		wantErr bool
	}{
		"正常系": {
			from:    "$$",
			to:      "$$$",
			want:    "$$" + "..." + "$$$",
			wantErr: false,
		},
		"エラー": {
			from:    "$$$",
			to:      "$$",
			want:    "",
			wantErr: true,
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			//Arange
			cr := CostRange{}

			//Act
			got, err := cr.String(tt.from, tt.to)

			//Assert
			if tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
