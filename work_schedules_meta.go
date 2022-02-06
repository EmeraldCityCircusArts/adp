package adp

type MetaArrayType struct {
	DataType            string                 `json:"dataType,omitempty"`
	Dependencies        MetaDependencyListType `json:"dependencies,omitempty"`
	DependencyLinks     []LinkType_v02         `json:"dependencyLinks,omitempty"`
	HelperMessage       MetaHelperMessageType  `json:"helperMessage,omitempty"`
	MaxItems            float64                `json:"maxItems,omitempty"`
	MinItems            float64                `json:"minItems,omitempty"`
	SameValueProperties []string               `json:"sameValueProperties,omitempty"`
	UniqueItems         bool                   `json:"uniqueItems,omitempty"`
	UniqueProperties    []string               `json:"uniqueProperties,omitempty"`
}

type MetaBooleanType struct {
	DataType        string                 `json:"dataType,omitempty"`
	Default         bool                   `json:"default,omitempty"`
	Dependencies    MetaDependencyListType `json:"dependencies,omitempty"`
	DependencyLinks []LinkType_v02         `json:"dependencyLinks,omitempty"`
	Disallow        []MetaDisallowItemType `json:"disallow,omitempty"`
	Equals          []string               `json:"equals,omitempty"`
	HelperMessage   MetaHelperMessageType  `json:"helperMessage,omitempty"`
	Hidden          bool                   `json:"hidden,omitempty"`
	LongLabelName   string                 `json:"longLabelName,omitempty"`
	Optional        bool                   `json:"optional,omitempty"`
	ReadOnly        bool                   `json:"readOnly,omitempty"`
	Sequence        float64                `json:"sequence,omitempty"`
	ShortLabelName  string                 `json:"shortLabelName,omitempty"`
}

// "description": "List conditional dependencies of supported properties (readOnly, pattern, hidden, disallow etc.) of an attribute or an object. Dependencies can be defined using oneOf (OR) / allOf (AND) object. oneOf - successful if one of the conditions satisfies, allOf - successful if all the conditions satisfies. Refer API specification for dependencies context object syntax and rules. Dependencies context cannot be validated against context schema if dependencies is defined context JSON, because the dependencies JSON structure is dynamically defined for attributes and the not defined in the context schema. Attribute level properties and dependencies properties are mutually exclusive, example - readOnly property can be set either at attribute level or inside dependencies. If both hidden and disallow properties are set inside dependencies or at attribute level, hidden property takes precedence.",
type MetaDependencyListType struct{}

type MetaDisallowItemType struct {
	ItemID string `json:"itemID,omitempty"`
}

// "description": "Helper text provided to the user when they take an action to affirm to be correct, true, or genuine.",
type MetaHelperMessageType struct {
	Links      []LinkType_v02 `json:"links,omitempty"`
	MessageTxt string         `json:"messageTxt,omitempty"`
}

type MetaLogicalOperator struct {
	LogicalOperatorCode                   string   `json:"logicalOperatorCode,omitempty"`
	MutuallyExclusiveIndicator            bool     `json:"mutuallyExclusiveIndicator,omitempty"`
	MutuallyExclusiveLogicalOperatorCodes []string `json:"mutuallyExclusiveLogicalOperatorCodes,omitempty"`
}

type MetaMaskingRuleType struct {
	AuthorizationRequiredIndicator bool           `json:"authorizationRequiredIndicator,omitempty"`
	AuthorizationResourceID        string         `json:"authorizationResourceID,omitempty"`
	Links                          []LinkType_v02 `json:"links,omitempty"`
}

type MetaNumberType struct {
	DataType        string                 `json:"dataType,omitempty"`
	Default         float64                `json:"default,omitempty"`
	Dependencies    MetaDependencyListType `json:"dependencies,omitempty"`
	DependencyLinks []LinkType_v02         `json:"dependencyLinks,omitempty"`
	Disallow        []MetaDisallowItemType `json:"disallow,omitempty"`
	Equals          []string               `json:"equals,omitempty"`
	GreaterThan     []string               `json:"greaterThan,omitempty"`
	HelperMessage   MetaHelperMessageType  `json:"helperMessage,omitempty"`
	Hidden          bool                   `json:"hidden,omitempty"`
	LessThan        []string               `json:"lessThan,omitempty"`
	LongLabelName   string                 `json:"longLabelName,omitempty"`
	Masking         MetaMaskingRuleType    `json:"masking,omitempty"`
	MaxLength       float64                `json:"maxLength,omitempty"`
	MaxValue        float64                `json:"maxValue,omitempty"`
	MinLength       float64                `json:"minLength,omitempty"`
	MinValue        float64                `json:"minValue,omitempty"`
	MultipleOf      float64                `json:"multipleOf,omitempty"`
	Optional        bool                   `json:"optional,omitempty"`
	Pattern         string                 `json:"pattern,omitempty"`
	ReadOnly        bool                   `json:"readOnly,omitempty"`
	Sequence        float64                `json:"sequence,omitempty"`
	ShortLabelName  string                 `json:"shortLabelName,omitempty"`
}

type MetaObjectType struct {
	DataType        string                 `json:"dataType,omitempty"`
	Dependencies    MetaDependencyListType `json:"dependencies,omitempty"`
	DependencyLinks []LinkType_v02         `json:"dependencyLinks,omitempty"`
	Disallow        []MetaDisallowItemType `json:"disallow,omitempty"`
	HelperMessage   MetaHelperMessageType  `json:"helperMessage,omitempty"`
	Hidden          bool                   `json:"hidden,omitempty"`
	Optional        bool                   `json:"optional,omitempty"`
	ReadOnly        bool                   `json:"readOnly,omitempty"`
	Sequence        float64                `json:"sequence,omitempty"`
}

type MetaQueryCriteria struct {
	DefaultNumberValue    float64               `json:"defaultNumberValue,omitempty"`
	DefaultStringValue    string                `json:"defaultStringValue,omitempty"`
	ExcludedResourcePaths []string              `json:"excludedResourcePaths,omitempty"`
	ItemID                string                `json:"itemID,omitempty"`
	LogicalOperators      []MetaLogicalOperator `json:"logicalOperators,omitempty"`
	ObligationCode        string                `json:"obligationCode,omitempty"`
	Pattern               string                `json:"pattern,omitempty"`
	QueryOptionCode       string                `json:"queryOptionCode,omitempty"`
	QueryOptionTypeCode   string                `json:"queryOptionTypeCode,omitempty"`
	QueryValueCodeList    CodeListType_v03      `json:"queryValueCodeList,omitempty"`
	ResourcePathMax       float64               `json:"resourcePathMax,omitempty"`
	ResourcePaths         []string              `json:"resourcePaths,omitempty"`
}

type MetaStringCodeListType struct {
	CodeList        CodeListType_v03       `json:"codeList,omitempty"`
	DataType        string                 `json:"dataType,omitempty"`
	Default         string                 `json:"default,omitempty"`
	Dependencies    MetaDependencyListType `json:"dependencies,omitempty"`
	DependencyLinks []LinkType_v02         `json:"dependencyLinks,omitempty"`
	Disallow        []MetaDisallowItemType `json:"disallow,omitempty"`
	Equals          []string               `json:"equals,omitempty"`
	GreaterThan     []string               `json:"greaterThan,omitempty"`
	HelperMessage   MetaHelperMessageType  `json:"helperMessage,omitempty"`
	Hidden          bool                   `json:"hidden,omitempty"`
	LessThan        []string               `json:"lessThan,omitempty"`
	LongLabelName   string                 `json:"longLabelName,omitempty"`
	Masking         MetaMaskingRuleType    `json:"masking,omitempty"`
	MaxLength       float64                `json:"maxLength,omitempty"`
	MinLength       float64                `json:"minLength,omitempty"`
	Optional        bool                   `json:"optional,omitempty"`
	Pattern         string                 `json:"pattern,omitempty"`
	ReadOnly        bool                   `json:"readOnly,omitempty"`
	Sequence        float64                `json:"sequence,omitempty"`
	ShortLabelName  string                 `json:"shortLabelName,omitempty"`
	ValueSet        MetaValueSetType       `json:"valueSet,omitempty"`
}

type MetaStringType struct {
	DataType        string                 `json:"dataType,omitempty"`
	Default         string                 `json:"default,omitempty"`
	Dependencies    MetaDependencyListType `json:"dependencies,omitempty"`
	DependencyLinks []LinkType_v02         `json:"dependencyLinks,omitempty"`
	Disallow        []MetaDisallowItemType `json:"disallow,omitempty"`
	Equals          []string               `json:"equals,omitempty"`
	GreaterThan     []string               `json:"greaterThan,omitempty"`
	HelperMessage   MetaHelperMessageType  `json:"helperMessage,omitempty"`
	Hidden          bool                   `json:"hidden,omitempty"`
	LessThan        []string               `json:"lessThan,omitempty"`
	LongLabelName   string                 `json:"longLabelName,omitempty"`
	Masking         MetaMaskingRuleType    `json:"masking,omitempty"`
	MaxLength       float64                `json:"maxLength,omitempty"`
	MinLength       float64                `json:"minLength,omitempty"`
	Optional        bool                   `json:"optional,omitempty"`
	Pattern         string                 `json:"pattern,omitempty"`
	ReadOnly        bool                   `json:"readOnly,omitempty"`
	Sequence        float64                `json:"sequence,omitempty"`
	ShortLabelName  string                 `json:"shortLabelName,omitempty"`
	ValueSet        MetaValueSetType       `json:"valueSet,omitempty"`
}

// "description": "The value set item type",
type MetaValueSetItemType struct {
	DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
	IdValue          string `json:"idValue,omitempty"`
	ItemDescription  string `json:"itemDescription,omitempty"`
}

type MetaValueSetMeta struct {
	IdValuePath         string `json:"idValuePath,omitempty"`
	ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
}

// "description": "A reference to the codelist extended with values",
type MetaValueSetType struct {
	Links             []LinkType_v02         `json:"links,omitempty"`
	Meta              MetaValueSetMeta       `json:"meta,omitempty"`
	ValueSetListItems []MetaValueSetItemType `json:"valueSetListItems,omitempty"`
}

type WorkSchedulesMeta struct {
	WorkSchedules                                                                              MetaArrayType          `json:"/workSchedules,omitempty"`
	WorkSchedules_Actions                                                                      MetaArrayType          `json:"/workSchedules/actions,omitempty"`
	WorkSchedules_Actions_ActionTypeCode                                                       MetaStringCodeListType `json:"/workSchedules/actions/actionTypeCode,omitempty"`
	WorkSchedules_Actions_Attestation                                                          MetaObjectType         `json:"/workSchedules/actions/attestation,omitempty"`
	WorkSchedules_Actions_Attestation_ActionBlockIndicator                                     MetaBooleanType        `json:"/workSchedules/actions/attestation/actionBlockIndicator,omitempty"`
	WorkSchedules_Actions_Attestation_MessageTxt                                               MetaStringType         `json:"/workSchedules/actions/attestation/messageTxt,omitempty"`
	WorkSchedules_Actions_Attestation_MessageUri                                               MetaStringType         `json:"/workSchedules/actions/attestation/messageUri,omitempty"`
	WorkSchedules_Actions_CanonicalUri                                                         MetaStringType         `json:"/workSchedules/actions/canonicalUri,omitempty"`
	WorkSchedules_Actions_CommentAllowedIndicator                                              MetaBooleanType        `json:"/workSchedules/actions/commentAllowedIndicator,omitempty"`
	WorkSchedules_Actions_ConfirmationRequiredIndicator                                        MetaBooleanType        `json:"/workSchedules/actions/confirmationRequiredIndicator,omitempty"`
	WorkSchedules_Actions_DefaultIndicator                                                     MetaBooleanType        `json:"/workSchedules/actions/defaultIndicator,omitempty"`
	WorkSchedules_Actions_Links                                                                MetaArrayType          `json:"/workSchedules/actions/links,omitempty"`
	WorkSchedules_Actions_Links_CanonicalUri                                                   MetaStringType         `json:"/workSchedules/actions/links/canonicalUri,omitempty"`
	WorkSchedules_Actions_Links_EncType                                                        MetaStringType         `json:"/workSchedules/actions/links/encType,omitempty"`
	WorkSchedules_Actions_Links_Href                                                           MetaStringType         `json:"/workSchedules/actions/links/href,omitempty"`
	WorkSchedules_Actions_Links_MediaType                                                      MetaStringType         `json:"/workSchedules/actions/links/mediaType,omitempty"`
	WorkSchedules_Actions_Links_Method                                                         MetaStringType         `json:"/workSchedules/actions/links/method,omitempty"`
	WorkSchedules_Actions_Links_PayLoadArguments                                               MetaArrayType          `json:"/workSchedules/actions/links/payLoadArguments,omitempty"`
	WorkSchedules_Actions_Links_PayLoadArguments_ArgumentPath                                  MetaStringType         `json:"/workSchedules/actions/links/payLoadArguments/argumentPath,omitempty"`
	WorkSchedules_Actions_Links_PayLoadArguments_ArgumentValue                                 MetaStringType         `json:"/workSchedules/actions/links/payLoadArguments/argumentValue,omitempty"`
	WorkSchedules_Actions_Links_Rel                                                            MetaStringType         `json:"/workSchedules/actions/links/rel,omitempty"`
	WorkSchedules_Actions_Links_Schema                                                         MetaStringType         `json:"/workSchedules/actions/links/schema,omitempty"`
	WorkSchedules_Actions_Links_TargetSchema                                                   MetaStringType         `json:"/workSchedules/actions/links/targetSchema,omitempty"`
	WorkSchedules_Actions_Links_Title                                                          MetaStringType         `json:"/workSchedules/actions/links/title,omitempty"`
	WorkSchedules_Actions_OperationID                                                          MetaStringType         `json:"/workSchedules/actions/operationID,omitempty"`
	WorkSchedules_Actions_Sequence                                                             MetaNumberType         `json:"/workSchedules/actions/sequence,omitempty"`
	WorkSchedules_AppliedTemplateID                                                            MetaStringType         `json:"/workSchedules/appliedTemplateID,omitempty"`
	WorkSchedules_AssociateOID                                                                 MetaStringType         `json:"/workSchedules/associateOID,omitempty"`
	WorkSchedules_ClockReferenceDateTime                                                       MetaStringType         `json:"/workSchedules/clockReferenceDateTime,omitempty"`
	WorkSchedules_ItemID                                                                       MetaStringType         `json:"/workSchedules/itemID,omitempty"`
	WorkSchedules_Links                                                                        MetaArrayType          `json:"/workSchedules/links,omitempty"`
	WorkSchedules_Links_CanonicalUri                                                           MetaStringType         `json:"/workSchedules/links/canonicalUri,omitempty"`
	WorkSchedules_Links_EncType                                                                MetaStringType         `json:"/workSchedules/links/encType,omitempty"`
	WorkSchedules_Links_Href                                                                   MetaStringType         `json:"/workSchedules/links/href,omitempty"`
	WorkSchedules_Links_MediaType                                                              MetaStringType         `json:"/workSchedules/links/mediaType,omitempty"`
	WorkSchedules_Links_Method                                                                 MetaStringType         `json:"/workSchedules/links/method,omitempty"`
	WorkSchedules_Links_PayLoadArguments                                                       MetaArrayType          `json:"/workSchedules/links/payLoadArguments,omitempty"`
	WorkSchedules_Links_PayLoadArguments_ArgumentPath                                          MetaStringType         `json:"/workSchedules/links/payLoadArguments/argumentPath,omitempty"`
	WorkSchedules_Links_PayLoadArguments_ArgumentValue                                         MetaStringType         `json:"/workSchedules/links/payLoadArguments/argumentValue,omitempty"`
	WorkSchedules_Links_Rel                                                                    MetaStringType         `json:"/workSchedules/links/rel,omitempty"`
	WorkSchedules_Links_Schema                                                                 MetaStringType         `json:"/workSchedules/links/schema,omitempty"`
	WorkSchedules_Links_TargetSchema                                                           MetaStringType         `json:"/workSchedules/links/targetSchema,omitempty"`
	WorkSchedules_Links_Title                                                                  MetaStringType         `json:"/workSchedules/links/title,omitempty"`
	WorkSchedules_ScheduleDays                                                                 MetaArrayType          `json:"/workSchedules/scheduleDays,omitempty"`
	WorkSchedules_ScheduleDays_Actions                                                         MetaArrayType          `json:"/workSchedules/scheduleDays/actions,omitempty"`
	WorkSchedules_ScheduleDays_Actions_ActionTypeCode                                          MetaStringCodeListType `json:"/workSchedules/scheduleDays/actions/actionTypeCode,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Attestation                                             MetaObjectType         `json:"/workSchedules/scheduleDays/actions/attestation,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Attestation_ActionBlockIndicator                        MetaBooleanType        `json:"/workSchedules/scheduleDays/actions/attestation/actionBlockIndicator,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Attestation_MessageTxt                                  MetaStringType         `json:"/workSchedules/scheduleDays/actions/attestation/messageTxt,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Attestation_MessageUri                                  MetaStringType         `json:"/workSchedules/scheduleDays/actions/attestation/messageUri,omitempty"`
	WorkSchedules_ScheduleDays_Actions_CanonicalUri                                            MetaStringType         `json:"/workSchedules/scheduleDays/actions/canonicalUri,omitempty"`
	WorkSchedules_ScheduleDays_Actions_CommentAllowedIndicator                                 MetaBooleanType        `json:"/workSchedules/scheduleDays/actions/commentAllowedIndicator,omitempty"`
	WorkSchedules_ScheduleDays_Actions_ConfirmationRequiredIndicator                           MetaBooleanType        `json:"/workSchedules/scheduleDays/actions/confirmationRequiredIndicator,omitempty"`
	WorkSchedules_ScheduleDays_Actions_DefaultIndicator                                        MetaBooleanType        `json:"/workSchedules/scheduleDays/actions/defaultIndicator,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links                                                   MetaArrayType          `json:"/workSchedules/scheduleDays/actions/links,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_CanonicalUri                                      MetaStringType         `json:"/workSchedules/scheduleDays/actions/links/canonicalUri,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_EncType                                           MetaStringType         `json:"/workSchedules/scheduleDays/actions/links/encType,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_Href                                              MetaStringType         `json:"/workSchedules/scheduleDays/actions/links/href,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_MediaType                                         MetaStringType         `json:"/workSchedules/scheduleDays/actions/links/mediaType,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_Method                                            MetaStringType         `json:"/workSchedules/scheduleDays/actions/links/method,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_PayLoadArguments                                  MetaArrayType          `json:"/workSchedules/scheduleDays/actions/links/payLoadArguments,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_PayLoadArguments_ArgumentPath                     MetaStringType         `json:"/workSchedules/scheduleDays/actions/links/payLoadArguments/argumentPath,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_PayLoadArguments_ArgumentValue                    MetaStringType         `json:"/workSchedules/scheduleDays/actions/links/payLoadArguments/argumentValue,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_Rel                                               MetaStringType         `json:"/workSchedules/scheduleDays/actions/links/rel,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_Schema                                            MetaStringType         `json:"/workSchedules/scheduleDays/actions/links/schema,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_TargetSchema                                      MetaStringType         `json:"/workSchedules/scheduleDays/actions/links/targetSchema,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Links_Title                                             MetaStringType         `json:"/workSchedules/scheduleDays/actions/links/title,omitempty"`
	WorkSchedules_ScheduleDays_Actions_OperationID                                             MetaStringType         `json:"/workSchedules/scheduleDays/actions/operationID,omitempty"`
	WorkSchedules_ScheduleDays_Actions_Sequence                                                MetaNumberType         `json:"/workSchedules/scheduleDays/actions/sequence,omitempty"`
	WorkSchedules_ScheduleDays_AppliedTempateID                                                MetaStringType         `json:"/workSchedules/scheduleDays/appliedTempateID,omitempty"`
	WorkSchedules_ScheduleDays_DayOfWeekCode                                                   MetaStringCodeListType `json:"/workSchedules/scheduleDays/dayOfWeekCode,omitempty"`
	WorkSchedules_ScheduleDays_DayOfWeekCode_CodeValue                                         MetaStringType         `json:"/workSchedules/scheduleDays/dayOfWeekCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_DayOfWeekCode_LongName                                          MetaStringType         `json:"/workSchedules/scheduleDays/dayOfWeekCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_DayOfWeekCode_ShortName                                         MetaStringType         `json:"/workSchedules/scheduleDays/dayOfWeekCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_DaySequenceNumber                                               MetaNumberType         `json:"/workSchedules/scheduleDays/daySequenceNumber,omitempty"`
	WorkSchedules_ScheduleDays_ItemID                                                          MetaStringType         `json:"/workSchedules/scheduleDays/itemID,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleDayDate                                                 MetaStringType         `json:"/workSchedules/scheduleDays/scheduleDayDate,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries                                                 MetaArrayType          `json:"/workSchedules/scheduleDays/scheduleEntries,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions                                         MetaArrayType          `json:"/workSchedules/scheduleDays/scheduleEntries/actions,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_ActionTypeCode                          MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/actions/actionTypeCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Attestation                             MetaObjectType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/attestation,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Attestation_ActionBlockIndicator        MetaBooleanType        `json:"/workSchedules/scheduleDays/scheduleEntries/actions/attestation/actionBlockIndicator,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Attestation_MessageTxt                  MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/attestation/messageTxt,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Attestation_MessageUri                  MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/attestation/messageUri,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_CanonicalUri                            MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/canonicalUri,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_CommentAllowedIndicator                 MetaBooleanType        `json:"/workSchedules/scheduleDays/scheduleEntries/actions/commentAllowedIndicator,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_ConfirmationRequiredIndicator           MetaBooleanType        `json:"/workSchedules/scheduleDays/scheduleEntries/actions/confirmationRequiredIndicator,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_DefaultIndicator                        MetaBooleanType        `json:"/workSchedules/scheduleDays/scheduleEntries/actions/defaultIndicator,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links                                   MetaArrayType          `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_CanonicalUri                      MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/canonicalUri,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_EncType                           MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/encType,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_Href                              MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/href,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_MediaType                         MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/mediaType,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_Method                            MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/method,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_PayLoadArguments                  MetaArrayType          `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/payLoadArguments,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_PayLoadArguments_ArgumentPath     MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/payLoadArguments/argumentPath,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_PayLoadArguments_ArgumentValue    MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/payLoadArguments/argumentValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_Rel                               MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/rel,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_Schema                            MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/schema,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_TargetSchema                      MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/targetSchema,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_Title                             MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/title,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_OperationID                             MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/operationID,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Sequence                                MetaNumberType         `json:"/workSchedules/scheduleDays/scheduleEntries/actions/sequence,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_CategoryTypeCode                                MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/categoryTypeCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_CategoryTypeCode_CodeValue                      MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/categoryTypeCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_CategoryTypeCode_LongName                       MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/categoryTypeCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_CategoryTypeCode_ShortName                      MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/categoryTypeCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DatePeriod                                      MetaObjectType         `json:"/workSchedules/scheduleDays/scheduleEntries/datePeriod,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DatePeriod_EndDate                              MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/datePeriod/endDate,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DatePeriod_EndTime                              MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/datePeriod/endTime,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DatePeriod_StartDate                            MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/datePeriod/startDate,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DatePeriod_StartTime                            MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/datePeriod/startTime,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DateTimePeriod                                  MetaObjectType         `json:"/workSchedules/scheduleDays/scheduleEntries/dateTimePeriod,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DateTimePeriod_EndDateTime                      MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/dateTimePeriod/endDateTime,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DateTimePeriod_StartDateTime                    MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/dateTimePeriod/startDateTime,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DayPeriodCode                                   MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/dayPeriodCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DayPeriodCode_CodeValue                         MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/dayPeriodCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DayPeriodCode_LongName                          MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/dayPeriodCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DayPeriodCode_ShortName                         MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/dayPeriodCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_Duration                                        MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/duration,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DurationTypeCode                                MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/durationTypeCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DurationTypeCode_CodeValue                      MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/durationTypeCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DurationTypeCode_LongName                       MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/durationTypeCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_DurationTypeCode_ShortName                      MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/durationTypeCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations                              MetaArrayType          `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationCode               MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationCode_CodeValue     MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationCode_LongName      MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationCode_ShortName     MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationTypeCode           MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationTypeCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationTypeCode_CodeValue MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationTypeCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationTypeCode_LongName  MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationTypeCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationTypeCode_ShortName MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationTypeCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_ItemID                       MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/itemID,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments                                   MetaArrayType          `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_AppliesToCode                     MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/appliesToCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_AppliesToCode_CodeValue           MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/appliesToCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_AppliesToCode_LongName            MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/appliesToCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_AppliesToCode_ShortName           MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/appliesToCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_CommentTypeCode                   MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/commentTypeCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_CommentTypeCode_CodeValue         MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/commentTypeCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_CommentTypeCode_LongName          MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/commentTypeCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_CommentTypeCode_ShortName         MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/commentTypeCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_ItemID                            MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/itemID,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_PrivateIndicator                  MetaBooleanType        `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/privateIndicator,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_TextValue                         MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/textValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryStatusCode                                 MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/entryStatusCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryStatusCode_CodeValue                       MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryStatusCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryStatusCode_EffectiveDate                   MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryStatusCode/effectiveDate,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryStatusCode_LongName                        MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryStatusCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryStatusCode_ShortName                       MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryStatusCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags                                       MetaObjectType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagCode                               MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagType                               MetaObjectType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagType,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagType_CurrencyCode                  MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagType/currencyCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagType_DataTypeCode                  MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagType/dataTypeCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagType_FormatCode                    MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagType/formatCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagValues                             MetaArrayType          `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagValues,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryTypeCode                                   MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/entryTypeCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryTypeCode_CodeValue                         MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryTypeCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryTypeCode_LongName                          MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryTypeCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_EntryTypeCode_ShortName                         MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/entryTypeCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_OverridePartialShiftIndicator                   MetaBooleanType        `json:"/workSchedules/scheduleDays/scheduleEntries/overridePartialShiftIndicator,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_OverrideWholeShiftIndicator                     MetaBooleanType        `json:"/workSchedules/scheduleDays/scheduleEntries/overrideWholeShiftIndicator,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_PayCode                                         MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/payCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_PayCode_CodeValue                               MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/payCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_PayCode_LongName                                MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/payCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_PayCode_ShortName                               MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/payCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_PayCodeCategory                                 MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/payCodeCategory,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_PayCodeCategory_CodeValue                       MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/payCodeCategory/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_PayCodeCategory_LongName                        MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/payCodeCategory/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_PayCodeCategory_ShortName                       MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/payCodeCategory/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_ScheduleEntryID                                 MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/scheduleEntryID,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_ShiftGroupID                                    MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/shiftGroupID,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_ShiftTypeCode                                   MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/shiftTypeCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_ShiftTypeCode_CodeValue                         MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/shiftTypeCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_ShiftTypeCode_LongName                          MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/shiftTypeCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_ShiftTypeCode_ShortName                         MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/shiftTypeCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity                                   MetaObjectType         `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity_QuantityValue                     MetaNumberType         `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity/quantityValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity_UnitTimeCode                      MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity/unitTimeCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity_UnitTimeCode_CodeValue            MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity/unitTimeCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity_UnitTimeCode_LongName             MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity/unitTimeCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity_UnitTimeCode_ShortName            MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity/unitTimeCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime                                       MetaObjectType         `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime_NameCode                              MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime/nameCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime_NameCode_CodeValue                    MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime/nameCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime_NameCode_LongName                     MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime/nameCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime_NameCode_ShortName                    MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime/nameCode/shortName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime_TimeValue                             MetaStringType         `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime/timeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduledHours                                                  MetaObjectType         `json:"/workSchedules/scheduleDays/scheduledHours,omitempty"`
	WorkSchedules_ScheduleDays_ScheduledHours_HoursQuantity                                    MetaNumberType         `json:"/workSchedules/scheduleDays/scheduledHours/hoursQuantity,omitempty"`
	WorkSchedules_ScheduleDays_ScheduledHours_UnitCode                                         MetaStringCodeListType `json:"/workSchedules/scheduleDays/scheduledHours/unitCode,omitempty"`
	WorkSchedules_ScheduleDays_ScheduledHours_UnitCode_CodeValue                               MetaStringType         `json:"/workSchedules/scheduleDays/scheduledHours/unitCode/codeValue,omitempty"`
	WorkSchedules_ScheduleDays_ScheduledHours_UnitCode_LongName                                MetaStringType         `json:"/workSchedules/scheduleDays/scheduledHours/unitCode/longName,omitempty"`
	WorkSchedules_ScheduleDays_ScheduledHours_UnitCode_ShortName                               MetaStringType         `json:"/workSchedules/scheduleDays/scheduledHours/unitCode/shortName,omitempty"`
	WorkSchedules_ScheduleID                                                                   MetaStringType         `json:"/workSchedules/scheduleID,omitempty"`
	WorkSchedules_ScheduleName                                                                 MetaStringType         `json:"/workSchedules/scheduleName,omitempty"`
	WorkSchedules_SchedulePeriod                                                               MetaObjectType         `json:"/workSchedules/schedulePeriod,omitempty"`
	WorkSchedules_SchedulePeriod_EndDate                                                       MetaStringType         `json:"/workSchedules/schedulePeriod/endDate,omitempty"`
	WorkSchedules_SchedulePeriod_StartDate                                                     MetaStringType         `json:"/workSchedules/schedulePeriod/startDate,omitempty"`
	WorkSchedules_ScheduledHours                                                               MetaObjectType         `json:"/workSchedules/scheduledHours,omitempty"`
	WorkSchedules_ScheduledHours_HoursQuantity                                                 MetaNumberType         `json:"/workSchedules/scheduledHours/hoursQuantity,omitempty"`
	WorkSchedules_ScheduledHours_UnitCode                                                      MetaStringCodeListType `json:"/workSchedules/scheduledHours/unitCode,omitempty"`
	WorkSchedules_ScheduledHours_UnitCode_CodeValue                                            MetaStringType         `json:"/workSchedules/scheduledHours/unitCode/codeValue,omitempty"`
	WorkSchedules_ScheduledHours_UnitCode_LongName                                             MetaStringType         `json:"/workSchedules/scheduledHours/unitCode/longName,omitempty"`
	WorkSchedules_ScheduledHours_UnitCode_ShortName                                            MetaStringType         `json:"/workSchedules/scheduledHours/unitCode/shortName,omitempty"`
	WorkSchedules_WorkAssignmentID                                                             MetaStringType         `json:"/workSchedules/workAssignmentID,omitempty"`
	WorkSchedules_WorkerName                                                                   MetaObjectType         `json:"/workSchedules/workerName,omitempty"`
	WorkSchedules_WorkerName_FamilyName1                                                       MetaStringType         `json:"/workSchedules/workerName/familyName1,omitempty"`
	WorkSchedules_WorkerName_FamilyName2                                                       MetaStringType         `json:"/workSchedules/workerName/familyName2,omitempty"`
	WorkSchedules_WorkerName_FormattedName                                                     MetaStringType         `json:"/workSchedules/workerName/formattedName,omitempty"`
	WorkSchedules_WorkerName_GivenName                                                         MetaStringType         `json:"/workSchedules/workerName/givenName,omitempty"`
	WorkSchedules_WorkerName_MiddleName                                                        MetaStringType         `json:"/workSchedules/workerName/middleName,omitempty"`
	QueryCriteria                                                                              []MetaQueryCriteria    `json:"queryCriteria,omitempty"`
}

type WorkSchedulesMetaResponse struct {
	Meta WorkSchedulesMeta `json:"meta,omitempty"`
}
