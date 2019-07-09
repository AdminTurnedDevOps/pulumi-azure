// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Firewall Rule for a PostgreSQL Server
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/postgresql_firewall_rule.html.markdown.
type FirewallRule struct {
	s *pulumi.ResourceState
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOpt) (*FirewallRule, error) {
	if args == nil || args.EndIpAddress == nil {
		return nil, errors.New("missing required argument 'EndIpAddress'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.StartIpAddress == nil {
		return nil, errors.New("missing required argument 'StartIpAddress'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["endIpAddress"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["serverName"] = nil
		inputs["startIpAddress"] = nil
	} else {
		inputs["endIpAddress"] = args.EndIpAddress
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["serverName"] = args.ServerName
		inputs["startIpAddress"] = args.StartIpAddress
	}
	s, err := ctx.RegisterResource("azure:postgresql/firewallRule:FirewallRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FirewallRule{s: s}, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FirewallRuleState, opts ...pulumi.ResourceOpt) (*FirewallRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["endIpAddress"] = state.EndIpAddress
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["serverName"] = state.ServerName
		inputs["startIpAddress"] = state.StartIpAddress
	}
	s, err := ctx.ReadResource("azure:postgresql/firewallRule:FirewallRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FirewallRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FirewallRule) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FirewallRule) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
func (r *FirewallRule) EndIpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endIpAddress"])
}

// Specifies the name of the PostgreSQL Firewall Rule. Changing this forces a
// new resource to be created.
func (r *FirewallRule) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
func (r *FirewallRule) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
func (r *FirewallRule) ServerName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serverName"])
}

// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
func (r *FirewallRule) StartIpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["startIpAddress"])
}

// Input properties used for looking up and filtering FirewallRule resources.
type FirewallRuleState struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress interface{}
	// Specifies the name of the PostgreSQL Firewall Rule. Changing this forces a
	// new resource to be created.
	Name interface{}
	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName interface{}
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress interface{}
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// Specifies the End IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	EndIpAddress interface{}
	// Specifies the name of the PostgreSQL Firewall Rule. Changing this forces a
	// new resource to be created.
	Name interface{}
	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName interface{}
	// Specifies the Start IP Address associated with this Firewall Rule. Changing this forces a new resource to be created.
	StartIpAddress interface{}
}
