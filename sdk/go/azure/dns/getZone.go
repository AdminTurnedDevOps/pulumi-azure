// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to obtain information about a DNS Zone.
func Lookupone(ctx *pulumi.Context, args *GetZoneArgs) (*GetZoneResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	outputs, err := ctx.Invoke("azure:dns/getZone:getZone", inputs)
	if err != nil {
		return nil, err
	}
	ret := GetZoneResult{}
	if v, ok := outputs["maxNumberOfRecordSets"]; ok {
		ret.MaxNumberOfRecordSets = v
	}
	if v, ok := outputs["nameServers"]; ok {
		ret.NameServers = v
	}
	if v, ok := outputs["numberOfRecordSets"]; ok {
		ret.NumberOfRecordSets = v
	}
	if v, ok := outputs["resourceGroupName"]; ok {
		ret.ResourceGroupName = v
	}
	if v, ok := outputs["tags"]; ok {
		ret.Tags = v
	}
	return &ret, nil
}

// A collection of arguments for invoking getZone.
type GetZoneArgs struct {
	// The name of the DNS Zone.
	Name interface{}
	// The Name of the Resource Group where the DNS Zone exists.
	// If the Name of the Resource Group is not provided, the first DNS Zone from the list of DNS Zones
	// in your subscription that matches `name` will be returned.
	ResourceGroupName interface{}
}

// A collection of values returned by getZone.
type GetZoneResult struct {
	// Maximum number of Records in the zone.
	MaxNumberOfRecordSets interface{}
	// A list of values that make up the NS record for the zone.
	NameServers interface{}
	// The number of records already in the zone.
	NumberOfRecordSets interface{}
	ResourceGroupName interface{}
	// A mapping of tags to assign to the EventHub Namespace.
	Tags interface{}
}
