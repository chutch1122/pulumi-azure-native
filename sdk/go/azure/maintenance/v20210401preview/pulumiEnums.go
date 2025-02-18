


package v20210401preview

type MaintenanceScope string

const (
	// This maintenance scope controls installation of azure platform updates i.e. services on physical nodes hosting customer VMs.
	MaintenanceScopeHost = MaintenanceScope("Host")
	// This maintenance scope controls os image installation on VM/VMSS
	MaintenanceScopeOSImage = MaintenanceScope("OSImage")
	// This maintenance scope controls extension installation on VM/VMSS
	MaintenanceScopeExtension = MaintenanceScope("Extension")
	// This maintenance scope controls installation of windows and linux packages on VM/VMSS
	MaintenanceScopeInGuestPatch = MaintenanceScope("InGuestPatch")
	// This maintenance scope controls installation of SQL server platform updates.
	MaintenanceScopeSQLDB = MaintenanceScope("SQLDB")
	// This maintenance scope controls installation of SQL managed instance platform update.
	MaintenanceScopeSQLManagedInstance = MaintenanceScope("SQLManagedInstance")
)

type RebootOptions string

const (
	RebootOptionsNeverReboot      = RebootOptions("NeverReboot")
	RebootOptionsRebootIfRequired = RebootOptions("RebootIfRequired")
	RebootOptionsAlwaysReboot     = RebootOptions("AlwaysReboot")
)

type TaskScope string

const (
	TaskScopeGlobal   = TaskScope("Global")
	TaskScopeResource = TaskScope("Resource")
)

type Visibility string

const (
	// Only visible to users with permissions.
	VisibilityCustom = Visibility("Custom")
	// Visible to all users.
	VisibilityPublic = Visibility("Public")
)

func init() {
}
