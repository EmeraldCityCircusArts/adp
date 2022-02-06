package adp

type TableColumn struct {
	ColumnName           string `json:"columnName,omitempty"`
	PrimarykeyIndicator  bool   `json:"primarykeyIndicator,omitempty"`
	UserVisibleIndicator bool   `json:"userVisibleIndicator,omitempty"`
}

type TableHeader struct {
	Columns         []TableColumn `json:"columns,omitempty"`
	FormatTypeCodes []string      `json:"formatTypeCodes,omitempty"`
	NumberOfColumns int           `json:"numberOfColumns,omitempty"`
	TitleName       string        `json:"titleName,omitempty"`
}

type TableRow struct {
	ColumnValues    []string `json:"columnValues,omitempty"`
	FormatTypeCodes []string `json:"formatTypeCodes,omitempty"`
}

type Table struct {
	Header    TableHeader `json:"header,omitempty"`
	TableName string      `json:"tableName,omitempty"`
	TableRows []TableRow  `json:"tableRows,omitempty"`
}

type TableResponse struct {
	Table `json:"table,omitempty"`
}
