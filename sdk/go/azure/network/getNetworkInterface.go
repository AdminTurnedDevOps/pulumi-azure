// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access the properties of an Azure Network Interface.
func LookupetworkInterface(ctx *pulumi.Context, args *GetNetworkInterfaceArgs) (*GetNetworkInterfaceResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	outputs, err := ctx.Invoke("azure:network/getNetworkInterface:getNetworkInterface", inputs)
	if err != nil {
		return nil, err
	}
	ret := GetNetworkInterfaceResult{}
	if v, ok := outputs["appliedDnsServers"]; ok {
		ret.AppliedDnsServers = v
	}
	if v, ok := outputs["dnsServers"]; ok {
		ret.DnsServers = v
	}
	if v, ok := outputs["enableAcceleratedNetworking"]; ok {
		ret.EnableAcceleratedNetworking = v
	}
	if v, ok := outputs["enableIpForwarding"]; ok {
		ret.EnableIpForwarding = v
	}
	if v, ok := outputs["internalDnsNameLabel"]; ok {
		ret.InternalDnsNameLabel = v
	}
	if v, ok := outputs["internalFqdn"]; ok {
		ret.InternalFqdn = v
	}
	if v, ok := outputs["ipConfigurations"]; ok {
		ret.IpConfigurations = v
	}
	if v, ok := outputs["location"]; ok {
		ret.Location = v
	}
	if v, ok := outputs["macAddress"]; ok {
		ret.MacAddress = v
	}
	if v, ok := outputs["networkSecurityGroupId"]; ok {
		ret.NetworkSecurityGroupId = v
	}
	if v, ok := outputs["privateIpAddress"]; ok {
		ret.PrivateIpAddress = v
	}
	if v, ok := outputs["privateIpAddresses"]; ok {
		ret.PrivateIpAddresses = v
	}
	if v, ok := outputs["tags"]; ok {
		ret.Tags = v
	}
	if v, ok := outputs["virtualMachineId"]; ok {
		ret.VirtualMachineId = v
	}
	return &ret, nil
}

// A collection of arguments for invoking getNetworkInterface.
type GetNetworkInterfaceArgs struct {
	// Specifies the name of the Network Interface.
	Name interface{}
	// Specifies the name of the resource group the Network Interface is located in.
	ResourceGroupName interface{}
}

// A collection of values returned by getNetworkInterface.
type GetNetworkInterfaceResult struct {
	// List of DNS servers applied to the specified network interface.
	AppliedDnsServers interface{}
	// The list of DNS servers used by the specified network interface.
	DnsServers interface{}
	// Indicates if accelerated networking is set on the specified network interface.
	EnableAcceleratedNetworking interface{}
	// Indicate if IP forwarding is set on the specified network interface.
	EnableIpForwarding interface{}
	// The internal dns name label of the specified network interface.
	InternalDnsNameLabel interface{}
	// The internal FQDN associated to the specified network interface.
	InternalFqdn interface{}
	// The list of IP configurations associated to the specified network interface.
	IpConfigurations interface{}
	// The location of the specified network interface.
	Location interface{}
	// The MAC address used by the specified network interface.
	MacAddress interface{}
	// The ID of the network security group associated to the specified network interface.
	NetworkSecurityGroupId interface{}
	// The primary private ip address associated to the specified network interface.
	PrivateIpAddress interface{}
	// The list of private ip addresses associates to the specified network interface.
	PrivateIpAddresses interface{}
	// List the tags assocatied to the specified network interface.
	Tags interface{}
	// The ID of the virtual machine that the specified network interface is attached to.
	VirtualMachineId interface{}
}
