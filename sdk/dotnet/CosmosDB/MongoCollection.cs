// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB
{
    /// <summary>
    /// Manages a Mongo Collection within a Cosmos DB Account.
    /// </summary>
    public partial class MongoCollection : Pulumi.CustomResource
    {
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// The default Time To Live in seconds. If the value is `0` items are not automatically expired.
        /// </summary>
        [Output("defaultTtlSeconds")]
        public Output<int?> DefaultTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The name of the key to partition on for sharding. There must not be any other unique index keys.
        /// </summary>
        [Output("shardKey")]
        public Output<string?> ShardKey { get; private set; } = null!;

        /// <summary>
        /// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
        /// </summary>
        [Output("throughput")]
        public Output<int> Throughput { get; private set; } = null!;


        /// <summary>
        /// Create a MongoCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MongoCollection(string name, MongoCollectionArgs args, CustomResourceOptions? options = null)
            : base("azure:cosmosdb/mongoCollection:MongoCollection", name, args ?? new MongoCollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MongoCollection(string name, Input<string> id, MongoCollectionState? state = null, CustomResourceOptions? options = null)
            : base("azure:cosmosdb/mongoCollection:MongoCollection", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MongoCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MongoCollection Get(string name, Input<string> id, MongoCollectionState? state = null, CustomResourceOptions? options = null)
        {
            return new MongoCollection(name, id, state, options);
        }
    }

    public sealed class MongoCollectionArgs : Pulumi.ResourceArgs
    {
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The default Time To Live in seconds. If the value is `0` items are not automatically expired.
        /// </summary>
        [Input("defaultTtlSeconds")]
        public Input<int>? DefaultTtlSeconds { get; set; }

        /// <summary>
        /// Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the key to partition on for sharding. There must not be any other unique index keys.
        /// </summary>
        [Input("shardKey")]
        public Input<string>? ShardKey { get; set; }

        /// <summary>
        /// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
        /// </summary>
        [Input("throughput")]
        public Input<int>? Throughput { get; set; }

        public MongoCollectionArgs()
        {
        }
    }

    public sealed class MongoCollectionState : Pulumi.ResourceArgs
    {
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// The name of the Cosmos DB Mongo Database in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The default Time To Live in seconds. If the value is `0` items are not automatically expired.
        /// </summary>
        [Input("defaultTtlSeconds")]
        public Input<int>? DefaultTtlSeconds { get; set; }

        /// <summary>
        /// Specifies the name of the Cosmos DB Mongo Collection. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which the Cosmos DB Mongo Collection is created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The name of the key to partition on for sharding. There must not be any other unique index keys.
        /// </summary>
        [Input("shardKey")]
        public Input<string>? ShardKey { get; set; }

        /// <summary>
        /// The throughput of the MongoDB collection (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon database creation otherwise it cannot be updated without a manual resource destroy-apply.
        /// </summary>
        [Input("throughput")]
        public Input<int>? Throughput { get; set; }

        public MongoCollectionState()
        {
        }
    }
}
