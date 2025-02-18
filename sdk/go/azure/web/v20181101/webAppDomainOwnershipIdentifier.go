


package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppDomainOwnershipIdentifier struct {
	pulumi.CustomResourceState

	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	Name pulumi.StringOutput    `pulumi:"name"`
	Type pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppDomainOwnershipIdentifier(ctx *pulumi.Context,
	name string, args *WebAppDomainOwnershipIdentifierArgs, opts ...pulumi.ResourceOption) (*WebAppDomainOwnershipIdentifier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppDomainOwnershipIdentifier"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppDomainOwnershipIdentifier
	err := ctx.RegisterResource("azure-native:web/v20181101:WebAppDomainOwnershipIdentifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppDomainOwnershipIdentifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppDomainOwnershipIdentifierState, opts ...pulumi.ResourceOption) (*WebAppDomainOwnershipIdentifier, error) {
	var resource WebAppDomainOwnershipIdentifier
	err := ctx.ReadResource("azure-native:web/v20181101:WebAppDomainOwnershipIdentifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppDomainOwnershipIdentifierState struct {
}

type WebAppDomainOwnershipIdentifierState struct {
}

func (WebAppDomainOwnershipIdentifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDomainOwnershipIdentifierState)(nil)).Elem()
}

type webAppDomainOwnershipIdentifierArgs struct {
	DomainOwnershipIdentifierName *string `pulumi:"domainOwnershipIdentifierName"`
	Id                            *string `pulumi:"id"`
	Kind                          *string `pulumi:"kind"`
	Name                          string  `pulumi:"name"`
	ResourceGroupName             string  `pulumi:"resourceGroupName"`
}


type WebAppDomainOwnershipIdentifierArgs struct {
	DomainOwnershipIdentifierName pulumi.StringPtrInput
	Id                            pulumi.StringPtrInput
	Kind                          pulumi.StringPtrInput
	Name                          pulumi.StringInput
	ResourceGroupName             pulumi.StringInput
}

func (WebAppDomainOwnershipIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDomainOwnershipIdentifierArgs)(nil)).Elem()
}

type WebAppDomainOwnershipIdentifierInput interface {
	pulumi.Input

	ToWebAppDomainOwnershipIdentifierOutput() WebAppDomainOwnershipIdentifierOutput
	ToWebAppDomainOwnershipIdentifierOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierOutput
}

func (*WebAppDomainOwnershipIdentifier) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppDomainOwnershipIdentifier)(nil)).Elem()
}

func (i *WebAppDomainOwnershipIdentifier) ToWebAppDomainOwnershipIdentifierOutput() WebAppDomainOwnershipIdentifierOutput {
	return i.ToWebAppDomainOwnershipIdentifierOutputWithContext(context.Background())
}

func (i *WebAppDomainOwnershipIdentifier) ToWebAppDomainOwnershipIdentifierOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppDomainOwnershipIdentifierOutput)
}

type WebAppDomainOwnershipIdentifierOutput struct{ *pulumi.OutputState }

func (WebAppDomainOwnershipIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppDomainOwnershipIdentifier)(nil)).Elem()
}

func (o WebAppDomainOwnershipIdentifierOutput) ToWebAppDomainOwnershipIdentifierOutput() WebAppDomainOwnershipIdentifierOutput {
	return o
}

func (o WebAppDomainOwnershipIdentifierOutput) ToWebAppDomainOwnershipIdentifierOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierOutput {
	return o
}

func (o WebAppDomainOwnershipIdentifierOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifier) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppDomainOwnershipIdentifierOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifier) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppDomainOwnershipIdentifierOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifier) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppDomainOwnershipIdentifierOutput{})
}
