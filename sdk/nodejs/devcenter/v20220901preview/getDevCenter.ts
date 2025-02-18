// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Represents a devcenter resource.
 */
export function getDevCenter(args: GetDevCenterArgs, opts?: pulumi.InvokeOptions): Promise<GetDevCenterResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure-native:devcenter/v20220901preview:getDevCenter", {
        "devCenterName": args.devCenterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDevCenterArgs {
    /**
     * The name of the devcenter.
     */
    devCenterName: string;
    /**
     * Name of the resource group within the Azure subscription.
     */
    resourceGroupName: string;
}

/**
 * Represents a devcenter resource.
 */
export interface GetDevCenterResult {
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * Managed identity properties
     */
    readonly identity?: outputs.devcenter.v20220901preview.ManagedServiceIdentityResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.devcenter.v20220901preview.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}

export function getDevCenterOutput(args: GetDevCenterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDevCenterResult> {
    return pulumi.output(args).apply(a => getDevCenter(a, opts))
}

export interface GetDevCenterOutputArgs {
    /**
     * The name of the devcenter.
     */
    devCenterName: pulumi.Input<string>;
    /**
     * Name of the resource group within the Azure subscription.
     */
    resourceGroupName: pulumi.Input<string>;
}
