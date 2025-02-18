


package v20191201

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
	case "azure-native:compute/v20191201:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:compute/v20191201:DedicatedHost":
		r = &DedicatedHost{}
	case "azure-native:compute/v20191201:DedicatedHostGroup":
		r = &DedicatedHostGroup{}
	case "azure-native:compute/v20191201:Gallery":
		r = &Gallery{}
	case "azure-native:compute/v20191201:GalleryApplication":
		r = &GalleryApplication{}
	case "azure-native:compute/v20191201:GalleryApplicationVersion":
		r = &GalleryApplicationVersion{}
	case "azure-native:compute/v20191201:GalleryImage":
		r = &GalleryImage{}
	case "azure-native:compute/v20191201:GalleryImageVersion":
		r = &GalleryImageVersion{}
	case "azure-native:compute/v20191201:Image":
		r = &Image{}
	case "azure-native:compute/v20191201:ProximityPlacementGroup":
		r = &ProximityPlacementGroup{}
	case "azure-native:compute/v20191201:SshPublicKey":
		r = &SshPublicKey{}
	case "azure-native:compute/v20191201:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:compute/v20191201:VirtualMachineExtension":
		r = &VirtualMachineExtension{}
	case "azure-native:compute/v20191201:VirtualMachineScaleSet":
		r = &VirtualMachineScaleSet{}
	case "azure-native:compute/v20191201:VirtualMachineScaleSetExtension":
		r = &VirtualMachineScaleSetExtension{}
	case "azure-native:compute/v20191201:VirtualMachineScaleSetVM":
		r = &VirtualMachineScaleSetVM{}
	case "azure-native:compute/v20191201:VirtualMachineScaleSetVMExtension":
		r = &VirtualMachineScaleSetVMExtension{}
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
		"compute/v20191201",
		&module{version},
	)
}
