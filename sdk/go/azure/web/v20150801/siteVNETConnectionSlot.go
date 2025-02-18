


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteVNETConnectionSlot struct {
	pulumi.CustomResourceState

	CertBlob       pulumi.StringPtrOutput       `pulumi:"certBlob"`
	CertThumbprint pulumi.StringPtrOutput       `pulumi:"certThumbprint"`
	DnsServers     pulumi.StringPtrOutput       `pulumi:"dnsServers"`
	Kind           pulumi.StringPtrOutput       `pulumi:"kind"`
	Location       pulumi.StringOutput          `pulumi:"location"`
	Name           pulumi.StringPtrOutput       `pulumi:"name"`
	ResyncRequired pulumi.BoolPtrOutput         `pulumi:"resyncRequired"`
	Routes         VnetRouteResponseArrayOutput `pulumi:"routes"`
	Tags           pulumi.StringMapOutput       `pulumi:"tags"`
	Type           pulumi.StringPtrOutput       `pulumi:"type"`
	VnetResourceId pulumi.StringPtrOutput       `pulumi:"vnetResourceId"`
}


func NewSiteVNETConnectionSlot(ctx *pulumi.Context,
	name string, args *SiteVNETConnectionSlotArgs, opts ...pulumi.ResourceOption) (*SiteVNETConnectionSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:SiteVNETConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:SiteVNETConnectionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteVNETConnectionSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteVNETConnectionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteVNETConnectionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteVNETConnectionSlotState, opts ...pulumi.ResourceOption) (*SiteVNETConnectionSlot, error) {
	var resource SiteVNETConnectionSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteVNETConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteVNETConnectionSlotState struct {
}

type SiteVNETConnectionSlotState struct {
}

func (SiteVNETConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteVNETConnectionSlotState)(nil)).Elem()
}

type siteVNETConnectionSlotArgs struct {
	CertBlob          *string           `pulumi:"certBlob"`
	CertThumbprint    *string           `pulumi:"certThumbprint"`
	DnsServers        *string           `pulumi:"dnsServers"`
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResyncRequired    *bool             `pulumi:"resyncRequired"`
	Routes            []VnetRoute       `pulumi:"routes"`
	Slot              string            `pulumi:"slot"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
	VnetName          *string           `pulumi:"vnetName"`
	VnetResourceId    *string           `pulumi:"vnetResourceId"`
}


type SiteVNETConnectionSlotArgs struct {
	CertBlob          pulumi.StringPtrInput
	CertThumbprint    pulumi.StringPtrInput
	DnsServers        pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ResyncRequired    pulumi.BoolPtrInput
	Routes            VnetRouteArrayInput
	Slot              pulumi.StringInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
	VnetName          pulumi.StringPtrInput
	VnetResourceId    pulumi.StringPtrInput
}

func (SiteVNETConnectionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteVNETConnectionSlotArgs)(nil)).Elem()
}

type SiteVNETConnectionSlotInput interface {
	pulumi.Input

	ToSiteVNETConnectionSlotOutput() SiteVNETConnectionSlotOutput
	ToSiteVNETConnectionSlotOutputWithContext(ctx context.Context) SiteVNETConnectionSlotOutput
}

func (*SiteVNETConnectionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteVNETConnectionSlot)(nil)).Elem()
}

func (i *SiteVNETConnectionSlot) ToSiteVNETConnectionSlotOutput() SiteVNETConnectionSlotOutput {
	return i.ToSiteVNETConnectionSlotOutputWithContext(context.Background())
}

func (i *SiteVNETConnectionSlot) ToSiteVNETConnectionSlotOutputWithContext(ctx context.Context) SiteVNETConnectionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteVNETConnectionSlotOutput)
}

type SiteVNETConnectionSlotOutput struct{ *pulumi.OutputState }

func (SiteVNETConnectionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteVNETConnectionSlot)(nil)).Elem()
}

func (o SiteVNETConnectionSlotOutput) ToSiteVNETConnectionSlotOutput() SiteVNETConnectionSlotOutput {
	return o
}

func (o SiteVNETConnectionSlotOutput) ToSiteVNETConnectionSlotOutputWithContext(ctx context.Context) SiteVNETConnectionSlotOutput {
	return o
}

func (o SiteVNETConnectionSlotOutput) CertBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteVNETConnectionSlot) pulumi.StringPtrOutput { return v.CertBlob }).(pulumi.StringPtrOutput)
}

func (o SiteVNETConnectionSlotOutput) CertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteVNETConnectionSlot) pulumi.StringPtrOutput { return v.CertThumbprint }).(pulumi.StringPtrOutput)
}

func (o SiteVNETConnectionSlotOutput) DnsServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteVNETConnectionSlot) pulumi.StringPtrOutput { return v.DnsServers }).(pulumi.StringPtrOutput)
}

func (o SiteVNETConnectionSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteVNETConnectionSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteVNETConnectionSlotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteVNETConnectionSlot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteVNETConnectionSlotOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteVNETConnectionSlot) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteVNETConnectionSlotOutput) ResyncRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteVNETConnectionSlot) pulumi.BoolPtrOutput { return v.ResyncRequired }).(pulumi.BoolPtrOutput)
}

func (o SiteVNETConnectionSlotOutput) Routes() VnetRouteResponseArrayOutput {
	return o.ApplyT(func(v *SiteVNETConnectionSlot) VnetRouteResponseArrayOutput { return v.Routes }).(VnetRouteResponseArrayOutput)
}

func (o SiteVNETConnectionSlotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteVNETConnectionSlot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteVNETConnectionSlotOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteVNETConnectionSlot) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o SiteVNETConnectionSlotOutput) VnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteVNETConnectionSlot) pulumi.StringPtrOutput { return v.VnetResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteVNETConnectionSlotOutput{})
}
