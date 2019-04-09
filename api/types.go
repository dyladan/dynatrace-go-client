package api

type NameDetectionResponse struct {
	Values []NameDetectionRule `json:"values"`
}

type NameDetectionRule struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type NameDetectionRuleDetail struct {
	Id                    string                               `json:"id,omitempty"`
	Metadata              NameDetectionRuleMetadata            `json:"metadata"`
	ApplicationIdentifier string                               `json:"applicationIdentifier"`
	FilterConfig          NameDetectionRuleFilterConfiguration `json:"filterConfig"`
}

type NameDetectionRuleMetadata struct {
	ClusterVersion        string  `json:"clusterVersion"`
	ConfigurationVersions []int64 `json:"configurationVersions"`
}

type NameDetectionRuleFilterConfiguration struct {
	Pattern                string `json:"pattern"`
	ApplicationMatchType   string `json:"applicationMatchType"`
	ApplicationMatchTarget string `json:"applicationMatchTarget"`
}

type NameDetectionRuleOrderRequest struct {
	Values []NameDetectionRule `json:"values"`
}

// AutoTags
type AutoTag struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Rules       []AutoTagRule `json:"rules"`
}

type AutoTagResponse struct {
	Values []AutoTag `json:"values"`
}

type AutoTagRule struct {
	Type             AutoTagRuleType              `json:"type"`
	Enabled          bool                         `json:"enabled"`
	ValueFormat      string                       `json:"valueFormat"`
	PropagationTypes []AutoTagRulePropagationType `json:"propagationTypes"`
	Conditions       []AutoTagRuleCondition       `json:"conditions"`
}

type AutoTagRuleType string

const (
	AutoTagRuleTypeCustomDevice AutoTagRuleType = "CUSTOM_DEVICE"
	AutoTagRuleTypeHost                         = "HOST"
	AutoTagRuleTypeProcessGroup                 = "PROCESS_GROUP"
	AutoTagRuleTypeService                      = "SERVICE"
)

type AutoTagRulePropagationType string

const (
	ServiceToProcessGroupLike  AutoTagRulePropagationType = "SERVICE_TO_PROCESS_GROUP_LIKE"
	ServiceToHostLike          AutoTagRulePropagationType = "SERVICE_TO_HOST_LIKE"
	ProcessGroupToHost         AutoTagRulePropagationType = "PROCESS_GROUP_TO_HOST"
	ProcessGroupToService      AutoTagRulePropagationType = "PROCESS_GROUP_TO_SERVICE"
	HostToProcessGroupInstance AutoTagRulePropagationType = "HOST_TO_PROCESS_GROUP_INSTANCE"
)

type AutoTagRuleConditionKey struct {
	Attribute string `json:"attribute"`
}

type AutoTagRuleConditionComparisonInfo struct {
	Type          string `json:"type"`
	Operator      string `json:"operator"`
	Value         string `json:"value"`
	Negate        bool   `json:"negate"`
	CaseSensitive bool   `json:"caseSensitive"`
}

type AutoTagRuleCondition struct {
	Key            AutoTagRuleConditionKey            `json:"key"`
	ComparisonInfo AutoTagRuleConditionComparisonInfo `json:"comparisonInfo"`
}

type ConstraintViolation struct {
	Path              string                               `json:"path,omitempty"`
	Message           string                               `json:"message,omitempty"`
	ParameterLocation ConstraintViolationParameterLocation `json:"parameterLocation,omitempty"`
	Location          string                               `json:"location,omitempty"`
}

type ConstraintViolationParameterLocation string

const (
	ConstraintViolationParameterLocationPath        ConstraintViolationParameterLocation = "PATH"
	ConstraintViolationParameterLocationPayloadBody ConstraintViolationParameterLocation = "PAYLOAD_BODY"
	ConstraintViolationParameterLocationQuery       ConstraintViolationParameterLocation = "QUERY"
)

// Dashboard Types
type DashboardResponse struct {
	Dashboards []DashboardStub `json:"dashboards"`
}

type Dashboard struct {
	ID                string                `json:"id,omitempty"`
	Metadata          ConfigurationMetadata `json:"metadata,omitempty"`
	DashboardMetadata DashboardMetadata     `json:"dashboardMetadata"`
	Tiles             []Tile                `json:"tiles"`
}

type DashboardStub struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Owner string `json:"owner,omitempty"`
}

type ConfigurationMetadata struct {
	ConfigurationVersions []int  `json:"configurationVersions"`
	ClusterVersion        string `json:"clusterVersion"`
}

type DashboardMetadata struct {
	Name            string          `json:"name"`
	Shared          bool            `json:"shared,omitempty"`
	Owner           string          `json:"owner,omitempty"`
	SharingDetails  SharingInfo     `json:"sharingDetails"`
	DashboardFilter DashboardFilter `json:"dashboardFilter"`
}

type SharingInfo struct {
	LinkShared bool `json:"linkShared,omitempty"`
	Published  bool `json:"published,omitempty"`
}

type DashboardFilter struct {
	Timeframe      string                    `json:"timeframe,omitempty"`
	ManagementZone EntityShortRepresentation `json:"anagementZone,omitempty"`
}

type EntityShortRepresentation struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Tile struct {
	Name       string     `json:"name"`
	TileType   TileType   `json:"tileType"`
	Configured bool       `json:"configured,omitempty"`
	Bounds     TileBounds `json:"bounds"`
	TileFilter TileFilter `json:"tileFilter"`
}

type TileType string

const (
	TileTypeApplication             TileType = "APPLICATION"
	TileTypeApplications            TileType = "APPLICATIONS"
	TileTypeApplicationsMostActive  TileType = "APPLICATIONS_MOST_ACTIVE"
	TileTypeApplicationMethod       TileType = "APPLICATION_METHOD"
	TileTypeApplicationWorldMap     TileType = "APPLICATION_WORLDMAP"
	TileTypeAWS                     TileType = "AWS"
	TileTypeBounceRate              TileType = "BOUNCE_RATE"
	TileTypeCustomApplication       TileType = "CUSTOM_APPLICATION"
	TileTypeCustomCharting          TileType = "CUSTOM_CHARTING"
	TileTypeDatabase                TileType = "DATABASE"
	TileTypeDAtabasesOverview       TileType = "DATABASES_OVERVIEW"
	TileTypeDCRUMServices           TileType = "DCRUM_SERVICES"
	TileTypeDocker                  TileType = "DOCKER"
	TileTypeDTAQL                   TileType = "DTAQL"
	TileTypeHeader                  TileType = "HEADER"
	TileTypeHost                    TileType = "HOST"
	TileTypeHosts                   TileType = "HOSTS"
	TileTypeLogAnalytics            TileType = "LOG_ANALYTICS"
	TileTypeMarkdown                TileType = "MARKDOWN"
	TileTypeMobileApplication       TileType = "MOBILE_APPLICATION"
	TileTypeNetwork                 TileType = "NETWORK"
	TileTypeNetworkMedium           TileType = "NETWORK_MEDIUM"
	TileTypeOpenProblems            TileType = "OPEN_PROBLEMS"
	TileTypeProcessGroupsOne        TileType = "PROCESS_GROUPS_ONE"
	TileTypePureModel               TileType = "PURE_MODEL"
	TileTypeResources               TileType = "RESOURCES"
	TileTypeServices                TileType = "SERVICES"
	TileTypeServiceVersatile        TileType = "SERVICE_VERSATILE"
	TileTypeSessionMetrics          TileType = "SESSION_METRICS"
	TileTypeSyntheticHTTPMonitor    TileType = "SYNTHETIC_HTTP_MONITOR"
	TileTypeSyntheticSingleExitTest TileType = "SYNTHETIC_SINGLE_EXT_TEST"
	TileTypeSyntheticSingleWebcheck TileType = "SYNTHETIC_SINGLE_WEBCHECK"
	TileTypeSyntheticTests          TileType = "SYNTHETIC_TESTS"
	TileTypeTechnologyLandscape     TileType = "TECHNOLOGY_LANDSCAPE"
	TileTypeThirdPartyMostActive    TileType = "THIRD_PARTY_MOST_ACTIVE"
	TileTypeUEMActiveSessions       TileType = "UEM_ACTIVE_SESSIONS"
	TileTypeUEMConversionsOverall   TileType = "UEM_CONVERSIONS_OVERALL"
	TileTypeUEMConversionsPerGoal   TileType = "UEM_CONVERSIONS_PER_GOAL"
	TileTypeUEMJSErrorsOverall      TileType = "UEM_JSERRORS_OVERALL"
	TileTypeUEMKeyUserActions       TileType = "UEM_KEY_USER_ACTIONS"
	TileTypeUsers                   TileType = "USERS"
	TileTypeVirtualization          TileType = "VIRTUALIZATION"
)

type TileBounds struct {
	Top    int `json:"top,omitempty"`
	Left   int `json:"left,omitempty"`
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type TileFilter struct {
	ManagementZone EntityShortRepresentation `json:"managementZone"`
}

// Error Types
type ErrorResponse struct {
	Detail *ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	Code                 int64                 `json:"code,omitempty"`
	Message              string                `json:"message,omitempty"`
	ConstraintViolations []ConstraintViolation `json:"constraintViolations"`
}

func (e *ErrorResponse) Error() string {
	if e != nil && e.Detail != nil {
		return e.Detail.Message
	}
	return "Unknown error"
}
