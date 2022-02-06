package adp

// "description": "A list of mobile telephone numbers",
type Mobile GenericTelephoneNumber

// "$ref": "#/definitions/emailUriType_v02"
// "description": "The URI of the associated email address",
type EmailUri string

// String returns the EmailUri formatted using the format string
func (e EmailUri) String() string {
	return string(e)
}

// "description": "A list of email addresses",
type Email struct {
	EmailUri              `json:"emailUri,omitempty"`
	ItemID                ItemIDType_v01 `json:"itemID,omitempty"`
	NameCode              CodeType_v02   `json:"nameCode,omitempty"`
	NotificationIndicator bool           `json:"notificationIndicator,omitempty"`
}

// "description": "A list of fax telephone numbers",
type Fax GenericTelephoneNumber

// "description": "A list of pager telephone numbers",
type Pager GenericTelephoneNumber

// "description": "A list of instant messages URIs",
type InstantMessage GenericURIAddress

// "description": "A list of internet URIs",
type InternetAddress struct {
	ItemID   ItemIDType_v01 `json:"itemID,omitempty"`
	NameCode CodeType_v02   `json:"nameCode,omitempty"`
	Uri      URIType_v01    `json:"uri,omitempty"`
}

// "description": "A list of social networks addresses",
type SocialNetwork GenericURIAddress

// "description": "A list of land-line telephone numbers",
type Landline GenericTelephoneNumber

// "description": "A list of custom amounts",
type AmountField struct {
	GenericField
	AmountValue  float64            `json:"amountValue,omitempty"`
	CurrencyCode SimpleCodeType_v02 `json:"currencyCode,omitempty"`
}

// "description": "A list of custom codes",
type CodeField struct {
	GenericField
	CodeValue string `json:"codeValue,omitempty"`
	LongName  string `json:"longName,omitempty"`
	ShortName string `json:"shortName,omitempty"`
}

// "description": "A list of custom dates",
type DateField struct {
	GenericField
	DateValue DateType_v01 `json:"dateValue,omitempty"`
}

// "description": "A list of custom datetimes",
type DateTimeField struct {
	GenericField
	DateTimeValue DateTimeType_v01 `json:"dateTimeValue,omitempty"`
}

// "description": "A list of custom indicators",
type IndicatorField struct {
	GenericField
	IndicatorValue IndicatorType_v01 `json:"indicatorValue,omitempty"`
}

// "description": "A list of custom numbers",
type NumberField struct {
	GenericField
	NumberValue float64 `json:"numberValue,omitempty"`
}

// "description": "A list of custom percentages",
type PercentField struct {
	GenericField
	PercentValue float64 `json:"percentValue,omitempty"`
}

// "description": "A list of custom strings",
type StringField struct {
	GenericField
	StringValue string `json:"stringValue,omitempty"`
}

// "description": "A list of custom telephone numbers",
type TelephoneField struct {
	GenericField
	Access          string `json:"access,omitempty"`
	AreaDialing     string `json:"areaDialing,omitempty"`
	CountryDialing  string `json:"countryDialing,omitempty"`
	DialNumber      string `json:"dialNumber,omitempty"`
	Extension       string `json:"extension,omitempty"`
	FormattedNumber string `json:"formattedNumber,omitempty"`
}

// "$ref": "#/definitions/governmentIDItemType_v02"
// "description": "A list of government identifiers",
type GovernmentID struct {
	CountryCode    SimpleCodeType_v02 `json:"countryCode,omitempty"`
	ExpirationDate DateType_v01       `json:"expirationDate,omitempty"`
	IdValue        string             `json:"idValue,omitempty"`
	ItemID         ItemIDType_v01     `json:"itemID,omitempty"`
	NameCode       CodeType_v02       `json:"nameCode,omitempty"`
	StatusCode     StatusType_v02     `json:"statusCode,omitempty"`
}

// "description": "A list of addresses related to the associated entity",
type OtherPersonalAddress struct {
	GenericAddress
	ItemID                 ItemIDType_v01    `json:"itemID,omitempty"`
	SameAsAddressIndicator IndicatorType_v01 `json:"sameAsAddressIndicator,omitempty"`
	SameAsAddressLink      LinkType_v02      `json:"sameAsAddressLink,omitempty"`
	TypeCode               CodeType_v02      `json:"typeCode,omitempty"`
}

// "$ref": "#/definitions/personalAddressType_v02"
// "description": "This is the primary residential address that  is used by an organization for payroll, taxation, and benefit program eligibility",
type LegalAddress struct {
	GenericAddress
	SameAsAddressIndicator IndicatorDefaultType_v01 `json:"sameAsAddressIndicator,omitempty"`
	SameAsAddressLink      LinkType_v02             `json:"sameAsAddressLink,omitempty"`
}

// "description": "The date range of the coverage for the related entity",
type CoveragePeriod struct {
	EndDate   DateType_v01 `json:"endDate,omitempty"`
	StartDate DateType_v01 `json:"startDate,omitempty"`
}

// "$ref": "#/definitions/socialInsuranceProgramItemType_v02"
// "description": "Social insurance is a government operated social security scheme that provides income loss, maternity, sickness benefits, and retirement pensions to the working population. It is funded by regular contributions based on wage or salary by both employees and employers, and by government through taxation. Examples in the USA are Medicare and Medicaid",
type SocialInsuranceProgram struct {
	CountryCode      SimpleCodeType_v02 `json:"countryCode,omitempty"`
	CoveragePeriod   `json:"coveragePeriod,omitempty"`
	CoveredIndicator IndicatorType_v01 `json:"coveredIndicator,omitempty"`
	ItemID           ItemIDType_v01    `json:"itemID,omitempty"`
	NameCode         CodeType_v02      `json:"nameCode,omitempty"`
}

// "description": "A list of identity documents, that serve as to confirm the identification of the holder , e.g. a birth certificate, an ID Card",
type IdentityDocument GenericDocument

//  "description": "A list of immigration documents for a person who is not a citizen of the country in which he or see resides , e.g. a passport, visa, alien registration card, aka, green card, residence permit, etc.",
type ImmigrationDocument struct {
	GenericDocument
	AuthorizedStayDuration     DurationType_v01 `json:"authorizedStayDuration,omitempty"`
	ReentryRequirementDuration DurationType_v01 `json:"reentryRequirementDuration,omitempty"`
}

//  "description": "A list of passport documents, issued by a national govenment which certifies the identiy and nationality of its holder",
type Passport GenericDocument

// "description": "A list of documents that authorize work, e.g. visas, a work permit, a labor card, etc.",
type WorkAuthorizationDocument GenericDocument

// "description": "Former names of the person in the associated context",
type FormerName struct {
	GenericPersonName
	ItemID   ItemIDType_v01 `json:"itemID,omitempty"`
	TypeCode CodeType_v02   `json:"typeCode,omitempty"`
}

// "description": "Other preferred names of the person expressed in a different language or script",
type AlternatePreferredName struct {
	GenericPersonName
	ItemID       ItemIDType_v01 `json:"itemID,omitempty"`
	LanguageCode CodeType_v02   `json:"languageCode,omitempty"`
}

// "$ref": "#/definitions/birthPlaceType_v02"
// "description": "The place a person was born. This can have taxation impacts",
type BirthPlace struct {
	CityName                 string                     `json:"cityName,omitempty"`
	CountryCode              SimpleCodeType_v02         `json:"countryCode,omitempty"`
	CountrySubdivisionLevel1 CountrySubdivisionType_v02 `json:"countrySubdivisionLevel1,omitempty"`
	CountrySubdivisionLevel2 CountrySubdivisionType_v02 `json:"countrySubdivisionLevel2,omitempty"`
	FormattedBirthPlace      string                     `json:"formattedBirthPlace,omitempty"`
	PostalCode               SimpleCodeType_v02         `json:"postalCode,omitempty"`
}

// "description": "Race is a classification system used to categorize humans into large and distinct populations or groups by heritable phenotypic characteristics, geographic ancestry, culture, history, language, physical appearance, ethnicity, and social status",
type RaceCode struct {
	GenericCode
	IdentificationMethodCode CodeType_v02 `json:"identificationMethodCode,omitempty"`
}

// "$ref": "#/definitions/personType_v02"
// "description": "An associate is a person with a relationship to an organization",
type Person struct {
	AlternatePreferredNames                   []AlternatePreferredName `json:"alternatePreferredNames,omitempty"`
	BirthDate                                 DateType_v01             `json:"birthDate,omitempty"`
	BirthName                                 PersonNameType_v02       `json:"birthName,omitempty"`
	BirthPlace                                `json:"birthPlace,omitempty"`
	CitizenshipCountryCodes                   []CodeType_v02               `json:"citizenshipCountryCodes,omitempty"`
	Communication                             CommunicationType_v02        `json:"communication,omitempty"`
	CustomFieldGroup                          CustomFieldContainerType_v02 `json:"customFieldGroup,omitempty"`
	DeathDate                                 DateType_v01                 `json:"deathDate,omitempty"`
	DeceasedIndicator                         IndicatorType_v01            `json:"deceasedIndicator,omitempty"`
	DisabilityIdentificationDeclinedIndicator IndicatorType_v01            `json:"disabilityIdentificationDeclinedIndicator,omitempty"`
	DisabilityPercentage                      PercentageType_v02           `json:"disabilityPercentage,omitempty"`
	DisabilityTypeCodes                       []CodeType_v02               `json:"disabilityTypeCodes,omitempty"`
	DisabledIndicator                         IndicatorType_v01            `json:"disabledIndicator,omitempty"`
	EthnicityCode                             CodeType_v02                 `json:"ethnicityCode,omitempty"`
	EthnicityGroupCode                        CodeType_v02                 `json:"ethnicityGroupCode,omitempty"`
	FormerNames                               []FormerName                 `json:"formerNames,omitempty"`
	GenderCode                                CodeType_v02                 `json:"genderCode,omitempty"`
	GenderIdentityCode                        CodeType_v02                 `json:"genderIdentityCode,omitempty"`
	GovernmentIDs                             []GovernmentID               `json:"governmentIDs,omitempty"`
	HighestEducationLevelCode                 CodeType_v02                 `json:"highestEducationLevelCode,omitempty"`
	IdentityDocuments                         []IdentityDocument           `json:"identityDocuments,omitempty"`
	ImmigrationDocuments                      []ImmigrationDocument        `json:"immigrationDocuments,omitempty"`
	LegalAddress                              `json:"legalAddress,omitempty"`
	LegalName                                 PersonNameType_v02     `json:"legalName,omitempty"`
	Links                                     []LinkType_v02         `json:"links,omitempty"`
	MaritalStatusCode                         StatusType_v02         `json:"maritalStatusCode,omitempty"`
	MilitaryClassificationCodes               []CodeType_v02         `json:"militaryClassificationCodes,omitempty"`
	MilitaryDischargeDate                     DateType_v01           `json:"militaryDischargeDate,omitempty"`
	MilitaryStatusCode                        StatusType_v02         `json:"militaryStatusCode,omitempty"`
	OtherPersonalAddresses                    []OtherPersonalAddress `json:"otherPersonalAddresses,omitempty"`
	Passports                                 []Passport             `json:"passports,omitempty"`
	PreferredName                             PersonNameType_v02     `json:"preferredName,omitempty"`
	RaceCode                                  `json:"raceCode,omitempty"`
	ReligionCode                              CodeType_v02                `json:"religionCode,omitempty"`
	ResidencyCountryCodes                     []CodeType_v02              `json:"residencyCountryCodes,omitempty"`
	SexualOrientationCode                     CodeType_v02                `json:"sexualOrientationCode,omitempty"`
	SocialInsurancePrograms                   []SocialInsuranceProgram    `json:"socialInsurancePrograms,omitempty"`
	StudentIndicator                          IndicatorType_v01           `json:"studentIndicator,omitempty"`
	StudentStatusCode                         StatusType_v02              `json:"studentStatusCode,omitempty"`
	TobaccoUserIndicator                      IndicatorType_v01           `json:"tobaccoUserIndicator,omitempty"`
	WorkAuthorizationDocuments                []WorkAuthorizationDocument `json:"workAuthorizationDocuments,omitempty"`
}

// "description": "A list of photos",
type Photo struct {
	ItemID   ItemIDType_v01 `json:"itemID,omitempty"`
	Links    []LinkType_v02 `json:"links,omitempty"`
	NameCode CodeType_v02   `json:"nameCode,omitempty"`
}

// "$ref": "#/definitions/costCenterType_v02"
// description": "The cost Center is used for taxing the costs of the related entity",
type AssignmentCostCenter struct {
	CostCenterID         SimpleIDType_v02   `json:"costCenterID,omitempty"`
	CostCenterName       string             `json:"costCenterName,omitempty"`
	CostCenterPercentage PercentageType_v02 `json:"costCenterPercentage,omitempty"`
}

// "description": "A list of worker groups",
type WorkerGroup struct {
	GroupCode CodeType_v02   `json:"groupCode,omitempty"`
	ItemID    ItemIDType_v01 `json:"itemID,omitempty"`
	NameCode  CodeType_v02   `json:"nameCode,omitempty"`
}

// "description": "A reference to a system used to classify workers into occupational categories for the purpose of collecting, calculating, or disseminating data. In the US, an example of this is the EEOC system; in the UK, an example of this is the SOC system. This is the at the assignment level in France",
type OccupationalClassification GenericClassification

// "description": "The occupational classification reflects the type of job or work that the person does, while the industry classification reflects the business activity of their employer or company. The occupational and industry classifications are based on a persons sole or primary job, unless otherwise specified. For the unemployed, the occupation and industry are based on the last job held. An example of an industry classification code in the US is the NAICS code.",
type IndustryClassification GenericClassification

// "$ref": "#/definitions/wageLawCoverageType_v02"
// "description": "A code indicating whether the work assignment is covered under applicable wage and labor laws relating to working hours, overtime, and similar protections, e.g. FLSA in the US",
type WageLawCoverage struct {
	CoverageCode    CodeType_v02 `json:"coverageCode,omitempty"`
	WageLawNameCode CodeType_v02 `json:"wageLawNameCode,omitempty"`
}

// "$ref": "#/definitions/positionIDType_v02"
// "description": "The position identifier",
type PositionID SimpleIDType_v02

// "$ref": "#/definitions/laborUnionType_v02"
// "description": "The associated labor union",
type LaborUnion struct {
	LaborUnionCode CodeType_v02 `json:"laborUnionCode,omitempty"`
	SeniorityDate  DateType_v01 `json:"seniorityDate,omitempty"`
}

// "$ref": "#/definitions/bargainingUnitType_v02"
// "description": "The associated bargaining unit (local labor union)",
type BargainingUnit struct {
	BargainingUnitCode CodeType_v02 `json:"bargainingUnitCode,omitempty"`
	SeniorityDate      DateType_v01 `json:"seniorityDate,omitempty"`
}

// "description": "The HOME organizational unit defines the work assignment and perhaps funds it. The should only be one home organizational unit for a given type, e.g. only one home department.",
type HomeOrganizationalUnit struct {
	ItemID   ItemIDType_v01 `json:"itemID,omitempty"`
	NameCode CodeType_v02   `json:"nameCode,omitempty"`
	TypeCode CodeType_v02   `json:"typeCode,omitempty"`
}

// "description": "The ASSIGNED organizational unit is one for which the work assignment is doing work",
type AssignedOrganizationalUnit struct {
	ItemID   ItemIDType_v01 `json:"itemID,omitempty"`
	NameCode CodeType_v02   `json:"nameCode,omitempty"`
	TypeCode CodeType_v02   `json:"typeCode,omitempty"`
}

// "$ref": "#/definitions/locationType_v02"
// "description": "The home work location is the location out of which the work assignment is based. It does not necessarily mean that the worker is physically located in that location. This data becomes relevant for affirmative action reasons. In the US for example, if a worker works from home or a location where there is not an EEO-1 form/location, then the worker is accounted for based on the location of the reports to manager or work unit. Additionally, other things like market salaries, taxes, etc. would be aligned to home office location",
type HomeWorkLocation struct {
	Address       OrganizationAddressType_v02 `json:"address,omitempty"`
	Communication CommunicationType_v02       `json:"communication,omitempty"`
	NameCode      CodeType_v02                `json:"nameCode,omitempty"`
}

// "description": "The assigned work location is the location (or locations) where the worker is physically located.",
type AssignedWorkLocation struct {
	Address       OrganizationAddressType_v02 `json:"address,omitempty"`
	Communication CommunicationType_v02       `json:"communication,omitempty"`
	ItemID        ItemIDType_v01              `json:"itemID,omitempty"`
	NameCode      CodeType_v02                `json:"nameCode,omitempty"`
}

// "$ref": "#/definitions/ratePercentageType_v02"
// "description": "A percentage rate",
type CommissionRatePercentage struct {
	BaseUnitCode    CodeType_v02 `json:"baseUnitCode,omitempty"`
	NameCode        CodeType_v02 `json:"nameCode,omitempty"`
	PercentageValue float64      `json:"percentageValue,omitempty"`
}

//  "description": "Describes the conditions under which the associated rate can be expected",
type AssociatedRateQualifier struct {
	ItemID                  ItemIDType_v01 `json:"itemID,omitempty"`
	QualifierObjectCode     CodeType_v02   `json:"qualifierObjectCode,omitempty"`
	QualifierObjectID       ItemIDType_v01 `json:"qualifierObjectID,omitempty"`
	QualifierObjectTypeCode CodeType_v02   `json:"qualifierObjectTypeCode,omitempty"`
}

// "description": "The base or primary pay (i.e. salary, wage) associated to a work assignment. The structure allows for the explicit communication of pre-defined equivalents for the base remuneration",
type BaseRemuneration struct {
	AnnualRateAmount         RateAmountType_v02        `json:"annualRateAmount,omitempty"`
	AssociatedRateQualifiers []AssociatedRateQualifier `json:"associatedRateQualifiers,omitempty"`
	BiweeklyRateAmount       RateAmountType_v02        `json:"biweeklyRateAmount,omitempty"`
	CommissionRatePercentage `json:"commissionRatePercentage,omitempty"`
	DailyRateAmount          RateAmountType_v02 `json:"dailyRateAmount,omitempty"`
	EffectiveDate            DateType_v01       `json:"effectiveDate,omitempty"`
	HourlyRateAmount         RateAmountType_v02 `json:"hourlyRateAmount,omitempty"`
	MonthlyRateAmount        RateAmountType_v02 `json:"monthlyRateAmount,omitempty"`
	PayPeriodRateAmount      RateAmountType_v02 `json:"payPeriodRateAmount,omitempty"`
	RecordingBasisCode       CodeType_v02       `json:"recordingBasisCode,omitempty"`
	SemiMonthlyRateAmount    RateAmountType_v02 `json:"semiMonthlyRateAmount,omitempty"`
	WeeklyRateAmount         RateAmountType_v02 `json:"weeklyRateAmount,omitempty"`
}

// "description": "A list of pay compensations associated to a given work assignment.  Additional remuneration includes all other forms of compensation beyond the base pay associated to a given work assignment.  The typeCode is used to further refine the type of compensation, e.g. bonus, commission, etc.",
type AdditionalRemuneration struct {
	AssociatedRateQualifiers []AssociatedRateQualifier `json:"associatedRateQualifiers,omitempty"`
	EffectiveDate            DateType_v01              `json:"effectiveDate,omitempty"`
	IntervalCode             CodeType_v02              `json:"intervalCode,omitempty"`
	ItemID                   ItemIDType_v01            `json:"itemID,omitempty"`
	NameCode                 CodeType_v02              `json:"nameCode,omitempty"`
	Rate                     RateType_v02              `json:"rate,omitempty"`
	TypeCode                 CodeType_v02              `json:"typeCode,omitempty"`
}

//   "description": "The work assignment to whom a given worker reports",
type ReportTo struct {
	AssociateOID              AssociateOIDType_v02 `json:"associateOID,omitempty"`
	ItemID                    ItemIDType_v01       `json:"itemID,omitempty"`
	PositionID                SimpleIDType_v02     `json:"positionID,omitempty"`
	PositionTitle             string               `json:"positionTitle,omitempty"`
	ReportsToRelationshipCode CodeType_v02         `json:"reportsToRelationshipCode,omitempty"`
	ReportsToWorkerName       ContactNameType_v02  `json:"reportsToWorkerName,omitempty"`
	WorkerID                  IDType_v02           `json:"workerID,omitempty"`
}

// "$ref": "#/definitions/rangeType_v02"
// "description": "The pay range of the pay grade",
type PayGradePayRange struct {
	MaximumRate RateType_v02 `json:"maximumRate,omitempty"`
	MedianRate  RateType_v02 `json:"medianRate,omitempty"`
	MinimumRate RateType_v02 `json:"minimumRate,omitempty"`
}

// "$ref": "#/definitions/workAssignmentItemType_v02"
// "description": "An object describing the details of a work assignment. A work assignment pertains to the duties and responsibilities that should be performed by the worker",
type WorkAssignment struct {
	ActualStartDate                     DateType_v01                 `json:"actualStartDate,omitempty"`
	AdditionalRemunerations             []AdditionalRemuneration     `json:"additionalRemunerations,omitempty"`
	AnnualBenefitBaseRate               RateAmountType_v02           `json:"annualBenefitBaseRate,omitempty"`
	AssignedOrganizationalUnits         []AssignedOrganizationalUnit `json:"assignedOrganizationalUnits,omitempty"`
	AssignedWorkLocations               []AssignedWorkLocation       `json:"assignedWorkLocations,omitempty"`
	AssignmentCostCenters               []AssignmentCostCenter       `json:"assignmentCostCenters,omitempty"`
	AssignmentStatus                    StatusReasonType_v02         `json:"assignmentStatus,omitempty"`
	AssignmentTermCode                  CodeType_v02                 `json:"assignmentTermCode,omitempty"`
	BargainingUnit                      `json:"bargainingUnit,omitempty"`
	BaseRemuneration                    `json:"baseRemuneration,omitempty"`
	CompaRatio                          RatioType_v02                `json:"compaRatio,omitempty"`
	CustomFieldGroup                    CustomFieldContainerType_v02 `json:"customFieldGroup,omitempty"`
	ExecutiveIndicator                  IndicatorType_v01            `json:"executiveIndicator,omitempty"`
	ExecutiveTypeCode                   CodeType_v02                 `json:"executiveTypeCode,omitempty"`
	ExpectedStartDate                   DateType_v01                 `json:"expectedStartDate,omitempty"`
	ExpectedTerminationDate             DateType_v01                 `json:"expectedTerminationDate,omitempty"`
	FullTimeEquivalenceRatio            RatioType_v02                `json:"fullTimeEquivalenceRatio,omitempty"`
	GeographicPayDifferentialCode       CodeType_v02                 `json:"geographicPayDifferentialCode,omitempty"`
	GeographicPayDifferentialPercentage PercentageType_v02           `json:"geographicPayDifferentialPercentage,omitempty"`
	HighlyCompensatedIndicator          IndicatorType_v01            `json:"highlyCompensatedIndicator,omitempty"`
	HighlyCompensatedTypeCode           CodeType_v02                 `json:"highlyCompensatedTypeCode,omitempty"`
	HireDate                            DateType_v01                 `json:"hireDate,omitempty"`
	HomeOrganizationalUnits             []HomeOrganizationalUnit     `json:"homeOrganizationalUnits,omitempty"`
	HomeWorkLocation                    `json:"homeWorkLocation,omitempty"`
	IndustryClassifications             []IndustryClassification `json:"industryClassifications,omitempty"`
	ItemID                              ItemIDType_v01           `json:"itemID,omitempty"`
	JobCode                             StatusType_v02           `json:"jobCode,omitempty"`
	JobTitle                            string                   `json:"jobTitle,omitempty"`
	LaborUnion                          `json:"laborUnion,omitempty"`
	LegalEntityID                       ItemIDType_v01               `json:"legalEntityID,omitempty"`
	Links                               []LinkType_v02               `json:"links,omitempty"`
	ManagementPositionIndicator         IndicatorType_v01            `json:"managementPositionIndicator,omitempty"`
	MinimumPayGradeStepDuration         DurationType_v01             `json:"minimumPayGradeStepDuration,omitempty"`
	NationalityContextCode              CodeType_v02                 `json:"nationalityContextCode,omitempty"`
	NextPayGradeStepDate                DateType_v01                 `json:"nextPayGradeStepDate,omitempty"`
	OccupationalClassifications         []OccupationalClassification `json:"occupationalClassifications,omitempty"`
	OfferAcceptanceDate                 DateType_v01                 `json:"offerAcceptanceDate,omitempty"`
	OfferExtensionDate                  DateType_v01                 `json:"offerExtensionDate,omitempty"`
	OfficerIndicator                    IndicatorType_v01            `json:"officerIndicator,omitempty"`
	OfficerTypeCode                     CodeType_v02                 `json:"officerTypeCode,omitempty"`
	PayCycleCode                        CodeType_v02                 `json:"payCycleCode,omitempty"`
	PayGradeCode                        CodeType_v02                 `json:"payGradeCode,omitempty"`
	PayGradePayRange                    `json:"payGradePayRange,omitempty"`
	PayGradeStepCode                    CodeType_v02       `json:"payGradeStepCode,omitempty"`
	PayGradeStepPayRate                 RateType_v02       `json:"payGradeStepPayRate,omitempty"`
	PayScaleCode                        CodeType_v02       `json:"payScaleCode,omitempty"`
	PayrollFileNumber                   string             `json:"payrollFileNumber,omitempty"`
	PayrollGroupCode                    SimpleCodeType_v02 `json:"payrollGroupCode,omitempty"`
	PayrollProcessingStatusCode         StatusType_v02     `json:"payrollProcessingStatusCode,omitempty"`
	PayrollRegionCode                   SimpleCodeType_v02 `json:"payrollRegionCode,omitempty"`
	PayrollScheduleGroupID              SimpleIDType_v02   `json:"payrollScheduleGroupID,omitempty"`
	PositionID                          `json:"positionID,omitempty"`
	PositionTitle                       string                `json:"positionTitle,omitempty"`
	PrimaryIndicator                    IndicatorType_v01     `json:"primaryIndicator,omitempty"`
	RemunerationBasisCode               CodeType_v02          `json:"remunerationBasisCode,omitempty"`
	ReportsTo                           []ReportTo            `json:"reportsTo,omitempty"`
	SeniorityDate                       DateType_v01          `json:"seniorityDate,omitempty"`
	StandardHours                       StandardHoursType_v02 `json:"standardHours,omitempty"`
	StandardPayPeriodHours              StandardHoursType_v02 `json:"standardPayPeriodHours,omitempty"`
	StockOwnerIndicator                 IndicatorType_v01     `json:"stockOwnerIndicator,omitempty"`
	StockOwnerPercentage                PercentageType_v02    `json:"stockOwnerPercentage,omitempty"`
	TerminationDate                     DateType_v01          `json:"terminationDate,omitempty"`
	VipIndicator                        IndicatorType_v01     `json:"vipIndicator,omitempty"`
	VipTypeCode                         CodeType_v02          `json:"vipTypeCode,omitempty"`
	WageLawCoverage                     `json:"wageLawCoverage,omitempty"`
	WorkArrangementCode                 CodeType_v02       `json:"workArrangementCode,omitempty"`
	WorkLevelCode                       CodeType_v02       `json:"workLevelCode,omitempty"`
	WorkShiftCode                       CodeType_v02       `json:"workShiftCode,omitempty"`
	WorkerGroups                        []WorkerGroup      `json:"workerGroups,omitempty"`
	WorkerProbationIndicator            IndicatorType_v01  `json:"workerProbationIndicator,omitempty"`
	WorkerProbationPeriod               DatePeriodType_v01 `json:"workerProbationPeriod,omitempty"`
	WorkerTypeCode                      CodeType_v02       `json:"workerTypeCode,omitempty"`
}

// "description": "These dates are related to the highest level legal relationship that a person has with the organization. If a given worker has multiple relationships with the organization, these are the dates associated with the primary relationship.",
type WorkerDates struct {
	AcquisitionDate         DateType_v01 `json:"acquisitionDate,omitempty"`
	AdjustedServiceDate     DateType_v01 `json:"adjustedServiceDate,omitempty"`
	ExpectedTerminationDate DateType_v01 `json:"expectedTerminationDate,omitempty"`
	OriginalHireDate        DateType_v01 `json:"originalHireDate,omitempty"`
	RehireDate              DateType_v01 `json:"rehireDate,omitempty"`
	RetirementDate          DateType_v01 `json:"retirementDate,omitempty"`
	TerminationDate         DateType_v01 `json:"terminationDate,omitempty"`
}

// "description": "A worker is someone who performs duties and has responsibilities for an organization",
type Worker struct {
	AssociateOID          AssociateOIDType_v02         `json:"associateOID,omitempty"`
	BusinessCommunication CommunicationType_v02        `json:"businessCommunication,omitempty"`
	CustomFieldGroup      CustomFieldContainerType_v02 `json:"customFieldGroup,omitempty"`
	Links                 []LinkType_v02               `json:"links,omitempty"`
	Person                `json:"person,omitempty"`
	Photos                []Photo          `json:"photos,omitempty"`
	WorkAssignments       []WorkAssignment `json:"workAssignments,omitempty"`
	WorkerDates           `json:"workerDates,omitempty"`
	WorkerID              IDType_v02           `json:"workerID,omitempty"`
	WorkerStatus          StatusReasonType_v02 `json:"workerStatus,omitempty"`
}
