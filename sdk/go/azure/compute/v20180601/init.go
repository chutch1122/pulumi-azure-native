


package v20180601

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
	case "azure-native:compute/v20180601:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:compute/v20180601:Disk":
		r = &Disk{}
	case "azure-native:compute/v20180601:Gallery":
		r = &Gallery{}
	case "azure-native:compute/v20180601:GalleryImage":
		r = &GalleryImage{}
	case "azure-native:compute/v20180601:GalleryImageVersion":
		r = &GalleryImageVersion{}
	case "azure-native:compute/v20180601:Image":
		r = &Image{}
	case "azure-native:compute/v20180601:ProximityPlacementGroup":
		r = &ProximityPlacementGroup{}
	case "azure-native:compute/v20180601:Snapshot":
		r = &Snapshot{}
	case "azure-native:compute/v20180601:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:compute/v20180601:VirtualMachineExtension":
		r = &VirtualMachineExtension{}
	case "azure-native:compute/v20180601:VirtualMachineScaleSet":
		r = &VirtualMachineScaleSet{}
	case "azure-native:compute/v20180601:VirtualMachineScaleSetExtension":
		r = &VirtualMachineScaleSetExtension{}
	case "azure-native:compute/v20180601:VirtualMachineScaleSetVM":
		r = &VirtualMachineScaleSetVM{}
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
		"compute/v20180601",
		&module{version},
	)
}
