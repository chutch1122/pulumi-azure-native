


package v20200921preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MsixPackageApplications struct {
	AppId          *string `pulumi:"appId"`
	AppUserModelID *string `pulumi:"appUserModelID"`
	Description    *string `pulumi:"description"`
	FriendlyName   *string `pulumi:"friendlyName"`
	IconImageName  *string `pulumi:"iconImageName"`
	RawIcon        *string `pulumi:"rawIcon"`
	RawPng         *string `pulumi:"rawPng"`
}





type MsixPackageApplicationsInput interface {
	pulumi.Input

	ToMsixPackageApplicationsOutput() MsixPackageApplicationsOutput
	ToMsixPackageApplicationsOutputWithContext(context.Context) MsixPackageApplicationsOutput
}

type MsixPackageApplicationsArgs struct {
	AppId          pulumi.StringPtrInput `pulumi:"appId"`
	AppUserModelID pulumi.StringPtrInput `pulumi:"appUserModelID"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	FriendlyName   pulumi.StringPtrInput `pulumi:"friendlyName"`
	IconImageName  pulumi.StringPtrInput `pulumi:"iconImageName"`
	RawIcon        pulumi.StringPtrInput `pulumi:"rawIcon"`
	RawPng         pulumi.StringPtrInput `pulumi:"rawPng"`
}

func (MsixPackageApplicationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageApplications)(nil)).Elem()
}

func (i MsixPackageApplicationsArgs) ToMsixPackageApplicationsOutput() MsixPackageApplicationsOutput {
	return i.ToMsixPackageApplicationsOutputWithContext(context.Background())
}

func (i MsixPackageApplicationsArgs) ToMsixPackageApplicationsOutputWithContext(ctx context.Context) MsixPackageApplicationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageApplicationsOutput)
}





type MsixPackageApplicationsArrayInput interface {
	pulumi.Input

	ToMsixPackageApplicationsArrayOutput() MsixPackageApplicationsArrayOutput
	ToMsixPackageApplicationsArrayOutputWithContext(context.Context) MsixPackageApplicationsArrayOutput
}

type MsixPackageApplicationsArray []MsixPackageApplicationsInput

func (MsixPackageApplicationsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageApplications)(nil)).Elem()
}

func (i MsixPackageApplicationsArray) ToMsixPackageApplicationsArrayOutput() MsixPackageApplicationsArrayOutput {
	return i.ToMsixPackageApplicationsArrayOutputWithContext(context.Background())
}

func (i MsixPackageApplicationsArray) ToMsixPackageApplicationsArrayOutputWithContext(ctx context.Context) MsixPackageApplicationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageApplicationsArrayOutput)
}

type MsixPackageApplicationsOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageApplications)(nil)).Elem()
}

func (o MsixPackageApplicationsOutput) ToMsixPackageApplicationsOutput() MsixPackageApplicationsOutput {
	return o
}

func (o MsixPackageApplicationsOutput) ToMsixPackageApplicationsOutputWithContext(ctx context.Context) MsixPackageApplicationsOutput {
	return o
}

func (o MsixPackageApplicationsOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) AppUserModelID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.AppUserModelID }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) IconImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.IconImageName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) RawIcon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.RawIcon }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) RawPng() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.RawPng }).(pulumi.StringPtrOutput)
}

type MsixPackageApplicationsArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageApplications)(nil)).Elem()
}

func (o MsixPackageApplicationsArrayOutput) ToMsixPackageApplicationsArrayOutput() MsixPackageApplicationsArrayOutput {
	return o
}

func (o MsixPackageApplicationsArrayOutput) ToMsixPackageApplicationsArrayOutputWithContext(ctx context.Context) MsixPackageApplicationsArrayOutput {
	return o
}

func (o MsixPackageApplicationsArrayOutput) Index(i pulumi.IntInput) MsixPackageApplicationsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageApplications {
		return vs[0].([]MsixPackageApplications)[vs[1].(int)]
	}).(MsixPackageApplicationsOutput)
}

type MsixPackageApplicationsResponse struct {
	AppId          *string `pulumi:"appId"`
	AppUserModelID *string `pulumi:"appUserModelID"`
	Description    *string `pulumi:"description"`
	FriendlyName   *string `pulumi:"friendlyName"`
	IconImageName  *string `pulumi:"iconImageName"`
	RawIcon        *string `pulumi:"rawIcon"`
	RawPng         *string `pulumi:"rawPng"`
}

type MsixPackageApplicationsResponseOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageApplicationsResponse)(nil)).Elem()
}

func (o MsixPackageApplicationsResponseOutput) ToMsixPackageApplicationsResponseOutput() MsixPackageApplicationsResponseOutput {
	return o
}

func (o MsixPackageApplicationsResponseOutput) ToMsixPackageApplicationsResponseOutputWithContext(ctx context.Context) MsixPackageApplicationsResponseOutput {
	return o
}

func (o MsixPackageApplicationsResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) AppUserModelID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.AppUserModelID }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) IconImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.IconImageName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) RawIcon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.RawIcon }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) RawPng() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.RawPng }).(pulumi.StringPtrOutput)
}

type MsixPackageApplicationsResponseArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageApplicationsResponse)(nil)).Elem()
}

func (o MsixPackageApplicationsResponseArrayOutput) ToMsixPackageApplicationsResponseArrayOutput() MsixPackageApplicationsResponseArrayOutput {
	return o
}

func (o MsixPackageApplicationsResponseArrayOutput) ToMsixPackageApplicationsResponseArrayOutputWithContext(ctx context.Context) MsixPackageApplicationsResponseArrayOutput {
	return o
}

func (o MsixPackageApplicationsResponseArrayOutput) Index(i pulumi.IntInput) MsixPackageApplicationsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageApplicationsResponse {
		return vs[0].([]MsixPackageApplicationsResponse)[vs[1].(int)]
	}).(MsixPackageApplicationsResponseOutput)
}

type MsixPackageDependencies struct {
	DependencyName *string `pulumi:"dependencyName"`
	MinVersion     *string `pulumi:"minVersion"`
	Publisher      *string `pulumi:"publisher"`
}





type MsixPackageDependenciesInput interface {
	pulumi.Input

	ToMsixPackageDependenciesOutput() MsixPackageDependenciesOutput
	ToMsixPackageDependenciesOutputWithContext(context.Context) MsixPackageDependenciesOutput
}

type MsixPackageDependenciesArgs struct {
	DependencyName pulumi.StringPtrInput `pulumi:"dependencyName"`
	MinVersion     pulumi.StringPtrInput `pulumi:"minVersion"`
	Publisher      pulumi.StringPtrInput `pulumi:"publisher"`
}

func (MsixPackageDependenciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageDependencies)(nil)).Elem()
}

func (i MsixPackageDependenciesArgs) ToMsixPackageDependenciesOutput() MsixPackageDependenciesOutput {
	return i.ToMsixPackageDependenciesOutputWithContext(context.Background())
}

func (i MsixPackageDependenciesArgs) ToMsixPackageDependenciesOutputWithContext(ctx context.Context) MsixPackageDependenciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageDependenciesOutput)
}





type MsixPackageDependenciesArrayInput interface {
	pulumi.Input

	ToMsixPackageDependenciesArrayOutput() MsixPackageDependenciesArrayOutput
	ToMsixPackageDependenciesArrayOutputWithContext(context.Context) MsixPackageDependenciesArrayOutput
}

type MsixPackageDependenciesArray []MsixPackageDependenciesInput

func (MsixPackageDependenciesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageDependencies)(nil)).Elem()
}

func (i MsixPackageDependenciesArray) ToMsixPackageDependenciesArrayOutput() MsixPackageDependenciesArrayOutput {
	return i.ToMsixPackageDependenciesArrayOutputWithContext(context.Background())
}

func (i MsixPackageDependenciesArray) ToMsixPackageDependenciesArrayOutputWithContext(ctx context.Context) MsixPackageDependenciesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageDependenciesArrayOutput)
}

type MsixPackageDependenciesOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageDependencies)(nil)).Elem()
}

func (o MsixPackageDependenciesOutput) ToMsixPackageDependenciesOutput() MsixPackageDependenciesOutput {
	return o
}

func (o MsixPackageDependenciesOutput) ToMsixPackageDependenciesOutputWithContext(ctx context.Context) MsixPackageDependenciesOutput {
	return o
}

func (o MsixPackageDependenciesOutput) DependencyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependencies) *string { return v.DependencyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesOutput) MinVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependencies) *string { return v.MinVersion }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependencies) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type MsixPackageDependenciesArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageDependencies)(nil)).Elem()
}

func (o MsixPackageDependenciesArrayOutput) ToMsixPackageDependenciesArrayOutput() MsixPackageDependenciesArrayOutput {
	return o
}

func (o MsixPackageDependenciesArrayOutput) ToMsixPackageDependenciesArrayOutputWithContext(ctx context.Context) MsixPackageDependenciesArrayOutput {
	return o
}

func (o MsixPackageDependenciesArrayOutput) Index(i pulumi.IntInput) MsixPackageDependenciesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageDependencies {
		return vs[0].([]MsixPackageDependencies)[vs[1].(int)]
	}).(MsixPackageDependenciesOutput)
}

type MsixPackageDependenciesResponse struct {
	DependencyName *string `pulumi:"dependencyName"`
	MinVersion     *string `pulumi:"minVersion"`
	Publisher      *string `pulumi:"publisher"`
}

type MsixPackageDependenciesResponseOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageDependenciesResponse)(nil)).Elem()
}

func (o MsixPackageDependenciesResponseOutput) ToMsixPackageDependenciesResponseOutput() MsixPackageDependenciesResponseOutput {
	return o
}

func (o MsixPackageDependenciesResponseOutput) ToMsixPackageDependenciesResponseOutputWithContext(ctx context.Context) MsixPackageDependenciesResponseOutput {
	return o
}

func (o MsixPackageDependenciesResponseOutput) DependencyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependenciesResponse) *string { return v.DependencyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesResponseOutput) MinVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependenciesResponse) *string { return v.MinVersion }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependenciesResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type MsixPackageDependenciesResponseArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageDependenciesResponse)(nil)).Elem()
}

func (o MsixPackageDependenciesResponseArrayOutput) ToMsixPackageDependenciesResponseArrayOutput() MsixPackageDependenciesResponseArrayOutput {
	return o
}

func (o MsixPackageDependenciesResponseArrayOutput) ToMsixPackageDependenciesResponseArrayOutputWithContext(ctx context.Context) MsixPackageDependenciesResponseArrayOutput {
	return o
}

func (o MsixPackageDependenciesResponseArrayOutput) Index(i pulumi.IntInput) MsixPackageDependenciesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageDependenciesResponse {
		return vs[0].([]MsixPackageDependenciesResponse)[vs[1].(int)]
	}).(MsixPackageDependenciesResponseOutput)
}

type RegistrationInfo struct {
	ExpirationTime             *string `pulumi:"expirationTime"`
	RegistrationTokenOperation *string `pulumi:"registrationTokenOperation"`
	Token                      *string `pulumi:"token"`
}





type RegistrationInfoInput interface {
	pulumi.Input

	ToRegistrationInfoOutput() RegistrationInfoOutput
	ToRegistrationInfoOutputWithContext(context.Context) RegistrationInfoOutput
}

type RegistrationInfoArgs struct {
	ExpirationTime             pulumi.StringPtrInput `pulumi:"expirationTime"`
	RegistrationTokenOperation pulumi.StringPtrInput `pulumi:"registrationTokenOperation"`
	Token                      pulumi.StringPtrInput `pulumi:"token"`
}

func (RegistrationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfo)(nil)).Elem()
}

func (i RegistrationInfoArgs) ToRegistrationInfoOutput() RegistrationInfoOutput {
	return i.ToRegistrationInfoOutputWithContext(context.Background())
}

func (i RegistrationInfoArgs) ToRegistrationInfoOutputWithContext(ctx context.Context) RegistrationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoOutput)
}

func (i RegistrationInfoArgs) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return i.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (i RegistrationInfoArgs) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoOutput).ToRegistrationInfoPtrOutputWithContext(ctx)
}









type RegistrationInfoPtrInput interface {
	pulumi.Input

	ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput
	ToRegistrationInfoPtrOutputWithContext(context.Context) RegistrationInfoPtrOutput
}

type registrationInfoPtrType RegistrationInfoArgs

func RegistrationInfoPtr(v *RegistrationInfoArgs) RegistrationInfoPtrInput {
	return (*registrationInfoPtrType)(v)
}

func (*registrationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfo)(nil)).Elem()
}

func (i *registrationInfoPtrType) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return i.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (i *registrationInfoPtrType) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoPtrOutput)
}

type RegistrationInfoOutput struct{ *pulumi.OutputState }

func (RegistrationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfo)(nil)).Elem()
}

func (o RegistrationInfoOutput) ToRegistrationInfoOutput() RegistrationInfoOutput {
	return o
}

func (o RegistrationInfoOutput) ToRegistrationInfoOutputWithContext(ctx context.Context) RegistrationInfoOutput {
	return o
}

func (o RegistrationInfoOutput) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return o.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (o RegistrationInfoOutput) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationInfo) *RegistrationInfo {
		return &v
	}).(RegistrationInfoPtrOutput)
}

func (o RegistrationInfoOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *string { return v.ExpirationTime }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *string { return v.RegistrationTokenOperation }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type RegistrationInfoPtrOutput struct{ *pulumi.OutputState }

func (RegistrationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfo)(nil)).Elem()
}

func (o RegistrationInfoPtrOutput) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return o
}

func (o RegistrationInfoPtrOutput) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return o
}

func (o RegistrationInfoPtrOutput) Elem() RegistrationInfoOutput {
	return o.ApplyT(func(v *RegistrationInfo) RegistrationInfo {
		if v != nil {
			return *v
		}
		var ret RegistrationInfo
		return ret
	}).(RegistrationInfoOutput)
}

func (o RegistrationInfoPtrOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoPtrOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationTokenOperation
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoPtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

type RegistrationInfoResponse struct {
	ExpirationTime             *string `pulumi:"expirationTime"`
	RegistrationTokenOperation *string `pulumi:"registrationTokenOperation"`
	Token                      *string `pulumi:"token"`
}

type RegistrationInfoResponseOutput struct{ *pulumi.OutputState }

func (RegistrationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfoResponse)(nil)).Elem()
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponseOutput() RegistrationInfoResponseOutput {
	return o
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponseOutputWithContext(ctx context.Context) RegistrationInfoResponseOutput {
	return o
}

func (o RegistrationInfoResponseOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *string { return v.ExpirationTime }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponseOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *string { return v.RegistrationTokenOperation }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponseOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type RegistrationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (RegistrationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfoResponse)(nil)).Elem()
}

func (o RegistrationInfoResponsePtrOutput) ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput {
	return o
}

func (o RegistrationInfoResponsePtrOutput) ToRegistrationInfoResponsePtrOutputWithContext(ctx context.Context) RegistrationInfoResponsePtrOutput {
	return o
}

func (o RegistrationInfoResponsePtrOutput) Elem() RegistrationInfoResponseOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) RegistrationInfoResponse {
		if v != nil {
			return *v
		}
		var ret RegistrationInfoResponse
		return ret
	}).(RegistrationInfoResponseOutput)
}

func (o RegistrationInfoResponsePtrOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponsePtrOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationTokenOperation
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponsePtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MsixPackageApplicationsOutput{})
	pulumi.RegisterOutputType(MsixPackageApplicationsArrayOutput{})
	pulumi.RegisterOutputType(MsixPackageApplicationsResponseOutput{})
	pulumi.RegisterOutputType(MsixPackageApplicationsResponseArrayOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesArrayOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesResponseOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesResponseArrayOutput{})
	pulumi.RegisterOutputType(RegistrationInfoOutput{})
	pulumi.RegisterOutputType(RegistrationInfoPtrOutput{})
	pulumi.RegisterOutputType(RegistrationInfoResponseOutput{})
	pulumi.RegisterOutputType(RegistrationInfoResponsePtrOutput{})
}
