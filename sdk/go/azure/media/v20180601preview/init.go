


package v20180601preview

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
	case "azure-native:media/v20180601preview:Asset":
		r = &Asset{}
	case "azure-native:media/v20180601preview:ContentKeyPolicy":
		r = &ContentKeyPolicy{}
	case "azure-native:media/v20180601preview:Job":
		r = &Job{}
	case "azure-native:media/v20180601preview:LiveEvent":
		r = &LiveEvent{}
	case "azure-native:media/v20180601preview:LiveOutput":
		r = &LiveOutput{}
	case "azure-native:media/v20180601preview:MediaService":
		r = &MediaService{}
	case "azure-native:media/v20180601preview:StreamingEndpoint":
		r = &StreamingEndpoint{}
	case "azure-native:media/v20180601preview:StreamingLocator":
		r = &StreamingLocator{}
	case "azure-native:media/v20180601preview:StreamingPolicy":
		r = &StreamingPolicy{}
	case "azure-native:media/v20180601preview:Transform":
		r = &Transform{}
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
		"media/v20180601preview",
		&module{version},
	)
}
