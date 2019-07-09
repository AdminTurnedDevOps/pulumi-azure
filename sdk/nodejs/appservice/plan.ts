// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage an App Service Plan component.
 * 
 * ## Example Usage (Dedicated)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West Europe",
 *     name: "api-rg-pro",
 * });
 * const testPlan = new azure.appservice.Plan("test", {
 *     location: testResourceGroup.location,
 *     name: "api-appserviceplan-pro",
 *     resourceGroupName: testResourceGroup.name,
 *     sku: {
 *         size: "S1",
 *         tier: "Standard",
 *     },
 * });
 * ```
 * 
 * ## Example Usage (Shared / Consumption Plan)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West Europe",
 *     name: "api-rg-pro",
 * });
 * const testPlan = new azure.appservice.Plan("test", {
 *     kind: "FunctionApp",
 *     location: testResourceGroup.location,
 *     name: "api-appserviceplan-pro",
 *     resourceGroupName: testResourceGroup.name,
 *     sku: {
 *         size: "Y1",
 *         tier: "Dynamic",
 *     },
 * });
 * ```
 * 
 * ## Example Usage (Linux)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West Europe",
 *     name: "api-rg-pro",
 * });
 * const testPlan = new azure.appservice.Plan("test", {
 *     kind: "Linux",
 *     location: testResourceGroup.location,
 *     name: "api-appserviceplan-pro",
 *     reserved: true,
 *     resourceGroupName: testResourceGroup.name,
 *     sku: {
 *         size: "S1",
 *         tier: "Standard",
 *     },
 * });
 * ```
 * 
 * ## Example Usage (Windows Container)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West Europe",
 *     name: "api-rg-pro",
 * });
 * const testPlan = new azure.appservice.Plan("test", {
 *     isXenon: true,
 *     kind: "xenon",
 *     location: testResourceGroup.location,
 *     name: "api-appserviceplan-pro",
 *     resourceGroupName: testResourceGroup.name,
 *     sku: {
 *         size: "PC2",
 *         tier: "PremiumContainer",
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/app_service_plan.html.markdown.
 */
export class Plan extends pulumi.CustomResource {
    /**
     * Get an existing Plan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PlanState, opts?: pulumi.CustomResourceOptions): Plan {
        return new Plan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appservice/plan:Plan';

    /**
     * Returns true if the given object is an instance of Plan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Plan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Plan.__pulumiType;
    }

    /**
     * The ID of the App Service Environment where the App Service Plan should be located. Changing forces a new resource to be created.
     */
    public readonly appServiceEnvironmentId!: pulumi.Output<string>;
    public readonly isXenon!: pulumi.Output<boolean | undefined>;
    /**
     * The kind of the App Service Plan to create. Possible values are `Windows` (also available as `App`), `Linux`, `elastic` (for Premium Consumption) and `FunctionApp` (for a Consumption Plan). Defaults to `Windows`. Changing this forces a new resource to be created.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
     */
    public readonly maximumElasticWorkerCount!: pulumi.Output<number>;
    /**
     * The maximum number of workers supported with the App Service Plan's sku.
     */
    public /*out*/ readonly maximumNumberOfWorkers!: pulumi.Output<number>;
    /**
     * Specifies the name of the App Service Plan component. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Can Apps assigned to this App Service Plan be scaled independently? If set to `false` apps assigned to this plan will scale to all instances of the plan.  Defaults to `false`.
     */
    public readonly perSiteScaling!: pulumi.Output<boolean>;
    public readonly properties!: pulumi.Output<{ appServiceEnvironmentId: string, perSiteScaling: boolean, reserved: boolean }>;
    /**
     * Is this App Service Plan `Reserved`. Defaults to `false`.
     */
    public readonly reserved!: pulumi.Output<boolean>;
    /**
     * The name of the resource group in which to create the App Service Plan component.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `sku` block as documented below.
     */
    public readonly sku!: pulumi.Output<{ capacity: number, size: string, tier: string }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a Plan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PlanArgs | PlanState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PlanState | undefined;
            inputs["appServiceEnvironmentId"] = state ? state.appServiceEnvironmentId : undefined;
            inputs["isXenon"] = state ? state.isXenon : undefined;
            inputs["kind"] = state ? state.kind : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["maximumElasticWorkerCount"] = state ? state.maximumElasticWorkerCount : undefined;
            inputs["maximumNumberOfWorkers"] = state ? state.maximumNumberOfWorkers : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["perSiteScaling"] = state ? state.perSiteScaling : undefined;
            inputs["properties"] = state ? state.properties : undefined;
            inputs["reserved"] = state ? state.reserved : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as PlanArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["appServiceEnvironmentId"] = args ? args.appServiceEnvironmentId : undefined;
            inputs["isXenon"] = args ? args.isXenon : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maximumElasticWorkerCount"] = args ? args.maximumElasticWorkerCount : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["perSiteScaling"] = args ? args.perSiteScaling : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["reserved"] = args ? args.reserved : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["maximumNumberOfWorkers"] = undefined /*out*/;
        }
        super(Plan.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Plan resources.
 */
export interface PlanState {
    /**
     * The ID of the App Service Environment where the App Service Plan should be located. Changing forces a new resource to be created.
     */
    readonly appServiceEnvironmentId?: pulumi.Input<string>;
    readonly isXenon?: pulumi.Input<boolean>;
    /**
     * The kind of the App Service Plan to create. Possible values are `Windows` (also available as `App`), `Linux`, `elastic` (for Premium Consumption) and `FunctionApp` (for a Consumption Plan). Defaults to `Windows`. Changing this forces a new resource to be created.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
     */
    readonly maximumElasticWorkerCount?: pulumi.Input<number>;
    /**
     * The maximum number of workers supported with the App Service Plan's sku.
     */
    readonly maximumNumberOfWorkers?: pulumi.Input<number>;
    /**
     * Specifies the name of the App Service Plan component. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Can Apps assigned to this App Service Plan be scaled independently? If set to `false` apps assigned to this plan will scale to all instances of the plan.  Defaults to `false`.
     */
    readonly perSiteScaling?: pulumi.Input<boolean>;
    readonly properties?: pulumi.Input<{ appServiceEnvironmentId?: pulumi.Input<string>, perSiteScaling?: pulumi.Input<boolean>, reserved?: pulumi.Input<boolean> }>;
    /**
     * Is this App Service Plan `Reserved`. Defaults to `false`.
     */
    readonly reserved?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the App Service Plan component.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `sku` block as documented below.
     */
    readonly sku?: pulumi.Input<{ capacity?: pulumi.Input<number>, size: pulumi.Input<string>, tier: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Plan resource.
 */
export interface PlanArgs {
    /**
     * The ID of the App Service Environment where the App Service Plan should be located. Changing forces a new resource to be created.
     */
    readonly appServiceEnvironmentId?: pulumi.Input<string>;
    readonly isXenon?: pulumi.Input<boolean>;
    /**
     * The kind of the App Service Plan to create. Possible values are `Windows` (also available as `App`), `Linux`, `elastic` (for Premium Consumption) and `FunctionApp` (for a Consumption Plan). Defaults to `Windows`. Changing this forces a new resource to be created.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.
     */
    readonly maximumElasticWorkerCount?: pulumi.Input<number>;
    /**
     * Specifies the name of the App Service Plan component. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Can Apps assigned to this App Service Plan be scaled independently? If set to `false` apps assigned to this plan will scale to all instances of the plan.  Defaults to `false`.
     */
    readonly perSiteScaling?: pulumi.Input<boolean>;
    readonly properties?: pulumi.Input<{ appServiceEnvironmentId?: pulumi.Input<string>, perSiteScaling?: pulumi.Input<boolean>, reserved?: pulumi.Input<boolean> }>;
    /**
     * Is this App Service Plan `Reserved`. Defaults to `false`.
     */
    readonly reserved?: pulumi.Input<boolean>;
    /**
     * The name of the resource group in which to create the App Service Plan component.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `sku` block as documented below.
     */
    readonly sku: pulumi.Input<{ capacity?: pulumi.Input<number>, size: pulumi.Input<string>, tier: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
