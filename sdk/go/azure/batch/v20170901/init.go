


package v20170901

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
	case "azure-native:batch/v20170901:Application":
		r = &Application{}
	case "azure-native:batch/v20170901:ApplicationPackage":
		r = &ApplicationPackage{}
	case "azure-native:batch/v20170901:BatchAccount":
		r = &BatchAccount{}
	case "azure-native:batch/v20170901:Certificate":
		r = &Certificate{}
	case "azure-native:batch/v20170901:Pool":
		r = &Pool{}
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
		"batch/v20170901",
		&module{version},
	)
}
