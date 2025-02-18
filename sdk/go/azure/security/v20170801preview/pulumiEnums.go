


package v20170801preview

type AlertNotifications string

const (
	// Get notifications on new alerts
	AlertNotificationsOn = AlertNotifications("On")
	// Don't get notifications on new alerts
	AlertNotificationsOff = AlertNotifications("Off")
)

type AlertsToAdmins string

const (
	// Send notification on new alerts to the subscription's admins
	AlertsToAdminsOn = AlertsToAdmins("On")
	// Don't send notification on new alerts to the subscription's admins
	AlertsToAdminsOff = AlertsToAdmins("Off")
)

type DataSource string

const (
	// Devices twin data
	DataSourceTwinData = DataSource("TwinData")
)

type ExportData string

const (
	// Agent raw events
	ExportDataRawEvents = ExportData("RawEvents")
)

type RecommendationConfigStatus string

const (
	RecommendationConfigStatusDisabled = RecommendationConfigStatus("Disabled")
	RecommendationConfigStatusEnabled  = RecommendationConfigStatus("Enabled")
)

type RecommendationType string

const (
	// Authentication schema used for pull an edge module from an ACR repository does not use Service Principal Authentication.
	RecommendationType_IoT_ACRAuthentication = RecommendationType("IoT_ACRAuthentication")
	// IoT agent message size capacity is currently underutilized, causing an increase in the number of sent messages. Adjust message intervals for better utilization.
	RecommendationType_IoT_AgentSendsUnutilizedMessages = RecommendationType("IoT_AgentSendsUnutilizedMessages")
	// Identified security related system configuration issues.
	RecommendationType_IoT_Baseline = RecommendationType("IoT_Baseline")
	// You can optimize Edge Hub memory usage by turning off protocol heads for any protocols not used by Edge modules in your solution.
	RecommendationType_IoT_EdgeHubMemOptimize = RecommendationType("IoT_EdgeHubMemOptimize")
	// Logging is disabled for this edge module.
	RecommendationType_IoT_EdgeLoggingOptions = RecommendationType("IoT_EdgeLoggingOptions")
	// A minority within a device security group has inconsistent Edge Module settings with the rest of their group.
	RecommendationType_IoT_InconsistentModuleSettings = RecommendationType("IoT_InconsistentModuleSettings")
	// Install the Azure Security of Things Agent.
	RecommendationType_IoT_InstallAgent = RecommendationType("IoT_InstallAgent")
	// IP Filter Configuration should have rules defined for allowed traffic and should deny all other traffic by default.
	RecommendationType_IoT_IPFilter_DenyAll = RecommendationType("IoT_IPFilter_DenyAll")
	// An Allow IP Filter rules source IP range is too large. Overly permissive rules might expose your IoT hub to malicious intenders.
	RecommendationType_IoT_IPFilter_PermissiveRule = RecommendationType("IoT_IPFilter_PermissiveRule")
	// A listening endpoint was found on the device.
	RecommendationType_IoT_OpenPorts = RecommendationType("IoT_OpenPorts")
	// An Allowed firewall policy was found (INPUT/OUTPUT). The policy should Deny all traffic by default and define rules to allow necessary communication to/from the device.
	RecommendationType_IoT_PermissiveFirewallPolicy = RecommendationType("IoT_PermissiveFirewallPolicy")
	// A rule in the firewall has been found that contains a permissive pattern for a wide range of IP addresses or Ports.
	RecommendationType_IoT_PermissiveInputFirewallRules = RecommendationType("IoT_PermissiveInputFirewallRules")
	// A rule in the firewall has been found that contains a permissive pattern for a wide range of IP addresses or Ports.
	RecommendationType_IoT_PermissiveOutputFirewallRules = RecommendationType("IoT_PermissiveOutputFirewallRules")
	// Edge module is configured to run in privileged mode, with extensive Linux capabilities or with host-level network access (send/receive data to host machine).
	RecommendationType_IoT_PrivilegedDockerOptions = RecommendationType("IoT_PrivilegedDockerOptions")
	// Same authentication credentials to the IoT Hub used by multiple devices. This could indicate an illegitimate device impersonating a legitimate device. It also exposes the risk of device impersonation by an attacker.
	RecommendationType_IoT_SharedCredentials = RecommendationType("IoT_SharedCredentials")
	// Insecure TLS configurations detected. Immediate upgrade recommended.
	RecommendationType_IoT_VulnerableTLSCipherSuite = RecommendationType("IoT_VulnerableTLSCipherSuite")
)

type SecuritySolutionStatus string

const (
	SecuritySolutionStatusEnabled  = SecuritySolutionStatus("Enabled")
	SecuritySolutionStatusDisabled = SecuritySolutionStatus("Disabled")
)

func init() {
}
