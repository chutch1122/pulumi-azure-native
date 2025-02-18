


package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppBackupConfiguration struct {
	pulumi.CustomResourceState

	BackupName        pulumi.StringPtrOutput                   `pulumi:"backupName"`
	BackupSchedule    BackupScheduleResponsePtrOutput          `pulumi:"backupSchedule"`
	Databases         DatabaseBackupSettingResponseArrayOutput `pulumi:"databases"`
	Enabled           pulumi.BoolPtrOutput                     `pulumi:"enabled"`
	Kind              pulumi.StringPtrOutput                   `pulumi:"kind"`
	Name              pulumi.StringOutput                      `pulumi:"name"`
	StorageAccountUrl pulumi.StringOutput                      `pulumi:"storageAccountUrl"`
	Type              pulumi.StringOutput                      `pulumi:"type"`
}


func NewWebAppBackupConfiguration(ctx *pulumi.Context,
	name string, args *WebAppBackupConfigurationArgs, opts ...pulumi.ResourceOption) (*WebAppBackupConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountUrl == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountUrl'")
	}
	if args.BackupSchedule != nil {
		args.BackupSchedule = args.BackupSchedule.ToBackupSchedulePtrOutput().ApplyT(func(v *BackupSchedule) *BackupSchedule { return v.Defaults() }).(BackupSchedulePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppBackupConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppBackupConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppBackupConfiguration
	err := ctx.RegisterResource("azure-native:web/v20181101:WebAppBackupConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppBackupConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppBackupConfigurationState, opts ...pulumi.ResourceOption) (*WebAppBackupConfiguration, error) {
	var resource WebAppBackupConfiguration
	err := ctx.ReadResource("azure-native:web/v20181101:WebAppBackupConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppBackupConfigurationState struct {
}

type WebAppBackupConfigurationState struct {
}

func (WebAppBackupConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppBackupConfigurationState)(nil)).Elem()
}

type webAppBackupConfigurationArgs struct {
	BackupName        *string                 `pulumi:"backupName"`
	BackupSchedule    *BackupSchedule         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSetting `pulumi:"databases"`
	Enabled           *bool                   `pulumi:"enabled"`
	Kind              *string                 `pulumi:"kind"`
	Name              string                  `pulumi:"name"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	StorageAccountUrl string                  `pulumi:"storageAccountUrl"`
}


type WebAppBackupConfigurationArgs struct {
	BackupName        pulumi.StringPtrInput
	BackupSchedule    BackupSchedulePtrInput
	Databases         DatabaseBackupSettingArrayInput
	Enabled           pulumi.BoolPtrInput
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	StorageAccountUrl pulumi.StringInput
}

func (WebAppBackupConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppBackupConfigurationArgs)(nil)).Elem()
}

type WebAppBackupConfigurationInput interface {
	pulumi.Input

	ToWebAppBackupConfigurationOutput() WebAppBackupConfigurationOutput
	ToWebAppBackupConfigurationOutputWithContext(ctx context.Context) WebAppBackupConfigurationOutput
}

func (*WebAppBackupConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppBackupConfiguration)(nil)).Elem()
}

func (i *WebAppBackupConfiguration) ToWebAppBackupConfigurationOutput() WebAppBackupConfigurationOutput {
	return i.ToWebAppBackupConfigurationOutputWithContext(context.Background())
}

func (i *WebAppBackupConfiguration) ToWebAppBackupConfigurationOutputWithContext(ctx context.Context) WebAppBackupConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppBackupConfigurationOutput)
}

type WebAppBackupConfigurationOutput struct{ *pulumi.OutputState }

func (WebAppBackupConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppBackupConfiguration)(nil)).Elem()
}

func (o WebAppBackupConfigurationOutput) ToWebAppBackupConfigurationOutput() WebAppBackupConfigurationOutput {
	return o
}

func (o WebAppBackupConfigurationOutput) ToWebAppBackupConfigurationOutputWithContext(ctx context.Context) WebAppBackupConfigurationOutput {
	return o
}

func (o WebAppBackupConfigurationOutput) BackupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.StringPtrOutput { return v.BackupName }).(pulumi.StringPtrOutput)
}

func (o WebAppBackupConfigurationOutput) BackupSchedule() BackupScheduleResponsePtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) BackupScheduleResponsePtrOutput { return v.BackupSchedule }).(BackupScheduleResponsePtrOutput)
}

func (o WebAppBackupConfigurationOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) DatabaseBackupSettingResponseArrayOutput { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o WebAppBackupConfigurationOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppBackupConfigurationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppBackupConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppBackupConfigurationOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.StringOutput { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

func (o WebAppBackupConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppBackupConfigurationOutput{})
}
