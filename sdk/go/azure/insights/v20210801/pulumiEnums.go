


package v20210801

type ConditionOperator string

const (
	ConditionOperatorEquals             = ConditionOperator("Equals")
	ConditionOperatorGreaterThan        = ConditionOperator("GreaterThan")
	ConditionOperatorGreaterThanOrEqual = ConditionOperator("GreaterThanOrEqual")
	ConditionOperatorLessThan           = ConditionOperator("LessThan")
	ConditionOperatorLessThanOrEqual    = ConditionOperator("LessThanOrEqual")
)

type DimensionOperator string

const (
	DimensionOperatorInclude = DimensionOperator("Include")
	DimensionOperatorExclude = DimensionOperator("Exclude")
)

type Kind string

const (
	KindUser   = Kind("user")
	KindShared = Kind("shared")
)

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned,UserAssigned")
)

type TimeAggregation string

const (
	TimeAggregationCount   = TimeAggregation("Count")
	TimeAggregationAverage = TimeAggregation("Average")
	TimeAggregationMinimum = TimeAggregation("Minimum")
	TimeAggregationMaximum = TimeAggregation("Maximum")
	TimeAggregationTotal   = TimeAggregation("Total")
)

func init() {
}
