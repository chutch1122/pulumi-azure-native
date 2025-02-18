


package v20180401

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
	case "azure-native:network/v20180401:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20180401:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20180401:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20180401:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20180401:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20180401:Endpoint":
		r = &Endpoint{}
	case "azure-native:network/v20180401:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20180401:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20180401:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20180401:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20180401:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20180401:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20180401:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20180401:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20180401:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20180401:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20180401:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20180401:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20180401:Profile":
		r = &Profile{}
	case "azure-native:network/v20180401:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20180401:Route":
		r = &Route{}
	case "azure-native:network/v20180401:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20180401:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20180401:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20180401:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20180401:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20180401:TrafficManagerUserMetricsKey":
		r = &TrafficManagerUserMetricsKey{}
	case "azure-native:network/v20180401:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20180401:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20180401:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20180401:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20180401:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20180401:VirtualWAN":
		r = &VirtualWAN{}
	case "azure-native:network/v20180401:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20180401:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20180401:VpnSite":
		r = &VpnSite{}
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
		"network/v20180401",
		&module{version},
	)
}
