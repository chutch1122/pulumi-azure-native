// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetSubscriptionAliasArgs, GetSubscriptionAliasResult, GetSubscriptionAliasOutputArgs } from "./getSubscriptionAlias";
export const getSubscriptionAlias: typeof import("./getSubscriptionAlias").getSubscriptionAlias = null as any;
export const getSubscriptionAliasOutput: typeof import("./getSubscriptionAlias").getSubscriptionAliasOutput = null as any;
utilities.lazyLoad(exports, ["getSubscriptionAlias","getSubscriptionAliasOutput"], () => require("./getSubscriptionAlias"));

export { SubscriptionAliasArgs } from "./subscriptionAlias";
export type SubscriptionAlias = import("./subscriptionAlias").SubscriptionAlias;
export const SubscriptionAlias: typeof import("./subscriptionAlias").SubscriptionAlias = null as any;
utilities.lazyLoad(exports, ["SubscriptionAlias"], () => require("./subscriptionAlias"));


// Export enums:
export * from "../../types/enums/subscription/v20191001preview";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:subscription/v20191001preview:SubscriptionAlias":
                return new SubscriptionAlias(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "subscription/v20191001preview", _module)
