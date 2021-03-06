// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Azure Machine Learning Workspace
type Workspace struct {
	pulumi.CustomResourceState

	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringOutput `pulumi:"applicationInsightsId"`
	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryId pulumi.StringPtrOutput `pulumi:"containerRegistryId"`
	// The description of this Machine Learning Workspace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Friendly name for this Machine Learning Workspace.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// An `identity` block defined below.
	Identity WorkspaceIdentityOutput `pulumi:"identity"`
	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// SKU/edition of the Machine Learning Workspace, possible values are `Basic` for a basic workspace or `Enterprise` for a feature rich workspace. Defaults to `Basic`.
	SkuName pulumi.StringPtrOutput `pulumi:"skuName"`
	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil || args.ApplicationInsightsId == nil {
		return nil, errors.New("missing required argument 'ApplicationInsightsId'")
	}
	if args == nil || args.Identity == nil {
		return nil, errors.New("missing required argument 'Identity'")
	}
	if args == nil || args.KeyVaultId == nil {
		return nil, errors.New("missing required argument 'KeyVaultId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageAccountId == nil {
		return nil, errors.New("missing required argument 'StorageAccountId'")
	}
	if args == nil {
		args = &WorkspaceArgs{}
	}
	var resource Workspace
	err := ctx.RegisterResource("azure:machinelearning/workspace:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure:machinelearning/workspace:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ApplicationInsightsId *string `pulumi:"applicationInsightsId"`
	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryId *string `pulumi:"containerRegistryId"`
	// The description of this Machine Learning Workspace.
	Description *string `pulumi:"description"`
	// Friendly name for this Machine Learning Workspace.
	FriendlyName *string `pulumi:"friendlyName"`
	// An `identity` block defined below.
	Identity *WorkspaceIdentity `pulumi:"identity"`
	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// SKU/edition of the Machine Learning Workspace, possible values are `Basic` for a basic workspace or `Enterprise` for a feature rich workspace. Defaults to `Basic`.
	SkuName *string `pulumi:"skuName"`
	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags map[string]string `pulumi:"tags"`
}

type WorkspaceState struct {
	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringPtrInput
	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryId pulumi.StringPtrInput
	// The description of this Machine Learning Workspace.
	Description pulumi.StringPtrInput
	// Friendly name for this Machine Learning Workspace.
	FriendlyName pulumi.StringPtrInput
	// An `identity` block defined below.
	Identity WorkspaceIdentityPtrInput
	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringPtrInput
	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// SKU/edition of the Machine Learning Workspace, possible values are `Basic` for a basic workspace or `Enterprise` for a feature rich workspace. Defaults to `Basic`.
	SkuName pulumi.StringPtrInput
	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags pulumi.StringMapInput
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ApplicationInsightsId string `pulumi:"applicationInsightsId"`
	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryId *string `pulumi:"containerRegistryId"`
	// The description of this Machine Learning Workspace.
	Description *string `pulumi:"description"`
	// Friendly name for this Machine Learning Workspace.
	FriendlyName *string `pulumi:"friendlyName"`
	// An `identity` block defined below.
	Identity WorkspaceIdentity `pulumi:"identity"`
	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	KeyVaultId string `pulumi:"keyVaultId"`
	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU/edition of the Machine Learning Workspace, possible values are `Basic` for a basic workspace or `Enterprise` for a feature rich workspace. Defaults to `Basic`.
	SkuName *string `pulumi:"skuName"`
	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	StorageAccountId string `pulumi:"storageAccountId"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The ID of the Application Insights associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringInput
	// The ID of the container registry associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	ContainerRegistryId pulumi.StringPtrInput
	// The description of this Machine Learning Workspace.
	Description pulumi.StringPtrInput
	// Friendly name for this Machine Learning Workspace.
	FriendlyName pulumi.StringPtrInput
	// An `identity` block defined below.
	Identity WorkspaceIdentityInput
	// The ID of key vault associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	KeyVaultId pulumi.StringInput
	// Specifies the supported Azure location where the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Machine Learning Workspace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which the Machine Learning Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// SKU/edition of the Machine Learning Workspace, possible values are `Basic` for a basic workspace or `Enterprise` for a feature rich workspace. Defaults to `Basic`.
	SkuName pulumi.StringPtrInput
	// The ID of the Storage Account associated with this Machine Learning Workspace. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringInput
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags pulumi.StringMapInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}
