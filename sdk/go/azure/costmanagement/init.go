


package costmanagement

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
	case "azure-native:costmanagement:CloudConnector":
		r = &CloudConnector{}
	case "azure-native:costmanagement:CostAllocationRule":
		r = &CostAllocationRule{}
	case "azure-native:costmanagement:Export":
		r = &Export{}
	case "azure-native:costmanagement:Report":
		r = &Report{}
	case "azure-native:costmanagement:ReportByBillingAccount":
		r = &ReportByBillingAccount{}
	case "azure-native:costmanagement:ReportByDepartment":
		r = &ReportByDepartment{}
	case "azure-native:costmanagement:ReportByResourceGroupName":
		r = &ReportByResourceGroupName{}
	case "azure-native:costmanagement:ScheduledAction":
		r = &ScheduledAction{}
	case "azure-native:costmanagement:ScheduledActionByScope":
		r = &ScheduledActionByScope{}
	case "azure-native:costmanagement:Setting":
		r = &Setting{}
	case "azure-native:costmanagement:View":
		r = &View{}
	case "azure-native:costmanagement:ViewByScope":
		r = &ViewByScope{}
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
		"costmanagement",
		&module{version},
	)
}
