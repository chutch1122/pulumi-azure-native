


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Queue struct {
	pulumi.CustomResourceState

	ApproximateMessageCount pulumi.IntOutput       `pulumi:"approximateMessageCount"`
	Metadata                pulumi.StringMapOutput `pulumi:"metadata"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	Type                    pulumi.StringOutput    `pulumi:"type"`
}


func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:Queue"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:Queue"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:Queue"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:Queue"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:Queue"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:Queue"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:Queue"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:Queue"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:Queue"),
		},
	})
	opts = append(opts, aliases)
	var resource Queue
	err := ctx.RegisterResource("azure-native:storage/v20210101:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("azure-native:storage/v20210101:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type queueState struct {
}

type QueueState struct {
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	AccountName       string            `pulumi:"accountName"`
	Metadata          map[string]string `pulumi:"metadata"`
	QueueName         *string           `pulumi:"queueName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
}


type QueueArgs struct {
	AccountName       pulumi.StringInput
	Metadata          pulumi.StringMapInput
	QueueName         pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (*Queue) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

type QueueOutput struct{ *pulumi.OutputState }

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

func (o QueueOutput) ApproximateMessageCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntOutput { return v.ApproximateMessageCount }).(pulumi.IntOutput)
}

func (o QueueOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o QueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o QueueOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueOutput{})
}
