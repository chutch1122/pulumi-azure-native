


package v20210401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobContainerImmutabilityPolicy struct {
	pulumi.CustomResourceState

	AllowProtectedAppendWrites            pulumi.BoolPtrOutput `pulumi:"allowProtectedAppendWrites"`
	Etag                                  pulumi.StringOutput  `pulumi:"etag"`
	ImmutabilityPeriodSinceCreationInDays pulumi.IntPtrOutput  `pulumi:"immutabilityPeriodSinceCreationInDays"`
	Name                                  pulumi.StringOutput  `pulumi:"name"`
	State                                 pulumi.StringOutput  `pulumi:"state"`
	Type                                  pulumi.StringOutput  `pulumi:"type"`
}


func NewBlobContainerImmutabilityPolicy(ctx *pulumi.Context,
	name string, args *BlobContainerImmutabilityPolicyArgs, opts ...pulumi.ResourceOption) (*BlobContainerImmutabilityPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180201:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180301preview:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180701:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20181101:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:BlobContainerImmutabilityPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobContainerImmutabilityPolicy
	err := ctx.RegisterResource("azure-native:storage/v20210401:BlobContainerImmutabilityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobContainerImmutabilityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobContainerImmutabilityPolicyState, opts ...pulumi.ResourceOption) (*BlobContainerImmutabilityPolicy, error) {
	var resource BlobContainerImmutabilityPolicy
	err := ctx.ReadResource("azure-native:storage/v20210401:BlobContainerImmutabilityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobContainerImmutabilityPolicyState struct {
}

type BlobContainerImmutabilityPolicyState struct {
}

func (BlobContainerImmutabilityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerImmutabilityPolicyState)(nil)).Elem()
}

type blobContainerImmutabilityPolicyArgs struct {
	AccountName                           string  `pulumi:"accountName"`
	AllowProtectedAppendWrites            *bool   `pulumi:"allowProtectedAppendWrites"`
	ContainerName                         string  `pulumi:"containerName"`
	ImmutabilityPeriodSinceCreationInDays *int    `pulumi:"immutabilityPeriodSinceCreationInDays"`
	ImmutabilityPolicyName                *string `pulumi:"immutabilityPolicyName"`
	ResourceGroupName                     string  `pulumi:"resourceGroupName"`
}


type BlobContainerImmutabilityPolicyArgs struct {
	AccountName                           pulumi.StringInput
	AllowProtectedAppendWrites            pulumi.BoolPtrInput
	ContainerName                         pulumi.StringInput
	ImmutabilityPeriodSinceCreationInDays pulumi.IntPtrInput
	ImmutabilityPolicyName                pulumi.StringPtrInput
	ResourceGroupName                     pulumi.StringInput
}

func (BlobContainerImmutabilityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerImmutabilityPolicyArgs)(nil)).Elem()
}

type BlobContainerImmutabilityPolicyInput interface {
	pulumi.Input

	ToBlobContainerImmutabilityPolicyOutput() BlobContainerImmutabilityPolicyOutput
	ToBlobContainerImmutabilityPolicyOutputWithContext(ctx context.Context) BlobContainerImmutabilityPolicyOutput
}

func (*BlobContainerImmutabilityPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainerImmutabilityPolicy)(nil)).Elem()
}

func (i *BlobContainerImmutabilityPolicy) ToBlobContainerImmutabilityPolicyOutput() BlobContainerImmutabilityPolicyOutput {
	return i.ToBlobContainerImmutabilityPolicyOutputWithContext(context.Background())
}

func (i *BlobContainerImmutabilityPolicy) ToBlobContainerImmutabilityPolicyOutputWithContext(ctx context.Context) BlobContainerImmutabilityPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobContainerImmutabilityPolicyOutput)
}

type BlobContainerImmutabilityPolicyOutput struct{ *pulumi.OutputState }

func (BlobContainerImmutabilityPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainerImmutabilityPolicy)(nil)).Elem()
}

func (o BlobContainerImmutabilityPolicyOutput) ToBlobContainerImmutabilityPolicyOutput() BlobContainerImmutabilityPolicyOutput {
	return o
}

func (o BlobContainerImmutabilityPolicyOutput) ToBlobContainerImmutabilityPolicyOutputWithContext(ctx context.Context) BlobContainerImmutabilityPolicyOutput {
	return o
}

func (o BlobContainerImmutabilityPolicyOutput) AllowProtectedAppendWrites() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlobContainerImmutabilityPolicy) pulumi.BoolPtrOutput { return v.AllowProtectedAppendWrites }).(pulumi.BoolPtrOutput)
}

func (o BlobContainerImmutabilityPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerImmutabilityPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BlobContainerImmutabilityPolicyOutput) ImmutabilityPeriodSinceCreationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BlobContainerImmutabilityPolicy) pulumi.IntPtrOutput {
		return v.ImmutabilityPeriodSinceCreationInDays
	}).(pulumi.IntPtrOutput)
}

func (o BlobContainerImmutabilityPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerImmutabilityPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BlobContainerImmutabilityPolicyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerImmutabilityPolicy) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o BlobContainerImmutabilityPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerImmutabilityPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobContainerImmutabilityPolicyOutput{})
}
