


package v20200515

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
	case "azure-native:timeseriesinsights/v20200515:AccessPolicy":
		r = &AccessPolicy{}
	case "azure-native:timeseriesinsights/v20200515:Environment":
		r = &Environment{}
	case "azure-native:timeseriesinsights/v20200515:EventHubEventSource":
		r = &EventHubEventSource{}
	case "azure-native:timeseriesinsights/v20200515:EventSource":
		r = &EventSource{}
	case "azure-native:timeseriesinsights/v20200515:Gen1Environment":
		r = &Gen1Environment{}
	case "azure-native:timeseriesinsights/v20200515:Gen2Environment":
		r = &Gen2Environment{}
	case "azure-native:timeseriesinsights/v20200515:IoTHubEventSource":
		r = &IoTHubEventSource{}
	case "azure-native:timeseriesinsights/v20200515:ReferenceDataSet":
		r = &ReferenceDataSet{}
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
		"timeseriesinsights/v20200515",
		&module{version},
	)
}
