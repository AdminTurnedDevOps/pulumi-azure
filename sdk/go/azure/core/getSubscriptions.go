// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about all the Subscriptions currently available.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/subscriptions.html.markdown.
func LookupSubscriptions(ctx *pulumi.Context, args *GetSubscriptionsArgs) (*GetSubscriptionsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["displayNameContains"] = args.DisplayNameContains
		inputs["displayNamePrefix"] = args.DisplayNamePrefix
	}
	outputs, err := ctx.Invoke("azure:core/getSubscriptions:getSubscriptions", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSubscriptionsResult{
		DisplayNameContains: outputs["displayNameContains"],
		DisplayNamePrefix: outputs["displayNamePrefix"],
		Subscriptions: outputs["subscriptions"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSubscriptions.
type GetSubscriptionsArgs struct {
	// A case-insensitive value which must be contained within the `display_name` field, used to filter the results
	DisplayNameContains interface{}
	// A case-insensitive prefix which can be used to filter on the `display_name` field
	DisplayNamePrefix interface{}
}

// A collection of values returned by getSubscriptions.
type GetSubscriptionsResult struct {
	DisplayNameContains interface{}
	DisplayNamePrefix interface{}
	// One or more `subscription` blocks as defined below.
	Subscriptions interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
