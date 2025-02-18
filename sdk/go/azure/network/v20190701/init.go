


package v20190701

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
	case "azure-native:network/v20190701:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20190701:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20190701:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20190701:BastionHost":
		r = &BastionHost{}
	case "azure-native:network/v20190701:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20190701:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20190701:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20190701:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20190701:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20190701:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20190701:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20190701:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20190701:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20190701:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20190701:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20190701:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network/v20190701:FirewallPolicyRuleGroup":
		r = &FirewallPolicyRuleGroup{}
	case "azure-native:network/v20190701:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20190701:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20190701:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20190701:NatGateway":
		r = &NatGateway{}
	case "azure-native:network/v20190701:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20190701:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20190701:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20190701:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20190701:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20190701:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20190701:P2sVpnServerConfiguration":
		r = &P2sVpnServerConfiguration{}
	case "azure-native:network/v20190701:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20190701:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network/v20190701:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network/v20190701:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20190701:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20190701:Route":
		r = &Route{}
	case "azure-native:network/v20190701:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20190701:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20190701:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20190701:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20190701:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20190701:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20190701:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20190701:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20190701:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20190701:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20190701:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20190701:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20190701:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20190701:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network/v20190701:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network/v20190701:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20190701:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20190701:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20190701:VpnSite":
		r = &VpnSite{}
	case "azure-native:network/v20190701:WebApplicationFirewallPolicy":
		r = &WebApplicationFirewallPolicy{}
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
		"network/v20190701",
		&module{version},
	)
}
