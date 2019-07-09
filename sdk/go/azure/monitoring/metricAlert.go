// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Metric Alert within Azure Monitor.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/monitor_metric_alert.html.markdown.
type MetricAlert struct {
	s *pulumi.ResourceState
}

// NewMetricAlert registers a new resource with the given unique name, arguments, and options.
func NewMetricAlert(ctx *pulumi.Context,
	name string, args *MetricAlertArgs, opts ...pulumi.ResourceOpt) (*MetricAlert, error) {
	if args == nil || args.Criterias == nil {
		return nil, errors.New("missing required argument 'Criterias'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Scopes == nil {
		return nil, errors.New("missing required argument 'Scopes'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["actions"] = nil
		inputs["autoMitigate"] = nil
		inputs["criterias"] = nil
		inputs["description"] = nil
		inputs["enabled"] = nil
		inputs["frequency"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["scopes"] = nil
		inputs["severity"] = nil
		inputs["tags"] = nil
		inputs["windowSize"] = nil
	} else {
		inputs["actions"] = args.Actions
		inputs["autoMitigate"] = args.AutoMitigate
		inputs["criterias"] = args.Criterias
		inputs["description"] = args.Description
		inputs["enabled"] = args.Enabled
		inputs["frequency"] = args.Frequency
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["scopes"] = args.Scopes
		inputs["severity"] = args.Severity
		inputs["tags"] = args.Tags
		inputs["windowSize"] = args.WindowSize
	}
	s, err := ctx.RegisterResource("azure:monitoring/metricAlert:MetricAlert", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MetricAlert{s: s}, nil
}

// GetMetricAlert gets an existing MetricAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricAlert(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MetricAlertState, opts ...pulumi.ResourceOpt) (*MetricAlert, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["actions"] = state.Actions
		inputs["autoMitigate"] = state.AutoMitigate
		inputs["criterias"] = state.Criterias
		inputs["description"] = state.Description
		inputs["enabled"] = state.Enabled
		inputs["frequency"] = state.Frequency
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["scopes"] = state.Scopes
		inputs["severity"] = state.Severity
		inputs["tags"] = state.Tags
		inputs["windowSize"] = state.WindowSize
	}
	s, err := ctx.ReadResource("azure:monitoring/metricAlert:MetricAlert", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MetricAlert{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *MetricAlert) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *MetricAlert) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// One or more `action` blocks as defined below.
func (r *MetricAlert) Actions() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["actions"])
}

// Should the alerts in this Metric Alert be auto resolved? Defaults to `false`.
func (r *MetricAlert) AutoMitigate() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoMitigate"])
}

// One or more `criteria` blocks as defined below.
func (r *MetricAlert) Criterias() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["criterias"])
}

// The description of this Metric Alert.
func (r *MetricAlert) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Should this Metric Alert be enabled? Defaults to `true`.
func (r *MetricAlert) Enabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enabled"])
}

// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
func (r *MetricAlert) Frequency() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["frequency"])
}

// The name of the Metric Alert. Changing this forces a new resource to be created.
func (r *MetricAlert) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which to create the Metric Alert instance.
func (r *MetricAlert) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The resource ID at which the metric criteria should be applied.
func (r *MetricAlert) Scopes() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["scopes"])
}

// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
func (r *MetricAlert) Severity() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["severity"])
}

// A mapping of tags to assign to the resource.
func (r *MetricAlert) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
func (r *MetricAlert) WindowSize() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["windowSize"])
}

// Input properties used for looking up and filtering MetricAlert resources.
type MetricAlertState struct {
	// One or more `action` blocks as defined below.
	Actions interface{}
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `false`.
	AutoMitigate interface{}
	// One or more `criteria` blocks as defined below.
	Criterias interface{}
	// The description of this Metric Alert.
	Description interface{}
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled interface{}
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency interface{}
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName interface{}
	// The resource ID at which the metric criteria should be applied.
	Scopes interface{}
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize interface{}
}

// The set of arguments for constructing a MetricAlert resource.
type MetricAlertArgs struct {
	// One or more `action` blocks as defined below.
	Actions interface{}
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `false`.
	AutoMitigate interface{}
	// One or more `criteria` blocks as defined below.
	Criterias interface{}
	// The description of this Metric Alert.
	Description interface{}
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled interface{}
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency interface{}
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName interface{}
	// The resource ID at which the metric criteria should be applied.
	Scopes interface{}
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize interface{}
}
