


package providerhub

type ExtendedLocationType string

const (
	ExtendedLocationTypeNotSpecified = ExtendedLocationType("NotSpecified")
	ExtendedLocationTypeEdgeZone     = ExtendedLocationType("EdgeZone")
	ExtendedLocationTypeArcZone      = ExtendedLocationType("ArcZone")
)

type ExtensionCategory string

const (
	ExtensionCategoryNotSpecified                      = ExtensionCategory("NotSpecified")
	ExtensionCategoryResourceCreationValidate          = ExtensionCategory("ResourceCreationValidate")
	ExtensionCategoryResourceCreationBegin             = ExtensionCategory("ResourceCreationBegin")
	ExtensionCategoryResourceCreationCompleted         = ExtensionCategory("ResourceCreationCompleted")
	ExtensionCategoryResourceReadValidate              = ExtensionCategory("ResourceReadValidate")
	ExtensionCategoryResourceReadBegin                 = ExtensionCategory("ResourceReadBegin")
	ExtensionCategoryResourcePatchValidate             = ExtensionCategory("ResourcePatchValidate")
	ExtensionCategoryResourcePatchCompleted            = ExtensionCategory("ResourcePatchCompleted")
	ExtensionCategoryResourceDeletionValidate          = ExtensionCategory("ResourceDeletionValidate")
	ExtensionCategoryResourceDeletionBegin             = ExtensionCategory("ResourceDeletionBegin")
	ExtensionCategoryResourceDeletionCompleted         = ExtensionCategory("ResourceDeletionCompleted")
	ExtensionCategoryResourcePostAction                = ExtensionCategory("ResourcePostAction")
	ExtensionCategorySubscriptionLifecycleNotification = ExtensionCategory("SubscriptionLifecycleNotification")
	ExtensionCategoryResourcePatchBegin                = ExtensionCategory("ResourcePatchBegin")
	ExtensionCategoryResourceMoveBegin                 = ExtensionCategory("ResourceMoveBegin")
	ExtensionCategoryResourceMoveCompleted             = ExtensionCategory("ResourceMoveCompleted")
)

type ExtensionOptionType string

const (
	ExtensionOptionTypeNotSpecified                                  = ExtensionOptionType("NotSpecified")
	ExtensionOptionTypeDoNotMergeExistingReadOnlyAndSecretProperties = ExtensionOptionType("DoNotMergeExistingReadOnlyAndSecretProperties")
	ExtensionOptionTypeIncludeInternalMetadata                       = ExtensionOptionType("IncludeInternalMetadata")
)

type FeaturesPolicy string

const (
	FeaturesPolicyAny = FeaturesPolicy("Any")
	FeaturesPolicyAll = FeaturesPolicy("All")
)

type IdentityManagementTypes string

const (
	IdentityManagementTypesNotSpecified              = IdentityManagementTypes("NotSpecified")
	IdentityManagementTypesSystemAssigned            = IdentityManagementTypes("SystemAssigned")
	IdentityManagementTypesUserAssigned              = IdentityManagementTypes("UserAssigned")
	IdentityManagementTypesActor                     = IdentityManagementTypes("Actor")
	IdentityManagementTypesDelegatedResourceIdentity = IdentityManagementTypes("DelegatedResourceIdentity")
)

type LoggingDetails string

const (
	LoggingDetailsNone = LoggingDetails("None")
	LoggingDetailsBody = LoggingDetails("Body")
)

type LoggingDirections string

const (
	LoggingDirectionsNone     = LoggingDirections("None")
	LoggingDirectionsRequest  = LoggingDirections("Request")
	LoggingDirectionsResponse = LoggingDirections("Response")
)

type MarketplaceType string

const (
	MarketplaceTypeNotSpecified = MarketplaceType("NotSpecified")
	MarketplaceTypeAddOn        = MarketplaceType("AddOn")
	MarketplaceTypeBypass       = MarketplaceType("Bypass")
	MarketplaceTypeStore        = MarketplaceType("Store")
)

type MessageScope string

const (
	MessageScopeNotSpecified            = MessageScope("NotSpecified")
	MessageScopeRegisteredSubscriptions = MessageScope("RegisteredSubscriptions")
)

type NotificationMode string

const (
	NotificationModeNotSpecified = NotificationMode("NotSpecified")
	NotificationModeEventHub     = NotificationMode("EventHub")
	NotificationModeWebHook      = NotificationMode("WebHook")
)

type OperationActionType string

const (
	OperationActionTypeNotSpecified = OperationActionType("NotSpecified")
	OperationActionTypeInternal     = OperationActionType("Internal")
)

type OperationOrigins string

const (
	OperationOriginsNotSpecified = OperationOrigins("NotSpecified")
	OperationOriginsUser         = OperationOrigins("User")
	OperationOriginsSystem       = OperationOrigins("System")
)

type OptInHeaderType string

const (
	OptInHeaderTypeNotSpecified                   = OptInHeaderType("NotSpecified")
	OptInHeaderTypeSignedUserToken                = OptInHeaderType("SignedUserToken")
	OptInHeaderTypeClientGroupMembership          = OptInHeaderType("ClientGroupMembership")
	OptInHeaderTypeSignedAuxiliaryTokens          = OptInHeaderType("SignedAuxiliaryTokens")
	OptInHeaderTypeUnboundedClientGroupMembership = OptInHeaderType("UnboundedClientGroupMembership")
)

type PreflightOption string

const (
	PreflightOptionNone                        = PreflightOption("None")
	PreflightOptionContinueDeploymentOnFailure = PreflightOption("ContinueDeploymentOnFailure")
	PreflightOptionDefaultValidationOnly       = PreflightOption("DefaultValidationOnly")
)

type ProvisioningState string

const (
	ProvisioningStateNotSpecified      = ProvisioningState("NotSpecified")
	ProvisioningStateAccepted          = ProvisioningState("Accepted")
	ProvisioningStateRunning           = ProvisioningState("Running")
	ProvisioningStateCreating          = ProvisioningState("Creating")
	ProvisioningStateCreated           = ProvisioningState("Created")
	ProvisioningStateDeleting          = ProvisioningState("Deleting")
	ProvisioningStateDeleted           = ProvisioningState("Deleted")
	ProvisioningStateCanceled          = ProvisioningState("Canceled")
	ProvisioningStateFailed            = ProvisioningState("Failed")
	ProvisioningStateSucceeded         = ProvisioningState("Succeeded")
	ProvisioningStateMovingResources   = ProvisioningState("MovingResources")
	ProvisioningStateTransientFailure  = ProvisioningState("TransientFailure")
	ProvisioningStateRolloutInProgress = ProvisioningState("RolloutInProgress")
)

type Regionality string

const (
	RegionalityNotSpecified = Regionality("NotSpecified")
	RegionalityGlobal       = Regionality("Global")
	RegionalityRegional     = Regionality("Regional")
)

type ResourceAccessPolicy string

const (
	ResourceAccessPolicyNotSpecified      = ResourceAccessPolicy("NotSpecified")
	ResourceAccessPolicyAcisReadAllowed   = ResourceAccessPolicy("AcisReadAllowed")
	ResourceAccessPolicyAcisActionAllowed = ResourceAccessPolicy("AcisActionAllowed")
)

type ResourceDeletionPolicy string

const (
	ResourceDeletionPolicyNotSpecified                   = ResourceDeletionPolicy("NotSpecified")
	ResourceDeletionPolicyCascadeDeleteAll               = ResourceDeletionPolicy("CascadeDeleteAll")
	ResourceDeletionPolicyCascadeDeleteProxyOnlyChildren = ResourceDeletionPolicy("CascadeDeleteProxyOnlyChildren")
)

type ResourceProviderCapabilitiesEffect string

const (
	ResourceProviderCapabilitiesEffectNotSpecified = ResourceProviderCapabilitiesEffect("NotSpecified")
	ResourceProviderCapabilitiesEffectAllow        = ResourceProviderCapabilitiesEffect("Allow")
	ResourceProviderCapabilitiesEffectDisallow     = ResourceProviderCapabilitiesEffect("Disallow")
)

type ResourceProviderType string

const (
	ResourceProviderTypeNotSpecified               = ResourceProviderType("NotSpecified")
	ResourceProviderTypeInternal                   = ResourceProviderType("Internal")
	ResourceProviderTypeExternal                   = ResourceProviderType("External")
	ResourceProviderTypeHidden                     = ResourceProviderType("Hidden")
	ResourceProviderTypeRegistrationFree           = ResourceProviderType("RegistrationFree")
	ResourceProviderTypeLegacyRegistrationRequired = ResourceProviderType("LegacyRegistrationRequired")
	ResourceProviderTypeTenantOnly                 = ResourceProviderType("TenantOnly")
	ResourceProviderTypeAuthorizationFree          = ResourceProviderType("AuthorizationFree")
)

type RoutingType string

const (
	RoutingTypeDefault          = RoutingType("Default")
	RoutingTypeProxyOnly        = RoutingType("ProxyOnly")
	RoutingTypeHostBased        = RoutingType("HostBased")
	RoutingTypeExtension        = RoutingType("Extension")
	RoutingTypeTenant           = RoutingType("Tenant")
	RoutingTypeFanout           = RoutingType("Fanout")
	RoutingTypeLocationBased    = RoutingType("LocationBased")
	RoutingTypeFailover         = RoutingType("Failover")
	RoutingTypeCascadeExtension = RoutingType("CascadeExtension")
)

type SkuScaleType string

const (
	SkuScaleTypeNone      = SkuScaleType("None")
	SkuScaleTypeManual    = SkuScaleType("Manual")
	SkuScaleTypeAutomatic = SkuScaleType("Automatic")
)

type SubscriptionNotificationOperation string

const (
	SubscriptionNotificationOperationNotDefined             = SubscriptionNotificationOperation("NotDefined")
	SubscriptionNotificationOperationDeleteAllResources     = SubscriptionNotificationOperation("DeleteAllResources")
	SubscriptionNotificationOperationSoftDeleteAllResources = SubscriptionNotificationOperation("SoftDeleteAllResources")
	SubscriptionNotificationOperationNoOp                   = SubscriptionNotificationOperation("NoOp")
	SubscriptionNotificationOperationBillingCancellation    = SubscriptionNotificationOperation("BillingCancellation")
	SubscriptionNotificationOperationUndoSoftDelete         = SubscriptionNotificationOperation("UndoSoftDelete")
)

type SubscriptionReregistrationResult string

const (
	SubscriptionReregistrationResultNotApplicable     = SubscriptionReregistrationResult("NotApplicable")
	SubscriptionReregistrationResultConditionalUpdate = SubscriptionReregistrationResult("ConditionalUpdate")
	SubscriptionReregistrationResultForcedUpdate      = SubscriptionReregistrationResult("ForcedUpdate")
	SubscriptionReregistrationResultFailed            = SubscriptionReregistrationResult("Failed")
)

type SubscriptionState string

const (
	SubscriptionStateNotDefined = SubscriptionState("NotDefined")
	SubscriptionStateEnabled    = SubscriptionState("Enabled")
	SubscriptionStateWarned     = SubscriptionState("Warned")
	SubscriptionStatePastDue    = SubscriptionState("PastDue")
	SubscriptionStateDisabled   = SubscriptionState("Disabled")
	SubscriptionStateDeleted    = SubscriptionState("Deleted")
)

type SubscriptionTransitioningState string

const (
	SubscriptionTransitioningStateRegistered              = SubscriptionTransitioningState("Registered")
	SubscriptionTransitioningStateUnregistered            = SubscriptionTransitioningState("Unregistered")
	SubscriptionTransitioningStateWarned                  = SubscriptionTransitioningState("Warned")
	SubscriptionTransitioningStateSuspended               = SubscriptionTransitioningState("Suspended")
	SubscriptionTransitioningStateDeleted                 = SubscriptionTransitioningState("Deleted")
	SubscriptionTransitioningStateWarnedToRegistered      = SubscriptionTransitioningState("WarnedToRegistered")
	SubscriptionTransitioningStateWarnedToSuspended       = SubscriptionTransitioningState("WarnedToSuspended")
	SubscriptionTransitioningStateWarnedToDeleted         = SubscriptionTransitioningState("WarnedToDeleted")
	SubscriptionTransitioningStateWarnedToUnregistered    = SubscriptionTransitioningState("WarnedToUnregistered")
	SubscriptionTransitioningStateSuspendedToRegistered   = SubscriptionTransitioningState("SuspendedToRegistered")
	SubscriptionTransitioningStateSuspendedToWarned       = SubscriptionTransitioningState("SuspendedToWarned")
	SubscriptionTransitioningStateSuspendedToDeleted      = SubscriptionTransitioningState("SuspendedToDeleted")
	SubscriptionTransitioningStateSuspendedToUnregistered = SubscriptionTransitioningState("SuspendedToUnregistered")
)

type ThrottlingMetricType string

const (
	ThrottlingMetricTypeNotSpecified      = ThrottlingMetricType("NotSpecified")
	ThrottlingMetricTypeNumberOfRequests  = ThrottlingMetricType("NumberOfRequests")
	ThrottlingMetricTypeNumberOfResources = ThrottlingMetricType("NumberOfResources")
)

type TrafficRegionCategory string

const (
	TrafficRegionCategoryNotSpecified           = TrafficRegionCategory("NotSpecified")
	TrafficRegionCategoryCanary                 = TrafficRegionCategory("Canary")
	TrafficRegionCategoryLowTraffic             = TrafficRegionCategory("LowTraffic")
	TrafficRegionCategoryMediumTraffic          = TrafficRegionCategory("MediumTraffic")
	TrafficRegionCategoryHighTraffic            = TrafficRegionCategory("HighTraffic")
	TrafficRegionCategoryNone                   = TrafficRegionCategory("None")
	TrafficRegionCategoryRestOfTheWorldGroupOne = TrafficRegionCategory("RestOfTheWorldGroupOne")
	TrafficRegionCategoryRestOfTheWorldGroupTwo = TrafficRegionCategory("RestOfTheWorldGroupTwo")
)

func init() {
}
