// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetMyWorkbookArgs, GetMyWorkbookResult, GetMyWorkbookOutputArgs } from "./getMyWorkbook";
export const getMyWorkbook: typeof import("./getMyWorkbook").getMyWorkbook = null as any;
export const getMyWorkbookOutput: typeof import("./getMyWorkbook").getMyWorkbookOutput = null as any;
utilities.lazyLoad(exports, ["getMyWorkbook","getMyWorkbookOutput"], () => require("./getMyWorkbook"));

export { GetWorkbookArgs, GetWorkbookResult, GetWorkbookOutputArgs } from "./getWorkbook";
export const getWorkbook: typeof import("./getWorkbook").getWorkbook = null as any;
export const getWorkbookOutput: typeof import("./getWorkbook").getWorkbookOutput = null as any;
utilities.lazyLoad(exports, ["getWorkbook","getWorkbookOutput"], () => require("./getWorkbook"));

export { MyWorkbookArgs } from "./myWorkbook";
export type MyWorkbook = import("./myWorkbook").MyWorkbook;
export const MyWorkbook: typeof import("./myWorkbook").MyWorkbook = null as any;
utilities.lazyLoad(exports, ["MyWorkbook"], () => require("./myWorkbook"));

export { WorkbookArgs } from "./workbook";
export type Workbook = import("./workbook").Workbook;
export const Workbook: typeof import("./workbook").Workbook = null as any;
utilities.lazyLoad(exports, ["Workbook"], () => require("./workbook"));


// Export enums:
export * from "../../types/enums/insights/v20201020";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:insights/v20201020:MyWorkbook":
                return new MyWorkbook(name, <any>undefined, { urn })
            case "azure-native:insights/v20201020:Workbook":
                return new Workbook(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "insights/v20201020", _module)
