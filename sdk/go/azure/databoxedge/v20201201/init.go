


package v20201201

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:databoxedge/v20201201:Addon":
		r = &Addon{}
	case "azure-native:databoxedge/v20201201:ArcAddon":
		r = &ArcAddon{}
	case "azure-native:databoxedge/v20201201:BandwidthSchedule":
		r = &BandwidthSchedule{}
	case "azure-native:databoxedge/v20201201:CloudEdgeManagementRole":
		r = &CloudEdgeManagementRole{}
	case "azure-native:databoxedge/v20201201:Container":
		r = &Container{}
	case "azure-native:databoxedge/v20201201:Device":
		r = &Device{}
	case "azure-native:databoxedge/v20201201:FileEventTrigger":
		r = &FileEventTrigger{}
	case "azure-native:databoxedge/v20201201:IoTAddon":
		r = &IoTAddon{}
	case "azure-native:databoxedge/v20201201:IoTRole":
		r = &IoTRole{}
	case "azure-native:databoxedge/v20201201:KubernetesRole":
		r = &KubernetesRole{}
	case "azure-native:databoxedge/v20201201:MECRole":
		r = &MECRole{}
	case "azure-native:databoxedge/v20201201:MonitoringConfig":
		r = &MonitoringConfig{}
	case "azure-native:databoxedge/v20201201:Order":
		r = &Order{}
	case "azure-native:databoxedge/v20201201:PeriodicTimerEventTrigger":
		r = &PeriodicTimerEventTrigger{}
	case "azure-native:databoxedge/v20201201:Role":
		r = &Role{}
	case "azure-native:databoxedge/v20201201:Share":
		r = &Share{}
	case "azure-native:databoxedge/v20201201:StorageAccount":
		r = &StorageAccount{}
	case "azure-native:databoxedge/v20201201:StorageAccountCredential":
		r = &StorageAccountCredential{}
	case "azure-native:databoxedge/v20201201:Trigger":
		r = &Trigger{}
	case "azure-native:databoxedge/v20201201:User":
		r = &User{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"databoxedge/v20201201",
		&module{version},
	)
}
