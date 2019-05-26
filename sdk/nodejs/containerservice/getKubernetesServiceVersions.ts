// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve the version of Kubernetes supported by Azure Kubernetes Service.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const current = pulumi.output(azure.containerservice.getKubernetesServiceVersions({
 *     location: "West Europe",
 * }));
 * 
 * export const latestVersion = current.latestVersion;
 * export const versions = current.versions;
 * ```
 */
export function getKubernetesServiceVersions(args: GetKubernetesServiceVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetKubernetesServiceVersionsResult> {
    return pulumi.runtime.invoke("azure:containerservice/getKubernetesServiceVersions:getKubernetesServiceVersions", {
        "location": args.location,
        "versionPrefix": args.versionPrefix,
    }, opts);
}

/**
 * A collection of arguments for invoking getKubernetesServiceVersions.
 */
export interface GetKubernetesServiceVersionsArgs {
    /**
     * Specifies the location in which to query for versions.
     */
    readonly location: string;
    /**
     * A prefix filter for the versions of Kubernetes which should be returned; for example `1.` will return `1.9` to `1.14`, whereas `1.12` will return `1.12.2`.
     */
    readonly versionPrefix?: string;
}

/**
 * A collection of values returned by getKubernetesServiceVersions.
 */
export interface GetKubernetesServiceVersionsResult {
    /**
     * The most recent version available.
     */
    readonly latestVersion: string;
    readonly location: string;
    readonly versionPrefix?: string;
    /**
     * The list of all supported versions.
     */
    readonly versions: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
