// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as v20160101 from "./v20160101";
import * as v20170601 from "./v20170601";
import * as v20200601preview from "./v20200601preview";
import * as v20220601 from "./v20220601";

export {
    v20160101,
    v20170601,
    v20200601preview,
    v20220601,
};

export const Location = {
    Global: "global",
} as const;

/**
 * Location of the resource.
 */
export type Location = (typeof Location)[keyof typeof Location];
