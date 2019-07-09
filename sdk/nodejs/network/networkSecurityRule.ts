// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class NetworkSecurityRule extends pulumi.CustomResource {
    /**
     * Get an existing NetworkSecurityRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkSecurityRuleState, opts?: pulumi.CustomResourceOptions): NetworkSecurityRule {
        return new NetworkSecurityRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/networkSecurityRule:NetworkSecurityRule';

    /**
     * Returns true if the given object is an instance of NetworkSecurityRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkSecurityRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkSecurityRule.__pulumiType;
    }

    /**
     * Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
     */
    public readonly access!: pulumi.Output<string>;
    /**
     * A description for this rule. Restricted to 140 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `destination_address_prefixes` is not specified.
     */
    public readonly destinationAddressPrefix!: pulumi.Output<string | undefined>;
    /**
     * List of destination address prefixes. Tags may not be used. This is required if `destination_address_prefix` is not specified.
     */
    public readonly destinationAddressPrefixes!: pulumi.Output<string[] | undefined>;
    /**
     * A List of destination Application Security Group ID's
     */
    public readonly destinationApplicationSecurityGroupIds!: pulumi.Output<string | undefined>;
    /**
     * Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destination_port_ranges` is not specified.
     */
    public readonly destinationPortRange!: pulumi.Output<string | undefined>;
    /**
     * List of destination ports or port ranges. This is required if `destination_port_range` is not specified.
     */
    public readonly destinationPortRanges!: pulumi.Output<string[] | undefined>;
    /**
     * The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
     */
    public readonly networkSecurityGroupName!: pulumi.Output<string>;
    /**
     * Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Network protocol this rule applies to. Possible values include `Tcp`, `Udp` or `*` (which matches both).
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `source_address_prefixes` is not specified.
     */
    public readonly sourceAddressPrefix!: pulumi.Output<string | undefined>;
    /**
     * List of source address prefixes. Tags may not be used. This is required if `source_address_prefix` is not specified.
     */
    public readonly sourceAddressPrefixes!: pulumi.Output<string[] | undefined>;
    /**
     * A List of source Application Security Group ID's
     */
    public readonly sourceApplicationSecurityGroupIds!: pulumi.Output<string | undefined>;
    /**
     * Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `source_port_ranges` is not specified.
     */
    public readonly sourcePortRange!: pulumi.Output<string | undefined>;
    /**
     * List of source ports or port ranges. This is required if `source_port_range` is not specified.
     */
    public readonly sourcePortRanges!: pulumi.Output<string[] | undefined>;

    /**
     * Create a NetworkSecurityRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkSecurityRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkSecurityRuleArgs | NetworkSecurityRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkSecurityRuleState | undefined;
            inputs["access"] = state ? state.access : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["destinationAddressPrefix"] = state ? state.destinationAddressPrefix : undefined;
            inputs["destinationAddressPrefixes"] = state ? state.destinationAddressPrefixes : undefined;
            inputs["destinationApplicationSecurityGroupIds"] = state ? state.destinationApplicationSecurityGroupIds : undefined;
            inputs["destinationPortRange"] = state ? state.destinationPortRange : undefined;
            inputs["destinationPortRanges"] = state ? state.destinationPortRanges : undefined;
            inputs["direction"] = state ? state.direction : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkSecurityGroupName"] = state ? state.networkSecurityGroupName : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sourceAddressPrefix"] = state ? state.sourceAddressPrefix : undefined;
            inputs["sourceAddressPrefixes"] = state ? state.sourceAddressPrefixes : undefined;
            inputs["sourceApplicationSecurityGroupIds"] = state ? state.sourceApplicationSecurityGroupIds : undefined;
            inputs["sourcePortRange"] = state ? state.sourcePortRange : undefined;
            inputs["sourcePortRanges"] = state ? state.sourcePortRanges : undefined;
        } else {
            const args = argsOrState as NetworkSecurityRuleArgs | undefined;
            if (!args || args.access === undefined) {
                throw new Error("Missing required property 'access'");
            }
            if (!args || args.direction === undefined) {
                throw new Error("Missing required property 'direction'");
            }
            if (!args || args.networkSecurityGroupName === undefined) {
                throw new Error("Missing required property 'networkSecurityGroupName'");
            }
            if (!args || args.priority === undefined) {
                throw new Error("Missing required property 'priority'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["access"] = args ? args.access : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["destinationAddressPrefix"] = args ? args.destinationAddressPrefix : undefined;
            inputs["destinationAddressPrefixes"] = args ? args.destinationAddressPrefixes : undefined;
            inputs["destinationApplicationSecurityGroupIds"] = args ? args.destinationApplicationSecurityGroupIds : undefined;
            inputs["destinationPortRange"] = args ? args.destinationPortRange : undefined;
            inputs["destinationPortRanges"] = args ? args.destinationPortRanges : undefined;
            inputs["direction"] = args ? args.direction : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkSecurityGroupName"] = args ? args.networkSecurityGroupName : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sourceAddressPrefix"] = args ? args.sourceAddressPrefix : undefined;
            inputs["sourceAddressPrefixes"] = args ? args.sourceAddressPrefixes : undefined;
            inputs["sourceApplicationSecurityGroupIds"] = args ? args.sourceApplicationSecurityGroupIds : undefined;
            inputs["sourcePortRange"] = args ? args.sourcePortRange : undefined;
            inputs["sourcePortRanges"] = args ? args.sourcePortRanges : undefined;
        }
        super(NetworkSecurityRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkSecurityRule resources.
 */
export interface NetworkSecurityRuleState {
    /**
     * Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
     */
    readonly access?: pulumi.Input<string>;
    /**
     * A description for this rule. Restricted to 140 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `destination_address_prefixes` is not specified.
     */
    readonly destinationAddressPrefix?: pulumi.Input<string>;
    /**
     * List of destination address prefixes. Tags may not be used. This is required if `destination_address_prefix` is not specified.
     */
    readonly destinationAddressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A List of destination Application Security Group ID's
     */
    readonly destinationApplicationSecurityGroupIds?: pulumi.Input<string>;
    /**
     * Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destination_port_ranges` is not specified.
     */
    readonly destinationPortRange?: pulumi.Input<string>;
    /**
     * List of destination ports or port ranges. This is required if `destination_port_range` is not specified.
     */
    readonly destinationPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
     */
    readonly direction?: pulumi.Input<string>;
    /**
     * The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
     */
    readonly networkSecurityGroupName?: pulumi.Input<string>;
    /**
     * Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Network protocol this rule applies to. Possible values include `Tcp`, `Udp` or `*` (which matches both).
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `source_address_prefixes` is not specified.
     */
    readonly sourceAddressPrefix?: pulumi.Input<string>;
    /**
     * List of source address prefixes. Tags may not be used. This is required if `source_address_prefix` is not specified.
     */
    readonly sourceAddressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A List of source Application Security Group ID's
     */
    readonly sourceApplicationSecurityGroupIds?: pulumi.Input<string>;
    /**
     * Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `source_port_ranges` is not specified.
     */
    readonly sourcePortRange?: pulumi.Input<string>;
    /**
     * List of source ports or port ranges. This is required if `source_port_range` is not specified.
     */
    readonly sourcePortRanges?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a NetworkSecurityRule resource.
 */
export interface NetworkSecurityRuleArgs {
    /**
     * Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
     */
    readonly access: pulumi.Input<string>;
    /**
     * A description for this rule. Restricted to 140 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `destination_address_prefixes` is not specified.
     */
    readonly destinationAddressPrefix?: pulumi.Input<string>;
    /**
     * List of destination address prefixes. Tags may not be used. This is required if `destination_address_prefix` is not specified.
     */
    readonly destinationAddressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A List of destination Application Security Group ID's
     */
    readonly destinationApplicationSecurityGroupIds?: pulumi.Input<string>;
    /**
     * Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destination_port_ranges` is not specified.
     */
    readonly destinationPortRange?: pulumi.Input<string>;
    /**
     * List of destination ports or port ranges. This is required if `destination_port_range` is not specified.
     */
    readonly destinationPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
     */
    readonly direction: pulumi.Input<string>;
    /**
     * The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
     */
    readonly networkSecurityGroupName: pulumi.Input<string>;
    /**
     * Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
     */
    readonly priority: pulumi.Input<number>;
    /**
     * Network protocol this rule applies to. Possible values include `Tcp`, `Udp` or `*` (which matches both).
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `source_address_prefixes` is not specified.
     */
    readonly sourceAddressPrefix?: pulumi.Input<string>;
    /**
     * List of source address prefixes. Tags may not be used. This is required if `source_address_prefix` is not specified.
     */
    readonly sourceAddressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A List of source Application Security Group ID's
     */
    readonly sourceApplicationSecurityGroupIds?: pulumi.Input<string>;
    /**
     * Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `source_port_ranges` is not specified.
     */
    readonly sourcePortRange?: pulumi.Input<string>;
    /**
     * List of source ports or port ranges. This is required if `source_port_range` is not specified.
     */
    readonly sourcePortRanges?: pulumi.Input<pulumi.Input<string>[]>;
}
