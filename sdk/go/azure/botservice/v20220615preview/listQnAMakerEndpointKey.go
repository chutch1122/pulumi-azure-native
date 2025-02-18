


package v20220615preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListQnAMakerEndpointKey(ctx *pulumi.Context, args *ListQnAMakerEndpointKeyArgs, opts ...pulumi.InvokeOption) (*ListQnAMakerEndpointKeyResult, error) {
	var rv ListQnAMakerEndpointKeyResult
	err := ctx.Invoke("azure-native:botservice/v20220615preview:listQnAMakerEndpointKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListQnAMakerEndpointKeyArgs struct {
	Authkey  *string `pulumi:"authkey"`
	Hostname *string `pulumi:"hostname"`
}


type ListQnAMakerEndpointKeyResult struct {
	InstalledVersion     *string `pulumi:"installedVersion"`
	LastStableVersion    *string `pulumi:"lastStableVersion"`
	PrimaryEndpointKey   *string `pulumi:"primaryEndpointKey"`
	SecondaryEndpointKey *string `pulumi:"secondaryEndpointKey"`
}

func ListQnAMakerEndpointKeyOutput(ctx *pulumi.Context, args ListQnAMakerEndpointKeyOutputArgs, opts ...pulumi.InvokeOption) ListQnAMakerEndpointKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListQnAMakerEndpointKeyResult, error) {
			args := v.(ListQnAMakerEndpointKeyArgs)
			r, err := ListQnAMakerEndpointKey(ctx, &args, opts...)
			var s ListQnAMakerEndpointKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListQnAMakerEndpointKeyResultOutput)
}

type ListQnAMakerEndpointKeyOutputArgs struct {
	Authkey  pulumi.StringPtrInput `pulumi:"authkey"`
	Hostname pulumi.StringPtrInput `pulumi:"hostname"`
}

func (ListQnAMakerEndpointKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListQnAMakerEndpointKeyArgs)(nil)).Elem()
}


type ListQnAMakerEndpointKeyResultOutput struct{ *pulumi.OutputState }

func (ListQnAMakerEndpointKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListQnAMakerEndpointKeyResult)(nil)).Elem()
}

func (o ListQnAMakerEndpointKeyResultOutput) ToListQnAMakerEndpointKeyResultOutput() ListQnAMakerEndpointKeyResultOutput {
	return o
}

func (o ListQnAMakerEndpointKeyResultOutput) ToListQnAMakerEndpointKeyResultOutputWithContext(ctx context.Context) ListQnAMakerEndpointKeyResultOutput {
	return o
}

func (o ListQnAMakerEndpointKeyResultOutput) InstalledVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListQnAMakerEndpointKeyResult) *string { return v.InstalledVersion }).(pulumi.StringPtrOutput)
}

func (o ListQnAMakerEndpointKeyResultOutput) LastStableVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListQnAMakerEndpointKeyResult) *string { return v.LastStableVersion }).(pulumi.StringPtrOutput)
}

func (o ListQnAMakerEndpointKeyResultOutput) PrimaryEndpointKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListQnAMakerEndpointKeyResult) *string { return v.PrimaryEndpointKey }).(pulumi.StringPtrOutput)
}

func (o ListQnAMakerEndpointKeyResultOutput) SecondaryEndpointKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListQnAMakerEndpointKeyResult) *string { return v.SecondaryEndpointKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListQnAMakerEndpointKeyResultOutput{})
}
