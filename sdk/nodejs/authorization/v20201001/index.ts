// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetRoleManagementPolicyAssignmentArgs, GetRoleManagementPolicyAssignmentResult, GetRoleManagementPolicyAssignmentOutputArgs } from "./getRoleManagementPolicyAssignment";
export const getRoleManagementPolicyAssignment: typeof import("./getRoleManagementPolicyAssignment").getRoleManagementPolicyAssignment = null as any;
export const getRoleManagementPolicyAssignmentOutput: typeof import("./getRoleManagementPolicyAssignment").getRoleManagementPolicyAssignmentOutput = null as any;
utilities.lazyLoad(exports, ["getRoleManagementPolicyAssignment","getRoleManagementPolicyAssignmentOutput"], () => require("./getRoleManagementPolicyAssignment"));

export { RoleManagementPolicyAssignmentArgs } from "./roleManagementPolicyAssignment";
export type RoleManagementPolicyAssignment = import("./roleManagementPolicyAssignment").RoleManagementPolicyAssignment;
export const RoleManagementPolicyAssignment: typeof import("./roleManagementPolicyAssignment").RoleManagementPolicyAssignment = null as any;
utilities.lazyLoad(exports, ["RoleManagementPolicyAssignment"], () => require("./roleManagementPolicyAssignment"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:authorization/v20201001:RoleManagementPolicyAssignment":
                return new RoleManagementPolicyAssignment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "authorization/v20201001", _module)
