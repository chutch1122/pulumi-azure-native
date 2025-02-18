


package v20200101

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
	case "azure-native:security/v20200101:AdaptiveApplicationControl":
		r = &AdaptiveApplicationControl{}
	case "azure-native:security/v20200101:Assessment":
		r = &Assessment{}
	case "azure-native:security/v20200101:AssessmentMetadataInSubscription":
		r = &AssessmentMetadataInSubscription{}
	case "azure-native:security/v20200101:JitNetworkAccessPolicy":
		r = &JitNetworkAccessPolicy{}
	case "azure-native:security/v20200101:ServerVulnerabilityAssessment":
		r = &ServerVulnerabilityAssessment{}
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
		"security/v20200101",
		&module{version},
	)
}
