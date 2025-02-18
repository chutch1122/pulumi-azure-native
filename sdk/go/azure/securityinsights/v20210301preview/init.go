


package v20210301preview

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
	case "azure-native:securityinsights/v20210301preview:AADDataConnector":
		r = &AADDataConnector{}
	case "azure-native:securityinsights/v20210301preview:AATPDataConnector":
		r = &AATPDataConnector{}
	case "azure-native:securityinsights/v20210301preview:ASCDataConnector":
		r = &ASCDataConnector{}
	case "azure-native:securityinsights/v20210301preview:Action":
		r = &Action{}
	case "azure-native:securityinsights/v20210301preview:ActivityCustomEntityQuery":
		r = &ActivityCustomEntityQuery{}
	case "azure-native:securityinsights/v20210301preview:AlertRule":
		r = &AlertRule{}
	case "azure-native:securityinsights/v20210301preview:Anomalies":
		r = &Anomalies{}
	case "azure-native:securityinsights/v20210301preview:AwsCloudTrailDataConnector":
		r = &AwsCloudTrailDataConnector{}
	case "azure-native:securityinsights/v20210301preview:CodelessUiDataConnector":
		r = &CodelessUiDataConnector{}
	case "azure-native:securityinsights/v20210301preview:DataConnector":
		r = &DataConnector{}
	case "azure-native:securityinsights/v20210301preview:Dynamics365DataConnector":
		r = &Dynamics365DataConnector{}
	case "azure-native:securityinsights/v20210301preview:EntityAnalytics":
		r = &EntityAnalytics{}
	case "azure-native:securityinsights/v20210301preview:EntityQuery":
		r = &EntityQuery{}
	case "azure-native:securityinsights/v20210301preview:EyesOn":
		r = &EyesOn{}
	case "azure-native:securityinsights/v20210301preview:FusionAlertRule":
		r = &FusionAlertRule{}
	case "azure-native:securityinsights/v20210301preview:Incident":
		r = &Incident{}
	case "azure-native:securityinsights/v20210301preview:IncidentComment":
		r = &IncidentComment{}
	case "azure-native:securityinsights/v20210301preview:IncidentRelation":
		r = &IncidentRelation{}
	case "azure-native:securityinsights/v20210301preview:MCASDataConnector":
		r = &MCASDataConnector{}
	case "azure-native:securityinsights/v20210301preview:MDATPDataConnector":
		r = &MDATPDataConnector{}
	case "azure-native:securityinsights/v20210301preview:MLBehaviorAnalyticsAlertRule":
		r = &MLBehaviorAnalyticsAlertRule{}
	case "azure-native:securityinsights/v20210301preview:MSTIDataConnector":
		r = &MSTIDataConnector{}
	case "azure-native:securityinsights/v20210301preview:MTPDataConnector":
		r = &MTPDataConnector{}
	case "azure-native:securityinsights/v20210301preview:Metadata":
		r = &Metadata{}
	case "azure-native:securityinsights/v20210301preview:MicrosoftSecurityIncidentCreationAlertRule":
		r = &MicrosoftSecurityIncidentCreationAlertRule{}
	case "azure-native:securityinsights/v20210301preview:OfficeATPDataConnector":
		r = &OfficeATPDataConnector{}
	case "azure-native:securityinsights/v20210301preview:OfficeDataConnector":
		r = &OfficeDataConnector{}
	case "azure-native:securityinsights/v20210301preview:ProductSetting":
		r = &ProductSetting{}
	case "azure-native:securityinsights/v20210301preview:ScheduledAlertRule":
		r = &ScheduledAlertRule{}
	case "azure-native:securityinsights/v20210301preview:SentinelOnboardingState":
		r = &SentinelOnboardingState{}
	case "azure-native:securityinsights/v20210301preview:SourceControl":
		r = &SourceControl{}
	case "azure-native:securityinsights/v20210301preview:TIDataConnector":
		r = &TIDataConnector{}
	case "azure-native:securityinsights/v20210301preview:ThreatIntelligenceAlertRule":
		r = &ThreatIntelligenceAlertRule{}
	case "azure-native:securityinsights/v20210301preview:TiTaxiiDataConnector":
		r = &TiTaxiiDataConnector{}
	case "azure-native:securityinsights/v20210301preview:Ueba":
		r = &Ueba{}
	case "azure-native:securityinsights/v20210301preview:Watchlist":
		r = &Watchlist{}
	case "azure-native:securityinsights/v20210301preview:WatchlistItem":
		r = &WatchlistItem{}
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
		"securityinsights/v20210301preview",
		&module{version},
	)
}
