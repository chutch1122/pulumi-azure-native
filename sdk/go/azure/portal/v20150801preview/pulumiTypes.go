


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DashboardLens struct {
	Metadata map[string]interface{}    `pulumi:"metadata"`
	Order    int                       `pulumi:"order"`
	Parts    map[string]DashboardParts `pulumi:"parts"`
}





type DashboardLensInput interface {
	pulumi.Input

	ToDashboardLensOutput() DashboardLensOutput
	ToDashboardLensOutputWithContext(context.Context) DashboardLensOutput
}

type DashboardLensArgs struct {
	Metadata pulumi.MapInput        `pulumi:"metadata"`
	Order    pulumi.IntInput        `pulumi:"order"`
	Parts    DashboardPartsMapInput `pulumi:"parts"`
}

func (DashboardLensArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardLens)(nil)).Elem()
}

func (i DashboardLensArgs) ToDashboardLensOutput() DashboardLensOutput {
	return i.ToDashboardLensOutputWithContext(context.Background())
}

func (i DashboardLensArgs) ToDashboardLensOutputWithContext(ctx context.Context) DashboardLensOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardLensOutput)
}





type DashboardLensMapInput interface {
	pulumi.Input

	ToDashboardLensMapOutput() DashboardLensMapOutput
	ToDashboardLensMapOutputWithContext(context.Context) DashboardLensMapOutput
}

type DashboardLensMap map[string]DashboardLensInput

func (DashboardLensMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardLens)(nil)).Elem()
}

func (i DashboardLensMap) ToDashboardLensMapOutput() DashboardLensMapOutput {
	return i.ToDashboardLensMapOutputWithContext(context.Background())
}

func (i DashboardLensMap) ToDashboardLensMapOutputWithContext(ctx context.Context) DashboardLensMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardLensMapOutput)
}

type DashboardLensOutput struct{ *pulumi.OutputState }

func (DashboardLensOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardLens)(nil)).Elem()
}

func (o DashboardLensOutput) ToDashboardLensOutput() DashboardLensOutput {
	return o
}

func (o DashboardLensOutput) ToDashboardLensOutputWithContext(ctx context.Context) DashboardLensOutput {
	return o
}

func (o DashboardLensOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardLens) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o DashboardLensOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardLens) int { return v.Order }).(pulumi.IntOutput)
}

func (o DashboardLensOutput) Parts() DashboardPartsMapOutput {
	return o.ApplyT(func(v DashboardLens) map[string]DashboardParts { return v.Parts }).(DashboardPartsMapOutput)
}

type DashboardLensMapOutput struct{ *pulumi.OutputState }

func (DashboardLensMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardLens)(nil)).Elem()
}

func (o DashboardLensMapOutput) ToDashboardLensMapOutput() DashboardLensMapOutput {
	return o
}

func (o DashboardLensMapOutput) ToDashboardLensMapOutputWithContext(ctx context.Context) DashboardLensMapOutput {
	return o
}

func (o DashboardLensMapOutput) MapIndex(k pulumi.StringInput) DashboardLensOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DashboardLens {
		return vs[0].(map[string]DashboardLens)[vs[1].(string)]
	}).(DashboardLensOutput)
}

type DashboardLensResponse struct {
	Metadata map[string]interface{}            `pulumi:"metadata"`
	Order    int                               `pulumi:"order"`
	Parts    map[string]DashboardPartsResponse `pulumi:"parts"`
}

type DashboardLensResponseOutput struct{ *pulumi.OutputState }

func (DashboardLensResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardLensResponse)(nil)).Elem()
}

func (o DashboardLensResponseOutput) ToDashboardLensResponseOutput() DashboardLensResponseOutput {
	return o
}

func (o DashboardLensResponseOutput) ToDashboardLensResponseOutputWithContext(ctx context.Context) DashboardLensResponseOutput {
	return o
}

func (o DashboardLensResponseOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardLensResponse) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o DashboardLensResponseOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardLensResponse) int { return v.Order }).(pulumi.IntOutput)
}

func (o DashboardLensResponseOutput) Parts() DashboardPartsResponseMapOutput {
	return o.ApplyT(func(v DashboardLensResponse) map[string]DashboardPartsResponse { return v.Parts }).(DashboardPartsResponseMapOutput)
}

type DashboardLensResponseMapOutput struct{ *pulumi.OutputState }

func (DashboardLensResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardLensResponse)(nil)).Elem()
}

func (o DashboardLensResponseMapOutput) ToDashboardLensResponseMapOutput() DashboardLensResponseMapOutput {
	return o
}

func (o DashboardLensResponseMapOutput) ToDashboardLensResponseMapOutputWithContext(ctx context.Context) DashboardLensResponseMapOutput {
	return o
}

func (o DashboardLensResponseMapOutput) MapIndex(k pulumi.StringInput) DashboardLensResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DashboardLensResponse {
		return vs[0].(map[string]DashboardLensResponse)[vs[1].(string)]
	}).(DashboardLensResponseOutput)
}

type DashboardParts struct {
	Metadata interface{}            `pulumi:"metadata"`
	Position DashboardPartsPosition `pulumi:"position"`
}





type DashboardPartsInput interface {
	pulumi.Input

	ToDashboardPartsOutput() DashboardPartsOutput
	ToDashboardPartsOutputWithContext(context.Context) DashboardPartsOutput
}

type DashboardPartsArgs struct {
	Metadata pulumi.Input                `pulumi:"metadata"`
	Position DashboardPartsPositionInput `pulumi:"position"`
}

func (DashboardPartsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardParts)(nil)).Elem()
}

func (i DashboardPartsArgs) ToDashboardPartsOutput() DashboardPartsOutput {
	return i.ToDashboardPartsOutputWithContext(context.Background())
}

func (i DashboardPartsArgs) ToDashboardPartsOutputWithContext(ctx context.Context) DashboardPartsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsOutput)
}





type DashboardPartsMapInput interface {
	pulumi.Input

	ToDashboardPartsMapOutput() DashboardPartsMapOutput
	ToDashboardPartsMapOutputWithContext(context.Context) DashboardPartsMapOutput
}

type DashboardPartsMap map[string]DashboardPartsInput

func (DashboardPartsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardParts)(nil)).Elem()
}

func (i DashboardPartsMap) ToDashboardPartsMapOutput() DashboardPartsMapOutput {
	return i.ToDashboardPartsMapOutputWithContext(context.Background())
}

func (i DashboardPartsMap) ToDashboardPartsMapOutputWithContext(ctx context.Context) DashboardPartsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsMapOutput)
}

type DashboardPartsOutput struct{ *pulumi.OutputState }

func (DashboardPartsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardParts)(nil)).Elem()
}

func (o DashboardPartsOutput) ToDashboardPartsOutput() DashboardPartsOutput {
	return o
}

func (o DashboardPartsOutput) ToDashboardPartsOutputWithContext(ctx context.Context) DashboardPartsOutput {
	return o
}

func (o DashboardPartsOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v DashboardParts) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o DashboardPartsOutput) Position() DashboardPartsPositionOutput {
	return o.ApplyT(func(v DashboardParts) DashboardPartsPosition { return v.Position }).(DashboardPartsPositionOutput)
}

type DashboardPartsMapOutput struct{ *pulumi.OutputState }

func (DashboardPartsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardParts)(nil)).Elem()
}

func (o DashboardPartsMapOutput) ToDashboardPartsMapOutput() DashboardPartsMapOutput {
	return o
}

func (o DashboardPartsMapOutput) ToDashboardPartsMapOutputWithContext(ctx context.Context) DashboardPartsMapOutput {
	return o
}

func (o DashboardPartsMapOutput) MapIndex(k pulumi.StringInput) DashboardPartsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DashboardParts {
		return vs[0].(map[string]DashboardParts)[vs[1].(string)]
	}).(DashboardPartsOutput)
}

type DashboardPartsPosition struct {
	ColSpan  int                    `pulumi:"colSpan"`
	Metadata map[string]interface{} `pulumi:"metadata"`
	RowSpan  int                    `pulumi:"rowSpan"`
	X        int                    `pulumi:"x"`
	Y        int                    `pulumi:"y"`
}





type DashboardPartsPositionInput interface {
	pulumi.Input

	ToDashboardPartsPositionOutput() DashboardPartsPositionOutput
	ToDashboardPartsPositionOutputWithContext(context.Context) DashboardPartsPositionOutput
}

type DashboardPartsPositionArgs struct {
	ColSpan  pulumi.IntInput `pulumi:"colSpan"`
	Metadata pulumi.MapInput `pulumi:"metadata"`
	RowSpan  pulumi.IntInput `pulumi:"rowSpan"`
	X        pulumi.IntInput `pulumi:"x"`
	Y        pulumi.IntInput `pulumi:"y"`
}

func (DashboardPartsPositionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsPosition)(nil)).Elem()
}

func (i DashboardPartsPositionArgs) ToDashboardPartsPositionOutput() DashboardPartsPositionOutput {
	return i.ToDashboardPartsPositionOutputWithContext(context.Background())
}

func (i DashboardPartsPositionArgs) ToDashboardPartsPositionOutputWithContext(ctx context.Context) DashboardPartsPositionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsPositionOutput)
}

type DashboardPartsPositionOutput struct{ *pulumi.OutputState }

func (DashboardPartsPositionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsPosition)(nil)).Elem()
}

func (o DashboardPartsPositionOutput) ToDashboardPartsPositionOutput() DashboardPartsPositionOutput {
	return o
}

func (o DashboardPartsPositionOutput) ToDashboardPartsPositionOutputWithContext(ctx context.Context) DashboardPartsPositionOutput {
	return o
}

func (o DashboardPartsPositionOutput) ColSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.ColSpan }).(pulumi.IntOutput)
}

func (o DashboardPartsPositionOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardPartsPosition) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o DashboardPartsPositionOutput) RowSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.RowSpan }).(pulumi.IntOutput)
}

func (o DashboardPartsPositionOutput) X() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.X }).(pulumi.IntOutput)
}

func (o DashboardPartsPositionOutput) Y() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.Y }).(pulumi.IntOutput)
}

type DashboardPartsResponse struct {
	Metadata interface{}                    `pulumi:"metadata"`
	Position DashboardPartsResponsePosition `pulumi:"position"`
}

type DashboardPartsResponseOutput struct{ *pulumi.OutputState }

func (DashboardPartsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsResponse)(nil)).Elem()
}

func (o DashboardPartsResponseOutput) ToDashboardPartsResponseOutput() DashboardPartsResponseOutput {
	return o
}

func (o DashboardPartsResponseOutput) ToDashboardPartsResponseOutputWithContext(ctx context.Context) DashboardPartsResponseOutput {
	return o
}

func (o DashboardPartsResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v DashboardPartsResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o DashboardPartsResponseOutput) Position() DashboardPartsResponsePositionOutput {
	return o.ApplyT(func(v DashboardPartsResponse) DashboardPartsResponsePosition { return v.Position }).(DashboardPartsResponsePositionOutput)
}

type DashboardPartsResponseMapOutput struct{ *pulumi.OutputState }

func (DashboardPartsResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardPartsResponse)(nil)).Elem()
}

func (o DashboardPartsResponseMapOutput) ToDashboardPartsResponseMapOutput() DashboardPartsResponseMapOutput {
	return o
}

func (o DashboardPartsResponseMapOutput) ToDashboardPartsResponseMapOutputWithContext(ctx context.Context) DashboardPartsResponseMapOutput {
	return o
}

func (o DashboardPartsResponseMapOutput) MapIndex(k pulumi.StringInput) DashboardPartsResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DashboardPartsResponse {
		return vs[0].(map[string]DashboardPartsResponse)[vs[1].(string)]
	}).(DashboardPartsResponseOutput)
}

type DashboardPartsResponsePosition struct {
	ColSpan  int                    `pulumi:"colSpan"`
	Metadata map[string]interface{} `pulumi:"metadata"`
	RowSpan  int                    `pulumi:"rowSpan"`
	X        int                    `pulumi:"x"`
	Y        int                    `pulumi:"y"`
}

type DashboardPartsResponsePositionOutput struct{ *pulumi.OutputState }

func (DashboardPartsResponsePositionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsResponsePosition)(nil)).Elem()
}

func (o DashboardPartsResponsePositionOutput) ToDashboardPartsResponsePositionOutput() DashboardPartsResponsePositionOutput {
	return o
}

func (o DashboardPartsResponsePositionOutput) ToDashboardPartsResponsePositionOutputWithContext(ctx context.Context) DashboardPartsResponsePositionOutput {
	return o
}

func (o DashboardPartsResponsePositionOutput) ColSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.ColSpan }).(pulumi.IntOutput)
}

func (o DashboardPartsResponsePositionOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o DashboardPartsResponsePositionOutput) RowSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.RowSpan }).(pulumi.IntOutput)
}

func (o DashboardPartsResponsePositionOutput) X() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.X }).(pulumi.IntOutput)
}

func (o DashboardPartsResponsePositionOutput) Y() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.Y }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(DashboardLensOutput{})
	pulumi.RegisterOutputType(DashboardLensMapOutput{})
	pulumi.RegisterOutputType(DashboardLensResponseOutput{})
	pulumi.RegisterOutputType(DashboardLensResponseMapOutput{})
	pulumi.RegisterOutputType(DashboardPartsOutput{})
	pulumi.RegisterOutputType(DashboardPartsMapOutput{})
	pulumi.RegisterOutputType(DashboardPartsPositionOutput{})
	pulumi.RegisterOutputType(DashboardPartsResponseOutput{})
	pulumi.RegisterOutputType(DashboardPartsResponseMapOutput{})
	pulumi.RegisterOutputType(DashboardPartsResponsePositionOutput{})
}
