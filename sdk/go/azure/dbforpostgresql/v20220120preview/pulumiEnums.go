


package v20220120preview

type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModeCreate             = CreateMode("Create")
	CreateModeUpdate             = CreateMode("Update")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
)

type GeoRedundantBackupEnum string

const (
	GeoRedundantBackupEnumEnabled  = GeoRedundantBackupEnum("Enabled")
	GeoRedundantBackupEnumDisabled = GeoRedundantBackupEnum("Disabled")
)

type HighAvailabilityMode string

const (
	HighAvailabilityModeDisabled      = HighAvailabilityMode("Disabled")
	HighAvailabilityModeZoneRedundant = HighAvailabilityMode("ZoneRedundant")
	HighAvailabilityModeSameZone      = HighAvailabilityMode("SameZone")
)

type ServerVersion string

const (
	ServerVersion_14 = ServerVersion("14")
	ServerVersion_13 = ServerVersion("13")
	ServerVersion_12 = ServerVersion("12")
	ServerVersion_11 = ServerVersion("11")
)

type SkuTier string

const (
	SkuTierBurstable       = SkuTier("Burstable")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
)

func init() {
}
