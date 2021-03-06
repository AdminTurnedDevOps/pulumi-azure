// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MariaDB
{
    public static class GetMariaDbServer
    {
        /// <summary>
        /// Use this data source to access information about an existing MariaDB Server.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMariaDbServerResult> InvokeAsync(GetMariaDbServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMariaDbServerResult>("azure:mariadb/getMariaDbServer:getMariaDbServer", args ?? new GetMariaDbServerArgs(), options.WithVersion());
    }


    public sealed class GetMariaDbServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the MariaDB Server to retrieve information about.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group where the MariaDB Server exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMariaDbServerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMariaDbServerResult
    {
        /// <summary>
        /// The Administrator Login for the MariaDB Server.
        /// </summary>
        public readonly string AdministratorLogin;
        /// <summary>
        /// The password associated with the `administrator_login` for the MariaDB Server.
        /// </summary>
        public readonly string AdministratorLoginPassword;
        /// <summary>
        /// The FQDN of the MariaDB Server.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The SKU Name for this MariaDB Server. 
        /// </summary>
        public readonly string SkuName;
        /// <summary>
        /// The SSL being enforced on connections.
        /// </summary>
        public readonly string SslEnforcement;
        /// <summary>
        /// A `storage_profile` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMariaDbServerStorageProfileResult> StorageProfiles;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// ---
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The version of MariaDB being used.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetMariaDbServerResult(
            string administratorLogin,

            string administratorLoginPassword,

            string fqdn,

            string id,

            string location,

            string name,

            string resourceGroupName,

            string skuName,

            string sslEnforcement,

            ImmutableArray<Outputs.GetMariaDbServerStorageProfileResult> storageProfiles,

            ImmutableDictionary<string, string> tags,

            string version)
        {
            AdministratorLogin = administratorLogin;
            AdministratorLoginPassword = administratorLoginPassword;
            Fqdn = fqdn;
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            SkuName = skuName;
            SslEnforcement = sslEnforcement;
            StorageProfiles = storageProfiles;
            Tags = tags;
            Version = version;
        }
    }
}
