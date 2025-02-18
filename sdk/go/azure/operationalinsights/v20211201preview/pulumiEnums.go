


package v20211201preview

type ColumnDataTypeHintEnum string

const (
	// A string that matches the pattern of a URI, for example, scheme://username:password@host:1234/this/is/a/path?k1=v1&k2=v2#fragment
	ColumnDataTypeHintEnumUri = ColumnDataTypeHintEnum("uri")
	// A standard 128-bit GUID following the standard shape, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ColumnDataTypeHintEnumGuid = ColumnDataTypeHintEnum("guid")
	// An Azure Resource Model (ARM) path: /subscriptions/{...}/resourceGroups/{...}/providers/Microsoft.{...}/{...}/{...}/{...}...
	ColumnDataTypeHintEnumArmPath = ColumnDataTypeHintEnum("armPath")
	// A standard V4/V6 ip address following the standard shape, x.x.x.x/y:y:y:y:y:y:y:y
	ColumnDataTypeHintEnumIp = ColumnDataTypeHintEnum("ip")
)

type ColumnTypeEnum string

const (
	ColumnTypeEnumString   = ColumnTypeEnum("string")
	ColumnTypeEnumInt      = ColumnTypeEnum("int")
	ColumnTypeEnumLong     = ColumnTypeEnum("long")
	ColumnTypeEnumReal     = ColumnTypeEnum("real")
	ColumnTypeEnumBoolean  = ColumnTypeEnum("boolean")
	ColumnTypeEnumDateTime = ColumnTypeEnum("dateTime")
	ColumnTypeEnumGuid     = ColumnTypeEnum("guid")
	ColumnTypeEnumDynamic  = ColumnTypeEnum("dynamic")
)

type PublicNetworkAccessType string

const (
	// Enables connectivity to Log Analytics through public DNS.
	PublicNetworkAccessTypeEnabled = PublicNetworkAccessType("Enabled")
	// Disables public connectivity to Log Analytics through public DNS.
	PublicNetworkAccessTypeDisabled = PublicNetworkAccessType("Disabled")
)

type TablePlanEnum string

const (
	// Logs  that are adjusted to support high volume low value verbose logs.
	TablePlanEnumBasic = TablePlanEnum("Basic")
	// Logs  that allow monitoring and analytics.
	TablePlanEnumAnalytics = TablePlanEnum("Analytics")
)

type WorkspaceSkuNameEnum string

const (
	WorkspaceSkuNameEnumFree                = WorkspaceSkuNameEnum("Free")
	WorkspaceSkuNameEnumStandard            = WorkspaceSkuNameEnum("Standard")
	WorkspaceSkuNameEnumPremium             = WorkspaceSkuNameEnum("Premium")
	WorkspaceSkuNameEnumPerNode             = WorkspaceSkuNameEnum("PerNode")
	WorkspaceSkuNameEnumPerGB2018           = WorkspaceSkuNameEnum("PerGB2018")
	WorkspaceSkuNameEnumStandalone          = WorkspaceSkuNameEnum("Standalone")
	WorkspaceSkuNameEnumCapacityReservation = WorkspaceSkuNameEnum("CapacityReservation")
	WorkspaceSkuNameEnumLACluster           = WorkspaceSkuNameEnum("LACluster")
)

func init() {
}
