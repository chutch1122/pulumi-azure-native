


package portal

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
	case "azure-native:portal:Console":
		r = &Console{}
	case "azure-native:portal:ConsoleWithLocation":
		r = &ConsoleWithLocation{}
	case "azure-native:portal:Dashboard":
		r = &Dashboard{}
	case "azure-native:portal:TenantConfiguration":
		r = &TenantConfiguration{}
	case "azure-native:portal:UserSettings":
		r = &UserSettings{}
	case "azure-native:portal:UserSettingsWithLocation":
		r = &UserSettingsWithLocation{}
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
		"portal",
		&module{version},
	)
}
