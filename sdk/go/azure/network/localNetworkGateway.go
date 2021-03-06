// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a local network gateway connection over which specific connections can be configured.
type LocalNetworkGateway struct {
	pulumi.CustomResourceState

	// The list of string CIDRs representing the
	// address spaces the gateway exposes.
	AddressSpaces pulumi.StringArrayOutput `pulumi:"addressSpaces"`
	// A `bgpSettings` block as defined below containing the
	// Local Network Gateway's BGP speaker settings.
	BgpSettings LocalNetworkGatewayBgpSettingsPtrOutput `pulumi:"bgpSettings"`
	// The IP address of the gateway to which to
	// connect.
	GatewayAddress pulumi.StringOutput `pulumi:"gatewayAddress"`
	// The location/region where the local network gateway is
	// created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the local network gateway. Changing this
	// forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to
	// create the local network gateway.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewLocalNetworkGateway registers a new resource with the given unique name, arguments, and options.
func NewLocalNetworkGateway(ctx *pulumi.Context,
	name string, args *LocalNetworkGatewayArgs, opts ...pulumi.ResourceOption) (*LocalNetworkGateway, error) {
	if args == nil || args.AddressSpaces == nil {
		return nil, errors.New("missing required argument 'AddressSpaces'")
	}
	if args == nil || args.GatewayAddress == nil {
		return nil, errors.New("missing required argument 'GatewayAddress'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &LocalNetworkGatewayArgs{}
	}
	var resource LocalNetworkGateway
	err := ctx.RegisterResource("azure:network/localNetworkGateway:LocalNetworkGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalNetworkGateway gets an existing LocalNetworkGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalNetworkGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalNetworkGatewayState, opts ...pulumi.ResourceOption) (*LocalNetworkGateway, error) {
	var resource LocalNetworkGateway
	err := ctx.ReadResource("azure:network/localNetworkGateway:LocalNetworkGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalNetworkGateway resources.
type localNetworkGatewayState struct {
	// The list of string CIDRs representing the
	// address spaces the gateway exposes.
	AddressSpaces []string `pulumi:"addressSpaces"`
	// A `bgpSettings` block as defined below containing the
	// Local Network Gateway's BGP speaker settings.
	BgpSettings *LocalNetworkGatewayBgpSettings `pulumi:"bgpSettings"`
	// The IP address of the gateway to which to
	// connect.
	GatewayAddress *string `pulumi:"gatewayAddress"`
	// The location/region where the local network gateway is
	// created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the local network gateway. Changing this
	// forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the local network gateway.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type LocalNetworkGatewayState struct {
	// The list of string CIDRs representing the
	// address spaces the gateway exposes.
	AddressSpaces pulumi.StringArrayInput
	// A `bgpSettings` block as defined below containing the
	// Local Network Gateway's BGP speaker settings.
	BgpSettings LocalNetworkGatewayBgpSettingsPtrInput
	// The IP address of the gateway to which to
	// connect.
	GatewayAddress pulumi.StringPtrInput
	// The location/region where the local network gateway is
	// created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the local network gateway. Changing this
	// forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the local network gateway.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (LocalNetworkGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*localNetworkGatewayState)(nil)).Elem()
}

type localNetworkGatewayArgs struct {
	// The list of string CIDRs representing the
	// address spaces the gateway exposes.
	AddressSpaces []string `pulumi:"addressSpaces"`
	// A `bgpSettings` block as defined below containing the
	// Local Network Gateway's BGP speaker settings.
	BgpSettings *LocalNetworkGatewayBgpSettings `pulumi:"bgpSettings"`
	// The IP address of the gateway to which to
	// connect.
	GatewayAddress string `pulumi:"gatewayAddress"`
	// The location/region where the local network gateway is
	// created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the local network gateway. Changing this
	// forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the local network gateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LocalNetworkGateway resource.
type LocalNetworkGatewayArgs struct {
	// The list of string CIDRs representing the
	// address spaces the gateway exposes.
	AddressSpaces pulumi.StringArrayInput
	// A `bgpSettings` block as defined below containing the
	// Local Network Gateway's BGP speaker settings.
	BgpSettings LocalNetworkGatewayBgpSettingsPtrInput
	// The IP address of the gateway to which to
	// connect.
	GatewayAddress pulumi.StringInput
	// The location/region where the local network gateway is
	// created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the local network gateway. Changing this
	// forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the local network gateway.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (LocalNetworkGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localNetworkGatewayArgs)(nil)).Elem()
}
