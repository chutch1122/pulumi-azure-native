


package v20201101preview

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
	case "azure-native:sql/v20201101preview:BackupShortTermRetentionPolicy":
		r = &BackupShortTermRetentionPolicy{}
	case "azure-native:sql/v20201101preview:Database":
		r = &Database{}
	case "azure-native:sql/v20201101preview:DatabaseAdvisor":
		r = &DatabaseAdvisor{}
	case "azure-native:sql/v20201101preview:DatabaseBlobAuditingPolicy":
		r = &DatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20201101preview:DatabaseSecurityAlertPolicy":
		r = &DatabaseSecurityAlertPolicy{}
	case "azure-native:sql/v20201101preview:DatabaseVulnerabilityAssessment":
		r = &DatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20201101preview:DatabaseVulnerabilityAssessmentRuleBaseline":
		r = &DatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20201101preview:ElasticPool":
		r = &ElasticPool{}
	case "azure-native:sql/v20201101preview:EncryptionProtector":
		r = &EncryptionProtector{}
	case "azure-native:sql/v20201101preview:ExtendedDatabaseBlobAuditingPolicy":
		r = &ExtendedDatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20201101preview:ExtendedServerBlobAuditingPolicy":
		r = &ExtendedServerBlobAuditingPolicy{}
	case "azure-native:sql/v20201101preview:FailoverGroup":
		r = &FailoverGroup{}
	case "azure-native:sql/v20201101preview:FirewallRule":
		r = &FirewallRule{}
	case "azure-native:sql/v20201101preview:InstanceFailoverGroup":
		r = &InstanceFailoverGroup{}
	case "azure-native:sql/v20201101preview:InstancePool":
		r = &InstancePool{}
	case "azure-native:sql/v20201101preview:Job":
		r = &Job{}
	case "azure-native:sql/v20201101preview:JobAgent":
		r = &JobAgent{}
	case "azure-native:sql/v20201101preview:JobCredential":
		r = &JobCredential{}
	case "azure-native:sql/v20201101preview:JobStep":
		r = &JobStep{}
	case "azure-native:sql/v20201101preview:JobTargetGroup":
		r = &JobTargetGroup{}
	case "azure-native:sql/v20201101preview:LongTermRetentionPolicy":
		r = &LongTermRetentionPolicy{}
	case "azure-native:sql/v20201101preview:ManagedDatabase":
		r = &ManagedDatabase{}
	case "azure-native:sql/v20201101preview:ManagedDatabaseSensitivityLabel":
		r = &ManagedDatabaseSensitivityLabel{}
	case "azure-native:sql/v20201101preview:ManagedDatabaseVulnerabilityAssessment":
		r = &ManagedDatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20201101preview:ManagedDatabaseVulnerabilityAssessmentRuleBaseline":
		r = &ManagedDatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20201101preview:ManagedInstance":
		r = &ManagedInstance{}
	case "azure-native:sql/v20201101preview:ManagedInstanceAdministrator":
		r = &ManagedInstanceAdministrator{}
	case "azure-native:sql/v20201101preview:ManagedInstanceAzureADOnlyAuthentication":
		r = &ManagedInstanceAzureADOnlyAuthentication{}
	case "azure-native:sql/v20201101preview:ManagedInstanceKey":
		r = &ManagedInstanceKey{}
	case "azure-native:sql/v20201101preview:ManagedInstancePrivateEndpointConnection":
		r = &ManagedInstancePrivateEndpointConnection{}
	case "azure-native:sql/v20201101preview:ManagedInstanceVulnerabilityAssessment":
		r = &ManagedInstanceVulnerabilityAssessment{}
	case "azure-native:sql/v20201101preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:sql/v20201101preview:SensitivityLabel":
		r = &SensitivityLabel{}
	case "azure-native:sql/v20201101preview:Server":
		r = &Server{}
	case "azure-native:sql/v20201101preview:ServerAdvisor":
		r = &ServerAdvisor{}
	case "azure-native:sql/v20201101preview:ServerAzureADAdministrator":
		r = &ServerAzureADAdministrator{}
	case "azure-native:sql/v20201101preview:ServerAzureADOnlyAuthentication":
		r = &ServerAzureADOnlyAuthentication{}
	case "azure-native:sql/v20201101preview:ServerBlobAuditingPolicy":
		r = &ServerBlobAuditingPolicy{}
	case "azure-native:sql/v20201101preview:ServerDnsAlias":
		r = &ServerDnsAlias{}
	case "azure-native:sql/v20201101preview:ServerKey":
		r = &ServerKey{}
	case "azure-native:sql/v20201101preview:ServerSecurityAlertPolicy":
		r = &ServerSecurityAlertPolicy{}
	case "azure-native:sql/v20201101preview:ServerTrustGroup":
		r = &ServerTrustGroup{}
	case "azure-native:sql/v20201101preview:ServerVulnerabilityAssessment":
		r = &ServerVulnerabilityAssessment{}
	case "azure-native:sql/v20201101preview:SyncAgent":
		r = &SyncAgent{}
	case "azure-native:sql/v20201101preview:SyncGroup":
		r = &SyncGroup{}
	case "azure-native:sql/v20201101preview:SyncMember":
		r = &SyncMember{}
	case "azure-native:sql/v20201101preview:TransparentDataEncryption":
		r = &TransparentDataEncryption{}
	case "azure-native:sql/v20201101preview:VirtualNetworkRule":
		r = &VirtualNetworkRule{}
	case "azure-native:sql/v20201101preview:WorkloadClassifier":
		r = &WorkloadClassifier{}
	case "azure-native:sql/v20201101preview:WorkloadGroup":
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
		"sql/v20201101preview",
		&module{version},
	)
}
