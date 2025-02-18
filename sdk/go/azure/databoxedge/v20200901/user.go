


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type User struct {
	pulumi.CustomResourceState

	EncryptedPassword AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"encryptedPassword"`
	Name              pulumi.StringOutput                        `pulumi:"name"`
	ShareAccessRights ShareAccessRightResponseArrayOutput        `pulumi:"shareAccessRights"`
	SystemData        SystemDataResponseOutput                   `pulumi:"systemData"`
	Type              pulumi.StringOutput                        `pulumi:"type"`
	UserType          pulumi.StringPtrOutput                     `pulumi:"userType"`
}


func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:User"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:User"),
		},
	})
	opts = append(opts, aliases)
	var resource User
	err := ctx.RegisterResource("azure-native:databoxedge/v20200901:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("azure-native:databoxedge/v20200901:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type userState struct {
}

type UserState struct {
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	DeviceName        string                     `pulumi:"deviceName"`
	EncryptedPassword *AsymmetricEncryptedSecret `pulumi:"encryptedPassword"`
	Name              *string                    `pulumi:"name"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	UserType          *string                    `pulumi:"userType"`
}


type UserArgs struct {
	DeviceName        pulumi.StringInput
	EncryptedPassword AsymmetricEncryptedSecretPtrInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	UserType          pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func (o UserOutput) EncryptedPassword() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v *User) AsymmetricEncryptedSecretResponsePtrOutput { return v.EncryptedPassword }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserOutput) ShareAccessRights() ShareAccessRightResponseArrayOutput {
	return o.ApplyT(func(v *User) ShareAccessRightResponseArrayOutput { return v.ShareAccessRights }).(ShareAccessRightResponseArrayOutput)
}

func (o UserOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *User) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o UserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o UserOutput) UserType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.UserType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(UserOutput{})
}
