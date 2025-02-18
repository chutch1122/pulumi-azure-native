


package v20140901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Namespace struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput            `pulumi:"location"`
	Name       pulumi.StringPtrOutput            `pulumi:"name"`
	Properties NamespacePropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput            `pulumi:"tags"`
	Type       pulumi.StringPtrOutput            `pulumi:"type"`
}


func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:notificationhubs:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20160301:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20170401:Namespace"),
		},
	})
	opts = append(opts, aliases)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:notificationhubs/v20140901:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:notificationhubs/v20140901:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type namespaceState struct {
}

type NamespaceState struct {
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	Location          *string             `pulumi:"location"`
	NamespaceName     *string             `pulumi:"namespaceName"`
	Properties        NamespaceProperties `pulumi:"properties"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	Tags              map[string]string   `pulumi:"tags"`
}


type NamespaceArgs struct {
	Location          pulumi.StringPtrInput
	NamespaceName     pulumi.StringPtrInput
	Properties        NamespacePropertiesInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

type NamespaceOutput struct{ *pulumi.OutputState }

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

func (o NamespaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NamespaceOutput) Properties() NamespacePropertiesResponseOutput {
	return o.ApplyT(func(v *Namespace) NamespacePropertiesResponseOutput { return v.Properties }).(NamespacePropertiesResponseOutput)
}

func (o NamespaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NamespaceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
}
