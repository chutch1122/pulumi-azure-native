


package v20150615

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
	case "azure-native:network/v20150615:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20150615:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20150615:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20150615:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20150615:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20150615:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20150615:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20150615:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20150615:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20150615:Route":
		r = &Route{}
	case "azure-native:network/v20150615:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20150615:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20150615:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20150615:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20150615:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20150615:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
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
		"network/v20150615",
		&module{version},
	)
}
