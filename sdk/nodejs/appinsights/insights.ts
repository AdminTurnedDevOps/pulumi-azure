// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Application Insights component.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/application_insights.html.markdown.
 */
export class Insights extends pulumi.CustomResource {
    /**
     * Get an existing Insights resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InsightsState, opts?: pulumi.CustomResourceOptions): Insights {
        return new Insights(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appinsights/insights:Insights';

    /**
     * Returns true if the given object is an instance of Insights.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Insights {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Insights.__pulumiType;
    }

    /**
     * The App ID associated with this Application Insights component.
     */
    public /*out*/ readonly appId!: pulumi.Output<string>;
    /**
     * Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
     */
    public readonly applicationType!: pulumi.Output<string>;
    /**
     * Specifies the Application Insights component daily data volume cap in GB.
     */
    public readonly dailyDataCapInGb!: pulumi.Output<number>;
    /**
     * Specifies if a notification email will be send when the daily data volume cap is met.
     */
    public readonly dailyDataCapNotificationsDisabled!: pulumi.Output<boolean>;
    /**
     * The Instrumentation Key for this Application Insights component.
     */
    public /*out*/ readonly instrumentationKey!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Application Insights component. Changing this forces a
     * new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to
     * create the Application Insights component.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`.
     */
    public readonly retentionInDays!: pulumi.Output<number | undefined>;
    /**
     * Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
     */
    public readonly samplingPercentage!: pulumi.Output<number | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Insights resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InsightsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InsightsArgs | InsightsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InsightsState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["applicationType"] = state ? state.applicationType : undefined;
            inputs["dailyDataCapInGb"] = state ? state.dailyDataCapInGb : undefined;
            inputs["dailyDataCapNotificationsDisabled"] = state ? state.dailyDataCapNotificationsDisabled : undefined;
            inputs["instrumentationKey"] = state ? state.instrumentationKey : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["retentionInDays"] = state ? state.retentionInDays : undefined;
            inputs["samplingPercentage"] = state ? state.samplingPercentage : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as InsightsArgs | undefined;
            if (!args || args.applicationType === undefined) {
                throw new Error("Missing required property 'applicationType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["applicationType"] = args ? args.applicationType : undefined;
            inputs["dailyDataCapInGb"] = args ? args.dailyDataCapInGb : undefined;
            inputs["dailyDataCapNotificationsDisabled"] = args ? args.dailyDataCapNotificationsDisabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["retentionInDays"] = args ? args.retentionInDays : undefined;
            inputs["samplingPercentage"] = args ? args.samplingPercentage : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["appId"] = undefined /*out*/;
            inputs["instrumentationKey"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Insights.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Insights resources.
 */
export interface InsightsState {
    /**
     * The App ID associated with this Application Insights component.
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
     */
    readonly applicationType?: pulumi.Input<string>;
    /**
     * Specifies the Application Insights component daily data volume cap in GB.
     */
    readonly dailyDataCapInGb?: pulumi.Input<number>;
    /**
     * Specifies if a notification email will be send when the daily data volume cap is met.
     */
    readonly dailyDataCapNotificationsDisabled?: pulumi.Input<boolean>;
    /**
     * The Instrumentation Key for this Application Insights component.
     */
    readonly instrumentationKey?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Application Insights component. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the Application Insights component.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`.
     */
    readonly retentionInDays?: pulumi.Input<number>;
    /**
     * Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
     */
    readonly samplingPercentage?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Insights resource.
 */
export interface InsightsArgs {
    /**
     * Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
     */
    readonly applicationType: pulumi.Input<string>;
    /**
     * Specifies the Application Insights component daily data volume cap in GB.
     */
    readonly dailyDataCapInGb?: pulumi.Input<number>;
    /**
     * Specifies if a notification email will be send when the daily data volume cap is met.
     */
    readonly dailyDataCapNotificationsDisabled?: pulumi.Input<boolean>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Application Insights component. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the Application Insights component.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the retention period in days. Possible values are `30`, `60`, `90`, `120`, `180`, `270`, `365`, `550` or `730`.
     */
    readonly retentionInDays?: pulumi.Input<number>;
    /**
     * Specifies the percentage of the data produced by the monitored application that is sampled for Application Insights telemetry.
     */
    readonly samplingPercentage?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
