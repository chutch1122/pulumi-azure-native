


package v20181015

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
	case "azure-native:labservices/v20181015:Environment":
		r = &Environment{}
	case "azure-native:labservices/v20181015:EnvironmentSetting":
		r = &EnvironmentSetting{}
	case "azure-native:labservices/v20181015:GalleryImage":
		r = &GalleryImage{}
	case "azure-native:labservices/v20181015:Lab":
		r = &Lab{}
	case "azure-native:labservices/v20181015:LabAccount":
		r = &LabAccount{}
	case "azure-native:labservices/v20181015:User":
		r = &User{}
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
		"labservices/v20181015",
		&module{version},
	)
}
