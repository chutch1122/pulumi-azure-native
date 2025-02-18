// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BudgetArgs } from "./budget";
export type Budget = import("./budget").Budget;
export const Budget: typeof import("./budget").Budget = null as any;
utilities.lazyLoad(exports, ["Budget"], () => require("./budget"));

export { GetBudgetArgs, GetBudgetResult, GetBudgetOutputArgs } from "./getBudget";
export const getBudget: typeof import("./getBudget").getBudget = null as any;
export const getBudgetOutput: typeof import("./getBudget").getBudgetOutput = null as any;
utilities.lazyLoad(exports, ["getBudget","getBudgetOutput"], () => require("./getBudget"));


// Export enums:
export * from "../types/enums/consumption";

// Export sub-modules:
import * as v20171230preview from "./v20171230preview";
import * as v20180131 from "./v20180131";
import * as v20180331 from "./v20180331";
import * as v20180630 from "./v20180630";
import * as v20180831 from "./v20180831";
import * as v20181001 from "./v20181001";
import * as v20190101 from "./v20190101";
import * as v20190401preview from "./v20190401preview";
import * as v20190501 from "./v20190501";
import * as v20190501preview from "./v20190501preview";
import * as v20190601 from "./v20190601";
import * as v20191001 from "./v20191001";
import * as v20191101 from "./v20191101";
import * as v20210501 from "./v20210501";
import * as v20211001 from "./v20211001";

export {
    v20171230preview,
    v20180131,
    v20180331,
    v20180630,
    v20180831,
    v20181001,
    v20190101,
    v20190401preview,
    v20190501,
    v20190501preview,
    v20190601,
    v20191001,
    v20191101,
    v20210501,
    v20211001,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:consumption:Budget":
                return new Budget(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "consumption", _module)
