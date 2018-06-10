// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access the properties of an Azure subscription.
func Lookupubscription(ctx *pulumi.Context, args *GetSubscriptionArgs) (*GetSubscriptionResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["subscriptionId"] = args.SubscriptionId
	}
	outputs, err := ctx.Invoke("azure:core/getSubscription:getSubscription", inputs)
	if err != nil {
		return nil, err
	}
	ret := GetSubscriptionResult{}
	if v, ok := outputs["displayName"]; ok {
		ret.DisplayName = v
	}
	if v, ok := outputs["locationPlacementId"]; ok {
		ret.LocationPlacementId = v
	}
	if v, ok := outputs["quotaId"]; ok {
		ret.QuotaId = v
	}
	if v, ok := outputs["spendingLimit"]; ok {
		ret.SpendingLimit = v
	}
	if v, ok := outputs["state"]; ok {
		ret.State = v
	}
	if v, ok := outputs["subscriptionId"]; ok {
		ret.SubscriptionId = v
	}
	return &ret, nil
}

// A collection of arguments for invoking getSubscription.
type GetSubscriptionArgs struct {
	// Specifies the ID of the subscription. If this argument is omitted, the subscription ID of the current Azure Resource Manager provider is used.
	SubscriptionId interface{}
}

// A collection of values returned by getSubscription.
type GetSubscriptionResult struct {
	// The subscription display name.
	DisplayName interface{}
	// The subscription location placement ID.
	LocationPlacementId interface{}
	// The subscription quota ID.
	QuotaId interface{}
	// The subscription spending limit.
	SpendingLimit interface{}
	// The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
	State interface{}
	SubscriptionId interface{}
}
