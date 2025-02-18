// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The iSCSI disk.
 */
/** @deprecated Version 2016-10-01 will be removed in v2 of the provider. */
export function getIscsiDisk(args: GetIscsiDiskArgs, opts?: pulumi.InvokeOptions): Promise<GetIscsiDiskResult> {
    pulumi.log.warn("getIscsiDisk is deprecated: Version 2016-10-01 will be removed in v2 of the provider.")
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure-native:storsimple/v20161001:getIscsiDisk", {
        "deviceName": args.deviceName,
        "diskName": args.diskName,
        "iscsiServerName": args.iscsiServerName,
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetIscsiDiskArgs {
    /**
     * The device name.
     */
    deviceName: string;
    /**
     * The disk name.
     */
    diskName: string;
    /**
     * The iSCSI server name.
     */
    iscsiServerName: string;
    /**
     * The manager name
     */
    managerName: string;
    /**
     * The resource group name
     */
    resourceGroupName: string;
}

/**
 * The iSCSI disk.
 */
export interface GetIscsiDiskResult {
    /**
     * The access control records.
     */
    readonly accessControlRecords: string[];
    /**
     * The data policy.
     */
    readonly dataPolicy: string;
    /**
     * The description.
     */
    readonly description?: string;
    /**
     * The disk status.
     */
    readonly diskStatus: string;
    /**
     * The identifier.
     */
    readonly id: string;
    /**
     * The local used capacity in bytes.
     */
    readonly localUsedCapacityInBytes: number;
    /**
     * The monitoring.
     */
    readonly monitoringStatus: string;
    /**
     * The name.
     */
    readonly name: string;
    /**
     * The provisioned capacity in bytes.
     */
    readonly provisionedCapacityInBytes: number;
    /**
     * The type.
     */
    readonly type: string;
    /**
     * The used capacity in bytes.
     */
    readonly usedCapacityInBytes: number;
}

export function getIscsiDiskOutput(args: GetIscsiDiskOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIscsiDiskResult> {
    return pulumi.output(args).apply(a => getIscsiDisk(a, opts))
}

export interface GetIscsiDiskOutputArgs {
    /**
     * The device name.
     */
    deviceName: pulumi.Input<string>;
    /**
     * The disk name.
     */
    diskName: pulumi.Input<string>;
    /**
     * The iSCSI server name.
     */
    iscsiServerName: pulumi.Input<string>;
    /**
     * The manager name
     */
    managerName: pulumi.Input<string>;
    /**
     * The resource group name
     */
    resourceGroupName: pulumi.Input<string>;
}
