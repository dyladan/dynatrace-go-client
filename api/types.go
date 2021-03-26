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
	ID                string                 `json:"id,omitempty"`
	Metadata          *ConfigurationMetadata `json:"metadata,omitempty"`
	DashboardMetadata *DashboardMetadata     `json:"dashboardMetadata,omitempty"`
	Tiles             []Tile                 `json:"tiles"`
}

type DashboardStub struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Owner string `json:"owner,omitempty"`
}

type ConfigurationMetadata struct {
	ConfigurationVersions []int  `json:"configurationVersions,omitempty"`
	ClusterVersion        string `json:"clusterVersion,omitempty"`
}

type DashboardMetadata struct {
	Name            string           `json:"name"`
	Shared          bool             `json:"shared,omitempty"`
	Owner           string           `json:"owner,omitempty"`
	SharingDetails  *SharingInfo     `json:"sharingDetails,omitempty"`
	DashboardFilter *DashboardFilter `json:"dashboardFilter"`
}

type SharingInfo struct {
	LinkShared bool `json:"linkShared,omitempty"`
	Published  bool `json:"published,omitempty"`
}

type DashboardFilter struct {
	Timeframe      string                     `json:"tismeframe,omitempty"`
	ManagementZone *EntityShortRepresentation `json:"managementZone"`
}

type EntityShortRepresentation struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Tile struct {
	Name       string      `json:"name"`
	TileType   TileType    `json:"tileType"`
	Configured bool        `json:"configured,omitempty"`
	Bounds     *TileBounds `json:"bounds"`
	TileFilter *TileFilter `json:"tileFilter"`

	// CustomChartingTile
	CustomFilterConfig *CustomFilterConfig `json:"filterConfig,omitempty"`
}

type CustomFilterConfig struct {
	Type                 string                         `json:"type"`
	CustomName           string                         `json:"customName"`
	DefaultName          string                         `json:"defaultName,omitempty"`
	ChartConfig          *CustomFilterChartConfig       `json:"chartConfig"`
	FiltersPerEntityType map[string]map[string][]string `json:"filtersPerEntityType,omitempty"`
}

type CustomFilterChartConfig struct {
	Type           CustomFilterChartConfigType      `json:"type"`
	Series         *[]CustomFilterChartSeriesConfig `json:"series"`
	ResultMetadata string                           `json:"resulMetadata,omitempty"`
}

type CustomFilterChartSeriesConfig struct {
	Metric        string                                   `json:"metric"`
	Aggregation   CustomFilterChartSeriesConfigAggregation `json:"aggregation"`
	Percentile    *int                                     `json:"percentile"`
	Type          CustomFilterChartSeriesConfigType        `json:"type"`
	EntityType    string                                   `json:"entityType"`
	Dimensions    []CustomFilterChartSeriesDimensionConfig `json:"dimensions"`
	SortAscending bool                                     `json:"sortAscending"`
	SortColumn    bool                                     `json:"sortColumn"`
}

type CustomFilterChartSeriesConfigAggregation string

const (
	AVG               CustomFilterChartSeriesConfigAggregation = "AVG"
	COUNT             CustomFilterChartSeriesConfigAggregation = "COUNT"
	DISTINCT          CustomFilterChartSeriesConfigAggregation = "DISTINCT"
	FASTEST10PERCENT  CustomFilterChartSeriesConfigAggregation = "FASTEST10PERCENT"
	MAX               CustomFilterChartSeriesConfigAggregation = "MAX"
	MEDIAN            CustomFilterChartSeriesConfigAggregation = "MEDIAN"
	MIN               CustomFilterChartSeriesConfigAggregation = "MIN"
	NONE              CustomFilterChartSeriesConfigAggregation = "NONE"
	OF_INTEREST_RATIO CustomFilterChartSeriesConfigAggregation = "OF_INTEREST_RATIO"
	OTHER_RATIO       CustomFilterChartSeriesConfigAggregation = "OTHER_RATIO"
	PERCENTILE        CustomFilterChartSeriesConfigAggregation = "PERCENTILE"
	PER_MIN           CustomFilterChartSeriesConfigAggregation = "PER_MIN"
	SLOWEST10PERCENT  CustomFilterChartSeriesConfigAggregation = "SLOWEST10PERCENT"
	SLOWEST5PERCENT   CustomFilterChartSeriesConfigAggregation = "SLOWEST5PERCENT"
	SUM               CustomFilterChartSeriesConfigAggregation = "SUM"
)

type CustomFilterChartSeriesDimensionConfig struct {
	ID               string   `json:"id"`
	Values           []string `json:"values"`
	UsedForSplitting bool     `json:"usedForSplitting,omitempty"`
	EntityDimension  bool     `json:"entityDimension,omitempty"`
}

type CustomFilterChartSeriesConfigType string

const (
	AREA CustomFilterChartSeriesConfigType = "AREA"
	BAR  CustomFilterChartSeriesConfigType = "BAR"
	LINE CustomFilterChartSeriesConfigType = "LINE"
)

type CustomFilterChartConfigType string

const (
	PIE          CustomFilterChartConfigType = "PIE"
	SINGLE_VALUE CustomFilterChartConfigType = "SINGLE_VALUE"
	TIMESERIES   CustomFilterChartConfigType = "TIMESERIES"
	TOP_LIST     CustomFilterChartConfigType = "TOP_LIST"
)

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
	Top    int `json:"top"`
	Left   int `json:"left"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type TileFilter struct {
	ManagementZone *EntityShortRepresentation `json:"managementZone"`
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

type EventType string

const (
	EventTypeAvailabilityEvent    = "AVAILABILITY_EVENT"
	EventTypeCustomAlert          = "CUSTOM_ALERT"
	EventTypeCustomAnnotation     = "CUSTOM_ANNOTATION"
	EventTypeCustomConfiguration  = "CUSTOM_CONFIGURATION"
	EventTypeCustomDeployment     = "CUSTOM_DEPLOYMENT"
	EventTypeCustomInfo           = "CUSTOM_INFO"
	EventTypeErrorEvent           = "ERROR_EVENT"
	EventTypeMarkedForTermination = "MARKED_FOR_TERMINATION"
	EventTypePerformanceEvent     = "PERFORMANCE_EVENT"
	EventTypeResourceContention   = "RESOURCE_CONTENTION"
)

type EventCreation struct {
	EventType             EventType            `json:"eventType,omitempty"`
	Start                 int64                `json:"start,omitempty"`
	End                   int64                `json:"end,omitempty"`
	TimeoutMinutes        int32                `json:"timeoutMinutes,omitempty"`
	Source                string               `json:"source,omitempty"`
	AnnotationType        string               `json:"annotationType,omitempty"`
	AnnotationDescription string               `json:"annotationDescription,omitempty"`
	AttachRules           PushEventAttachRules `json:"attachRules,omitempty"`
	Description           string               `json:"description,omitempty"`
	Title                 string               `json:"title,omitempty"`
	CustomProperties      map[string]string    `json:"customProperties,omitempty"`
	AllowDavisMerge       bool                 `json:"allowDavisMerge,omitempty"`
}

type PushEventAttachRules struct {
	EntityIds     []string       `json:"entityIds,omitempty"`
	TagMatchRules []TagMatchRule `json:"tagRule,omitempty"`
}

type TagMatchRule struct {
	MeTypes []string  `json:"meTypes,omitempty"`
	Tags    []TagInfo `json:"tags,omitempty"`
}

type TagInfo struct {
	Context string `json:"context,omitempty"`
	Key     string `json:"key,omitempty"`
}

type EventStoreResult struct {
	StoredEventIds       []int    `json:"storedEventIds,omitempty"`
	StoredIds            []string `json:"storedIds,omitempty"`
	StoredCorrelationIds []string `json:"storedCorrelationIds,omitempty"`
}

type CustomDevicePushMessage struct {
	DisplayName string            `json:"displayName,omitempty"`
	Group       string            `json:"group,omitempty"`
	IPAddresses []string          `json:"ipAddresses,omitempty"`
	ListenPorts []int             `json:"listenPorts,omitempty"`
	Favicon     string            `json:"favicon,omitempty"`
	ConfigURL   string            `json:"configUrl,omitempty"`
	Type        string            `json:"type,omitempty"`
	Properties  map[string]string `json:"properties,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	HostNames   []string          `json:"hostNames,omitempty"`
}

type CustomDevicePushResult struct {
	EntityID string `json:"entityId"`
	GroupID  string `json:"groupId"`
}

type ProblemsResponse struct {
	TotalCount int       `json:"totalCount"`
	PageSize   int       `json:"pageSize"`
	Problems   []Problem `json:"problems"`
}
type EntityID struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type AffectedEntities struct {
	EntityID EntityID `json:"entityId"`
	Name     string   `json:"name"`
}
type ImpactedEntities struct {
	EntityID EntityID `json:"entityId"`
	Name     string   `json:"name"`
}
type ProblemFilters struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Entity struct {
	EntityID EntityID `json:"entityId"`
	Name     string   `json:"name"`
}
type GroupingEntity struct {
	EntityID EntityID `json:"entityId"`
	Name     string   `json:"name"`
}
type Details struct {
	EvidenceType      string         `json:"evidenceType"`
	DisplayName       string         `json:"displayName"`
	Entity            Entity         `json:"entity"`
	GroupingEntity    GroupingEntity `json:"groupingEntity"`
	RootCauseRelevant bool           `json:"rootCauseRelevant"`
	EventID           string         `json:"eventId"`
	EventType         string         `json:"eventType"`
	StartTime         int64          `json:"startTime"`
}
type EvidenceDetails struct {
	TotalCount int       `json:"totalCount"`
	Details    []Details `json:"details"`
}

type Problem struct {
	ProblemID        string             `json:"problemId"`
	DisplayID        string             `json:"displayId"`
	Title            string             `json:"title"`
	ImpactLevel      string             `json:"impactLevel"`
	SeverityLevel    string             `json:"severityLevel"`
	Status           string             `json:"status"`
	AffectedEntities []AffectedEntities `json:"affectedEntities"`
	ImpactedEntities []ImpactedEntities `json:"impactedEntities"`
	RootCauseEntity  interface{}        `json:"rootCauseEntity"`
	ManagementZones  []interface{}      `json:"managementZones"`
	EntityTags       []interface{}      `json:"entityTags"`
	ProblemFilters   []ProblemFilters   `json:"problemFilters"`
	StartTime        int64              `json:"startTime"`
	EndTime          int                `json:"endTime"`
	EvidenceDetails  EvidenceDetails    `json:"evidenceDetails"`
}
