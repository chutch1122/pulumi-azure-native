


package eventhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CaptureDescription struct {
	Destination       *Destination                `pulumi:"destination"`
	Enabled           *bool                       `pulumi:"enabled"`
	Encoding          *EncodingCaptureDescription `pulumi:"encoding"`
	IntervalInSeconds *int                        `pulumi:"intervalInSeconds"`
	SizeLimitInBytes  *int                        `pulumi:"sizeLimitInBytes"`
	SkipEmptyArchives *bool                       `pulumi:"skipEmptyArchives"`
}





type CaptureDescriptionInput interface {
	pulumi.Input

	ToCaptureDescriptionOutput() CaptureDescriptionOutput
	ToCaptureDescriptionOutputWithContext(context.Context) CaptureDescriptionOutput
}

type CaptureDescriptionArgs struct {
	Destination       DestinationPtrInput                `pulumi:"destination"`
	Enabled           pulumi.BoolPtrInput                `pulumi:"enabled"`
	Encoding          EncodingCaptureDescriptionPtrInput `pulumi:"encoding"`
	IntervalInSeconds pulumi.IntPtrInput                 `pulumi:"intervalInSeconds"`
	SizeLimitInBytes  pulumi.IntPtrInput                 `pulumi:"sizeLimitInBytes"`
	SkipEmptyArchives pulumi.BoolPtrInput                `pulumi:"skipEmptyArchives"`
}

func (CaptureDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CaptureDescription)(nil)).Elem()
}

func (i CaptureDescriptionArgs) ToCaptureDescriptionOutput() CaptureDescriptionOutput {
	return i.ToCaptureDescriptionOutputWithContext(context.Background())
}

func (i CaptureDescriptionArgs) ToCaptureDescriptionOutputWithContext(ctx context.Context) CaptureDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaptureDescriptionOutput)
}

func (i CaptureDescriptionArgs) ToCaptureDescriptionPtrOutput() CaptureDescriptionPtrOutput {
	return i.ToCaptureDescriptionPtrOutputWithContext(context.Background())
}

func (i CaptureDescriptionArgs) ToCaptureDescriptionPtrOutputWithContext(ctx context.Context) CaptureDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaptureDescriptionOutput).ToCaptureDescriptionPtrOutputWithContext(ctx)
}









type CaptureDescriptionPtrInput interface {
	pulumi.Input

	ToCaptureDescriptionPtrOutput() CaptureDescriptionPtrOutput
	ToCaptureDescriptionPtrOutputWithContext(context.Context) CaptureDescriptionPtrOutput
}

type captureDescriptionPtrType CaptureDescriptionArgs

func CaptureDescriptionPtr(v *CaptureDescriptionArgs) CaptureDescriptionPtrInput {
	return (*captureDescriptionPtrType)(v)
}

func (*captureDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CaptureDescription)(nil)).Elem()
}

func (i *captureDescriptionPtrType) ToCaptureDescriptionPtrOutput() CaptureDescriptionPtrOutput {
	return i.ToCaptureDescriptionPtrOutputWithContext(context.Background())
}

func (i *captureDescriptionPtrType) ToCaptureDescriptionPtrOutputWithContext(ctx context.Context) CaptureDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaptureDescriptionPtrOutput)
}

type CaptureDescriptionOutput struct{ *pulumi.OutputState }

func (CaptureDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CaptureDescription)(nil)).Elem()
}

func (o CaptureDescriptionOutput) ToCaptureDescriptionOutput() CaptureDescriptionOutput {
	return o
}

func (o CaptureDescriptionOutput) ToCaptureDescriptionOutputWithContext(ctx context.Context) CaptureDescriptionOutput {
	return o
}

func (o CaptureDescriptionOutput) ToCaptureDescriptionPtrOutput() CaptureDescriptionPtrOutput {
	return o.ToCaptureDescriptionPtrOutputWithContext(context.Background())
}

func (o CaptureDescriptionOutput) ToCaptureDescriptionPtrOutputWithContext(ctx context.Context) CaptureDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CaptureDescription) *CaptureDescription {
		return &v
	}).(CaptureDescriptionPtrOutput)
}

func (o CaptureDescriptionOutput) Destination() DestinationPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *Destination { return v.Destination }).(DestinationPtrOutput)
}

func (o CaptureDescriptionOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o CaptureDescriptionOutput) Encoding() EncodingCaptureDescriptionPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *EncodingCaptureDescription { return v.Encoding }).(EncodingCaptureDescriptionPtrOutput)
}

func (o CaptureDescriptionOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionOutput) SizeLimitInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *int { return v.SizeLimitInBytes }).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionOutput) SkipEmptyArchives() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CaptureDescription) *bool { return v.SkipEmptyArchives }).(pulumi.BoolPtrOutput)
}

type CaptureDescriptionPtrOutput struct{ *pulumi.OutputState }

func (CaptureDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaptureDescription)(nil)).Elem()
}

func (o CaptureDescriptionPtrOutput) ToCaptureDescriptionPtrOutput() CaptureDescriptionPtrOutput {
	return o
}

func (o CaptureDescriptionPtrOutput) ToCaptureDescriptionPtrOutputWithContext(ctx context.Context) CaptureDescriptionPtrOutput {
	return o
}

func (o CaptureDescriptionPtrOutput) Elem() CaptureDescriptionOutput {
	return o.ApplyT(func(v *CaptureDescription) CaptureDescription {
		if v != nil {
			return *v
		}
		var ret CaptureDescription
		return ret
	}).(CaptureDescriptionOutput)
}

func (o CaptureDescriptionPtrOutput) Destination() DestinationPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *Destination {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(DestinationPtrOutput)
}

func (o CaptureDescriptionPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o CaptureDescriptionPtrOutput) Encoding() EncodingCaptureDescriptionPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *EncodingCaptureDescription {
		if v == nil {
			return nil
		}
		return v.Encoding
	}).(EncodingCaptureDescriptionPtrOutput)
}

func (o CaptureDescriptionPtrOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *int {
		if v == nil {
			return nil
		}
		return v.IntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionPtrOutput) SizeLimitInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *int {
		if v == nil {
			return nil
		}
		return v.SizeLimitInBytes
	}).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionPtrOutput) SkipEmptyArchives() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CaptureDescription) *bool {
		if v == nil {
			return nil
		}
		return v.SkipEmptyArchives
	}).(pulumi.BoolPtrOutput)
}

type CaptureDescriptionResponse struct {
	Destination       *DestinationResponse `pulumi:"destination"`
	Enabled           *bool                `pulumi:"enabled"`
	Encoding          *string              `pulumi:"encoding"`
	IntervalInSeconds *int                 `pulumi:"intervalInSeconds"`
	SizeLimitInBytes  *int                 `pulumi:"sizeLimitInBytes"`
	SkipEmptyArchives *bool                `pulumi:"skipEmptyArchives"`
}

type CaptureDescriptionResponseOutput struct{ *pulumi.OutputState }

func (CaptureDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CaptureDescriptionResponse)(nil)).Elem()
}

func (o CaptureDescriptionResponseOutput) ToCaptureDescriptionResponseOutput() CaptureDescriptionResponseOutput {
	return o
}

func (o CaptureDescriptionResponseOutput) ToCaptureDescriptionResponseOutputWithContext(ctx context.Context) CaptureDescriptionResponseOutput {
	return o
}

func (o CaptureDescriptionResponseOutput) Destination() DestinationResponsePtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *DestinationResponse { return v.Destination }).(DestinationResponsePtrOutput)
}

func (o CaptureDescriptionResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o CaptureDescriptionResponseOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *string { return v.Encoding }).(pulumi.StringPtrOutput)
}

func (o CaptureDescriptionResponseOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionResponseOutput) SizeLimitInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *int { return v.SizeLimitInBytes }).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionResponseOutput) SkipEmptyArchives() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CaptureDescriptionResponse) *bool { return v.SkipEmptyArchives }).(pulumi.BoolPtrOutput)
}

type CaptureDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (CaptureDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaptureDescriptionResponse)(nil)).Elem()
}

func (o CaptureDescriptionResponsePtrOutput) ToCaptureDescriptionResponsePtrOutput() CaptureDescriptionResponsePtrOutput {
	return o
}

func (o CaptureDescriptionResponsePtrOutput) ToCaptureDescriptionResponsePtrOutputWithContext(ctx context.Context) CaptureDescriptionResponsePtrOutput {
	return o
}

func (o CaptureDescriptionResponsePtrOutput) Elem() CaptureDescriptionResponseOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) CaptureDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret CaptureDescriptionResponse
		return ret
	}).(CaptureDescriptionResponseOutput)
}

func (o CaptureDescriptionResponsePtrOutput) Destination() DestinationResponsePtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *DestinationResponse {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(DestinationResponsePtrOutput)
}

func (o CaptureDescriptionResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o CaptureDescriptionResponsePtrOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Encoding
	}).(pulumi.StringPtrOutput)
}

func (o CaptureDescriptionResponsePtrOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return v.IntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionResponsePtrOutput) SizeLimitInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return v.SizeLimitInBytes
	}).(pulumi.IntPtrOutput)
}

func (o CaptureDescriptionResponsePtrOutput) SkipEmptyArchives() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CaptureDescriptionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SkipEmptyArchives
	}).(pulumi.BoolPtrOutput)
}

type ClusterSku struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}





type ClusterSkuInput interface {
	pulumi.Input

	ToClusterSkuOutput() ClusterSkuOutput
	ToClusterSkuOutputWithContext(context.Context) ClusterSkuOutput
}

type ClusterSkuArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (ClusterSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (i ClusterSkuArgs) ToClusterSkuOutput() ClusterSkuOutput {
	return i.ToClusterSkuOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput)
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput).ToClusterSkuPtrOutputWithContext(ctx)
}









type ClusterSkuPtrInput interface {
	pulumi.Input

	ToClusterSkuPtrOutput() ClusterSkuPtrOutput
	ToClusterSkuPtrOutputWithContext(context.Context) ClusterSkuPtrOutput
}

type clusterSkuPtrType ClusterSkuArgs

func ClusterSkuPtr(v *ClusterSkuArgs) ClusterSkuPtrInput {
	return (*clusterSkuPtrType)(v)
}

func (*clusterSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuPtrOutput)
}

type ClusterSkuOutput struct{ *pulumi.OutputState }

func (ClusterSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (o ClusterSkuOutput) ToClusterSkuOutput() ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterSku) *ClusterSku {
		return &v
	}).(ClusterSkuPtrOutput)
}

func (o ClusterSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ClusterSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterSku) string { return v.Name }).(pulumi.StringOutput)
}

type ClusterSkuPtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) Elem() ClusterSkuOutput {
	return o.ApplyT(func(v *ClusterSku) ClusterSku {
		if v != nil {
			return *v
		}
		var ret ClusterSku
		return ret
	}).(ClusterSkuOutput)
}

func (o ClusterSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ClusterSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ClusterSkuResponse struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}

type ClusterSkuResponseOutput struct{ *pulumi.OutputState }

func (ClusterSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSkuResponse)(nil)).Elem()
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponseOutput() ClusterSkuResponseOutput {
	return o
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponseOutputWithContext(ctx context.Context) ClusterSkuResponseOutput {
	return o
}

func (o ClusterSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ClusterSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ClusterSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSkuResponse)(nil)).Elem()
}

func (o ClusterSkuResponsePtrOutput) ToClusterSkuResponsePtrOutput() ClusterSkuResponsePtrOutput {
	return o
}

func (o ClusterSkuResponsePtrOutput) ToClusterSkuResponsePtrOutputWithContext(ctx context.Context) ClusterSkuResponsePtrOutput {
	return o
}

func (o ClusterSkuResponsePtrOutput) Elem() ClusterSkuResponseOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) ClusterSkuResponse {
		if v != nil {
			return *v
		}
		var ret ClusterSkuResponse
		return ret
	}).(ClusterSkuResponseOutput)
}

func (o ClusterSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ClusterSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ConnectionState struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}





type ConnectionStateInput interface {
	pulumi.Input

	ToConnectionStateOutput() ConnectionStateOutput
	ToConnectionStateOutputWithContext(context.Context) ConnectionStateOutput
}

type ConnectionStateArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
}

func (ConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (i ConnectionStateArgs) ToConnectionStateOutput() ConnectionStateOutput {
	return i.ToConnectionStateOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput)
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput).ToConnectionStatePtrOutputWithContext(ctx)
}









type ConnectionStatePtrInput interface {
	pulumi.Input

	ToConnectionStatePtrOutput() ConnectionStatePtrOutput
	ToConnectionStatePtrOutputWithContext(context.Context) ConnectionStatePtrOutput
}

type connectionStatePtrType ConnectionStateArgs

func ConnectionStatePtr(v *ConnectionStateArgs) ConnectionStatePtrInput {
	return (*connectionStatePtrType)(v)
}

func (*connectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatePtrOutput)
}

type ConnectionStateOutput struct{ *pulumi.OutputState }

func (ConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (o ConnectionStateOutput) ToConnectionStateOutput() ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionState) *ConnectionState {
		return &v
	}).(ConnectionStatePtrOutput)
}

func (o ConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) Elem() ConnectionStateOutput {
	return o.ApplyT(func(v *ConnectionState) ConnectionState {
		if v != nil {
			return *v
		}
		var ret ConnectionState
		return ret
	}).(ConnectionStateOutput)
}

func (o ConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ConnectionStateResponse struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}

type ConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutput() ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutputWithContext(ctx context.Context) ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutput() ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutputWithContext(ctx context.Context) ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) Elem() ConnectionStateResponseOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) ConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionStateResponse
		return ret
	}).(ConnectionStateResponseOutput)
}

func (o ConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type Destination struct {
	ArchiveNameFormat        *string `pulumi:"archiveNameFormat"`
	BlobContainer            *string `pulumi:"blobContainer"`
	Name                     *string `pulumi:"name"`
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
}





type DestinationInput interface {
	pulumi.Input

	ToDestinationOutput() DestinationOutput
	ToDestinationOutputWithContext(context.Context) DestinationOutput
}

type DestinationArgs struct {
	ArchiveNameFormat        pulumi.StringPtrInput `pulumi:"archiveNameFormat"`
	BlobContainer            pulumi.StringPtrInput `pulumi:"blobContainer"`
	Name                     pulumi.StringPtrInput `pulumi:"name"`
	StorageAccountResourceId pulumi.StringPtrInput `pulumi:"storageAccountResourceId"`
}

func (DestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Destination)(nil)).Elem()
}

func (i DestinationArgs) ToDestinationOutput() DestinationOutput {
	return i.ToDestinationOutputWithContext(context.Background())
}

func (i DestinationArgs) ToDestinationOutputWithContext(ctx context.Context) DestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationOutput)
}

func (i DestinationArgs) ToDestinationPtrOutput() DestinationPtrOutput {
	return i.ToDestinationPtrOutputWithContext(context.Background())
}

func (i DestinationArgs) ToDestinationPtrOutputWithContext(ctx context.Context) DestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationOutput).ToDestinationPtrOutputWithContext(ctx)
}









type DestinationPtrInput interface {
	pulumi.Input

	ToDestinationPtrOutput() DestinationPtrOutput
	ToDestinationPtrOutputWithContext(context.Context) DestinationPtrOutput
}

type destinationPtrType DestinationArgs

func DestinationPtr(v *DestinationArgs) DestinationPtrInput {
	return (*destinationPtrType)(v)
}

func (*destinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Destination)(nil)).Elem()
}

func (i *destinationPtrType) ToDestinationPtrOutput() DestinationPtrOutput {
	return i.ToDestinationPtrOutputWithContext(context.Background())
}

func (i *destinationPtrType) ToDestinationPtrOutputWithContext(ctx context.Context) DestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationPtrOutput)
}

type DestinationOutput struct{ *pulumi.OutputState }

func (DestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Destination)(nil)).Elem()
}

func (o DestinationOutput) ToDestinationOutput() DestinationOutput {
	return o
}

func (o DestinationOutput) ToDestinationOutputWithContext(ctx context.Context) DestinationOutput {
	return o
}

func (o DestinationOutput) ToDestinationPtrOutput() DestinationPtrOutput {
	return o.ToDestinationPtrOutputWithContext(context.Background())
}

func (o DestinationOutput) ToDestinationPtrOutputWithContext(ctx context.Context) DestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Destination) *Destination {
		return &v
	}).(DestinationPtrOutput)
}

func (o DestinationOutput) ArchiveNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Destination) *string { return v.ArchiveNameFormat }).(pulumi.StringPtrOutput)
}

func (o DestinationOutput) BlobContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Destination) *string { return v.BlobContainer }).(pulumi.StringPtrOutput)
}

func (o DestinationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Destination) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DestinationOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Destination) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

type DestinationPtrOutput struct{ *pulumi.OutputState }

func (DestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Destination)(nil)).Elem()
}

func (o DestinationPtrOutput) ToDestinationPtrOutput() DestinationPtrOutput {
	return o
}

func (o DestinationPtrOutput) ToDestinationPtrOutputWithContext(ctx context.Context) DestinationPtrOutput {
	return o
}

func (o DestinationPtrOutput) Elem() DestinationOutput {
	return o.ApplyT(func(v *Destination) Destination {
		if v != nil {
			return *v
		}
		var ret Destination
		return ret
	}).(DestinationOutput)
}

func (o DestinationPtrOutput) ArchiveNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Destination) *string {
		if v == nil {
			return nil
		}
		return v.ArchiveNameFormat
	}).(pulumi.StringPtrOutput)
}

func (o DestinationPtrOutput) BlobContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Destination) *string {
		if v == nil {
			return nil
		}
		return v.BlobContainer
	}).(pulumi.StringPtrOutput)
}

func (o DestinationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Destination) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DestinationPtrOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Destination) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountResourceId
	}).(pulumi.StringPtrOutput)
}

type DestinationResponse struct {
	ArchiveNameFormat        *string `pulumi:"archiveNameFormat"`
	BlobContainer            *string `pulumi:"blobContainer"`
	Name                     *string `pulumi:"name"`
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
}

type DestinationResponseOutput struct{ *pulumi.OutputState }

func (DestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationResponse)(nil)).Elem()
}

func (o DestinationResponseOutput) ToDestinationResponseOutput() DestinationResponseOutput {
	return o
}

func (o DestinationResponseOutput) ToDestinationResponseOutputWithContext(ctx context.Context) DestinationResponseOutput {
	return o
}

func (o DestinationResponseOutput) ArchiveNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationResponse) *string { return v.ArchiveNameFormat }).(pulumi.StringPtrOutput)
}

func (o DestinationResponseOutput) BlobContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationResponse) *string { return v.BlobContainer }).(pulumi.StringPtrOutput)
}

func (o DestinationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DestinationResponseOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationResponse) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

type DestinationResponsePtrOutput struct{ *pulumi.OutputState }

func (DestinationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationResponse)(nil)).Elem()
}

func (o DestinationResponsePtrOutput) ToDestinationResponsePtrOutput() DestinationResponsePtrOutput {
	return o
}

func (o DestinationResponsePtrOutput) ToDestinationResponsePtrOutputWithContext(ctx context.Context) DestinationResponsePtrOutput {
	return o
}

func (o DestinationResponsePtrOutput) Elem() DestinationResponseOutput {
	return o.ApplyT(func(v *DestinationResponse) DestinationResponse {
		if v != nil {
			return *v
		}
		var ret DestinationResponse
		return ret
	}).(DestinationResponseOutput)
}

func (o DestinationResponsePtrOutput) ArchiveNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArchiveNameFormat
	}).(pulumi.StringPtrOutput)
}

func (o DestinationResponsePtrOutput) BlobContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.BlobContainer
	}).(pulumi.StringPtrOutput)
}

func (o DestinationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DestinationResponsePtrOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountResourceId
	}).(pulumi.StringPtrOutput)
}

type NWRuleSetIpRules struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}


func (val *NWRuleSetIpRules) Defaults() *NWRuleSetIpRules {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type NWRuleSetIpRulesInput interface {
	pulumi.Input

	ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput
	ToNWRuleSetIpRulesOutputWithContext(context.Context) NWRuleSetIpRulesOutput
}

type NWRuleSetIpRulesArgs struct {
	Action pulumi.StringPtrInput `pulumi:"action"`
	IpMask pulumi.StringPtrInput `pulumi:"ipMask"`
}


func (val *NWRuleSetIpRulesArgs) Defaults() *NWRuleSetIpRulesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		tmp.Action = pulumi.StringPtr("Allow")
	}
	return &tmp
}
func (NWRuleSetIpRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRules)(nil)).Elem()
}

func (i NWRuleSetIpRulesArgs) ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput {
	return i.ToNWRuleSetIpRulesOutputWithContext(context.Background())
}

func (i NWRuleSetIpRulesArgs) ToNWRuleSetIpRulesOutputWithContext(ctx context.Context) NWRuleSetIpRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetIpRulesOutput)
}





type NWRuleSetIpRulesArrayInput interface {
	pulumi.Input

	ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput
	ToNWRuleSetIpRulesArrayOutputWithContext(context.Context) NWRuleSetIpRulesArrayOutput
}

type NWRuleSetIpRulesArray []NWRuleSetIpRulesInput

func (NWRuleSetIpRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRules)(nil)).Elem()
}

func (i NWRuleSetIpRulesArray) ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput {
	return i.ToNWRuleSetIpRulesArrayOutputWithContext(context.Background())
}

func (i NWRuleSetIpRulesArray) ToNWRuleSetIpRulesArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetIpRulesArrayOutput)
}

type NWRuleSetIpRulesOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRules)(nil)).Elem()
}

func (o NWRuleSetIpRulesOutput) ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput {
	return o
}

func (o NWRuleSetIpRulesOutput) ToNWRuleSetIpRulesOutputWithContext(ctx context.Context) NWRuleSetIpRulesOutput {
	return o
}

func (o NWRuleSetIpRulesOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRules) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NWRuleSetIpRulesOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRules) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type NWRuleSetIpRulesArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRules)(nil)).Elem()
}

func (o NWRuleSetIpRulesArrayOutput) ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput {
	return o
}

func (o NWRuleSetIpRulesArrayOutput) ToNWRuleSetIpRulesArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesArrayOutput {
	return o
}

func (o NWRuleSetIpRulesArrayOutput) Index(i pulumi.IntInput) NWRuleSetIpRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetIpRules {
		return vs[0].([]NWRuleSetIpRules)[vs[1].(int)]
	}).(NWRuleSetIpRulesOutput)
}

type NWRuleSetIpRulesResponse struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}


func (val *NWRuleSetIpRulesResponse) Defaults() *NWRuleSetIpRulesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type NWRuleSetIpRulesResponseOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRulesResponse)(nil)).Elem()
}

func (o NWRuleSetIpRulesResponseOutput) ToNWRuleSetIpRulesResponseOutput() NWRuleSetIpRulesResponseOutput {
	return o
}

func (o NWRuleSetIpRulesResponseOutput) ToNWRuleSetIpRulesResponseOutputWithContext(ctx context.Context) NWRuleSetIpRulesResponseOutput {
	return o
}

func (o NWRuleSetIpRulesResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRulesResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NWRuleSetIpRulesResponseOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRulesResponse) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type NWRuleSetIpRulesResponseArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRulesResponse)(nil)).Elem()
}

func (o NWRuleSetIpRulesResponseArrayOutput) ToNWRuleSetIpRulesResponseArrayOutput() NWRuleSetIpRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetIpRulesResponseArrayOutput) ToNWRuleSetIpRulesResponseArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetIpRulesResponseArrayOutput) Index(i pulumi.IntInput) NWRuleSetIpRulesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetIpRulesResponse {
		return vs[0].([]NWRuleSetIpRulesResponse)[vs[1].(int)]
	}).(NWRuleSetIpRulesResponseOutput)
}

type NWRuleSetVirtualNetworkRules struct {
	IgnoreMissingVnetServiceEndpoint *bool   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           *Subnet `pulumi:"subnet"`
}





type NWRuleSetVirtualNetworkRulesInput interface {
	pulumi.Input

	ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput
	ToNWRuleSetVirtualNetworkRulesOutputWithContext(context.Context) NWRuleSetVirtualNetworkRulesOutput
}

type NWRuleSetVirtualNetworkRulesArgs struct {
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           SubnetPtrInput      `pulumi:"subnet"`
}

func (NWRuleSetVirtualNetworkRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (i NWRuleSetVirtualNetworkRulesArgs) ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput {
	return i.ToNWRuleSetVirtualNetworkRulesOutputWithContext(context.Background())
}

func (i NWRuleSetVirtualNetworkRulesArgs) ToNWRuleSetVirtualNetworkRulesOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetVirtualNetworkRulesOutput)
}





type NWRuleSetVirtualNetworkRulesArrayInput interface {
	pulumi.Input

	ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput
	ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(context.Context) NWRuleSetVirtualNetworkRulesArrayOutput
}

type NWRuleSetVirtualNetworkRulesArray []NWRuleSetVirtualNetworkRulesInput

func (NWRuleSetVirtualNetworkRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (i NWRuleSetVirtualNetworkRulesArray) ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput {
	return i.ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(context.Background())
}

func (i NWRuleSetVirtualNetworkRulesArray) ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetVirtualNetworkRulesArrayOutput)
}

type NWRuleSetVirtualNetworkRulesOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesOutput) ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesOutput) ToNWRuleSetVirtualNetworkRulesOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRules) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o NWRuleSetVirtualNetworkRulesOutput) Subnet() SubnetPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRules) *Subnet { return v.Subnet }).(SubnetPtrOutput)
}

type NWRuleSetVirtualNetworkRulesArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) Index(i pulumi.IntInput) NWRuleSetVirtualNetworkRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetVirtualNetworkRules {
		return vs[0].([]NWRuleSetVirtualNetworkRules)[vs[1].(int)]
	}).(NWRuleSetVirtualNetworkRulesOutput)
}

type NWRuleSetVirtualNetworkRulesResponse struct {
	IgnoreMissingVnetServiceEndpoint *bool           `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           *SubnetResponse `pulumi:"subnet"`
}

type NWRuleSetVirtualNetworkRulesResponseOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRulesResponse)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) ToNWRuleSetVirtualNetworkRulesResponseOutput() NWRuleSetVirtualNetworkRulesResponseOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) ToNWRuleSetVirtualNetworkRulesResponseOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesResponseOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRulesResponse) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRulesResponse) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

type NWRuleSetVirtualNetworkRulesResponseArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRulesResponse)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) ToNWRuleSetVirtualNetworkRulesResponseArrayOutput() NWRuleSetVirtualNetworkRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) ToNWRuleSetVirtualNetworkRulesResponseArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) Index(i pulumi.IntInput) NWRuleSetVirtualNetworkRulesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetVirtualNetworkRulesResponse {
		return vs[0].([]NWRuleSetVirtualNetworkRulesResponse)[vs[1].(int)]
	}).(NWRuleSetVirtualNetworkRulesResponseOutput)
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponse struct {
	Id *string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type Subnet struct {
	Id string `pulumi:"id"`
}





type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(context.Context) SubnetOutput
}

type SubnetArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (i SubnetArgs) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i SubnetArgs) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}

func (i SubnetArgs) ToSubnetPtrOutput() SubnetPtrOutput {
	return i.ToSubnetPtrOutputWithContext(context.Background())
}

func (i SubnetArgs) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput).ToSubnetPtrOutputWithContext(ctx)
}









type SubnetPtrInput interface {
	pulumi.Input

	ToSubnetPtrOutput() SubnetPtrOutput
	ToSubnetPtrOutputWithContext(context.Context) SubnetPtrOutput
}

type subnetPtrType SubnetArgs

func SubnetPtr(v *SubnetArgs) SubnetPtrInput {
	return (*subnetPtrType)(v)
}

func (*subnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (i *subnetPtrType) ToSubnetPtrOutput() SubnetPtrOutput {
	return i.ToSubnetPtrOutputWithContext(context.Background())
}

func (i *subnetPtrType) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPtrOutput)
}

type SubnetOutput struct{ *pulumi.OutputState }

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetPtrOutput() SubnetPtrOutput {
	return o.ToSubnetPtrOutputWithContext(context.Background())
}

func (o SubnetOutput) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Subnet) *Subnet {
		return &v
	}).(SubnetPtrOutput)
}

func (o SubnetOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Subnet) string { return v.Id }).(pulumi.StringOutput)
}

type SubnetPtrOutput struct{ *pulumi.OutputState }

func (SubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (o SubnetPtrOutput) ToSubnetPtrOutput() SubnetPtrOutput {
	return o
}

func (o SubnetPtrOutput) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return o
}

func (o SubnetPtrOutput) Elem() SubnetOutput {
	return o.ApplyT(func(v *Subnet) Subnet {
		if v != nil {
			return *v
		}
		var ret Subnet
		return ret
	}).(SubnetOutput)
}

func (o SubnetPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SubnetResponse struct {
	Id string `pulumi:"id"`
}

type SubnetResponseOutput struct{ *pulumi.OutputState }

func (SubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseOutput) ToSubnetResponseOutput() SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetResponse) string { return v.Id }).(pulumi.StringOutput)
}

type SubnetResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetResponse)(nil)).Elem()
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutput() SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutputWithContext(ctx context.Context) SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) Elem() SubnetResponseOutput {
	return o.ApplyT(func(v *SubnetResponse) SubnetResponse {
		if v != nil {
			return *v
		}
		var ret SubnetResponse
		return ret
	}).(SubnetResponseOutput)
}

func (o SubnetResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
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

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type ThrottlingPolicy struct {
	MetricId           string  `pulumi:"metricId"`
	Name               string  `pulumi:"name"`
	RateLimitThreshold float64 `pulumi:"rateLimitThreshold"`
	Type               string  `pulumi:"type"`
}





type ThrottlingPolicyInput interface {
	pulumi.Input

	ToThrottlingPolicyOutput() ThrottlingPolicyOutput
	ToThrottlingPolicyOutputWithContext(context.Context) ThrottlingPolicyOutput
}

type ThrottlingPolicyArgs struct {
	MetricId           pulumi.StringInput  `pulumi:"metricId"`
	Name               pulumi.StringInput  `pulumi:"name"`
	RateLimitThreshold pulumi.Float64Input `pulumi:"rateLimitThreshold"`
	Type               pulumi.StringInput  `pulumi:"type"`
}

func (ThrottlingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingPolicy)(nil)).Elem()
}

func (i ThrottlingPolicyArgs) ToThrottlingPolicyOutput() ThrottlingPolicyOutput {
	return i.ToThrottlingPolicyOutputWithContext(context.Background())
}

func (i ThrottlingPolicyArgs) ToThrottlingPolicyOutputWithContext(ctx context.Context) ThrottlingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingPolicyOutput)
}





type ThrottlingPolicyArrayInput interface {
	pulumi.Input

	ToThrottlingPolicyArrayOutput() ThrottlingPolicyArrayOutput
	ToThrottlingPolicyArrayOutputWithContext(context.Context) ThrottlingPolicyArrayOutput
}

type ThrottlingPolicyArray []ThrottlingPolicyInput

func (ThrottlingPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThrottlingPolicy)(nil)).Elem()
}

func (i ThrottlingPolicyArray) ToThrottlingPolicyArrayOutput() ThrottlingPolicyArrayOutput {
	return i.ToThrottlingPolicyArrayOutputWithContext(context.Background())
}

func (i ThrottlingPolicyArray) ToThrottlingPolicyArrayOutputWithContext(ctx context.Context) ThrottlingPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingPolicyArrayOutput)
}

type ThrottlingPolicyOutput struct{ *pulumi.OutputState }

func (ThrottlingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingPolicy)(nil)).Elem()
}

func (o ThrottlingPolicyOutput) ToThrottlingPolicyOutput() ThrottlingPolicyOutput {
	return o
}

func (o ThrottlingPolicyOutput) ToThrottlingPolicyOutputWithContext(ctx context.Context) ThrottlingPolicyOutput {
	return o
}

func (o ThrottlingPolicyOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v ThrottlingPolicy) string { return v.MetricId }).(pulumi.StringOutput)
}

func (o ThrottlingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ThrottlingPolicy) string { return v.Name }).(pulumi.StringOutput)
}

func (o ThrottlingPolicyOutput) RateLimitThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v ThrottlingPolicy) float64 { return v.RateLimitThreshold }).(pulumi.Float64Output)
}

func (o ThrottlingPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ThrottlingPolicy) string { return v.Type }).(pulumi.StringOutput)
}

type ThrottlingPolicyArrayOutput struct{ *pulumi.OutputState }

func (ThrottlingPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThrottlingPolicy)(nil)).Elem()
}

func (o ThrottlingPolicyArrayOutput) ToThrottlingPolicyArrayOutput() ThrottlingPolicyArrayOutput {
	return o
}

func (o ThrottlingPolicyArrayOutput) ToThrottlingPolicyArrayOutputWithContext(ctx context.Context) ThrottlingPolicyArrayOutput {
	return o
}

func (o ThrottlingPolicyArrayOutput) Index(i pulumi.IntInput) ThrottlingPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThrottlingPolicy {
		return vs[0].([]ThrottlingPolicy)[vs[1].(int)]
	}).(ThrottlingPolicyOutput)
}

type ThrottlingPolicyResponse struct {
	MetricId           string  `pulumi:"metricId"`
	Name               string  `pulumi:"name"`
	RateLimitThreshold float64 `pulumi:"rateLimitThreshold"`
	Type               string  `pulumi:"type"`
}

type ThrottlingPolicyResponseOutput struct{ *pulumi.OutputState }

func (ThrottlingPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingPolicyResponse)(nil)).Elem()
}

func (o ThrottlingPolicyResponseOutput) ToThrottlingPolicyResponseOutput() ThrottlingPolicyResponseOutput {
	return o
}

func (o ThrottlingPolicyResponseOutput) ToThrottlingPolicyResponseOutputWithContext(ctx context.Context) ThrottlingPolicyResponseOutput {
	return o
}

func (o ThrottlingPolicyResponseOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v ThrottlingPolicyResponse) string { return v.MetricId }).(pulumi.StringOutput)
}

func (o ThrottlingPolicyResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ThrottlingPolicyResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ThrottlingPolicyResponseOutput) RateLimitThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v ThrottlingPolicyResponse) float64 { return v.RateLimitThreshold }).(pulumi.Float64Output)
}

func (o ThrottlingPolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ThrottlingPolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ThrottlingPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (ThrottlingPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThrottlingPolicyResponse)(nil)).Elem()
}

func (o ThrottlingPolicyResponseArrayOutput) ToThrottlingPolicyResponseArrayOutput() ThrottlingPolicyResponseArrayOutput {
	return o
}

func (o ThrottlingPolicyResponseArrayOutput) ToThrottlingPolicyResponseArrayOutputWithContext(ctx context.Context) ThrottlingPolicyResponseArrayOutput {
	return o
}

func (o ThrottlingPolicyResponseArrayOutput) Index(i pulumi.IntInput) ThrottlingPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThrottlingPolicyResponse {
		return vs[0].([]ThrottlingPolicyResponse)[vs[1].(int)]
	}).(ThrottlingPolicyResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CaptureDescriptionOutput{})
	pulumi.RegisterOutputType(CaptureDescriptionPtrOutput{})
	pulumi.RegisterOutputType(CaptureDescriptionResponseOutput{})
	pulumi.RegisterOutputType(CaptureDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterSkuOutput{})
	pulumi.RegisterOutputType(ClusterSkuPtrOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponseOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(DestinationOutput{})
	pulumi.RegisterOutputType(DestinationPtrOutput{})
	pulumi.RegisterOutputType(DestinationResponseOutput{})
	pulumi.RegisterOutputType(DestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesResponseArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetOutput{})
	pulumi.RegisterOutputType(SubnetPtrOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(ThrottlingPolicyOutput{})
	pulumi.RegisterOutputType(ThrottlingPolicyArrayOutput{})
	pulumi.RegisterOutputType(ThrottlingPolicyResponseOutput{})
	pulumi.RegisterOutputType(ThrottlingPolicyResponseArrayOutput{})
}
