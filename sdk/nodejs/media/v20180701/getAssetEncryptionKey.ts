// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Data needed to decrypt asset files encrypted with legacy storage encryption.
 */
export function getAssetEncryptionKey(args: GetAssetEncryptionKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetAssetEncryptionKeyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure-native:media/v20180701:getAssetEncryptionKey", {
        "accountName": args.accountName,
        "assetName": args.assetName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAssetEncryptionKeyArgs {
    /**
     * The Media Services account name.
     */
    accountName: string;
    /**
     * The Asset name.
     */
    assetName: string;
    /**
     * The name of the resource group within the Azure subscription.
     */
    resourceGroupName: string;
}

/**
 * Data needed to decrypt asset files encrypted with legacy storage encryption.
 */
export interface GetAssetEncryptionKeyResult {
    /**
     * Asset File encryption metadata.
     */
    readonly assetFileEncryptionMetadata?: outputs.media.v20180701.AssetFileEncryptionMetadataResponse[];
    /**
     * The Asset File storage encryption key.
     */
    readonly key?: string;
}

export function getAssetEncryptionKeyOutput(args: GetAssetEncryptionKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAssetEncryptionKeyResult> {
    return pulumi.output(args).apply(a => getAssetEncryptionKey(a, opts))
}

export interface GetAssetEncryptionKeyOutputArgs {
    /**
     * The Media Services account name.
     */
    accountName: pulumi.Input<string>;
    /**
     * The Asset name.
     */
    assetName: pulumi.Input<string>;
    /**
     * The name of the resource group within the Azure subscription.
     */
    resourceGroupName: pulumi.Input<string>;
}
