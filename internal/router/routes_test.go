package router

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SortByTimeAndDistance(t *testing.T) {
	routeShortest := Route{
		Destination: Location{},
		Duration:    1,
		Distance:    1,
	}
	routeLongest := Route{
		Destination: Location{},
		Duration:    5,
		Distance:    5,
	}
	routeLongTimeButShort := Route{
		Destination: Location{},
		Duration:    5,
		Distance:    1,
	}
	tests := []struct {
		routes         []Route
		expectedRoutes []Route
	}{
		{
			[]Route{
				routeLongTimeButShort,
				routeLongest,
				routeShortest,
			},
			[]Route{
				routeShortest,
				routeLongTimeButShort,
				routeLongest,
			},
		},
		{
			[]Route{
				routeLongest,
				routeLongTimeButShort,
			},
			[]Route{
				routeLongTimeButShort,
				routeLongest,
			},
		},
	}

	for _, test := range tests {
		sort.Sort(ByTimeAndDistance(test.routes))
		assert.Equal(t, test.expectedRoutes, test.routes)
	}

}
