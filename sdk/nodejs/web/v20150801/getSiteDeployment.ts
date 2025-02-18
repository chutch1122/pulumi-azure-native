// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Represents user credentials used for publishing activity
 */
/** @deprecated Version 2015-08-01 will be removed in v2 of the provider. */
export function getSiteDeployment(args: GetSiteDeploymentArgs, opts?: pulumi.InvokeOptions): Promise<GetSiteDeploymentResult> {
    pulumi.log.warn("getSiteDeployment is deprecated: Version 2015-08-01 will be removed in v2 of the provider.")
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure-native:web/v20150801:getSiteDeployment", {
        "id": args.id,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetSiteDeploymentArgs {
    /**
     * Id of the deployment
     */
    id: string;
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
 * Represents user credentials used for publishing activity
 */
export interface GetSiteDeploymentResult {
    /**
     * Active
     */
    readonly active?: boolean;
    /**
     * Author
     */
    readonly author?: string;
    /**
     * AuthorEmail
     */
    readonly authorEmail?: string;
    /**
     * Deployer
     */
    readonly deployer?: string;
    /**
     * Detail
     */
    readonly details?: string;
    /**
     * EndTime
     */
    readonly endTime?: string;
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
     * Message
     */
    readonly message?: string;
    /**
     * Resource Name
     */
    readonly name?: string;
    /**
     * StartTime
     */
    readonly startTime?: string;
    /**
     * Status
     */
    readonly status?: number;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type?: string;
}

export function getSiteDeploymentOutput(args: GetSiteDeploymentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSiteDeploymentResult> {
    return pulumi.output(args).apply(a => getSiteDeployment(a, opts))
}

export interface GetSiteDeploymentOutputArgs {
    /**
     * Id of the deployment
     */
    id: pulumi.Input<string>;
    /**
     * Name of web app
     */
    name: pulumi.Input<string>;
    /**
     * Name of resource group
     */
    resourceGroupName: pulumi.Input<string>;
}
