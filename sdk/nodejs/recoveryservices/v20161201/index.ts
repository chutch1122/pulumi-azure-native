// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetProtectionContainerArgs, GetProtectionContainerResult, GetProtectionContainerOutputArgs } from "./getProtectionContainer";
export const getProtectionContainer: typeof import("./getProtectionContainer").getProtectionContainer = null as any;
export const getProtectionContainerOutput: typeof import("./getProtectionContainer").getProtectionContainerOutput = null as any;
utilities.lazyLoad(exports, ["getProtectionContainer","getProtectionContainerOutput"], () => require("./getProtectionContainer"));

export { ProtectionContainerArgs } from "./protectionContainer";
export type ProtectionContainer = import("./protectionContainer").ProtectionContainer;
export const ProtectionContainer: typeof import("./protectionContainer").ProtectionContainer = null as any;
utilities.lazyLoad(exports, ["ProtectionContainer"], () => require("./protectionContainer"));


// Export enums:
export * from "../../types/enums/recoveryservices/v20161201";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:recoveryservices/v20161201:ProtectionContainer":
                return new ProtectionContainer(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "recoveryservices/v20161201", _module)
