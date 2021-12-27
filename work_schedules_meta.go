package adp

type WorkSchedulesMetaResponse struct {
	Meta struct {
		WorkSchedules struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules,omitempty"`
		WorkSchedules_Actions struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/actions,omitempty"`
		WorkSchedules_Actions_ActionTypeCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/actionTypeCode,omitempty"`
		WorkSchedules_Actions_Attestation struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/actions/attestation,omitempty"`
		WorkSchedules_Actions_Attestation_ActionBlockIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/actions/attestation/actionBlockIndicator,omitempty"`
		WorkSchedules_Actions_Attestation_MessageTxt struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/attestation/messageTxt,omitempty"`
		WorkSchedules_Actions_Attestation_MessageUri struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/attestation/messageUri,omitempty"`
		WorkSchedules_Actions_CanonicalUri struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/canonicalUri,omitempty"`
		WorkSchedules_Actions_CommentAllowedIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/actions/commentAllowedIndicator,omitempty"`
		WorkSchedules_Actions_ConfirmationRequiredIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/actions/confirmationRequiredIndicator,omitempty"`
		WorkSchedules_Actions_DefaultIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/actions/defaultIndicator,omitempty"`
		WorkSchedules_Actions_Links struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/actions/links,omitempty"`
		WorkSchedules_Actions_Links_CanonicalUri struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/links/canonicalUri,omitempty"`
		WorkSchedules_Actions_Links_EncType struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/links/encType,omitempty"`
		WorkSchedules_Actions_Links_Href struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/links/href,omitempty"`
		WorkSchedules_Actions_Links_MediaType struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/links/mediaType,omitempty"`
		WorkSchedules_Actions_Links_Method struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/links/method,omitempty"`
		WorkSchedules_Actions_Links_PayLoadArguments struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/actions/links/payLoadArguments,omitempty"`
		WorkSchedules_Actions_Links_PayLoadArguments_ArgumentPath struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/links/payLoadArguments/argumentPath,omitempty"`
		WorkSchedules_Actions_Links_PayLoadArguments_ArgumentValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/links/payLoadArguments/argumentValue,omitempty"`
		WorkSchedules_Actions_Links_Rel struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/links/rel,omitempty"`
		WorkSchedules_Actions_Links_Schema struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/links/schema,omitempty"`
		WorkSchedules_Actions_Links_TargetSchema struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/links/targetSchema,omitempty"`
		WorkSchedules_Actions_Links_Title struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/links/title,omitempty"`
		WorkSchedules_Actions_OperationID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/actions/operationID,omitempty"`
		WorkSchedules_Actions_Sequence struct {
			DataType     string  `json:"dataType,omitempty"`
			Default      float64 `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MaxValue       float64 `json:"maxValue,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			MinValue       float64 `json:"minValue,omitempty"`
			MultipleOf     float64 `json:"multipleOf,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/actions/sequence,omitempty"`
		WorkSchedules_AppliedTemplateID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/appliedTemplateID,omitempty"`
		WorkSchedules_AssociateOID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/associateOID,omitempty"`
		WorkSchedules_ClockReferenceDateTime struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/clockReferenceDateTime,omitempty"`
		WorkSchedules_ItemID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/itemID,omitempty"`
		WorkSchedules_Links struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/links,omitempty"`
		WorkSchedules_Links_CanonicalUri struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/links/canonicalUri,omitempty"`
		WorkSchedules_Links_EncType struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/links/encType,omitempty"`
		WorkSchedules_Links_Href struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/links/href,omitempty"`
		WorkSchedules_Links_MediaType struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/links/mediaType,omitempty"`
		WorkSchedules_Links_Method struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/links/method,omitempty"`
		WorkSchedules_Links_PayLoadArguments struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/links/payLoadArguments,omitempty"`
		WorkSchedules_Links_PayLoadArguments_ArgumentPath struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/links/payLoadArguments/argumentPath,omitempty"`
		WorkSchedules_Links_PayLoadArguments_ArgumentValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/links/payLoadArguments/argumentValue,omitempty"`
		WorkSchedules_Links_Rel struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/links/rel,omitempty"`
		WorkSchedules_Links_Schema struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/links/schema,omitempty"`
		WorkSchedules_Links_TargetSchema struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/links/targetSchema,omitempty"`
		WorkSchedules_Links_Title struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/links/title,omitempty"`
		WorkSchedules_ScheduleDays struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/scheduleDays,omitempty"`
		WorkSchedules_ScheduleDays_Actions struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions,omitempty"`
		WorkSchedules_ScheduleDays_Actions_ActionTypeCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/actionTypeCode,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Attestation struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/attestation,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Attestation_ActionBlockIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/attestation/actionBlockIndicator,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Attestation_MessageTxt struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/attestation/messageTxt,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Attestation_MessageUri struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/attestation/messageUri,omitempty"`
		WorkSchedules_ScheduleDays_Actions_CanonicalUri struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/canonicalUri,omitempty"`
		WorkSchedules_ScheduleDays_Actions_CommentAllowedIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/commentAllowedIndicator,omitempty"`
		WorkSchedules_ScheduleDays_Actions_ConfirmationRequiredIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/confirmationRequiredIndicator,omitempty"`
		WorkSchedules_ScheduleDays_Actions_DefaultIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/defaultIndicator,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_CanonicalUri struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/canonicalUri,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_EncType struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/encType,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_Href struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/href,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_MediaType struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/mediaType,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_Method struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/method,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_PayLoadArguments struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/payLoadArguments,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_PayLoadArguments_ArgumentPath struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/payLoadArguments/argumentPath,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_PayLoadArguments_ArgumentValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/payLoadArguments/argumentValue,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_Rel struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/rel,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_Schema struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/schema,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_TargetSchema struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/targetSchema,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Links_Title struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/links/title,omitempty"`
		WorkSchedules_ScheduleDays_Actions_OperationID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/operationID,omitempty"`
		WorkSchedules_ScheduleDays_Actions_Sequence struct {
			DataType     string  `json:"dataType,omitempty"`
			Default      float64 `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MaxValue       float64 `json:"maxValue,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			MinValue       float64 `json:"minValue,omitempty"`
			MultipleOf     float64 `json:"multipleOf,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/actions/sequence,omitempty"`
		WorkSchedules_ScheduleDays_AppliedTempateID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/appliedTempateID,omitempty"`
		WorkSchedules_ScheduleDays_DayOfWeekCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/dayOfWeekCode,omitempty"`
		WorkSchedules_ScheduleDays_DayOfWeekCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/dayOfWeekCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_DayOfWeekCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/dayOfWeekCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_DayOfWeekCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/dayOfWeekCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_DaySequenceNumber struct {
			DataType     string  `json:"dataType,omitempty"`
			Default      float64 `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MaxValue       float64 `json:"maxValue,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			MinValue       float64 `json:"minValue,omitempty"`
			MultipleOf     float64 `json:"multipleOf,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/daySequenceNumber,omitempty"`
		WorkSchedules_ScheduleDays_ItemID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/itemID,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleDayDate struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleDayDate,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_ActionTypeCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/actionTypeCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Attestation struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/attestation,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Attestation_ActionBlockIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/attestation/actionBlockIndicator,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Attestation_MessageTxt struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/attestation/messageTxt,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Attestation_MessageUri struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/attestation/messageUri,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_CanonicalUri struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/canonicalUri,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_CommentAllowedIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/commentAllowedIndicator,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_ConfirmationRequiredIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/confirmationRequiredIndicator,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_DefaultIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/defaultIndicator,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_CanonicalUri struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/canonicalUri,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_EncType struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/encType,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_Href struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/href,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_MediaType struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/mediaType,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_Method struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/method,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_PayLoadArguments struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/payLoadArguments,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_PayLoadArguments_ArgumentPath struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/payLoadArguments/argumentPath,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_PayLoadArguments_ArgumentValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/payLoadArguments/argumentValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_Rel struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/rel,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_Schema struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/schema,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_TargetSchema struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/targetSchema,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Links_Title struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/links/title,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_OperationID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/operationID,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Actions_Sequence struct {
			DataType     string  `json:"dataType,omitempty"`
			Default      float64 `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MaxValue       float64 `json:"maxValue,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			MinValue       float64 `json:"minValue,omitempty"`
			MultipleOf     float64 `json:"multipleOf,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/actions/sequence,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_CategoryTypeCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/categoryTypeCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_CategoryTypeCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/categoryTypeCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_CategoryTypeCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/categoryTypeCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_CategoryTypeCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/categoryTypeCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DatePeriod struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/datePeriod,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DatePeriod_EndDate struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/datePeriod/endDate,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DatePeriod_EndTime struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/datePeriod/endTime,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DatePeriod_StartDate struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/datePeriod/startDate,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DatePeriod_StartTime struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/datePeriod/startTime,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DateTimePeriod struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/dateTimePeriod,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DateTimePeriod_EndDateTime struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/dateTimePeriod/endDateTime,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DateTimePeriod_StartDateTime struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/dateTimePeriod/startDateTime,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DayPeriodCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/dayPeriodCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DayPeriodCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/dayPeriodCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DayPeriodCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/dayPeriodCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DayPeriodCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/dayPeriodCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_Duration struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/duration,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DurationTypeCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/durationTypeCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DurationTypeCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/durationTypeCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DurationTypeCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/durationTypeCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_DurationTypeCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/durationTypeCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationTypeCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationTypeCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationTypeCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationTypeCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationTypeCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationTypeCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_AllocationTypeCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/allocationTypeCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EarningAllocations_ItemID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/earningAllocations/itemID,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_AppliesToCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/appliesToCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_AppliesToCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/appliesToCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_AppliesToCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/appliesToCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_AppliesToCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/appliesToCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_CommentTypeCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/commentTypeCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_CommentTypeCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/commentTypeCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_CommentTypeCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/commentTypeCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_CommentTypeCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/commentTypeCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_ItemID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/itemID,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_PrivateIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/privateIndicator,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryComments_TextValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryComments/textValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryStatusCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryStatusCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryStatusCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryStatusCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryStatusCode_EffectiveDate struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryStatusCode/effectiveDate,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryStatusCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryStatusCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryStatusCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryStatusCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagType struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagType,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagType_CurrencyCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagType/currencyCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagType_DataTypeCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagType/dataTypeCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagType_FormatCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagType/formatCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryTags_TagValues struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			MaxItems            float64  `json:"maxItems,omitempty"`
			MinItems            float64  `json:"minItems,omitempty"`
			SameValueProperties []string `json:"sameValueProperties,omitempty"`
			UniqueItems         bool     `json:"uniqueItems,omitempty"`
			UniqueProperties    []string `json:"uniqueProperties,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryTags/tagValues,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryTypeCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryTypeCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryTypeCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryTypeCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryTypeCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryTypeCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_EntryTypeCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/entryTypeCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_OverridePartialShiftIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/overridePartialShiftIndicator,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_OverrideWholeShiftIndicator struct {
			DataType     string `json:"dataType,omitempty"`
			Default      bool   `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden         bool    `json:"hidden,omitempty"`
			LongLabelName  string  `json:"longLabelName,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/overrideWholeShiftIndicator,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_PayCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/payCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_PayCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/payCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_PayCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/payCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_PayCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/payCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_PayCodeCategory struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/payCodeCategory,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_PayCodeCategory_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/payCodeCategory/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_PayCodeCategory_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/payCodeCategory/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_PayCodeCategory_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/payCodeCategory/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_ScheduleEntryID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/scheduleEntryID,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_ShiftGroupID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/shiftGroupID,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_ShiftTypeCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/shiftTypeCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_ShiftTypeCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/shiftTypeCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_ShiftTypeCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/shiftTypeCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_ShiftTypeCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/shiftTypeCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity_QuantityValue struct {
			DataType     string  `json:"dataType,omitempty"`
			Default      float64 `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MaxValue       float64 `json:"maxValue,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			MinValue       float64 `json:"minValue,omitempty"`
			MultipleOf     float64 `json:"multipleOf,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity/quantityValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity_UnitTimeCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity/unitTimeCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity_UnitTimeCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity/unitTimeCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity_UnitTimeCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity/unitTimeCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalQuantity_UnitTimeCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalQuantity/unitTimeCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime_NameCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime/nameCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime_NameCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime/nameCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime_NameCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime/nameCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime_NameCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime/nameCode/shortName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduleEntries_TotalTime_TimeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduleEntries/totalTime/timeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduledHours struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduledHours,omitempty"`
		WorkSchedules_ScheduleDays_ScheduledHours_HoursQuantity struct {
			DataType     string  `json:"dataType,omitempty"`
			Default      float64 `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MaxValue       float64 `json:"maxValue,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			MinValue       float64 `json:"minValue,omitempty"`
			MultipleOf     float64 `json:"multipleOf,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduledHours/hoursQuantity,omitempty"`
		WorkSchedules_ScheduleDays_ScheduledHours_UnitCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduledHours/unitCode,omitempty"`
		WorkSchedules_ScheduleDays_ScheduledHours_UnitCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduledHours/unitCode/codeValue,omitempty"`
		WorkSchedules_ScheduleDays_ScheduledHours_UnitCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduledHours/unitCode/longName,omitempty"`
		WorkSchedules_ScheduleDays_ScheduledHours_UnitCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleDays/scheduledHours/unitCode/shortName,omitempty"`
		WorkSchedules_ScheduleID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleID,omitempty"`
		WorkSchedules_ScheduleName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduleName,omitempty"`
		WorkSchedules_SchedulePeriod struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/schedulePeriod,omitempty"`
		WorkSchedules_SchedulePeriod_EndDate struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/schedulePeriod/endDate,omitempty"`
		WorkSchedules_SchedulePeriod_StartDate struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/schedulePeriod/startDate,omitempty"`
		WorkSchedules_ScheduledHours struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/scheduledHours,omitempty"`
		WorkSchedules_ScheduledHours_HoursQuantity struct {
			DataType     string  `json:"dataType,omitempty"`
			Default      float64 `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MaxValue       float64 `json:"maxValue,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			MinValue       float64 `json:"minValue,omitempty"`
			MultipleOf     float64 `json:"multipleOf,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
		} `json:"/workSchedules/scheduledHours/hoursQuantity,omitempty"`
		WorkSchedules_ScheduledHours_UnitCode struct {
			CodeList struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"codeList,omitempty"`
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduledHours/unitCode,omitempty"`
		WorkSchedules_ScheduledHours_UnitCode_CodeValue struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduledHours/unitCode/codeValue,omitempty"`
		WorkSchedules_ScheduledHours_UnitCode_LongName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduledHours/unitCode/longName,omitempty"`
		WorkSchedules_ScheduledHours_UnitCode_ShortName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/scheduledHours/unitCode/shortName,omitempty"`
		WorkSchedules_WorkAssignmentID struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/workAssignmentID,omitempty"`
		WorkSchedules_WorkerName struct {
			DataType     string `json:"dataType,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden   bool    `json:"hidden,omitempty"`
			Optional bool    `json:"optional,omitempty"`
			ReadOnly bool    `json:"readOnly,omitempty"`
			Sequence float64 `json:"sequence,omitempty"`
		} `json:"/workSchedules/workerName,omitempty"`
		WorkSchedules_WorkerName_FamilyName1 struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/workerName/familyName1,omitempty"`
		WorkSchedules_WorkerName_FamilyName2 struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/workerName/familyName2,omitempty"`
		WorkSchedules_WorkerName_FormattedName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/workerName/formattedName,omitempty"`
		WorkSchedules_WorkerName_GivenName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/workerName/givenName,omitempty"`
		WorkSchedules_WorkerName_MiddleName struct {
			DataType     string `json:"dataType,omitempty"`
			Default      string `json:"default,omitempty"`
			Dependencies struct {
			} `json:"dependencies,omitempty"`
			DependencyLinks []struct {
				CanonicalUri     string `json:"canonicalUri,omitempty"`
				EncType          string `json:"encType,omitempty"`
				Href             string `json:"href,omitempty"`
				MediaType        string `json:"mediaType,omitempty"`
				Method           string `json:"method,omitempty"`
				PayLoadArguments []struct {
					ArgumentPath  string `json:"argumentPath,omitempty"`
					ArgumentValue string `json:"argumentValue,omitempty"`
				} `json:"payLoadArguments,omitempty"`
				Rel          string `json:"rel,omitempty"`
				Schema       string `json:"schema,omitempty"`
				TargetSchema string `json:"targetSchema,omitempty"`
				Title        string `json:"title,omitempty"`
			} `json:"dependencyLinks,omitempty"`
			Disallow []struct {
				ItemID string `json:"itemID,omitempty"`
			} `json:"disallow,omitempty"`
			Equals        []string `json:"equals,omitempty"`
			GreaterThan   []string `json:"greaterThan,omitempty"`
			HelperMessage struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				MessageTxt string `json:"messageTxt,omitempty"`
			} `json:"helperMessage,omitempty"`
			Hidden        bool     `json:"hidden,omitempty"`
			LessThan      []string `json:"lessThan,omitempty"`
			LongLabelName string   `json:"longLabelName,omitempty"`
			Masking       struct {
				AuthorizationRequiredIndicator bool   `json:"authorizationRequiredIndicator,omitempty"`
				AuthorizationResourceID        string `json:"authorizationResourceID,omitempty"`
				Links                          []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
			} `json:"masking,omitempty"`
			MaxLength      float64 `json:"maxLength,omitempty"`
			MinLength      float64 `json:"minLength,omitempty"`
			Optional       bool    `json:"optional,omitempty"`
			Pattern        string  `json:"pattern,omitempty"`
			ReadOnly       bool    `json:"readOnly,omitempty"`
			Sequence       float64 `json:"sequence,omitempty"`
			ShortLabelName string  `json:"shortLabelName,omitempty"`
			ValueSet       struct {
				Links []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				Meta struct {
					IdValuePath         string `json:"idValuePath,omitempty"`
					ItemDescriptionPath string `json:"itemDescriptionPath,omitempty"`
				} `json:"meta,omitempty"`
				ValueSetListItems []struct {
					DefaultIndicator bool   `json:"defaultIndicator,omitempty"`
					IdValue          string `json:"idValue,omitempty"`
					ItemDescription  string `json:"itemDescription,omitempty"`
				} `json:"valueSetListItems,omitempty"`
			} `json:"valueSet,omitempty"`
		} `json:"/workSchedules/workerName/middleName,omitempty"`
		QueryCriteria []struct {
			DefaultNumberValue    float64  `json:"defaultNumberValue,omitempty"`
			DefaultStringValue    string   `json:"defaultStringValue,omitempty"`
			ExcludedResourcePaths []string `json:"excludedResourcePaths,omitempty"`
			ItemID                string   `json:"itemID,omitempty"`
			LogicalOperators      []struct {
				LogicalOperatorCode                   string   `json:"logicalOperatorCode,omitempty"`
				MutuallyExclusiveIndicator            bool     `json:"mutuallyExclusiveIndicator,omitempty"`
				MutuallyExclusiveLogicalOperatorCodes []string `json:"mutuallyExclusiveLogicalOperatorCodes,omitempty"`
			} `json:"logicalOperators,omitempty"`
			ObligationCode      string `json:"obligationCode,omitempty"`
			Pattern             string `json:"pattern,omitempty"`
			QueryOptionCode     string `json:"queryOptionCode,omitempty"`
			QueryOptionTypeCode string `json:"queryOptionTypeCode,omitempty"`
			QueryValueCodeList  struct {
				CodeListTitle      string `json:"codeListTitle,omitempty"`
				ExclusiveIndicator bool   `json:"exclusiveIndicator,omitempty"`
				Links              []struct {
					CanonicalUri     string `json:"canonicalUri,omitempty"`
					EncType          string `json:"encType,omitempty"`
					Href             string `json:"href,omitempty"`
					MediaType        string `json:"mediaType,omitempty"`
					Method           string `json:"method,omitempty"`
					PayLoadArguments []struct {
						ArgumentPath  string `json:"argumentPath,omitempty"`
						ArgumentValue string `json:"argumentValue,omitempty"`
					} `json:"payLoadArguments,omitempty"`
					Rel          string `json:"rel,omitempty"`
					Schema       string `json:"schema,omitempty"`
					TargetSchema string `json:"targetSchema,omitempty"`
					Title        string `json:"title,omitempty"`
				} `json:"links,omitempty"`
				ListItems []struct {
					CodeValue         string `json:"codeValue,omitempty"`
					DefaultIndicator  bool   `json:"defaultIndicator,omitempty"`
					ForeignKey        string `json:"foreignKey,omitempty"`
					InactiveIndicator bool   `json:"inactiveIndicator,omitempty"`
					ItemID            string `json:"itemID,omitempty"`
					LongName          string `json:"longName,omitempty"`
					ShortName         string `json:"shortName,omitempty"`
					ValueDescription  string `json:"valueDescription,omitempty"`
				} `json:"listItems,omitempty"`
			} `json:"queryValueCodeList,omitempty"`
			ResourcePathMax float64  `json:"resourcePathMax,omitempty"`
			ResourcePaths   []string `json:"resourcePaths,omitempty"`
		} `json:"queryCriteria,omitempty"`
	} `json:"meta,omitempty"`
}
