// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetPolicyAssignmentArgs, GetPolicyAssignmentResult, GetPolicyAssignmentOutputArgs } from "./getPolicyAssignment";
export const getPolicyAssignment: typeof import("./getPolicyAssignment").getPolicyAssignment = null as any;
export const getPolicyAssignmentOutput: typeof import("./getPolicyAssignment").getPolicyAssignmentOutput = null as any;
utilities.lazyLoad(exports, ["getPolicyAssignment","getPolicyAssignmentOutput"], () => require("./getPolicyAssignment"));

export { GetPolicyDefinitionArgs, GetPolicyDefinitionResult, GetPolicyDefinitionOutputArgs } from "./getPolicyDefinition";
export const getPolicyDefinition: typeof import("./getPolicyDefinition").getPolicyDefinition = null as any;
export const getPolicyDefinitionOutput: typeof import("./getPolicyDefinition").getPolicyDefinitionOutput = null as any;
utilities.lazyLoad(exports, ["getPolicyDefinition","getPolicyDefinitionOutput"], () => require("./getPolicyDefinition"));

export { GetPolicyDefinitionAtManagementGroupArgs, GetPolicyDefinitionAtManagementGroupResult, GetPolicyDefinitionAtManagementGroupOutputArgs } from "./getPolicyDefinitionAtManagementGroup";
export const getPolicyDefinitionAtManagementGroup: typeof import("./getPolicyDefinitionAtManagementGroup").getPolicyDefinitionAtManagementGroup = null as any;
export const getPolicyDefinitionAtManagementGroupOutput: typeof import("./getPolicyDefinitionAtManagementGroup").getPolicyDefinitionAtManagementGroupOutput = null as any;
utilities.lazyLoad(exports, ["getPolicyDefinitionAtManagementGroup","getPolicyDefinitionAtManagementGroupOutput"], () => require("./getPolicyDefinitionAtManagementGroup"));

export { GetPolicySetDefinitionArgs, GetPolicySetDefinitionResult, GetPolicySetDefinitionOutputArgs } from "./getPolicySetDefinition";
export const getPolicySetDefinition: typeof import("./getPolicySetDefinition").getPolicySetDefinition = null as any;
export const getPolicySetDefinitionOutput: typeof import("./getPolicySetDefinition").getPolicySetDefinitionOutput = null as any;
utilities.lazyLoad(exports, ["getPolicySetDefinition","getPolicySetDefinitionOutput"], () => require("./getPolicySetDefinition"));

export { GetPolicySetDefinitionAtManagementGroupArgs, GetPolicySetDefinitionAtManagementGroupResult, GetPolicySetDefinitionAtManagementGroupOutputArgs } from "./getPolicySetDefinitionAtManagementGroup";
export const getPolicySetDefinitionAtManagementGroup: typeof import("./getPolicySetDefinitionAtManagementGroup").getPolicySetDefinitionAtManagementGroup = null as any;
export const getPolicySetDefinitionAtManagementGroupOutput: typeof import("./getPolicySetDefinitionAtManagementGroup").getPolicySetDefinitionAtManagementGroupOutput = null as any;
utilities.lazyLoad(exports, ["getPolicySetDefinitionAtManagementGroup","getPolicySetDefinitionAtManagementGroupOutput"], () => require("./getPolicySetDefinitionAtManagementGroup"));

export { PolicyAssignmentArgs } from "./policyAssignment";
export type PolicyAssignment = import("./policyAssignment").PolicyAssignment;
export const PolicyAssignment: typeof import("./policyAssignment").PolicyAssignment = null as any;
utilities.lazyLoad(exports, ["PolicyAssignment"], () => require("./policyAssignment"));

export { PolicyDefinitionArgs } from "./policyDefinition";
export type PolicyDefinition = import("./policyDefinition").PolicyDefinition;
export const PolicyDefinition: typeof import("./policyDefinition").PolicyDefinition = null as any;
utilities.lazyLoad(exports, ["PolicyDefinition"], () => require("./policyDefinition"));

export { PolicyDefinitionAtManagementGroupArgs } from "./policyDefinitionAtManagementGroup";
export type PolicyDefinitionAtManagementGroup = import("./policyDefinitionAtManagementGroup").PolicyDefinitionAtManagementGroup;
export const PolicyDefinitionAtManagementGroup: typeof import("./policyDefinitionAtManagementGroup").PolicyDefinitionAtManagementGroup = null as any;
utilities.lazyLoad(exports, ["PolicyDefinitionAtManagementGroup"], () => require("./policyDefinitionAtManagementGroup"));

export { PolicySetDefinitionArgs } from "./policySetDefinition";
export type PolicySetDefinition = import("./policySetDefinition").PolicySetDefinition;
export const PolicySetDefinition: typeof import("./policySetDefinition").PolicySetDefinition = null as any;
utilities.lazyLoad(exports, ["PolicySetDefinition"], () => require("./policySetDefinition"));

export { PolicySetDefinitionAtManagementGroupArgs } from "./policySetDefinitionAtManagementGroup";
export type PolicySetDefinitionAtManagementGroup = import("./policySetDefinitionAtManagementGroup").PolicySetDefinitionAtManagementGroup;
export const PolicySetDefinitionAtManagementGroup: typeof import("./policySetDefinitionAtManagementGroup").PolicySetDefinitionAtManagementGroup = null as any;
utilities.lazyLoad(exports, ["PolicySetDefinitionAtManagementGroup"], () => require("./policySetDefinitionAtManagementGroup"));


// Export enums:
export * from "../../types/enums/authorization/v20210601";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:authorization/v20210601:PolicyAssignment":
                return new PolicyAssignment(name, <any>undefined, { urn })
            case "azure-native:authorization/v20210601:PolicyDefinition":
                return new PolicyDefinition(name, <any>undefined, { urn })
            case "azure-native:authorization/v20210601:PolicyDefinitionAtManagementGroup":
                return new PolicyDefinitionAtManagementGroup(name, <any>undefined, { urn })
            case "azure-native:authorization/v20210601:PolicySetDefinition":
                return new PolicySetDefinition(name, <any>undefined, { urn })
            case "azure-native:authorization/v20210601:PolicySetDefinitionAtManagementGroup":
                return new PolicySetDefinitionAtManagementGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "authorization/v20210601", _module)
