


package storage

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
	case "azure-native:storage:Blob":
		r = &Blob{}
	case "azure-native:storage:BlobContainer":
		r = &BlobContainer{}
	case "azure-native:storage:BlobContainerImmutabilityPolicy":
		r = &BlobContainerImmutabilityPolicy{}
	case "azure-native:storage:BlobInventoryPolicy":
		r = &BlobInventoryPolicy{}
	case "azure-native:storage:BlobServiceProperties":
		r = &BlobServiceProperties{}
	case "azure-native:storage:EncryptionScope":
		r = &EncryptionScope{}
	case "azure-native:storage:FileServiceProperties":
		r = &FileServiceProperties{}
	case "azure-native:storage:FileShare":
		r = &FileShare{}
	case "azure-native:storage:LocalUser":
		r = &LocalUser{}
	case "azure-native:storage:ManagementPolicy":
		r = &ManagementPolicy{}
	case "azure-native:storage:ObjectReplicationPolicy":
		r = &ObjectReplicationPolicy{}
	case "azure-native:storage:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:storage:Queue":
		r = &Queue{}
	case "azure-native:storage:QueueServiceProperties":
		r = &QueueServiceProperties{}
	case "azure-native:storage:StorageAccount":
		r = &StorageAccount{}
	case "azure-native:storage:StorageAccountStaticWebsite":
		r = &StorageAccountStaticWebsite{}
	case "azure-native:storage:Table":
		r = &Table{}
	case "azure-native:storage:TableServiceProperties":
		r = &TableServiceProperties{}
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
		"storage",
		&module{version},
	)
}
