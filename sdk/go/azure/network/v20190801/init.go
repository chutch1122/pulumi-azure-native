


package v20190801

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
	case "azure-native:network/v20190801:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20190801:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20190801:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20190801:BastionHost":
		r = &BastionHost{}
	case "azure-native:network/v20190801:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20190801:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20190801:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20190801:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20190801:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20190801:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20190801:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20190801:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20190801:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20190801:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20190801:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20190801:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network/v20190801:FirewallPolicyRuleGroup":
		r = &FirewallPolicyRuleGroup{}
	case "azure-native:network/v20190801:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20190801:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20190801:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20190801:NatGateway":
		r = &NatGateway{}
	case "azure-native:network/v20190801:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20190801:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20190801:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20190801:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20190801:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20190801:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20190801:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20190801:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network/v20190801:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network/v20190801:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20190801:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20190801:Route":
		r = &Route{}
	case "azure-native:network/v20190801:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20190801:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20190801:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20190801:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20190801:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20190801:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20190801:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20190801:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20190801:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20190801:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20190801:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20190801:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20190801:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20190801:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network/v20190801:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network/v20190801:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20190801:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20190801:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20190801:VpnServerConfiguration":
		r = &VpnServerConfiguration{}
	case "azure-native:network/v20190801:VpnSite":
		r = &VpnSite{}
	case "azure-native:network/v20190801:WebApplicationFirewallPolicy":
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
		"network/v20190801",
		&module{version},
	)
}
