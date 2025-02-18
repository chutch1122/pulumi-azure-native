


package v20211115preview

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
	case "azure-native:documentdb/v20211115preview:CassandraCluster":
		r = &CassandraCluster{}
	case "azure-native:documentdb/v20211115preview:CassandraDataCenter":
		r = &CassandraDataCenter{}
	case "azure-native:documentdb/v20211115preview:CassandraResourceCassandraKeyspace":
		r = &CassandraResourceCassandraKeyspace{}
	case "azure-native:documentdb/v20211115preview:CassandraResourceCassandraTable":
		r = &CassandraResourceCassandraTable{}
	case "azure-native:documentdb/v20211115preview:CassandraResourceCassandraView":
		r = &CassandraResourceCassandraView{}
	case "azure-native:documentdb/v20211115preview:DatabaseAccount":
		r = &DatabaseAccount{}
	case "azure-native:documentdb/v20211115preview:GraphResourceGraph":
		r = &GraphResourceGraph{}
	case "azure-native:documentdb/v20211115preview:GremlinResourceGremlinDatabase":
		r = &GremlinResourceGremlinDatabase{}
	case "azure-native:documentdb/v20211115preview:GremlinResourceGremlinGraph":
		r = &GremlinResourceGremlinGraph{}
	case "azure-native:documentdb/v20211115preview:MongoDBResourceMongoDBCollection":
		r = &MongoDBResourceMongoDBCollection{}
	case "azure-native:documentdb/v20211115preview:MongoDBResourceMongoDBDatabase":
		r = &MongoDBResourceMongoDBDatabase{}
	case "azure-native:documentdb/v20211115preview:MongoDBResourceMongoRoleDefinition":
		r = &MongoDBResourceMongoRoleDefinition{}
	case "azure-native:documentdb/v20211115preview:MongoDBResourceMongoUserDefinition":
		r = &MongoDBResourceMongoUserDefinition{}
	case "azure-native:documentdb/v20211115preview:NotebookWorkspace":
		r = &NotebookWorkspace{}
	case "azure-native:documentdb/v20211115preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:documentdb/v20211115preview:Service":
		r = &Service{}
	case "azure-native:documentdb/v20211115preview:SqlResourceSqlContainer":
		r = &SqlResourceSqlContainer{}
	case "azure-native:documentdb/v20211115preview:SqlResourceSqlDatabase":
		r = &SqlResourceSqlDatabase{}
	case "azure-native:documentdb/v20211115preview:SqlResourceSqlRoleAssignment":
		r = &SqlResourceSqlRoleAssignment{}
	case "azure-native:documentdb/v20211115preview:SqlResourceSqlRoleDefinition":
		r = &SqlResourceSqlRoleDefinition{}
	case "azure-native:documentdb/v20211115preview:SqlResourceSqlStoredProcedure":
		r = &SqlResourceSqlStoredProcedure{}
	case "azure-native:documentdb/v20211115preview:SqlResourceSqlTrigger":
		r = &SqlResourceSqlTrigger{}
	case "azure-native:documentdb/v20211115preview:SqlResourceSqlUserDefinedFunction":
		r = &SqlResourceSqlUserDefinedFunction{}
	case "azure-native:documentdb/v20211115preview:TableResourceTable":
		r = &TableResourceTable{}
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
		"documentdb/v20211115preview",
		&module{version},
	)
}
