// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an HTTP Action within a Logic App Workflow
type ActionHttp struct {
	s *pulumi.ResourceState
}

// NewActionHttp registers a new resource with the given unique name, arguments, and options.
func NewActionHttp(ctx *pulumi.Context,
	name string, args *ActionHttpArgs, opts ...pulumi.ResourceOpt) (*ActionHttp, error) {
	if args == nil || args.LogicAppId == nil {
		return nil, errors.New("missing required argument 'LogicAppId'")
	}
	if args == nil || args.Method == nil {
		return nil, errors.New("missing required argument 'Method'")
	}
	if args == nil || args.Uri == nil {
		return nil, errors.New("missing required argument 'Uri'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["body"] = nil
		inputs["headers"] = nil
		inputs["logicAppId"] = nil
		inputs["method"] = nil
		inputs["name"] = nil
		inputs["uri"] = nil
	} else {
		inputs["body"] = args.Body
		inputs["headers"] = args.Headers
		inputs["logicAppId"] = args.LogicAppId
		inputs["method"] = args.Method
		inputs["name"] = args.Name
		inputs["uri"] = args.Uri
	}
	s, err := ctx.RegisterResource("azure:logicapps/actionHttp:ActionHttp", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ActionHttp{s: s}, nil
}

// GetActionHttp gets an existing ActionHttp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionHttp(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ActionHttpState, opts ...pulumi.ResourceOpt) (*ActionHttp, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["body"] = state.Body
		inputs["headers"] = state.Headers
		inputs["logicAppId"] = state.LogicAppId
		inputs["method"] = state.Method
		inputs["name"] = state.Name
		inputs["uri"] = state.Uri
	}
	s, err := ctx.ReadResource("azure:logicapps/actionHttp:ActionHttp", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ActionHttp{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ActionHttp) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ActionHttp) ID() *pulumi.IDOutput {
	return r.s.ID
}

// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
func (r *ActionHttp) Body() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["body"])
}

// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
func (r *ActionHttp) Headers() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["headers"])
}

// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
func (r *ActionHttp) LogicAppId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["logicAppId"])
}

// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
func (r *ActionHttp) Method() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["method"])
}

// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
func (r *ActionHttp) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Specifies the URI which will be called when this HTTP Action is triggered.
func (r *ActionHttp) Uri() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["uri"])
}

// Input properties used for looking up and filtering ActionHttp resources.
type ActionHttpState struct {
	// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
	Body interface{}
	// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
	Headers interface{}
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId interface{}
	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
	Method interface{}
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name interface{}
	// Specifies the URI which will be called when this HTTP Action is triggered.
	Uri interface{}
}

// The set of arguments for constructing a ActionHttp resource.
type ActionHttpArgs struct {
	// Specifies the HTTP Body that should be sent to the `uri` when this HTTP Action is triggered.
	Body interface{}
	// Specifies a Map of Key-Value Pairs that should be sent to the `uri` when this HTTP Action is triggered.
	Headers interface{}
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId interface{}
	// Specifies the HTTP Method which should be used for this HTTP Action. Possible values include `DELETE`, `GET`, `PATCH`, `POST` and `PUT`.
	Method interface{}
	// Specifies the name of the HTTP Action to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name interface{}
	// Specifies the URI which will be called when this HTTP Action is triggered.
	Uri interface{}
}