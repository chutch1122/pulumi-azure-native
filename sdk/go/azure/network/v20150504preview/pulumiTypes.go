


package v20150504preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ARecord struct {
	Ipv4Address *string `pulumi:"ipv4Address"`
}





type ARecordInput interface {
	pulumi.Input

	ToARecordOutput() ARecordOutput
	ToARecordOutputWithContext(context.Context) ARecordOutput
}

type ARecordArgs struct {
	Ipv4Address pulumi.StringPtrInput `pulumi:"ipv4Address"`
}

func (ARecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecord)(nil)).Elem()
}

func (i ARecordArgs) ToARecordOutput() ARecordOutput {
	return i.ToARecordOutputWithContext(context.Background())
}

func (i ARecordArgs) ToARecordOutputWithContext(ctx context.Context) ARecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordOutput)
}





type ARecordArrayInput interface {
	pulumi.Input

	ToARecordArrayOutput() ARecordArrayOutput
	ToARecordArrayOutputWithContext(context.Context) ARecordArrayOutput
}

type ARecordArray []ARecordInput

func (ARecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ARecord)(nil)).Elem()
}

func (i ARecordArray) ToARecordArrayOutput() ARecordArrayOutput {
	return i.ToARecordArrayOutputWithContext(context.Background())
}

func (i ARecordArray) ToARecordArrayOutputWithContext(ctx context.Context) ARecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordArrayOutput)
}

type ARecordOutput struct{ *pulumi.OutputState }

func (ARecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecord)(nil)).Elem()
}

func (o ARecordOutput) ToARecordOutput() ARecordOutput {
	return o
}

func (o ARecordOutput) ToARecordOutputWithContext(ctx context.Context) ARecordOutput {
	return o
}

func (o ARecordOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ARecord) *string { return v.Ipv4Address }).(pulumi.StringPtrOutput)
}

type ARecordArrayOutput struct{ *pulumi.OutputState }

func (ARecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ARecord)(nil)).Elem()
}

func (o ARecordArrayOutput) ToARecordArrayOutput() ARecordArrayOutput {
	return o
}

func (o ARecordArrayOutput) ToARecordArrayOutputWithContext(ctx context.Context) ARecordArrayOutput {
	return o
}

func (o ARecordArrayOutput) Index(i pulumi.IntInput) ARecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ARecord {
		return vs[0].([]ARecord)[vs[1].(int)]
	}).(ARecordOutput)
}

type ARecordResponse struct {
	Ipv4Address *string `pulumi:"ipv4Address"`
}

type ARecordResponseOutput struct{ *pulumi.OutputState }

func (ARecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecordResponse)(nil)).Elem()
}

func (o ARecordResponseOutput) ToARecordResponseOutput() ARecordResponseOutput {
	return o
}

func (o ARecordResponseOutput) ToARecordResponseOutputWithContext(ctx context.Context) ARecordResponseOutput {
	return o
}

func (o ARecordResponseOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ARecordResponse) *string { return v.Ipv4Address }).(pulumi.StringPtrOutput)
}

type ARecordResponseArrayOutput struct{ *pulumi.OutputState }

func (ARecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ARecordResponse)(nil)).Elem()
}

func (o ARecordResponseArrayOutput) ToARecordResponseArrayOutput() ARecordResponseArrayOutput {
	return o
}

func (o ARecordResponseArrayOutput) ToARecordResponseArrayOutputWithContext(ctx context.Context) ARecordResponseArrayOutput {
	return o
}

func (o ARecordResponseArrayOutput) Index(i pulumi.IntInput) ARecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ARecordResponse {
		return vs[0].([]ARecordResponse)[vs[1].(int)]
	}).(ARecordResponseOutput)
}

type AaaaRecord struct {
	Ipv6Address *string `pulumi:"ipv6Address"`
}





type AaaaRecordInput interface {
	pulumi.Input

	ToAaaaRecordOutput() AaaaRecordOutput
	ToAaaaRecordOutputWithContext(context.Context) AaaaRecordOutput
}

type AaaaRecordArgs struct {
	Ipv6Address pulumi.StringPtrInput `pulumi:"ipv6Address"`
}

func (AaaaRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AaaaRecord)(nil)).Elem()
}

func (i AaaaRecordArgs) ToAaaaRecordOutput() AaaaRecordOutput {
	return i.ToAaaaRecordOutputWithContext(context.Background())
}

func (i AaaaRecordArgs) ToAaaaRecordOutputWithContext(ctx context.Context) AaaaRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AaaaRecordOutput)
}





type AaaaRecordArrayInput interface {
	pulumi.Input

	ToAaaaRecordArrayOutput() AaaaRecordArrayOutput
	ToAaaaRecordArrayOutputWithContext(context.Context) AaaaRecordArrayOutput
}

type AaaaRecordArray []AaaaRecordInput

func (AaaaRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AaaaRecord)(nil)).Elem()
}

func (i AaaaRecordArray) ToAaaaRecordArrayOutput() AaaaRecordArrayOutput {
	return i.ToAaaaRecordArrayOutputWithContext(context.Background())
}

func (i AaaaRecordArray) ToAaaaRecordArrayOutputWithContext(ctx context.Context) AaaaRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AaaaRecordArrayOutput)
}

type AaaaRecordOutput struct{ *pulumi.OutputState }

func (AaaaRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AaaaRecord)(nil)).Elem()
}

func (o AaaaRecordOutput) ToAaaaRecordOutput() AaaaRecordOutput {
	return o
}

func (o AaaaRecordOutput) ToAaaaRecordOutputWithContext(ctx context.Context) AaaaRecordOutput {
	return o
}

func (o AaaaRecordOutput) Ipv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AaaaRecord) *string { return v.Ipv6Address }).(pulumi.StringPtrOutput)
}

type AaaaRecordArrayOutput struct{ *pulumi.OutputState }

func (AaaaRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AaaaRecord)(nil)).Elem()
}

func (o AaaaRecordArrayOutput) ToAaaaRecordArrayOutput() AaaaRecordArrayOutput {
	return o
}

func (o AaaaRecordArrayOutput) ToAaaaRecordArrayOutputWithContext(ctx context.Context) AaaaRecordArrayOutput {
	return o
}

func (o AaaaRecordArrayOutput) Index(i pulumi.IntInput) AaaaRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AaaaRecord {
		return vs[0].([]AaaaRecord)[vs[1].(int)]
	}).(AaaaRecordOutput)
}

type AaaaRecordResponse struct {
	Ipv6Address *string `pulumi:"ipv6Address"`
}

type AaaaRecordResponseOutput struct{ *pulumi.OutputState }

func (AaaaRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AaaaRecordResponse)(nil)).Elem()
}

func (o AaaaRecordResponseOutput) ToAaaaRecordResponseOutput() AaaaRecordResponseOutput {
	return o
}

func (o AaaaRecordResponseOutput) ToAaaaRecordResponseOutputWithContext(ctx context.Context) AaaaRecordResponseOutput {
	return o
}

func (o AaaaRecordResponseOutput) Ipv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AaaaRecordResponse) *string { return v.Ipv6Address }).(pulumi.StringPtrOutput)
}

type AaaaRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (AaaaRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AaaaRecordResponse)(nil)).Elem()
}

func (o AaaaRecordResponseArrayOutput) ToAaaaRecordResponseArrayOutput() AaaaRecordResponseArrayOutput {
	return o
}

func (o AaaaRecordResponseArrayOutput) ToAaaaRecordResponseArrayOutputWithContext(ctx context.Context) AaaaRecordResponseArrayOutput {
	return o
}

func (o AaaaRecordResponseArrayOutput) Index(i pulumi.IntInput) AaaaRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AaaaRecordResponse {
		return vs[0].([]AaaaRecordResponse)[vs[1].(int)]
	}).(AaaaRecordResponseOutput)
}

type CnameRecord struct {
	Cname *string `pulumi:"cname"`
}





type CnameRecordInput interface {
	pulumi.Input

	ToCnameRecordOutput() CnameRecordOutput
	ToCnameRecordOutputWithContext(context.Context) CnameRecordOutput
}

type CnameRecordArgs struct {
	Cname pulumi.StringPtrInput `pulumi:"cname"`
}

func (CnameRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CnameRecord)(nil)).Elem()
}

func (i CnameRecordArgs) ToCnameRecordOutput() CnameRecordOutput {
	return i.ToCnameRecordOutputWithContext(context.Background())
}

func (i CnameRecordArgs) ToCnameRecordOutputWithContext(ctx context.Context) CnameRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordOutput)
}

func (i CnameRecordArgs) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return i.ToCnameRecordPtrOutputWithContext(context.Background())
}

func (i CnameRecordArgs) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordOutput).ToCnameRecordPtrOutputWithContext(ctx)
}









type CnameRecordPtrInput interface {
	pulumi.Input

	ToCnameRecordPtrOutput() CnameRecordPtrOutput
	ToCnameRecordPtrOutputWithContext(context.Context) CnameRecordPtrOutput
}

type cnameRecordPtrType CnameRecordArgs

func CnameRecordPtr(v *CnameRecordArgs) CnameRecordPtrInput {
	return (*cnameRecordPtrType)(v)
}

func (*cnameRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecord)(nil)).Elem()
}

func (i *cnameRecordPtrType) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return i.ToCnameRecordPtrOutputWithContext(context.Background())
}

func (i *cnameRecordPtrType) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordPtrOutput)
}

type CnameRecordOutput struct{ *pulumi.OutputState }

func (CnameRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CnameRecord)(nil)).Elem()
}

func (o CnameRecordOutput) ToCnameRecordOutput() CnameRecordOutput {
	return o
}

func (o CnameRecordOutput) ToCnameRecordOutputWithContext(ctx context.Context) CnameRecordOutput {
	return o
}

func (o CnameRecordOutput) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return o.ToCnameRecordPtrOutputWithContext(context.Background())
}

func (o CnameRecordOutput) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CnameRecord) *CnameRecord {
		return &v
	}).(CnameRecordPtrOutput)
}

func (o CnameRecordOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CnameRecord) *string { return v.Cname }).(pulumi.StringPtrOutput)
}

type CnameRecordPtrOutput struct{ *pulumi.OutputState }

func (CnameRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecord)(nil)).Elem()
}

func (o CnameRecordPtrOutput) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return o
}

func (o CnameRecordPtrOutput) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return o
}

func (o CnameRecordPtrOutput) Elem() CnameRecordOutput {
	return o.ApplyT(func(v *CnameRecord) CnameRecord {
		if v != nil {
			return *v
		}
		var ret CnameRecord
		return ret
	}).(CnameRecordOutput)
}

func (o CnameRecordPtrOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CnameRecord) *string {
		if v == nil {
			return nil
		}
		return v.Cname
	}).(pulumi.StringPtrOutput)
}

type CnameRecordResponse struct {
	Cname *string `pulumi:"cname"`
}

type CnameRecordResponseOutput struct{ *pulumi.OutputState }

func (CnameRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CnameRecordResponse)(nil)).Elem()
}

func (o CnameRecordResponseOutput) ToCnameRecordResponseOutput() CnameRecordResponseOutput {
	return o
}

func (o CnameRecordResponseOutput) ToCnameRecordResponseOutputWithContext(ctx context.Context) CnameRecordResponseOutput {
	return o
}

func (o CnameRecordResponseOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CnameRecordResponse) *string { return v.Cname }).(pulumi.StringPtrOutput)
}

type CnameRecordResponsePtrOutput struct{ *pulumi.OutputState }

func (CnameRecordResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecordResponse)(nil)).Elem()
}

func (o CnameRecordResponsePtrOutput) ToCnameRecordResponsePtrOutput() CnameRecordResponsePtrOutput {
	return o
}

func (o CnameRecordResponsePtrOutput) ToCnameRecordResponsePtrOutputWithContext(ctx context.Context) CnameRecordResponsePtrOutput {
	return o
}

func (o CnameRecordResponsePtrOutput) Elem() CnameRecordResponseOutput {
	return o.ApplyT(func(v *CnameRecordResponse) CnameRecordResponse {
		if v != nil {
			return *v
		}
		var ret CnameRecordResponse
		return ret
	}).(CnameRecordResponseOutput)
}

func (o CnameRecordResponsePtrOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CnameRecordResponse) *string {
		if v == nil {
			return nil
		}
		return v.Cname
	}).(pulumi.StringPtrOutput)
}

type MxRecord struct {
	Exchange   *string `pulumi:"exchange"`
	Preference *int    `pulumi:"preference"`
}





type MxRecordInput interface {
	pulumi.Input

	ToMxRecordOutput() MxRecordOutput
	ToMxRecordOutputWithContext(context.Context) MxRecordOutput
}

type MxRecordArgs struct {
	Exchange   pulumi.StringPtrInput `pulumi:"exchange"`
	Preference pulumi.IntPtrInput    `pulumi:"preference"`
}

func (MxRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecord)(nil)).Elem()
}

func (i MxRecordArgs) ToMxRecordOutput() MxRecordOutput {
	return i.ToMxRecordOutputWithContext(context.Background())
}

func (i MxRecordArgs) ToMxRecordOutputWithContext(ctx context.Context) MxRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordOutput)
}





type MxRecordArrayInput interface {
	pulumi.Input

	ToMxRecordArrayOutput() MxRecordArrayOutput
	ToMxRecordArrayOutputWithContext(context.Context) MxRecordArrayOutput
}

type MxRecordArray []MxRecordInput

func (MxRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecord)(nil)).Elem()
}

func (i MxRecordArray) ToMxRecordArrayOutput() MxRecordArrayOutput {
	return i.ToMxRecordArrayOutputWithContext(context.Background())
}

func (i MxRecordArray) ToMxRecordArrayOutputWithContext(ctx context.Context) MxRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordArrayOutput)
}

type MxRecordOutput struct{ *pulumi.OutputState }

func (MxRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecord)(nil)).Elem()
}

func (o MxRecordOutput) ToMxRecordOutput() MxRecordOutput {
	return o
}

func (o MxRecordOutput) ToMxRecordOutputWithContext(ctx context.Context) MxRecordOutput {
	return o
}

func (o MxRecordOutput) Exchange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MxRecord) *string { return v.Exchange }).(pulumi.StringPtrOutput)
}

func (o MxRecordOutput) Preference() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MxRecord) *int { return v.Preference }).(pulumi.IntPtrOutput)
}

type MxRecordArrayOutput struct{ *pulumi.OutputState }

func (MxRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecord)(nil)).Elem()
}

func (o MxRecordArrayOutput) ToMxRecordArrayOutput() MxRecordArrayOutput {
	return o
}

func (o MxRecordArrayOutput) ToMxRecordArrayOutputWithContext(ctx context.Context) MxRecordArrayOutput {
	return o
}

func (o MxRecordArrayOutput) Index(i pulumi.IntInput) MxRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MxRecord {
		return vs[0].([]MxRecord)[vs[1].(int)]
	}).(MxRecordOutput)
}

type MxRecordResponse struct {
	Exchange   *string `pulumi:"exchange"`
	Preference *int    `pulumi:"preference"`
}

type MxRecordResponseOutput struct{ *pulumi.OutputState }

func (MxRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecordResponse)(nil)).Elem()
}

func (o MxRecordResponseOutput) ToMxRecordResponseOutput() MxRecordResponseOutput {
	return o
}

func (o MxRecordResponseOutput) ToMxRecordResponseOutputWithContext(ctx context.Context) MxRecordResponseOutput {
	return o
}

func (o MxRecordResponseOutput) Exchange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MxRecordResponse) *string { return v.Exchange }).(pulumi.StringPtrOutput)
}

func (o MxRecordResponseOutput) Preference() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MxRecordResponse) *int { return v.Preference }).(pulumi.IntPtrOutput)
}

type MxRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (MxRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecordResponse)(nil)).Elem()
}

func (o MxRecordResponseArrayOutput) ToMxRecordResponseArrayOutput() MxRecordResponseArrayOutput {
	return o
}

func (o MxRecordResponseArrayOutput) ToMxRecordResponseArrayOutputWithContext(ctx context.Context) MxRecordResponseArrayOutput {
	return o
}

func (o MxRecordResponseArrayOutput) Index(i pulumi.IntInput) MxRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MxRecordResponse {
		return vs[0].([]MxRecordResponse)[vs[1].(int)]
	}).(MxRecordResponseOutput)
}

type NsRecord struct {
	Nsdname *string `pulumi:"nsdname"`
}





type NsRecordInput interface {
	pulumi.Input

	ToNsRecordOutput() NsRecordOutput
	ToNsRecordOutputWithContext(context.Context) NsRecordOutput
}

type NsRecordArgs struct {
	Nsdname pulumi.StringPtrInput `pulumi:"nsdname"`
}

func (NsRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NsRecord)(nil)).Elem()
}

func (i NsRecordArgs) ToNsRecordOutput() NsRecordOutput {
	return i.ToNsRecordOutputWithContext(context.Background())
}

func (i NsRecordArgs) ToNsRecordOutputWithContext(ctx context.Context) NsRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsRecordOutput)
}





type NsRecordArrayInput interface {
	pulumi.Input

	ToNsRecordArrayOutput() NsRecordArrayOutput
	ToNsRecordArrayOutputWithContext(context.Context) NsRecordArrayOutput
}

type NsRecordArray []NsRecordInput

func (NsRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NsRecord)(nil)).Elem()
}

func (i NsRecordArray) ToNsRecordArrayOutput() NsRecordArrayOutput {
	return i.ToNsRecordArrayOutputWithContext(context.Background())
}

func (i NsRecordArray) ToNsRecordArrayOutputWithContext(ctx context.Context) NsRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsRecordArrayOutput)
}

type NsRecordOutput struct{ *pulumi.OutputState }

func (NsRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsRecord)(nil)).Elem()
}

func (o NsRecordOutput) ToNsRecordOutput() NsRecordOutput {
	return o
}

func (o NsRecordOutput) ToNsRecordOutputWithContext(ctx context.Context) NsRecordOutput {
	return o
}

func (o NsRecordOutput) Nsdname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsRecord) *string { return v.Nsdname }).(pulumi.StringPtrOutput)
}

type NsRecordArrayOutput struct{ *pulumi.OutputState }

func (NsRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NsRecord)(nil)).Elem()
}

func (o NsRecordArrayOutput) ToNsRecordArrayOutput() NsRecordArrayOutput {
	return o
}

func (o NsRecordArrayOutput) ToNsRecordArrayOutputWithContext(ctx context.Context) NsRecordArrayOutput {
	return o
}

func (o NsRecordArrayOutput) Index(i pulumi.IntInput) NsRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NsRecord {
		return vs[0].([]NsRecord)[vs[1].(int)]
	}).(NsRecordOutput)
}

type NsRecordResponse struct {
	Nsdname *string `pulumi:"nsdname"`
}

type NsRecordResponseOutput struct{ *pulumi.OutputState }

func (NsRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsRecordResponse)(nil)).Elem()
}

func (o NsRecordResponseOutput) ToNsRecordResponseOutput() NsRecordResponseOutput {
	return o
}

func (o NsRecordResponseOutput) ToNsRecordResponseOutputWithContext(ctx context.Context) NsRecordResponseOutput {
	return o
}

func (o NsRecordResponseOutput) Nsdname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsRecordResponse) *string { return v.Nsdname }).(pulumi.StringPtrOutput)
}

type NsRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (NsRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NsRecordResponse)(nil)).Elem()
}

func (o NsRecordResponseArrayOutput) ToNsRecordResponseArrayOutput() NsRecordResponseArrayOutput {
	return o
}

func (o NsRecordResponseArrayOutput) ToNsRecordResponseArrayOutputWithContext(ctx context.Context) NsRecordResponseArrayOutput {
	return o
}

func (o NsRecordResponseArrayOutput) Index(i pulumi.IntInput) NsRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NsRecordResponse {
		return vs[0].([]NsRecordResponse)[vs[1].(int)]
	}).(NsRecordResponseOutput)
}

type PtrRecord struct {
	Ptrdname *string `pulumi:"ptrdname"`
}





type PtrRecordInput interface {
	pulumi.Input

	ToPtrRecordOutput() PtrRecordOutput
	ToPtrRecordOutputWithContext(context.Context) PtrRecordOutput
}

type PtrRecordArgs struct {
	Ptrdname pulumi.StringPtrInput `pulumi:"ptrdname"`
}

func (PtrRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecord)(nil)).Elem()
}

func (i PtrRecordArgs) ToPtrRecordOutput() PtrRecordOutput {
	return i.ToPtrRecordOutputWithContext(context.Background())
}

func (i PtrRecordArgs) ToPtrRecordOutputWithContext(ctx context.Context) PtrRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtrRecordOutput)
}





type PtrRecordArrayInput interface {
	pulumi.Input

	ToPtrRecordArrayOutput() PtrRecordArrayOutput
	ToPtrRecordArrayOutputWithContext(context.Context) PtrRecordArrayOutput
}

type PtrRecordArray []PtrRecordInput

func (PtrRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PtrRecord)(nil)).Elem()
}

func (i PtrRecordArray) ToPtrRecordArrayOutput() PtrRecordArrayOutput {
	return i.ToPtrRecordArrayOutputWithContext(context.Background())
}

func (i PtrRecordArray) ToPtrRecordArrayOutputWithContext(ctx context.Context) PtrRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtrRecordArrayOutput)
}

type PtrRecordOutput struct{ *pulumi.OutputState }

func (PtrRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecord)(nil)).Elem()
}

func (o PtrRecordOutput) ToPtrRecordOutput() PtrRecordOutput {
	return o
}

func (o PtrRecordOutput) ToPtrRecordOutputWithContext(ctx context.Context) PtrRecordOutput {
	return o
}

func (o PtrRecordOutput) Ptrdname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PtrRecord) *string { return v.Ptrdname }).(pulumi.StringPtrOutput)
}

type PtrRecordArrayOutput struct{ *pulumi.OutputState }

func (PtrRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PtrRecord)(nil)).Elem()
}

func (o PtrRecordArrayOutput) ToPtrRecordArrayOutput() PtrRecordArrayOutput {
	return o
}

func (o PtrRecordArrayOutput) ToPtrRecordArrayOutputWithContext(ctx context.Context) PtrRecordArrayOutput {
	return o
}

func (o PtrRecordArrayOutput) Index(i pulumi.IntInput) PtrRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PtrRecord {
		return vs[0].([]PtrRecord)[vs[1].(int)]
	}).(PtrRecordOutput)
}

type PtrRecordResponse struct {
	Ptrdname *string `pulumi:"ptrdname"`
}

type PtrRecordResponseOutput struct{ *pulumi.OutputState }

func (PtrRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecordResponse)(nil)).Elem()
}

func (o PtrRecordResponseOutput) ToPtrRecordResponseOutput() PtrRecordResponseOutput {
	return o
}

func (o PtrRecordResponseOutput) ToPtrRecordResponseOutputWithContext(ctx context.Context) PtrRecordResponseOutput {
	return o
}

func (o PtrRecordResponseOutput) Ptrdname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PtrRecordResponse) *string { return v.Ptrdname }).(pulumi.StringPtrOutput)
}

type PtrRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (PtrRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PtrRecordResponse)(nil)).Elem()
}

func (o PtrRecordResponseArrayOutput) ToPtrRecordResponseArrayOutput() PtrRecordResponseArrayOutput {
	return o
}

func (o PtrRecordResponseArrayOutput) ToPtrRecordResponseArrayOutputWithContext(ctx context.Context) PtrRecordResponseArrayOutput {
	return o
}

func (o PtrRecordResponseArrayOutput) Index(i pulumi.IntInput) PtrRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PtrRecordResponse {
		return vs[0].([]PtrRecordResponse)[vs[1].(int)]
	}).(PtrRecordResponseOutput)
}

type SoaRecord struct {
	Email        *string  `pulumi:"email"`
	ExpireTime   *float64 `pulumi:"expireTime"`
	Host         *string  `pulumi:"host"`
	MinimumTTL   *float64 `pulumi:"minimumTTL"`
	RefreshTime  *float64 `pulumi:"refreshTime"`
	RetryTime    *float64 `pulumi:"retryTime"`
	SerialNumber *float64 `pulumi:"serialNumber"`
}





type SoaRecordInput interface {
	pulumi.Input

	ToSoaRecordOutput() SoaRecordOutput
	ToSoaRecordOutputWithContext(context.Context) SoaRecordOutput
}

type SoaRecordArgs struct {
	Email        pulumi.StringPtrInput  `pulumi:"email"`
	ExpireTime   pulumi.Float64PtrInput `pulumi:"expireTime"`
	Host         pulumi.StringPtrInput  `pulumi:"host"`
	MinimumTTL   pulumi.Float64PtrInput `pulumi:"minimumTTL"`
	RefreshTime  pulumi.Float64PtrInput `pulumi:"refreshTime"`
	RetryTime    pulumi.Float64PtrInput `pulumi:"retryTime"`
	SerialNumber pulumi.Float64PtrInput `pulumi:"serialNumber"`
}

func (SoaRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecord)(nil)).Elem()
}

func (i SoaRecordArgs) ToSoaRecordOutput() SoaRecordOutput {
	return i.ToSoaRecordOutputWithContext(context.Background())
}

func (i SoaRecordArgs) ToSoaRecordOutputWithContext(ctx context.Context) SoaRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordOutput)
}

func (i SoaRecordArgs) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return i.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (i SoaRecordArgs) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordOutput).ToSoaRecordPtrOutputWithContext(ctx)
}









type SoaRecordPtrInput interface {
	pulumi.Input

	ToSoaRecordPtrOutput() SoaRecordPtrOutput
	ToSoaRecordPtrOutputWithContext(context.Context) SoaRecordPtrOutput
}

type soaRecordPtrType SoaRecordArgs

func SoaRecordPtr(v *SoaRecordArgs) SoaRecordPtrInput {
	return (*soaRecordPtrType)(v)
}

func (*soaRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecord)(nil)).Elem()
}

func (i *soaRecordPtrType) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return i.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (i *soaRecordPtrType) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordPtrOutput)
}

type SoaRecordOutput struct{ *pulumi.OutputState }

func (SoaRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecord)(nil)).Elem()
}

func (o SoaRecordOutput) ToSoaRecordOutput() SoaRecordOutput {
	return o
}

func (o SoaRecordOutput) ToSoaRecordOutputWithContext(ctx context.Context) SoaRecordOutput {
	return o
}

func (o SoaRecordOutput) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return o.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (o SoaRecordOutput) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SoaRecord) *SoaRecord {
		return &v
	}).(SoaRecordPtrOutput)
}

func (o SoaRecordOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecord) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o SoaRecordOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.ExpireTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecord) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o SoaRecordOutput) MinimumTTL() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.MinimumTTL }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.RefreshTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.RetryTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.SerialNumber }).(pulumi.Float64PtrOutput)
}

type SoaRecordPtrOutput struct{ *pulumi.OutputState }

func (SoaRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecord)(nil)).Elem()
}

func (o SoaRecordPtrOutput) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return o
}

func (o SoaRecordPtrOutput) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return o
}

func (o SoaRecordPtrOutput) Elem() SoaRecordOutput {
	return o.ApplyT(func(v *SoaRecord) SoaRecord {
		if v != nil {
			return *v
		}
		var ret SoaRecord
		return ret
	}).(SoaRecordOutput)
}

func (o SoaRecordPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecord) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordPtrOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.ExpireTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecord) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordPtrOutput) MinimumTTL() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumTTL
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.Float64PtrOutput)
}

type SoaRecordResponse struct {
	Email        *string  `pulumi:"email"`
	ExpireTime   *float64 `pulumi:"expireTime"`
	Host         *string  `pulumi:"host"`
	MinimumTTL   *float64 `pulumi:"minimumTTL"`
	RefreshTime  *float64 `pulumi:"refreshTime"`
	RetryTime    *float64 `pulumi:"retryTime"`
	SerialNumber *float64 `pulumi:"serialNumber"`
}

type SoaRecordResponseOutput struct{ *pulumi.OutputState }

func (SoaRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecordResponse)(nil)).Elem()
}

func (o SoaRecordResponseOutput) ToSoaRecordResponseOutput() SoaRecordResponseOutput {
	return o
}

func (o SoaRecordResponseOutput) ToSoaRecordResponseOutputWithContext(ctx context.Context) SoaRecordResponseOutput {
	return o
}

func (o SoaRecordResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponseOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.ExpireTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponseOutput) MinimumTTL() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.MinimumTTL }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.RefreshTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.RetryTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.SerialNumber }).(pulumi.Float64PtrOutput)
}

type SoaRecordResponsePtrOutput struct{ *pulumi.OutputState }

func (SoaRecordResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecordResponse)(nil)).Elem()
}

func (o SoaRecordResponsePtrOutput) ToSoaRecordResponsePtrOutput() SoaRecordResponsePtrOutput {
	return o
}

func (o SoaRecordResponsePtrOutput) ToSoaRecordResponsePtrOutputWithContext(ctx context.Context) SoaRecordResponsePtrOutput {
	return o
}

func (o SoaRecordResponsePtrOutput) Elem() SoaRecordResponseOutput {
	return o.ApplyT(func(v *SoaRecordResponse) SoaRecordResponse {
		if v != nil {
			return *v
		}
		var ret SoaRecordResponse
		return ret
	}).(SoaRecordResponseOutput)
}

func (o SoaRecordResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponsePtrOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.ExpireTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponsePtrOutput) MinimumTTL() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumTTL
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.Float64PtrOutput)
}

type SrvRecord struct {
	Port     *int    `pulumi:"port"`
	Priority *int    `pulumi:"priority"`
	Target   *string `pulumi:"target"`
	Weight   *int    `pulumi:"weight"`
}





type SrvRecordInput interface {
	pulumi.Input

	ToSrvRecordOutput() SrvRecordOutput
	ToSrvRecordOutputWithContext(context.Context) SrvRecordOutput
}

type SrvRecordArgs struct {
	Port     pulumi.IntPtrInput    `pulumi:"port"`
	Priority pulumi.IntPtrInput    `pulumi:"priority"`
	Target   pulumi.StringPtrInput `pulumi:"target"`
	Weight   pulumi.IntPtrInput    `pulumi:"weight"`
}

func (SrvRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecord)(nil)).Elem()
}

func (i SrvRecordArgs) ToSrvRecordOutput() SrvRecordOutput {
	return i.ToSrvRecordOutputWithContext(context.Background())
}

func (i SrvRecordArgs) ToSrvRecordOutputWithContext(ctx context.Context) SrvRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordOutput)
}





type SrvRecordArrayInput interface {
	pulumi.Input

	ToSrvRecordArrayOutput() SrvRecordArrayOutput
	ToSrvRecordArrayOutputWithContext(context.Context) SrvRecordArrayOutput
}

type SrvRecordArray []SrvRecordInput

func (SrvRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecord)(nil)).Elem()
}

func (i SrvRecordArray) ToSrvRecordArrayOutput() SrvRecordArrayOutput {
	return i.ToSrvRecordArrayOutputWithContext(context.Background())
}

func (i SrvRecordArray) ToSrvRecordArrayOutputWithContext(ctx context.Context) SrvRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordArrayOutput)
}

type SrvRecordOutput struct{ *pulumi.OutputState }

func (SrvRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecord)(nil)).Elem()
}

func (o SrvRecordOutput) ToSrvRecordOutput() SrvRecordOutput {
	return o
}

func (o SrvRecordOutput) ToSrvRecordOutputWithContext(ctx context.Context) SrvRecordOutput {
	return o
}

func (o SrvRecordOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o SrvRecordOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SrvRecordOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SrvRecord) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o SrvRecordOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type SrvRecordArrayOutput struct{ *pulumi.OutputState }

func (SrvRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecord)(nil)).Elem()
}

func (o SrvRecordArrayOutput) ToSrvRecordArrayOutput() SrvRecordArrayOutput {
	return o
}

func (o SrvRecordArrayOutput) ToSrvRecordArrayOutputWithContext(ctx context.Context) SrvRecordArrayOutput {
	return o
}

func (o SrvRecordArrayOutput) Index(i pulumi.IntInput) SrvRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SrvRecord {
		return vs[0].([]SrvRecord)[vs[1].(int)]
	}).(SrvRecordOutput)
}

type SrvRecordResponse struct {
	Port     *int    `pulumi:"port"`
	Priority *int    `pulumi:"priority"`
	Target   *string `pulumi:"target"`
	Weight   *int    `pulumi:"weight"`
}

type SrvRecordResponseOutput struct{ *pulumi.OutputState }

func (SrvRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecordResponse)(nil)).Elem()
}

func (o SrvRecordResponseOutput) ToSrvRecordResponseOutput() SrvRecordResponseOutput {
	return o
}

func (o SrvRecordResponseOutput) ToSrvRecordResponseOutputWithContext(ctx context.Context) SrvRecordResponseOutput {
	return o
}

func (o SrvRecordResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o SrvRecordResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SrvRecordResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o SrvRecordResponseOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type SrvRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (SrvRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecordResponse)(nil)).Elem()
}

func (o SrvRecordResponseArrayOutput) ToSrvRecordResponseArrayOutput() SrvRecordResponseArrayOutput {
	return o
}

func (o SrvRecordResponseArrayOutput) ToSrvRecordResponseArrayOutputWithContext(ctx context.Context) SrvRecordResponseArrayOutput {
	return o
}

func (o SrvRecordResponseArrayOutput) Index(i pulumi.IntInput) SrvRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SrvRecordResponse {
		return vs[0].([]SrvRecordResponse)[vs[1].(int)]
	}).(SrvRecordResponseOutput)
}

type TxtRecord struct {
	Value []string `pulumi:"value"`
}





type TxtRecordInput interface {
	pulumi.Input

	ToTxtRecordOutput() TxtRecordOutput
	ToTxtRecordOutputWithContext(context.Context) TxtRecordOutput
}

type TxtRecordArgs struct {
	Value pulumi.StringArrayInput `pulumi:"value"`
}

func (TxtRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecord)(nil)).Elem()
}

func (i TxtRecordArgs) ToTxtRecordOutput() TxtRecordOutput {
	return i.ToTxtRecordOutputWithContext(context.Background())
}

func (i TxtRecordArgs) ToTxtRecordOutputWithContext(ctx context.Context) TxtRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordOutput)
}





type TxtRecordArrayInput interface {
	pulumi.Input

	ToTxtRecordArrayOutput() TxtRecordArrayOutput
	ToTxtRecordArrayOutputWithContext(context.Context) TxtRecordArrayOutput
}

type TxtRecordArray []TxtRecordInput

func (TxtRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecord)(nil)).Elem()
}

func (i TxtRecordArray) ToTxtRecordArrayOutput() TxtRecordArrayOutput {
	return i.ToTxtRecordArrayOutputWithContext(context.Background())
}

func (i TxtRecordArray) ToTxtRecordArrayOutputWithContext(ctx context.Context) TxtRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordArrayOutput)
}

type TxtRecordOutput struct{ *pulumi.OutputState }

func (TxtRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecord)(nil)).Elem()
}

func (o TxtRecordOutput) ToTxtRecordOutput() TxtRecordOutput {
	return o
}

func (o TxtRecordOutput) ToTxtRecordOutputWithContext(ctx context.Context) TxtRecordOutput {
	return o
}

func (o TxtRecordOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TxtRecord) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type TxtRecordArrayOutput struct{ *pulumi.OutputState }

func (TxtRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecord)(nil)).Elem()
}

func (o TxtRecordArrayOutput) ToTxtRecordArrayOutput() TxtRecordArrayOutput {
	return o
}

func (o TxtRecordArrayOutput) ToTxtRecordArrayOutputWithContext(ctx context.Context) TxtRecordArrayOutput {
	return o
}

func (o TxtRecordArrayOutput) Index(i pulumi.IntInput) TxtRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TxtRecord {
		return vs[0].([]TxtRecord)[vs[1].(int)]
	}).(TxtRecordOutput)
}

type TxtRecordResponse struct {
	Value []string `pulumi:"value"`
}

type TxtRecordResponseOutput struct{ *pulumi.OutputState }

func (TxtRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecordResponse)(nil)).Elem()
}

func (o TxtRecordResponseOutput) ToTxtRecordResponseOutput() TxtRecordResponseOutput {
	return o
}

func (o TxtRecordResponseOutput) ToTxtRecordResponseOutputWithContext(ctx context.Context) TxtRecordResponseOutput {
	return o
}

func (o TxtRecordResponseOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TxtRecordResponse) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type TxtRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (TxtRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecordResponse)(nil)).Elem()
}

func (o TxtRecordResponseArrayOutput) ToTxtRecordResponseArrayOutput() TxtRecordResponseArrayOutput {
	return o
}

func (o TxtRecordResponseArrayOutput) ToTxtRecordResponseArrayOutputWithContext(ctx context.Context) TxtRecordResponseArrayOutput {
	return o
}

func (o TxtRecordResponseArrayOutput) Index(i pulumi.IntInput) TxtRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TxtRecordResponse {
		return vs[0].([]TxtRecordResponse)[vs[1].(int)]
	}).(TxtRecordResponseOutput)
}

type ZoneProperties struct {
	MaxNumberOfRecordSets *float64 `pulumi:"maxNumberOfRecordSets"`
	NumberOfRecordSets    *float64 `pulumi:"numberOfRecordSets"`
}





type ZonePropertiesInput interface {
	pulumi.Input

	ToZonePropertiesOutput() ZonePropertiesOutput
	ToZonePropertiesOutputWithContext(context.Context) ZonePropertiesOutput
}

type ZonePropertiesArgs struct {
	MaxNumberOfRecordSets pulumi.Float64PtrInput `pulumi:"maxNumberOfRecordSets"`
	NumberOfRecordSets    pulumi.Float64PtrInput `pulumi:"numberOfRecordSets"`
}

func (ZonePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneProperties)(nil)).Elem()
}

func (i ZonePropertiesArgs) ToZonePropertiesOutput() ZonePropertiesOutput {
	return i.ToZonePropertiesOutputWithContext(context.Background())
}

func (i ZonePropertiesArgs) ToZonePropertiesOutputWithContext(ctx context.Context) ZonePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZonePropertiesOutput)
}

func (i ZonePropertiesArgs) ToZonePropertiesPtrOutput() ZonePropertiesPtrOutput {
	return i.ToZonePropertiesPtrOutputWithContext(context.Background())
}

func (i ZonePropertiesArgs) ToZonePropertiesPtrOutputWithContext(ctx context.Context) ZonePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZonePropertiesOutput).ToZonePropertiesPtrOutputWithContext(ctx)
}









type ZonePropertiesPtrInput interface {
	pulumi.Input

	ToZonePropertiesPtrOutput() ZonePropertiesPtrOutput
	ToZonePropertiesPtrOutputWithContext(context.Context) ZonePropertiesPtrOutput
}

type zonePropertiesPtrType ZonePropertiesArgs

func ZonePropertiesPtr(v *ZonePropertiesArgs) ZonePropertiesPtrInput {
	return (*zonePropertiesPtrType)(v)
}

func (*zonePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneProperties)(nil)).Elem()
}

func (i *zonePropertiesPtrType) ToZonePropertiesPtrOutput() ZonePropertiesPtrOutput {
	return i.ToZonePropertiesPtrOutputWithContext(context.Background())
}

func (i *zonePropertiesPtrType) ToZonePropertiesPtrOutputWithContext(ctx context.Context) ZonePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZonePropertiesPtrOutput)
}

type ZonePropertiesOutput struct{ *pulumi.OutputState }

func (ZonePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneProperties)(nil)).Elem()
}

func (o ZonePropertiesOutput) ToZonePropertiesOutput() ZonePropertiesOutput {
	return o
}

func (o ZonePropertiesOutput) ToZonePropertiesOutputWithContext(ctx context.Context) ZonePropertiesOutput {
	return o
}

func (o ZonePropertiesOutput) ToZonePropertiesPtrOutput() ZonePropertiesPtrOutput {
	return o.ToZonePropertiesPtrOutputWithContext(context.Background())
}

func (o ZonePropertiesOutput) ToZonePropertiesPtrOutputWithContext(ctx context.Context) ZonePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ZoneProperties) *ZoneProperties {
		return &v
	}).(ZonePropertiesPtrOutput)
}

func (o ZonePropertiesOutput) MaxNumberOfRecordSets() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ZoneProperties) *float64 { return v.MaxNumberOfRecordSets }).(pulumi.Float64PtrOutput)
}

func (o ZonePropertiesOutput) NumberOfRecordSets() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ZoneProperties) *float64 { return v.NumberOfRecordSets }).(pulumi.Float64PtrOutput)
}

type ZonePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ZonePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneProperties)(nil)).Elem()
}

func (o ZonePropertiesPtrOutput) ToZonePropertiesPtrOutput() ZonePropertiesPtrOutput {
	return o
}

func (o ZonePropertiesPtrOutput) ToZonePropertiesPtrOutputWithContext(ctx context.Context) ZonePropertiesPtrOutput {
	return o
}

func (o ZonePropertiesPtrOutput) Elem() ZonePropertiesOutput {
	return o.ApplyT(func(v *ZoneProperties) ZoneProperties {
		if v != nil {
			return *v
		}
		var ret ZoneProperties
		return ret
	}).(ZonePropertiesOutput)
}

func (o ZonePropertiesPtrOutput) MaxNumberOfRecordSets() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ZoneProperties) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxNumberOfRecordSets
	}).(pulumi.Float64PtrOutput)
}

func (o ZonePropertiesPtrOutput) NumberOfRecordSets() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ZoneProperties) *float64 {
		if v == nil {
			return nil
		}
		return v.NumberOfRecordSets
	}).(pulumi.Float64PtrOutput)
}

type ZonePropertiesResponse struct {
	MaxNumberOfRecordSets          *float64 `pulumi:"maxNumberOfRecordSets"`
	MaxNumberOfRecordsPerRecordSet float64  `pulumi:"maxNumberOfRecordsPerRecordSet"`
	NumberOfRecordSets             *float64 `pulumi:"numberOfRecordSets"`
}

type ZonePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ZonePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZonePropertiesResponse)(nil)).Elem()
}

func (o ZonePropertiesResponseOutput) ToZonePropertiesResponseOutput() ZonePropertiesResponseOutput {
	return o
}

func (o ZonePropertiesResponseOutput) ToZonePropertiesResponseOutputWithContext(ctx context.Context) ZonePropertiesResponseOutput {
	return o
}

func (o ZonePropertiesResponseOutput) MaxNumberOfRecordSets() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ZonePropertiesResponse) *float64 { return v.MaxNumberOfRecordSets }).(pulumi.Float64PtrOutput)
}

func (o ZonePropertiesResponseOutput) MaxNumberOfRecordsPerRecordSet() pulumi.Float64Output {
	return o.ApplyT(func(v ZonePropertiesResponse) float64 { return v.MaxNumberOfRecordsPerRecordSet }).(pulumi.Float64Output)
}

func (o ZonePropertiesResponseOutput) NumberOfRecordSets() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ZonePropertiesResponse) *float64 { return v.NumberOfRecordSets }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ARecordOutput{})
	pulumi.RegisterOutputType(ARecordArrayOutput{})
	pulumi.RegisterOutputType(ARecordResponseOutput{})
	pulumi.RegisterOutputType(ARecordResponseArrayOutput{})
	pulumi.RegisterOutputType(AaaaRecordOutput{})
	pulumi.RegisterOutputType(AaaaRecordArrayOutput{})
	pulumi.RegisterOutputType(AaaaRecordResponseOutput{})
	pulumi.RegisterOutputType(AaaaRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(CnameRecordOutput{})
	pulumi.RegisterOutputType(CnameRecordPtrOutput{})
	pulumi.RegisterOutputType(CnameRecordResponseOutput{})
	pulumi.RegisterOutputType(CnameRecordResponsePtrOutput{})
	pulumi.RegisterOutputType(MxRecordOutput{})
	pulumi.RegisterOutputType(MxRecordArrayOutput{})
	pulumi.RegisterOutputType(MxRecordResponseOutput{})
	pulumi.RegisterOutputType(MxRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(NsRecordOutput{})
	pulumi.RegisterOutputType(NsRecordArrayOutput{})
	pulumi.RegisterOutputType(NsRecordResponseOutput{})
	pulumi.RegisterOutputType(NsRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(PtrRecordOutput{})
	pulumi.RegisterOutputType(PtrRecordArrayOutput{})
	pulumi.RegisterOutputType(PtrRecordResponseOutput{})
	pulumi.RegisterOutputType(PtrRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(SoaRecordOutput{})
	pulumi.RegisterOutputType(SoaRecordPtrOutput{})
	pulumi.RegisterOutputType(SoaRecordResponseOutput{})
	pulumi.RegisterOutputType(SoaRecordResponsePtrOutput{})
	pulumi.RegisterOutputType(SrvRecordOutput{})
	pulumi.RegisterOutputType(SrvRecordArrayOutput{})
	pulumi.RegisterOutputType(SrvRecordResponseOutput{})
	pulumi.RegisterOutputType(SrvRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(TxtRecordOutput{})
	pulumi.RegisterOutputType(TxtRecordArrayOutput{})
	pulumi.RegisterOutputType(TxtRecordResponseOutput{})
	pulumi.RegisterOutputType(TxtRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(ZonePropertiesOutput{})
	pulumi.RegisterOutputType(ZonePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ZonePropertiesResponseOutput{})
}
