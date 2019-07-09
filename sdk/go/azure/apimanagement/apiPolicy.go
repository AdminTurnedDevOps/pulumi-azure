// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Management API Policy
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_api_policy.html.markdown.
type ApiPolicy struct {
	s *pulumi.ResourceState
}

// NewApiPolicy registers a new resource with the given unique name, arguments, and options.
func NewApiPolicy(ctx *pulumi.Context,
	name string, args *ApiPolicyArgs, opts ...pulumi.ResourceOpt) (*ApiPolicy, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.ApiName == nil {
		return nil, errors.New("missing required argument 'ApiName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["apiManagementName"] = nil
		inputs["apiName"] = nil
		inputs["resourceGroupName"] = nil
		inputs["xmlContent"] = nil
		inputs["xmlLink"] = nil
	} else {
		inputs["apiManagementName"] = args.ApiManagementName
		inputs["apiName"] = args.ApiName
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["xmlContent"] = args.XmlContent
		inputs["xmlLink"] = args.XmlLink
	}
	s, err := ctx.RegisterResource("azure:apimanagement/apiPolicy:ApiPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiPolicy{s: s}, nil
}

// GetApiPolicy gets an existing ApiPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApiPolicyState, opts ...pulumi.ResourceOpt) (*ApiPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiManagementName"] = state.ApiManagementName
		inputs["apiName"] = state.ApiName
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["xmlContent"] = state.XmlContent
		inputs["xmlLink"] = state.XmlLink
	}
	s, err := ctx.ReadResource("azure:apimanagement/apiPolicy:ApiPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ApiPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ApiPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ApiPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the API Management Service. Changing this forces a new resource to be created.
func (r *ApiPolicy) ApiManagementName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["apiManagementName"])
}

// The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
func (r *ApiPolicy) ApiName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["apiName"])
}

// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
func (r *ApiPolicy) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The XML Content for this Policy.
func (r *ApiPolicy) XmlContent() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["xmlContent"])
}

// A link to a Policy XML Document, which must be publicly available.
func (r *ApiPolicy) XmlLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["xmlLink"])
}

// Input properties used for looking up and filtering ApiPolicy resources.
type ApiPolicyState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName interface{}
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The XML Content for this Policy.
	XmlContent interface{}
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink interface{}
}

// The set of arguments for constructing a ApiPolicy resource.
type ApiPolicyArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName interface{}
	// The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	ApiName interface{}
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The XML Content for this Policy.
	XmlContent interface{}
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink interface{}
}
