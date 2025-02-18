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

export { GetPatchScheduleArgs, GetPatchScheduleResult, GetPatchScheduleOutputArgs } from "./getPatchSchedule";
export const getPatchSchedule: typeof import("./getPatchSchedule").getPatchSchedule = null as any;
export const getPatchScheduleOutput: typeof import("./getPatchSchedule").getPatchScheduleOutput = null as any;
utilities.lazyLoad(exports, ["getPatchSchedule","getPatchScheduleOutput"], () => require("./getPatchSchedule"));

export { GetRedisArgs, GetRedisResult, GetRedisOutputArgs } from "./getRedis";
export const getRedis: typeof import("./getRedis").getRedis = null as any;
export const getRedisOutput: typeof import("./getRedis").getRedisOutput = null as any;
utilities.lazyLoad(exports, ["getRedis","getRedisOutput"], () => require("./getRedis"));

export { GetRedisLinkedServerArgs, GetRedisLinkedServerResult, GetRedisLinkedServerOutputArgs } from "./getRedisLinkedServer";
export const getRedisLinkedServer: typeof import("./getRedisLinkedServer").getRedisLinkedServer = null as any;
export const getRedisLinkedServerOutput: typeof import("./getRedisLinkedServer").getRedisLinkedServerOutput = null as any;
utilities.lazyLoad(exports, ["getRedisLinkedServer","getRedisLinkedServerOutput"], () => require("./getRedisLinkedServer"));

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

export { RedisLinkedServerArgs } from "./redisLinkedServer";
export type RedisLinkedServer = import("./redisLinkedServer").RedisLinkedServer;
export const RedisLinkedServer: typeof import("./redisLinkedServer").RedisLinkedServer = null as any;
utilities.lazyLoad(exports, ["RedisLinkedServer"], () => require("./redisLinkedServer"));


// Export enums:
export * from "../../types/enums/cache/v20170201";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:cache/v20170201:FirewallRule":
                return new FirewallRule(name, <any>undefined, { urn })
            case "azure-native:cache/v20170201:PatchSchedule":
                return new PatchSchedule(name, <any>undefined, { urn })
            case "azure-native:cache/v20170201:Redis":
                return new Redis(name, <any>undefined, { urn })
            case "azure-native:cache/v20170201:RedisLinkedServer":
                return new RedisLinkedServer(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "cache/v20170201", _module)
