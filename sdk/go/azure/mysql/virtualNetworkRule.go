// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a MySQL Virtual Network Rule.
// 
// > **NOTE:** MySQL Virtual Network Rules [can only be used with SKU Tiers of `GeneralPurpose` or `MemoryOptimized`](https://docs.microsoft.com/en-us/azure/mysql/concepts-data-access-and-security-vnet)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/mysql_virtual_network_rule.html.markdown.
type VirtualNetworkRule struct {
	s *pulumi.ResourceState
}

// NewVirtualNetworkRule registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkRule(ctx *pulumi.Context,
	name string, args *VirtualNetworkRuleArgs, opts ...pulumi.ResourceOpt) (*VirtualNetworkRule, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["serverName"] = nil
		inputs["subnetId"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["serverName"] = args.ServerName
		inputs["subnetId"] = args.SubnetId
	}
	s, err := ctx.RegisterResource("azure:mysql/virtualNetworkRule:VirtualNetworkRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualNetworkRule{s: s}, nil
}

// GetVirtualNetworkRule gets an existing VirtualNetworkRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VirtualNetworkRuleState, opts ...pulumi.ResourceOpt) (*VirtualNetworkRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["serverName"] = state.ServerName
		inputs["subnetId"] = state.SubnetId
	}
	s, err := ctx.ReadResource("azure:mysql/virtualNetworkRule:VirtualNetworkRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualNetworkRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VirtualNetworkRule) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VirtualNetworkRule) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the MySQL Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
func (r *VirtualNetworkRule) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
func (r *VirtualNetworkRule) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
func (r *VirtualNetworkRule) ServerName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serverName"])
}

// The ID of the subnet that the MySQL server will be connected to.
func (r *VirtualNetworkRule) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// Input properties used for looking up and filtering VirtualNetworkRule resources.
type VirtualNetworkRuleState struct {
	// The name of the MySQL Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName interface{}
	// The ID of the subnet that the MySQL server will be connected to.
	SubnetId interface{}
}

// The set of arguments for constructing a VirtualNetworkRule resource.
type VirtualNetworkRuleArgs struct {
	// The name of the MySQL Virtual Network Rule. Cannot be empty and must only contain alphanumeric characters and hyphens. Cannot start with a number, and cannot start or end with a hyphen. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName interface{}
	// The ID of the subnet that the MySQL server will be connected to.
	SubnetId interface{}
}
