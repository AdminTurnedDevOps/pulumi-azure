// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package streamanalytics

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Stream Analytics Stream Input IoTHub.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_stream_input_iothub.html.markdown.
type StreamInputIotHub struct {
	s *pulumi.ResourceState
}

// NewStreamInputIotHub registers a new resource with the given unique name, arguments, and options.
func NewStreamInputIotHub(ctx *pulumi.Context,
	name string, args *StreamInputIotHubArgs, opts ...pulumi.ResourceOpt) (*StreamInputIotHub, error) {
	if args == nil || args.Endpoint == nil {
		return nil, errors.New("missing required argument 'Endpoint'")
	}
	if args == nil || args.EventhubConsumerGroupName == nil {
		return nil, errors.New("missing required argument 'EventhubConsumerGroupName'")
	}
	if args == nil || args.IothubNamespace == nil {
		return nil, errors.New("missing required argument 'IothubNamespace'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Serialization == nil {
		return nil, errors.New("missing required argument 'Serialization'")
	}
	if args == nil || args.SharedAccessPolicyKey == nil {
		return nil, errors.New("missing required argument 'SharedAccessPolicyKey'")
	}
	if args == nil || args.SharedAccessPolicyName == nil {
		return nil, errors.New("missing required argument 'SharedAccessPolicyName'")
	}
	if args == nil || args.StreamAnalyticsJobName == nil {
		return nil, errors.New("missing required argument 'StreamAnalyticsJobName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["endpoint"] = nil
		inputs["eventhubConsumerGroupName"] = nil
		inputs["iothubNamespace"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["serialization"] = nil
		inputs["sharedAccessPolicyKey"] = nil
		inputs["sharedAccessPolicyName"] = nil
		inputs["streamAnalyticsJobName"] = nil
	} else {
		inputs["endpoint"] = args.Endpoint
		inputs["eventhubConsumerGroupName"] = args.EventhubConsumerGroupName
		inputs["iothubNamespace"] = args.IothubNamespace
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["serialization"] = args.Serialization
		inputs["sharedAccessPolicyKey"] = args.SharedAccessPolicyKey
		inputs["sharedAccessPolicyName"] = args.SharedAccessPolicyName
		inputs["streamAnalyticsJobName"] = args.StreamAnalyticsJobName
	}
	s, err := ctx.RegisterResource("azure:streamanalytics/streamInputIotHub:StreamInputIotHub", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &StreamInputIotHub{s: s}, nil
}

// GetStreamInputIotHub gets an existing StreamInputIotHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamInputIotHub(ctx *pulumi.Context,
	name string, id pulumi.ID, state *StreamInputIotHubState, opts ...pulumi.ResourceOpt) (*StreamInputIotHub, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["endpoint"] = state.Endpoint
		inputs["eventhubConsumerGroupName"] = state.EventhubConsumerGroupName
		inputs["iothubNamespace"] = state.IothubNamespace
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["serialization"] = state.Serialization
		inputs["sharedAccessPolicyKey"] = state.SharedAccessPolicyKey
		inputs["sharedAccessPolicyName"] = state.SharedAccessPolicyName
		inputs["streamAnalyticsJobName"] = state.StreamAnalyticsJobName
	}
	s, err := ctx.ReadResource("azure:streamanalytics/streamInputIotHub:StreamInputIotHub", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &StreamInputIotHub{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *StreamInputIotHub) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *StreamInputIotHub) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
func (r *StreamInputIotHub) Endpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endpoint"])
}

// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
func (r *StreamInputIotHub) EventhubConsumerGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["eventhubConsumerGroupName"])
}

// The name or the URI of the IoT Hub.
func (r *StreamInputIotHub) IothubNamespace() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["iothubNamespace"])
}

// The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
func (r *StreamInputIotHub) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
func (r *StreamInputIotHub) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `serialization` block as defined below.
func (r *StreamInputIotHub) Serialization() *pulumi.Output {
	return r.s.State["serialization"]
}

// The shared access policy key for the specified shared access policy.
func (r *StreamInputIotHub) SharedAccessPolicyKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sharedAccessPolicyKey"])
}

// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
func (r *StreamInputIotHub) SharedAccessPolicyName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sharedAccessPolicyName"])
}

// The name of the Stream Analytics Job. Changing this forces a new resource to be created. 
func (r *StreamInputIotHub) StreamAnalyticsJobName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["streamAnalyticsJobName"])
}

// Input properties used for looking up and filtering StreamInputIotHub resources.
type StreamInputIotHubState struct {
	// The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
	Endpoint interface{}
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName interface{}
	// The name or the URI of the IoT Hub.
	IothubNamespace interface{}
	// The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `serialization` block as defined below.
	Serialization interface{}
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey interface{}
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName interface{}
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created. 
	StreamAnalyticsJobName interface{}
}

// The set of arguments for constructing a StreamInputIotHub resource.
type StreamInputIotHubArgs struct {
	// The IoT Hub endpoint to connect to (ie. messages/events, messages/operationsMonitoringEvents, etc.).
	Endpoint interface{}
	// The name of an Event Hub Consumer Group that should be used to read events from the Event Hub. Specifying distinct consumer group names for multiple inputs allows each of those inputs to receive the same events from the Event Hub.
	EventhubConsumerGroupName interface{}
	// The name or the URI of the IoT Hub.
	IothubNamespace interface{}
	// The name of the Stream Input IoTHub. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `serialization` block as defined below.
	Serialization interface{}
	// The shared access policy key for the specified shared access policy.
	SharedAccessPolicyKey interface{}
	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
	SharedAccessPolicyName interface{}
	// The name of the Stream Analytics Job. Changing this forces a new resource to be created. 
	StreamAnalyticsJobName interface{}
}
