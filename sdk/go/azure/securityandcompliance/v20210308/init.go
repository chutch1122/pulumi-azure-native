


package v20210308

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
	case "azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsAdtAPI":
		r = &PrivateEndpointConnectionsAdtAPI{}
	case "azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsComp":
		r = &PrivateEndpointConnectionsComp{}
	case "azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsForEDM":
		r = &PrivateEndpointConnectionsForEDM{}
	case "azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsForMIPPolicySync":
		r = &PrivateEndpointConnectionsForMIPPolicySync{}
	case "azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsForSCCPowershell":
		r = &PrivateEndpointConnectionsForSCCPowershell{}
	case "azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsSec":
		r = &PrivateEndpointConnectionsSec{}
	case "azure-native:securityandcompliance/v20210308:privateLinkServicesForEDMUpload":
		r = &PrivateLinkServicesForEDMUpload{}
	case "azure-native:securityandcompliance/v20210308:privateLinkServicesForM365ComplianceCenter":
		r = &PrivateLinkServicesForM365ComplianceCenter{}
	case "azure-native:securityandcompliance/v20210308:privateLinkServicesForM365SecurityCenter":
		r = &PrivateLinkServicesForM365SecurityCenter{}
	case "azure-native:securityandcompliance/v20210308:privateLinkServicesForMIPPolicySync":
		r = &PrivateLinkServicesForMIPPolicySync{}
	case "azure-native:securityandcompliance/v20210308:privateLinkServicesForO365ManagementActivityAPI":
		r = &PrivateLinkServicesForO365ManagementActivityAPI{}
	case "azure-native:securityandcompliance/v20210308:privateLinkServicesForSCCPowershell":
		r = &PrivateLinkServicesForSCCPowershell{}
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
		"securityandcompliance/v20210308",
		&module{version},
	)
}
