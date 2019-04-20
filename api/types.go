package api

// General types
type RuleType string

const (
	AppmonServer                 RuleType = "APPMON_SERVER"
	AppmonSystemProfile                   = "APPMON_SYSTEM_PROFILE"
	AWSAccount                            = "AWS_ACCOUNT"
	AWSAutoScalingGroup                   = "AWS_AUTO_SCALING_GROUP"
	AWSClassicLoadBalancer                = "AWS_CLASSIC_LOAD_BALANCER"
	AWSRelationalDatabaseService          = "AWS_RELATIONAL_DATABASE_SERVICE"
	BrowserMonitor                        = "BROWSER_MONITOR"
	CloudFoundryFoundation                = "CLOUD_FOUNDRY_FOUNDATION"
	CustomApplication                     = "CUSTOM_APPLICATION"
	CustomDevice                          = "CUSTOM_DEVICE"
	CustomDeviceGroup                     = "CUSTOM_DEVICE_GROUP"
	DataCenterService                     = "DATA_CENTER_SERVICE"
	EnterpriseApplication                 = "ENTERPRISE_APPLICATION"
	ESXiHost                              = "ESXI_HOST"
	ExternalMonitor                       = "EXTERNAL_MONITOR"
	Host                                  = "HOST"
	HTTPMonitor                           = "HTTP_MONITOR"
	KubernetesCluster                     = "KUBERNETES_CLUSTER"
	MobileApplication                     = "MOBILE_APPLICATION"
	OpenstackAccount                      = "OPENSTACK_ACCOUNT"
	ProcessGroup                          = "PROCESS_GROUP"
	Service                               = "SERVICE"
	WebApplication                        = "WEB_APPLICATION"
)

type Operator string

const (
	Equals Operator = "EQUALS"
	Exists          = "EXISTS"
)

type Comparison string

const (
	BeginsWith         Comparison = "BEGINS_WITH"
	Contains                      = "CONTAINS"
	EndsWith                      = "ENDS_WITH"
	EqualsCompairson              = "EQUALS"
	ExistsComparison              = "EXISTS"
	RegexMatches                  = "REGEX_MATCHES"
	GreaterThan                   = "GREATER_THAN"
	GreaterThanOrEqual            = "GREATER_THAN_OR_EQUAL"
	LowerThan                     = "LOWER_THAN"
	LowerThanOrEqual              = "LOWER_THAN_OR_EQUAL"
)

type PropagationType string

const (
	ServiceToProcessGroupLike       PropagationType = "SERVICE_TO_PROCESS_GROUP_LIKE"
	ServiceToHostLike                               = "SERVICE_TO_HOST_LIKE"
	ProcessGroupToHost                              = "PROCESS_GROUP_TO_HOST"
	ProcessGroupToService                           = "PROCESS_GROUP_TO_SERVICE"
	HostToProcessGroupInstance                      = "HOST_TO_PROCESS_GROUP_INSTANCE"
	CustomDeviceGroupToCustomDevice                 = "CUSTOM_DEVICE_GROUP_TO_CUSTOM_DEVICE"
)

type Attribute string

const (
	AppmonServerName                          Attribute = "APPMON_SERVER_NAME"
	AppmonSystemProfileName                             = "APPMON_SYSTEM_PROFILE_NAME"
	AWSAccountId                                        = "AWS_ACCOUNT_ID"
	AWSAccountName                                      = "AWS_ACCOUNT_NAME"
	AWSApplicationLoadBalancerName                      = "AWS_APPLICATION_LOAD_BALANCER_NAME"
	AWSApplicationLoadBalancerTags                      = "AWS_APPLICATION_LOAD_BALANCER_TAGS"
	AWSAutoScalingGroupName                             = "AWS_AUTO_SCALING_GROUP_NAME"
	AWSAutoScalingGroupTags                             = "AWS_AUTO_SCALING_GROUP_TAGS"
	AWSAvailabilityZoneName                             = "AWS_AVAILABILITY_ZONE_NAME"
	AWSClassicLoadBalancerBackendPorts                  = "AWS_CLASSIC_LOAD_BALANCER_BACKEND_PORTS"
	AWSClassicLoadBalancerFrontendPorts                 = "AWS_CLASSIC_LOAD_BALANCER_FRONTEND_PORTS"
	AWSClassicLoadBalancerName                          = "AWS_CLASSIC_LOAD_BALANCER_NAME"
	AWSClassicLoadbalacerTags                           = "AWS_CLASSIC_LOAD_BALANCER_TAGS"
	AWSNetworkLoadBalancerName                          = "AWS_NETWORK_LOAD_BALANCER_NAME"
	AWSNetworkLoadBalancerTags                          = "AWS_NETWORK_LOAD_BALANCER_TAGS"
	AWSRelationalDatabaseServiceDBName                  = "AWS_RELATIONAL_DATABASE_SERVICE_DB_NAME"
	AWSRelationalDatabaseServiceEndpoint                = "AWS_RELATIONAL_DATABASE_SERVICE_ENDPOINT"
	AWSRelationalDatabaseServiceEngine                  = "AWS_RELATIONAL_DATABASE_SERVICE_ENGINE"
	AWSRelationalDatabaseServiceInstanceClass           = "AWS_RELATIONAL_DATABASE_SERVICE_INSTANCE_CLASS"
	AWSRelationalDatabaseServiceName                    = "AWS_RELATIONAL_DATABASE_SERVICE_NAME"
	AWSRelationalDatabaseServicePort                    = "AWS_RELATIONAL_DATABASE_SERVICE_PORT"
	AWSRelationalDatabaseServiceTags                    = "AWS_RELATIONAL_DATABASE_SERVICE_TAGS"
	AzureRegionName                                     = "AZURE_REGION_NAME"
	AzureScaleSetName                                   = "AZURE_SCALE_SET_NAME"
	AzureVMName                                         = "AZURE_VM_NAME"
	BrowserMonitorName                                  = "BROWSER_MONITOR_NAME"
	BrowserMonitorTags                                  = "BROWSER_MONITOR_TAGS"
	CloudFoundryFoundationName                          = "CLOUD_FOUNDRY_FOUNDATION_NAME"
	CloudFoundryOrgName                                 = "CLOUD_FOUNDRY_ORG_NAME"
	CustomApplicationName                               = "CUSTOM_APPLICATION_NAME"
	CustomApplicationPlatform                           = "CUSTOM_APPLICATION_PLATFORM"
	CustomApplicationTags                               = "CUSTOM_APPLICATION_TAGS"
	CustomApplicationType                               = "CUSTOM_APPLICATION_TYPE"
	CustomDeviceDetectedName                            = "CUSTOM_DEVICE_DETECTED_NAME"
	CustomDeviceDNSAddress                              = "CUSTOM_DEVICE_DNS_ADDRESS"
	CustomDeviceGroupName                               = "CUSTOM_DEVICE_GROUP_NAME"
	CustomDeviceIPAddress                               = "CUSTOM_DEVICE_IP_ADDRESS"
	CustomDeviceMetadata                                = "CUSTOM_DEVICE_METADATA"
	CustomDeviceName                                    = "CUSTOM_DEVICE_NAME"
	CustomDevicePort                                    = "CUSTOM_DEVICE_PORT"
	CustomDeviceTags                                    = "CUSTOM_DEVICE_TAGS"
	CustomDeviceTechnology                              = "CUSTOM_DEVICE_TECHNOLOGY"
	DataCenterServiceDecoderType                        = "DATA_CENTER_SERVICE_DECODER_TYPE"
	DataCenterServiceIPAddress                          = "DATA_CENTER_SERVICE_IP_ADDRESS"
	DataCenterServiceMetadata                           = "DATA_CENTER_SERVICE_METADATA"
	DataCenterServiceName                               = "DATA_CENTER_SERVICE_NAME"
	DataCenterServicePort                               = "DATA_CENTER_SERVICE_PORT"
	DataCenterServiceTags                               = "DATA_CENTER_SERVICE_TAGS"
	DockerContainerName                                 = "DOCKER_CONTAINER_NAME"
	DockerFullImageName                                 = "DOCKER_FULL_IMAGE_NAME"
	DockerImageVersion                                  = "DOCKER_IMAGE_VERSION"
	DockerStrippedImageName                             = "DOCKER_STRIPPED_IMAGE_NAME"
	EC2InstanceAMIId                                    = "EC2_INSTANCE_AMI_ID"
	EC2InstanceAWSInstanceType                          = "EC2_INSTANCE_AWS_INSTANCE_TYPE"
	EC2InstanceAWSSecurityGroup                         = "EC2_INSTANCE_AWS_SECURITY_GROUP"
	EC2InstanceBeanstalkEnvName                         = "EC2_INSTANCE_BEANSTALK_ENV_NAME"
	EC2InstanceId                                       = "EC2_INSTANCE_ID"
	EC2InstanceName                                     = "EC2_INSTANCE_NAME"
	EC2InstancePrivateHostName                          = "EC2_INSTANCE_PRIVATE_HOST_NAME"
	EC2InstancePublicHostName                           = "EC2_INSTANCE_PUBLIC_HOST_NAME"
	EC2InstanceTags                                     = "EC2_INSTANCE_TAGS"
	EnterpriseApplicationDecoderType                    = "ENTERPRISE_APPLICATION_DECODER_TYPE"
	EnterpriseApplicationIPAddress                      = "ENTERPRISE_APPLICATION_IP_ADDRESS"
	EnterpriseApplicationMetadata                       = "ENTERPRISE_APPLICATION_METADATA"
	EnterpriseApplicationName                           = "ENTERPRISE_APPLICATION_NAME"
	EnterpriseApplicationPort                           = "ENTERPRISE_APPLICATION_PORT"
	EnterpriseApplicationTags                           = "ENTERPRISE_APPLICATION_TAGS"
	ESXiHostClusterName                                 = "ESXI_HOST_CLUSTER_NAME"
	ESXiHostHardwareModel                               = "ESXI_HOST_HARDWARE_MODEL"
	ESXiHostHardwareVendor                              = "ESXI_HOST_HARDWARE_VENDOR"
	ESXiHostName                                        = "ESXI_HOST_NAME"
	ESXiHostProductName                                 = "ESXI_HOST_PRODUCT_NAME"
	ESXiHostProductVersion                              = "ESXI_HOST_PRODUCT_VERSION"
	ESXiHostTags                                        = "ESXI_HOST_TAGS"
	ExternalMonitorEngineDiscription                    = "EXTERNAL_MONITOR_ENGINE_DESCRIPTION"
	ExternalMonitorEngineName                           = "EXTERNAL_MONITOR_ENGINE_NAME"
	ExternalMonitorEngineType                           = "EXTERNAL_MONITOR_ENGINE_TYPE"
	ExternalMonitorName                                 = "EXTERNAL_MONITOR_NAME"
	ExternalMonitorTags                                 = "EXTERNAL_MONITOR_TAGS"
	GeolocationSiteName                                 = "GEOLOCATION_SITE_NAME"
	GoogleCloudPlatformZoneName                         = "GOOGLE_CLOUD_PLATFORM_ZONE_NAME"
	GoogleComputeInstanceId                             = "GOOGLE_COMPUTE_INSTANCE_ID"
	GoogleComputeInstanceMachineType                    = "GOOGLE_COMPUTE_INSTANCE_MACHINE_TYPE"
	GoogleComputeInstanceName                           = "GOOGLE_COMPUTE_INSTANCE_NAME"
	GoogleComputeInstanceProject                        = "GOOGLE_COMPUTE_INSTANCE_PROJECT"
	GoogleComputeInstanceProjectId                      = "GOOGLE_COMPUTE_INSTANCE_PROJECT_ID"
	GoogleComputeInstancePublicIPAddresses              = "GOOGLE_COMPUTE_INSTANCE_PUBLIC_IP_ADDRESSES"
	HostAIXLogicalCPUCount                              = "HOST_AIX_LOGICAL_CPU_COUNT"
	HostAIXSimultaneousThreads                          = "HOST_AIX_SIMULTANEOUS_THREADS"
	HostAIXCPUCount                                     = "HOST_AIX_VIRTUAL_CPU_COUNT"
	HostArchitecture                                    = "HOST_ARCHITECTURE"
	HostAWSNameTag                                      = "HOST_AWS_NAME_TAG"
	HostAzureComputeMode                                = "HOST_AZURE_COMPUTE_MODE"
	HostAzureSKU                                        = "HOST_AZURE_SKU"
	HostAzureWebApplicationHostNames                    = "HOST_AZURE_WEB_APPLICATION_HOST_NAMES"
	HostAzureWebApplicationSiteNames                    = "HOST_AZURE_WEB_APPLICATION_SITE_NAMES"
	HostBitness                                         = "HOST_BITNESS"
	HostBOSHAvailabilityZone                            = "HOST_BOSH_AVAILABILITY_ZONE"
	HostBOSHDeploymentId                                = "HOST_BOSH_DEPLOYMENT_ID"
	HostBOSHInstanceId                                  = "HOST_BOSH_INSTANCE_ID"
	HostBOSHInstanceName                                = "HOST_BOSH_INSTANCE_NAME"
	HostBOSHName                                        = "HOST_BOSH_NAME"
	HostBOSHStemcellVersion                             = "HOST_BOSH_STEMCELL_VERSION"
	HostCloudType                                       = "HOST_CLOUD_TYPE"
	HostCustomMetadata                                  = "HOST_CUSTOM_METADATA"
	HostDetectedName                                    = "HOST_DETECTED_NAME"
	HostGroupId                                         = "HOST_GROUP_ID"
	HostGroupName                                       = "HOST_GROUP_NAME"
	HostHypervisorType                                  = "HOST_HYPERVISOR_TYPE"
	HostIPAddress                                       = "HOST_IP_ADDRESS"
	HostKubernetesLabels                                = "HOST_KUBERNETES_LABELS"
	HostName                                            = "HOST_NAME"
	HostOneAgentCustomHostName                          = "HOST_ONEAGENT_CUSTOM_HOST_NAME"
	HostOSType                                          = "HOST_OS_TYPE"
	HostOSVersion                                       = "HOST_OS_VERSION"
	HostPaaSMemoryLimit                                 = "HOST_PAAS_MEMORY_LIMIT"
	HostPaaSType                                        = "HOST_PAAS_TYPE"
	HostTags                                            = "HOST_TAGS"
	HostTechnology                                      = "HOST_TECHNOLOGY"
	HTTPMonitorName                                     = "HTTP_MONITOR_NAME"
	HTTPMonitorTags                                     = "HTTP_MONITOR_TAGS"
	KubernetesClusterName                               = "KUBERNETES_CLUSTER_NAME"
	KubernetesNodeName                                  = "KUBERNETES_NODE_NAME"
	MobileApplicationName                               = "MOBILE_APPLICATION_NAME"
	MobileApplicationPlatform                           = "MOBILE_APPLICATION_PLATFORM"
	MobileApplicationTags                               = "MOBILE_APPLICATION_TAGS"
	NameOfComputeNode                                   = "NAME_OF_COMPUTE_NODE"
	OpenStackAccountName                                = "OPENSTACK_ACCOUNT_NAME"
	OpenStackAccountProjectName                         = "OPENSTACK_ACCOUNT_PROJECT_NAME"
	OpenStackAvailabilityZoneName                       = "OPENSTACK_AVAILABILITY_ZONE_NAME"
	OpenStackProjectName                                = "OPENSTACK_PROJECT_NAME"
	OpenStackRegionName                                 = "OPENSTACK_REGION_NAME"
	OpenStackVMInstanceType                             = "OPENSTACK_VM_INSTANCE_TYPE"
	OpenStackVMName                                     = "OPENSTACK_VM_NAME"
	OpenStackVMSecurityGroup                            = "OPENSTACK_VM_SECURITY_GROUP"
	ProcessGroupAzureHostName                           = "PROCESS_GROUP_AZURE_HOST_NAME"
	ProcessGroupAzureSiteName                           = "PROCESS_GROUP_AZURE_SITE_NAME"
	ProcessGroupCustomMetadata                          = "PROCESS_GROUP_CUSTOM_METADATA"
	ProcessGroupDetectedName                            = "PROCESS_GROUP_DETECTED_NAME"
	ProcessGroupId                                      = "PROCESS_GROUP_ID"
	ProcessGroupInstanceCustomMetadata                  = "PROCESS_GROUP_INSTANCE_CUSTOM_METADATA"
	ProcessGroupListenPort                              = "PROCESS_GROUP_LISTEN_PORT"
	ProcessGroupName                                    = "PROCESS_GROUP_NAME"
	ProcessGroupPredefinedMetadata                      = "PROCESS_GROUP_PREDEFINED_METADATA"
	ProcessGroupTags                                    = "PROCESS_GROUP_TAGS"
	ProcessGroupTechnology                              = "PROCESS_GROUP_TECHNOLOGY"
	ProcessGroupTechnologyEdition                       = "PROCESS_GROUP_TECHNOLOGY_EDITION"
	ProcessGroupTechnologyVersion                       = "PROCESS_GROUP_TECHNOLOGY_VERSION"
	ServiceAkkaActorSystem                              = "SERVICE_AKKA_ACTOR_SYSTEM"
	ServiceCTGServiceName                               = "SERVICE_CTG_SERVICE_NAME"
	ServiceDatabaseHostName                             = "SERVICE_DATABASE_HOST_NAME"
	ServiceDatabaseName                                 = "SERVICE_DATABASE_NAME"
	ServiceDatabaseTopology                             = "SERVICE_DATABASE_TOPOLOGY"
	ServiceDatabaseVendor                               = "SERVICE_DATABASE_VENDOR"
	ServiceDetectedName                                 = "SERVICE_DETECTED_NAME"
	ServiceIBMCTGGatewayURL                             = "SERVICE_IBM_CTG_GATEWAY_URL"
	ServiceIIBApplicationName                           = "SERVICE_IIB_APPLICATION_NAME"
	ServiceMessagingListenerClassName                   = "SERVICE_MESSAGING_LISTENER_CLASS_NAME"
	ServiceName                                         = "SERVICE_NAME"
	ServicePort                                         = "SERVICE_PORT"
	ServicePublicDomainName                             = "SERVICE_PUBLIC_DOMAIN_NAME"
	ServiceRemoteEndpoint                               = "SERVICE_REMOTE_ENDPOINT"
	ServiceRemoteServiceName                            = "SERVICE_REMOTE_SERVICE_NAME"
	ServiceTags                                         = "SERVICE_TAGS"
	ServiceTechnology                                   = "SERVICE_TECHNOLOGY"
	ServiceTechnologyEdition                            = "SERVICE_TECHNOLOGY_EDITION"
	ServiceTechnologyVersion                            = "SERVICE_TECHNOLOGY_VERSION"
	ServiceTopology                                     = "SERVICE_TOPOLOGY"
	ServiceType                                         = "SERVICE_TYPE"
	ServiceWebApplicationId                             = "SERVICE_WEB_APPLICATION_ID"
	ServiceWebContextRoot                               = "SERVICE_WEB_CONTEXT_ROOT"
	ServiceWebServerEndpoint                            = "SERVICE_WEB_SERVER_ENDPOINT"
	ServiceWebServerName                                = "SERVICE_WEB_SERVER_NAME"
	ServiceWebServiceName                               = "SERVICE_WEB_SERVICE_NAME"
	ServiceWebServiceNamespace                          = "SERVICE_WEB_SERVICE_NAMESPACE"
	VMWareDatacenterName                                = "VMWARE_DATACENTER_NAME"
	VMWareVMName                                        = "VMWARE_VM_NAME"
	WebApplicationName                                  = "WEB_APPLICATION_NAME"
	WebApplicationNamePattern                           = "WEB_APPLICATION_NAME_PATTERN"
	WebApplicationTags                                  = "WEB_APPLICATION_TAGS"
	WebApplicationType                                  = "WEB_APPLICATION_TYPE"
)

const (
	AkkaActorClassName        Attribute = "AKKA_ACTOR_CLASS_NAME"
	AkkaActorMessageType                = "AKKA_ACTOR_MESSAGE_TYPE"
	AkkaActorPath                       = "AKKA_ACTOR_PATH"
	CICSProgramName                     = "CICS_PROGRAM_NAME"
	CICSSystemId                        = "CICS_SYSTEM_ID"
	CICSTaskId                          = "CICS_TASK_ID"
	CICSTransactionId                   = "CICS_TRANSACTION_ID"
	CTGProgram                          = "CTG_PROGRAM"
	CTGTransactionId                    = "CTG_TRANSACTION_ID"
	CustomServiceClass                  = "CUSTOMSERVICE_CLASS"
	CustomServiceMethod                 = "CUSTOMSERVICE_METHOD"
	HTTPRequestMethod                   = "HTTP_REQUEST_METHOD"
	HTTPStatus                          = "HTTP_STATUS"
	IIBApplicationName                  = "IIB_APPLICATION_NAME"
	IIBInputType                        = "IIB_INPUT_TYPE"
	IIBLibraryName                      = "IIB_LIBRARY_NAME"
	IIBMessageFlowName                  = "IIB_MESSAGE_FLOW_NAME"
	IMSProgramName                      = "IMS_PROGRAM_NAME"
	IMSTransactionId                    = "IMS_TRANSACTION_ID"
	MessagingDestinationType            = "MESSAGING_DESTINATION_TYPE"
	MessagingIsTemporaryQueue           = "MESSAGING_IS_TEMPORARY_QUEUE"
	MessagingQueueName                  = "MESSAGING_QUEUE_NAME"
	MessagingQueueVendor                = "MESSAGING_QUEUE_VENDOR"
	RemoteEndpoint                      = "REMOTE_ENDPOINT"
	RemoteMethod                        = "REMOTE_METHOD"
	RequestName                         = "REQUEST_NAME"
	RMIClass                            = "RMI_CLASS"
	RMIMethod                           = "RMI_METHOD"
	ServiceRequestAttribute             = "SERVICE_REQUEST_ATTRIBUTE"
	WebrequestQuery                     = "WEBREQUEST_QUERY"
	WebrequestURL                       = "WEBREQUEST_URL"
	WebrequestURLHost                   = "WEBREQUEST_URL_HOST"
	WebrequestURLPath                   = "WEBREQUEST_URL_PATH"
	WebrequestURLPort                   = "WEBREQUEST_URL_PORT"
	WebserviceEndpoint                  = "WEBSERVICE_ENDPOINT"
	WebserviceMethod                    = "WEBSERVICE_METHOD"
	ZOSCallType                         = "ZOS_CALL_TYPE"
)

const (
	AmazonECRImageAccountId                  Attribute = "AMAZON_ECR_IMAGE_ACCOUNT_ID"
	AmazonECRImageRegion                               = "AMAZON_ECR_IMAGE_REGION"
	AmazonLambdaFunctionName                           = "AMAZON_LAMBDA_FUNCTION_NAME"
	AmazonRegion                                       = "AMAZON_REGION"
	ApacheConfigPath                                   = "APACHE_CONFIG_PATH"
	ApacheSparkMasterIPAddress                         = "APACHE_SPARK_MASTER_IP_ADDRESS"
	ASPDotNetCoreApplicationPath                       = "ASP_DOT_NET_CORE_APPLICATION_PATH"
	AWSECSCluster                                      = "AWS_ECS_CLUSTER"
	AWSECSContainerName                                = "AWS_ECS_CONTAINERNAME"
	AWSECSFamily                                       = "AWS_ECS_FAMILY"
	AWSECSRevision                                     = "AWS_ECS_REVISION"
	CassandraClusterName                               = "CASSANDRA_CLUSTER_NAME"
	CatalinaBase                                       = "CATALINA_BASE"
	CatalinaHome                                       = "CATALINA_HOME"
	CloudFoundryAppId                                  = "CLOUD_FOUNDRY_APP_ID"
	CloudFoundryAppName                                = "CLOUD_FOUNDRY_APP_NAME"
	CloudFoundryInstanceIndex                          = "CLOUD_FOUNDRY_INSTANCE_INDEX"
	CloudFoundrySpaceId                                = "CLOUD_FOUNDRY_SPACE_ID"
	CloudFoundrySpaceName                              = "CLOUD_FOUNDRY_SPACE_NAME"
	ColdfusionJVMConfigFile                            = "COLDFUSION_JVM_CONFIG_FILE"
	ColdfusionServiceName                              = "COLDFUSION_SERVICE_NAME"
	CommandLineArgs                                    = "COMMAND_LINE_ARGS"
	DotNetCommand                                      = "DOTNET_COMMAND"
	DotNetCommandPath                                  = "DOTNET_COMMAND_PATH"
	DynatraceClusterId                                 = "DYNATRACE_CLUSTER_ID"
	DynatraceNodeId                                    = "DYNATRACE_NODE_ID"
	ElasticsearchClustName                             = "ELASTICSEARCH_CLUSTER_NAME"
	ElasticsearchNodeName                              = "ELASTICSEARCH_NODE_NAME"
	EquinoxConfigPath                                  = "EQUINOX_CONFIG_PATH"
	EXEName                                            = "EXE_NAME"
	EXEPath                                            = "EXE_PATH"
	GlassFishDomainName                                = "GLASS_FISH_DOMAIN_NAME"
	GlassFishInstanceName                              = "GLASS_FISH_INSTANCE_NAME"
	GoogleAppEngineInstance                            = "GOOGLE_APP_ENGINE_INSTANCE"
	GoogleAppEngineService                             = "GOOGLE_APP_ENGINE_SERVICE"
	GoogleCloudProject                                 = "GOOGLE_CLOUD_PROJECT"
	HybrisBinDirectory                                 = "HYBRIS_BIN_DIRECTORY"
	HybrisConfigDirectory                              = "HYBRIS_CONFIG_DIRECTORY"
	HybrisDataDirectory                                = "HYBRIS_DATA_DIRECTORY"
	IBMCICSRegion                                      = "IBM_CICS_REGION"
	IBMCTGName                                         = "IBM_CTG_NAME"
	IBMIMSConnectRegion                                = "IBM_IMS_CONNECT_REGION"
	IBMIMSControlRegion                                = "IBM_IMS_CONTROL_REGION"
	IBMIMSMessageProcessingRegion                      = "IBM_IMS_MESSAGE_PROCESSING_REGION"
	IBMIMSSoapGWName                                   = "IBM_IMS_SOAP_GW_NAME"
	IBMIntegrationNodeName                             = "IBM_INTEGRATION_NODE_NAME"
	IBMIntegrationServerName                           = "IBM_INTEGRATION_SERVER_NAME"
	IISAppPool                                         = "IIS_APP_POOL"
	IISRoleName                                        = "IIS_ROLE_NAME"
	JavaJarFile                                        = "JAVA_JAR_FILE"
	JavaJarPath                                        = "JAVA_JAR_PATH"
	JavaMainClass                                      = "JAVA_MAIN_CLASS"
	JavaMainModule                                     = "JAVA_MAIN_MODULE"
	JBossHome                                          = "JBOSS_HOME"
	JBossMode                                          = "JBOSS_MODE"
	JBossServerName                                    = "JBOSS_SERVER_NAME"
	KubernetesBasePodName                              = "KUBERNETES_BASE_POD_NAME"
	KubernetesContainerName                            = "KUBERNETES_CONTAINER_NAME"
	KubernetesFullPodName                              = "KUBERNETES_FULL_POD_NAME"
	KubernetesNamespace                                = "KUBERNETES_NAMESPACE"
	KubernetesPodUId                                   = "KUBERNETES_POD_UID"
	NodeJSAppBaseDirectory                             = "NODE_JS_APP_BASE_DIRECTORY"
	NodeJSAppName                                      = "NODE_JS_APP_NAME"
	NodeJSScriptName                                   = "NODE_JS_SCRIPT_NAME"
	PHPScriptPath                                      = "PHP_SCRIPT_PATH"
	PHPWorkingDirectory                                = "PHP_WORKING_DIRECTORY"
	RubyAppRootPath                                    = "RUBY_APP_ROOT_PATH"
	RubyScriptPath                                     = "RUBY_SCRIPT_PATH"
	SoftwareAGInstallRoot                              = "SOFTWAREAG_INSTALL_ROOT"
	SoftwareAGProductPropName                          = "SOFTWAREAG_PRODUCTPROPNAME"
	TibcoBusinessWorksAppNodeName                      = "TIBCO_BUSINESS_WORKS_APP_NODE_NAME"
	TibcoBusinessWorksAppSpaceName                     = "TIBCO_BUSINESS_WORKS_APP_SPACE_NAME"
	TibcoBusinessWorksDomainName                       = "TIBCO_BUSINESS_WORKS_DOMAIN_NAME"
	TibcoBusinessWorksEnginePropertyFile               = "TIBCO_BUSINESS_WORKS_ENGINE_PROPERTY_FILE"
	TibcoBusinessWorksEnginePropertyFilePath           = "TIBCO_BUSINESS_WORKS_ENGINE_PROPERTY_FILE_PATH"
	TibcoBusinessWorksHome                             = "TIBCO_BUSINESS_WORKS_HOME"
	VarnishInstanceName                                = "VARNISH_INSTANCE_NAME"
	WebLogicClusterName                                = "WEB_LOGIC_CLUSTER_NAME"
	WebLogicDomainName                                 = "WEB_LOGIC_DOMAIN_NAME"
	WebLogicHome                                       = "WEB_LOGIC_HOME"
	WebLogicName                                       = "WEB_LOGIC_NAME"
	WebSphereCellname                                  = "WEB_SPHERE_CELL_NAME"
	WebSphereClusterName                               = "WEB_SPHERE_CLUSTER_NAME"
	WebSphereNodeName                                  = "WEB_SPHERE_NODE_NAME"
	WebSphereServerName                                = "WEB_SPHERE_SERVER_NAME"
)

type Source string

const (
	CloudFoundry Source = "CLOUD_FOUNDRY"
	Environment         = "ENVIRONMENT"
	GoogleCloud         = "GOOGLE_CLOUD"
	Kubernetes          = "KUBERNETES"
	Plugin              = "PLUGIN"
)

type DynamicKey struct {
	Key    string `json:"key"`
	Source Source `json:"source"`
}

type ConditionKey struct {
	Attribute  Attribute  `json:"attribute"`
	Type       string     `json:"type,omitempty"`
	DynamicKey DynamicKey `json:"dynamicKey,omitempty"`
}

// Application Detection Rule
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
	Type             RuleType                   `json:"type"`
	Enabled          bool                       `json:"enabled"`
	PropagationTypes []PropagationType          `json:"propagationTypes"`
	Conditions       []ManagementZoneConditions `json:"conditions"`
}

type ManagementZoneConditions struct {
	Key            ConditionKey                 `json:"key"`
	ComparisonInfo ManagementZoneComparisonInfo `json:"comparisonInfo"`
}

type ManagementZoneMetadata struct {
	ConfigurationVersions []int64 `json:"convigurationVersions"`
	ClusterVersion        string  `json:"clusterVersion"`
}

type ManagementZoneComparisonInfo struct {
	Type          string   `json:"type,omitempty"`
	Operator      Operator `json:"operator"`
	Negate        bool     `json:"negate"`
	Value         string   `json:"value"`
	CaseSensitive bool     `json:"caseSensitive,omitempty"`
}

type ManagementZoneType struct {
	Metadata ManagementZoneMetadata `json:"metadata"`
	Id       string                 `json:"id"`
	Name     string                 `json:"name"`
	Rules    []ManagementZoneRule   `json:"rules"`
}

type ManagementZoneResponse struct {
	Values []ManagementZoneType `json:"values"`
}
