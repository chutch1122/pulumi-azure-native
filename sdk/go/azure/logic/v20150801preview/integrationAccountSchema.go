


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountSchema struct {
	pulumi.CustomResourceState

	ChangedTime     pulumi.StringOutput                         `pulumi:"changedTime"`
	Content         pulumi.AnyOutput                            `pulumi:"content"`
	ContentLink     IntegrationAccountContentLinkResponseOutput `pulumi:"contentLink"`
	ContentType     pulumi.StringPtrOutput                      `pulumi:"contentType"`
	CreatedTime     pulumi.StringOutput                         `pulumi:"createdTime"`
	Location        pulumi.StringPtrOutput                      `pulumi:"location"`
	Metadata        pulumi.AnyOutput                            `pulumi:"metadata"`
	Name            pulumi.StringPtrOutput                      `pulumi:"name"`
	SchemaType      pulumi.StringPtrOutput                      `pulumi:"schemaType"`
	Tags            pulumi.StringMapOutput                      `pulumi:"tags"`
	TargetNamespace pulumi.StringPtrOutput                      `pulumi:"targetNamespace"`
	Type            pulumi.StringPtrOutput                      `pulumi:"type"`
}


func NewIntegrationAccountSchema(ctx *pulumi.Context,
	name string, args *IntegrationAccountSchemaArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountSchema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountSchema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountSchema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountSchema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountSchema"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountSchema
	err := ctx.RegisterResource("azure-native:logic/v20150801preview:IntegrationAccountSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountSchemaState, opts ...pulumi.ResourceOption) (*IntegrationAccountSchema, error) {
	var resource IntegrationAccountSchema
	err := ctx.ReadResource("azure-native:logic/v20150801preview:IntegrationAccountSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountSchemaState struct {
}

type IntegrationAccountSchemaState struct {
}

func (IntegrationAccountSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountSchemaState)(nil)).Elem()
}

type integrationAccountSchemaArgs struct {
	Content                interface{}       `pulumi:"content"`
	ContentType            *string           `pulumi:"contentType"`
	Id                     *string           `pulumi:"id"`
	IntegrationAccountName string            `pulumi:"integrationAccountName"`
	Location               *string           `pulumi:"location"`
	Metadata               interface{}       `pulumi:"metadata"`
	Name                   *string           `pulumi:"name"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	SchemaName             *string           `pulumi:"schemaName"`
	SchemaType             *SchemaType       `pulumi:"schemaType"`
	Tags                   map[string]string `pulumi:"tags"`
	TargetNamespace        *string           `pulumi:"targetNamespace"`
	Type                   *string           `pulumi:"type"`
}


type IntegrationAccountSchemaArgs struct {
	Content                pulumi.Input
	ContentType            pulumi.StringPtrInput
	Id                     pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	Metadata               pulumi.Input
	Name                   pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	SchemaName             pulumi.StringPtrInput
	SchemaType             SchemaTypePtrInput
	Tags                   pulumi.StringMapInput
	TargetNamespace        pulumi.StringPtrInput
	Type                   pulumi.StringPtrInput
}

func (IntegrationAccountSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountSchemaArgs)(nil)).Elem()
}

type IntegrationAccountSchemaInput interface {
	pulumi.Input

	ToIntegrationAccountSchemaOutput() IntegrationAccountSchemaOutput
	ToIntegrationAccountSchemaOutputWithContext(ctx context.Context) IntegrationAccountSchemaOutput
}

func (*IntegrationAccountSchema) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSchema)(nil)).Elem()
}

func (i *IntegrationAccountSchema) ToIntegrationAccountSchemaOutput() IntegrationAccountSchemaOutput {
	return i.ToIntegrationAccountSchemaOutputWithContext(context.Background())
}

func (i *IntegrationAccountSchema) ToIntegrationAccountSchemaOutputWithContext(ctx context.Context) IntegrationAccountSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountSchemaOutput)
}

type IntegrationAccountSchemaOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSchema)(nil)).Elem()
}

func (o IntegrationAccountSchemaOutput) ToIntegrationAccountSchemaOutput() IntegrationAccountSchemaOutput {
	return o
}

func (o IntegrationAccountSchemaOutput) ToIntegrationAccountSchemaOutputWithContext(ctx context.Context) IntegrationAccountSchemaOutput {
	return o
}

func (o IntegrationAccountSchemaOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o IntegrationAccountSchemaOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.AnyOutput { return v.Content }).(pulumi.AnyOutput)
}

func (o IntegrationAccountSchemaOutput) ContentLink() IntegrationAccountContentLinkResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) IntegrationAccountContentLinkResponseOutput { return v.ContentLink }).(IntegrationAccountContentLinkResponseOutput)
}

func (o IntegrationAccountSchemaOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountSchemaOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o IntegrationAccountSchemaOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountSchemaOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o IntegrationAccountSchemaOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountSchemaOutput) SchemaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.SchemaType }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountSchemaOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o IntegrationAccountSchemaOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountSchemaOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountSchemaOutput{})
}
