// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing Virtual Machine.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const test = pulumi.output(azure.compute.getVirtualMachine({
 *     name: "production",
 *     resourceGroupName: "networking",
 * }));
 * 
 * export const virtualMachineId = test.apply(test => test.id);
 * ```
 */
export function getVirtualMachine(args: GetVirtualMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineResult> {
    return pulumi.runtime.invoke("azure:compute/getVirtualMachine:getVirtualMachine", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualMachine.
 */
export interface GetVirtualMachineArgs {
    /**
     * Specifies the name of the Virtual Machine.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group the Virtual Machine is located in.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getVirtualMachine.
 */
export interface GetVirtualMachineResult {
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
