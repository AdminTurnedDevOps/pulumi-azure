// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Disk Snapshot.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/snapshot.html.markdown.
type Snapshot struct {
	s *pulumi.ResourceState
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOpt) (*Snapshot, error) {
	if args == nil || args.CreateOption == nil {
		return nil, errors.New("missing required argument 'CreateOption'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["createOption"] = nil
		inputs["diskSizeGb"] = nil
		inputs["encryptionSettings"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["sourceResourceId"] = nil
		inputs["sourceUri"] = nil
		inputs["storageAccountId"] = nil
		inputs["tags"] = nil
	} else {
		inputs["createOption"] = args.CreateOption
		inputs["diskSizeGb"] = args.DiskSizeGb
		inputs["encryptionSettings"] = args.EncryptionSettings
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["sourceResourceId"] = args.SourceResourceId
		inputs["sourceUri"] = args.SourceUri
		inputs["storageAccountId"] = args.StorageAccountId
		inputs["tags"] = args.Tags
	}
	s, err := ctx.RegisterResource("azure:compute/snapshot:Snapshot", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Snapshot{s: s}, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SnapshotState, opts ...pulumi.ResourceOpt) (*Snapshot, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["createOption"] = state.CreateOption
		inputs["diskSizeGb"] = state.DiskSizeGb
		inputs["encryptionSettings"] = state.EncryptionSettings
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["sourceResourceId"] = state.SourceResourceId
		inputs["sourceUri"] = state.SourceUri
		inputs["storageAccountId"] = state.StorageAccountId
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:compute/snapshot:Snapshot", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Snapshot{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Snapshot) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Snapshot) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
func (r *Snapshot) CreateOption() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createOption"])
}

// The size of the Snapshotted Disk in GB.
func (r *Snapshot) DiskSizeGb() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["diskSizeGb"])
}

func (r *Snapshot) EncryptionSettings() *pulumi.Output {
	return r.s.State["encryptionSettings"]
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *Snapshot) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
func (r *Snapshot) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
func (r *Snapshot) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
func (r *Snapshot) SourceResourceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceResourceId"])
}

// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
func (r *Snapshot) SourceUri() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceUri"])
}

// Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
func (r *Snapshot) StorageAccountId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["storageAccountId"])
}

// A mapping of tags to assign to the resource.
func (r *Snapshot) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Snapshot resources.
type SnapshotState struct {
	// Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
	CreateOption interface{}
	// The size of the Snapshotted Disk in GB.
	DiskSizeGb interface{}
	EncryptionSettings interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
	SourceResourceId interface{}
	// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
	SourceUri interface{}
	// Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
	StorageAccountId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// Indicates how the snapshot is to be created. Possible values are `Copy` or `Import`. Changing this forces a new resource to be created.
	CreateOption interface{}
	// The size of the Snapshotted Disk in GB.
	DiskSizeGb interface{}
	EncryptionSettings interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Snapshot resource. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which to create the Snapshot. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Specifies a reference to an existing snapshot, when `create_option` is `Copy`. Changing this forces a new resource to be created.
	SourceResourceId interface{}
	// Specifies the URI to a Managed or Unmanaged Disk. Changing this forces a new resource to be created.
	SourceUri interface{}
	// Specifies the ID of an storage account. Used with `source_uri` to allow authorization during import of unmanaged blobs from a different subscription. Changing this forces a new resource to be created.
	StorageAccountId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
