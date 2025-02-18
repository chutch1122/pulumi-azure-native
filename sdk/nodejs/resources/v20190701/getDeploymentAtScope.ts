// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Deployment information.
 */
export function getDeploymentAtScope(args: GetDeploymentAtScopeArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentAtScopeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure-native:resources/v20190701:getDeploymentAtScope", {
        "deploymentName": args.deploymentName,
        "scope": args.scope,
    }, opts);
}

export interface GetDeploymentAtScopeArgs {
    /**
     * The name of the deployment.
     */
    deploymentName: string;
    /**
     * The scope of a deployment.
     */
    scope: string;
}

/**
 * Deployment information.
 */
export interface GetDeploymentAtScopeResult {
    /**
     * The ID of the deployment.
     */
    readonly id: string;
    /**
     * the location of the deployment.
     */
    readonly location?: string;
    /**
     * The name of the deployment.
     */
    readonly name: string;
    /**
     * Deployment properties.
     */
    readonly properties: outputs.resources.v20190701.DeploymentPropertiesExtendedResponse;
    /**
     * The type of the deployment.
     */
    readonly type: string;
}

export function getDeploymentAtScopeOutput(args: GetDeploymentAtScopeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeploymentAtScopeResult> {
    return pulumi.output(args).apply(a => getDeploymentAtScope(a, opts))
}

export interface GetDeploymentAtScopeOutputArgs {
    /**
     * The name of the deployment.
     */
    deploymentName: pulumi.Input<string>;
    /**
     * The scope of a deployment.
     */
    scope: pulumi.Input<string>;
}
