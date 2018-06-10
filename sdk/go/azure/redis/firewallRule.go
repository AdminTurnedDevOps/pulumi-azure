// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Firewall Rule associated with a Premium Redis Cache.
// 
// ~> **Note:** Redis Firewall Rules can only be assigned to a Redis Cache with a `Premium` SKU.
type FirewallRule struct {
	s *pulumi.ResourceState
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOpt) (*FirewallRule, error) {
	if args == nil || args.EndIp == nil {
		return nil, errors.New("missing required argument 'EndIp'")
	}
	if args == nil || args.RedisCacheName == nil {
		return nil, errors.New("missing required argument 'RedisCacheName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StartIp == nil {
		return nil, errors.New("missing required argument 'StartIp'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["endIp"] = nil
		inputs["name"] = nil
		inputs["redisCacheName"] = nil
		inputs["resourceGroupName"] = nil
		inputs["startIp"] = nil
	} else {
		inputs["endIp"] = args.EndIp
		inputs["name"] = args.Name
		inputs["redisCacheName"] = args.RedisCacheName
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["startIp"] = args.StartIp
	}
	s, err := ctx.RegisterResource("azure:redis/firewallRule:FirewallRule", name, true, inputs, opts...)
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
		inputs["endIp"] = state.EndIp
		inputs["name"] = state.Name
		inputs["redisCacheName"] = state.RedisCacheName
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["startIp"] = state.StartIp
	}
	s, err := ctx.ReadResource("azure:redis/firewallRule:FirewallRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FirewallRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FirewallRule) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FirewallRule) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The highest IP address included in the range.
func (r *FirewallRule) EndIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endIp"])
}

// The name of the Firewall Rule. Changing this forces a new resource to be created.
func (r *FirewallRule) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the Redis Cache. Changing this forces a new resource to be created.
func (r *FirewallRule) RedisCacheName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["redisCacheName"])
}

// The name of the resource group in which this Redis Cache exists.
func (r *FirewallRule) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The lowest IP address included in the range
func (r *FirewallRule) StartIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["startIp"])
}

// Input properties used for looking up and filtering FirewallRule resources.
type FirewallRuleState struct {
	// The highest IP address included in the range.
	EndIp interface{}
	// The name of the Firewall Rule. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Redis Cache. Changing this forces a new resource to be created.
	RedisCacheName interface{}
	// The name of the resource group in which this Redis Cache exists.
	ResourceGroupName interface{}
	// The lowest IP address included in the range
	StartIp interface{}
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// The highest IP address included in the range.
	EndIp interface{}
	// The name of the Firewall Rule. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Redis Cache. Changing this forces a new resource to be created.
	RedisCacheName interface{}
	// The name of the resource group in which this Redis Cache exists.
	ResourceGroupName interface{}
	// The lowest IP address included in the range
	StartIp interface{}
}
