// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Private Endpoint.
//
// > **NOTE** Private Endpoint is currently in Public Preview.
//
// Azure Private Endpoint is a network interface that connects you privately and securely to a service powered by Azure Private Link. Private Endpoint uses a private IP address from your VNet, effectively bringing the service into your VNet. The service could be an Azure service such as Azure Storage, SQL, etc. or your own Private Link Service.
type Endpoint struct {
	pulumi.CustomResourceState

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `privateServiceConnection` block as defined below.
	PrivateServiceConnection EndpointPrivateServiceConnectionOutput `pulumi:"privateServiceConnection"`
	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil || args.PrivateServiceConnection == nil {
		return nil, errors.New("missing required argument 'PrivateServiceConnection'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	if args == nil {
		args = &EndpointArgs{}
	}
	var resource Endpoint
	err := ctx.RegisterResource("azure:privatelink/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azure:privatelink/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `privateServiceConnection` block as defined below.
	PrivateServiceConnection *EndpointPrivateServiceConnection `pulumi:"privateServiceConnection"`
	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	SubnetId *string `pulumi:"subnetId"`
}

type EndpointState struct {
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `privateServiceConnection` block as defined below.
	PrivateServiceConnection EndpointPrivateServiceConnectionPtrInput
	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	SubnetId pulumi.StringPtrInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `privateServiceConnection` block as defined below.
	PrivateServiceConnection EndpointPrivateServiceConnection `pulumi:"privateServiceConnection"`
	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the Name of the Private Endpoint. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `privateServiceConnection` block as defined below.
	PrivateServiceConnection EndpointPrivateServiceConnectionInput
	// Specifies the Name of the Resource Group within which the Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The ID of the Subnet from which Private IP Addresses will be allocated for this Private Endpoint. Changing this forces a new resource to be created.
	SubnetId pulumi.StringInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}
