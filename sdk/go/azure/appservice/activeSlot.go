// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Promotes an App Service Slot to Production within an App Service.
// 
// -> **Note:** When using Slots - the `app_settings`, `connection_string` and `site_config` blocks on the `azurerm_app_service` resource will be overwritten when promoting a Slot using the `azurerm_app_service_active_slot` resource.
type ActiveSlot struct {
	s *pulumi.ResourceState
}

// NewActiveSlot registers a new resource with the given unique name, arguments, and options.
func NewActiveSlot(ctx *pulumi.Context,
	name string, args *ActiveSlotArgs, opts ...pulumi.ResourceOpt) (*ActiveSlot, error) {
	if args == nil || args.AppServiceName == nil {
		return nil, errors.New("missing required argument 'AppServiceName'")
	}
	if args == nil || args.AppServiceSlotName == nil {
		return nil, errors.New("missing required argument 'AppServiceSlotName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appServiceName"] = nil
		inputs["appServiceSlotName"] = nil
		inputs["resourceGroupName"] = nil
	} else {
		inputs["appServiceName"] = args.AppServiceName
		inputs["appServiceSlotName"] = args.AppServiceSlotName
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	s, err := ctx.RegisterResource("azure:appservice/activeSlot:ActiveSlot", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ActiveSlot{s: s}, nil
}

// GetActiveSlot gets an existing ActiveSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActiveSlot(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ActiveSlotState, opts ...pulumi.ResourceOpt) (*ActiveSlot, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appServiceName"] = state.AppServiceName
		inputs["appServiceSlotName"] = state.AppServiceSlotName
		inputs["resourceGroupName"] = state.ResourceGroupName
	}
	s, err := ctx.ReadResource("azure:appservice/activeSlot:ActiveSlot", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ActiveSlot{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ActiveSlot) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ActiveSlot) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
func (r *ActiveSlot) AppServiceName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["appServiceName"])
}

// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
func (r *ActiveSlot) AppServiceSlotName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["appServiceSlotName"])
}

// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
func (r *ActiveSlot) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Input properties used for looking up and filtering ActiveSlot resources.
type ActiveSlotState struct {
	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName interface{}
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName interface{}
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
}

// The set of arguments for constructing a ActiveSlot resource.
type ActiveSlotArgs struct {
	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName interface{}
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName interface{}
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
}