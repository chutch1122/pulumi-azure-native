// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Description of a backup which will be performed
 */
export function listSiteBackupConfiguration(args: ListSiteBackupConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<ListSiteBackupConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure-native:web/v20150801:listSiteBackupConfiguration", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListSiteBackupConfigurationArgs {
    /**
     * Name of web app
     */
    name: string;
    /**
     * Name of resource group
     */
    resourceGroupName: string;
}

/**
 * Description of a backup which will be performed
 */
export interface ListSiteBackupConfigurationResult {
    /**
     * Schedule for the backup if it is executed periodically
     */
    readonly backupSchedule?: outputs.web.v20150801.BackupScheduleResponse;
    /**
     * Databases included in the backup
     */
    readonly databases?: outputs.web.v20150801.DatabaseBackupSettingResponse[];
    /**
     * True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled
     */
    readonly enabled?: boolean;
    /**
     * Resource Id
     */
    readonly id?: string;
    /**
     * Kind of resource
     */
    readonly kind?: string;
    /**
     * Resource Location
     */
    readonly location: string;
    /**
     * Resource Name
     */
    readonly name?: string;
    /**
     * SAS URL to the container
     */
    readonly storageAccountUrl?: string;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}

export function listSiteBackupConfigurationOutput(args: ListSiteBackupConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListSiteBackupConfigurationResult> {
    return pulumi.output(args).apply(a => listSiteBackupConfiguration(a, opts))
}

export interface ListSiteBackupConfigurationOutputArgs {
    /**
     * Name of web app
     */
    name: pulumi.Input<string>;
    /**
     * Name of resource group
     */
    resourceGroupName: pulumi.Input<string>;
}
