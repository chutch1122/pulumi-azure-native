


package v20200801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cluster struct {
	pulumi.CustomResourceState

	ClusterId          pulumi.StringOutput                 `pulumi:"clusterId"`
	Identity           IdentityResponsePtrOutput           `pulumi:"identity"`
	KeyVaultProperties KeyVaultPropertiesResponsePtrOutput `pulumi:"keyVaultProperties"`
	Location           pulumi.StringOutput                 `pulumi:"location"`
	Name               pulumi.StringOutput                 `pulumi:"name"`
	NextLink           pulumi.StringPtrOutput              `pulumi:"nextLink"`
	ProvisioningState  pulumi.StringOutput                 `pulumi:"provisioningState"`
	Sku                ClusterSkuResponsePtrOutput         `pulumi:"sku"`
	Tags               pulumi.StringMapOutput              `pulumi:"tags"`
	Type               pulumi.StringOutput                 `pulumi:"type"`
}


func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190801preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20201001:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20210601:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azure-native:operationalinsights/v20200801:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure-native:operationalinsights/v20200801:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	ClusterName        *string             `pulumi:"clusterName"`
	Identity           *Identity           `pulumi:"identity"`
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	Location           *string             `pulumi:"location"`
	NextLink           *string             `pulumi:"nextLink"`
	ResourceGroupName  string              `pulumi:"resourceGroupName"`
	Sku                *ClusterSku         `pulumi:"sku"`
	Tags               map[string]string   `pulumi:"tags"`
}


type ClusterArgs struct {
	ClusterName        pulumi.StringPtrInput
	Identity           IdentityPtrInput
	KeyVaultProperties KeyVaultPropertiesPtrInput
	Location           pulumi.StringPtrInput
	NextLink           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	Sku                ClusterSkuPtrInput
	Tags               pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

func (o ClusterOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o ClusterOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) KeyVaultPropertiesResponsePtrOutput { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ClusterOutput) Sku() ClusterSkuResponsePtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterSkuResponsePtrOutput { return v.Sku }).(ClusterSkuResponsePtrOutput)
}

func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
