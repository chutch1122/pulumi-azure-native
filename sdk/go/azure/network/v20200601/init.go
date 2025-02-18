


package v20200601

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
	case "azure-native:network/v20200601:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20200601:ApplicationGatewayPrivateEndpointConnection":
		r = &ApplicationGatewayPrivateEndpointConnection{}
	case "azure-native:network/v20200601:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20200601:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20200601:BastionHost":
		r = &BastionHost{}
	case "azure-native:network/v20200601:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20200601:CustomIPPrefix":
		r = &CustomIPPrefix{}
	case "azure-native:network/v20200601:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20200601:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20200601:DscpConfiguration":
		r = &DscpConfiguration{}
	case "azure-native:network/v20200601:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20200601:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20200601:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20200601:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20200601:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20200601:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20200601:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20200601:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20200601:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network/v20200601:FirewallPolicyRuleCollectionGroup":
		r = &FirewallPolicyRuleCollectionGroup{}
	case "azure-native:network/v20200601:FlowLog":
		r = &FlowLog{}
	case "azure-native:network/v20200601:HubRouteTable":
		r = &HubRouteTable{}
	case "azure-native:network/v20200601:HubVirtualNetworkConnection":
		r = &HubVirtualNetworkConnection{}
	case "azure-native:network/v20200601:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20200601:IpAllocation":
		r = &IpAllocation{}
	case "azure-native:network/v20200601:IpGroup":
		r = &IpGroup{}
	case "azure-native:network/v20200601:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20200601:LoadBalancerBackendAddressPool":
		r = &LoadBalancerBackendAddressPool{}
	case "azure-native:network/v20200601:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20200601:NatGateway":
		r = &NatGateway{}
	case "azure-native:network/v20200601:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20200601:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20200601:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20200601:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20200601:NetworkVirtualAppliance":
		r = &NetworkVirtualAppliance{}
	case "azure-native:network/v20200601:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20200601:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20200601:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20200601:PrivateDnsZoneGroup":
		r = &PrivateDnsZoneGroup{}
	case "azure-native:network/v20200601:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network/v20200601:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network/v20200601:PrivateLinkServicePrivateEndpointConnection":
		r = &PrivateLinkServicePrivateEndpointConnection{}
	case "azure-native:network/v20200601:PrivateRecordSet":
		r = &PrivateRecordSet{}
	case "azure-native:network/v20200601:PrivateZone":
		r = &PrivateZone{}
	case "azure-native:network/v20200601:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20200601:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20200601:Route":
		r = &Route{}
	case "azure-native:network/v20200601:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20200601:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20200601:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20200601:SecurityPartnerProvider":
		r = &SecurityPartnerProvider{}
	case "azure-native:network/v20200601:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20200601:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20200601:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20200601:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20200601:VirtualApplianceSite":
		r = &VirtualApplianceSite{}
	case "azure-native:network/v20200601:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20200601:VirtualHubBgpConnection":
		r = &VirtualHubBgpConnection{}
	case "azure-native:network/v20200601:VirtualHubIpConfiguration":
		r = &VirtualHubIpConfiguration{}
	case "azure-native:network/v20200601:VirtualHubRouteTableV2":
		r = &VirtualHubRouteTableV2{}
	case "azure-native:network/v20200601:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20200601:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20200601:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20200601:VirtualNetworkLink":
		r = &VirtualNetworkLink{}
	case "azure-native:network/v20200601:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20200601:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20200601:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network/v20200601:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network/v20200601:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20200601:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20200601:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20200601:VpnServerConfiguration":
		r = &VpnServerConfiguration{}
	case "azure-native:network/v20200601:VpnSite":
		r = &VpnSite{}
	case "azure-native:network/v20200601:WebApplicationFirewallPolicy":
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
		"network/v20200601",
		&module{version},
	)
}
