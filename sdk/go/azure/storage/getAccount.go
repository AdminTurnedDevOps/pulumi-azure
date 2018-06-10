// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Gets information about the specified Storage Account.
func Lookupccount(ctx *pulumi.Context, args *GetAccountArgs) (*GetAccountResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	outputs, err := ctx.Invoke("azure:storage/getAccount:getAccount", inputs)
	if err != nil {
		return nil, err
	}
	ret := GetAccountResult{}
	if v, ok := outputs["accessTier"]; ok {
		ret.AccessTier = v
	}
	if v, ok := outputs["accountEncryptionSource"]; ok {
		ret.AccountEncryptionSource = v
	}
	if v, ok := outputs["accountKind"]; ok {
		ret.AccountKind = v
	}
	if v, ok := outputs["accountReplicationType"]; ok {
		ret.AccountReplicationType = v
	}
	if v, ok := outputs["accountTier"]; ok {
		ret.AccountTier = v
	}
	if v, ok := outputs["customDomain"]; ok {
		ret.CustomDomain = v
	}
	if v, ok := outputs["enableBlobEncryption"]; ok {
		ret.EnableBlobEncryption = v
	}
	if v, ok := outputs["enableFileEncryption"]; ok {
		ret.EnableFileEncryption = v
	}
	if v, ok := outputs["enableHttpsTrafficOnly"]; ok {
		ret.EnableHttpsTrafficOnly = v
	}
	if v, ok := outputs["location"]; ok {
		ret.Location = v
	}
	if v, ok := outputs["primaryAccessKey"]; ok {
		ret.PrimaryAccessKey = v
	}
	if v, ok := outputs["primaryBlobConnectionString"]; ok {
		ret.PrimaryBlobConnectionString = v
	}
	if v, ok := outputs["primaryBlobEndpoint"]; ok {
		ret.PrimaryBlobEndpoint = v
	}
	if v, ok := outputs["primaryConnectionString"]; ok {
		ret.PrimaryConnectionString = v
	}
	if v, ok := outputs["primaryFileEndpoint"]; ok {
		ret.PrimaryFileEndpoint = v
	}
	if v, ok := outputs["primaryLocation"]; ok {
		ret.PrimaryLocation = v
	}
	if v, ok := outputs["primaryQueueEndpoint"]; ok {
		ret.PrimaryQueueEndpoint = v
	}
	if v, ok := outputs["primaryTableEndpoint"]; ok {
		ret.PrimaryTableEndpoint = v
	}
	if v, ok := outputs["secondaryAccessKey"]; ok {
		ret.SecondaryAccessKey = v
	}
	if v, ok := outputs["secondaryBlobConnectionString"]; ok {
		ret.SecondaryBlobConnectionString = v
	}
	if v, ok := outputs["secondaryBlobEndpoint"]; ok {
		ret.SecondaryBlobEndpoint = v
	}
	if v, ok := outputs["secondaryConnectionString"]; ok {
		ret.SecondaryConnectionString = v
	}
	if v, ok := outputs["secondaryLocation"]; ok {
		ret.SecondaryLocation = v
	}
	if v, ok := outputs["secondaryQueueEndpoint"]; ok {
		ret.SecondaryQueueEndpoint = v
	}
	if v, ok := outputs["secondaryTableEndpoint"]; ok {
		ret.SecondaryTableEndpoint = v
	}
	if v, ok := outputs["tags"]; ok {
		ret.Tags = v
	}
	return &ret, nil
}

// A collection of arguments for invoking getAccount.
type GetAccountArgs struct {
	// Specifies the name of the Storage Account
	Name interface{}
	// Specifies the name of the resource group the Storage Account is located in.
	ResourceGroupName interface{}
}

// A collection of values returned by getAccount.
type GetAccountResult struct {
	// Defines the access tier for `BlobStorage` accounts.
	AccessTier interface{}
	// The Encryption Source for this Storage Account.
	AccountEncryptionSource interface{}
	// Defines the Kind of account, either `BlobStorage` or `Storage`.
	AccountKind interface{}
	// Defines the type of replication used for this storage account.
	AccountReplicationType interface{}
	// Defines the Tier of this storage account.
	AccountTier interface{}
	// A `custom_domain` block as documented below.
	CustomDomain interface{}
	// Are Encryption Services are enabled for Blob storage? See [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/)
	// for more information.
	EnableBlobEncryption interface{}
	// Are Encryption Services are enabled for File storage? See [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/)
	// for more information.
	EnableFileEncryption interface{}
	// Is traffic only allowed via HTTPS? See [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information.
	EnableHttpsTrafficOnly interface{}
	// The Azure location where the Storage Account exists
	Location interface{}
	// The primary access key for the Storage Account.
	PrimaryAccessKey interface{}
	// The connection string associated with the primary blob location
	PrimaryBlobConnectionString interface{}
	// The endpoint URL for blob storage in the primary location.
	PrimaryBlobEndpoint interface{}
	// The connection string associated with the primary location
	PrimaryConnectionString interface{}
	// The endpoint URL for file storage in the primary location.
	PrimaryFileEndpoint interface{}
	// The primary location of the Storage Account.
	PrimaryLocation interface{}
	// The endpoint URL for queue storage in the primary location.
	PrimaryQueueEndpoint interface{}
	// The endpoint URL for table storage in the primary location.
	PrimaryTableEndpoint interface{}
	// The secondary access key for the Storage Account.
	SecondaryAccessKey interface{}
	// The connection string associated with the secondary blob location
	SecondaryBlobConnectionString interface{}
	// The endpoint URL for blob storage in the secondary location.
	SecondaryBlobEndpoint interface{}
	// The connection string associated with the secondary location
	SecondaryConnectionString interface{}
	// The secondary location of the Storage Account.
	SecondaryLocation interface{}
	// The endpoint URL for queue storage in the secondary location.
	SecondaryQueueEndpoint interface{}
	// The endpoint URL for table storage in the secondary location.
	SecondaryTableEndpoint interface{}
	// A mapping of tags to assigned to the resource.
	Tags interface{}
}
