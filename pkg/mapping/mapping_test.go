package mapping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mapping = &Mapping{
		XLSX: XLSX{
			GlobPattern: "spreadsheet*.xlsx",
			Worksheet: []Worksheet{
				{
					Name: "Sheet1",
					Cell: []Cell{
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
		SQL: SQL{
			Database: "db1",
			Table:    "table1",
			Attribute: []Attribute{
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
	}
)

// Test returning of a new Mapping struct.
func TestNew(t *testing.T) {
	expected := &Mapping{}
	actual := New()
	if !assert.ObjectsAreEqual(expected, actual) {
		assert.Equal(t, expected, actual)
	}
}

// Test reading of the XML mapping file.
func TestMapping_ReadXML(t *testing.T) {
	expected := mapping
	actual := new(Mapping)
	actual.ReadXML("mapping.xml")
	if !assert.ObjectsAreEqual(expected, actual) {
		assert.Equal(t, expected, actual)
	}
}

// Test reading of the XML mapping file with an invalid file path.
func TestMapping_ReadXML2(t *testing.T) {
	err := new(Mapping).ReadXML(".xml")
	if !assert.Error(t, err) {
		assert.Equal(t, nil, err)
	}
}

// Test reading of the XML mapping file with an invalid file and data format.
func TestMapping_ReadXML3(t *testing.T) {
	err := new(Mapping).ReadXML("mapping.txt")
	if !assert.Error(t, err) {
		assert.Equal(t, nil, err)
	}
}

// Test generating of a SQL insert into statement.
func TestSQL_GenerateInsertIntoStmt(t *testing.T) {
	expected := "INSERT INTO db1.table1 (attr1, attr2) VALUES ('a', 1);"
	sql := mapping.SQL
	actual := sql.GenerateInsertIntoStmt()
	assert.Equal(t, expected, actual)
}
