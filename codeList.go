package adp

type CodeListItem struct {
	CodeValue        string `json:"codeValue,omitempty"`
	ForeignKey       string `json:"foreignKey,omitempty"`
	ShortName        string `json:"shortName,omitempty"`
	ValueDescription string `json:"valueDescription,omitempty"`
}

type CodeList struct {
	CodeListTitle string         `json:"codeListTitle,omitempty"`
	ListItems     []CodeListItem `json:"listItems,omitempty"`
}

type CodeListResponse struct {
	CodeLists []CodeList `json:"codeLists,omitempty"`
}
