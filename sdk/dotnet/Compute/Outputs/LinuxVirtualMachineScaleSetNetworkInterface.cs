// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Outputs
{

    [OutputType]
    public sealed class LinuxVirtualMachineScaleSetNetworkInterface
    {
        /// <summary>
        /// A list of IP Addresses of DNS Servers which should be assigned to the Network Interface.
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// Does this Network Interface support Accelerated Networking? Defaults to `false`.
        /// </summary>
        public readonly bool? EnableAcceleratedNetworking;
        /// <summary>
        /// Does this Network Interface support IP Forwarding? Defaults to `false`.
        /// </summary>
        public readonly bool? EnableIpForwarding;
        /// <summary>
        /// One or more `ip_configuration` blocks as defined above.
        /// </summary>
        public readonly ImmutableArray<Outputs.LinuxVirtualMachineScaleSetNetworkInterfaceIpConfiguration> IpConfigurations;
        /// <summary>
        /// The Name which should be used for this Network Interface. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of a Network Security Group which should be assigned to this Network Interface.
        /// </summary>
        public readonly string? NetworkSecurityGroupId;
        /// <summary>
        /// Is this the Primary IP Configuration?
        /// </summary>
        public readonly bool? Primary;

        [OutputConstructor]
        private LinuxVirtualMachineScaleSetNetworkInterface(
            ImmutableArray<string> dnsServers,

            bool? enableAcceleratedNetworking,

            bool? enableIpForwarding,

            ImmutableArray<Outputs.LinuxVirtualMachineScaleSetNetworkInterfaceIpConfiguration> ipConfigurations,

            string name,

            string? networkSecurityGroupId,

            bool? primary)
        {
            DnsServers = dnsServers;
            EnableAcceleratedNetworking = enableAcceleratedNetworking;
            EnableIpForwarding = enableIpForwarding;
            IpConfigurations = ipConfigurations;
            Name = name;
            NetworkSecurityGroupId = networkSecurityGroupId;
            Primary = primary;
        }
    }
}