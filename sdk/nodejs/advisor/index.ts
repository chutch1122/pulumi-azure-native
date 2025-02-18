// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetSuppressionArgs, GetSuppressionResult, GetSuppressionOutputArgs } from "./getSuppression";
export const getSuppression: typeof import("./getSuppression").getSuppression = null as any;
export const getSuppressionOutput: typeof import("./getSuppression").getSuppressionOutput = null as any;
utilities.lazyLoad(exports, ["getSuppression","getSuppressionOutput"], () => require("./getSuppression"));

export { SuppressionArgs } from "./suppression";
export type Suppression = import("./suppression").Suppression;
export const Suppression: typeof import("./suppression").Suppression = null as any;
utilities.lazyLoad(exports, ["Suppression"], () => require("./suppression"));


// Export sub-modules:
import * as v20160712preview from "./v20160712preview";
import * as v20170331 from "./v20170331";
import * as v20170419 from "./v20170419";
import * as v20200101 from "./v20200101";
import * as v20220901 from "./v20220901";

export {
    v20160712preview,
    v20170331,
    v20170419,
    v20200101,
    v20220901,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:advisor:Suppression":
                return new Suppression(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "advisor", _module)
