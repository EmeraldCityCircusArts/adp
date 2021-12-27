package adp

// "$ref": "#/definitions/attestationType_v02"
// "description": "Optional text provided to the user when they take an action to affirm to be correct, true, or genuine. Requires confirmation indicator to be set to true.",
type Attestation struct {
	ActionBlockIndicator IndicatorDefaultType_v01 `json:"actionBlockIndicator,omitempty"`
	MessageTxt           string                   `json:"messageTxt,omitempty"`
	MessageUri           URIType_v01              `json:"messageUri,omitempty"`
}

// "$ref": "#/definitions/datePeriodType_v02"
// "description": "The string reprersentation of the date range. Used to specify effective period.  Follows the ISO-8601:2000 format",
type DatePeriod struct {
	EndDate   DateType_v01 `json:"endDate,omitempty"`
	EndTime   string       `json:"endTime,omitempty"`
	StartDate DateType_v01 `json:"startDate,omitempty"`
	StartTime string       `json:"startTime,omitempty"`
}

// "$ref": "#/definitions/dateTimePeriodType_v01"
// "description": "The string reprersentation of the date range. Used to specify effective period.  Follows the ISO-8601:2000 format.",
type DateTimePeriod struct {
	EndDateTime   DateTimeType_v01 `json:"endDateTime,omitempty"`
	StartDateTime DateTimeType_v01 `json:"startDateTime,omitempty"`
}

type EarningAllocation struct {
	AllocationCode     CodeType_v02   `json:"allocationCode,omitempty"`
	AllocationTypeCode CodeType_v02   `json:"allocationTypeCode,omitempty"`
	ItemID             ItemIDType_v01 `json:"itemID,omitempty"`
}

type EntryComment struct {
	AppliesToCode    CodeType_v02      `json:"appliesToCode,omitempty"`
	CommentTypeCode  CodeType_v02      `json:"commentTypeCode,omitempty"`
	ItemID           ItemIDType_v01    `json:"itemID,omitempty"`
	PrivateIndicator IndicatorType_v01 `json:"privateIndicator,omitempty"`
	TextValue        string            `json:"textValue,omitempty"`
}

// "description": "The type of tag",
type TagType struct {
	CurrencyCode SimpleCodeType_v02 `json:"currencyCode,omitempty"`
	DataTypeCode string             `json:"dataTypeCode,omitempty"`
	FormatCode   string             `json:"formatCode,omitempty"`
}

// "description": "This is genric object which can be use forvarious purpose.  e.g employee job in organzational map",
type EntryTags struct {
	TagCode   string `json:"tagCode,omitempty"`
	TagType   `json:"tagType,omitempty"`
	TagValues []string `json:"tagValues,omitempty"`
}

// "$ref": "#/definitions/quantityTimeType_v02"
// "description": "The total quantity of of the time expressed in the units as per unit type",
type TotalQuantity struct {
	QuantityValue float64      `json:"quantityValue,omitempty"`
	UnitTimeCode  CodeType_v02 `json:"unitTimeCode,omitempty"`
}

// "description": "Total time represented as hh:mm",
type TotalTime struct {
	NameCode  CodeType_v02 `json:"nameCode,omitempty"`
	TimeValue TimeType_v01 `json:"timeValue,omitempty"`
}

// "description": "A collection of entries within a single day, e.g. shifts, paycode entries, etc.",
type ScheduleEntry struct {
	Actions                       []ActionType_v02 `json:"actions,omitempty"`
	CategoryTypeCode              CodeType_v02     `json:"categoryTypeCode,omitempty"`
	DatePeriod                    `json:"datePeriod,omitempty"`
	DateTimePeriod                `json:"dateTimePeriod,omitempty"`
	DayPeriodCode                 CodeType_v02        `json:"dayPeriodCode,omitempty"`
	Duration                      DurationType_v01    `json:"duration,omitempty"`
	DurationTypeCode              CodeType_v02        `json:"durationTypeCode,omitempty"`
	EarningAllocations            []EarningAllocation `json:"earningAllocations,omitempty"`
	EntryComments                 []EntryComment      `json:"entryComments,omitempty"`
	EntryStatusCode               StatusType_v02      `json:"entryStatusCode,omitempty"`
	EntryTags                     `json:"entryTags,omitempty"`
	EntryTypeCode                 CodeType_v02      `json:"entryTypeCode,omitempty"`
	OverridePartialShiftIndicator IndicatorType_v01 `json:"overridePartialShiftIndicator,omitempty"`
	OverrideWholeShiftIndicator   IndicatorType_v01 `json:"overrideWholeShiftIndicator,omitempty"`
	PayCode                       CodeType_v02      `json:"payCode,omitempty"`
	PayCodeCategory               CodeType_v02      `json:"payCodeCategory,omitempty"`
	ScheduleEntryID               SimpleIDType_v02  `json:"scheduleEntryID,omitempty"`
	ShiftGroupID                  SimpleIDType_v02  `json:"shiftGroupID,omitempty"`
	ShiftTypeCode                 CodeType_v02      `json:"shiftTypeCode,omitempty"`
	TotalQuantity                 `json:"totalQuantity,omitempty"`
	TotalTime                     `json:"totalTime,omitempty"`
}

// "description": "A collection of days within the schedule",
type ScheduleDay struct {
	Actions           []ActionType_v02      `json:"actions,omitempty"`
	AppliedTempateID  SimpleIDType_v02      `json:"appliedTempateID,omitempty"`
	DayOfWeekCode     CodeType_v02          `json:"dayOfWeekCode,omitempty"`
	DaySequenceNumber SequenceType_v01      `json:"daySequenceNumber,omitempty"`
	ItemID            ItemIDType_v01        `json:"itemID,omitempty"`
	ScheduleDayDate   DateType_v01          `json:"scheduleDayDate,omitempty"`
	ScheduleEntries   []ScheduleEntry       `json:"scheduleEntries,omitempty"`
	ScheduledHours    StandardHoursType_v02 `json:"scheduledHours,omitempty"`
}

// "description": "Collection of the work schedules applicable to a single employee or multiple team members",
type WorkSchedule struct {
	Actions                []ActionType_v02      `json:"actions,omitempty"`
	AppliedTemplateID      SimpleIDType_v02      `json:"appliedTemplateID,omitempty"`
	AssociateOID           AssociateOIDType_v02  `json:"associateOID,omitempty"`
	ClockReferenceDateTime DateTimeType_v01      `json:"clockReferenceDateTime,omitempty"`
	ItemID                 ItemIDType_v01        `json:"itemID,omitempty"`
	Links                  []LinkType_v02        `json:"links,omitempty"`
	ScheduleDays           []ScheduleDay         `json:"scheduleDays,omitempty"`
	ScheduleID             SimpleIDType_v02      `json:"scheduleID,omitempty"`
	ScheduleName           string                `json:"scheduleName,omitempty"`
	SchedulePeriod         DatePeriodType_v01    `json:"schedulePeriod,omitempty"`
	ScheduledHours         StandardHoursType_v02 `json:"scheduledHours,omitempty"`
	WorkAssignmentID       SimpleIDType_v02      `json:"workAssignmentID,omitempty"`
	WorkerName             ContactNameType_v02   `json:"workerName,omitempty"`
}
