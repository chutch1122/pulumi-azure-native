


package v20200918

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
	case "azure-native:kusto/v20200918:AttachedDatabaseConfiguration":
		r = &AttachedDatabaseConfiguration{}
	case "azure-native:kusto/v20200918:Cluster":
		r = &Cluster{}
	case "azure-native:kusto/v20200918:ClusterPrincipalAssignment":
		r = &ClusterPrincipalAssignment{}
	case "azure-native:kusto/v20200918:DataConnection":
		r = &DataConnection{}
	case "azure-native:kusto/v20200918:Database":
		r = &Database{}
	case "azure-native:kusto/v20200918:DatabasePrincipalAssignment":
		r = &DatabasePrincipalAssignment{}
	case "azure-native:kusto/v20200918:EventGridDataConnection":
		r = &EventGridDataConnection{}
	case "azure-native:kusto/v20200918:EventHubDataConnection":
		r = &EventHubDataConnection{}
	case "azure-native:kusto/v20200918:IotHubDataConnection":
		r = &IotHubDataConnection{}
	case "azure-native:kusto/v20200918:ReadOnlyFollowingDatabase":
		r = &ReadOnlyFollowingDatabase{}
	case "azure-native:kusto/v20200918:ReadWriteDatabase":
		r = &ReadWriteDatabase{}
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
		"kusto/v20200918",
		&module{version},
	)
}
