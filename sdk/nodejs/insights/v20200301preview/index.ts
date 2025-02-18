// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { ComponentLinkedStorageAccountArgs } from "./componentLinkedStorageAccount";
export type ComponentLinkedStorageAccount = import("./componentLinkedStorageAccount").ComponentLinkedStorageAccount;
export const ComponentLinkedStorageAccount: typeof import("./componentLinkedStorageAccount").ComponentLinkedStorageAccount = null as any;
utilities.lazyLoad(exports, ["ComponentLinkedStorageAccount"], () => require("./componentLinkedStorageAccount"));

export { GetComponentLinkedStorageAccountArgs, GetComponentLinkedStorageAccountResult, GetComponentLinkedStorageAccountOutputArgs } from "./getComponentLinkedStorageAccount";
export const getComponentLinkedStorageAccount: typeof import("./getComponentLinkedStorageAccount").getComponentLinkedStorageAccount = null as any;
export const getComponentLinkedStorageAccountOutput: typeof import("./getComponentLinkedStorageAccount").getComponentLinkedStorageAccountOutput = null as any;
utilities.lazyLoad(exports, ["getComponentLinkedStorageAccount","getComponentLinkedStorageAccountOutput"], () => require("./getComponentLinkedStorageAccount"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:insights/v20200301preview:ComponentLinkedStorageAccount":
                return new ComponentLinkedStorageAccount(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "insights/v20200301preview", _module)
