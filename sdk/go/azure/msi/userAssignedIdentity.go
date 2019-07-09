// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package msi

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a user assigned identity.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/user_assigned_identity.html.markdown.
type UserAssignedIdentity struct {
	s *pulumi.ResourceState
}

// NewUserAssignedIdentity registers a new resource with the given unique name, arguments, and options.
func NewUserAssignedIdentity(ctx *pulumi.Context,
	name string, args *UserAssignedIdentityArgs, opts ...pulumi.ResourceOpt) (*UserAssignedIdentity, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
	} else {
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
	}
	inputs["clientId"] = nil
	inputs["principalId"] = nil
	s, err := ctx.RegisterResource("azure:msi/userAssignedIdentity:UserAssignedIdentity", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserAssignedIdentity{s: s}, nil
}

// GetUserAssignedIdentity gets an existing UserAssignedIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserAssignedIdentity(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserAssignedIdentityState, opts ...pulumi.ResourceOpt) (*UserAssignedIdentity, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["clientId"] = state.ClientId
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["principalId"] = state.PrincipalId
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:msi/userAssignedIdentity:UserAssignedIdentity", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserAssignedIdentity{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *UserAssignedIdentity) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *UserAssignedIdentity) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Client ID associated with the user assigned identity.
func (r *UserAssignedIdentity) ClientId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clientId"])
}

// The location/region where the user assigned identity is
// created.
func (r *UserAssignedIdentity) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The name of the user assigned identity. Changing this forces a
// new identity to be created.
func (r *UserAssignedIdentity) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Service Principal ID associated with the user assigned identity.
func (r *UserAssignedIdentity) PrincipalId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["principalId"])
}

// The name of the resource group in which to
// create the user assigned identity.
func (r *UserAssignedIdentity) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *UserAssignedIdentity) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering UserAssignedIdentity resources.
type UserAssignedIdentityState struct {
	// Client ID associated with the user assigned identity.
	ClientId interface{}
	// The location/region where the user assigned identity is
	// created.
	Location interface{}
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name interface{}
	// Service Principal ID associated with the user assigned identity.
	PrincipalId interface{}
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a UserAssignedIdentity resource.
type UserAssignedIdentityArgs struct {
	// The location/region where the user assigned identity is
	// created.
	Location interface{}
	// The name of the user assigned identity. Changing this forces a
	// new identity to be created.
	Name interface{}
	// The name of the resource group in which to
	// create the user assigned identity.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
