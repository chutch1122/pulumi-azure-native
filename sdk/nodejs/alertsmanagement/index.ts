// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ActionRuleByNameArgs } from "./actionRuleByName";
export type ActionRuleByName = import("./actionRuleByName").ActionRuleByName;
export const ActionRuleByName: typeof import("./actionRuleByName").ActionRuleByName = null as any;
utilities.lazyLoad(exports, ["ActionRuleByName"], () => require("./actionRuleByName"));

export { GetActionRuleByNameArgs, GetActionRuleByNameResult, GetActionRuleByNameOutputArgs } from "./getActionRuleByName";
export const getActionRuleByName: typeof import("./getActionRuleByName").getActionRuleByName = null as any;
export const getActionRuleByNameOutput: typeof import("./getActionRuleByName").getActionRuleByNameOutput = null as any;
utilities.lazyLoad(exports, ["getActionRuleByName","getActionRuleByNameOutput"], () => require("./getActionRuleByName"));

export { GetSmartDetectorAlertRuleArgs, GetSmartDetectorAlertRuleResult, GetSmartDetectorAlertRuleOutputArgs } from "./getSmartDetectorAlertRule";
export const getSmartDetectorAlertRule: typeof import("./getSmartDetectorAlertRule").getSmartDetectorAlertRule = null as any;
export const getSmartDetectorAlertRuleOutput: typeof import("./getSmartDetectorAlertRule").getSmartDetectorAlertRuleOutput = null as any;
utilities.lazyLoad(exports, ["getSmartDetectorAlertRule","getSmartDetectorAlertRuleOutput"], () => require("./getSmartDetectorAlertRule"));

export { SmartDetectorAlertRuleArgs } from "./smartDetectorAlertRule";
export type SmartDetectorAlertRule = import("./smartDetectorAlertRule").SmartDetectorAlertRule;
export const SmartDetectorAlertRule: typeof import("./smartDetectorAlertRule").SmartDetectorAlertRule = null as any;
utilities.lazyLoad(exports, ["SmartDetectorAlertRule"], () => require("./smartDetectorAlertRule"));


// Export enums:
export * from "../types/enums/alertsmanagement";

// Export sub-modules:
import * as v20181102privatepreview from "./v20181102privatepreview";
import * as v20190301 from "./v20190301";
import * as v20190505preview from "./v20190505preview";
import * as v20190601 from "./v20190601";
import * as v20210401 from "./v20210401";
import * as v20210808 from "./v20210808";
import * as v20210808preview from "./v20210808preview";

export {
    v20181102privatepreview,
    v20190301,
    v20190505preview,
    v20190601,
    v20210401,
    v20210808,
    v20210808preview,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:alertsmanagement:ActionRuleByName":
                return new ActionRuleByName(name, <any>undefined, { urn })
            case "azure-native:alertsmanagement:SmartDetectorAlertRule":
                return new SmartDetectorAlertRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "alertsmanagement", _module)
