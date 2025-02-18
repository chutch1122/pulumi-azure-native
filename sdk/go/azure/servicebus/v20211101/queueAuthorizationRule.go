


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type QueueAuthorizationRule struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput      `pulumi:"location"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	Rights     pulumi.StringArrayOutput `pulumi:"rights"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewQueueAuthorizationRule(ctx *pulumi.Context,
	name string, args *QueueAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*QueueAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.QueueName == nil {
		return nil, errors.New("invalid value for required argument 'QueueName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rights == nil {
		return nil, errors.New("invalid value for required argument 'Rights'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20140901:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:QueueAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20220101preview:QueueAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource QueueAuthorizationRule
	err := ctx.RegisterResource("azure-native:servicebus/v20211101:QueueAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetQueueAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueAuthorizationRuleState, opts ...pulumi.ResourceOption) (*QueueAuthorizationRule, error) {
	var resource QueueAuthorizationRule
	err := ctx.ReadResource("azure-native:servicebus/v20211101:QueueAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type queueAuthorizationRuleState struct {
}

type QueueAuthorizationRuleState struct {
}

func (QueueAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleState)(nil)).Elem()
}

type queueAuthorizationRuleArgs struct {
	AuthorizationRuleName *string        `pulumi:"authorizationRuleName"`
	NamespaceName         string         `pulumi:"namespaceName"`
	QueueName             string         `pulumi:"queueName"`
	ResourceGroupName     string         `pulumi:"resourceGroupName"`
	Rights                []AccessRights `pulumi:"rights"`
}


type QueueAuthorizationRuleArgs struct {
	AuthorizationRuleName pulumi.StringPtrInput
	NamespaceName         pulumi.StringInput
	QueueName             pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Rights                AccessRightsArrayInput
}

func (QueueAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueAuthorizationRuleArgs)(nil)).Elem()
}

type QueueAuthorizationRuleInput interface {
	pulumi.Input

	ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput
	ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput
}

func (*QueueAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueAuthorizationRule)(nil)).Elem()
}

func (i *QueueAuthorizationRule) ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput {
	return i.ToQueueAuthorizationRuleOutputWithContext(context.Background())
}

func (i *QueueAuthorizationRule) ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueAuthorizationRuleOutput)
}

type QueueAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (QueueAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueAuthorizationRule)(nil)).Elem()
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRuleOutput() QueueAuthorizationRuleOutput {
	return o
}

func (o QueueAuthorizationRuleOutput) ToQueueAuthorizationRuleOutputWithContext(ctx context.Context) QueueAuthorizationRuleOutput {
	return o
}

func (o QueueAuthorizationRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o QueueAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o QueueAuthorizationRuleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o QueueAuthorizationRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o QueueAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *QueueAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueAuthorizationRuleOutput{})
}
