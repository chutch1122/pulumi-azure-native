


package v20220103

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
	case "azure-native:compute/v20220103:Gallery":
		r = &Gallery{}
	case "azure-native:compute/v20220103:GalleryApplication":
		r = &GalleryApplication{}
	case "azure-native:compute/v20220103:GalleryApplicationVersion":
		r = &GalleryApplicationVersion{}
	case "azure-native:compute/v20220103:GalleryImage":
		r = &GalleryImage{}
	case "azure-native:compute/v20220103:GalleryImageVersion":
		r = &GalleryImageVersion{}
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
		"compute/v20220103",
		&module{version},
	)
}
