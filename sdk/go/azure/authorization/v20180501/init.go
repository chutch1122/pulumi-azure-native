


package v20180501

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
	case "azure-native:authorization/v20180501:PolicyAssignment":
		r = &PolicyAssignment{}
	case "azure-native:authorization/v20180501:PolicyDefinition":
		r = &PolicyDefinition{}
	case "azure-native:authorization/v20180501:PolicyDefinitionAtManagementGroup":
		r = &PolicyDefinitionAtManagementGroup{}
	case "azure-native:authorization/v20180501:PolicySetDefinition":
		r = &PolicySetDefinition{}
	case "azure-native:authorization/v20180501:PolicySetDefinitionAtManagementGroup":
		r = &PolicySetDefinitionAtManagementGroup{}
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
		"authorization/v20180501",
		&module{version},
	)
}
