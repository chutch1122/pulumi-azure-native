


package v20180201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BuildTask struct {
	pulumi.CustomResourceState

	Alias             pulumi.StringOutput                      `pulumi:"alias"`
	CreationDate      pulumi.StringOutput                      `pulumi:"creationDate"`
	Location          pulumi.StringOutput                      `pulumi:"location"`
	Name              pulumi.StringOutput                      `pulumi:"name"`
	Platform          PlatformPropertiesResponseOutput         `pulumi:"platform"`
	ProvisioningState pulumi.StringOutput                      `pulumi:"provisioningState"`
	SourceRepository  SourceRepositoryPropertiesResponseOutput `pulumi:"sourceRepository"`
	Status            pulumi.StringPtrOutput                   `pulumi:"status"`
	Tags              pulumi.StringMapOutput                   `pulumi:"tags"`
	Timeout           pulumi.IntPtrOutput                      `pulumi:"timeout"`
	Type              pulumi.StringOutput                      `pulumi:"type"`
}


func NewBuildTask(ctx *pulumi.Context,
	name string, args *BuildTaskArgs, opts ...pulumi.ResourceOption) (*BuildTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.Platform == nil {
		return nil, errors.New("invalid value for required argument 'Platform'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceRepository == nil {
		return nil, errors.New("invalid value for required argument 'SourceRepository'")
	}
	args.SourceRepository = args.SourceRepository.ToSourceRepositoryPropertiesOutput().ApplyT(func(v SourceRepositoryProperties) SourceRepositoryProperties { return *v.Defaults() }).(SourceRepositoryPropertiesOutput)
	if isZero(args.Timeout) {
		args.Timeout = pulumi.IntPtr(3600)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:BuildTask"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20180901:BuildTask"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190401:BuildTask"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190601preview:BuildTask"),
		},
	})
	opts = append(opts, aliases)
	var resource BuildTask
	err := ctx.RegisterResource("azure-native:containerregistry/v20180201preview:BuildTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBuildTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BuildTaskState, opts ...pulumi.ResourceOption) (*BuildTask, error) {
	var resource BuildTask
	err := ctx.ReadResource("azure-native:containerregistry/v20180201preview:BuildTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type buildTaskState struct {
}

type BuildTaskState struct {
}

func (BuildTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*buildTaskState)(nil)).Elem()
}

type buildTaskArgs struct {
	Alias             string                     `pulumi:"alias"`
	BuildTaskName     *string                    `pulumi:"buildTaskName"`
	Location          *string                    `pulumi:"location"`
	Platform          PlatformProperties         `pulumi:"platform"`
	RegistryName      string                     `pulumi:"registryName"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	SourceRepository  SourceRepositoryProperties `pulumi:"sourceRepository"`
	Status            *string                    `pulumi:"status"`
	Tags              map[string]string          `pulumi:"tags"`
	Timeout           *int                       `pulumi:"timeout"`
}


type BuildTaskArgs struct {
	Alias             pulumi.StringInput
	BuildTaskName     pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Platform          PlatformPropertiesInput
	RegistryName      pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SourceRepository  SourceRepositoryPropertiesInput
	Status            pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	Timeout           pulumi.IntPtrInput
}

func (BuildTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*buildTaskArgs)(nil)).Elem()
}

type BuildTaskInput interface {
	pulumi.Input

	ToBuildTaskOutput() BuildTaskOutput
	ToBuildTaskOutputWithContext(ctx context.Context) BuildTaskOutput
}

func (*BuildTask) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildTask)(nil)).Elem()
}

func (i *BuildTask) ToBuildTaskOutput() BuildTaskOutput {
	return i.ToBuildTaskOutputWithContext(context.Background())
}

func (i *BuildTask) ToBuildTaskOutputWithContext(ctx context.Context) BuildTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildTaskOutput)
}

type BuildTaskOutput struct{ *pulumi.OutputState }

func (BuildTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildTask)(nil)).Elem()
}

func (o BuildTaskOutput) ToBuildTaskOutput() BuildTaskOutput {
	return o
}

func (o BuildTaskOutput) ToBuildTaskOutputWithContext(ctx context.Context) BuildTaskOutput {
	return o
}

func (o BuildTaskOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildTask) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

func (o BuildTaskOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildTask) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o BuildTaskOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildTask) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BuildTaskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildTask) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BuildTaskOutput) Platform() PlatformPropertiesResponseOutput {
	return o.ApplyT(func(v *BuildTask) PlatformPropertiesResponseOutput { return v.Platform }).(PlatformPropertiesResponseOutput)
}

func (o BuildTaskOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildTask) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o BuildTaskOutput) SourceRepository() SourceRepositoryPropertiesResponseOutput {
	return o.ApplyT(func(v *BuildTask) SourceRepositoryPropertiesResponseOutput { return v.SourceRepository }).(SourceRepositoryPropertiesResponseOutput)
}

func (o BuildTaskOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildTask) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o BuildTaskOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BuildTask) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BuildTaskOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BuildTask) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o BuildTaskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BuildTask) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildTaskOutput{})
}
