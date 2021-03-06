// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logicapps

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Recurrence Trigger within a Logic App Workflow
type TriggerRecurrence struct {
	pulumi.CustomResourceState

	// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
	Frequency pulumi.StringOutput `pulumi:"frequency"`
	// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
	Interval pulumi.IntOutput `pulumi:"interval"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringOutput `pulumi:"logicAppId"`
	// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
}

// NewTriggerRecurrence registers a new resource with the given unique name, arguments, and options.
func NewTriggerRecurrence(ctx *pulumi.Context,
	name string, args *TriggerRecurrenceArgs, opts ...pulumi.ResourceOption) (*TriggerRecurrence, error) {
	if args == nil || args.Frequency == nil {
		return nil, errors.New("missing required argument 'Frequency'")
	}
	if args == nil || args.Interval == nil {
		return nil, errors.New("missing required argument 'Interval'")
	}
	if args == nil || args.LogicAppId == nil {
		return nil, errors.New("missing required argument 'LogicAppId'")
	}
	if args == nil {
		args = &TriggerRecurrenceArgs{}
	}
	var resource TriggerRecurrence
	err := ctx.RegisterResource("azure:logicapps/triggerRecurrence:TriggerRecurrence", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTriggerRecurrence gets an existing TriggerRecurrence resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTriggerRecurrence(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerRecurrenceState, opts ...pulumi.ResourceOption) (*TriggerRecurrence, error) {
	var resource TriggerRecurrence
	err := ctx.ReadResource("azure:logicapps/triggerRecurrence:TriggerRecurrence", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TriggerRecurrence resources.
type triggerRecurrenceState struct {
	// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
	Frequency *string `pulumi:"frequency"`
	// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
	Interval *int `pulumi:"interval"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId *string `pulumi:"logicAppId"`
	// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
	StartTime *string `pulumi:"startTime"`
}

type TriggerRecurrenceState struct {
	// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
	Frequency pulumi.StringPtrInput
	// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
	Interval pulumi.IntPtrInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringPtrInput
	// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
	StartTime pulumi.StringPtrInput
}

func (TriggerRecurrenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerRecurrenceState)(nil)).Elem()
}

type triggerRecurrenceArgs struct {
	// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
	Frequency string `pulumi:"frequency"`
	// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
	Interval int `pulumi:"interval"`
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId string `pulumi:"logicAppId"`
	// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
	StartTime *string `pulumi:"startTime"`
}

// The set of arguments for constructing a TriggerRecurrence resource.
type TriggerRecurrenceArgs struct {
	// Specifies the Frequency at which this Trigger should be run. Possible values include `Month`, `Week`, `Day`, `Hour`, `Minute` and `Second`.
	Frequency pulumi.StringInput
	// Specifies interval used for the Frequency, for example a value of `4` for `interval` and `hour` for `frequency` would run the Trigger every 4 hours.
	Interval pulumi.IntInput
	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppId pulumi.StringInput
	// Specifies the name of the Recurrence Triggers to be created within the Logic App Workflow. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the start date and time for this trigger in RFC3339 format: `2000-01-02T03:04:05Z`.
	StartTime pulumi.StringPtrInput
}

func (TriggerRecurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerRecurrenceArgs)(nil)).Elem()
}
