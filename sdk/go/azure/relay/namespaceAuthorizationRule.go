


package relay

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NamespaceAuthorizationRule struct {
	pulumi.CustomResourceState

	Name   pulumi.StringOutput      `pulumi:"name"`
	Rights pulumi.StringArrayOutput `pulumi:"rights"`
	Type   pulumi.StringOutput      `pulumi:"type"`
}


func NewNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, args *NamespaceAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rights == nil {
		return nil, errors.New("invalid value for required argument 'Rights'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:relay/v20160701:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20170401:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20211101:NamespaceAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NamespaceAuthorizationRule
	err := ctx.RegisterResource("azure-native:relay:NamespaceAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceAuthorizationRuleState, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	var resource NamespaceAuthorizationRule
	err := ctx.ReadResource("azure-native:relay:NamespaceAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type namespaceAuthorizationRuleState struct {
}

type NamespaceAuthorizationRuleState struct {
}

func (NamespaceAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceAuthorizationRuleState)(nil)).Elem()
}

type namespaceAuthorizationRuleArgs struct {
	AuthorizationRuleName *string        `pulumi:"authorizationRuleName"`
	NamespaceName         string         `pulumi:"namespaceName"`
	ResourceGroupName     string         `pulumi:"resourceGroupName"`
	Rights                []AccessRights `pulumi:"rights"`
}


type NamespaceAuthorizationRuleArgs struct {
	AuthorizationRuleName pulumi.StringPtrInput
	NamespaceName         pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Rights                AccessRightsArrayInput
}

func (NamespaceAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceAuthorizationRuleArgs)(nil)).Elem()
}

type NamespaceAuthorizationRuleInput interface {
	pulumi.Input

	ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput
	ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput
}

func (*NamespaceAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceAuthorizationRule)(nil)).Elem()
}

func (i *NamespaceAuthorizationRule) ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput {
	return i.ToNamespaceAuthorizationRuleOutputWithContext(context.Background())
}

func (i *NamespaceAuthorizationRule) ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceAuthorizationRuleOutput)
}

type NamespaceAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (NamespaceAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceAuthorizationRule)(nil)).Elem()
}

func (o NamespaceAuthorizationRuleOutput) ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput {
	return o
}

func (o NamespaceAuthorizationRuleOutput) ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput {
	return o
}

func (o NamespaceAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NamespaceAuthorizationRuleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o NamespaceAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceAuthorizationRuleOutput{})
}
