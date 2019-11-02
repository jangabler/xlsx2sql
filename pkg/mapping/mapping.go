package mapping

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

// Mapping holds all the information from the source XLSX files and the target
// SQL statements.
type Mapping struct {
	XLSX XLSX `xml:"xlsx"`
	SQL  SQL  `xml:"sql"`
}

// XLSX holds all the information for the XLSX files based on a glob pattern
// and the selected worksheets.
type XLSX struct {
	GlobPattern string      `xml:"globPattern,attr"`
	Worksheet   []Worksheet `xml:"worksheet"`
}

// Worksheet holds all the information for name of the worksheet and the
// selected cells.
type Worksheet struct {
	Name string `xml:"name,attr"`
	Cell []Cell `xml:"cell"`
}

// Cell holds all the information of the coordinate and the reference to the SQL
// table attribute in the insert into statement.
type Cell struct {
	Coordinate   string `xml:"coordinate,attr"`
	RefAttribute string `xml:"refAttribute,attr"`
}

// SQL holds all the information for the target database, table and the SQL
// table attributes in the insert into statement.
type SQL struct {
	Database  string      `xml:"database,attr"`
	Table     string      `xml:"table,attr"`
	Attribute []Attribute `xml:"attribute"`
}

// Attribute holds all the information for SQL table attribute in the insert
// into statement based on the attribute name, data type and value.
type Attribute struct {
	Name     string `xml:"name,attr"`
	DataType string `xml:"dataType,attr"`
	Value    string `xml:"value,attr"`
}

// New returns a new Mapping struct.
func New() *Mapping {
	return &Mapping{}
}

// ReadXML reads the XML mapping file and unmarshal it to the Mapping struct.
func (m *Mapping) ReadXML(path string) error {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = xml.Unmarshal([]byte(f), &m)
	if err != nil {
		return err
	}
	return err
}

// GenerateInsertIntoStmt generates SQL insert into statement of the SQL struct.
func (s *SQL) GenerateInsertIntoStmt() string {
	var (
		nameArr []string
		valArr  []string
	)
	for _, a := range s.Attribute {
		nameArr = append(nameArr, a.Name)
		valFmt := a.Value
		if strings.Compare(a.DataType, "string") == 0 {
			valFmt = fmt.Sprintf("'%s'", a.Value)
		}
		valArr = append(valArr, valFmt)
	}
	return fmt.Sprintf(
		"INSERT INTO %s.%s (%s) VALUES (%s);",
		s.Database,
		s.Table,
		strings.Join(nameArr, ", "),
		strings.Join(valArr, ", "),
	)
}
