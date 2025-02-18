


package datafactory

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
	case "azure-native:datafactory:DataFlow":
		r = &DataFlow{}
	case "azure-native:datafactory:Dataset":
		r = &Dataset{}
	case "azure-native:datafactory:Factory":
		r = &Factory{}
	case "azure-native:datafactory:GlobalParameter":
		r = &GlobalParameter{}
	case "azure-native:datafactory:IntegrationRuntime":
		r = &IntegrationRuntime{}
	case "azure-native:datafactory:LinkedService":
		r = &LinkedService{}
	case "azure-native:datafactory:ManagedPrivateEndpoint":
		r = &ManagedPrivateEndpoint{}
	case "azure-native:datafactory:Pipeline":
		r = &Pipeline{}
	case "azure-native:datafactory:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:datafactory:Trigger":
		r = &Trigger{}
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
		"datafactory",
		&module{version},
	)
}
