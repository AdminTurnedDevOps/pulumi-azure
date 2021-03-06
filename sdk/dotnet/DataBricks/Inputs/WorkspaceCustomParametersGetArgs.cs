// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataBricks.Inputs
{

    public sealed class WorkspaceCustomParametersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Are public IP Addresses not allowed?
        /// </summary>
        [Input("noPublicIp")]
        public Input<bool>? NoPublicIp { get; set; }

        /// <summary>
        /// The name of the Private Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        /// </summary>
        [Input("privateSubnetName")]
        public Input<string>? PrivateSubnetName { get; set; }

        /// <summary>
        /// The name of the Public Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        /// </summary>
        [Input("publicSubnetName")]
        public Input<string>? PublicSubnetName { get; set; }

        /// <summary>
        /// The ID of a Virtual Network where this Databricks Cluster should be created.
        /// </summary>
        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        public WorkspaceCustomParametersGetArgs()
        {
        }
    }
}
