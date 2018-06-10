// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Enables you to manage DNS CNAME Records within Azure DNS.
type CNameRecord struct {
	s *pulumi.ResourceState
}

// NewCNameRecord registers a new resource with the given unique name, arguments, and options.
func NewCNameRecord(ctx *pulumi.Context,
	name string, args *CNameRecordArgs, opts ...pulumi.ResourceOpt) (*CNameRecord, error) {
	if args == nil || args.Record == nil {
		return nil, errors.New("missing required argument 'Record'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Ttl == nil {
		return nil, errors.New("missing required argument 'Ttl'")
	}
	if args == nil || args.ZoneName == nil {
		return nil, errors.New("missing required argument 'ZoneName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["record"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
		inputs["ttl"] = nil
		inputs["zoneName"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["record"] = args.Record
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
		inputs["ttl"] = args.Ttl
		inputs["zoneName"] = args.ZoneName
	}
	s, err := ctx.RegisterResource("azure:dns/cNameRecord:CNameRecord", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CNameRecord{s: s}, nil
}

// GetCNameRecord gets an existing CNameRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCNameRecord(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CNameRecordState, opts ...pulumi.ResourceOpt) (*CNameRecord, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["record"] = state.Record
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
		inputs["ttl"] = state.Ttl
		inputs["zoneName"] = state.ZoneName
	}
	s, err := ctx.ReadResource("azure:dns/cNameRecord:CNameRecord", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CNameRecord{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *CNameRecord) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *CNameRecord) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The name of the DNS CNAME Record.
func (r *CNameRecord) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The target of the CNAME.
func (r *CNameRecord) Record() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["record"])
}

// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
func (r *CNameRecord) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *CNameRecord) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

func (r *CNameRecord) Ttl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ttl"])
}

// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
func (r *CNameRecord) ZoneName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zoneName"])
}

// Input properties used for looking up and filtering CNameRecord resources.
type CNameRecordState struct {
	// The name of the DNS CNAME Record.
	Name interface{}
	// The target of the CNAME.
	Record interface{}
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	Ttl interface{}
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName interface{}
}

// The set of arguments for constructing a CNameRecord resource.
type CNameRecordArgs struct {
	// The name of the DNS CNAME Record.
	Name interface{}
	// The target of the CNAME.
	Record interface{}
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	Ttl interface{}
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName interface{}
}
