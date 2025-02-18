// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { AccessPolicyArgs } from "./accessPolicy";
export type AccessPolicy = import("./accessPolicy").AccessPolicy;
export const AccessPolicy: typeof import("./accessPolicy").AccessPolicy = null as any;
utilities.lazyLoad(exports, ["AccessPolicy"], () => require("./accessPolicy"));

export { EnvironmentArgs } from "./environment";
export type Environment = import("./environment").Environment;
export const Environment: typeof import("./environment").Environment = null as any;
utilities.lazyLoad(exports, ["Environment"], () => require("./environment"));

export { EventHubEventSourceArgs } from "./eventHubEventSource";
export type EventHubEventSource = import("./eventHubEventSource").EventHubEventSource;
export const EventHubEventSource: typeof import("./eventHubEventSource").EventHubEventSource = null as any;
utilities.lazyLoad(exports, ["EventHubEventSource"], () => require("./eventHubEventSource"));

export { EventSourceArgs } from "./eventSource";
export type EventSource = import("./eventSource").EventSource;
export const EventSource: typeof import("./eventSource").EventSource = null as any;
utilities.lazyLoad(exports, ["EventSource"], () => require("./eventSource"));

export { GetAccessPolicyArgs, GetAccessPolicyResult, GetAccessPolicyOutputArgs } from "./getAccessPolicy";
export const getAccessPolicy: typeof import("./getAccessPolicy").getAccessPolicy = null as any;
export const getAccessPolicyOutput: typeof import("./getAccessPolicy").getAccessPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getAccessPolicy","getAccessPolicyOutput"], () => require("./getAccessPolicy"));

export { GetEnvironmentArgs, GetEnvironmentResult, GetEnvironmentOutputArgs } from "./getEnvironment";
export const getEnvironment: typeof import("./getEnvironment").getEnvironment = null as any;
export const getEnvironmentOutput: typeof import("./getEnvironment").getEnvironmentOutput = null as any;
utilities.lazyLoad(exports, ["getEnvironment","getEnvironmentOutput"], () => require("./getEnvironment"));

export { GetEventHubEventSourceArgs, GetEventHubEventSourceResult, GetEventHubEventSourceOutputArgs } from "./getEventHubEventSource";
export const getEventHubEventSource: typeof import("./getEventHubEventSource").getEventHubEventSource = null as any;
export const getEventHubEventSourceOutput: typeof import("./getEventHubEventSource").getEventHubEventSourceOutput = null as any;
utilities.lazyLoad(exports, ["getEventHubEventSource","getEventHubEventSourceOutput"], () => require("./getEventHubEventSource"));

export { GetEventSourceArgs, GetEventSourceResult, GetEventSourceOutputArgs } from "./getEventSource";
export const getEventSource: typeof import("./getEventSource").getEventSource = null as any;
export const getEventSourceOutput: typeof import("./getEventSource").getEventSourceOutput = null as any;
utilities.lazyLoad(exports, ["getEventSource","getEventSourceOutput"], () => require("./getEventSource"));

export { GetIoTHubEventSourceArgs, GetIoTHubEventSourceResult, GetIoTHubEventSourceOutputArgs } from "./getIoTHubEventSource";
export const getIoTHubEventSource: typeof import("./getIoTHubEventSource").getIoTHubEventSource = null as any;
export const getIoTHubEventSourceOutput: typeof import("./getIoTHubEventSource").getIoTHubEventSourceOutput = null as any;
utilities.lazyLoad(exports, ["getIoTHubEventSource","getIoTHubEventSourceOutput"], () => require("./getIoTHubEventSource"));

export { GetLongTermEnvironmentArgs, GetLongTermEnvironmentResult, GetLongTermEnvironmentOutputArgs } from "./getLongTermEnvironment";
export const getLongTermEnvironment: typeof import("./getLongTermEnvironment").getLongTermEnvironment = null as any;
export const getLongTermEnvironmentOutput: typeof import("./getLongTermEnvironment").getLongTermEnvironmentOutput = null as any;
utilities.lazyLoad(exports, ["getLongTermEnvironment","getLongTermEnvironmentOutput"], () => require("./getLongTermEnvironment"));

export { GetReferenceDataSetArgs, GetReferenceDataSetResult, GetReferenceDataSetOutputArgs } from "./getReferenceDataSet";
export const getReferenceDataSet: typeof import("./getReferenceDataSet").getReferenceDataSet = null as any;
export const getReferenceDataSetOutput: typeof import("./getReferenceDataSet").getReferenceDataSetOutput = null as any;
utilities.lazyLoad(exports, ["getReferenceDataSet","getReferenceDataSetOutput"], () => require("./getReferenceDataSet"));

export { GetStandardEnvironmentArgs, GetStandardEnvironmentResult, GetStandardEnvironmentOutputArgs } from "./getStandardEnvironment";
export const getStandardEnvironment: typeof import("./getStandardEnvironment").getStandardEnvironment = null as any;
export const getStandardEnvironmentOutput: typeof import("./getStandardEnvironment").getStandardEnvironmentOutput = null as any;
utilities.lazyLoad(exports, ["getStandardEnvironment","getStandardEnvironmentOutput"], () => require("./getStandardEnvironment"));

export { IoTHubEventSourceArgs } from "./ioTHubEventSource";
export type IoTHubEventSource = import("./ioTHubEventSource").IoTHubEventSource;
export const IoTHubEventSource: typeof import("./ioTHubEventSource").IoTHubEventSource = null as any;
utilities.lazyLoad(exports, ["IoTHubEventSource"], () => require("./ioTHubEventSource"));

export { LongTermEnvironmentArgs } from "./longTermEnvironment";
export type LongTermEnvironment = import("./longTermEnvironment").LongTermEnvironment;
export const LongTermEnvironment: typeof import("./longTermEnvironment").LongTermEnvironment = null as any;
utilities.lazyLoad(exports, ["LongTermEnvironment"], () => require("./longTermEnvironment"));

export { ReferenceDataSetArgs } from "./referenceDataSet";
export type ReferenceDataSet = import("./referenceDataSet").ReferenceDataSet;
export const ReferenceDataSet: typeof import("./referenceDataSet").ReferenceDataSet = null as any;
utilities.lazyLoad(exports, ["ReferenceDataSet"], () => require("./referenceDataSet"));

export { StandardEnvironmentArgs } from "./standardEnvironment";
export type StandardEnvironment = import("./standardEnvironment").StandardEnvironment;
export const StandardEnvironment: typeof import("./standardEnvironment").StandardEnvironment = null as any;
utilities.lazyLoad(exports, ["StandardEnvironment"], () => require("./standardEnvironment"));


// Export enums:
export * from "../../types/enums/timeseriesinsights/v20180815preview";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:timeseriesinsights/v20180815preview:AccessPolicy":
                return new AccessPolicy(name, <any>undefined, { urn })
            case "azure-native:timeseriesinsights/v20180815preview:Environment":
                return new Environment(name, <any>undefined, { urn })
            case "azure-native:timeseriesinsights/v20180815preview:EventHubEventSource":
                return new EventHubEventSource(name, <any>undefined, { urn })
            case "azure-native:timeseriesinsights/v20180815preview:EventSource":
                return new EventSource(name, <any>undefined, { urn })
            case "azure-native:timeseriesinsights/v20180815preview:IoTHubEventSource":
                return new IoTHubEventSource(name, <any>undefined, { urn })
            case "azure-native:timeseriesinsights/v20180815preview:LongTermEnvironment":
                return new LongTermEnvironment(name, <any>undefined, { urn })
            case "azure-native:timeseriesinsights/v20180815preview:ReferenceDataSet":
                return new ReferenceDataSet(name, <any>undefined, { urn })
            case "azure-native:timeseriesinsights/v20180815preview:StandardEnvironment":
                return new StandardEnvironment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "timeseriesinsights/v20180815preview", _module)
