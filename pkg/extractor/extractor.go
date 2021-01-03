package extractor

import (
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/jangabler/xlsx2sql/pkg/mapping"
)

var (
	openFile = excelize.OpenFile
)

// Extractor holds all insert into statements.
type Extractor struct {
	Results []mapping.SQL
}

// New returns a new Extractor struct.
func New() *Extractor {
	return &Extractor{}
}

// Run starts the extraction process.
func (e *Extractor) Run(m mapping.Mapping) {
	paths, _ := filepath.Glob(m.XLSX.GlobPattern)
	dataTypeMap := getDataTypeMap(m.SQL.Attribute)
	// Iterate over all files with there worksheets and fill Result slice of
	// Extractor struct.
	for _, p := range paths {
		xlsx, err := openFile(p)
		if err != nil {
			continue
		}
		for _, w := range m.XLSX.Worksheet {
			var attrs []mapping.Attribute
			for _, c := range w.Cell {
				val, _ := xlsx.GetCellValue(
					w.Name,
					c.Coordinate,
				)
				attrs = append(attrs, mapping.Attribute{
					Name:     c.RefAttribute,
					DataType: dataTypeMap[c.RefAttribute],
					Value:    val,
				})
			}
			e.Results = append(e.Results, mapping.SQL{
				Database:  m.SQL.Database,
				Table:     m.SQL.Table,
				Attribute: attrs,
			})
		}
	}
}

// Fill map of data types with SQL attributes of mapping.
func getDataTypeMap(attrs []mapping.Attribute) map[string]string {
	dataTypeMap := make(map[string]string)
	for _, a := range attrs {
		dataTypeMap[a.Name] = a.DataType
	}
	return dataTypeMap
}
