// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Policy
{
    public static class GetPolicyDefintion
    {
        /// <summary>
        /// Use this data source to access information about a Policy Definition, both custom and built in. Retrieves Policy Definitions from your current subscription by default.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPolicyDefintionResult> InvokeAsync(GetPolicyDefintionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyDefintionResult>("azure:policy/getPolicyDefintion:getPolicyDefintion", args ?? new GetPolicyDefintionArgs(), options.WithVersion());
    }


    public sealed class GetPolicyDefintionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Policy Definition.
        /// </summary>
        [Input("displayName", required: true)]
        public string DisplayName { get; set; } = null!;

        /// <summary>
        /// Only retrieve Policy Definitions from this Management Group.
        /// </summary>
        [Input("managementGroupId")]
        public string? ManagementGroupId { get; set; }

        public GetPolicyDefintionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyDefintionResult
    {
        /// <summary>
        /// The Description of the Policy.
        /// </summary>
        public readonly string Description;
        public readonly string DisplayName;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ManagementGroupId;
        /// <summary>
        /// Any Metadata defined in the Policy.
        /// </summary>
        public readonly string Metadata;
        /// <summary>
        /// The Name of the Policy Definition.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Any Parameters defined in the Policy.
        /// </summary>
        public readonly string Parameters;
        /// <summary>
        /// The Rule as defined (in JSON) in the Policy.
        /// </summary>
        public readonly string PolicyRule;
        /// <summary>
        /// The Type of the Policy, such as `Microsoft.Authorization/policyDefinitions`.
        /// </summary>
        public readonly string PolicyType;
        /// <summary>
        /// The Type of Policy.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPolicyDefintionResult(
            string description,

            string displayName,

            string id,

            string? managementGroupId,

            string metadata,

            string name,

            string parameters,

            string policyRule,

            string policyType,

            string type)
        {
            Description = description;
            DisplayName = displayName;
            Id = id;
            ManagementGroupId = managementGroupId;
            Metadata = metadata;
            Name = name;
            Parameters = parameters;
            PolicyRule = policyRule;
            PolicyType = policyType;
            Type = type;
        }
    }
}
