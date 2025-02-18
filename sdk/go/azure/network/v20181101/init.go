


package v20181101

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
	case "azure-native:network/v20181101:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20181101:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20181101:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20181101:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20181101:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20181101:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20181101:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20181101:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20181101:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20181101:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20181101:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20181101:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20181101:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20181101:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20181101:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20181101:InterfaceEndpoint":
		r = &InterfaceEndpoint{}
	case "azure-native:network/v20181101:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20181101:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20181101:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20181101:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20181101:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20181101:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20181101:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20181101:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20181101:P2sVpnServerConfiguration":
		r = &P2sVpnServerConfiguration{}
	case "azure-native:network/v20181101:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20181101:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20181101:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20181101:Route":
		r = &Route{}
	case "azure-native:network/v20181101:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20181101:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20181101:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20181101:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20181101:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20181101:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20181101:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20181101:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20181101:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20181101:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20181101:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20181101:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20181101:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20181101:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20181101:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20181101:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20181101:VpnSite":
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
		"network/v20181101",
		&module{version},
	)
}
