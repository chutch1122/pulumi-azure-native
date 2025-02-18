// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { FluidRelayServerArgs } from "./fluidRelayServer";
export type FluidRelayServer = import("./fluidRelayServer").FluidRelayServer;
export const FluidRelayServer: typeof import("./fluidRelayServer").FluidRelayServer = null as any;
utilities.lazyLoad(exports, ["FluidRelayServer"], () => require("./fluidRelayServer"));

export { GetFluidRelayServerArgs, GetFluidRelayServerResult, GetFluidRelayServerOutputArgs } from "./getFluidRelayServer";
export const getFluidRelayServer: typeof import("./getFluidRelayServer").getFluidRelayServer = null as any;
export const getFluidRelayServerOutput: typeof import("./getFluidRelayServer").getFluidRelayServerOutput = null as any;
utilities.lazyLoad(exports, ["getFluidRelayServer","getFluidRelayServerOutput"], () => require("./getFluidRelayServer"));

export { GetFluidRelayServerKeysArgs, GetFluidRelayServerKeysResult, GetFluidRelayServerKeysOutputArgs } from "./getFluidRelayServerKeys";
export const getFluidRelayServerKeys: typeof import("./getFluidRelayServerKeys").getFluidRelayServerKeys = null as any;
export const getFluidRelayServerKeysOutput: typeof import("./getFluidRelayServerKeys").getFluidRelayServerKeysOutput = null as any;
utilities.lazyLoad(exports, ["getFluidRelayServerKeys","getFluidRelayServerKeysOutput"], () => require("./getFluidRelayServerKeys"));

export { ListFluidRelayServerKeysArgs, ListFluidRelayServerKeysResult, ListFluidRelayServerKeysOutputArgs } from "./listFluidRelayServerKeys";
export const listFluidRelayServerKeys: typeof import("./listFluidRelayServerKeys").listFluidRelayServerKeys = null as any;
export const listFluidRelayServerKeysOutput: typeof import("./listFluidRelayServerKeys").listFluidRelayServerKeysOutput = null as any;
utilities.lazyLoad(exports, ["listFluidRelayServerKeys","listFluidRelayServerKeysOutput"], () => require("./listFluidRelayServerKeys"));


// Export enums:
export * from "../types/enums/fluidrelay";

// Export sub-modules:
import * as v20210312preview from "./v20210312preview";
import * as v20210615preview from "./v20210615preview";
import * as v20210830preview from "./v20210830preview";
import * as v20210910preview from "./v20210910preview";
import * as v20220215 from "./v20220215";
import * as v20220421 from "./v20220421";
import * as v20220511 from "./v20220511";
import * as v20220526 from "./v20220526";
import * as v20220601 from "./v20220601";

export {
    v20210312preview,
    v20210615preview,
    v20210830preview,
    v20210910preview,
    v20220215,
    v20220421,
    v20220511,
    v20220526,
    v20220601,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:fluidrelay:FluidRelayServer":
                return new FluidRelayServer(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "fluidrelay", _module)
