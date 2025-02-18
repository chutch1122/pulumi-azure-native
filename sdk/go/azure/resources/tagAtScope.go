


package resources

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TagAtScope struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput `pulumi:"name"`
	Properties TagsResponseOutput  `pulumi:"properties"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewTagAtScope(ctx *pulumi.Context,
	name string, args *TagAtScopeArgs, opts ...pulumi.ResourceOption) (*TagAtScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources/v20191001:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:TagAtScope"),
		},
	})
	opts = append(opts, aliases)
	var resource TagAtScope
	err := ctx.RegisterResource("azure-native:resources:TagAtScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTagAtScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagAtScopeState, opts ...pulumi.ResourceOption) (*TagAtScope, error) {
	var resource TagAtScope
	err := ctx.ReadResource("azure-native:resources:TagAtScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type tagAtScopeState struct {
}

type TagAtScopeState struct {
}

func (TagAtScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagAtScopeState)(nil)).Elem()
}

type tagAtScopeArgs struct {
	Properties Tags   `pulumi:"properties"`
	Scope      string `pulumi:"scope"`
}


type TagAtScopeArgs struct {
	Properties TagsInput
	Scope      pulumi.StringInput
}

func (TagAtScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagAtScopeArgs)(nil)).Elem()
}

type TagAtScopeInput interface {
	pulumi.Input

	ToTagAtScopeOutput() TagAtScopeOutput
	ToTagAtScopeOutputWithContext(ctx context.Context) TagAtScopeOutput
}

func (*TagAtScope) ElementType() reflect.Type {
	return reflect.TypeOf((**TagAtScope)(nil)).Elem()
}

func (i *TagAtScope) ToTagAtScopeOutput() TagAtScopeOutput {
	return i.ToTagAtScopeOutputWithContext(context.Background())
}

func (i *TagAtScope) ToTagAtScopeOutputWithContext(ctx context.Context) TagAtScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagAtScopeOutput)
}

type TagAtScopeOutput struct{ *pulumi.OutputState }

func (TagAtScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagAtScope)(nil)).Elem()
}

func (o TagAtScopeOutput) ToTagAtScopeOutput() TagAtScopeOutput {
	return o
}

func (o TagAtScopeOutput) ToTagAtScopeOutputWithContext(ctx context.Context) TagAtScopeOutput {
	return o
}

func (o TagAtScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagAtScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TagAtScopeOutput) Properties() TagsResponseOutput {
	return o.ApplyT(func(v *TagAtScope) TagsResponseOutput { return v.Properties }).(TagsResponseOutput)
}

func (o TagAtScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TagAtScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TagAtScopeOutput{})
}
