package tests

import (
	"app/internal/db/clickhouse"
	"app/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildFilterQuery(t *testing.T) {

	client := &clickhouse.ClickhouseClient{}

	// Define test cases
	tests := []struct {
		name          string
		cond          model.FilterCondition
		fields        []string
		expectedQuery string
	}{
		{
			name:          "Query with one condition",
			cond:          model.FilterCondition{Sport: "basketball"},
			fields:        []string{"sport", "age"},
			expectedQuery: "SELECT sport,age FROM events WHERE sport = @sport;",
		},
		{
			name:          "Query with no conditions",
			cond:          model.FilterCondition{},
			fields:        []string{"id", "name"},
			expectedQuery: "SELECT id,name FROM events ;",
		},
		{
			name:          "Query with multiple conditions",
			cond:          model.FilterCondition{Sport: "basketball", Gender: "male"},
			fields:        []string{"id", "name"},
			expectedQuery: "SELECT id,name FROM events WHERE gender = @gender AND sport = @sport;",
		},
	}

	// Execute tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query, _ := client.BuildFilterQuery(tt.cond, tt.fields)
			assert.Equal(t, tt.expectedQuery, query)
		})
	}
}
