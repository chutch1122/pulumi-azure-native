


package storsimple

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
	case "azure-native:storsimple:AccessControlRecord":
		r = &AccessControlRecord{}
	case "azure-native:storsimple:BackupPolicy":
		r = &BackupPolicy{}
	case "azure-native:storsimple:BackupSchedule":
		r = &BackupSchedule{}
	case "azure-native:storsimple:BandwidthSetting":
		r = &BandwidthSetting{}
	case "azure-native:storsimple:Manager":
		r = &Manager{}
	case "azure-native:storsimple:ManagerExtendedInfo":
		r = &ManagerExtendedInfo{}
	case "azure-native:storsimple:StorageAccountCredential":
		r = &StorageAccountCredential{}
	case "azure-native:storsimple:Volume":
		r = &Volume{}
	case "azure-native:storsimple:VolumeContainer":
		r = &VolumeContainer{}
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
		"storsimple",
		&module{version},
	)
}
