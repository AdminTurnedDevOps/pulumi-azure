// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Enables you to manage DNS TXT Records within Azure Private DNS.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/private_dns_txt_record.html.markdown.
 */
export class TxtRecord extends pulumi.CustomResource {
    /**
     * Get an existing TxtRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TxtRecordState, opts?: pulumi.CustomResourceOptions): TxtRecord {
        return new TxtRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:privatedns/txtRecord:TxtRecord';

    /**
     * Returns true if the given object is an instance of TxtRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TxtRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TxtRecord.__pulumiType;
    }

    /**
     * The FQDN of the DNS TXT Record.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * The name of the DNS TXT Record. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * One or more `record` blocks as defined below.
     */
    public readonly records!: pulumi.Output<outputs.privatedns.TxtRecordRecord[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly ttl!: pulumi.Output<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly zoneName!: pulumi.Output<string>;

    /**
     * Create a TxtRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TxtRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TxtRecordArgs | TxtRecordState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TxtRecordState | undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["records"] = state ? state.records : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["zoneName"] = state ? state.zoneName : undefined;
        } else {
            const args = argsOrState as TxtRecordArgs | undefined;
            if (!args || args.records === undefined) {
                throw new Error("Missing required property 'records'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.ttl === undefined) {
                throw new Error("Missing required property 'ttl'");
            }
            if (!args || args.zoneName === undefined) {
                throw new Error("Missing required property 'zoneName'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["records"] = args ? args.records : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["zoneName"] = args ? args.zoneName : undefined;
            inputs["fqdn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TxtRecord.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TxtRecord resources.
 */
export interface TxtRecordState {
    /**
     * The FQDN of the DNS TXT Record.
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * The name of the DNS TXT Record. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more `record` blocks as defined below.
     */
    readonly records?: pulumi.Input<pulumi.Input<inputs.privatedns.TxtRecordRecord>[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly ttl?: pulumi.Input<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    readonly zoneName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TxtRecord resource.
 */
export interface TxtRecordArgs {
    /**
     * The name of the DNS TXT Record. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more `record` blocks as defined below.
     */
    readonly records: pulumi.Input<pulumi.Input<inputs.privatedns.TxtRecordRecord>[]>;
    /**
     * Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly ttl: pulumi.Input<number>;
    /**
     * Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
     */
    readonly zoneName: pulumi.Input<string>;
}