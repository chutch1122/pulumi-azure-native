


package datashare

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
	case "azure-native:datashare:ADLSGen1FileDataSet":
		r = &ADLSGen1FileDataSet{}
	case "azure-native:datashare:ADLSGen1FolderDataSet":
		r = &ADLSGen1FolderDataSet{}
	case "azure-native:datashare:ADLSGen2FileDataSet":
		r = &ADLSGen2FileDataSet{}
	case "azure-native:datashare:ADLSGen2FileDataSetMapping":
		r = &ADLSGen2FileDataSetMapping{}
	case "azure-native:datashare:ADLSGen2FileSystemDataSet":
		r = &ADLSGen2FileSystemDataSet{}
	case "azure-native:datashare:ADLSGen2FileSystemDataSetMapping":
		r = &ADLSGen2FileSystemDataSetMapping{}
	case "azure-native:datashare:ADLSGen2FolderDataSet":
		r = &ADLSGen2FolderDataSet{}
	case "azure-native:datashare:ADLSGen2FolderDataSetMapping":
		r = &ADLSGen2FolderDataSetMapping{}
	case "azure-native:datashare:Account":
		r = &Account{}
	case "azure-native:datashare:BlobContainerDataSet":
		r = &BlobContainerDataSet{}
	case "azure-native:datashare:BlobContainerDataSetMapping":
		r = &BlobContainerDataSetMapping{}
	case "azure-native:datashare:BlobDataSet":
		r = &BlobDataSet{}
	case "azure-native:datashare:BlobDataSetMapping":
		r = &BlobDataSetMapping{}
	case "azure-native:datashare:BlobFolderDataSet":
		r = &BlobFolderDataSet{}
	case "azure-native:datashare:BlobFolderDataSetMapping":
		r = &BlobFolderDataSetMapping{}
	case "azure-native:datashare:DataSet":
		r = &DataSet{}
	case "azure-native:datashare:DataSetMapping":
		r = &DataSetMapping{}
	case "azure-native:datashare:Invitation":
		r = &Invitation{}
	case "azure-native:datashare:KustoClusterDataSet":
		r = &KustoClusterDataSet{}
	case "azure-native:datashare:KustoClusterDataSetMapping":
		r = &KustoClusterDataSetMapping{}
	case "azure-native:datashare:KustoDatabaseDataSet":
		r = &KustoDatabaseDataSet{}
	case "azure-native:datashare:KustoDatabaseDataSetMapping":
		r = &KustoDatabaseDataSetMapping{}
	case "azure-native:datashare:ScheduledSynchronizationSetting":
		r = &ScheduledSynchronizationSetting{}
	case "azure-native:datashare:ScheduledTrigger":
		r = &ScheduledTrigger{}
	case "azure-native:datashare:Share":
		r = &Share{}
	case "azure-native:datashare:ShareSubscription":
		r = &ShareSubscription{}
	case "azure-native:datashare:SqlDBTableDataSet":
		r = &SqlDBTableDataSet{}
	case "azure-native:datashare:SqlDBTableDataSetMapping":
		r = &SqlDBTableDataSetMapping{}
	case "azure-native:datashare:SqlDWTableDataSet":
		r = &SqlDWTableDataSet{}
	case "azure-native:datashare:SqlDWTableDataSetMapping":
		r = &SqlDWTableDataSetMapping{}
	case "azure-native:datashare:SynapseWorkspaceSqlPoolTableDataSet":
		r = &SynapseWorkspaceSqlPoolTableDataSet{}
	case "azure-native:datashare:SynapseWorkspaceSqlPoolTableDataSetMapping":
		r = &SynapseWorkspaceSqlPoolTableDataSetMapping{}
	case "azure-native:datashare:SynchronizationSetting":
		r = &SynchronizationSetting{}
	case "azure-native:datashare:Trigger":
		r = &Trigger{}
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
		"datashare",
		&module{version},
	)
}
