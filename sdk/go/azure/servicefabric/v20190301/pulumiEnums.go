


package v20190301

type ArmServicePackageActivationMode string

const (
	// Indicates the application package activation mode will use shared process.
	ArmServicePackageActivationModeSharedProcess = ArmServicePackageActivationMode("SharedProcess")
	// Indicates the application package activation mode will use exclusive process.
	ArmServicePackageActivationModeExclusiveProcess = ArmServicePackageActivationMode("ExclusiveProcess")
)

type ArmUpgradeFailureAction string

const (
	// Indicates that a rollback of the upgrade will be performed by Service Fabric if the upgrade fails.
	ArmUpgradeFailureActionRollback = ArmUpgradeFailureAction("Rollback")
	// Indicates that a manual repair will need to be performed by the administrator if the upgrade fails. Service Fabric will not proceed to the next upgrade domain automatically.
	ArmUpgradeFailureActionManual = ArmUpgradeFailureAction("Manual")
)

type MoveCost string

const (
	// Zero move cost. This value is zero.
	MoveCostZero = MoveCost("Zero")
	// Specifies the move cost of the service as Low. The value is 1.
	MoveCostLow = MoveCost("Low")
	// Specifies the move cost of the service as Medium. The value is 2.
	MoveCostMedium = MoveCost("Medium")
	// Specifies the move cost of the service as High. The value is 3.
	MoveCostHigh = MoveCost("High")
)

type PartitionScheme string

const (
	// Indicates the partition kind is invalid. All Service Fabric enumerations have the invalid type. The value is zero.
	PartitionSchemeInvalid = PartitionScheme("Invalid")
	// Indicates that the partition is based on string names, and is a SingletonPartitionSchemeDescription object, The value is 1.
	PartitionSchemeSingleton = PartitionScheme("Singleton")
	// Indicates that the partition is based on Int64 key ranges, and is a UniformInt64RangePartitionSchemeDescription object. The value is 2.
	PartitionSchemeUniformInt64Range = PartitionScheme("UniformInt64Range")
	// Indicates that the partition is based on string names, and is a NamedPartitionSchemeDescription object. The value is 3
	PartitionSchemeNamed = PartitionScheme("Named")
)

type ServiceCorrelationScheme string

const (
	// An invalid correlation scheme. Cannot be used. The value is zero.
	ServiceCorrelationSchemeInvalid = ServiceCorrelationScheme("Invalid")
	// Indicates that this service has an affinity relationship with another service. Provided for backwards compatibility, consider preferring the Aligned or NonAlignedAffinity options. The value is 1.
	ServiceCorrelationSchemeAffinity = ServiceCorrelationScheme("Affinity")
	// Aligned affinity ensures that the primaries of the partitions of the affinitized services are collocated on the same nodes. This is the default and is the same as selecting the Affinity scheme. The value is 2.
	ServiceCorrelationSchemeAlignedAffinity = ServiceCorrelationScheme("AlignedAffinity")
	// Non-Aligned affinity guarantees that all replicas of each service will be placed on the same nodes. Unlike Aligned Affinity, this does not guarantee that replicas of particular role will be collocated. The value is 3.
	ServiceCorrelationSchemeNonAlignedAffinity = ServiceCorrelationScheme("NonAlignedAffinity")
)

type ServiceKind string

const (
	// Indicates the service kind is invalid. All Service Fabric enumerations have the invalid type. The value is zero.
	ServiceKindInvalid = ServiceKind("Invalid")
	// Does not use Service Fabric to make its state highly available or reliable. The value is 1.
	ServiceKindStateless = ServiceKind("Stateless")
	// Uses Service Fabric to make its state or part of its state highly available and reliable. The value is 2.
	ServiceKindStateful = ServiceKind("Stateful")
)

type ServiceLoadMetricWeight string

const (
	// Disables resource balancing for this metric. This value is zero.
	ServiceLoadMetricWeightZero = ServiceLoadMetricWeight("Zero")
	// Specifies the metric weight of the service load as Low. The value is 1.
	ServiceLoadMetricWeightLow = ServiceLoadMetricWeight("Low")
	// Specifies the metric weight of the service load as Medium. The value is 2.
	ServiceLoadMetricWeightMedium = ServiceLoadMetricWeight("Medium")
	// Specifies the metric weight of the service load as High. The value is 3.
	ServiceLoadMetricWeightHigh = ServiceLoadMetricWeight("High")
)

type ServicePlacementPolicyType string

const (
	// Indicates the type of the placement policy is invalid. All Service Fabric enumerations have the invalid type. The value is zero.
	ServicePlacementPolicyTypeInvalid = ServicePlacementPolicyType("Invalid")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementInvalidDomainPolicyDescription, which indicates that a particular fault or upgrade domain cannot be used for placement of this service. The value is 1.
	ServicePlacementPolicyTypeInvalidDomain = ServicePlacementPolicyType("InvalidDomain")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementRequireDomainDistributionPolicyDescription indicating that the replicas of the service must be placed in a specific domain. The value is 2.
	ServicePlacementPolicyTypeRequiredDomain = ServicePlacementPolicyType("RequiredDomain")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementPreferPrimaryDomainPolicyDescription, which indicates that if possible the Primary replica for the partitions of the service should be located in a particular domain as an optimization. The value is 3.
	ServicePlacementPolicyTypePreferredPrimaryDomain = ServicePlacementPolicyType("PreferredPrimaryDomain")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementRequireDomainDistributionPolicyDescription, indicating that the system will disallow placement of any two replicas from the same partition in the same domain at any time. The value is 4.
	ServicePlacementPolicyTypeRequiredDomainDistribution = ServicePlacementPolicyType("RequiredDomainDistribution")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementNonPartiallyPlaceServicePolicyDescription, which indicates that if possible all replicas of a particular partition of the service should be placed atomically. The value is 5.
	ServicePlacementPolicyTypeNonPartiallyPlaceService = ServicePlacementPolicyType("NonPartiallyPlaceService")
)

func init() {
}
