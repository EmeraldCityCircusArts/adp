package adp

// "description": "Indicates the nature of the relationship of the related resource to the resource that generated this representation",
/* "enum": [
  "alternate",
  "create",
  "canonical",
  "describedby",
  "edit-form",
  "enclosure",
  "full",
  "related",
  "root",
  "self",
  "up",
  "search",
  "first",
  "last",
  "next",
  "previous",
  "/adp/invoke",
  "/adp/image",
  "/adp/confirm-message",
  "/adp/status-monitor",
  "/adp/codelist",
  "/adp/template",
  "/adp/externalLink",
  "/adp/validation",
  "/adp/deeplink",
  "/adp/attachment",
  "/adp/workflow"
]
*/
type Rel string

// "description": "The media type that the linked resource will return (Response)",
/* "enum": [
     "application/gzip",
     "application/json",
     "application/msword",
     "application/pdf",
     "application/postscript",
     "application/vnd.ms-excel",
     "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
     "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
     "application/xml",
     "application/x-www-form-urlencoded",
     "image/gif",
     "image/jpeg",
     "image/png",
     "image/tiff",
     "multipart/mixed",
     "text/html",
     "text/plain",
     "application/vnd.visio",
     "image/bmp",
     "image/x-ms-bmp",
     "application/vnd.openxmlformats-officedocument.presentationml.presentation",
     "application/vnd.ms-powerpoint",
     "video/mp4",
     "audio/mpeg",
     "video/x-msvideo",
     "video/x-ms-wmv",
     "application/rtf",
     "application/vnd.ms-outlook",
     "text/csv",
     "video/quicktime",
     "application/zip",
     "application/illustrator",
     "text/xml"
   ]
*/
type MediaType string

// "description": "The HTTP method code (HTTP Verb) traverse the link (GET, POST, PUT or DELETE)",
/* "enum": [
     "GET",
     "POST",
     "DELETE",
     "PUT"
   ]
*/
type Method string

// "description": "If present, this property indicates a query media type format that the server supports for querying or posting to the collection of instances at the target resource.  The query can be suffixed to the target URI to query the collection with property-based constraints on the resources that SHOULD be returned from the server or used to post data to the resource (depending on the method). (Request)",
/* "enum": [
     "application/json",
     "application/x-www-form-urlencoded"
   ]
*/
type EncType string

type PayLoadArgument struct {
	ArgumentPath  string `json:"argumentPath,omitempty"`
	ArgumentValue string `json:"argumentValue,omitempty"`
}

// "$ref": "#/definitions/messageTypeCodeType_v01"
// "description": "Process Message [codeValue] instances may be of type: success, warning, error, or info.",
type MessageTypeCode struct {
	CodeValue string `json:"codeValue,omitempty"`
	LongName  string `json:"longName,omitempty"`
	ShortName string `json:"shortName,omitempty"`
}

// "$ref": "#/definitions/processMessage_v01"
// "description": "processMessage is optional and contains processing information for either the request (at the ConfirmMessage level) or a resource being managed in the request (at the ResourceMessage level).  Most often there will be a single instance for an associated request or resource, but this structure allows for more than one if needed, for example, if multiple errors exist for a single resource. "
type ProcessMessage struct {
	DeveloperMessage         MessageType_v01 `json:"developerMessage,omitempty"`
	ExpressionLanguageCode   GenericCode     `json:"expressionLanguageCode,omitempty"`
	Links                    []LinkType_v02  `json:"links,omitempty"`
	MessageTypeCode          `json:"messageTypeCode,omitempty"`
	ProcessMessageID         IDType_v01      `json:"processMessageID,omitempty"`
	SourceLocationExpression string          `json:"sourceLocationExpression,omitempty"`
	UserMessage              MessageType_v01 `json:"userMessage,omitempty"`
}

// "description": "resource Message contains the resource-specific processing results for resources being managed in the request.  A resource may have its processing reported as:  succeeded or failed.  It must be used to represent resource-specific messages.  In the case of a request managing multiple resources and resource-specific messages are to be returned, then this array will contain one object for each such resource of the request.",
type ResourceMessage struct {
	ProcessMessages    []ProcessMessage `json:"processMessages,omitempty"`
	ResourceLink       LinkType_v02     `json:"resourceLink,omitempty"`
	ResourceMessageID  IDType_v01       `json:"resourceMessageID,omitempty"`
	ResourceStatusCode GenericCode      `json:"resourceStatusCode,omitempty"`
}

type WorkSchedulesResponse struct {
	ConfirmMessage ConfirmMessageType_v01 `json:"confirmMessage,omitempty"`
	Meta           MetaType_v02           `json:"meta,omitempty"`
	WorkSchedules  []WorkSchedule         `json:"workSchedules,omitempty"`
}

type WorkerResponse struct {
	ConfirmMessage ConfirmMessageType_v01 `json:"confirmMessage,omitempty"`
	Meta           MetaType_v02           `json:"meta,omitempty"`
	Workers        []Worker               `json:"workers,omitempty"`
}
