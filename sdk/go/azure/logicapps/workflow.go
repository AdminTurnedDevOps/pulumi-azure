// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Logic App Workflow.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/logic_app_workflow.html.markdown.
type Workflow struct {
	s *pulumi.ResourceState
}

// NewWorkflow registers a new resource with the given unique name, arguments, and options.
func NewWorkflow(ctx *pulumi.Context,
	name string, args *WorkflowArgs, opts ...pulumi.ResourceOpt) (*Workflow, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["parameters"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
		inputs["workflowSchema"] = nil
		inputs["workflowVersion"] = nil
	} else {
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["parameters"] = args.Parameters
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
		inputs["workflowSchema"] = args.WorkflowSchema
		inputs["workflowVersion"] = args.WorkflowVersion
	}
	inputs["accessEndpoint"] = nil
	s, err := ctx.RegisterResource("azure:logicapps/workflow:Workflow", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Workflow{s: s}, nil
}

// GetWorkflow gets an existing Workflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.ID, state *WorkflowState, opts ...pulumi.ResourceOpt) (*Workflow, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessEndpoint"] = state.AccessEndpoint
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["parameters"] = state.Parameters
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
		inputs["workflowSchema"] = state.WorkflowSchema
		inputs["workflowVersion"] = state.WorkflowVersion
	}
	s, err := ctx.ReadResource("azure:logicapps/workflow:Workflow", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Workflow{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Workflow) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Workflow) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The Access Endpoint for the Logic App Workflow
func (r *Workflow) AccessEndpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessEndpoint"])
}

// Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
func (r *Workflow) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
func (r *Workflow) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A map of Key-Value pairs.
func (r *Workflow) Parameters() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["parameters"])
}

// The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
func (r *Workflow) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *Workflow) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
func (r *Workflow) WorkflowSchema() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["workflowSchema"])
}

// Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be create.d
func (r *Workflow) WorkflowVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["workflowVersion"])
}

// Input properties used for looking up and filtering Workflow resources.
type WorkflowState struct {
	// The Access Endpoint for the Logic App Workflow
	AccessEndpoint interface{}
	// Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
	Name interface{}
	// A map of Key-Value pairs.
	Parameters interface{}
	// The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
	WorkflowSchema interface{}
	// Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be create.d
	WorkflowVersion interface{}
}

// The set of arguments for constructing a Workflow resource.
type WorkflowArgs struct {
	// Specifies the supported Azure location where the Logic App Workflow exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Logic App Workflow. Changing this forces a new resource to be created.
	Name interface{}
	// A map of Key-Value pairs.
	Parameters interface{}
	// The name of the Resource Group in which the Logic App Workflow should be created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the Schema to use for this Logic App Workflow. Defaults to `https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#`. Changing this forces a new resource to be created.
	WorkflowSchema interface{}
	// Specifies the version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`. Changing this forces a new resource to be create.d
	WorkflowVersion interface{}
}
