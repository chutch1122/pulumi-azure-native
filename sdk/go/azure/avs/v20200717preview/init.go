


package v20200717preview

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
	case "azure-native:avs/v20200717preview:Addon":
		r = &Addon{}
	case "azure-native:avs/v20200717preview:Authorization":
		r = &Authorization{}
	case "azure-native:avs/v20200717preview:Cluster":
		r = &Cluster{}
	case "azure-native:avs/v20200717preview:GlobalReachConnection":
		r = &GlobalReachConnection{}
	case "azure-native:avs/v20200717preview:HcxEnterpriseSite":
		r = &HcxEnterpriseSite{}
	case "azure-native:avs/v20200717preview:PrivateCloud":
		r = &PrivateCloud{}
	case "azure-native:avs/v20200717preview:WorkloadNetworkDhcp":
		r = &WorkloadNetworkDhcp{}
	case "azure-native:avs/v20200717preview:WorkloadNetworkDnsService":
		r = &WorkloadNetworkDnsService{}
	case "azure-native:avs/v20200717preview:WorkloadNetworkDnsZone":
		r = &WorkloadNetworkDnsZone{}
	case "azure-native:avs/v20200717preview:WorkloadNetworkPortMirroring":
		r = &WorkloadNetworkPortMirroring{}
	case "azure-native:avs/v20200717preview:WorkloadNetworkSegment":
		r = &WorkloadNetworkSegment{}
	case "azure-native:avs/v20200717preview:WorkloadNetworkVMGroup":
		r = &WorkloadNetworkVMGroup{}
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
		"avs/v20200717preview",
		&module{version},
	)
}
