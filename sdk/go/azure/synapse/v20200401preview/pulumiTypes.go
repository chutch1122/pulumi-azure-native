


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SkuV3 struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SkuV3Input interface {
	pulumi.Input

	ToSkuV3Output() SkuV3Output
	ToSkuV3OutputWithContext(context.Context) SkuV3Output
}

type SkuV3Args struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuV3)(nil)).Elem()
}

func (i SkuV3Args) ToSkuV3Output() SkuV3Output {
	return i.ToSkuV3OutputWithContext(context.Background())
}

func (i SkuV3Args) ToSkuV3OutputWithContext(ctx context.Context) SkuV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(SkuV3Output)
}

func (i SkuV3Args) ToSkuV3PtrOutput() SkuV3PtrOutput {
	return i.ToSkuV3PtrOutputWithContext(context.Background())
}

func (i SkuV3Args) ToSkuV3PtrOutputWithContext(ctx context.Context) SkuV3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuV3Output).ToSkuV3PtrOutputWithContext(ctx)
}









type SkuV3PtrInput interface {
	pulumi.Input

	ToSkuV3PtrOutput() SkuV3PtrOutput
	ToSkuV3PtrOutputWithContext(context.Context) SkuV3PtrOutput
}

type skuV3PtrType SkuV3Args

func SkuV3Ptr(v *SkuV3Args) SkuV3PtrInput {
	return (*skuV3PtrType)(v)
}

func (*skuV3PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuV3)(nil)).Elem()
}

func (i *skuV3PtrType) ToSkuV3PtrOutput() SkuV3PtrOutput {
	return i.ToSkuV3PtrOutputWithContext(context.Background())
}

func (i *skuV3PtrType) ToSkuV3PtrOutputWithContext(ctx context.Context) SkuV3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuV3PtrOutput)
}

type SkuV3Output struct{ *pulumi.OutputState }

func (SkuV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuV3)(nil)).Elem()
}

func (o SkuV3Output) ToSkuV3Output() SkuV3Output {
	return o
}

func (o SkuV3Output) ToSkuV3OutputWithContext(ctx context.Context) SkuV3Output {
	return o
}

func (o SkuV3Output) ToSkuV3PtrOutput() SkuV3PtrOutput {
	return o.ToSkuV3PtrOutputWithContext(context.Background())
}

func (o SkuV3Output) ToSkuV3PtrOutputWithContext(ctx context.Context) SkuV3PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuV3) *SkuV3 {
		return &v
	}).(SkuV3PtrOutput)
}

func (o SkuV3Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuV3) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuV3Output) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuV3) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuV3PtrOutput struct{ *pulumi.OutputState }

func (SkuV3PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuV3)(nil)).Elem()
}

func (o SkuV3PtrOutput) ToSkuV3PtrOutput() SkuV3PtrOutput {
	return o
}

func (o SkuV3PtrOutput) ToSkuV3PtrOutputWithContext(ctx context.Context) SkuV3PtrOutput {
	return o
}

func (o SkuV3PtrOutput) Elem() SkuV3Output {
	return o.ApplyT(func(v *SkuV3) SkuV3 {
		if v != nil {
			return *v
		}
		var ret SkuV3
		return ret
	}).(SkuV3Output)
}

func (o SkuV3PtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuV3) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuV3PtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuV3) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuV3Response struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}

type SkuV3ResponseOutput struct{ *pulumi.OutputState }

func (SkuV3ResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuV3Response)(nil)).Elem()
}

func (o SkuV3ResponseOutput) ToSkuV3ResponseOutput() SkuV3ResponseOutput {
	return o
}

func (o SkuV3ResponseOutput) ToSkuV3ResponseOutputWithContext(ctx context.Context) SkuV3ResponseOutput {
	return o
}

func (o SkuV3ResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuV3Response) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuV3ResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuV3Response) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuV3ResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuV3ResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuV3Response)(nil)).Elem()
}

func (o SkuV3ResponsePtrOutput) ToSkuV3ResponsePtrOutput() SkuV3ResponsePtrOutput {
	return o
}

func (o SkuV3ResponsePtrOutput) ToSkuV3ResponsePtrOutputWithContext(ctx context.Context) SkuV3ResponsePtrOutput {
	return o
}

func (o SkuV3ResponsePtrOutput) Elem() SkuV3ResponseOutput {
	return o.ApplyT(func(v *SkuV3Response) SkuV3Response {
		if v != nil {
			return *v
		}
		var ret SkuV3Response
		return ret
	}).(SkuV3ResponseOutput)
}

func (o SkuV3ResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuV3Response) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuV3ResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuV3Response) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SqlDatabaseDataRetention struct {
	DropRetentionPeriod *string `pulumi:"dropRetentionPeriod"`
	RetentionPeriod     *string `pulumi:"retentionPeriod"`
}





type SqlDatabaseDataRetentionInput interface {
	pulumi.Input

	ToSqlDatabaseDataRetentionOutput() SqlDatabaseDataRetentionOutput
	ToSqlDatabaseDataRetentionOutputWithContext(context.Context) SqlDatabaseDataRetentionOutput
}

type SqlDatabaseDataRetentionArgs struct {
	DropRetentionPeriod pulumi.StringPtrInput `pulumi:"dropRetentionPeriod"`
	RetentionPeriod     pulumi.StringPtrInput `pulumi:"retentionPeriod"`
}

func (SqlDatabaseDataRetentionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseDataRetention)(nil)).Elem()
}

func (i SqlDatabaseDataRetentionArgs) ToSqlDatabaseDataRetentionOutput() SqlDatabaseDataRetentionOutput {
	return i.ToSqlDatabaseDataRetentionOutputWithContext(context.Background())
}

func (i SqlDatabaseDataRetentionArgs) ToSqlDatabaseDataRetentionOutputWithContext(ctx context.Context) SqlDatabaseDataRetentionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseDataRetentionOutput)
}

func (i SqlDatabaseDataRetentionArgs) ToSqlDatabaseDataRetentionPtrOutput() SqlDatabaseDataRetentionPtrOutput {
	return i.ToSqlDatabaseDataRetentionPtrOutputWithContext(context.Background())
}

func (i SqlDatabaseDataRetentionArgs) ToSqlDatabaseDataRetentionPtrOutputWithContext(ctx context.Context) SqlDatabaseDataRetentionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseDataRetentionOutput).ToSqlDatabaseDataRetentionPtrOutputWithContext(ctx)
}









type SqlDatabaseDataRetentionPtrInput interface {
	pulumi.Input

	ToSqlDatabaseDataRetentionPtrOutput() SqlDatabaseDataRetentionPtrOutput
	ToSqlDatabaseDataRetentionPtrOutputWithContext(context.Context) SqlDatabaseDataRetentionPtrOutput
}

type sqlDatabaseDataRetentionPtrType SqlDatabaseDataRetentionArgs

func SqlDatabaseDataRetentionPtr(v *SqlDatabaseDataRetentionArgs) SqlDatabaseDataRetentionPtrInput {
	return (*sqlDatabaseDataRetentionPtrType)(v)
}

func (*sqlDatabaseDataRetentionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabaseDataRetention)(nil)).Elem()
}

func (i *sqlDatabaseDataRetentionPtrType) ToSqlDatabaseDataRetentionPtrOutput() SqlDatabaseDataRetentionPtrOutput {
	return i.ToSqlDatabaseDataRetentionPtrOutputWithContext(context.Background())
}

func (i *sqlDatabaseDataRetentionPtrType) ToSqlDatabaseDataRetentionPtrOutputWithContext(ctx context.Context) SqlDatabaseDataRetentionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseDataRetentionPtrOutput)
}

type SqlDatabaseDataRetentionOutput struct{ *pulumi.OutputState }

func (SqlDatabaseDataRetentionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseDataRetention)(nil)).Elem()
}

func (o SqlDatabaseDataRetentionOutput) ToSqlDatabaseDataRetentionOutput() SqlDatabaseDataRetentionOutput {
	return o
}

func (o SqlDatabaseDataRetentionOutput) ToSqlDatabaseDataRetentionOutputWithContext(ctx context.Context) SqlDatabaseDataRetentionOutput {
	return o
}

func (o SqlDatabaseDataRetentionOutput) ToSqlDatabaseDataRetentionPtrOutput() SqlDatabaseDataRetentionPtrOutput {
	return o.ToSqlDatabaseDataRetentionPtrOutputWithContext(context.Background())
}

func (o SqlDatabaseDataRetentionOutput) ToSqlDatabaseDataRetentionPtrOutputWithContext(ctx context.Context) SqlDatabaseDataRetentionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlDatabaseDataRetention) *SqlDatabaseDataRetention {
		return &v
	}).(SqlDatabaseDataRetentionPtrOutput)
}

func (o SqlDatabaseDataRetentionOutput) DropRetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseDataRetention) *string { return v.DropRetentionPeriod }).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseDataRetentionOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseDataRetention) *string { return v.RetentionPeriod }).(pulumi.StringPtrOutput)
}

type SqlDatabaseDataRetentionPtrOutput struct{ *pulumi.OutputState }

func (SqlDatabaseDataRetentionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabaseDataRetention)(nil)).Elem()
}

func (o SqlDatabaseDataRetentionPtrOutput) ToSqlDatabaseDataRetentionPtrOutput() SqlDatabaseDataRetentionPtrOutput {
	return o
}

func (o SqlDatabaseDataRetentionPtrOutput) ToSqlDatabaseDataRetentionPtrOutputWithContext(ctx context.Context) SqlDatabaseDataRetentionPtrOutput {
	return o
}

func (o SqlDatabaseDataRetentionPtrOutput) Elem() SqlDatabaseDataRetentionOutput {
	return o.ApplyT(func(v *SqlDatabaseDataRetention) SqlDatabaseDataRetention {
		if v != nil {
			return *v
		}
		var ret SqlDatabaseDataRetention
		return ret
	}).(SqlDatabaseDataRetentionOutput)
}

func (o SqlDatabaseDataRetentionPtrOutput) DropRetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseDataRetention) *string {
		if v == nil {
			return nil
		}
		return v.DropRetentionPeriod
	}).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseDataRetentionPtrOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseDataRetention) *string {
		if v == nil {
			return nil
		}
		return v.RetentionPeriod
	}).(pulumi.StringPtrOutput)
}

type SqlDatabaseDataRetentionResponse struct {
	DropRetentionPeriod *string `pulumi:"dropRetentionPeriod"`
	RetentionPeriod     *string `pulumi:"retentionPeriod"`
}

type SqlDatabaseDataRetentionResponseOutput struct{ *pulumi.OutputState }

func (SqlDatabaseDataRetentionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseDataRetentionResponse)(nil)).Elem()
}

func (o SqlDatabaseDataRetentionResponseOutput) ToSqlDatabaseDataRetentionResponseOutput() SqlDatabaseDataRetentionResponseOutput {
	return o
}

func (o SqlDatabaseDataRetentionResponseOutput) ToSqlDatabaseDataRetentionResponseOutputWithContext(ctx context.Context) SqlDatabaseDataRetentionResponseOutput {
	return o
}

func (o SqlDatabaseDataRetentionResponseOutput) DropRetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseDataRetentionResponse) *string { return v.DropRetentionPeriod }).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseDataRetentionResponseOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseDataRetentionResponse) *string { return v.RetentionPeriod }).(pulumi.StringPtrOutput)
}

type SqlDatabaseDataRetentionResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlDatabaseDataRetentionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabaseDataRetentionResponse)(nil)).Elem()
}

func (o SqlDatabaseDataRetentionResponsePtrOutput) ToSqlDatabaseDataRetentionResponsePtrOutput() SqlDatabaseDataRetentionResponsePtrOutput {
	return o
}

func (o SqlDatabaseDataRetentionResponsePtrOutput) ToSqlDatabaseDataRetentionResponsePtrOutputWithContext(ctx context.Context) SqlDatabaseDataRetentionResponsePtrOutput {
	return o
}

func (o SqlDatabaseDataRetentionResponsePtrOutput) Elem() SqlDatabaseDataRetentionResponseOutput {
	return o.ApplyT(func(v *SqlDatabaseDataRetentionResponse) SqlDatabaseDataRetentionResponse {
		if v != nil {
			return *v
		}
		var ret SqlDatabaseDataRetentionResponse
		return ret
	}).(SqlDatabaseDataRetentionResponseOutput)
}

func (o SqlDatabaseDataRetentionResponsePtrOutput) DropRetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseDataRetentionResponse) *string {
		if v == nil {
			return nil
		}
		return v.DropRetentionPeriod
	}).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseDataRetentionResponsePtrOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseDataRetentionResponse) *string {
		if v == nil {
			return nil
		}
		return v.RetentionPeriod
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          string `pulumi:"createdAt"`
	CreatedBy          string `pulumi:"createdBy"`
	CreatedByType      string `pulumi:"createdByType"`
	LastModifiedAt     string `pulumi:"lastModifiedAt"`
	LastModifiedBy     string `pulumi:"lastModifiedBy"`
	LastModifiedByType string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedByType }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedByType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SkuV3Output{})
	pulumi.RegisterOutputType(SkuV3PtrOutput{})
	pulumi.RegisterOutputType(SkuV3ResponseOutput{})
	pulumi.RegisterOutputType(SkuV3ResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlDatabaseDataRetentionOutput{})
	pulumi.RegisterOutputType(SqlDatabaseDataRetentionPtrOutput{})
	pulumi.RegisterOutputType(SqlDatabaseDataRetentionResponseOutput{})
	pulumi.RegisterOutputType(SqlDatabaseDataRetentionResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
