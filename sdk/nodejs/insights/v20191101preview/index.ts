// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { DataCollectionRuleArgs } from "./dataCollectionRule";
export type DataCollectionRule = import("./dataCollectionRule").DataCollectionRule;
export const DataCollectionRule: typeof import("./dataCollectionRule").DataCollectionRule = null as any;
utilities.lazyLoad(exports, ["DataCollectionRule"], () => require("./dataCollectionRule"));

export { DataCollectionRuleAssociationArgs } from "./dataCollectionRuleAssociation";
export type DataCollectionRuleAssociation = import("./dataCollectionRuleAssociation").DataCollectionRuleAssociation;
export const DataCollectionRuleAssociation: typeof import("./dataCollectionRuleAssociation").DataCollectionRuleAssociation = null as any;
utilities.lazyLoad(exports, ["DataCollectionRuleAssociation"], () => require("./dataCollectionRuleAssociation"));

export { GetDataCollectionRuleArgs, GetDataCollectionRuleResult, GetDataCollectionRuleOutputArgs } from "./getDataCollectionRule";
export const getDataCollectionRule: typeof import("./getDataCollectionRule").getDataCollectionRule = null as any;
export const getDataCollectionRuleOutput: typeof import("./getDataCollectionRule").getDataCollectionRuleOutput = null as any;
utilities.lazyLoad(exports, ["getDataCollectionRule","getDataCollectionRuleOutput"], () => require("./getDataCollectionRule"));

export { GetDataCollectionRuleAssociationArgs, GetDataCollectionRuleAssociationResult, GetDataCollectionRuleAssociationOutputArgs } from "./getDataCollectionRuleAssociation";
export const getDataCollectionRuleAssociation: typeof import("./getDataCollectionRuleAssociation").getDataCollectionRuleAssociation = null as any;
export const getDataCollectionRuleAssociationOutput: typeof import("./getDataCollectionRuleAssociation").getDataCollectionRuleAssociationOutput = null as any;
utilities.lazyLoad(exports, ["getDataCollectionRuleAssociation","getDataCollectionRuleAssociationOutput"], () => require("./getDataCollectionRuleAssociation"));


// Export enums:
export * from "../../types/enums/insights/v20191101preview";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:insights/v20191101preview:DataCollectionRule":
                return new DataCollectionRule(name, <any>undefined, { urn })
            case "azure-native:insights/v20191101preview:DataCollectionRuleAssociation":
                return new DataCollectionRuleAssociation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "insights/v20191101preview", _module)
