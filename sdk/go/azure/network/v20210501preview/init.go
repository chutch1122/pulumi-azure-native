


package v20210501preview

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
	case "azure-native:network/v20210501preview:AdminRule":
		r = &AdminRule{}
	case "azure-native:network/v20210501preview:AdminRuleCollection":
		r = &AdminRuleCollection{}
	case "azure-native:network/v20210501preview:ConnectivityConfiguration":
		r = &ConnectivityConfiguration{}
	case "azure-native:network/v20210501preview:DefaultAdminRule":
		r = &DefaultAdminRule{}
	case "azure-native:network/v20210501preview:DefaultUserRule":
		r = &DefaultUserRule{}
	case "azure-native:network/v20210501preview:ManagementGroupNetworkManagerConnection":
		r = &ManagementGroupNetworkManagerConnection{}
	case "azure-native:network/v20210501preview:NetworkGroup":
		r = &NetworkGroup{}
	case "azure-native:network/v20210501preview:NetworkManager":
		r = &NetworkManager{}
	case "azure-native:network/v20210501preview:ScopeConnection":
		r = &ScopeConnection{}
	case "azure-native:network/v20210501preview:SecurityAdminConfiguration":
		r = &SecurityAdminConfiguration{}
	case "azure-native:network/v20210501preview:SecurityUserConfiguration":
		r = &SecurityUserConfiguration{}
	case "azure-native:network/v20210501preview:StaticMember":
		r = &StaticMember{}
	case "azure-native:network/v20210501preview:SubscriptionNetworkManagerConnection":
		r = &SubscriptionNetworkManagerConnection{}
	case "azure-native:network/v20210501preview:UserRule":
		r = &UserRule{}
	case "azure-native:network/v20210501preview:UserRuleCollection":
		r = &UserRuleCollection{}
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
		"network/v20210501preview",
		&module{version},
	)
}
