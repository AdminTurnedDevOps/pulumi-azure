// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a HDInsight HBase Cluster.
type HBaseCluster struct {
	s *pulumi.ResourceState
}

// NewHBaseCluster registers a new resource with the given unique name, arguments, and options.
func NewHBaseCluster(ctx *pulumi.Context,
	name string, args *HBaseClusterArgs, opts ...pulumi.ResourceOpt) (*HBaseCluster, error) {
	if args == nil || args.ClusterVersion == nil {
		return nil, errors.New("missing required argument 'ClusterVersion'")
	}
	if args == nil || args.ComponentVersion == nil {
		return nil, errors.New("missing required argument 'ComponentVersion'")
	}
	if args == nil || args.Gateway == nil {
		return nil, errors.New("missing required argument 'Gateway'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Roles == nil {
		return nil, errors.New("missing required argument 'Roles'")
	}
	if args == nil || args.StorageAccounts == nil {
		return nil, errors.New("missing required argument 'StorageAccounts'")
	}
	if args == nil || args.Tier == nil {
		return nil, errors.New("missing required argument 'Tier'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["clusterVersion"] = nil
		inputs["componentVersion"] = nil
		inputs["gateway"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["roles"] = nil
		inputs["storageAccounts"] = nil
		inputs["tags"] = nil
		inputs["tier"] = nil
	} else {
		inputs["clusterVersion"] = args.ClusterVersion
		inputs["componentVersion"] = args.ComponentVersion
		inputs["gateway"] = args.Gateway
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["roles"] = args.Roles
		inputs["storageAccounts"] = args.StorageAccounts
		inputs["tags"] = args.Tags
		inputs["tier"] = args.Tier
	}
	inputs["httpsEndpoint"] = nil
	inputs["sshEndpoint"] = nil
	s, err := ctx.RegisterResource("azure:hdinsight/hBaseCluster:HBaseCluster", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HBaseCluster{s: s}, nil
}

// GetHBaseCluster gets an existing HBaseCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHBaseCluster(ctx *pulumi.Context,
	name string, id pulumi.ID, state *HBaseClusterState, opts ...pulumi.ResourceOpt) (*HBaseCluster, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["clusterVersion"] = state.ClusterVersion
		inputs["componentVersion"] = state.ComponentVersion
		inputs["gateway"] = state.Gateway
		inputs["httpsEndpoint"] = state.HttpsEndpoint
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["roles"] = state.Roles
		inputs["sshEndpoint"] = state.SshEndpoint
		inputs["storageAccounts"] = state.StorageAccounts
		inputs["tags"] = state.Tags
		inputs["tier"] = state.Tier
	}
	s, err := ctx.ReadResource("azure:hdinsight/hBaseCluster:HBaseCluster", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HBaseCluster{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *HBaseCluster) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *HBaseCluster) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
func (r *HBaseCluster) ClusterVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterVersion"])
}

// A `component_version` block as defined below.
func (r *HBaseCluster) ComponentVersion() *pulumi.Output {
	return r.s.State["componentVersion"]
}

// A `gateway` block as defined below.
func (r *HBaseCluster) Gateway() *pulumi.Output {
	return r.s.State["gateway"]
}

// The HTTPS Connectivity Endpoint for this HDInsight HBase Cluster.
func (r *HBaseCluster) HttpsEndpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["httpsEndpoint"])
}

// Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
func (r *HBaseCluster) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
func (r *HBaseCluster) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
func (r *HBaseCluster) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `roles` block as defined below.
func (r *HBaseCluster) Roles() *pulumi.Output {
	return r.s.State["roles"]
}

// The SSH Connectivity Endpoint for this HDInsight HBase Cluster.
func (r *HBaseCluster) SshEndpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sshEndpoint"])
}

// One or more `storage_account` block as defined below.
func (r *HBaseCluster) StorageAccounts() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["storageAccounts"])
}

// A map of Tags which should be assigned to this HDInsight HBase Cluster.
func (r *HBaseCluster) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
func (r *HBaseCluster) Tier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tier"])
}

// Input properties used for looking up and filtering HBaseCluster resources.
type HBaseClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion interface{}
	// A `component_version` block as defined below.
	ComponentVersion interface{}
	// A `gateway` block as defined below.
	Gateway interface{}
	// The HTTPS Connectivity Endpoint for this HDInsight HBase Cluster.
	HttpsEndpoint interface{}
	// Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
	Name interface{}
	// Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `roles` block as defined below.
	Roles interface{}
	// The SSH Connectivity Endpoint for this HDInsight HBase Cluster.
	SshEndpoint interface{}
	// One or more `storage_account` block as defined below.
	StorageAccounts interface{}
	// A map of Tags which should be assigned to this HDInsight HBase Cluster.
	Tags interface{}
	// Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier interface{}
}

// The set of arguments for constructing a HBaseCluster resource.
type HBaseClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion interface{}
	// A `component_version` block as defined below.
	ComponentVersion interface{}
	// A `gateway` block as defined below.
	Gateway interface{}
	// Specifies the Azure Region which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
	Name interface{}
	// Specifies the name of the Resource Group in which this HDInsight HBase Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `roles` block as defined below.
	Roles interface{}
	// One or more `storage_account` block as defined below.
	StorageAccounts interface{}
	// A map of Tags which should be assigned to this HDInsight HBase Cluster.
	Tags interface{}
	// Specifies the Tier which should be used for this HDInsight HBase Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier interface{}
}