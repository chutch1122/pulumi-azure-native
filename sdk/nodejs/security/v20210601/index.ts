// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { AssessmentArgs } from "./assessment";
export type Assessment = import("./assessment").Assessment;
export const Assessment: typeof import("./assessment").Assessment = null as any;
utilities.lazyLoad(exports, ["Assessment"], () => require("./assessment"));

export { AssessmentMetadataInSubscriptionArgs } from "./assessmentMetadataInSubscription";
export type AssessmentMetadataInSubscription = import("./assessmentMetadataInSubscription").AssessmentMetadataInSubscription;
export const AssessmentMetadataInSubscription: typeof import("./assessmentMetadataInSubscription").AssessmentMetadataInSubscription = null as any;
utilities.lazyLoad(exports, ["AssessmentMetadataInSubscription"], () => require("./assessmentMetadataInSubscription"));

export { GetAssessmentArgs, GetAssessmentResult, GetAssessmentOutputArgs } from "./getAssessment";
export const getAssessment: typeof import("./getAssessment").getAssessment = null as any;
export const getAssessmentOutput: typeof import("./getAssessment").getAssessmentOutput = null as any;
utilities.lazyLoad(exports, ["getAssessment","getAssessmentOutput"], () => require("./getAssessment"));

export { GetAssessmentMetadataInSubscriptionArgs, GetAssessmentMetadataInSubscriptionResult, GetAssessmentMetadataInSubscriptionOutputArgs } from "./getAssessmentMetadataInSubscription";
export const getAssessmentMetadataInSubscription: typeof import("./getAssessmentMetadataInSubscription").getAssessmentMetadataInSubscription = null as any;
export const getAssessmentMetadataInSubscriptionOutput: typeof import("./getAssessmentMetadataInSubscription").getAssessmentMetadataInSubscriptionOutput = null as any;
utilities.lazyLoad(exports, ["getAssessmentMetadataInSubscription","getAssessmentMetadataInSubscriptionOutput"], () => require("./getAssessmentMetadataInSubscription"));


// Export enums:
export * from "../../types/enums/security/v20210601";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:security/v20210601:Assessment":
                return new Assessment(name, <any>undefined, { urn })
            case "azure-native:security/v20210601:AssessmentMetadataInSubscription":
                return new AssessmentMetadataInSubscription(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "security/v20210601", _module)
