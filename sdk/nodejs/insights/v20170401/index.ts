// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { ActionGroupArgs } from "./actionGroup";
export type ActionGroup = import("./actionGroup").ActionGroup;
export const ActionGroup: typeof import("./actionGroup").ActionGroup = null as any;
utilities.lazyLoad(exports, ["ActionGroup"], () => require("./actionGroup"));

export { ActivityLogAlertArgs } from "./activityLogAlert";
export type ActivityLogAlert = import("./activityLogAlert").ActivityLogAlert;
export const ActivityLogAlert: typeof import("./activityLogAlert").ActivityLogAlert = null as any;
utilities.lazyLoad(exports, ["ActivityLogAlert"], () => require("./activityLogAlert"));

export { GetActionGroupArgs, GetActionGroupResult, GetActionGroupOutputArgs } from "./getActionGroup";
export const getActionGroup: typeof import("./getActionGroup").getActionGroup = null as any;
export const getActionGroupOutput: typeof import("./getActionGroup").getActionGroupOutput = null as any;
utilities.lazyLoad(exports, ["getActionGroup","getActionGroupOutput"], () => require("./getActionGroup"));

export { GetActivityLogAlertArgs, GetActivityLogAlertResult, GetActivityLogAlertOutputArgs } from "./getActivityLogAlert";
export const getActivityLogAlert: typeof import("./getActivityLogAlert").getActivityLogAlert = null as any;
export const getActivityLogAlertOutput: typeof import("./getActivityLogAlert").getActivityLogAlertOutput = null as any;
utilities.lazyLoad(exports, ["getActivityLogAlert","getActivityLogAlertOutput"], () => require("./getActivityLogAlert"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:insights/v20170401:ActionGroup":
                return new ActionGroup(name, <any>undefined, { urn })
            case "azure-native:insights/v20170401:ActivityLogAlert":
                return new ActivityLogAlert(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "insights/v20170401", _module)
