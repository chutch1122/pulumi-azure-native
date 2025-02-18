


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GuestDiagnosticsSettingsAssociation struct {
	pulumi.CustomResourceState

	GuestDiagnosticSettingsName pulumi.StringOutput    `pulumi:"guestDiagnosticSettingsName"`
	Location                    pulumi.StringOutput    `pulumi:"location"`
	Name                        pulumi.StringOutput    `pulumi:"name"`
	Tags                        pulumi.StringMapOutput `pulumi:"tags"`
	Type                        pulumi.StringOutput    `pulumi:"type"`
}


func NewGuestDiagnosticsSettingsAssociation(ctx *pulumi.Context,
	name string, args *GuestDiagnosticsSettingsAssociationArgs, opts ...pulumi.ResourceOption) (*GuestDiagnosticsSettingsAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GuestDiagnosticSettingsName == nil {
		return nil, errors.New("invalid value for required argument 'GuestDiagnosticSettingsName'")
	}
	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:GuestDiagnosticsSettingsAssociation"),
		},
	})
	opts = append(opts, aliases)
	var resource GuestDiagnosticsSettingsAssociation
	err := ctx.RegisterResource("azure-native:insights/v20180601preview:GuestDiagnosticsSettingsAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGuestDiagnosticsSettingsAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestDiagnosticsSettingsAssociationState, opts ...pulumi.ResourceOption) (*GuestDiagnosticsSettingsAssociation, error) {
	var resource GuestDiagnosticsSettingsAssociation
	err := ctx.ReadResource("azure-native:insights/v20180601preview:GuestDiagnosticsSettingsAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type guestDiagnosticsSettingsAssociationState struct {
}

type GuestDiagnosticsSettingsAssociationState struct {
}

func (GuestDiagnosticsSettingsAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestDiagnosticsSettingsAssociationState)(nil)).Elem()
}

type guestDiagnosticsSettingsAssociationArgs struct {
	AssociationName             *string           `pulumi:"associationName"`
	GuestDiagnosticSettingsName string            `pulumi:"guestDiagnosticSettingsName"`
	Location                    *string           `pulumi:"location"`
	ResourceUri                 string            `pulumi:"resourceUri"`
	Tags                        map[string]string `pulumi:"tags"`
}


type GuestDiagnosticsSettingsAssociationArgs struct {
	AssociationName             pulumi.StringPtrInput
	GuestDiagnosticSettingsName pulumi.StringInput
	Location                    pulumi.StringPtrInput
	ResourceUri                 pulumi.StringInput
	Tags                        pulumi.StringMapInput
}

func (GuestDiagnosticsSettingsAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestDiagnosticsSettingsAssociationArgs)(nil)).Elem()
}

type GuestDiagnosticsSettingsAssociationInput interface {
	pulumi.Input

	ToGuestDiagnosticsSettingsAssociationOutput() GuestDiagnosticsSettingsAssociationOutput
	ToGuestDiagnosticsSettingsAssociationOutputWithContext(ctx context.Context) GuestDiagnosticsSettingsAssociationOutput
}

func (*GuestDiagnosticsSettingsAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestDiagnosticsSettingsAssociation)(nil)).Elem()
}

func (i *GuestDiagnosticsSettingsAssociation) ToGuestDiagnosticsSettingsAssociationOutput() GuestDiagnosticsSettingsAssociationOutput {
	return i.ToGuestDiagnosticsSettingsAssociationOutputWithContext(context.Background())
}

func (i *GuestDiagnosticsSettingsAssociation) ToGuestDiagnosticsSettingsAssociationOutputWithContext(ctx context.Context) GuestDiagnosticsSettingsAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestDiagnosticsSettingsAssociationOutput)
}

type GuestDiagnosticsSettingsAssociationOutput struct{ *pulumi.OutputState }

func (GuestDiagnosticsSettingsAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestDiagnosticsSettingsAssociation)(nil)).Elem()
}

func (o GuestDiagnosticsSettingsAssociationOutput) ToGuestDiagnosticsSettingsAssociationOutput() GuestDiagnosticsSettingsAssociationOutput {
	return o
}

func (o GuestDiagnosticsSettingsAssociationOutput) ToGuestDiagnosticsSettingsAssociationOutputWithContext(ctx context.Context) GuestDiagnosticsSettingsAssociationOutput {
	return o
}

func (o GuestDiagnosticsSettingsAssociationOutput) GuestDiagnosticSettingsName() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestDiagnosticsSettingsAssociation) pulumi.StringOutput { return v.GuestDiagnosticSettingsName }).(pulumi.StringOutput)
}

func (o GuestDiagnosticsSettingsAssociationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestDiagnosticsSettingsAssociation) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GuestDiagnosticsSettingsAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestDiagnosticsSettingsAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GuestDiagnosticsSettingsAssociationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GuestDiagnosticsSettingsAssociation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GuestDiagnosticsSettingsAssociationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestDiagnosticsSettingsAssociation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GuestDiagnosticsSettingsAssociationOutput{})
}
