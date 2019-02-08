package api

// type NameDetectionRule struct {
// 	Id                    string                               `json:"id"`
// 	ApplicationIdentifier string                               `json:"applicationIdentifier"`
// 	FilterConfig          NameDetectionRuleFilterConfiguration `json:"filterConfig"`
// }

// type NameDetectionRuleFilterConfiguration struct {
// 	Pattern                string `json:"pattern"`
// 	ApplicationMatchType   string `json:"applicationMatchType"`
// 	ApplicationMatchTarget string `json:"applicationMatchTarget"`
// }

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
