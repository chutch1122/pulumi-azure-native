


package v20180101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Backend struct {
	pulumi.CustomResourceState

	Credentials BackendCredentialsContractResponsePtrOutput `pulumi:"credentials"`
	Description pulumi.StringPtrOutput                      `pulumi:"description"`
	Name        pulumi.StringOutput                         `pulumi:"name"`
	Properties  BackendPropertiesResponseOutput             `pulumi:"properties"`
	Protocol    pulumi.StringOutput                         `pulumi:"protocol"`
	Proxy       BackendProxyContractResponsePtrOutput       `pulumi:"proxy"`
	ResourceId  pulumi.StringPtrOutput                      `pulumi:"resourceId"`
	Title       pulumi.StringPtrOutput                      `pulumi:"title"`
	Tls         BackendTlsPropertiesResponsePtrOutput       `pulumi:"tls"`
	Type        pulumi.StringOutput                         `pulumi:"type"`
	Url         pulumi.StringOutput                         `pulumi:"url"`
}


func NewBackend(ctx *pulumi.Context,
	name string, args *BackendArgs, opts ...pulumi.ResourceOption) (*Backend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Tls != nil {
		args.Tls = args.Tls.ToBackendTlsPropertiesPtrOutput().ApplyT(func(v *BackendTlsProperties) *BackendTlsProperties { return v.Defaults() }).(BackendTlsPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:Backend"),
		},
	})
	opts = append(opts, aliases)
	var resource Backend
	err := ctx.RegisterResource("azure-native:apimanagement/v20180101:Backend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendState, opts ...pulumi.ResourceOption) (*Backend, error) {
	var resource Backend
	err := ctx.ReadResource("azure-native:apimanagement/v20180101:Backend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type backendState struct {
}

type BackendState struct {
}

func (BackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendState)(nil)).Elem()
}

type backendArgs struct {
	Backendid         *string                     `pulumi:"backendid"`
	Credentials       *BackendCredentialsContract `pulumi:"credentials"`
	Description       *string                     `pulumi:"description"`
	Properties        *BackendProperties          `pulumi:"properties"`
	Protocol          string                      `pulumi:"protocol"`
	Proxy             *BackendProxyContract       `pulumi:"proxy"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
	ResourceId        *string                     `pulumi:"resourceId"`
	ServiceName       string                      `pulumi:"serviceName"`
	Title             *string                     `pulumi:"title"`
	Tls               *BackendTlsProperties       `pulumi:"tls"`
	Url               string                      `pulumi:"url"`
}


type BackendArgs struct {
	Backendid         pulumi.StringPtrInput
	Credentials       BackendCredentialsContractPtrInput
	Description       pulumi.StringPtrInput
	Properties        BackendPropertiesPtrInput
	Protocol          pulumi.StringInput
	Proxy             BackendProxyContractPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceId        pulumi.StringPtrInput
	ServiceName       pulumi.StringInput
	Title             pulumi.StringPtrInput
	Tls               BackendTlsPropertiesPtrInput
	Url               pulumi.StringInput
}

func (BackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendArgs)(nil)).Elem()
}

type BackendInput interface {
	pulumi.Input

	ToBackendOutput() BackendOutput
	ToBackendOutputWithContext(ctx context.Context) BackendOutput
}

func (*Backend) ElementType() reflect.Type {
	return reflect.TypeOf((**Backend)(nil)).Elem()
}

func (i *Backend) ToBackendOutput() BackendOutput {
	return i.ToBackendOutputWithContext(context.Background())
}

func (i *Backend) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendOutput)
}

type BackendOutput struct{ *pulumi.OutputState }

func (BackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backend)(nil)).Elem()
}

func (o BackendOutput) ToBackendOutput() BackendOutput {
	return o
}

func (o BackendOutput) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return o
}

func (o BackendOutput) Credentials() BackendCredentialsContractResponsePtrOutput {
	return o.ApplyT(func(v *Backend) BackendCredentialsContractResponsePtrOutput { return v.Credentials }).(BackendCredentialsContractResponsePtrOutput)
}

func (o BackendOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BackendOutput) Properties() BackendPropertiesResponseOutput {
	return o.ApplyT(func(v *Backend) BackendPropertiesResponseOutput { return v.Properties }).(BackendPropertiesResponseOutput)
}

func (o BackendOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o BackendOutput) Proxy() BackendProxyContractResponsePtrOutput {
	return o.ApplyT(func(v *Backend) BackendProxyContractResponsePtrOutput { return v.Proxy }).(BackendProxyContractResponsePtrOutput)
}

func (o BackendOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) Tls() BackendTlsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Backend) BackendTlsPropertiesResponsePtrOutput { return v.Tls }).(BackendTlsPropertiesResponsePtrOutput)
}

func (o BackendOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o BackendOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BackendOutput{})
}
