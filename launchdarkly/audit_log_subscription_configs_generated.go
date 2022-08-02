// Code generated by github.com/stelminator/terraform-provider-launchdarkly/scripts/codegen DO NOT EDIT.

package launchdarkly

var SUBSCRIPTION_CONFIGURATION_FIELDS = map[string]IntegrationConfig{
	"datadog": {
		"apiKey": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter your Datadog [API key](https://app.datadoghq.com/organization-settings/api-keys).",
			IsOptional:    false,
			IsSecret:      true,
			Type:          "string",
		},
		"hideMemberDetails": {
			AllowedValues: []string{},
			DefaultValue:  false,
			Description:   "Don't send related member email and names.",
			IsOptional:    true,
			IsSecret:      false,
			Type:          "boolean",
		},
		"hostURL": {
			AllowedValues: []string{"https://api.datadoghq.com", "https://api.datadoghq.eu", "https://us3.datadoghq.com", "https://us5.datadoghq.com", "https://app.ddog-gov.com"},
			DefaultValue:  "https://api.datadoghq.com",
			Description:   "Your Datadog host URL. Read [How do I tell which Datadog site I am on?](https://docs.datadoghq.com/getting_started/site/#how-do-i-tell-which-datadog-site-i-am-on) if you are unsure which host URL to select.",
			IsOptional:    true,
			IsSecret:      false,
			Type:          "enum",
		},
	},
	"dynatrace": {
		"apiToken": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter your Dynatrace API token. [Learn more](https://www.dynatrace.com/support/help/shortlink/api-authentication#generate-a-token) about generating tokens. The 'access problem and event feed, metrics, and topology' scope is required.",
			IsOptional:    false,
			IsSecret:      true,
			Type:          "string",
		},
		"entity": {
			AllowedValues: []string{"APPLICATION", "APPLICATION_METHOD", "APPLICATION_METHOD_GROUP", "AUTO_SCALING_GROUP", "AUXILIARY_SYNTHETIC_TEST", "AWS_APPLICATION_LOAD_BALANCER", "AWS_AVAILABILITY_ZONE", "AWS_CREDENTIALS", "AWS_LAMBDA_FUNCTION", "AWS_NETWORK_LOAD_BALANCER", "AZURE_API_MANAGEMENT_SERVICE", "AZURE_APPLICATION_GATEWAY", "AZURE_COSMOS_DB", "AZURE_CREDENTIALS", "AZURE_EVENT_HUB", "AZURE_EVENT_HUB_NAMESPACE", "AZURE_FUNCTION_APP", "AZURE_IOT_HUB", "AZURE_LOAD_BALANCER", "AZURE_MGMT_GROUP", "AZURE_REDIS_CACHE", "AZURE_REGION", "AZURE_SERVICE_BUS_NAMESPACE", "AZURE_SERVICE_BUS_QUEUE", "AZURE_SERVICE_BUS_TOPIC", "AZURE_SQL_DATABASE", "AZURE_SQL_ELASTIC_POOL", "AZURE_SQL_SERVER", "AZURE_STORAGE_ACCOUNT", "AZURE_SUBSCRIPTION", "AZURE_TENANT", "AZURE_VM", "AZURE_VM_SCALE_SET", "AZURE_WEB_APP", "CF_APPLICATION", "CF_FOUNDATION", "CINDER_VOLUME", "CLOUD_APPLICATION", "CLOUD_APPLICATION_INSTANCE", "CLOUD_APPLICATION_NAMESPACE", "CONTAINER_GROUP", "CONTAINER_GROUP_INSTANCE", "CUSTOM_APPLICATION", "CUSTOM_DEVICE", "CUSTOM_DEVICE_GROUP", "DCRUM_APPLICATION", "DCRUM_SERVICE", "DCRUM_SERVICE_INSTANCE", "DEVICE_APPLICATION_METHOD", "DISK", "DOCKER_CONTAINER_GROUP_INSTANCE", "DYNAMO_DB_TABLE", "EBS_VOLUME", "EC2_INSTANCE", "ELASTIC_LOAD_BALANCER", "ENVIRONMENT", "EXTERNAL_SYNTHETIC_TEST_STEP", "GCP_ZONE", "GEOLOCATION", "GEOLOC_SITE", "GOOGLE_COMPUTE_ENGINE", "HOST", "HOST_GROUP", "HTTP_CHECK", "HTTP_CHECK_STEP", "HYPERVISOR", "KUBERNETES_CLUSTER", "KUBERNETES_NODE", "MOBILE_APPLICATION", "NETWORK_INTERFACE", "NEUTRON_SUBNET", "OPENSTACK_PROJECT", "OPENSTACK_REGION", "OPENSTACK_VM", "OS", "PROCESS_GROUP", "PROCESS_GROUP_INSTANCE", "RELATIONAL_DATABASE_SERVICE", "SERVICE", "SERVICE_INSTANCE", "SERVICE_METHOD", "SERVICE_METHOD_GROUP", "SWIFT_CONTAINER", "SYNTHETIC_LOCATION", "SYNTHETIC_TEST", "SYNTHETIC_TEST_STEP", "VIRTUALMACHINE", "VMWARE_DATACENTER"},
			DefaultValue:  "APPLICATION",
			Description:   "Select the Dynatrace entity `meType` to associate with all events.",
			IsOptional:    true,
			IsSecret:      false,
			Type:          "enum",
		},
		"url": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter the URL used to access your Dynatrace (managed or hosted) service. Follow the pattern shown in the placeholder text.",
			IsOptional:    false,
			IsSecret:      false,
			Type:          "uri",
		},
	},
	"elastic": {
		"index": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter the index name where you want to store your LaunchDarkly data",
			IsOptional:    false,
			IsSecret:      false,
			Type:          "string",
		},
		"token": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter the base64 _credentials_ based on your [API Key](https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html)",
			IsOptional:    false,
			IsSecret:      true,
			Type:          "string",
		},
		"url": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter the URL for your Elasticsearch endpoint, including the socket",
			IsOptional:    false,
			IsSecret:      false,
			Type:          "uri",
		},
	},
	"honeycomb": {
		"apiKey": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter your [Honeycomb API key](https://ui.honeycomb.io/teams).",
			IsOptional:    false,
			IsSecret:      true,
			Type:          "string",
		},
		"datasetName": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter the name of your Honeycomb dataset. This associates LaunchDarkly data with the corresponding Honeycomb dataset.",
			IsOptional:    false,
			IsSecret:      false,
			Type:          "string",
		},
	},
	"logdna": {
		"ingestionKey": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter your [LogDNA ingestion key](https://app.logdna.com/manage/api-keys).",
			IsOptional:    false,
			IsSecret:      true,
			Type:          "string",
		},
		"level": {
			AllowedValues: []string{},
			DefaultValue:  "INFO",
			Description:   "The log level with which LaunchDarkly messages will be published.",
			IsOptional:    true,
			IsSecret:      false,
			Type:          "string",
		},
	},
	"msteams": {"url": {
		AllowedValues: []string{},
		DefaultValue:  nil,
		Description:   "Enter your Microsoft Teams [incoming webhook URL](https://docs.launchdarkly.com/integrations/microsoft-teams#setting-up-a-connector-in-microsoft-teams).",
		IsOptional:    false,
		IsSecret:      false,
		Type:          "uri",
	}},
	"new-relic-apm": {
		"apiKey": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter your [New Relic REST API key](https://docs.newrelic.com/docs/apis/get-started/intro-apis/types-new-relic-api-keys#rest-api-key).",
			IsOptional:    false,
			IsSecret:      true,
			Type:          "string",
		},
		"applicationId": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter your [New Relic application ID](https://docs.newrelic.com/docs/accounts/install-new-relic/account-setup/app-id-other-product-ids#ui).",
			IsOptional:    false,
			IsSecret:      false,
			Type:          "string",
		},
		"domain": {
			AllowedValues: []string{"api.newrelic.com", "api.eu.newrelic.com"},
			DefaultValue:  "api.newrelic.com",
			Description:   "Your New Relic data center. The default(US) is \"api.newrelic.com\". Use \"api.eu.newrelic.com\" if you are using the EU data center.",
			IsOptional:    true,
			IsSecret:      false,
			Type:          "enum",
		},
	},
	"signalfx": {
		"accessToken": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter your [SignalFx access token](https://docs.signalfx.com/en/latest/admin-guide/tokens.html#working-with-access-tokens).",
			IsOptional:    false,
			IsSecret:      true,
			Type:          "string",
		},
		"realm": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter your [SignalFx realm](https://developers.signalfx.com/#realms-in-endpoints).",
			IsOptional:    false,
			IsSecret:      false,
			Type:          "string",
		},
	},
	"splunk": {
		"base-url": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter your [Splunk HTTP event collector base URL](https://docs.splunk.com/Documentation/Splunk/latest/Data/UsetheHTTPEventCollector).",
			IsOptional:    false,
			IsSecret:      false,
			Type:          "string",
		},
		"skip-ca-verification": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Splunk Cloud instances sign their own SSL certificates by default. If you use Splunk Cloud, you may need to skip certificate validation in order for this integration to work.",
			IsOptional:    false,
			IsSecret:      false,
			Type:          "boolean",
		},
		"token": {
			AllowedValues: []string{},
			DefaultValue:  nil,
			Description:   "Enter your HTTP event collector token value.",
			IsOptional:    false,
			IsSecret:      true,
			Type:          "string",
		},
	},
}
