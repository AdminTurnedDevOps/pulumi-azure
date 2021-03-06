// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Event Hubs Authorization Rule within an Event Hub.
func LookupAuthorizationRule(ctx *pulumi.Context, args *LookupAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationRuleResult, error) {
	var rv LookupAuthorizationRuleResult
	err := ctx.Invoke("azure:eventhub/getAuthorizationRule:getAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthorizationRule.
type LookupAuthorizationRuleArgs struct {
	// Specifies the name of the EventHub.
	EventhubName string `pulumi:"eventhubName"`
	Listen       *bool  `pulumi:"listen"`
	Manage       *bool  `pulumi:"manage"`
	// Specifies the name of the EventHub Authorization Rule resource. be created.
	Name string `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group in which the EventHub Authorization Rule's grandparent Namespace exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Send              *bool  `pulumi:"send"`
}

// A collection of values returned by getAuthorizationRule.
type LookupAuthorizationRuleResult struct {
	EventhubName string `pulumi:"eventhubName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	Listen        *bool  `pulumi:"listen"`
	Location      string `pulumi:"location"`
	Manage        *bool  `pulumi:"manage"`
	Name          string `pulumi:"name"`
	NamespaceName string `pulumi:"namespaceName"`
	// The Primary Connection String for the Event Hubs Authorization Rule.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// The Primary Key for the Event Hubs Authorization Rule.
	PrimaryKey        string `pulumi:"primaryKey"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Event Hubs Authorization Rule.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the Event Hubs Authorization Rule.
	SecondaryKey string `pulumi:"secondaryKey"`
	Send         *bool  `pulumi:"send"`
}
