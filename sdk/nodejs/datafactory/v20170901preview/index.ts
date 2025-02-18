// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { DatasetArgs } from "./dataset";
export type Dataset = import("./dataset").Dataset;
export const Dataset: typeof import("./dataset").Dataset = null as any;
utilities.lazyLoad(exports, ["Dataset"], () => require("./dataset"));

export { FactoryArgs } from "./factory";
export type Factory = import("./factory").Factory;
export const Factory: typeof import("./factory").Factory = null as any;
utilities.lazyLoad(exports, ["Factory"], () => require("./factory"));

export { GetDatasetArgs, GetDatasetResult, GetDatasetOutputArgs } from "./getDataset";
export const getDataset: typeof import("./getDataset").getDataset = null as any;
export const getDatasetOutput: typeof import("./getDataset").getDatasetOutput = null as any;
utilities.lazyLoad(exports, ["getDataset","getDatasetOutput"], () => require("./getDataset"));

export { GetFactoryArgs, GetFactoryResult, GetFactoryOutputArgs } from "./getFactory";
export const getFactory: typeof import("./getFactory").getFactory = null as any;
export const getFactoryOutput: typeof import("./getFactory").getFactoryOutput = null as any;
utilities.lazyLoad(exports, ["getFactory","getFactoryOutput"], () => require("./getFactory"));

export { GetIntegrationRuntimeArgs, GetIntegrationRuntimeResult, GetIntegrationRuntimeOutputArgs } from "./getIntegrationRuntime";
export const getIntegrationRuntime: typeof import("./getIntegrationRuntime").getIntegrationRuntime = null as any;
export const getIntegrationRuntimeOutput: typeof import("./getIntegrationRuntime").getIntegrationRuntimeOutput = null as any;
utilities.lazyLoad(exports, ["getIntegrationRuntime","getIntegrationRuntimeOutput"], () => require("./getIntegrationRuntime"));

export { GetIntegrationRuntimeConnectionInfoArgs, GetIntegrationRuntimeConnectionInfoResult, GetIntegrationRuntimeConnectionInfoOutputArgs } from "./getIntegrationRuntimeConnectionInfo";
export const getIntegrationRuntimeConnectionInfo: typeof import("./getIntegrationRuntimeConnectionInfo").getIntegrationRuntimeConnectionInfo = null as any;
export const getIntegrationRuntimeConnectionInfoOutput: typeof import("./getIntegrationRuntimeConnectionInfo").getIntegrationRuntimeConnectionInfoOutput = null as any;
utilities.lazyLoad(exports, ["getIntegrationRuntimeConnectionInfo","getIntegrationRuntimeConnectionInfoOutput"], () => require("./getIntegrationRuntimeConnectionInfo"));

export { GetIntegrationRuntimeStatusArgs, GetIntegrationRuntimeStatusResult, GetIntegrationRuntimeStatusOutputArgs } from "./getIntegrationRuntimeStatus";
export const getIntegrationRuntimeStatus: typeof import("./getIntegrationRuntimeStatus").getIntegrationRuntimeStatus = null as any;
export const getIntegrationRuntimeStatusOutput: typeof import("./getIntegrationRuntimeStatus").getIntegrationRuntimeStatusOutput = null as any;
utilities.lazyLoad(exports, ["getIntegrationRuntimeStatus","getIntegrationRuntimeStatusOutput"], () => require("./getIntegrationRuntimeStatus"));

export { GetLinkedServiceArgs, GetLinkedServiceResult, GetLinkedServiceOutputArgs } from "./getLinkedService";
export const getLinkedService: typeof import("./getLinkedService").getLinkedService = null as any;
export const getLinkedServiceOutput: typeof import("./getLinkedService").getLinkedServiceOutput = null as any;
utilities.lazyLoad(exports, ["getLinkedService","getLinkedServiceOutput"], () => require("./getLinkedService"));

export { GetPipelineArgs, GetPipelineResult, GetPipelineOutputArgs } from "./getPipeline";
export const getPipeline: typeof import("./getPipeline").getPipeline = null as any;
export const getPipelineOutput: typeof import("./getPipeline").getPipelineOutput = null as any;
utilities.lazyLoad(exports, ["getPipeline","getPipelineOutput"], () => require("./getPipeline"));

export { GetTriggerArgs, GetTriggerResult, GetTriggerOutputArgs } from "./getTrigger";
export const getTrigger: typeof import("./getTrigger").getTrigger = null as any;
export const getTriggerOutput: typeof import("./getTrigger").getTriggerOutput = null as any;
utilities.lazyLoad(exports, ["getTrigger","getTriggerOutput"], () => require("./getTrigger"));

export { IntegrationRuntimeArgs } from "./integrationRuntime";
export type IntegrationRuntime = import("./integrationRuntime").IntegrationRuntime;
export const IntegrationRuntime: typeof import("./integrationRuntime").IntegrationRuntime = null as any;
utilities.lazyLoad(exports, ["IntegrationRuntime"], () => require("./integrationRuntime"));

export { LinkedServiceArgs } from "./linkedService";
export type LinkedService = import("./linkedService").LinkedService;
export const LinkedService: typeof import("./linkedService").LinkedService = null as any;
utilities.lazyLoad(exports, ["LinkedService"], () => require("./linkedService"));

export { ListIntegrationRuntimeAuthKeysArgs, ListIntegrationRuntimeAuthKeysResult, ListIntegrationRuntimeAuthKeysOutputArgs } from "./listIntegrationRuntimeAuthKeys";
export const listIntegrationRuntimeAuthKeys: typeof import("./listIntegrationRuntimeAuthKeys").listIntegrationRuntimeAuthKeys = null as any;
export const listIntegrationRuntimeAuthKeysOutput: typeof import("./listIntegrationRuntimeAuthKeys").listIntegrationRuntimeAuthKeysOutput = null as any;
utilities.lazyLoad(exports, ["listIntegrationRuntimeAuthKeys","listIntegrationRuntimeAuthKeysOutput"], () => require("./listIntegrationRuntimeAuthKeys"));

export { PipelineArgs } from "./pipeline";
export type Pipeline = import("./pipeline").Pipeline;
export const Pipeline: typeof import("./pipeline").Pipeline = null as any;
utilities.lazyLoad(exports, ["Pipeline"], () => require("./pipeline"));

export { TriggerArgs } from "./trigger";
export type Trigger = import("./trigger").Trigger;
export const Trigger: typeof import("./trigger").Trigger = null as any;
utilities.lazyLoad(exports, ["Trigger"], () => require("./trigger"));


// Export enums:
export * from "../../types/enums/datafactory/v20170901preview";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:datafactory/v20170901preview:Dataset":
                return new Dataset(name, <any>undefined, { urn })
            case "azure-native:datafactory/v20170901preview:Factory":
                return new Factory(name, <any>undefined, { urn })
            case "azure-native:datafactory/v20170901preview:IntegrationRuntime":
                return new IntegrationRuntime(name, <any>undefined, { urn })
            case "azure-native:datafactory/v20170901preview:LinkedService":
                return new LinkedService(name, <any>undefined, { urn })
            case "azure-native:datafactory/v20170901preview:Pipeline":
                return new Pipeline(name, <any>undefined, { urn })
            case "azure-native:datafactory/v20170901preview:Trigger":
                return new Trigger(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "datafactory/v20170901preview", _module)
