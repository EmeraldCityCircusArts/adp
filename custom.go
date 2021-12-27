package adp

import "time"

// "description": "The string reprersentation of the date-time value. Follows the ISO-8601:2000 format: 2008-05-11T15:30:00-06:00."
type DateTimeType_v01 string

// String returns the DateTimeType_v01 formatted using the format string
func (d DateTimeType_v01) String() string {
	return string(d)
}

// Time parses a DateTimeType_v01 and returns the time value it represents.
func (d DateTimeType_v01) Time() (t time.Time) {
	t, err := time.Parse("2006-01-02T15:04:05-07:00", string(d))
	if err != nil {
		panic(err)
	}
	return
}

// "description": "The string representation of the date value. Follows the ISO-8601:2000 format: 2008-05-11",
type DateType_v01 string

// String returns the DateType_v01 formatted using the format string
func (d DateType_v01) String() string {
	return string(d)
}

// Time parses a DateType_v01 and returns the time value it represents.
func (d DateType_v01) Time() (t time.Time) {
	t, err := time.Parse("2006-01-02", string(d))
	if err != nil {
		panic(err)
	}
	return
}

type GenericAddress struct {
	AttentionOfName          string                     `json:"attentionOfName,omitempty"`
	BlockName                string                     `json:"blockName,omitempty"`
	BuildingName             string                     `json:"buildingName,omitempty"`
	BuildingNumber           string                     `json:"buildingNumber,omitempty"`
	CareOfName               string                     `json:"careOfName,omitempty"`
	CityName                 string                     `json:"cityName,omitempty"`
	CountryCode              SimpleCodeType_v02         `json:"countryCode,omitempty"`
	CountrySubdivisionLevel1 CountrySubdivisionType_v02 `json:"countrySubdivisionLevel1,omitempty"`
	CountrySubdivisionLevel2 CountrySubdivisionType_v02 `json:"countrySubdivisionLevel2,omitempty"`
	DeliveryPoint            string                     `json:"deliveryPoint,omitempty"`
	Door                     string                     `json:"door,omitempty"`
	Floor                    string                     `json:"floor,omitempty"`
	GeoCoordinate            CoordinateType_v02         `json:"geoCoordinate,omitempty"`
	LineFive                 string                     `json:"lineFive,omitempty"`
	LineFour                 string                     `json:"lineFour,omitempty"`
	LineOne                  string                     `json:"lineOne,omitempty"`
	LineThree                string                     `json:"lineThree,omitempty"`
	LineTwo                  string                     `json:"lineTwo,omitempty"`
	NameCode                 CodeType_v02               `json:"nameCode,omitempty"`
	PlotID                   string                     `json:"plotID,omitempty"`
	PostOfficeBox            string                     `json:"postOfficeBox,omitempty"`
	PostalCode               string                     `json:"postalCode,omitempty"`
	ScriptCode               CodeType_v02               `json:"scriptCode,omitempty"`
	StairCase                string                     `json:"stairCase,omitempty"`
	StreetName               string                     `json:"streetName,omitempty"`
	StreetTypeCode           CodeType_v02               `json:"streetTypeCode,omitempty"`
	Unit                     string                     `json:"unit,omitempty"`
}

type GenericClassification struct {
	ClassificationCode CodeType_v02   `json:"classificationCode,omitempty"`
	ItemID             ItemIDType_v01 `json:"itemID,omitempty"`
	NameCode           CodeType_v02   `json:"nameCode,omitempty"`
}

// codeValue is an item from an enumerated list in the json schema that will be different for all of the instances of this generic struct
type GenericCode struct {
	CodeValue string `json:"codeValue,omitempty"`
	LongName  string `json:"longName,omitempty"`
	ShortName string `json:"shortName,omitempty"`
}

type GenericDocument struct {
	CountryCode    SimpleCodeType_v02 `json:"countryCode,omitempty"`
	DocumentID     ItemIDType_v01     `json:"documentID,omitempty"`
	ExpirationDate DateType_v01       `json:"expirationDate,omitempty"`
	IssueDate      DateType_v01       `json:"issueDate,omitempty"`
	IssuingParty   PartyBaseType_v02  `json:"issuingParty,omitempty"`
	ItemID         ItemIDType_v01     `json:"itemID,omitempty"`
	TypeCode       CodeType_v02       `json:"typeCode,omitempty"`
}

type GenericField struct {
	CategoryCode CodeType_v02   `json:"categoryCode,omitempty"`
	ItemID       ItemIDType_v01 `json:"itemID,omitempty"`
	NameCode     CodeType_v02   `json:"nameCode,omitempty"`
}

type GenericPersonName struct {
	FamilyName1            string               `json:"familyName1,omitempty"`
	FamilyName1Prefix      string               `json:"familyName1Prefix,omitempty"`
	FamilyName2            string               `json:"familyName2,omitempty"`
	FamilyName2Prefix      string               `json:"familyName2Prefix,omitempty"`
	FormattedName          string               `json:"formattedName,omitempty"`
	GenerationAffixCode    CodeType_v02         `json:"generationAffixCode,omitempty"`
	GivenName              string               `json:"givenName,omitempty"`
	Initials               string               `json:"initials,omitempty"`
	MiddleName             string               `json:"middleName,omitempty"`
	NameCode               CodeType_v02         `json:"nameCode,omitempty"`
	NickName               string               `json:"nickName,omitempty"`
	PreferredSalutations   []SalutationType_v02 `json:"preferredSalutations,omitempty"`
	QualificationAffixCode CodeType_v02         `json:"qualificationAffixCode,omitempty"`
	ScriptCode             CodeType_v02         `json:"scriptCode,omitempty"`
	TitleAffixCodes        []TitleAffixType_v02 `json:"titleAffixCodes,omitempty"`
	TitlePrefixCodes       []TitleAffixType_v02 `json:"titlePrefixCodes,omitempty"`
}

type GenericTelephoneNumber struct {
	Access                string         `json:"access,omitempty"`
	AreaDialing           string         `json:"areaDialing,omitempty"`
	CountryDialing        string         `json:"countryDialing,omitempty"`
	DialNumber            string         `json:"dialNumber,omitempty"`
	Extension             string         `json:"extension,omitempty"`
	FormattedNumber       string         `json:"formattedNumber,omitempty"`
	ItemID                ItemIDType_v01 `json:"itemID,omitempty"`
	NameCode              CodeType_v02   `json:"nameCode,omitempty"`
	NotificationIndicator bool           `json:"notificationIndicator,omitempty"`
}

type GenericURIAddress struct {
	ItemID                ItemIDType_v01 `json:"itemID,omitempty"`
	NameCode              CodeType_v02   `json:"nameCode,omitempty"`
	NotificationIndicator bool           `json:"notificationIndicator,omitempty"`
	Uri                   URIType_v01    `json:"uri,omitempty"`
}
