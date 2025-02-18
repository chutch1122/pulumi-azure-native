


package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteFunctionAppSettings(ctx *pulumi.Context, args *ListStaticSiteFunctionAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteFunctionAppSettingsResult, error) {
	var rv ListStaticSiteFunctionAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20200601:listStaticSiteFunctionAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteFunctionAppSettingsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteFunctionAppSettingsResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}

func ListStaticSiteFunctionAppSettingsOutput(ctx *pulumi.Context, args ListStaticSiteFunctionAppSettingsOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteFunctionAppSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteFunctionAppSettingsResult, error) {
			args := v.(ListStaticSiteFunctionAppSettingsArgs)
			r, err := ListStaticSiteFunctionAppSettings(ctx, &args, opts...)
			var s ListStaticSiteFunctionAppSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteFunctionAppSettingsResultOutput)
}

type ListStaticSiteFunctionAppSettingsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteFunctionAppSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteFunctionAppSettingsArgs)(nil)).Elem()
}


type ListStaticSiteFunctionAppSettingsResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteFunctionAppSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteFunctionAppSettingsResult)(nil)).Elem()
}

func (o ListStaticSiteFunctionAppSettingsResultOutput) ToListStaticSiteFunctionAppSettingsResultOutput() ListStaticSiteFunctionAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteFunctionAppSettingsResultOutput) ToListStaticSiteFunctionAppSettingsResultOutputWithContext(ctx context.Context) ListStaticSiteFunctionAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteFunctionAppSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteFunctionAppSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListStaticSiteFunctionAppSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListStaticSiteFunctionAppSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListStaticSiteFunctionAppSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteFunctionAppSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListStaticSiteFunctionAppSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListStaticSiteFunctionAppSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListStaticSiteFunctionAppSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteFunctionAppSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteFunctionAppSettingsResultOutput{})
}
