


package v20210101preview

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
	case "azure-native:apimanagement/v20210101preview:Api":
		r = &Api{}
	case "azure-native:apimanagement/v20210101preview:ApiDiagnostic":
		r = &ApiDiagnostic{}
	case "azure-native:apimanagement/v20210101preview:ApiIssue":
		r = &ApiIssue{}
	case "azure-native:apimanagement/v20210101preview:ApiIssueAttachment":
		r = &ApiIssueAttachment{}
	case "azure-native:apimanagement/v20210101preview:ApiIssueComment":
		r = &ApiIssueComment{}
	case "azure-native:apimanagement/v20210101preview:ApiManagementService":
		r = &ApiManagementService{}
	case "azure-native:apimanagement/v20210101preview:ApiOperation":
		r = &ApiOperation{}
	case "azure-native:apimanagement/v20210101preview:ApiOperationPolicy":
		r = &ApiOperationPolicy{}
	case "azure-native:apimanagement/v20210101preview:ApiPolicy":
		r = &ApiPolicy{}
	case "azure-native:apimanagement/v20210101preview:ApiRelease":
		r = &ApiRelease{}
	case "azure-native:apimanagement/v20210101preview:ApiSchema":
		r = &ApiSchema{}
	case "azure-native:apimanagement/v20210101preview:ApiTagDescription":
		r = &ApiTagDescription{}
	case "azure-native:apimanagement/v20210101preview:ApiVersionSet":
		r = &ApiVersionSet{}
	case "azure-native:apimanagement/v20210101preview:AuthorizationServer":
		r = &AuthorizationServer{}
	case "azure-native:apimanagement/v20210101preview:Backend":
		r = &Backend{}
	case "azure-native:apimanagement/v20210101preview:Cache":
		r = &Cache{}
	case "azure-native:apimanagement/v20210101preview:Certificate":
		r = &Certificate{}
	case "azure-native:apimanagement/v20210101preview:ContentItem":
		r = &ContentItem{}
	case "azure-native:apimanagement/v20210101preview:ContentType":
		r = &ContentType{}
	case "azure-native:apimanagement/v20210101preview:Diagnostic":
		r = &Diagnostic{}
	case "azure-native:apimanagement/v20210101preview:EmailTemplate":
		r = &EmailTemplate{}
	case "azure-native:apimanagement/v20210101preview:Gateway":
		r = &Gateway{}
	case "azure-native:apimanagement/v20210101preview:GatewayApiEntityTag":
		r = &GatewayApiEntityTag{}
	case "azure-native:apimanagement/v20210101preview:GatewayCertificateAuthority":
		r = &GatewayCertificateAuthority{}
	case "azure-native:apimanagement/v20210101preview:GatewayHostnameConfiguration":
		r = &GatewayHostnameConfiguration{}
	case "azure-native:apimanagement/v20210101preview:Group":
		r = &Group{}
	case "azure-native:apimanagement/v20210101preview:GroupUser":
		r = &GroupUser{}
	case "azure-native:apimanagement/v20210101preview:IdentityProvider":
		r = &IdentityProvider{}
	case "azure-native:apimanagement/v20210101preview:Logger":
		r = &Logger{}
	case "azure-native:apimanagement/v20210101preview:NamedValue":
		r = &NamedValue{}
	case "azure-native:apimanagement/v20210101preview:NotificationRecipientEmail":
		r = &NotificationRecipientEmail{}
	case "azure-native:apimanagement/v20210101preview:NotificationRecipientUser":
		r = &NotificationRecipientUser{}
	case "azure-native:apimanagement/v20210101preview:OpenIdConnectProvider":
		r = &OpenIdConnectProvider{}
	case "azure-native:apimanagement/v20210101preview:Policy":
		r = &Policy{}
	case "azure-native:apimanagement/v20210101preview:Product":
		r = &Product{}
	case "azure-native:apimanagement/v20210101preview:ProductApi":
		r = &ProductApi{}
	case "azure-native:apimanagement/v20210101preview:ProductGroup":
		r = &ProductGroup{}
	case "azure-native:apimanagement/v20210101preview:ProductPolicy":
		r = &ProductPolicy{}
	case "azure-native:apimanagement/v20210101preview:Subscription":
		r = &Subscription{}
	case "azure-native:apimanagement/v20210101preview:Tag":
		r = &Tag{}
	case "azure-native:apimanagement/v20210101preview:TagByApi":
		r = &TagByApi{}
	case "azure-native:apimanagement/v20210101preview:TagByOperation":
		r = &TagByOperation{}
	case "azure-native:apimanagement/v20210101preview:TagByProduct":
		r = &TagByProduct{}
	case "azure-native:apimanagement/v20210101preview:User":
		r = &User{}
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
		"apimanagement/v20210101preview",
		&module{version},
	)
}
