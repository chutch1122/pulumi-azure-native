


package databoxedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMECRole(ctx *pulumi.Context, args *LookupMECRoleArgs, opts ...pulumi.InvokeOption) (*LookupMECRoleResult, error) {
	var rv LookupMECRoleResult
	err := ctx.Invoke("azure-native:databoxedge:getMECRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMECRoleArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMECRoleResult struct {
	ConnectionString *AsymmetricEncryptedSecretResponse `pulumi:"connectionString"`
	Id               string                             `pulumi:"id"`
	Kind             string                             `pulumi:"kind"`
	Name             string                             `pulumi:"name"`
	RoleStatus       string                             `pulumi:"roleStatus"`
	SystemData       SystemDataResponse                 `pulumi:"systemData"`
	Type             string                             `pulumi:"type"`
}

func LookupMECRoleOutput(ctx *pulumi.Context, args LookupMECRoleOutputArgs, opts ...pulumi.InvokeOption) LookupMECRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMECRoleResult, error) {
			args := v.(LookupMECRoleArgs)
			r, err := LookupMECRole(ctx, &args, opts...)
			var s LookupMECRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMECRoleResultOutput)
}

type LookupMECRoleOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMECRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMECRoleArgs)(nil)).Elem()
}


type LookupMECRoleResultOutput struct{ *pulumi.OutputState }

func (LookupMECRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMECRoleResult)(nil)).Elem()
}

func (o LookupMECRoleResultOutput) ToLookupMECRoleResultOutput() LookupMECRoleResultOutput {
	return o
}

func (o LookupMECRoleResultOutput) ToLookupMECRoleResultOutputWithContext(ctx context.Context) LookupMECRoleResultOutput {
	return o
}

func (o LookupMECRoleResultOutput) ConnectionString() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v LookupMECRoleResult) *AsymmetricEncryptedSecretResponse { return v.ConnectionString }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o LookupMECRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMECRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMECRoleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMECRoleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupMECRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMECRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMECRoleResultOutput) RoleStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMECRoleResult) string { return v.RoleStatus }).(pulumi.StringOutput)
}

func (o LookupMECRoleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMECRoleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMECRoleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMECRoleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMECRoleResultOutput{})
}
