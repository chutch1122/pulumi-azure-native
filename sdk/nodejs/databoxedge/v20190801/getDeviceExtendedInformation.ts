// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The extended Info of the Data Box Edge/Gateway device.
 */
export function getDeviceExtendedInformation(args: GetDeviceExtendedInformationArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceExtendedInformationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure-native:databoxedge/v20190801:getDeviceExtendedInformation", {
        "deviceName": args.deviceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDeviceExtendedInformationArgs {
    /**
     * The device name.
     */
    deviceName: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
}

/**
 * The extended Info of the Data Box Edge/Gateway device.
 */
export interface GetDeviceExtendedInformationResult {
    /**
     * The public part of the encryption certificate. Client uses this to encrypt any secret.
     */
    readonly encryptionKey?: string;
    /**
     * The digital signature of encrypted certificate.
     */
    readonly encryptionKeyThumbprint?: string;
    /**
     * The path ID that uniquely identifies the object.
     */
    readonly id: string;
    /**
     * The object name.
     */
    readonly name: string;
    /**
     * The Resource ID of the Resource.
     */
    readonly resourceKey: string;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
}

export function getDeviceExtendedInformationOutput(args: GetDeviceExtendedInformationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeviceExtendedInformationResult> {
    return pulumi.output(args).apply(a => getDeviceExtendedInformation(a, opts))
}

export interface GetDeviceExtendedInformationOutputArgs {
    /**
     * The device name.
     */
    deviceName: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
}
