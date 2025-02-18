


package v20220101preview

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
	case "azure-native:logz/v20220101preview:MetricsSource":
		r = &MetricsSource{}
	case "azure-native:logz/v20220101preview:MetricsSourceTagRule":
		r = &MetricsSourceTagRule{}
	case "azure-native:logz/v20220101preview:Monitor":
		r = &Monitor{}
	case "azure-native:logz/v20220101preview:SubAccount":
		r = &SubAccount{}
	case "azure-native:logz/v20220101preview:SubAccountTagRule":
		r = &SubAccountTagRule{}
	case "azure-native:logz/v20220101preview:TagRule":
		r = &TagRule{}
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
		"logz/v20220101preview",
		&module{version},
	)
}
