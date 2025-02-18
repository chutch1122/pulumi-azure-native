


package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionType string

const (
	// Modify an object's properties
	ActionTypeModifyProperties = ActionType("ModifyProperties")
	// Run a playbook on an object
	ActionTypeRunPlaybook = ActionType("RunPlaybook")
)

type AlertDetail string

const (
	// Alert display name
	AlertDetailDisplayName = AlertDetail("DisplayName")
	// Alert severity
	AlertDetailSeverity = AlertDetail("Severity")
)

type AlertRuleKind string

const (
	AlertRuleKindScheduled                         = AlertRuleKind("Scheduled")
	AlertRuleKindMicrosoftSecurityIncidentCreation = AlertRuleKind("MicrosoftSecurityIncidentCreation")
	AlertRuleKindFusion                            = AlertRuleKind("Fusion")
)

type AlertSeverity string

const (
	// High severity
	AlertSeverityHigh = AlertSeverity("High")
	// Medium severity
	AlertSeverityMedium = AlertSeverity("Medium")
	// Low severity
	AlertSeverityLow = AlertSeverity("Low")
	// Informational severity
	AlertSeverityInformational = AlertSeverity("Informational")
)

type AttackTactic string

const (
	AttackTacticReconnaissance          = AttackTactic("Reconnaissance")
	AttackTacticResourceDevelopment     = AttackTactic("ResourceDevelopment")
	AttackTacticInitialAccess           = AttackTactic("InitialAccess")
	AttackTacticExecution               = AttackTactic("Execution")
	AttackTacticPersistence             = AttackTactic("Persistence")
	AttackTacticPrivilegeEscalation     = AttackTactic("PrivilegeEscalation")
	AttackTacticDefenseEvasion          = AttackTactic("DefenseEvasion")
	AttackTacticCredentialAccess        = AttackTactic("CredentialAccess")
	AttackTacticDiscovery               = AttackTactic("Discovery")
	AttackTacticLateralMovement         = AttackTactic("LateralMovement")
	AttackTacticCollection              = AttackTactic("Collection")
	AttackTacticExfiltration            = AttackTactic("Exfiltration")
	AttackTacticCommandAndControl       = AttackTactic("CommandAndControl")
	AttackTacticImpact                  = AttackTactic("Impact")
	AttackTacticPreAttack               = AttackTactic("PreAttack")
	AttackTacticImpairProcessControl    = AttackTactic("ImpairProcessControl")
	AttackTacticInhibitResponseFunction = AttackTactic("InhibitResponseFunction")
)

type AutomationRulePropertyConditionSupportedOperator string

const (
	// Evaluates if the property equals at least one of the condition values
	AutomationRulePropertyConditionSupportedOperatorEquals = AutomationRulePropertyConditionSupportedOperator("Equals")
	// Evaluates if the property does not equal any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotEquals = AutomationRulePropertyConditionSupportedOperator("NotEquals")
	// Evaluates if the property contains at least one of the condition values
	AutomationRulePropertyConditionSupportedOperatorContains = AutomationRulePropertyConditionSupportedOperator("Contains")
	// Evaluates if the property does not contain any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotContains = AutomationRulePropertyConditionSupportedOperator("NotContains")
	// Evaluates if the property starts with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorStartsWith = AutomationRulePropertyConditionSupportedOperator("StartsWith")
	// Evaluates if the property does not start with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotStartsWith = AutomationRulePropertyConditionSupportedOperator("NotStartsWith")
	// Evaluates if the property ends with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorEndsWith = AutomationRulePropertyConditionSupportedOperator("EndsWith")
	// Evaluates if the property does not end with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotEndsWith = AutomationRulePropertyConditionSupportedOperator("NotEndsWith")
)

type AutomationRulePropertyConditionSupportedProperty string

const (
	// The title of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentTitle = AutomationRulePropertyConditionSupportedProperty("IncidentTitle")
	// The description of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentDescription = AutomationRulePropertyConditionSupportedProperty("IncidentDescription")
	// The severity of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentSeverity = AutomationRulePropertyConditionSupportedProperty("IncidentSeverity")
	// The status of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentStatus = AutomationRulePropertyConditionSupportedProperty("IncidentStatus")
	// The related Analytic rule ids of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentRelatedAnalyticRuleIds = AutomationRulePropertyConditionSupportedProperty("IncidentRelatedAnalyticRuleIds")
	// The tactics of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentTactics = AutomationRulePropertyConditionSupportedProperty("IncidentTactics")
	// The labels of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentLabel = AutomationRulePropertyConditionSupportedProperty("IncidentLabel")
	// The provider name of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentProviderName = AutomationRulePropertyConditionSupportedProperty("IncidentProviderName")
	// The account Azure Active Directory tenant id
	AutomationRulePropertyConditionSupportedPropertyAccountAadTenantId = AutomationRulePropertyConditionSupportedProperty("AccountAadTenantId")
	// The account Azure Active Directory user id
	AutomationRulePropertyConditionSupportedPropertyAccountAadUserId = AutomationRulePropertyConditionSupportedProperty("AccountAadUserId")
	// The account name
	AutomationRulePropertyConditionSupportedPropertyAccountName = AutomationRulePropertyConditionSupportedProperty("AccountName")
	// The account NetBIOS domain name
	AutomationRulePropertyConditionSupportedPropertyAccountNTDomain = AutomationRulePropertyConditionSupportedProperty("AccountNTDomain")
	// The account Azure Active Directory Passport User ID
	AutomationRulePropertyConditionSupportedPropertyAccountPUID = AutomationRulePropertyConditionSupportedProperty("AccountPUID")
	// The account security identifier
	AutomationRulePropertyConditionSupportedPropertyAccountSid = AutomationRulePropertyConditionSupportedProperty("AccountSid")
	// The account unique identifier
	AutomationRulePropertyConditionSupportedPropertyAccountObjectGuid = AutomationRulePropertyConditionSupportedProperty("AccountObjectGuid")
	// The account user principal name suffix
	AutomationRulePropertyConditionSupportedPropertyAccountUPNSuffix = AutomationRulePropertyConditionSupportedProperty("AccountUPNSuffix")
	// The name of the product of the alert
	AutomationRulePropertyConditionSupportedPropertyAlertProductNames = AutomationRulePropertyConditionSupportedProperty("AlertProductNames")
	// The Azure resource id
	AutomationRulePropertyConditionSupportedPropertyAzureResourceResourceId = AutomationRulePropertyConditionSupportedProperty("AzureResourceResourceId")
	// The Azure resource subscription id
	AutomationRulePropertyConditionSupportedPropertyAzureResourceSubscriptionId = AutomationRulePropertyConditionSupportedProperty("AzureResourceSubscriptionId")
	// The cloud application identifier
	AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppId = AutomationRulePropertyConditionSupportedProperty("CloudApplicationAppId")
	// The cloud application name
	AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppName = AutomationRulePropertyConditionSupportedProperty("CloudApplicationAppName")
	// The dns record domain name
	AutomationRulePropertyConditionSupportedPropertyDNSDomainName = AutomationRulePropertyConditionSupportedProperty("DNSDomainName")
	// The file directory full path
	AutomationRulePropertyConditionSupportedPropertyFileDirectory = AutomationRulePropertyConditionSupportedProperty("FileDirectory")
	// The file name without path
	AutomationRulePropertyConditionSupportedPropertyFileName = AutomationRulePropertyConditionSupportedProperty("FileName")
	// The file hash value
	AutomationRulePropertyConditionSupportedPropertyFileHashValue = AutomationRulePropertyConditionSupportedProperty("FileHashValue")
	// The host Azure resource id
	AutomationRulePropertyConditionSupportedPropertyHostAzureID = AutomationRulePropertyConditionSupportedProperty("HostAzureID")
	// The host name without domain
	AutomationRulePropertyConditionSupportedPropertyHostName = AutomationRulePropertyConditionSupportedProperty("HostName")
	// The host NetBIOS name
	AutomationRulePropertyConditionSupportedPropertyHostNetBiosName = AutomationRulePropertyConditionSupportedProperty("HostNetBiosName")
	// The host NT domain
	AutomationRulePropertyConditionSupportedPropertyHostNTDomain = AutomationRulePropertyConditionSupportedProperty("HostNTDomain")
	// The host operating system
	AutomationRulePropertyConditionSupportedPropertyHostOSVersion = AutomationRulePropertyConditionSupportedProperty("HostOSVersion")
	// "The IoT device id
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceId = AutomationRulePropertyConditionSupportedProperty("IoTDeviceId")
	// The IoT device name
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceName = AutomationRulePropertyConditionSupportedProperty("IoTDeviceName")
	// The IoT device type
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceType = AutomationRulePropertyConditionSupportedProperty("IoTDeviceType")
	// The IoT device vendor
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceVendor = AutomationRulePropertyConditionSupportedProperty("IoTDeviceVendor")
	// The IoT device model
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceModel = AutomationRulePropertyConditionSupportedProperty("IoTDeviceModel")
	// The IoT device operating system
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceOperatingSystem = AutomationRulePropertyConditionSupportedProperty("IoTDeviceOperatingSystem")
	// The IP address
	AutomationRulePropertyConditionSupportedPropertyIPAddress = AutomationRulePropertyConditionSupportedProperty("IPAddress")
	// The mailbox display name
	AutomationRulePropertyConditionSupportedPropertyMailboxDisplayName = AutomationRulePropertyConditionSupportedProperty("MailboxDisplayName")
	// The mailbox primary address
	AutomationRulePropertyConditionSupportedPropertyMailboxPrimaryAddress = AutomationRulePropertyConditionSupportedProperty("MailboxPrimaryAddress")
	// The mailbox user principal name
	AutomationRulePropertyConditionSupportedPropertyMailboxUPN = AutomationRulePropertyConditionSupportedProperty("MailboxUPN")
	// The mail message delivery action
	AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryAction = AutomationRulePropertyConditionSupportedProperty("MailMessageDeliveryAction")
	// The mail message delivery location
	AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryLocation = AutomationRulePropertyConditionSupportedProperty("MailMessageDeliveryLocation")
	// The mail message recipient
	AutomationRulePropertyConditionSupportedPropertyMailMessageRecipient = AutomationRulePropertyConditionSupportedProperty("MailMessageRecipient")
	// The mail message sender IP address
	AutomationRulePropertyConditionSupportedPropertyMailMessageSenderIP = AutomationRulePropertyConditionSupportedProperty("MailMessageSenderIP")
	// The mail message subject
	AutomationRulePropertyConditionSupportedPropertyMailMessageSubject = AutomationRulePropertyConditionSupportedProperty("MailMessageSubject")
	// The mail message P1 sender
	AutomationRulePropertyConditionSupportedPropertyMailMessageP1Sender = AutomationRulePropertyConditionSupportedProperty("MailMessageP1Sender")
	// The mail message P2 sender
	AutomationRulePropertyConditionSupportedPropertyMailMessageP2Sender = AutomationRulePropertyConditionSupportedProperty("MailMessageP2Sender")
	// The malware category
	AutomationRulePropertyConditionSupportedPropertyMalwareCategory = AutomationRulePropertyConditionSupportedProperty("MalwareCategory")
	// The malware name
	AutomationRulePropertyConditionSupportedPropertyMalwareName = AutomationRulePropertyConditionSupportedProperty("MalwareName")
	// The process execution command line
	AutomationRulePropertyConditionSupportedPropertyProcessCommandLine = AutomationRulePropertyConditionSupportedProperty("ProcessCommandLine")
	// The process id
	AutomationRulePropertyConditionSupportedPropertyProcessId = AutomationRulePropertyConditionSupportedProperty("ProcessId")
	// The registry key path
	AutomationRulePropertyConditionSupportedPropertyRegistryKey = AutomationRulePropertyConditionSupportedProperty("RegistryKey")
	// The registry key value in string formatted representation
	AutomationRulePropertyConditionSupportedPropertyRegistryValueData = AutomationRulePropertyConditionSupportedProperty("RegistryValueData")
	// The url
	AutomationRulePropertyConditionSupportedPropertyUrl = AutomationRulePropertyConditionSupportedProperty("Url")
)

type ConditionType string

const (
	// Evaluate an object property value
	ConditionTypeProperty = ConditionType("Property")
)

type DataConnectorKind string

const (
	DataConnectorKindAzureActiveDirectory                      = DataConnectorKind("AzureActiveDirectory")
	DataConnectorKindAzureSecurityCenter                       = DataConnectorKind("AzureSecurityCenter")
	DataConnectorKindMicrosoftCloudAppSecurity                 = DataConnectorKind("MicrosoftCloudAppSecurity")
	DataConnectorKindThreatIntelligence                        = DataConnectorKind("ThreatIntelligence")
	DataConnectorKindOffice365                                 = DataConnectorKind("Office365")
	DataConnectorKindAmazonWebServicesCloudTrail               = DataConnectorKind("AmazonWebServicesCloudTrail")
	DataConnectorKindAzureAdvancedThreatProtection             = DataConnectorKind("AzureAdvancedThreatProtection")
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection = DataConnectorKind("MicrosoftDefenderAdvancedThreatProtection")
)

type DataTypeState string

const (
	DataTypeStateEnabled  = DataTypeState("Enabled")
	DataTypeStateDisabled = DataTypeState("Disabled")
)

type EntityMappingType string

const (
	// User account entity type
	EntityMappingTypeAccount = EntityMappingType("Account")
	// Host entity type
	EntityMappingTypeHost = EntityMappingType("Host")
	// IP address entity type
	EntityMappingTypeIP = EntityMappingType("IP")
	// Malware entity type
	EntityMappingTypeMalware = EntityMappingType("Malware")
	// System file entity type
	EntityMappingTypeFile = EntityMappingType("File")
	// Process entity type
	EntityMappingTypeProcess = EntityMappingType("Process")
	// Cloud app entity type
	EntityMappingTypeCloudApplication = EntityMappingType("CloudApplication")
	// DNS entity type
	EntityMappingTypeDNS = EntityMappingType("DNS")
	// Azure resource entity type
	EntityMappingTypeAzureResource = EntityMappingType("AzureResource")
	// File-hash entity type
	EntityMappingTypeFileHash = EntityMappingType("FileHash")
	// Registry key entity type
	EntityMappingTypeRegistryKey = EntityMappingType("RegistryKey")
	// Registry value entity type
	EntityMappingTypeRegistryValue = EntityMappingType("RegistryValue")
	// Security group entity type
	EntityMappingTypeSecurityGroup = EntityMappingType("SecurityGroup")
	// URL entity type
	EntityMappingTypeURL = EntityMappingType("URL")
	// Mailbox entity type
	EntityMappingTypeMailbox = EntityMappingType("Mailbox")
	// Mail cluster entity type
	EntityMappingTypeMailCluster = EntityMappingType("MailCluster")
	// Mail message entity type
	EntityMappingTypeMailMessage = EntityMappingType("MailMessage")
	// Submission mail entity type
	EntityMappingTypeSubmissionMail = EntityMappingType("SubmissionMail")
)

type EventGroupingAggregationKind string

const (
	EventGroupingAggregationKindSingleAlert    = EventGroupingAggregationKind("SingleAlert")
	EventGroupingAggregationKindAlertPerResult = EventGroupingAggregationKind("AlertPerResult")
)

type IncidentClassification string

const (
	// Incident classification was undetermined
	IncidentClassificationUndetermined = IncidentClassification("Undetermined")
	// Incident was true positive
	IncidentClassificationTruePositive = IncidentClassification("TruePositive")
	// Incident was benign positive
	IncidentClassificationBenignPositive = IncidentClassification("BenignPositive")
	// Incident was false positive
	IncidentClassificationFalsePositive = IncidentClassification("FalsePositive")
)

type IncidentClassificationReason string

const (
	// Classification reason was suspicious activity
	IncidentClassificationReasonSuspiciousActivity = IncidentClassificationReason("SuspiciousActivity")
	// Classification reason was suspicious but expected
	IncidentClassificationReasonSuspiciousButExpected = IncidentClassificationReason("SuspiciousButExpected")
	// Classification reason was incorrect alert logic
	IncidentClassificationReasonIncorrectAlertLogic = IncidentClassificationReason("IncorrectAlertLogic")
	// Classification reason was inaccurate data
	IncidentClassificationReasonInaccurateData = IncidentClassificationReason("InaccurateData")
)

type IncidentSeverity string

const (
	// High severity
	IncidentSeverityHigh = IncidentSeverity("High")
	// Medium severity
	IncidentSeverityMedium = IncidentSeverity("Medium")
	// Low severity
	IncidentSeverityLow = IncidentSeverity("Low")
	// Informational severity
	IncidentSeverityInformational = IncidentSeverity("Informational")
)

type IncidentStatus string

const (
	// An active incident which isn't being handled currently
	IncidentStatusNew = IncidentStatus("New")
	// An active incident which is being handled
	IncidentStatusActive = IncidentStatus("Active")
	// A non-active incident
	IncidentStatusClosed = IncidentStatus("Closed")
)

type MatchingMethod string

const (
	// Grouping alerts into a single incident if all the entities match
	MatchingMethodAllEntities = MatchingMethod("AllEntities")
	// Grouping any alerts triggered by this rule into a single incident
	MatchingMethodAnyAlert = MatchingMethod("AnyAlert")
	// Grouping alerts into a single incident if the selected entities, custom details and alert details match
	MatchingMethodSelected = MatchingMethod("Selected")
)

type MicrosoftSecurityProductName string

const (
	MicrosoftSecurityProductName_Microsoft_Cloud_App_Security               = MicrosoftSecurityProductName("Microsoft Cloud App Security")
	MicrosoftSecurityProductName_Azure_Security_Center                      = MicrosoftSecurityProductName("Azure Security Center")
	MicrosoftSecurityProductName_Azure_Advanced_Threat_Protection           = MicrosoftSecurityProductName("Azure Advanced Threat Protection")
	MicrosoftSecurityProductName_Azure_Active_Directory_Identity_Protection = MicrosoftSecurityProductName("Azure Active Directory Identity Protection")
	MicrosoftSecurityProductName_Azure_Security_Center_for_IoT              = MicrosoftSecurityProductName("Azure Security Center for IoT")
)

type OwnerType string

const (
	// The incident owner type is unknown
	OwnerTypeUnknown = OwnerType("Unknown")
	// The incident owner type is an AAD user
	OwnerTypeUser = OwnerType("User")
	// The incident owner type is an AAD group
	OwnerTypeGroup = OwnerType("Group")
)

type Source string

const (
	Source_Local_file     = Source("Local file")
	Source_Remote_storage = Source("Remote storage")
)

type ThreatIntelligenceResourceInnerKind string

const (
	// Entity represents threat intelligence indicator in the system.
	ThreatIntelligenceResourceInnerKindIndicator = ThreatIntelligenceResourceInnerKind("indicator")
)

type TriggerOperator string

const (
	TriggerOperatorGreaterThan = TriggerOperator("GreaterThan")
	TriggerOperatorLessThan    = TriggerOperator("LessThan")
	TriggerOperatorEqual       = TriggerOperator("Equal")
	TriggerOperatorNotEqual    = TriggerOperator("NotEqual")
)

func (TriggerOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerOperator)(nil)).Elem()
}

func (e TriggerOperator) ToTriggerOperatorOutput() TriggerOperatorOutput {
	return pulumi.ToOutput(e).(TriggerOperatorOutput)
}

func (e TriggerOperator) ToTriggerOperatorOutputWithContext(ctx context.Context) TriggerOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TriggerOperatorOutput)
}

func (e TriggerOperator) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return e.ToTriggerOperatorPtrOutputWithContext(context.Background())
}

func (e TriggerOperator) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return TriggerOperator(e).ToTriggerOperatorOutputWithContext(ctx).ToTriggerOperatorPtrOutputWithContext(ctx)
}

func (e TriggerOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggerOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TriggerOperatorOutput struct{ *pulumi.OutputState }

func (TriggerOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerOperator)(nil)).Elem()
}

func (o TriggerOperatorOutput) ToTriggerOperatorOutput() TriggerOperatorOutput {
	return o
}

func (o TriggerOperatorOutput) ToTriggerOperatorOutputWithContext(ctx context.Context) TriggerOperatorOutput {
	return o
}

func (o TriggerOperatorOutput) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return o.ToTriggerOperatorPtrOutputWithContext(context.Background())
}

func (o TriggerOperatorOutput) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerOperator) *TriggerOperator {
		return &v
	}).(TriggerOperatorPtrOutput)
}

func (o TriggerOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TriggerOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TriggerOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TriggerOperatorPtrOutput struct{ *pulumi.OutputState }

func (TriggerOperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerOperator)(nil)).Elem()
}

func (o TriggerOperatorPtrOutput) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return o
}

func (o TriggerOperatorPtrOutput) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return o
}

func (o TriggerOperatorPtrOutput) Elem() TriggerOperatorOutput {
	return o.ApplyT(func(v *TriggerOperator) TriggerOperator {
		if v != nil {
			return *v
		}
		var ret TriggerOperator
		return ret
	}).(TriggerOperatorOutput)
}

func (o TriggerOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TriggerOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TriggerOperatorInput interface {
	pulumi.Input

	ToTriggerOperatorOutput() TriggerOperatorOutput
	ToTriggerOperatorOutputWithContext(context.Context) TriggerOperatorOutput
}

var triggerOperatorPtrType = reflect.TypeOf((**TriggerOperator)(nil)).Elem()

type TriggerOperatorPtrInput interface {
	pulumi.Input

	ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput
	ToTriggerOperatorPtrOutputWithContext(context.Context) TriggerOperatorPtrOutput
}

type triggerOperatorPtr string

func TriggerOperatorPtr(v string) TriggerOperatorPtrInput {
	return (*triggerOperatorPtr)(&v)
}

func (*triggerOperatorPtr) ElementType() reflect.Type {
	return triggerOperatorPtrType
}

func (in *triggerOperatorPtr) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return pulumi.ToOutput(in).(TriggerOperatorPtrOutput)
}

func (in *triggerOperatorPtr) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TriggerOperatorPtrOutput)
}

type TriggersOn string

const (
	// Trigger on Incidents
	TriggersOnIncidents = TriggersOn("Incidents")
)

type TriggersWhen string

const (
	// Trigger on created objects
	TriggersWhenCreated = TriggersWhen("Created")
)

func init() {
	pulumi.RegisterOutputType(TriggerOperatorOutput{})
	pulumi.RegisterOutputType(TriggerOperatorPtrOutput{})
}
