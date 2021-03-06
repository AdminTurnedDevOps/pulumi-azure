// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package msi

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a user assigned identity.
type UserAssignedIdentity struct {
	pulumi.CustomResourceState

	// Client ID associated with the user assigned identity.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The location/region where the user assigned identity is
	// created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Service Principal ID associated with the user assigned identity.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewUserAssignedIdentity registers a new resource with the given unique name, arguments, and options.
func NewUserAssignedIdentity(ctx *pulumi.Context,
	name string, args *UserAssignedIdentityArgs, opts ...pulumi.ResourceOption) (*UserAssignedIdentity, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &UserAssignedIdentityArgs{}
	}
	var resource UserAssignedIdentity
	err := ctx.RegisterResource("azure:msi/userAssignedIdentity:UserAssignedIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserAssignedIdentity gets an existing UserAssignedIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserAssignedIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserAssignedIdentityState, opts ...pulumi.ResourceOption) (*UserAssignedIdentity, error) {
	var resource UserAssignedIdentity
	err := ctx.ReadResource("azure:msi/userAssignedIdentity:UserAssignedIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserAssignedIdentity resources.
type userAssignedIdentityState struct {
	// Client ID associated with the user assigned identity.
	ClientId *string `pulumi:"clientId"`
	// The location/region where the user assigned identity is
	// created.
	Location *string `pulumi:"location"`
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name *string `pulumi:"name"`
	// Service Principal ID associated with the user assigned identity.
	PrincipalId *string `pulumi:"principalId"`
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type UserAssignedIdentityState struct {
	// Client ID associated with the user assigned identity.
	ClientId pulumi.StringPtrInput
	// The location/region where the user assigned identity is
	// created.
	Location pulumi.StringPtrInput
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name pulumi.StringPtrInput
	// Service Principal ID associated with the user assigned identity.
	PrincipalId pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (UserAssignedIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*userAssignedIdentityState)(nil)).Elem()
}

type userAssignedIdentityArgs struct {
	// The location/region where the user assigned identity is
	// created.
	Location *string `pulumi:"location"`
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a UserAssignedIdentity resource.
type UserAssignedIdentityArgs struct {
	// The location/region where the user assigned identity is
	// created.
	Location pulumi.StringPtrInput
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (UserAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userAssignedIdentityArgs)(nil)).Elem()
}
