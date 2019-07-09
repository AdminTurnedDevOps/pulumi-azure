// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Public IP Address.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/public_ip.html.markdown.
func LookupPublicIP(ctx *pulumi.Context, args *GetPublicIPArgs) (*GetPublicIPResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
		inputs["zones"] = args.Zones
	}
	outputs, err := ctx.Invoke("azure:network/getPublicIP:getPublicIP", inputs)
	if err != nil {
		return nil, err
	}
	return &GetPublicIPResult{
		AllocationMethod: outputs["allocationMethod"],
		DomainNameLabel: outputs["domainNameLabel"],
		Fqdn: outputs["fqdn"],
		IdleTimeoutInMinutes: outputs["idleTimeoutInMinutes"],
		IpAddress: outputs["ipAddress"],
		IpVersion: outputs["ipVersion"],
		Location: outputs["location"],
		Name: outputs["name"],
		ResourceGroupName: outputs["resourceGroupName"],
		ReverseFqdn: outputs["reverseFqdn"],
		Sku: outputs["sku"],
		Tags: outputs["tags"],
		Zones: outputs["zones"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getPublicIP.
type GetPublicIPArgs struct {
	// Specifies the name of the public IP address.
	Name interface{}
	// Specifies the name of the resource group.
	ResourceGroupName interface{}
	Tags interface{}
	Zones interface{}
}

// A collection of values returned by getPublicIP.
type GetPublicIPResult struct {
	AllocationMethod interface{}
	// The label for the Domain Name.
	DomainNameLabel interface{}
	// Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
	Fqdn interface{}
	// Specifies the timeout for the TCP idle connection.
	IdleTimeoutInMinutes interface{}
	// The IP address value that was allocated.
	IpAddress interface{}
	// The IP version being used, for example `IPv4` or `IPv6`.
	IpVersion interface{}
	Location interface{}
	Name interface{}
	ResourceGroupName interface{}
	ReverseFqdn interface{}
	Sku interface{}
	// A mapping of tags to assigned to the resource.
	Tags interface{}
	Zones interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
