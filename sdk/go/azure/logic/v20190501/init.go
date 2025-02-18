


package v20190501

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
	case "azure-native:logic/v20190501:IntegrationAccount":
		r = &IntegrationAccount{}
	case "azure-native:logic/v20190501:IntegrationAccountAgreement":
		r = &IntegrationAccountAgreement{}
	case "azure-native:logic/v20190501:IntegrationAccountAssembly":
		r = &IntegrationAccountAssembly{}
	case "azure-native:logic/v20190501:IntegrationAccountBatchConfiguration":
		r = &IntegrationAccountBatchConfiguration{}
	case "azure-native:logic/v20190501:IntegrationAccountCertificate":
		r = &IntegrationAccountCertificate{}
	case "azure-native:logic/v20190501:IntegrationAccountMap":
		r = &IntegrationAccountMap{}
	case "azure-native:logic/v20190501:IntegrationAccountPartner":
		r = &IntegrationAccountPartner{}
	case "azure-native:logic/v20190501:IntegrationAccountSchema":
		r = &IntegrationAccountSchema{}
	case "azure-native:logic/v20190501:IntegrationAccountSession":
		r = &IntegrationAccountSession{}
	case "azure-native:logic/v20190501:IntegrationServiceEnvironment":
		r = &IntegrationServiceEnvironment{}
	case "azure-native:logic/v20190501:IntegrationServiceEnvironmentManagedApi":
		r = &IntegrationServiceEnvironmentManagedApi{}
	case "azure-native:logic/v20190501:Workflow":
		r = &Workflow{}
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
		"logic/v20190501",
		&module{version},
	)
}
