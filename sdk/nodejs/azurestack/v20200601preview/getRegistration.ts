// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Registration information.
 */
export function getRegistration(args: GetRegistrationArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistrationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure-native:azurestack/v20200601preview:getRegistration", {
        "registrationName": args.registrationName,
        "resourceGroup": args.resourceGroup,
    }, opts);
}

export interface GetRegistrationArgs {
    /**
     * Name of the Azure Stack registration.
     */
    registrationName: string;
    /**
     * Name of the resource group.
     */
    resourceGroup: string;
}

/**
 * Registration information.
 */
export interface GetRegistrationResult {
    /**
     * Specifies the billing mode for the Azure Stack registration.
     */
    readonly billingModel?: string;
    /**
     * The identifier of the registered Azure Stack.
     */
    readonly cloudId?: string;
    /**
     * The entity tag used for optimistic concurrency when modifying the resource.
     */
    readonly etag?: string;
    /**
     * ID of the resource.
     */
    readonly id: string;
    /**
     * The kind of the resource.
     */
    readonly kind: string;
    /**
     * Location of the resource.
     */
    readonly location: string;
    /**
     * Name of the resource.
     */
    readonly name: string;
    /**
     * The object identifier associated with the Azure Stack connecting to Azure.
     */
    readonly objectId?: string;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.azurestack.v20200601preview.SystemDataResponse;
    /**
     * Custom tags for the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Type of Resource.
     */
    readonly type: string;
}

export function getRegistrationOutput(args: GetRegistrationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegistrationResult> {
    return pulumi.output(args).apply(a => getRegistration(a, opts))
}

export interface GetRegistrationOutputArgs {
    /**
     * Name of the Azure Stack registration.
     */
    registrationName: pulumi.Input<string>;
    /**
     * Name of the resource group.
     */
    resourceGroup: pulumi.Input<string>;
}
