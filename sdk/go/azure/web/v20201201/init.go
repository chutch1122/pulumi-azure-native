


package v20201201

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
	case "azure-native:web/v20201201:AppServiceEnvironment":
		r = &AppServiceEnvironment{}
	case "azure-native:web/v20201201:AppServiceEnvironmentPrivateEndpointConnection":
		r = &AppServiceEnvironmentPrivateEndpointConnection{}
	case "azure-native:web/v20201201:AppServicePlan":
		r = &AppServicePlan{}
	case "azure-native:web/v20201201:AppServicePlanRouteForVnet":
		r = &AppServicePlanRouteForVnet{}
	case "azure-native:web/v20201201:Certificate":
		r = &Certificate{}
	case "azure-native:web/v20201201:StaticSite":
		r = &StaticSite{}
	case "azure-native:web/v20201201:StaticSiteCustomDomain":
		r = &StaticSiteCustomDomain{}
	case "azure-native:web/v20201201:StaticSitePrivateEndpointConnection":
		r = &StaticSitePrivateEndpointConnection{}
	case "azure-native:web/v20201201:StaticSiteUserProvidedFunctionAppForStaticSite":
		r = &StaticSiteUserProvidedFunctionAppForStaticSite{}
	case "azure-native:web/v20201201:StaticSiteUserProvidedFunctionAppForStaticSiteBuild":
		r = &StaticSiteUserProvidedFunctionAppForStaticSiteBuild{}
	case "azure-native:web/v20201201:WebApp":
		r = &WebApp{}
	case "azure-native:web/v20201201:WebAppApplicationSettings":
		r = &WebAppApplicationSettings{}
	case "azure-native:web/v20201201:WebAppApplicationSettingsSlot":
		r = &WebAppApplicationSettingsSlot{}
	case "azure-native:web/v20201201:WebAppAuthSettings":
		r = &WebAppAuthSettings{}
	case "azure-native:web/v20201201:WebAppAuthSettingsSlot":
		r = &WebAppAuthSettingsSlot{}
	case "azure-native:web/v20201201:WebAppAuthSettingsV2":
		r = &WebAppAuthSettingsV2{}
	case "azure-native:web/v20201201:WebAppAuthSettingsV2Slot":
		r = &WebAppAuthSettingsV2Slot{}
	case "azure-native:web/v20201201:WebAppAzureStorageAccounts":
		r = &WebAppAzureStorageAccounts{}
	case "azure-native:web/v20201201:WebAppAzureStorageAccountsSlot":
		r = &WebAppAzureStorageAccountsSlot{}
	case "azure-native:web/v20201201:WebAppBackupConfiguration":
		r = &WebAppBackupConfiguration{}
	case "azure-native:web/v20201201:WebAppBackupConfigurationSlot":
		r = &WebAppBackupConfigurationSlot{}
	case "azure-native:web/v20201201:WebAppConnectionStrings":
		r = &WebAppConnectionStrings{}
	case "azure-native:web/v20201201:WebAppConnectionStringsSlot":
		r = &WebAppConnectionStringsSlot{}
	case "azure-native:web/v20201201:WebAppDeployment":
		r = &WebAppDeployment{}
	case "azure-native:web/v20201201:WebAppDeploymentSlot":
		r = &WebAppDeploymentSlot{}
	case "azure-native:web/v20201201:WebAppDiagnosticLogsConfiguration":
		r = &WebAppDiagnosticLogsConfiguration{}
	case "azure-native:web/v20201201:WebAppDomainOwnershipIdentifier":
		r = &WebAppDomainOwnershipIdentifier{}
	case "azure-native:web/v20201201:WebAppDomainOwnershipIdentifierSlot":
		r = &WebAppDomainOwnershipIdentifierSlot{}
	case "azure-native:web/v20201201:WebAppFunction":
		r = &WebAppFunction{}
	case "azure-native:web/v20201201:WebAppHostNameBinding":
		r = &WebAppHostNameBinding{}
	case "azure-native:web/v20201201:WebAppHostNameBindingSlot":
		r = &WebAppHostNameBindingSlot{}
	case "azure-native:web/v20201201:WebAppHybridConnection":
		r = &WebAppHybridConnection{}
	case "azure-native:web/v20201201:WebAppHybridConnectionSlot":
		r = &WebAppHybridConnectionSlot{}
	case "azure-native:web/v20201201:WebAppInstanceFunctionSlot":
		r = &WebAppInstanceFunctionSlot{}
	case "azure-native:web/v20201201:WebAppMetadata":
		r = &WebAppMetadata{}
	case "azure-native:web/v20201201:WebAppMetadataSlot":
		r = &WebAppMetadataSlot{}
	case "azure-native:web/v20201201:WebAppPremierAddOn":
		r = &WebAppPremierAddOn{}
	case "azure-native:web/v20201201:WebAppPremierAddOnSlot":
		r = &WebAppPremierAddOnSlot{}
	case "azure-native:web/v20201201:WebAppPrivateEndpointConnection":
		r = &WebAppPrivateEndpointConnection{}
	case "azure-native:web/v20201201:WebAppPrivateEndpointConnectionSlot":
		r = &WebAppPrivateEndpointConnectionSlot{}
	case "azure-native:web/v20201201:WebAppPublicCertificate":
		r = &WebAppPublicCertificate{}
	case "azure-native:web/v20201201:WebAppPublicCertificateSlot":
		r = &WebAppPublicCertificateSlot{}
	case "azure-native:web/v20201201:WebAppRelayServiceConnection":
		r = &WebAppRelayServiceConnection{}
	case "azure-native:web/v20201201:WebAppRelayServiceConnectionSlot":
		r = &WebAppRelayServiceConnectionSlot{}
	case "azure-native:web/v20201201:WebAppSiteExtension":
		r = &WebAppSiteExtension{}
	case "azure-native:web/v20201201:WebAppSiteExtensionSlot":
		r = &WebAppSiteExtensionSlot{}
	case "azure-native:web/v20201201:WebAppSitePushSettings":
		r = &WebAppSitePushSettings{}
	case "azure-native:web/v20201201:WebAppSitePushSettingsSlot":
		r = &WebAppSitePushSettingsSlot{}
	case "azure-native:web/v20201201:WebAppSlot":
		r = &WebAppSlot{}
	case "azure-native:web/v20201201:WebAppSlotConfigurationNames":
		r = &WebAppSlotConfigurationNames{}
	case "azure-native:web/v20201201:WebAppSourceControl":
		r = &WebAppSourceControl{}
	case "azure-native:web/v20201201:WebAppSourceControlSlot":
		r = &WebAppSourceControlSlot{}
	case "azure-native:web/v20201201:WebAppSwiftVirtualNetworkConnection":
		r = &WebAppSwiftVirtualNetworkConnection{}
	case "azure-native:web/v20201201:WebAppVnetConnection":
		r = &WebAppVnetConnection{}
	case "azure-native:web/v20201201:WebAppVnetConnectionSlot":
		r = &WebAppVnetConnectionSlot{}
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
		"web/v20201201",
		&module{version},
	)
}
