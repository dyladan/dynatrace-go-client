package api

// General types
type RuleType string

const (
	AppmonServer                 RuleType = "APPMON_SERVER"
	AppmonSystemProfile          RuleType = "APPMON_SYSTEM_PROFILE"
	AWSAccount                   RuleType = "AWS_ACCOUNT"
	AWSAutoScalingGroup          RuleType = "AWS_AUTO_SCALING_GROUP"
	AWSClassicLoadBalancer       RuleType = "AWS_CLASSIC_LOAD_BALANCER"
	AWSRelationalDatabaseService RuleType = "AWS_RELATIONAL_DATABASE_SERVICE"
	BrowserMonitor               RuleType = "BROWSER_MONITOR"
	CloudFoundryFoundation       RuleType = "CLOUD_FOUNDRY_FOUNDATION"
	CustomApplication            RuleType = "CUSTOM_APPLICATION"
	CustomDevice                 RuleType = "CUSTOM_DEVICE"
	CustomDeviceGroup            RuleType = "CUSTOM_DEVICE_GROUP"
	DataCenterService            RuleType = "DATA_CENTER_SERVICE"
	EnterpriseApplication        RuleType = "ENTERPRISE_APPLICATION"
	ESXiHost                     RuleType = "ESXI_HOST"
	ExternalMonitor              RuleType = "EXTERNAL_MONITOR"
	Host                         RuleType = "HOST"
	HTTPMonitor                  RuleType = "HTTP_MONITOR"
	KubernetesCluster            RuleType = "KUBERNETES_CLUSTER"
	MobileApplication            RuleType = "MOBILE_APPLICATION"
	OpenstackAccount             RuleType = "OPENSTACK_ACCOUNT"
	ProcessGroup                 RuleType = "PROCESS_GROUP"
	Service                      RuleType = "SERVICE"
	WebApplication               RuleType = "WEB_APPLICATION"
)

type PropagationType string

const (
	ServiceToProcessGroupLike       PropagationType = "SERVICE_TO_PROCESS_GROUP_LIKE"
	ServiceToHostLike               PropagationType = "SERVICE_TO_HOST_LIKE"
	ProcessGroupToHost              PropagationType = "PROCESS_GROUP_TO_HOST"
	ProcessGroupToService           PropagationType = "PROCESS_GROUP_TO_SERVICE"
	HostToProcessGroupInstance      PropagationType = "HOST_TO_PROCESS_GROUP_INSTANCE"
	CustomDeviceGroupToCustomDevice PropagationType = "CUSTOM_DEVICE_GROUP_TO_CUSTOM_DEVICE"
)

type Attribute string

const (
	AppmonServerName                          Attribute = "APPMON_SERVER_NAME"
	AppmonSystemProfileName                   Attribute = "APPMON_SYSTEM_PROFILE_NAME"
	AWSAccountId                              Attribute = "AWS_ACCOUNT_ID"
	AWSAccountName                            Attribute = "AWS_ACCOUNT_NAME"
	AWSApplicationLoadBalancerName            Attribute = "AWS_APPLICATION_LOAD_BALANCER_NAME"
	AWSApplicationLoadBalancerTags            Attribute = "AWS_APPLICATION_LOAD_BALANCER_TAGS"
	AWSAutoScalingGroupName                   Attribute = "AWS_AUTO_SCALING_GROUP_NAME"
	AWSAutoScalingGroupTags                   Attribute = "AWS_AUTO_SCALING_GROUP_TAGS"
	AWSAvailabilityZoneName                   Attribute = "AWS_AVAILABILITY_ZONE_NAME"
	AWSClassicLoadBalancerBackendPorts        Attribute = "AWS_CLASSIC_LOAD_BALANCER_BACKEND_PORTS"
	AWSClassicLoadBalancerFrontendPorts       Attribute = "AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS"
	AWSClassicLoadBalancerName                Attribute = "AWS_CLASSIC_LOAD_BALANCER_NAME"
	AWSClassicLoadbalacerTags                 Attribute = "AWS_CLASSIC_LOAD_BALANCER_TAGS"
	AWSNetworkLoadBalancerName                Attribute = "AWS_NETWORK_LOAD_BALANCER_NAME"
	AWSNetworkLoadBalancerTags                Attribute = "AWS_NETWORK_LOAD_BALANCER_TAGS"
	AWSRelationalDatabaseServiceDBName        Attribute = "AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME"
	AWSRelationalDatabaseServiceEndpoint      Attribute = "AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT"
	AWSRelationalDatabaseServiceEngine        Attribute = "AWS_RELATIONAL_DATABASE_SERVICE_ENGINE"
	AWSRelationalDatabaseServiceInstanceClass Attribute = "AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS"
	AWSRelationalDatabaseServiceName          Attribute = "AWS_RELATIONAL_DATABASE_SERVICE_NAME"
	AWSRelationalDatabaseServicePort          Attribute = "AWS_RELATIONAL_DATABASE_SERVICE_PORT"
	AWSRelationalDatabaseServiceTags          Attribute = "AWS_RELATIONAL_DATABASE_SERVICE_TAGS"
	AzureRegionName                           Attribute = "AZURE_REGION_NAME"
	AzureScaleSetName                         Attribute = "AZURE_SCALE_SET_NAME"
	AzureVMName                               Attribute = "AZURE_VM_NAME"
	BrowserMonitorName                        Attribute = "BROWSER_MONITOR_NAME"
	BrowserMonitorTags                        Attribute = "BROWSER_MONITOR_TAGS"
	CloudFoundryFoundationName                Attribute = "CLOUD_FOUNDRY_FOUNDATION_NAME"
	CloudFoundryOrgName                       Attribute = "CLOUD_FOUNDRY_ORG_NAME"
	CustomApplicationName                     Attribute = "CUSTOM_APPLICATION_NAME"
	CustomApplicationPlatform                 Attribute = "CUSTOM_APPLICATION_PLATFORM"
	CustomApplicationTags                     Attribute = "CUSTOM_APPLICATION_TAGS"
	CustomApplicationType                     Attribute = "CUSTOM_APPLICATION_TYPE"
	CustomDeviceDetectedName                  Attribute = "CUSTOM_DEVICE_DETECTED_NAME"
	CustomDeviceDNSAddress                    Attribute = "CUSTOM_DEVICE_DNS_ADDRESS"
	CustomDeviceGroupName                     Attribute = "CUSTOM_DEVICE_GROUP_NAME"
	CustomDeviceIPAddress                     Attribute = "CUSTOM_DEVICE_IP_ADDRESS"
	CustomDeviceMetadata                      Attribute = "CUSTOM_DEVICE_METADATA"
	CustomDeviceName                          Attribute = "CUSTOM_DEVICE_NAME"
	CustomDevicePort                          Attribute = "CUSTOM_DEVICE_PORT"
	CustomDeviceTags                          Attribute = "CUSTOM_DEVICE_TAGS"
	CustomDeviceTechnology                    Attribute = "CUSTOM_DEVICE_TECHNOLOGY"
	DataCenterServiceDecoderType              Attribute = "DATA_CENTER_SERVICE_DECODER_TYPE"
	DataCenterServiceIPAddress                Attribute = "DATA_CENTER_SERVICE_IP_ADDRESS"
	DataCenterServiceMetadata                 Attribute = "DATA_CENTER_SERVICE_METADATA"
	DataCenterServiceName                     Attribute = "DATA_CENTER_SERVICE_NAME"
	DataCenterServicePort                     Attribute = "DATA_CENTER_SERVICE_PORT"
	DataCenterServiceTags                     Attribute = "DATA_CENTER_SERVICE_TAGS"
	DockerContainerName                       Attribute = "DOCKER_CONTAINER_NAME"
	DockerFullImageName                       Attribute = "DOCKER_FULL_IMAGE_NAME"
	DockerImageVersion                        Attribute = "DOCKER_IMAGE_VERSION"
	DockerStrippedImageName                   Attribute = "DOCKER_STRIPPED_IMAGE_NAME"
	EC2InstanceAMIId                          Attribute = "EC2_INSTANCE_AMI_ID"
	EC2InstanceAWSInstanceType                Attribute = "EC2_INSTANCE_AWS_INSTANCE_TYPE"
	EC2InstanceAWSSecurityGroup               Attribute = "EC2_INSTANCE_AWS_SECURITY_GROUP"
	EC2InstanceBeanstalkEnvName               Attribute = "EC2_INSTANCE_BEANSTALK_ENV_NAME"
	EC2InstanceId                             Attribute = "EC2_INSTANCE_ID"
	EC2InstanceName                           Attribute = "EC2_INSTANCE_NAME"
	EC2InstancePrivateHostName                Attribute = "EC2_INSTANCE_PRIVATE_HOST_NAME"
	EC2InstancePublicHostName                 Attribute = "EC2_INSTANCE_PUBLIC_HOST_NAME"
	EC2InstanceTags                           Attribute = "EC2_INSTANCE_TAGS"
	EnterpriseApplicationDecoderType          Attribute = "ENTERPRISE_APPLICATION_DECODER_TYPE"
	EnterpriseApplicationIPAddress            Attribute = "ENTERPRISE_APPLICATION_IP_ADDRESS"
	EnterpriseApplicationMetadata             Attribute = "ENTERPRISE_APPLICATION_METADATA"
	EnterpriseApplicationName                 Attribute = "ENTERPRISE_APPLICATION_NAME"
	EnterpriseApplicationPort                 Attribute = "ENTERPRISE_APPLICATION_PORT"
	EnterpriseApplicationTags                 Attribute = "ENTERPRISE_APPLICATION_TAGS"
	ESXiHostClusterName                       Attribute = "ESXI_HOST_CLUSTER_NAME"
	ESXiHostHardwareModel                     Attribute = "ESXI_HOST_HARDWARE_MODEL"
	ESXiHostHardwareVendor                    Attribute = "ESXI_HOST_HARDWARE_VENDOR"
	ESXiHostName                              Attribute = "ESXI_HOST_NAME"
	ESXiHostProductName                       Attribute = "ESXI_HOST_PRODUCT_NAME"
	ESXiHostProductVersion                    Attribute = "ESXI_HOST_PRODUCT_VERSION"
	ESXiHostTags                              Attribute = "ESXI_HOST_TAGS"
	ExternalMonitorEngineDiscription          Attribute = "EXTERNAL_MONITOR_ENGINE_DESCRIPTION"
	ExternalMonitorEngineName                 Attribute = "EXTERNAL_MONITOR_ENGINE_NAME"
	ExternalMonitorEngineType                 Attribute = "EXTERNAL_MONITOR_ENGINE_TYPE"
	ExternalMonitorName                       Attribute = "EXTERNAL_MONITOR_NAME"
	ExternalMonitorTags                       Attribute = "EXTERNAL_MONITOR_TAGS"
	GeolocationSiteName                       Attribute = "GEOLOCATION_SITE_NAME"
	GoogleCloudPlatformZoneName               Attribute = "GOOGLE_CLOUD_PLATFORM_ZONE_NAME"
	GoogleComputeInstanceId                   Attribute = "GOOGLE_COMPUTE_INSTANCE_ID"
	GoogleComputeInstanceMachineType          Attribute = "GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE"
	GoogleComputeInstanceName                 Attribute = "GOOGLE_COMPUTE_INSTANCE_NAME"
	GoogleComputeInstanceProject              Attribute = "GOOGLE_COMPUTE_INSTANCE_PROJECT"
	GoogleComputeInstanceProjectId            Attribute = "GOOGLE_COMPUTE_INSTANCE_PROJECT_ID"
	GoogleComputeInstancePublicIPAddresses    Attribute = "GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES"
	HostAIXLogicalCPUCount                    Attribute = "HOST_AIX_LOGICAL_CPU_COUNT"
	HostAIXSimultaneousThreads                Attribute = "HOST_AIX_SIMULTANEOUS_THREADS"
	HostAIXCPUCount                           Attribute = "HOST_AIX_VIRTUAL_CPU_COUNT"
	HostArchitecture                          Attribute = "HOST_ARCHITECTURE"
	HostAWSNameTag                            Attribute = "HOST_AWS_NAME_TAG"
	HostAzureComputeMode                      Attribute = "HOST_AZURE_COMPUTE_MODE"
	HostAzureSKU                              Attribute = "HOST_AZURE_SKU"
	HostAzureWebApplicationHostNames          Attribute = "HOST_AZURE_WEB_APPLICATION_HOST_NAMES"
	HostAzureWebApplicationSiteNames          Attribute = "HOST_AZURE_WEB_APPLICATION_SITE_NAMES"
	HostBitness                               Attribute = "HOST_BITNESS"
	HostBOSHAvailabilityZone                  Attribute = "HOST_BOSH_AVAILABILITY_ZONE"
	HostBOSHDeploymentId                      Attribute = "HOST_BOSH_DEPLOYMENT_ID"
	HostBOSHInstanceId                        Attribute = "HOST_BOSH_INSTANCE_ID"
	HostBOSHInstanceName                      Attribute = "HOST_BOSH_INSTANCE_NAME"
	HostBOSHName                              Attribute = "HOST_BOSH_NAME"
	HostBOSHStemcellVersion                   Attribute = "HOST_BOSH_STEMCELL_VERSION"
	HostCloudType                             Attribute = "HOST_CLOUD_TYPE"
	HostCustomMetadata                        Attribute = "HOST_CUSTOM_METADATA"
	HostDetectedName                          Attribute = "HOST_DETECTED_NAME"
	HostGroupId                               Attribute = "HOST_GROUP_ID"
	HostGroupName                             Attribute = "HOST_GROUP_NAME"
	HostHypervisorType                        Attribute = "HOST_HYPERVISOR_TYPE"
	HostIPAddress                             Attribute = "HOST_IP_ADDRESS"
	HostKubernetesLabels                      Attribute = "HOST_KUBERNETES_LABELS"
	HostName                                  Attribute = "HOST_NAME"
	HostOneAgentCustomHostName                Attribute = "HOST_ONEAGENT_CUSTOM_HOST_NAME"
	HostOSType                                Attribute = "HOST_OS_TYPE"
	HostOSVersion                             Attribute = "HOST_OS_VERSION"
	HostPaaSMemoryLimit                       Attribute = "HOST_PAAS_MEMORY_LIMIT"
	HostPaaSType                              Attribute = "HOST_PAAS_TYPE"
	HostTags                                  Attribute = "HOST_TAGS"
	HostTechnology                            Attribute = "HOST_TECHNOLOGY"
	HTTPMonitorName                           Attribute = "HTTP_MONITOR_NAME"
	HTTPMonitorTags                           Attribute = "HTTP_MONITOR_TAGS"
	KubernetesClusterName                     Attribute = "KUBERNETES_CLUSTER_NAME"
	KubernetesNodeName                        Attribute = "KUBERNETES_NODE_NAME"
	MobileApplicationName                     Attribute = "MOBILE_APPLICATION_NAME"
	MobileApplicationPlatform                 Attribute = "MOBILE_APPLICATION_PLATFORM"
	MobileApplicationTags                     Attribute = "MOBILE_APPLICATION_TAGS"
	NameOfComputeNode                         Attribute = "NAME_OF_COMPUTE_NODE"
	OpenStackAccountName                      Attribute = "OPENSTACK_ACCOUNT_NAME"
	OpenStackAccountProjectName               Attribute = "OPENSTACK_ACCOUNT_PROJECT_NAME"
	OpenStackAvailabilityZoneName             Attribute = "OPENSTACK_AVAILABILITY_ZONE_NAME"
	OpenStackProjectName                      Attribute = "OPENSTACK_PROJECT_NAME"
	OpenStackRegionName                       Attribute = "OPENSTACK_REGION_NAME"
	OpenStackVMInstanceType                   Attribute = "OPENSTACK_VM_INSTANCE_TYPE"
	OpenStackVMName                           Attribute = "OPENSTACK_VM_NAME"
	OpenStackVMSecurityGroup                  Attribute = "OPENSTACK_VM_SECURITY_GROUP"
	ProcessGroupAzureHostName                 Attribute = "PROCESS_GROUP_AZURE_HOST_NAME"
	ProcessGroupAzureSiteName                 Attribute = "PROCESS_GROUP_AZURE_SITE_NAME"
	ProcessGroupCustomMetadata                Attribute = "PROCESS_GROUP_CUSTOM_METADATA"
	ProcessGroupDetectedName                  Attribute = "PROCESS_GROUP_DETECTED_NAME"
	ProcessGroupId                            Attribute = "PROCESS_GROUP_ID"
	ProcessGroupInstanceCustomMetadata        Attribute = "PROCESS_GROUP_INSTANCE_CUSTOM_METADATA"
	ProcessGroupListenPort                    Attribute = "PROCESS_GROUP_LISTEN_PORT"
	ProcessGroupName                          Attribute = "PROCESS_GROUP_NAME"
	ProcessGroupPredefinedMetadata            Attribute = "PROCESS_GROUP_PREDEFINED_METADATA"
	ProcessGroupTags                          Attribute = "PROCESS_GROUP_TAGS"
	ProcessGroupTechnology                    Attribute = "PROCESS_GROUP_TECHNOLOGY"
	ProcessGroupTechnologyEdition             Attribute = "PROCESS_GROUP_TECHNOLOGY_EDITION"
	ProcessGroupTechnologyVersion             Attribute = "PROCESS_GROUP_TECHNOLOGY_VERSION"
	ServiceAkkaActorSystem                    Attribute = "SERVICE_AKKA_ACTOR_SYSTEM"
	ServiceCTGServiceName                     Attribute = "SERVICE_CTG_SERVICE_NAME"
	ServiceDatabaseHostName                   Attribute = "SERVICE_DATABASE_HOST_NAME"
	ServiceDatabaseName                       Attribute = "SERVICE_DATABASE_NAME"
	ServiceDatabaseTopology                   Attribute = "SERVICE_DATABASE_TOPOLOGY"
	ServiceDatabaseVendor                     Attribute = "SERVICE_DATABASE_VENDOR"
	ServiceDetectedName                       Attribute = "SERVICE_DETECTED_NAME"
	ServiceIBMCTGGatewayURL                   Attribute = "SERVICE_IBM_CTG_GATEWAY_URL"
	ServiceIIBApplicationName                 Attribute = "SERVICE_IIB_APPLICATION_NAME"
	ServiceMessagingListenerClassName         Attribute = "SERVICE_MESSAGING_LISTENER_CLASS_NAME"
	ServiceName                               Attribute = "SERVICE_NAME"
	ServicePort                               Attribute = "SERVICE_PORT"
	ServicePublicDomainName                   Attribute = "SERVICE_PUBLIC_DOMAIN_NAME"
	ServiceRemoteEndpoint                     Attribute = "SERVICE_REMOTE_ENDPOINT"
	ServiceRemoteServiceName                  Attribute = "SERVICE_REMOTE_SERVICE_NAME"
	ServiceTags                               Attribute = "SERVICE_TAGS"
	ServiceTechnology                         Attribute = "SERVICE_TECHNOLOGY"
	ServiceTechnologyEdition                  Attribute = "SERVICE_TECHNOLOGY_EDITION"
	ServiceTechnologyVersion                  Attribute = "SERVICE_TECHNOLOGY_VERSION"
	ServiceTopology                           Attribute = "SERVICE_TOPOLOGY"
	ServiceType                               Attribute = "SERVICE_TYPE"
	ServiceWebApplicationId                   Attribute = "SERVICE_WEB_APPLICATION_ID"
	ServiceWebContextRoot                     Attribute = "SERVICE_WEB_CONTEXT_ROOT"
	ServiceWebServerEndpoint                  Attribute = "SERVICE_WEB_SERVER_ENDPOINT"
	ServiceWebServerName                      Attribute = "SERVICE_WEB_SERVER_NAME"
	ServiceWebServiceName                     Attribute = "SERVICE_WEB_SERVICE_NAME"
	ServiceWebServiceNamespace                Attribute = "SERVICE_WEB_SERVICE_NAMESPACE"
	VMWareDatacenterName                      Attribute = "VMWARE_DATACENTER_NAME"
	VMWareVMName                              Attribute = "VMWARE_VM_NAME"
	WebApplicationName                        Attribute = "WEB_APPLICATION_NAME"
	WebApplicationNamePattern                 Attribute = "WEB_APPLICATION_NAME_PATTERN"
	WebApplicationTags                        Attribute = "WEB_APPLICATION_TAGS"
	WebApplicationType                        Attribute = "WEB_APPLICATION_TYPE"
)

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
	Type             AutoTagRuleType        `json:"type"`
	Enabled          bool                   `json:"enabled"`
	ValueFormat      string                 `json:"valueFormat"`
	PropagationTypes []PropagationType      `json:"propagationTypes"`
	Conditions       []AutoTagRuleCondition `json:"conditions"`
}

type AutoTagRuleType string

const (
	AutoTagRuleTypeCustomDevice AutoTagRuleType = "CUSTOM_DEVICE"
	AutoTagRuleTypeHost                         = "HOST"
	AutoTagRuleTypeProcessGroup                 = "PROCESS_GROUP"
	AutoTagRuleTypeService                      = "SERVICE"
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

// Management zone types

type ManagementZoneRule struct {
	Type             RuleType          `json:"type"`
	Enabled          bool              `json:"enabled"`
	PropagationTypes []PropagationType `json:"propagationTypes"`
}

type ManagementZoneMetadata struct {
	ConfigurationVersions []int64 `json:"convigurationVersions"`
	ClusterVersion        string  `json:"clusterVersion"`
}

type ManagementZoneType struct {
	Metadata ManagementZoneMetadata `json:"metadata"`
	Id       string                 `json:"id"`
	Name     string                 `json:"name"`
	Rules    []ManagementZoneRule   `json:"rules"`
}
