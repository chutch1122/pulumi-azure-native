


package v20200301

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
	case "azure-native:documentdb/v20200301:CassandraResourceCassandraKeyspace":
		r = &CassandraResourceCassandraKeyspace{}
	case "azure-native:documentdb/v20200301:CassandraResourceCassandraTable":
		r = &CassandraResourceCassandraTable{}
	case "azure-native:documentdb/v20200301:DatabaseAccount":
		r = &DatabaseAccount{}
	case "azure-native:documentdb/v20200301:GremlinResourceGremlinDatabase":
		r = &GremlinResourceGremlinDatabase{}
	case "azure-native:documentdb/v20200301:GremlinResourceGremlinGraph":
		r = &GremlinResourceGremlinGraph{}
	case "azure-native:documentdb/v20200301:MongoDBResourceMongoDBCollection":
		r = &MongoDBResourceMongoDBCollection{}
	case "azure-native:documentdb/v20200301:MongoDBResourceMongoDBDatabase":
		r = &MongoDBResourceMongoDBDatabase{}
	case "azure-native:documentdb/v20200301:NotebookWorkspace":
		r = &NotebookWorkspace{}
	case "azure-native:documentdb/v20200301:SqlResourceSqlContainer":
		r = &SqlResourceSqlContainer{}
	case "azure-native:documentdb/v20200301:SqlResourceSqlDatabase":
		r = &SqlResourceSqlDatabase{}
	case "azure-native:documentdb/v20200301:SqlResourceSqlStoredProcedure":
		r = &SqlResourceSqlStoredProcedure{}
	case "azure-native:documentdb/v20200301:SqlResourceSqlTrigger":
		r = &SqlResourceSqlTrigger{}
	case "azure-native:documentdb/v20200301:SqlResourceSqlUserDefinedFunction":
		r = &SqlResourceSqlUserDefinedFunction{}
	case "azure-native:documentdb/v20200301:TableResourceTable":
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
		"documentdb/v20200301",
		&module{version},
	)
}
