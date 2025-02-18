


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListPrivateStoreNewPlansNotifications(ctx *pulumi.Context, args *ListPrivateStoreNewPlansNotificationsArgs, opts ...pulumi.InvokeOption) (*ListPrivateStoreNewPlansNotificationsResult, error) {
	var rv ListPrivateStoreNewPlansNotificationsResult
	err := ctx.Invoke("azure-native:marketplace/v20220301:listPrivateStoreNewPlansNotifications", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPrivateStoreNewPlansNotificationsArgs struct {
	PrivateStoreId string `pulumi:"privateStoreId"`
}


type ListPrivateStoreNewPlansNotificationsResult struct {
	NewPlansNotifications []NewNotificationsResponse `pulumi:"newPlansNotifications"`
}

func ListPrivateStoreNewPlansNotificationsOutput(ctx *pulumi.Context, args ListPrivateStoreNewPlansNotificationsOutputArgs, opts ...pulumi.InvokeOption) ListPrivateStoreNewPlansNotificationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPrivateStoreNewPlansNotificationsResult, error) {
			args := v.(ListPrivateStoreNewPlansNotificationsArgs)
			r, err := ListPrivateStoreNewPlansNotifications(ctx, &args, opts...)
			var s ListPrivateStoreNewPlansNotificationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPrivateStoreNewPlansNotificationsResultOutput)
}

type ListPrivateStoreNewPlansNotificationsOutputArgs struct {
	PrivateStoreId pulumi.StringInput `pulumi:"privateStoreId"`
}

func (ListPrivateStoreNewPlansNotificationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateStoreNewPlansNotificationsArgs)(nil)).Elem()
}


type ListPrivateStoreNewPlansNotificationsResultOutput struct{ *pulumi.OutputState }

func (ListPrivateStoreNewPlansNotificationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateStoreNewPlansNotificationsResult)(nil)).Elem()
}

func (o ListPrivateStoreNewPlansNotificationsResultOutput) ToListPrivateStoreNewPlansNotificationsResultOutput() ListPrivateStoreNewPlansNotificationsResultOutput {
	return o
}

func (o ListPrivateStoreNewPlansNotificationsResultOutput) ToListPrivateStoreNewPlansNotificationsResultOutputWithContext(ctx context.Context) ListPrivateStoreNewPlansNotificationsResultOutput {
	return o
}

func (o ListPrivateStoreNewPlansNotificationsResultOutput) NewPlansNotifications() NewNotificationsResponseArrayOutput {
	return o.ApplyT(func(v ListPrivateStoreNewPlansNotificationsResult) []NewNotificationsResponse {
		return v.NewPlansNotifications
	}).(NewNotificationsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPrivateStoreNewPlansNotificationsResultOutput{})
}
