// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Compute node information related to a AmlCompute.
 */
export function listComputeNodes(args: ListComputeNodesArgs, opts?: pulumi.InvokeOptions): Promise<ListComputeNodesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure-native:machinelearningservices/v20210301preview:listComputeNodes", {
        "computeName": args.computeName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface ListComputeNodesArgs {
    /**
     * Name of the Azure Machine Learning compute.
     */
    computeName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Name of Azure Machine Learning workspace.
     */
    workspaceName: string;
}

/**
 * Compute node information related to a AmlCompute.
 */
export interface ListComputeNodesResult {
    /**
     * The type of compute
     * Expected value is 'AmlCompute'.
     */
    readonly computeType: "AmlCompute";
    /**
     * The continuation token.
     */
    readonly nextLink: string;
    /**
     * The collection of returned AmlCompute nodes details.
     */
    readonly nodes: outputs.machinelearningservices.v20210301preview.AmlComputeNodeInformationResponse[];
}

export function listComputeNodesOutput(args: ListComputeNodesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListComputeNodesResult> {
    return pulumi.output(args).apply(a => listComputeNodes(a, opts))
}

export interface ListComputeNodesOutputArgs {
    /**
     * Name of the Azure Machine Learning compute.
     */
    computeName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of Azure Machine Learning workspace.
     */
    workspaceName: pulumi.Input<string>;
}
