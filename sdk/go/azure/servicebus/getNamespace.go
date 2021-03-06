// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing ServiceBus Namespace.
func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure:servicebus/getNamespace:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNamespace.
type LookupNamespaceArgs struct {
	// Specifies the name of the ServiceBus Namespace.
	Name string `pulumi:"name"`
	// Specifies the name of the Resource Group where the ServiceBus Namespace exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getNamespace.
type LookupNamespaceResult struct {
	// The capacity of the ServiceBus Namespace.
	Capacity int `pulumi:"capacity"`
	// The primary connection string for the authorization
	// rule `RootManageSharedAccessKey`.
	DefaultPrimaryConnectionString string `pulumi:"defaultPrimaryConnectionString"`
	// The primary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultPrimaryKey string `pulumi:"defaultPrimaryKey"`
	// The secondary connection string for the
	// authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryConnectionString string `pulumi:"defaultSecondaryConnectionString"`
	// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
	DefaultSecondaryKey string `pulumi:"defaultSecondaryKey"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The location of the Resource Group in which the ServiceBus Namespace exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Tier used for the ServiceBus Namespace.
	Sku string `pulumi:"sku"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this ServiceBus Namespace is zone redundant.
	ZoneRedundant bool `pulumi:"zoneRedundant"`
}
