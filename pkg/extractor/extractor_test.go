package extractor

import (
	"testing"

	"github.com/jangabler/xlsx2sql/pkg/mapping"
	"github.com/stretchr/testify/assert"
)

// Test returning of a new Extractor struct.
func TestNew(t *testing.T) {
	expected := &Extractor{}
	actual := New()
	if !assert.ObjectsAreEqual(expected, actual) {
		assert.Equal(t, expected, actual)
	}
}

func TestExtractor_Run(t *testing.T) {
	expected := &Extractor{
		Results: []mapping.SQL{
			{
				Database: "db1",
				Table:    "table1",
				Attribute: []mapping.Attribute{
					{
						Name:     "attr1",
						DataType: "string",
						Value:    "a",
					},
					{
						Name:     "attr2",
						DataType: "integer",
						Value:    "1",
					},
				},
			},
			{
				Database: "db1",
				Table:    "table1",
				Attribute: []mapping.Attribute{
					{
						Name:     "attr1",
						DataType: "string",
						Value:    "b",
					},
					{
						Name:     "attr2",
						DataType: "integer",
						Value:    "2",
					},
				},
			},
		},
	}
	m := mapping.Mapping{
		XLSX: mapping.XLSX{
			GlobPattern: "spreadsheet*.xlsx",
			Worksheet: []mapping.Worksheet{
				{
					Name: "Sheet1",
					Cell: []mapping.Cell{
						{
							Coordinate:   "A1",
							RefAttribute: "attr1",
						},
						{
							Coordinate:   "B1",
							RefAttribute: "attr2",
						},
					},
				},
			},
		},
		SQL: mapping.SQL{
			Database: "db1",
			Table:    "table1",
			Attribute: []mapping.Attribute{
				{
					Name:     "attr1",
					DataType: "string",
				},
				{
					Name:     "attr2",
					DataType: "integer",
				},
			},
		},
	}
	actual := new(Extractor)
	actual.Run(m)
	if !assert.ObjectsAreEqual(expected, actual) {
		assert.Equal(t, expected, actual)
	}
}

func TestGetDataTypeMap(t *testing.T) {
	expected := map[string]string{
		"attr1": "string",
		"attr2": "integer",
	}
	actual := getDataTypeMap([]mapping.Attribute{
		{
			Name:     "attr1",
			DataType: "string",
		},
		{
			Name:     "attr2",
			DataType: "integer",
		},
	})
	if !assert.ObjectsAreEqual(expected, actual) {
		assert.Equal(t, expected, actual)
	}
}
