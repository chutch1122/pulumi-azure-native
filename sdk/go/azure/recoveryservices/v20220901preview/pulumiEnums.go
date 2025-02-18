


package v20220901preview

type AcquireStorageAccountLock string

const (
	AcquireStorageAccountLockAcquire    = AcquireStorageAccountLock("Acquire")
	AcquireStorageAccountLockNotAcquire = AcquireStorageAccountLock("NotAcquire")
)

type BackupItemType string

const (
	BackupItemTypeInvalid           = BackupItemType("Invalid")
	BackupItemTypeVM                = BackupItemType("VM")
	BackupItemTypeFileFolder        = BackupItemType("FileFolder")
	BackupItemTypeAzureSqlDb        = BackupItemType("AzureSqlDb")
	BackupItemTypeSQLDB             = BackupItemType("SQLDB")
	BackupItemTypeExchange          = BackupItemType("Exchange")
	BackupItemTypeSharepoint        = BackupItemType("Sharepoint")
	BackupItemTypeVMwareVM          = BackupItemType("VMwareVM")
	BackupItemTypeSystemState       = BackupItemType("SystemState")
	BackupItemTypeClient            = BackupItemType("Client")
	BackupItemTypeGenericDataSource = BackupItemType("GenericDataSource")
	BackupItemTypeSQLDataBase       = BackupItemType("SQLDataBase")
	BackupItemTypeAzureFileShare    = BackupItemType("AzureFileShare")
	BackupItemTypeSAPHanaDatabase   = BackupItemType("SAPHanaDatabase")
	BackupItemTypeSAPAseDatabase    = BackupItemType("SAPAseDatabase")
	BackupItemTypeSAPHanaDBInstance = BackupItemType("SAPHanaDBInstance")
)

type BackupManagementType string

const (
	BackupManagementTypeInvalid           = BackupManagementType("Invalid")
	BackupManagementTypeAzureIaasVM       = BackupManagementType("AzureIaasVM")
	BackupManagementTypeMAB               = BackupManagementType("MAB")
	BackupManagementTypeDPM               = BackupManagementType("DPM")
	BackupManagementTypeAzureBackupServer = BackupManagementType("AzureBackupServer")
	BackupManagementTypeAzureSql          = BackupManagementType("AzureSql")
	BackupManagementTypeAzureStorage      = BackupManagementType("AzureStorage")
	BackupManagementTypeAzureWorkload     = BackupManagementType("AzureWorkload")
	BackupManagementTypeDefaultBackup     = BackupManagementType("DefaultBackup")
)

type CreateMode string

const (
	CreateModeInvalid = CreateMode("Invalid")
	CreateModeDefault = CreateMode("Default")
	CreateModeRecover = CreateMode("Recover")
)

type DayOfWeek string

const (
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
)

type IAASVMPolicyType string

const (
	IAASVMPolicyTypeInvalid = IAASVMPolicyType("Invalid")
	IAASVMPolicyTypeV1      = IAASVMPolicyType("V1")
	IAASVMPolicyTypeV2      = IAASVMPolicyType("V2")
)

type LastBackupStatus string

const (
	LastBackupStatusInvalid   = LastBackupStatus("Invalid")
	LastBackupStatusHealthy   = LastBackupStatus("Healthy")
	LastBackupStatusUnhealthy = LastBackupStatus("Unhealthy")
	LastBackupStatusIRPending = LastBackupStatus("IRPending")
)

type MonthOfYear string

const (
	MonthOfYearInvalid   = MonthOfYear("Invalid")
	MonthOfYearJanuary   = MonthOfYear("January")
	MonthOfYearFebruary  = MonthOfYear("February")
	MonthOfYearMarch     = MonthOfYear("March")
	MonthOfYearApril     = MonthOfYear("April")
	MonthOfYearMay       = MonthOfYear("May")
	MonthOfYearJune      = MonthOfYear("June")
	MonthOfYearJuly      = MonthOfYear("July")
	MonthOfYearAugust    = MonthOfYear("August")
	MonthOfYearSeptember = MonthOfYear("September")
	MonthOfYearOctober   = MonthOfYear("October")
	MonthOfYearNovember  = MonthOfYear("November")
	MonthOfYearDecember  = MonthOfYear("December")
)

type OperationType string

const (
	OperationTypeInvalid    = OperationType("Invalid")
	OperationTypeRegister   = OperationType("Register")
	OperationTypeReregister = OperationType("Reregister")
)

type PolicyType string

const (
	PolicyTypeInvalid              = PolicyType("Invalid")
	PolicyTypeFull                 = PolicyType("Full")
	PolicyTypeDifferential         = PolicyType("Differential")
	PolicyTypeLog                  = PolicyType("Log")
	PolicyTypeCopyOnlyFull         = PolicyType("CopyOnlyFull")
	PolicyTypeIncremental          = PolicyType("Incremental")
	PolicyTypeSnapshotFull         = PolicyType("SnapshotFull")
	PolicyTypeSnapshotCopyOnlyFull = PolicyType("SnapshotCopyOnlyFull")
)

type PrivateEndpointConnectionStatus string

const (
	PrivateEndpointConnectionStatusPending      = PrivateEndpointConnectionStatus("Pending")
	PrivateEndpointConnectionStatusApproved     = PrivateEndpointConnectionStatus("Approved")
	PrivateEndpointConnectionStatusRejected     = PrivateEndpointConnectionStatus("Rejected")
	PrivateEndpointConnectionStatusDisconnected = PrivateEndpointConnectionStatus("Disconnected")
)

type ProtectableContainerType string

const (
	ProtectableContainerTypeInvalid                                   = ProtectableContainerType("Invalid")
	ProtectableContainerTypeUnknown                                   = ProtectableContainerType("Unknown")
	ProtectableContainerTypeIaasVMContainer                           = ProtectableContainerType("IaasVMContainer")
	ProtectableContainerTypeIaasVMServiceContainer                    = ProtectableContainerType("IaasVMServiceContainer")
	ProtectableContainerTypeDPMContainer                              = ProtectableContainerType("DPMContainer")
	ProtectableContainerTypeAzureBackupServerContainer                = ProtectableContainerType("AzureBackupServerContainer")
	ProtectableContainerTypeMABContainer                              = ProtectableContainerType("MABContainer")
	ProtectableContainerTypeCluster                                   = ProtectableContainerType("Cluster")
	ProtectableContainerTypeAzureSqlContainer                         = ProtectableContainerType("AzureSqlContainer")
	ProtectableContainerTypeWindows                                   = ProtectableContainerType("Windows")
	ProtectableContainerTypeVCenter                                   = ProtectableContainerType("VCenter")
	ProtectableContainerTypeVMAppContainer                            = ProtectableContainerType("VMAppContainer")
	ProtectableContainerTypeSQLAGWorkLoadContainer                    = ProtectableContainerType("SQLAGWorkLoadContainer")
	ProtectableContainerTypeStorageContainer                          = ProtectableContainerType("StorageContainer")
	ProtectableContainerTypeGenericContainer                          = ProtectableContainerType("GenericContainer")
	ProtectableContainerType_Microsoft_ClassicCompute_virtualMachines = ProtectableContainerType("Microsoft.ClassicCompute/virtualMachines")
	ProtectableContainerType_Microsoft_Compute_virtualMachines        = ProtectableContainerType("Microsoft.Compute/virtualMachines")
	ProtectableContainerTypeAzureWorkloadContainer                    = ProtectableContainerType("AzureWorkloadContainer")
)

type ProtectedItemHealthStatus string

const (
	ProtectedItemHealthStatusInvalid      = ProtectedItemHealthStatus("Invalid")
	ProtectedItemHealthStatusHealthy      = ProtectedItemHealthStatus("Healthy")
	ProtectedItemHealthStatusUnhealthy    = ProtectedItemHealthStatus("Unhealthy")
	ProtectedItemHealthStatusNotReachable = ProtectedItemHealthStatus("NotReachable")
	ProtectedItemHealthStatusIRPending    = ProtectedItemHealthStatus("IRPending")
)

type ProtectedItemStateEnum string

const (
	ProtectedItemStateEnumInvalid           = ProtectedItemStateEnum("Invalid")
	ProtectedItemStateEnumIRPending         = ProtectedItemStateEnum("IRPending")
	ProtectedItemStateEnumProtected         = ProtectedItemStateEnum("Protected")
	ProtectedItemStateEnumProtectionError   = ProtectedItemStateEnum("ProtectionError")
	ProtectedItemStateEnumProtectionStopped = ProtectedItemStateEnum("ProtectionStopped")
	ProtectedItemStateEnumProtectionPaused  = ProtectedItemStateEnum("ProtectionPaused")
)

type ProtectionIntentItemType string

const (
	ProtectionIntentItemTypeInvalid                                    = ProtectionIntentItemType("Invalid")
	ProtectionIntentItemTypeAzureResourceItem                          = ProtectionIntentItemType("AzureResourceItem")
	ProtectionIntentItemTypeRecoveryServiceVaultItem                   = ProtectionIntentItemType("RecoveryServiceVaultItem")
	ProtectionIntentItemTypeAzureWorkloadContainerAutoProtectionIntent = ProtectionIntentItemType("AzureWorkloadContainerAutoProtectionIntent")
	ProtectionIntentItemTypeAzureWorkloadAutoProtectionIntent          = ProtectionIntentItemType("AzureWorkloadAutoProtectionIntent")
	ProtectionIntentItemTypeAzureWorkloadSQLAutoProtectionIntent       = ProtectionIntentItemType("AzureWorkloadSQLAutoProtectionIntent")
)

type ProtectionState string

const (
	ProtectionStateInvalid           = ProtectionState("Invalid")
	ProtectionStateIRPending         = ProtectionState("IRPending")
	ProtectionStateProtected         = ProtectionState("Protected")
	ProtectionStateProtectionError   = ProtectionState("ProtectionError")
	ProtectionStateProtectionStopped = ProtectionState("ProtectionStopped")
	ProtectionStateProtectionPaused  = ProtectionState("ProtectionPaused")
)

type ProtectionStatus string

const (
	ProtectionStatusInvalid          = ProtectionStatus("Invalid")
	ProtectionStatusNotProtected     = ProtectionStatus("NotProtected")
	ProtectionStatusProtecting       = ProtectionStatus("Protecting")
	ProtectionStatusProtected        = ProtectionStatus("Protected")
	ProtectionStatusProtectionFailed = ProtectionStatus("ProtectionFailed")
)

type ProvisioningState string

const (
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
	ProvisioningStateDeleting  = ProvisioningState("Deleting")
	ProvisioningStateFailed    = ProvisioningState("Failed")
	ProvisioningStatePending   = ProvisioningState("Pending")
)

type ResourceHealthStatus string

const (
	ResourceHealthStatusHealthy             = ResourceHealthStatus("Healthy")
	ResourceHealthStatusTransientDegraded   = ResourceHealthStatus("TransientDegraded")
	ResourceHealthStatusPersistentDegraded  = ResourceHealthStatus("PersistentDegraded")
	ResourceHealthStatusTransientUnhealthy  = ResourceHealthStatus("TransientUnhealthy")
	ResourceHealthStatusPersistentUnhealthy = ResourceHealthStatus("PersistentUnhealthy")
	ResourceHealthStatusInvalid             = ResourceHealthStatus("Invalid")
)

type RetentionDurationType string

const (
	RetentionDurationTypeInvalid = RetentionDurationType("Invalid")
	RetentionDurationTypeDays    = RetentionDurationType("Days")
	RetentionDurationTypeWeeks   = RetentionDurationType("Weeks")
	RetentionDurationTypeMonths  = RetentionDurationType("Months")
	RetentionDurationTypeYears   = RetentionDurationType("Years")
)

type RetentionScheduleFormat string

const (
	RetentionScheduleFormatInvalid = RetentionScheduleFormat("Invalid")
	RetentionScheduleFormatDaily   = RetentionScheduleFormat("Daily")
	RetentionScheduleFormatWeekly  = RetentionScheduleFormat("Weekly")
)

type ScheduleRunType string

const (
	ScheduleRunTypeInvalid = ScheduleRunType("Invalid")
	ScheduleRunTypeDaily   = ScheduleRunType("Daily")
	ScheduleRunTypeWeekly  = ScheduleRunType("Weekly")
	ScheduleRunTypeHourly  = ScheduleRunType("Hourly")
)

type TieringMode string

const (
	TieringModeInvalid         = TieringMode("Invalid")
	TieringModeTierRecommended = TieringMode("TierRecommended")
	TieringModeTierAfter       = TieringMode("TierAfter")
	TieringModeDoNotTier       = TieringMode("DoNotTier")
)

type WeekOfMonth string

const (
	WeekOfMonthFirst   = WeekOfMonth("First")
	WeekOfMonthSecond  = WeekOfMonth("Second")
	WeekOfMonthThird   = WeekOfMonth("Third")
	WeekOfMonthFourth  = WeekOfMonth("Fourth")
	WeekOfMonthLast    = WeekOfMonth("Last")
	WeekOfMonthInvalid = WeekOfMonth("Invalid")
)

type WorkloadItemType string

const (
	WorkloadItemTypeInvalid           = WorkloadItemType("Invalid")
	WorkloadItemTypeSQLInstance       = WorkloadItemType("SQLInstance")
	WorkloadItemTypeSQLDataBase       = WorkloadItemType("SQLDataBase")
	WorkloadItemTypeSAPHanaSystem     = WorkloadItemType("SAPHanaSystem")
	WorkloadItemTypeSAPHanaDatabase   = WorkloadItemType("SAPHanaDatabase")
	WorkloadItemTypeSAPAseSystem      = WorkloadItemType("SAPAseSystem")
	WorkloadItemTypeSAPAseDatabase    = WorkloadItemType("SAPAseDatabase")
	WorkloadItemTypeSAPHanaDBInstance = WorkloadItemType("SAPHanaDBInstance")
)

type WorkloadType string

const (
	WorkloadTypeInvalid           = WorkloadType("Invalid")
	WorkloadTypeVM                = WorkloadType("VM")
	WorkloadTypeFileFolder        = WorkloadType("FileFolder")
	WorkloadTypeAzureSqlDb        = WorkloadType("AzureSqlDb")
	WorkloadTypeSQLDB             = WorkloadType("SQLDB")
	WorkloadTypeExchange          = WorkloadType("Exchange")
	WorkloadTypeSharepoint        = WorkloadType("Sharepoint")
	WorkloadTypeVMwareVM          = WorkloadType("VMwareVM")
	WorkloadTypeSystemState       = WorkloadType("SystemState")
	WorkloadTypeClient            = WorkloadType("Client")
	WorkloadTypeGenericDataSource = WorkloadType("GenericDataSource")
	WorkloadTypeSQLDataBase       = WorkloadType("SQLDataBase")
	WorkloadTypeAzureFileShare    = WorkloadType("AzureFileShare")
	WorkloadTypeSAPHanaDatabase   = WorkloadType("SAPHanaDatabase")
	WorkloadTypeSAPAseDatabase    = WorkloadType("SAPAseDatabase")
	WorkloadTypeSAPHanaDBInstance = WorkloadType("SAPHanaDBInstance")
)

func init() {
}
