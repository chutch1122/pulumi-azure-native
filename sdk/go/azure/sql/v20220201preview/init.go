


package v20220201preview

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
	case "azure-native:sql/v20220201preview:BackupShortTermRetentionPolicy":
		r = &BackupShortTermRetentionPolicy{}
	case "azure-native:sql/v20220201preview:DataMaskingPolicy":
		r = &DataMaskingPolicy{}
	case "azure-native:sql/v20220201preview:Database":
		r = &Database{}
	case "azure-native:sql/v20220201preview:DatabaseAdvisor":
		r = &DatabaseAdvisor{}
	case "azure-native:sql/v20220201preview:DatabaseBlobAuditingPolicy":
		r = &DatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20220201preview:DatabaseSecurityAlertPolicy":
		r = &DatabaseSecurityAlertPolicy{}
	case "azure-native:sql/v20220201preview:DatabaseSqlVulnerabilityAssessmentRuleBaseline":
		r = &DatabaseSqlVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220201preview:DatabaseVulnerabilityAssessment":
		r = &DatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20220201preview:DatabaseVulnerabilityAssessmentRuleBaseline":
		r = &DatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220201preview:DistributedAvailabilityGroup":
		r = &DistributedAvailabilityGroup{}
	case "azure-native:sql/v20220201preview:ElasticPool":
		r = &ElasticPool{}
	case "azure-native:sql/v20220201preview:EncryptionProtector":
		r = &EncryptionProtector{}
	case "azure-native:sql/v20220201preview:ExtendedDatabaseBlobAuditingPolicy":
		r = &ExtendedDatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20220201preview:ExtendedServerBlobAuditingPolicy":
		r = &ExtendedServerBlobAuditingPolicy{}
	case "azure-native:sql/v20220201preview:FailoverGroup":
		r = &FailoverGroup{}
	case "azure-native:sql/v20220201preview:FirewallRule":
		r = &FirewallRule{}
	case "azure-native:sql/v20220201preview:GeoBackupPolicy":
		r = &GeoBackupPolicy{}
	case "azure-native:sql/v20220201preview:IPv6FirewallRule":
		r = &IPv6FirewallRule{}
	case "azure-native:sql/v20220201preview:InstanceFailoverGroup":
		r = &InstanceFailoverGroup{}
	case "azure-native:sql/v20220201preview:InstancePool":
		r = &InstancePool{}
	case "azure-native:sql/v20220201preview:Job":
		r = &Job{}
	case "azure-native:sql/v20220201preview:JobAgent":
		r = &JobAgent{}
	case "azure-native:sql/v20220201preview:JobCredential":
		r = &JobCredential{}
	case "azure-native:sql/v20220201preview:JobStep":
		r = &JobStep{}
	case "azure-native:sql/v20220201preview:JobTargetGroup":
		r = &JobTargetGroup{}
	case "azure-native:sql/v20220201preview:LongTermRetentionPolicy":
		r = &LongTermRetentionPolicy{}
	case "azure-native:sql/v20220201preview:ManagedDatabase":
		r = &ManagedDatabase{}
	case "azure-native:sql/v20220201preview:ManagedDatabaseSensitivityLabel":
		r = &ManagedDatabaseSensitivityLabel{}
	case "azure-native:sql/v20220201preview:ManagedDatabaseVulnerabilityAssessment":
		r = &ManagedDatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20220201preview:ManagedDatabaseVulnerabilityAssessmentRuleBaseline":
		r = &ManagedDatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220201preview:ManagedInstance":
		r = &ManagedInstance{}
	case "azure-native:sql/v20220201preview:ManagedInstanceAdministrator":
		r = &ManagedInstanceAdministrator{}
	case "azure-native:sql/v20220201preview:ManagedInstanceAzureADOnlyAuthentication":
		r = &ManagedInstanceAzureADOnlyAuthentication{}
	case "azure-native:sql/v20220201preview:ManagedInstanceKey":
		r = &ManagedInstanceKey{}
	case "azure-native:sql/v20220201preview:ManagedInstancePrivateEndpointConnection":
		r = &ManagedInstancePrivateEndpointConnection{}
	case "azure-native:sql/v20220201preview:ManagedInstanceVulnerabilityAssessment":
		r = &ManagedInstanceVulnerabilityAssessment{}
	case "azure-native:sql/v20220201preview:ManagedServerDnsAlias":
		r = &ManagedServerDnsAlias{}
	case "azure-native:sql/v20220201preview:OutboundFirewallRule":
		r = &OutboundFirewallRule{}
	case "azure-native:sql/v20220201preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:sql/v20220201preview:SensitivityLabel":
		r = &SensitivityLabel{}
	case "azure-native:sql/v20220201preview:Server":
		r = &Server{}
	case "azure-native:sql/v20220201preview:ServerAdvisor":
		r = &ServerAdvisor{}
	case "azure-native:sql/v20220201preview:ServerAzureADAdministrator":
		r = &ServerAzureADAdministrator{}
	case "azure-native:sql/v20220201preview:ServerAzureADOnlyAuthentication":
		r = &ServerAzureADOnlyAuthentication{}
	case "azure-native:sql/v20220201preview:ServerBlobAuditingPolicy":
		r = &ServerBlobAuditingPolicy{}
	case "azure-native:sql/v20220201preview:ServerDnsAlias":
		r = &ServerDnsAlias{}
	case "azure-native:sql/v20220201preview:ServerKey":
		r = &ServerKey{}
	case "azure-native:sql/v20220201preview:ServerSecurityAlertPolicy":
		r = &ServerSecurityAlertPolicy{}
	case "azure-native:sql/v20220201preview:ServerTrustCertificate":
		r = &ServerTrustCertificate{}
	case "azure-native:sql/v20220201preview:ServerTrustGroup":
		r = &ServerTrustGroup{}
	case "azure-native:sql/v20220201preview:ServerVulnerabilityAssessment":
		r = &ServerVulnerabilityAssessment{}
	case "azure-native:sql/v20220201preview:SqlVulnerabilityAssessmentRuleBaseline":
		r = &SqlVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220201preview:SqlVulnerabilityAssessmentsSetting":
		r = &SqlVulnerabilityAssessmentsSetting{}
	case "azure-native:sql/v20220201preview:SyncAgent":
		r = &SyncAgent{}
	case "azure-native:sql/v20220201preview:SyncGroup":
		r = &SyncGroup{}
	case "azure-native:sql/v20220201preview:SyncMember":
		r = &SyncMember{}
	case "azure-native:sql/v20220201preview:TransparentDataEncryption":
		r = &TransparentDataEncryption{}
	case "azure-native:sql/v20220201preview:VirtualNetworkRule":
		r = &VirtualNetworkRule{}
	case "azure-native:sql/v20220201preview:WorkloadClassifier":
		r = &WorkloadClassifier{}
	case "azure-native:sql/v20220201preview:WorkloadGroup":
		r = &WorkloadGroup{}
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
		"sql/v20220201preview",
		&module{version},
	)
}
