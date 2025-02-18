


package v20150501preview

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
	case "azure-native:sql/v20150501preview:DatabaseAdvisor":
		r = &DatabaseAdvisor{}
	case "azure-native:sql/v20150501preview:DatabaseBlobAuditingPolicy":
		r = &DatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20150501preview:EncryptionProtector":
		r = &EncryptionProtector{}
	case "azure-native:sql/v20150501preview:FailoverGroup":
		r = &FailoverGroup{}
	case "azure-native:sql/v20150501preview:FirewallRule":
		r = &FirewallRule{}
	case "azure-native:sql/v20150501preview:ManagedInstance":
		r = &ManagedInstance{}
	case "azure-native:sql/v20150501preview:Server":
		r = &Server{}
	case "azure-native:sql/v20150501preview:ServerAdvisor":
		r = &ServerAdvisor{}
	case "azure-native:sql/v20150501preview:ServerKey":
		r = &ServerKey{}
	case "azure-native:sql/v20150501preview:SyncAgent":
		r = &SyncAgent{}
	case "azure-native:sql/v20150501preview:SyncGroup":
		r = &SyncGroup{}
	case "azure-native:sql/v20150501preview:SyncMember":
		r = &SyncMember{}
	case "azure-native:sql/v20150501preview:VirtualNetworkRule":
		r = &VirtualNetworkRule{}
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
		"sql/v20150501preview",
		&module{version},
	)
}
