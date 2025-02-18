


package v20190101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiIssue struct {
	pulumi.CustomResourceState

	ApiId       pulumi.StringPtrOutput `pulumi:"apiId"`
	CreatedDate pulumi.StringPtrOutput `pulumi:"createdDate"`
	Description pulumi.StringOutput    `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	State       pulumi.StringPtrOutput `pulumi:"state"`
	Title       pulumi.StringOutput    `pulumi:"title"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	UserId      pulumi.StringOutput    `pulumi:"userId"`
}


func NewApiIssue(ctx *pulumi.Context,
	name string, args *ApiIssueArgs, opts ...pulumi.ResourceOption) (*ApiIssue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ApiIssue"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiIssue
	err := ctx.RegisterResource("azure-native:apimanagement/v20190101:ApiIssue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiIssue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIssueState, opts ...pulumi.ResourceOption) (*ApiIssue, error) {
	var resource ApiIssue
	err := ctx.ReadResource("azure-native:apimanagement/v20190101:ApiIssue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiIssueState struct {
}

type ApiIssueState struct {
}

func (ApiIssueState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueState)(nil)).Elem()
}

type apiIssueArgs struct {
	ApiId             string  `pulumi:"apiId"`
	CreatedDate       *string `pulumi:"createdDate"`
	Description       string  `pulumi:"description"`
	IssueId           *string `pulumi:"issueId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	State             *string `pulumi:"state"`
	Title             string  `pulumi:"title"`
	UserId            string  `pulumi:"userId"`
}


type ApiIssueArgs struct {
	ApiId             pulumi.StringInput
	CreatedDate       pulumi.StringPtrInput
	Description       pulumi.StringInput
	IssueId           pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	State             pulumi.StringPtrInput
	Title             pulumi.StringInput
	UserId            pulumi.StringInput
}

func (ApiIssueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueArgs)(nil)).Elem()
}

type ApiIssueInput interface {
	pulumi.Input

	ToApiIssueOutput() ApiIssueOutput
	ToApiIssueOutputWithContext(ctx context.Context) ApiIssueOutput
}

func (*ApiIssue) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIssue)(nil)).Elem()
}

func (i *ApiIssue) ToApiIssueOutput() ApiIssueOutput {
	return i.ToApiIssueOutputWithContext(context.Background())
}

func (i *ApiIssue) ToApiIssueOutputWithContext(ctx context.Context) ApiIssueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIssueOutput)
}

type ApiIssueOutput struct{ *pulumi.OutputState }

func (ApiIssueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIssue)(nil)).Elem()
}

func (o ApiIssueOutput) ToApiIssueOutput() ApiIssueOutput {
	return o
}

func (o ApiIssueOutput) ToApiIssueOutputWithContext(ctx context.Context) ApiIssueOutput {
	return o
}

func (o ApiIssueOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringPtrOutput { return v.ApiId }).(pulumi.StringPtrOutput)
}

func (o ApiIssueOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringPtrOutput { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

func (o ApiIssueOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ApiIssueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApiIssueOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o ApiIssueOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

func (o ApiIssueOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApiIssueOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiIssueOutput{})
}
