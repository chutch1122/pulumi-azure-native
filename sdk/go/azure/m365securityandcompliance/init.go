


package m365securityandcompliance

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
	case "azure-native:m365securityandcompliance:PrivateEndpointConnectionsAdtAPI":
		r = &PrivateEndpointConnectionsAdtAPI{}
	case "azure-native:m365securityandcompliance:PrivateEndpointConnectionsComp":
		r = &PrivateEndpointConnectionsComp{}
	case "azure-native:m365securityandcompliance:PrivateEndpointConnectionsForEDM":
		r = &PrivateEndpointConnectionsForEDM{}
	case "azure-native:m365securityandcompliance:PrivateEndpointConnectionsForMIPPolicySync":
		r = &PrivateEndpointConnectionsForMIPPolicySync{}
	case "azure-native:m365securityandcompliance:PrivateEndpointConnectionsForSCCPowershell":
		r = &PrivateEndpointConnectionsForSCCPowershell{}
	case "azure-native:m365securityandcompliance:PrivateEndpointConnectionsSec":
		r = &PrivateEndpointConnectionsSec{}
	case "azure-native:m365securityandcompliance:privateLinkServicesForEDMUpload":
		r = &PrivateLinkServicesForEDMUpload{}
	case "azure-native:m365securityandcompliance:privateLinkServicesForM365ComplianceCenter":
		r = &PrivateLinkServicesForM365ComplianceCenter{}
	case "azure-native:m365securityandcompliance:privateLinkServicesForM365SecurityCenter":
		r = &PrivateLinkServicesForM365SecurityCenter{}
	case "azure-native:m365securityandcompliance:privateLinkServicesForMIPPolicySync":
		r = &PrivateLinkServicesForMIPPolicySync{}
	case "azure-native:m365securityandcompliance:privateLinkServicesForO365ManagementActivityAPI":
		r = &PrivateLinkServicesForO365ManagementActivityAPI{}
	case "azure-native:m365securityandcompliance:privateLinkServicesForSCCPowershell":
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
		"m365securityandcompliance",
		&module{version},
	)
}
