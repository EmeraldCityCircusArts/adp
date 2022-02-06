// Package adp implements a simple Golang library for working with the adp.com API.
// Based on ADP API documentation available at
// https://developers.adp.com/articles/api/all/apiexplorer
package adp

import "time"

// "description": "Actions which the user is allowed to initiate against the related entity",
type ActionType_v02 struct {
	ActionTypeCode                string `json:"actionTypeCode,omitempty"`
	Attestation                   `json:"attestation,omitempty"`
	CanonicalUri                  URIType_v01              `json:"canonicalUri,omitempty"`
	CommentAllowedIndicator       IndicatorDefaultType_v01 `json:"commentAllowedIndicator,omitempty"`
	ConfirmationRequiredIndicator IndicatorDefaultType_v01 `json:"confirmationRequiredIndicator,omitempty"`
	DefaultIndicator              IndicatorDefaultType_v01 `json:"defaultIndicator,omitempty"`
	Links                         []LinkType_v02           `json:"links,omitempty"`
	OperationID                   string                   `json:"operationID,omitempty"`
	Sequence                      SequenceType_v01         `json:"sequence,omitempty"`
}

// description": "A unique identifier of a worker the schedule is related to",
type AssociateOIDType_v02 SimpleIDType_v02

// String returns the AssociateOIDType_v02 formatted using the format string
func (a AssociateOIDType_v02) String() string {
	return string(a)
}

type CodeListItemType_v02 struct {
	CodeValue         string `json:"codeValue,omitempty"`
	DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
	ForeignKey        string `json:"foreignKey,omitempty"`
	InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
	ItemID            string `json:"itemID,omitempty"`
	LongName          string `json:"longName,omitempty"`
	ShortName         string `json:"shortName,omitempty"`
	ValueDescription  string `json:"valueDescription,omitempty"`
}

type CodeListType_v03 struct {
	CodeListTitle      string                 `json:"codeListTitle,omitempty"`
	ExclusiveIndicator bool                   `json:"exclusiveIndicator,omitempty"`
	Links              []LinkType_v02         `json:"links,omitempty"`
	ListItems          []CodeListItemType_v02 `json:"listItems,omitempty"`
}

// "description": "The code for the related entity",
type CodeType_v01 struct {
	GenericCode
}

// "description": "The code for the related entity.  If this is a coded value, codeValue and shortName should be used. If this is just a string value, only shortName is necessary",
type CodeType_v02 struct {
	GenericCode
}

// "description": "Communication objects to include various communication mechanisms, e,g, phone, email, etc.",
type CommunicationType_v02 struct {
	Emails            []Email           `json:"emails,omitempty"`
	Faxes             []Fax             `json:"faxes,omitempty"`
	InstantMessages   []InstantMessage  `json:"instantMessages,omitempty"`
	InternetAddresses []InternetAddress `json:"internetAddresses,omitempty"`
	Landlines         []Landline        `json:"landlines,omitempty"`
	Mobiles           []Mobile          `json:"mobiles,omitempty"`
	Pagers            []Pager           `json:"pagers,omitempty"`
	SocialNetworks    []SocialNetwork   `json:"socialNetworks,omitempty"`
}

// "description": "The ConfirmMessage contains the processing results for the corresponding request.  A request may have its processing reported as: succeeded, partially failed, or failed.",
type ConfirmMessageType_v01 struct {
	ConfirmMessageID       IDType_v01        `json:"confirmMessageID,omitempty"`
	CreateDateTime         DateTimeType_v01  `json:"createDateTime,omitempty"`
	ProcessMessages        []ProcessMessage  `json:"processMessages,omitempty"`
	ProcessingStatusCode   GenericCode       `json:"processingStatusCode,omitempty"`
	ProtocolCode           CodeType_v01      `json:"protocolCode,omitempty"`
	ProtocolStatusCode     CodeType_v01      `json:"protocolStatusCode,omitempty"`
	RequestETag            string            `json:"requestETag,omitempty"`
	RequestID              IDType_v01        `json:"requestID,omitempty"`
	RequestLink            LinkType_v02      `json:"requestLink,omitempty"`
	RequestMethodCode      GenericCode       `json:"requestMethodCode,omitempty"`
	RequestReceiptDateTime DateTimeType_v01  `json:"requestReceiptDateTime,omitempty"`
	RequestStatusCode      GenericCode       `json:"requestStatusCode,omitempty"`
	ResourceMessages       []ResourceMessage `json:"resourceMessages,omitempty"`
	SessionID              IDType_v01        `json:"sessionID,omitempty"`
}

// "description": "The name of the employee the schedule request is made for",
type ContactNameType_v02 struct {
	FamilyName1   string `json:"familyName1,omitempty"`
	FamilyName2   string `json:"familyName2,omitempty"`
	FormattedName string `json:"formattedName,omitempty"`
	GivenName     string `json:"givenName,omitempty"`
	MiddleName    string `json:"middleName,omitempty"`
}

// "description": "The coordinates of the location",
type CoordinateType_v02 struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

// "description": "A container of custom fields. The category within each custom object allows for the categorizing or grouping of the associated custom field for the purpose of data mining. For example, the typeCode on a custom string object might be valued as SIZE and the correlating name codes might be t-shirt, pants, jacket, etc."
type CustomFieldContainerType_v02 struct {
	AmountFields    []AmountField    `json:"amountFields,omitempty"`
	CodeFields      []CodeField      `json:"codeFields,omitempty"`
	DateFields      []DateField      `json:"dateFields,omitempty"`
	DateTimeFields  []DateTimeField  `json:"dateTimeFields,omitempty"`
	IndicatorFields []IndicatorField `json:"indicatorFields,omitempty"`
	Links           []LinkType_v02   `json:"links,omitempty"`
	NumberFields    []NumberField    `json:"numberFields,omitempty"`
	PercentFields   []PercentField   `json:"percentFields,omitempty"`
	StringFields    []StringField    `json:"stringFields,omitempty"`
	TelephoneFields []TelephoneField `json:"telephoneFields,omitempty"`
}

// "description": "A country sub-division correlating to a administrative level one or two levels below country. For example, in the United States, this might be a state or a county",
type CountrySubdivisionType_v02 struct {
	GenericCode
	SubdivisionType string `json:"subdivisionType,omitempty"`
}

// "description": "The string reprersentation of the date range. Used to specify effective period.  Follows the ISO-8601:2000 format",
type DatePeriodType_v01 struct {
	EndDate   DateType_v01 `json:"endDate,omitempty"`
	StartDate DateType_v01 `json:"startDate,omitempty"`
}

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

// "description": "Represents the duration of time as represented by ISO 8601. Where the value space is a six dimensional space where the coordinates designate the Gregorian year, month, day, hour, minute, and second. The number of seconds can include decimal digits to arbitary pecision. PnYn MnDTnH nMnS, where: [n] is replaced by the value for each of the date and time elements that follow the [n]. Leading zeros are not required, but the maximum number of digits for each element should be agreed to by the communicating parties. The capital letters 'P', 'Y', 'M', 'W', 'D', 'T', 'H', 'M', and 'S' are designators for each of the date and time elements and are not replaced. P is the duration designator (historically called 'period') placed at the start of the duration representation. Y is the year designator that follows the value for the number of years. M is the month designator that follows the value for the number of months. D is the day designator that follows the value for the number of days. T is the time designator that precedes the time components of the representation. H is the hour designator that follows the value for the number of hours. M is the minute designator that follows the value for the number of minutes. S is the second designator that follows the value for the number of seconds. The number of seconds can include decimal digits to arbitrary precision",
type DurationType_v01 string

// "description": "The unique identifier of the related entity ",
type IDType_v01 struct {
	IdValue          string `json:"idValue,omitempty"`
	SchemeAgencyName string `json:"schemeAgencyName,omitempty"`
	SchemeName       string `json:"schemeName,omitempty"`
}

//  "description": "The unique identifier of the related entity ",
type IDType_v02 struct {
	IdValue    string       `json:"idValue,omitempty"`
	SchemeCode CodeType_v02 `json:"schemeCode,omitempty"`
}

type IndicatorDefaultType_v01 bool

// "description": "Boolean expression",
type IndicatorType_v01 bool

// "description": "The unique identifier of the object in a collection",
type ItemIDType_v01 SimpleIDType_v02

// "description": "A link description object is used to describe link relations.  In the context of a schema, it defines the link relations of the instances of the schema, and can be parameterized by the instance values.  The link description format can be used on its own in regular (non-schema documents), and use of this format can be declared by referencing the normative link description schema as the the schema for the data structure that uses the links",
type LinkType_v02 struct {
	CanonicalUri     URIType_v01 `json:"canonicalUri,omitempty"`
	EncType          `json:"encType,omitempty"`
	Href             URIType_v01 `json:"href,omitempty"`
	MediaType        `json:"mediaType,omitempty"`
	Method           `json:"method,omitempty"`
	PayLoadArguments []PayLoadArgument `json:"payLoadArguments,omitempty"`
	Rel              `json:"rel,omitempty"`
	Schema           string `json:"schema,omitempty"`
	TargetSchema     string `json:"targetSchema,omitempty"`
	Title            string `json:"title,omitempty"`
}

// "description": "Representation of a message."
type MessageType_v01 struct {
	CodeValue  string         `json:"codeValue,omitempty"`
	Links      []LinkType_v02 `json:"links,omitempty"`
	MessageTxt string         `json:"messageTxt,omitempty"`
	Title      string         `json:"title,omitempty"`
}

// "description": "Meta contains the metadata associated with a response.  A GET response may include metadata to support pagination",
type MetaType_v02 struct {
	CompleteIndicator IndicatorType_v01 `json:"completeIndicator,omitempty"`
	Links             []LinkType_v02    `json:"links,omitempty"`
	ResourceSetID     SimpleIDType_v02  `json:"resourceSetID,omitempty"`
	StartSequence     SequenceType_v01  `json:"startSequence,omitempty"`
	TotalNumber       float64           `json:"totalNumber,omitempty"`
}

// "description": "organization address objects / components",
type OrganizationAddressType_v02 GenericAddress

// "description": "Schema describing the details of a party - an agency, bureau, or other organization issuing the associated document; the body or persons exercising power or command. An issuing party is often is a government agency or official authority, but sometimes might be a private organization",
type PartyBaseType_v02 struct {
	Address  OrganizationAddressType_v02 `json:"address,omitempty"`
	NameCode CodeType_v02                `json:"nameCode,omitempty"`
}

// "description": "A representation of a percentage value",
type PercentageType_v02 float64

// "description": "Object describing the details of a person name",
// Exported and altered all consumers. Original: type personNameType_v02 GenericPersonName
type PersonNameType_v02 GenericPersonName

// "description": "A representation of a rate amount value",
type RateAmountType_v02 struct {
	AmountValue  float64            `json:"amountValue,omitempty"`
	CurrencyCode SimpleCodeType_v02 `json:"currencyCode,omitempty"`
	NameCode     CodeType_v02       `json:"nameCode,omitempty"`
}

// "description": "A representation of a rate value",
type RateType_v02 struct {
	AmountValue         float64            `json:"amountValue,omitempty"`
	BaseMultiplierValue float64            `json:"baseMultiplierValue,omitempty"`
	BaseUnitCode        CodeType_v02       `json:"baseUnitCode,omitempty"`
	CurrencyCode        SimpleCodeType_v02 `json:"currencyCode,omitempty"`
	UnitCode            CodeType_v02       `json:"unitCode,omitempty"`
}

// "description": "The ratio",
type RatioType_v02 float64

// "description": "A preferred salutation, e.g. Mr., Dr., etc.",
type SalutationType_v02 struct {
	SalutationCode CodeType_v02     `json:"salutationCode,omitempty"`
	SequenceNumber SequenceType_v01 `json:"sequenceNumber,omitempty"`
	TypeCode       CodeType_v02     `json:"typeCode,omitempty"`
}

// "description": "Sequence of a related entity when included in the collection or a group",
type SequenceType_v01 float64

// "description": "A simple (string) code.  Can have a code list reference",
type SimpleCodeType_v02 string

// "description": "The simple (string) identifier of an object",
type SimpleIDType_v02 string

// "description": "The standard number of hours of work associated to a position, typically used to drive the definition of a full time assignment, e.g. 40 in the US, 37.5 in the UK, 35 in FR. This number is usually expressed based on a week. The unitCode is used to convey that measurement.",
type StandardHoursType_v02 struct {
	HoursQuantity float64      `json:"hoursQuantity,omitempty"`
	UnitCode      CodeType_v02 `json:"unitCode,omitempty"`
}

// "description": "Status to include the status code, effective date (ISO-8601:2000 format) and the reason for change",
type StatusReasonType_v02 struct {
	EffectiveDate DateType_v01 `json:"effectiveDate,omitempty"`
	ReasonCode    CodeType_v02 `json:"reasonCode,omitempty"`
	StatusCode    CodeType_v02 `json:"statusCode,omitempty"`
}

// "description": "Status to include the status code and effective date (ISO-8601:2000 format)",
type StatusType_v02 struct {
	EffectiveDate DateType_v01 `json:"effectiveDate,omitempty"`
	GenericCode
}

// "description": "The string reprersentation of the time value usually expressed as hh:mm",
type TimeType_v01 string

// "description": "Person titles, e.g. social, professional, generational, etc.",
type TitleAffixType_v02 struct {
	AffixCode      CodeType_v02     `json:"affixCode,omitempty"`
	SequenceNumber SequenceType_v01 `json:"sequenceNumber,omitempty"`
}

// "description": "The URI of the related entity",
type URIType_v01 string
