


package v20210801

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
	case "azure-native:recoveryservices/v20210801:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:recoveryservices/v20210801:ProtectedItem":
		r = &ProtectedItem{}
	case "azure-native:recoveryservices/v20210801:ProtectionContainer":
		r = &ProtectionContainer{}
	case "azure-native:recoveryservices/v20210801:ProtectionIntent":
		r = &ProtectionIntent{}
	case "azure-native:recoveryservices/v20210801:ProtectionPolicy":
		r = &ProtectionPolicy{}
	case "azure-native:recoveryservices/v20210801:ReplicationFabric":
		r = &ReplicationFabric{}
	case "azure-native:recoveryservices/v20210801:ReplicationMigrationItem":
		r = &ReplicationMigrationItem{}
	case "azure-native:recoveryservices/v20210801:ReplicationNetworkMapping":
		r = &ReplicationNetworkMapping{}
	case "azure-native:recoveryservices/v20210801:ReplicationPolicy":
		r = &ReplicationPolicy{}
	case "azure-native:recoveryservices/v20210801:ReplicationProtectedItem":
		r = &ReplicationProtectedItem{}
	case "azure-native:recoveryservices/v20210801:ReplicationProtectionContainerMapping":
		r = &ReplicationProtectionContainerMapping{}
	case "azure-native:recoveryservices/v20210801:ReplicationRecoveryPlan":
		r = &ReplicationRecoveryPlan{}
	case "azure-native:recoveryservices/v20210801:ReplicationRecoveryServicesProvider":
		r = &ReplicationRecoveryServicesProvider{}
	case "azure-native:recoveryservices/v20210801:ReplicationStorageClassificationMapping":
		r = &ReplicationStorageClassificationMapping{}
	case "azure-native:recoveryservices/v20210801:ReplicationvCenter":
		r = &ReplicationvCenter{}
	case "azure-native:recoveryservices/v20210801:ResourceGuardProxy":
		r = &ResourceGuardProxy{}
	case "azure-native:recoveryservices/v20210801:Vault":
		r = &Vault{}
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
		"recoveryservices/v20210801",
		&module{version},
	)
}
