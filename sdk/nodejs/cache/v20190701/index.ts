// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { FirewallRuleArgs } from "./firewallRule";
export type FirewallRule = import("./firewallRule").FirewallRule;
export const FirewallRule: typeof import("./firewallRule").FirewallRule = null as any;
utilities.lazyLoad(exports, ["FirewallRule"], () => require("./firewallRule"));

export { GetFirewallRuleArgs, GetFirewallRuleResult, GetFirewallRuleOutputArgs } from "./getFirewallRule";
export const getFirewallRule: typeof import("./getFirewallRule").getFirewallRule = null as any;
export const getFirewallRuleOutput: typeof import("./getFirewallRule").getFirewallRuleOutput = null as any;
utilities.lazyLoad(exports, ["getFirewallRule","getFirewallRuleOutput"], () => require("./getFirewallRule"));

export { GetLinkedServerArgs, GetLinkedServerResult, GetLinkedServerOutputArgs } from "./getLinkedServer";
export const getLinkedServer: typeof import("./getLinkedServer").getLinkedServer = null as any;
export const getLinkedServerOutput: typeof import("./getLinkedServer").getLinkedServerOutput = null as any;
utilities.lazyLoad(exports, ["getLinkedServer","getLinkedServerOutput"], () => require("./getLinkedServer"));

export { GetPatchScheduleArgs, GetPatchScheduleResult, GetPatchScheduleOutputArgs } from "./getPatchSchedule";
export const getPatchSchedule: typeof import("./getPatchSchedule").getPatchSchedule = null as any;
export const getPatchScheduleOutput: typeof import("./getPatchSchedule").getPatchScheduleOutput = null as any;
utilities.lazyLoad(exports, ["getPatchSchedule","getPatchScheduleOutput"], () => require("./getPatchSchedule"));

export { GetRedisArgs, GetRedisResult, GetRedisOutputArgs } from "./getRedis";
export const getRedis: typeof import("./getRedis").getRedis = null as any;
export const getRedisOutput: typeof import("./getRedis").getRedisOutput = null as any;
utilities.lazyLoad(exports, ["getRedis","getRedisOutput"], () => require("./getRedis"));

export { LinkedServerArgs } from "./linkedServer";
export type LinkedServer = import("./linkedServer").LinkedServer;
export const LinkedServer: typeof import("./linkedServer").LinkedServer = null as any;
utilities.lazyLoad(exports, ["LinkedServer"], () => require("./linkedServer"));

export { ListRedisKeysArgs, ListRedisKeysResult, ListRedisKeysOutputArgs } from "./listRedisKeys";
export const listRedisKeys: typeof import("./listRedisKeys").listRedisKeys = null as any;
export const listRedisKeysOutput: typeof import("./listRedisKeys").listRedisKeysOutput = null as any;
utilities.lazyLoad(exports, ["listRedisKeys","listRedisKeysOutput"], () => require("./listRedisKeys"));

export { PatchScheduleArgs } from "./patchSchedule";
export type PatchSchedule = import("./patchSchedule").PatchSchedule;
export const PatchSchedule: typeof import("./patchSchedule").PatchSchedule = null as any;
utilities.lazyLoad(exports, ["PatchSchedule"], () => require("./patchSchedule"));

export { RedisArgs } from "./redis";
export type Redis = import("./redis").Redis;
export const Redis: typeof import("./redis").Redis = null as any;
utilities.lazyLoad(exports, ["Redis"], () => require("./redis"));


// Export enums:
export * from "../../types/enums/cache/v20190701";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:cache/v20190701:FirewallRule":
                return new FirewallRule(name, <any>undefined, { urn })
            case "azure-native:cache/v20190701:LinkedServer":
                return new LinkedServer(name, <any>undefined, { urn })
            case "azure-native:cache/v20190701:PatchSchedule":
                return new PatchSchedule(name, <any>undefined, { urn })
            case "azure-native:cache/v20190701:Redis":
                return new Redis(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "cache/v20190701", _module)
