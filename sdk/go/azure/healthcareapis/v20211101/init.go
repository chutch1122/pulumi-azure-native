


package v20211101

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
	case "azure-native:healthcareapis/v20211101:DicomService":
		r = &DicomService{}
	case "azure-native:healthcareapis/v20211101:FhirService":
		r = &FhirService{}
	case "azure-native:healthcareapis/v20211101:IotConnector":
		r = &IotConnector{}
	case "azure-native:healthcareapis/v20211101:IotConnectorFhirDestination":
		r = &IotConnectorFhirDestination{}
	case "azure-native:healthcareapis/v20211101:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:healthcareapis/v20211101:Service":
		r = &Service{}
	case "azure-native:healthcareapis/v20211101:Workspace":
		r = &Workspace{}
	case "azure-native:healthcareapis/v20211101:WorkspacePrivateEndpointConnection":
		r = &WorkspacePrivateEndpointConnection{}
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
		"healthcareapis/v20211101",
		&module{version},
	)
}
