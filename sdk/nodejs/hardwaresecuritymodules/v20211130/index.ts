// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { DedicatedHsmArgs } from "./dedicatedHsm";
export type DedicatedHsm = import("./dedicatedHsm").DedicatedHsm;
export const DedicatedHsm: typeof import("./dedicatedHsm").DedicatedHsm = null as any;
utilities.lazyLoad(exports, ["DedicatedHsm"], () => require("./dedicatedHsm"));

export { GetDedicatedHsmArgs, GetDedicatedHsmResult, GetDedicatedHsmOutputArgs } from "./getDedicatedHsm";
export const getDedicatedHsm: typeof import("./getDedicatedHsm").getDedicatedHsm = null as any;
export const getDedicatedHsmOutput: typeof import("./getDedicatedHsm").getDedicatedHsmOutput = null as any;
utilities.lazyLoad(exports, ["getDedicatedHsm","getDedicatedHsmOutput"], () => require("./getDedicatedHsm"));


// Export enums:
export * from "../../types/enums/hardwaresecuritymodules/v20211130";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:hardwaresecuritymodules/v20211130:DedicatedHsm":
                return new DedicatedHsm(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "hardwaresecuritymodules/v20211130", _module)
