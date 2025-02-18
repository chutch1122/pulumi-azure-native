


package v20150114privatepreview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IoMAMPolicyByName struct {
	pulumi.CustomResourceState

	AccessRecheckOfflineTimeout pulumi.StringPtrOutput `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  pulumi.StringPtrOutput `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         pulumi.StringPtrOutput `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           pulumi.StringPtrOutput `pulumi:"appSharingToLevel"`
	Authentication              pulumi.StringPtrOutput `pulumi:"authentication"`
	ClipboardSharingLevel       pulumi.StringPtrOutput `pulumi:"clipboardSharingLevel"`
	DataBackup                  pulumi.StringPtrOutput `pulumi:"dataBackup"`
	Description                 pulumi.StringPtrOutput `pulumi:"description"`
	DeviceCompliance            pulumi.StringPtrOutput `pulumi:"deviceCompliance"`
	FileEncryptionLevel         pulumi.StringPtrOutput `pulumi:"fileEncryptionLevel"`
	FileSharingSaveAs           pulumi.StringPtrOutput `pulumi:"fileSharingSaveAs"`
	FriendlyName                pulumi.StringOutput    `pulumi:"friendlyName"`
	GroupStatus                 pulumi.StringOutput    `pulumi:"groupStatus"`
	LastModifiedTime            pulumi.StringOutput    `pulumi:"lastModifiedTime"`
	Location                    pulumi.StringPtrOutput `pulumi:"location"`
	ManagedBrowser              pulumi.StringPtrOutput `pulumi:"managedBrowser"`
	Name                        pulumi.StringOutput    `pulumi:"name"`
	NumOfApps                   pulumi.IntOutput       `pulumi:"numOfApps"`
	OfflineWipeTimeout          pulumi.StringPtrOutput `pulumi:"offlineWipeTimeout"`
	Pin                         pulumi.StringPtrOutput `pulumi:"pin"`
	PinNumRetry                 pulumi.IntPtrOutput    `pulumi:"pinNumRetry"`
	Tags                        pulumi.StringMapOutput `pulumi:"tags"`
	TouchId                     pulumi.StringPtrOutput `pulumi:"touchId"`
	Type                        pulumi.StringOutput    `pulumi:"type"`
}


func NewIoMAMPolicyByName(ctx *pulumi.Context,
	name string, args *IoMAMPolicyByNameArgs, opts ...pulumi.ResourceOption) (*IoMAMPolicyByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FriendlyName == nil {
		return nil, errors.New("invalid value for required argument 'FriendlyName'")
	}
	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if isZero(args.AppSharingFromLevel) {
		args.AppSharingFromLevel = pulumi.StringPtr("none")
	}
	if isZero(args.AppSharingToLevel) {
		args.AppSharingToLevel = pulumi.StringPtr("none")
	}
	if isZero(args.Authentication) {
		args.Authentication = pulumi.StringPtr("required")
	}
	if isZero(args.ClipboardSharingLevel) {
		args.ClipboardSharingLevel = pulumi.StringPtr("blocked")
	}
	if isZero(args.DataBackup) {
		args.DataBackup = pulumi.StringPtr("allow")
	}
	if isZero(args.DeviceCompliance) {
		args.DeviceCompliance = pulumi.StringPtr("enable")
	}
	if isZero(args.FileEncryptionLevel) {
		args.FileEncryptionLevel = pulumi.StringPtr("deviceLocked")
	}
	if isZero(args.FileSharingSaveAs) {
		args.FileSharingSaveAs = pulumi.StringPtr("allow")
	}
	if isZero(args.ManagedBrowser) {
		args.ManagedBrowser = pulumi.StringPtr("required")
	}
	if isZero(args.Pin) {
		args.Pin = pulumi.StringPtr("required")
	}
	if isZero(args.TouchId) {
		args.TouchId = pulumi.StringPtr("enable")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:intune:IoMAMPolicyByName"),
		},
		{
			Type: pulumi.String("azure-native:intune/v20150114preview:IoMAMPolicyByName"),
		},
	})
	opts = append(opts, aliases)
	var resource IoMAMPolicyByName
	err := ctx.RegisterResource("azure-native:intune/v20150114privatepreview:IoMAMPolicyByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIoMAMPolicyByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IoMAMPolicyByNameState, opts ...pulumi.ResourceOption) (*IoMAMPolicyByName, error) {
	var resource IoMAMPolicyByName
	err := ctx.ReadResource("azure-native:intune/v20150114privatepreview:IoMAMPolicyByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ioMAMPolicyByNameState struct {
}

type IoMAMPolicyByNameState struct {
}

func (IoMAMPolicyByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*ioMAMPolicyByNameState)(nil)).Elem()
}

type ioMAMPolicyByNameArgs struct {
	AccessRecheckOfflineTimeout *string           `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  *string           `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         *string           `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           *string           `pulumi:"appSharingToLevel"`
	Authentication              *string           `pulumi:"authentication"`
	ClipboardSharingLevel       *string           `pulumi:"clipboardSharingLevel"`
	DataBackup                  *string           `pulumi:"dataBackup"`
	Description                 *string           `pulumi:"description"`
	DeviceCompliance            *string           `pulumi:"deviceCompliance"`
	FileEncryptionLevel         *string           `pulumi:"fileEncryptionLevel"`
	FileSharingSaveAs           *string           `pulumi:"fileSharingSaveAs"`
	FriendlyName                string            `pulumi:"friendlyName"`
	HostName                    string            `pulumi:"hostName"`
	Location                    *string           `pulumi:"location"`
	ManagedBrowser              *string           `pulumi:"managedBrowser"`
	OfflineWipeTimeout          *string           `pulumi:"offlineWipeTimeout"`
	Pin                         *string           `pulumi:"pin"`
	PinNumRetry                 *int              `pulumi:"pinNumRetry"`
	PolicyName                  *string           `pulumi:"policyName"`
	Tags                        map[string]string `pulumi:"tags"`
	TouchId                     *string           `pulumi:"touchId"`
}


type IoMAMPolicyByNameArgs struct {
	AccessRecheckOfflineTimeout pulumi.StringPtrInput
	AccessRecheckOnlineTimeout  pulumi.StringPtrInput
	AppSharingFromLevel         pulumi.StringPtrInput
	AppSharingToLevel           pulumi.StringPtrInput
	Authentication              pulumi.StringPtrInput
	ClipboardSharingLevel       pulumi.StringPtrInput
	DataBackup                  pulumi.StringPtrInput
	Description                 pulumi.StringPtrInput
	DeviceCompliance            pulumi.StringPtrInput
	FileEncryptionLevel         pulumi.StringPtrInput
	FileSharingSaveAs           pulumi.StringPtrInput
	FriendlyName                pulumi.StringInput
	HostName                    pulumi.StringInput
	Location                    pulumi.StringPtrInput
	ManagedBrowser              pulumi.StringPtrInput
	OfflineWipeTimeout          pulumi.StringPtrInput
	Pin                         pulumi.StringPtrInput
	PinNumRetry                 pulumi.IntPtrInput
	PolicyName                  pulumi.StringPtrInput
	Tags                        pulumi.StringMapInput
	TouchId                     pulumi.StringPtrInput
}

func (IoMAMPolicyByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ioMAMPolicyByNameArgs)(nil)).Elem()
}

type IoMAMPolicyByNameInput interface {
	pulumi.Input

	ToIoMAMPolicyByNameOutput() IoMAMPolicyByNameOutput
	ToIoMAMPolicyByNameOutputWithContext(ctx context.Context) IoMAMPolicyByNameOutput
}

func (*IoMAMPolicyByName) ElementType() reflect.Type {
	return reflect.TypeOf((**IoMAMPolicyByName)(nil)).Elem()
}

func (i *IoMAMPolicyByName) ToIoMAMPolicyByNameOutput() IoMAMPolicyByNameOutput {
	return i.ToIoMAMPolicyByNameOutputWithContext(context.Background())
}

func (i *IoMAMPolicyByName) ToIoMAMPolicyByNameOutputWithContext(ctx context.Context) IoMAMPolicyByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoMAMPolicyByNameOutput)
}

type IoMAMPolicyByNameOutput struct{ *pulumi.OutputState }

func (IoMAMPolicyByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IoMAMPolicyByName)(nil)).Elem()
}

func (o IoMAMPolicyByNameOutput) ToIoMAMPolicyByNameOutput() IoMAMPolicyByNameOutput {
	return o
}

func (o IoMAMPolicyByNameOutput) ToIoMAMPolicyByNameOutputWithContext(ctx context.Context) IoMAMPolicyByNameOutput {
	return o
}

func (o IoMAMPolicyByNameOutput) AccessRecheckOfflineTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.AccessRecheckOfflineTimeout }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) AccessRecheckOnlineTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.AccessRecheckOnlineTimeout }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) AppSharingFromLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.AppSharingFromLevel }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) AppSharingToLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.AppSharingToLevel }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.Authentication }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) ClipboardSharingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.ClipboardSharingLevel }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) DataBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.DataBackup }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) DeviceCompliance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.DeviceCompliance }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) FileEncryptionLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.FileEncryptionLevel }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) FileSharingSaveAs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.FileSharingSaveAs }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o IoMAMPolicyByNameOutput) GroupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringOutput { return v.GroupStatus }).(pulumi.StringOutput)
}

func (o IoMAMPolicyByNameOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o IoMAMPolicyByNameOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) ManagedBrowser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.ManagedBrowser }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IoMAMPolicyByNameOutput) NumOfApps() pulumi.IntOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.IntOutput { return v.NumOfApps }).(pulumi.IntOutput)
}

func (o IoMAMPolicyByNameOutput) OfflineWipeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.OfflineWipeTimeout }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) Pin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.Pin }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) PinNumRetry() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.IntPtrOutput { return v.PinNumRetry }).(pulumi.IntPtrOutput)
}

func (o IoMAMPolicyByNameOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o IoMAMPolicyByNameOutput) TouchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringPtrOutput { return v.TouchId }).(pulumi.StringPtrOutput)
}

func (o IoMAMPolicyByNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IoMAMPolicyByName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IoMAMPolicyByNameOutput{})
}
