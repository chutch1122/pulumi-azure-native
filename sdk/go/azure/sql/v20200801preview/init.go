


package v20200801preview

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
	case "azure-native:sql/v20200801preview:BackupShortTermRetentionPolicy":
		r = &BackupShortTermRetentionPolicy{}
	case "azure-native:sql/v20200801preview:Database":
		r = &Database{}
	case "azure-native:sql/v20200801preview:DatabaseAdvisor":
		r = &DatabaseAdvisor{}
	case "azure-native:sql/v20200801preview:DatabaseBlobAuditingPolicy":
		r = &DatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20200801preview:DatabaseSecurityAlertPolicy":
		r = &DatabaseSecurityAlertPolicy{}
	case "azure-native:sql/v20200801preview:DatabaseVulnerabilityAssessment":
		r = &DatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20200801preview:DatabaseVulnerabilityAssessmentRuleBaseline":
		r = &DatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20200801preview:ElasticPool":
		r = &ElasticPool{}
	case "azure-native:sql/v20200801preview:EncryptionProtector":
		r = &EncryptionProtector{}
	case "azure-native:sql/v20200801preview:ExtendedDatabaseBlobAuditingPolicy":
		r = &ExtendedDatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20200801preview:ExtendedServerBlobAuditingPolicy":
		r = &ExtendedServerBlobAuditingPolicy{}
	case "azure-native:sql/v20200801preview:FailoverGroup":
		r = &FailoverGroup{}
	case "azure-native:sql/v20200801preview:FirewallRule":
		r = &FirewallRule{}
	case "azure-native:sql/v20200801preview:InstanceFailoverGroup":
		r = &InstanceFailoverGroup{}
	case "azure-native:sql/v20200801preview:InstancePool":
		r = &InstancePool{}
	case "azure-native:sql/v20200801preview:Job":
		r = &Job{}
	case "azure-native:sql/v20200801preview:JobAgent":
		r = &JobAgent{}
	case "azure-native:sql/v20200801preview:JobCredential":
		r = &JobCredential{}
	case "azure-native:sql/v20200801preview:JobStep":
		r = &JobStep{}
	case "azure-native:sql/v20200801preview:JobTargetGroup":
		r = &JobTargetGroup{}
	case "azure-native:sql/v20200801preview:LongTermRetentionPolicy":
		r = &LongTermRetentionPolicy{}
	case "azure-native:sql/v20200801preview:ManagedDatabase":
		r = &ManagedDatabase{}
	case "azure-native:sql/v20200801preview:ManagedDatabaseSensitivityLabel":
		r = &ManagedDatabaseSensitivityLabel{}
	case "azure-native:sql/v20200801preview:ManagedDatabaseVulnerabilityAssessment":
		r = &ManagedDatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20200801preview:ManagedDatabaseVulnerabilityAssessmentRuleBaseline":
		r = &ManagedDatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20200801preview:ManagedInstance":
		r = &ManagedInstance{}
	case "azure-native:sql/v20200801preview:ManagedInstanceAdministrator":
		r = &ManagedInstanceAdministrator{}
	case "azure-native:sql/v20200801preview:ManagedInstanceAzureADOnlyAuthentication":
		r = &ManagedInstanceAzureADOnlyAuthentication{}
	case "azure-native:sql/v20200801preview:ManagedInstanceKey":
		r = &ManagedInstanceKey{}
	case "azure-native:sql/v20200801preview:ManagedInstancePrivateEndpointConnection":
		r = &ManagedInstancePrivateEndpointConnection{}
	case "azure-native:sql/v20200801preview:ManagedInstanceVulnerabilityAssessment":
		r = &ManagedInstanceVulnerabilityAssessment{}
	case "azure-native:sql/v20200801preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:sql/v20200801preview:SensitivityLabel":
		r = &SensitivityLabel{}
	case "azure-native:sql/v20200801preview:Server":
		r = &Server{}
	case "azure-native:sql/v20200801preview:ServerAdvisor":
		r = &ServerAdvisor{}
	case "azure-native:sql/v20200801preview:ServerAzureADAdministrator":
		r = &ServerAzureADAdministrator{}
	case "azure-native:sql/v20200801preview:ServerAzureADOnlyAuthentication":
		r = &ServerAzureADOnlyAuthentication{}
	case "azure-native:sql/v20200801preview:ServerBlobAuditingPolicy":
		r = &ServerBlobAuditingPolicy{}
	case "azure-native:sql/v20200801preview:ServerDnsAlias":
		r = &ServerDnsAlias{}
	case "azure-native:sql/v20200801preview:ServerKey":
		r = &ServerKey{}
	case "azure-native:sql/v20200801preview:ServerSecurityAlertPolicy":
		r = &ServerSecurityAlertPolicy{}
	case "azure-native:sql/v20200801preview:ServerTrustGroup":
		r = &ServerTrustGroup{}
	case "azure-native:sql/v20200801preview:ServerVulnerabilityAssessment":
		r = &ServerVulnerabilityAssessment{}
	case "azure-native:sql/v20200801preview:SyncAgent":
		r = &SyncAgent{}
	case "azure-native:sql/v20200801preview:SyncGroup":
		r = &SyncGroup{}
	case "azure-native:sql/v20200801preview:SyncMember":
		r = &SyncMember{}
	case "azure-native:sql/v20200801preview:TransparentDataEncryption":
		r = &TransparentDataEncryption{}
	case "azure-native:sql/v20200801preview:VirtualNetworkRule":
		r = &VirtualNetworkRule{}
	case "azure-native:sql/v20200801preview:WorkloadClassifier":
		r = &WorkloadClassifier{}
	case "azure-native:sql/v20200801preview:WorkloadGroup":
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
		"sql/v20200801preview",
		&module{version},
	)
}
