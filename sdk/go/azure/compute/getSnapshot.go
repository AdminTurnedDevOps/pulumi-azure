// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Snapshot.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/snapshot.html.markdown.
func LookupSnapshot(ctx *pulumi.Context, args *GetSnapshotArgs) (*GetSnapshotResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	outputs, err := ctx.Invoke("azure:compute/getSnapshot:getSnapshot", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSnapshotResult{
		CreationOption: outputs["creationOption"],
		DiskSizeGb: outputs["diskSizeGb"],
		EncryptionSettings: outputs["encryptionSettings"],
		Name: outputs["name"],
		OsType: outputs["osType"],
		ResourceGroupName: outputs["resourceGroupName"],
		SourceResourceId: outputs["sourceResourceId"],
		SourceUri: outputs["sourceUri"],
		StorageAccountId: outputs["storageAccountId"],
		TimeCreated: outputs["timeCreated"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSnapshot.
type GetSnapshotArgs struct {
	// Specifies the name of the Snapshot.
	Name interface{}
	// Specifies the name of the resource group the Snapshot is located in.
	ResourceGroupName interface{}
}

// A collection of values returned by getSnapshot.
type GetSnapshotResult struct {
	CreationOption interface{}
	// The size of the Snapshotted Disk in GB.
	DiskSizeGb interface{}
	EncryptionSettings interface{}
	Name interface{}
	OsType interface{}
	ResourceGroupName interface{}
	// The reference to an existing snapshot.
	SourceResourceId interface{}
	// The URI to a Managed or Unmanaged Disk.
	SourceUri interface{}
	// The ID of an storage account.
	StorageAccountId interface{}
	TimeCreated interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
