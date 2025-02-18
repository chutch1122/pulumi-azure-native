


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MoveCollection struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput                    `pulumi:"etag"`
	Identity   IdentityResponsePtrOutput              `pulumi:"identity"`
	Location   pulumi.StringPtrOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                    `pulumi:"name"`
	Properties MoveCollectionPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput               `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                 `pulumi:"tags"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewMoveCollection(ctx *pulumi.Context,
	name string, args *MoveCollectionArgs, opts ...pulumi.ResourceOption) (*MoveCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate:MoveCollection"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20191001preview:MoveCollection"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20210101:MoveCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource MoveCollection
	err := ctx.RegisterResource("azure-native:migrate/v20210801:MoveCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMoveCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MoveCollectionState, opts ...pulumi.ResourceOption) (*MoveCollection, error) {
	var resource MoveCollection
	err := ctx.ReadResource("azure-native:migrate/v20210801:MoveCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type moveCollectionState struct {
}

type MoveCollectionState struct {
}

func (MoveCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*moveCollectionState)(nil)).Elem()
}

type moveCollectionArgs struct {
	Identity           *Identity                 `pulumi:"identity"`
	Location           *string                   `pulumi:"location"`
	MoveCollectionName *string                   `pulumi:"moveCollectionName"`
	Properties         *MoveCollectionProperties `pulumi:"properties"`
	ResourceGroupName  string                    `pulumi:"resourceGroupName"`
	Tags               map[string]string         `pulumi:"tags"`
}


type MoveCollectionArgs struct {
	Identity           IdentityPtrInput
	Location           pulumi.StringPtrInput
	MoveCollectionName pulumi.StringPtrInput
	Properties         MoveCollectionPropertiesPtrInput
	ResourceGroupName  pulumi.StringInput
	Tags               pulumi.StringMapInput
}

func (MoveCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moveCollectionArgs)(nil)).Elem()
}

type MoveCollectionInput interface {
	pulumi.Input

	ToMoveCollectionOutput() MoveCollectionOutput
	ToMoveCollectionOutputWithContext(ctx context.Context) MoveCollectionOutput
}

func (*MoveCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollection)(nil)).Elem()
}

func (i *MoveCollection) ToMoveCollectionOutput() MoveCollectionOutput {
	return i.ToMoveCollectionOutputWithContext(context.Background())
}

func (i *MoveCollection) ToMoveCollectionOutputWithContext(ctx context.Context) MoveCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionOutput)
}

type MoveCollectionOutput struct{ *pulumi.OutputState }

func (MoveCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollection)(nil)).Elem()
}

func (o MoveCollectionOutput) ToMoveCollectionOutput() MoveCollectionOutput {
	return o
}

func (o MoveCollectionOutput) ToMoveCollectionOutputWithContext(ctx context.Context) MoveCollectionOutput {
	return o
}

func (o MoveCollectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *MoveCollection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o MoveCollectionOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *MoveCollection) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o MoveCollectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o MoveCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MoveCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MoveCollectionOutput) Properties() MoveCollectionPropertiesResponseOutput {
	return o.ApplyT(func(v *MoveCollection) MoveCollectionPropertiesResponseOutput { return v.Properties }).(MoveCollectionPropertiesResponseOutput)
}

func (o MoveCollectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MoveCollection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MoveCollectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MoveCollection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MoveCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MoveCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MoveCollectionOutput{})
}
