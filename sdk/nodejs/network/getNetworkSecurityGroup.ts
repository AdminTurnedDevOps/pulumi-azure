// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Network Security Group.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const test = pulumi.all([azurerm_network_security_group_test.name, azurerm_resource_group_test.name]).apply(([azurerm_network_security_group_testName, azurerm_resource_group_testName]) => azure.network.getNetworkSecurityGroup({
 *     name: azurerm_network_security_group_testName,
 *     resourceGroupName: azurerm_resource_group_testName,
 * }));
 * 
 * export const location = test.location;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/network_security_group.html.markdown.
 */
export function getNetworkSecurityGroup(args: GetNetworkSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkSecurityGroupResult> {
    return pulumi.runtime.invoke("azure:network/getNetworkSecurityGroup:getNetworkSecurityGroup", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworkSecurityGroup.
 */
export interface GetNetworkSecurityGroupArgs {
    /**
     * Specifies the Name of the Network Security Group.
     */
    readonly name: string;
    /**
     * Specifies the Name of the Resource Group within which the Network Security Group exists
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getNetworkSecurityGroup.
 */
export interface GetNetworkSecurityGroupResult {
    /**
     * The supported Azure location where the resource exists.
     */
    readonly location: string;
    /**
     * The name of the security rule.
     */
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * One or more `security_rule` blocks as defined below.
     */
    readonly securityRules: { access: string, description: string, destinationAddressPrefix: string, destinationAddressPrefixes: string[], destinationApplicationSecurityGroupIds?: string[], destinationPortRange: string, destinationPortRanges: string[], direction: string, name: string, priority: number, protocol: string, sourceAddressPrefix: string, sourceAddressPrefixes: string[], sourceApplicationSecurityGroupIds?: string[], sourcePortRange: string, sourcePortRanges: string[] }[];
    /**
     * A mapping of tags assigned to the resource.
     */
    readonly tags: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
