


package v20200801

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
	case "azure-native:operationalinsights/v20200801:Cluster":
		r = &Cluster{}
	case "azure-native:operationalinsights/v20200801:DataExport":
		r = &DataExport{}
	case "azure-native:operationalinsights/v20200801:DataSource":
		r = &DataSource{}
	case "azure-native:operationalinsights/v20200801:LinkedService":
		r = &LinkedService{}
	case "azure-native:operationalinsights/v20200801:LinkedStorageAccount":
		r = &LinkedStorageAccount{}
	case "azure-native:operationalinsights/v20200801:SavedSearch":
		r = &SavedSearch{}
	case "azure-native:operationalinsights/v20200801:StorageInsightConfig":
		r = &StorageInsightConfig{}
	case "azure-native:operationalinsights/v20200801:Workspace":
		r = &Workspace{}
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
		"operationalinsights/v20200801",
		&module{version},
	)
}
