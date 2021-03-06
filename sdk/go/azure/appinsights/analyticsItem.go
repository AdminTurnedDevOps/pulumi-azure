// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appinsights

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Application Insights Analytics Item component.
type AnalyticsItem struct {
	pulumi.CustomResourceState

	// The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringOutput `pulumi:"applicationInsightsId"`
	// The content for the Analytics Item, for example the query text if `type` is `query`.
	Content pulumi.StringOutput `pulumi:"content"`
	// The alias to use for the function. Required when `type` is `function`.
	FunctionAlias pulumi.StringPtrOutput `pulumi:"functionAlias"`
	// Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// A string containing the time the Analytics Item was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// A string containing the time the Analytics Item was last modified.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
	Type pulumi.StringOutput `pulumi:"type"`
	// A string indicating the version of the query format
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewAnalyticsItem registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsItem(ctx *pulumi.Context,
	name string, args *AnalyticsItemArgs, opts ...pulumi.ResourceOption) (*AnalyticsItem, error) {
	if args == nil || args.ApplicationInsightsId == nil {
		return nil, errors.New("missing required argument 'ApplicationInsightsId'")
	}
	if args == nil || args.Content == nil {
		return nil, errors.New("missing required argument 'Content'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &AnalyticsItemArgs{}
	}
	var resource AnalyticsItem
	err := ctx.RegisterResource("azure:appinsights/analyticsItem:AnalyticsItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyticsItem gets an existing AnalyticsItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsItemState, opts ...pulumi.ResourceOption) (*AnalyticsItem, error) {
	var resource AnalyticsItem
	err := ctx.ReadResource("azure:appinsights/analyticsItem:AnalyticsItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalyticsItem resources.
type analyticsItemState struct {
	// The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
	ApplicationInsightsId *string `pulumi:"applicationInsightsId"`
	// The content for the Analytics Item, for example the query text if `type` is `query`.
	Content *string `pulumi:"content"`
	// The alias to use for the function. Required when `type` is `function`.
	FunctionAlias *string `pulumi:"functionAlias"`
	// Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
	Scope *string `pulumi:"scope"`
	// A string containing the time the Analytics Item was created.
	TimeCreated *string `pulumi:"timeCreated"`
	// A string containing the time the Analytics Item was last modified.
	TimeModified *string `pulumi:"timeModified"`
	// The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
	Type *string `pulumi:"type"`
	// A string indicating the version of the query format
	Version *string `pulumi:"version"`
}

type AnalyticsItemState struct {
	// The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringPtrInput
	// The content for the Analytics Item, for example the query text if `type` is `query`.
	Content pulumi.StringPtrInput
	// The alias to use for the function. Required when `type` is `function`.
	FunctionAlias pulumi.StringPtrInput
	// Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
	Scope pulumi.StringPtrInput
	// A string containing the time the Analytics Item was created.
	TimeCreated pulumi.StringPtrInput
	// A string containing the time the Analytics Item was last modified.
	TimeModified pulumi.StringPtrInput
	// The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
	Type pulumi.StringPtrInput
	// A string indicating the version of the query format
	Version pulumi.StringPtrInput
}

func (AnalyticsItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsItemState)(nil)).Elem()
}

type analyticsItemArgs struct {
	// The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
	ApplicationInsightsId string `pulumi:"applicationInsightsId"`
	// The content for the Analytics Item, for example the query text if `type` is `query`.
	Content string `pulumi:"content"`
	// The alias to use for the function. Required when `type` is `function`.
	FunctionAlias *string `pulumi:"functionAlias"`
	// Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
	Scope string `pulumi:"scope"`
	// The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a AnalyticsItem resource.
type AnalyticsItemArgs struct {
	// The ID of the Application Insights component on which the Analytics Item exists. Changing this forces a new resource to be created.
	ApplicationInsightsId pulumi.StringInput
	// The content for the Analytics Item, for example the query text if `type` is `query`.
	Content pulumi.StringInput
	// The alias to use for the function. Required when `type` is `function`.
	FunctionAlias pulumi.StringPtrInput
	// Specifies the name of the Application Insights Analytics Item. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The scope for the Analytics Item. Can be `shared` or `user`. Changing this forces a new resource to be created. Must be `shared` for functions.
	Scope pulumi.StringInput
	// The type of Analytics Item to create. Can be one of `query`, `function`, `folder`, `recent`. Changing this forces a new resource to be created.
	Type pulumi.StringInput
}

func (AnalyticsItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsItemArgs)(nil)).Elem()
}
