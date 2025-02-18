// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetManagementGroupDiagnosticSettingArgs, GetManagementGroupDiagnosticSettingResult, GetManagementGroupDiagnosticSettingOutputArgs } from "./getManagementGroupDiagnosticSetting";
export const getManagementGroupDiagnosticSetting: typeof import("./getManagementGroupDiagnosticSetting").getManagementGroupDiagnosticSetting = null as any;
export const getManagementGroupDiagnosticSettingOutput: typeof import("./getManagementGroupDiagnosticSetting").getManagementGroupDiagnosticSettingOutput = null as any;
utilities.lazyLoad(exports, ["getManagementGroupDiagnosticSetting","getManagementGroupDiagnosticSettingOutput"], () => require("./getManagementGroupDiagnosticSetting"));

export { ManagementGroupDiagnosticSettingArgs } from "./managementGroupDiagnosticSetting";
export type ManagementGroupDiagnosticSetting = import("./managementGroupDiagnosticSetting").ManagementGroupDiagnosticSetting;
export const ManagementGroupDiagnosticSetting: typeof import("./managementGroupDiagnosticSetting").ManagementGroupDiagnosticSetting = null as any;
utilities.lazyLoad(exports, ["ManagementGroupDiagnosticSetting"], () => require("./managementGroupDiagnosticSetting"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:insights/v20200101preview:ManagementGroupDiagnosticSetting":
                return new ManagementGroupDiagnosticSetting(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "insights/v20200101preview", _module)
