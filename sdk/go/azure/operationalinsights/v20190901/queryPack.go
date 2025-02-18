


package v20190901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type QueryPack struct {
	pulumi.CustomResourceState

	Location          pulumi.StringOutput    `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	QueryPackId       pulumi.StringOutput    `pulumi:"queryPackId"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	TimeCreated       pulumi.StringOutput    `pulumi:"timeCreated"`
	TimeModified      pulumi.StringOutput    `pulumi:"timeModified"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewQueryPack(ctx *pulumi.Context,
	name string, args *QueryPackArgs, opts ...pulumi.ResourceOption) (*QueryPack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:QueryPack"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190901preview:QueryPack"),
		},
	})
	opts = append(opts, aliases)
	var resource QueryPack
	err := ctx.RegisterResource("azure-native:operationalinsights/v20190901:QueryPack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetQueryPack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueryPackState, opts ...pulumi.ResourceOption) (*QueryPack, error) {
	var resource QueryPack
	err := ctx.ReadResource("azure-native:operationalinsights/v20190901:QueryPack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type queryPackState struct {
}

type QueryPackState struct {
}

func (QueryPackState) ElementType() reflect.Type {
	return reflect.TypeOf((*queryPackState)(nil)).Elem()
}

type queryPackArgs struct {
	Location          *string           `pulumi:"location"`
	QueryPackName     *string           `pulumi:"queryPackName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type QueryPackArgs struct {
	Location          pulumi.StringPtrInput
	QueryPackName     pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (QueryPackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queryPackArgs)(nil)).Elem()
}

type QueryPackInput interface {
	pulumi.Input

	ToQueryPackOutput() QueryPackOutput
	ToQueryPackOutputWithContext(ctx context.Context) QueryPackOutput
}

func (*QueryPack) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryPack)(nil)).Elem()
}

func (i *QueryPack) ToQueryPackOutput() QueryPackOutput {
	return i.ToQueryPackOutputWithContext(context.Background())
}

func (i *QueryPack) ToQueryPackOutputWithContext(ctx context.Context) QueryPackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryPackOutput)
}

type QueryPackOutput struct{ *pulumi.OutputState }

func (QueryPackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryPack)(nil)).Elem()
}

func (o QueryPackOutput) ToQueryPackOutput() QueryPackOutput {
	return o
}

func (o QueryPackOutput) ToQueryPackOutputWithContext(ctx context.Context) QueryPackOutput {
	return o
}

func (o QueryPackOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o QueryPackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o QueryPackOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o QueryPackOutput) QueryPackId() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.QueryPackId }).(pulumi.StringOutput)
}

func (o QueryPackOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o QueryPackOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o QueryPackOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.TimeModified }).(pulumi.StringOutput)
}

func (o QueryPackOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueryPackOutput{})
}
