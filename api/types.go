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
	CustomDevice AutoTagRuleType = "CUSTOM_DEVICE"
	Host                         = "HOST"
	ProcessGroup                 = "PROCESS_GROUP"
	Service                      = "SERVICE"
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
