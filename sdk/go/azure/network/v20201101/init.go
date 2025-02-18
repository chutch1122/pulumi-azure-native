


package v20201101

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
	case "azure-native:network/v20201101:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20201101:ApplicationGatewayPrivateEndpointConnection":
		r = &ApplicationGatewayPrivateEndpointConnection{}
	case "azure-native:network/v20201101:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20201101:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20201101:BastionHost":
		r = &BastionHost{}
	case "azure-native:network/v20201101:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20201101:CustomIPPrefix":
		r = &CustomIPPrefix{}
	case "azure-native:network/v20201101:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20201101:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20201101:DscpConfiguration":
		r = &DscpConfiguration{}
	case "azure-native:network/v20201101:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20201101:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20201101:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20201101:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20201101:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20201101:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20201101:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20201101:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20201101:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network/v20201101:FirewallPolicyRuleCollectionGroup":
		r = &FirewallPolicyRuleCollectionGroup{}
	case "azure-native:network/v20201101:FlowLog":
		r = &FlowLog{}
	case "azure-native:network/v20201101:HubRouteTable":
		r = &HubRouteTable{}
	case "azure-native:network/v20201101:HubVirtualNetworkConnection":
		r = &HubVirtualNetworkConnection{}
	case "azure-native:network/v20201101:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20201101:IpAllocation":
		r = &IpAllocation{}
	case "azure-native:network/v20201101:IpGroup":
		r = &IpGroup{}
	case "azure-native:network/v20201101:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20201101:LoadBalancerBackendAddressPool":
		r = &LoadBalancerBackendAddressPool{}
	case "azure-native:network/v20201101:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20201101:NatGateway":
		r = &NatGateway{}
	case "azure-native:network/v20201101:NatRule":
		r = &NatRule{}
	case "azure-native:network/v20201101:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20201101:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20201101:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20201101:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20201101:NetworkVirtualAppliance":
		r = &NetworkVirtualAppliance{}
	case "azure-native:network/v20201101:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20201101:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20201101:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20201101:Policy":
		r = &Policy{}
	case "azure-native:network/v20201101:PrivateDnsZoneGroup":
		r = &PrivateDnsZoneGroup{}
	case "azure-native:network/v20201101:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network/v20201101:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network/v20201101:PrivateLinkServicePrivateEndpointConnection":
		r = &PrivateLinkServicePrivateEndpointConnection{}
	case "azure-native:network/v20201101:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20201101:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20201101:Route":
		r = &Route{}
	case "azure-native:network/v20201101:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20201101:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20201101:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20201101:SecurityPartnerProvider":
		r = &SecurityPartnerProvider{}
	case "azure-native:network/v20201101:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20201101:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20201101:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20201101:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20201101:VirtualApplianceSite":
		r = &VirtualApplianceSite{}
	case "azure-native:network/v20201101:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20201101:VirtualHubBgpConnection":
		r = &VirtualHubBgpConnection{}
	case "azure-native:network/v20201101:VirtualHubIpConfiguration":
		r = &VirtualHubIpConfiguration{}
	case "azure-native:network/v20201101:VirtualHubRouteTableV2":
		r = &VirtualHubRouteTableV2{}
	case "azure-native:network/v20201101:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20201101:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20201101:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20201101:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20201101:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20201101:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network/v20201101:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network/v20201101:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20201101:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20201101:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20201101:VpnServerConfiguration":
		r = &VpnServerConfiguration{}
	case "azure-native:network/v20201101:VpnSite":
		r = &VpnSite{}
	case "azure-native:network/v20201101:WebApplicationFirewallPolicy":
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
		"network/v20201101",
		&module{version},
	)
}
