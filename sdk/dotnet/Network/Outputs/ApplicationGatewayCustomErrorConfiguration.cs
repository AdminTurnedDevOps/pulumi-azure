// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayCustomErrorConfiguration
    {
        /// <summary>
        /// Error page URL of the application gateway customer error.
        /// </summary>
        public readonly string CustomErrorPageUrl;
        /// <summary>
        /// The ID of the Rewrite Rule Set
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Status code of the application gateway customer error. Possible values are `HttpStatus403` and `HttpStatus502`
        /// </summary>
        public readonly string StatusCode;

        [OutputConstructor]
        private ApplicationGatewayCustomErrorConfiguration(
            string customErrorPageUrl,

            string? id,

            string statusCode)
        {
            CustomErrorPageUrl = customErrorPageUrl;
            Id = id;
            StatusCode = statusCode;
        }
    }
}
