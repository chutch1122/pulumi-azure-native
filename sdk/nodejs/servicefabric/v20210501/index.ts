// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { ApplicationArgs } from "./application";
export type Application = import("./application").Application;
export const Application: typeof import("./application").Application = null as any;
utilities.lazyLoad(exports, ["Application"], () => require("./application"));

export { ApplicationTypeArgs } from "./applicationType";
export type ApplicationType = import("./applicationType").ApplicationType;
export const ApplicationType: typeof import("./applicationType").ApplicationType = null as any;
utilities.lazyLoad(exports, ["ApplicationType"], () => require("./applicationType"));

export { ApplicationTypeVersionArgs } from "./applicationTypeVersion";
export type ApplicationTypeVersion = import("./applicationTypeVersion").ApplicationTypeVersion;
export const ApplicationTypeVersion: typeof import("./applicationTypeVersion").ApplicationTypeVersion = null as any;
utilities.lazyLoad(exports, ["ApplicationTypeVersion"], () => require("./applicationTypeVersion"));

export { GetApplicationArgs, GetApplicationResult, GetApplicationOutputArgs } from "./getApplication";
export const getApplication: typeof import("./getApplication").getApplication = null as any;
export const getApplicationOutput: typeof import("./getApplication").getApplicationOutput = null as any;
utilities.lazyLoad(exports, ["getApplication","getApplicationOutput"], () => require("./getApplication"));

export { GetApplicationTypeArgs, GetApplicationTypeResult, GetApplicationTypeOutputArgs } from "./getApplicationType";
export const getApplicationType: typeof import("./getApplicationType").getApplicationType = null as any;
export const getApplicationTypeOutput: typeof import("./getApplicationType").getApplicationTypeOutput = null as any;
utilities.lazyLoad(exports, ["getApplicationType","getApplicationTypeOutput"], () => require("./getApplicationType"));

export { GetApplicationTypeVersionArgs, GetApplicationTypeVersionResult, GetApplicationTypeVersionOutputArgs } from "./getApplicationTypeVersion";
export const getApplicationTypeVersion: typeof import("./getApplicationTypeVersion").getApplicationTypeVersion = null as any;
export const getApplicationTypeVersionOutput: typeof import("./getApplicationTypeVersion").getApplicationTypeVersionOutput = null as any;
utilities.lazyLoad(exports, ["getApplicationTypeVersion","getApplicationTypeVersionOutput"], () => require("./getApplicationTypeVersion"));

export { GetManagedClusterArgs, GetManagedClusterResult, GetManagedClusterOutputArgs } from "./getManagedCluster";
export const getManagedCluster: typeof import("./getManagedCluster").getManagedCluster = null as any;
export const getManagedClusterOutput: typeof import("./getManagedCluster").getManagedClusterOutput = null as any;
utilities.lazyLoad(exports, ["getManagedCluster","getManagedClusterOutput"], () => require("./getManagedCluster"));

export { GetNodeTypeArgs, GetNodeTypeResult, GetNodeTypeOutputArgs } from "./getNodeType";
export const getNodeType: typeof import("./getNodeType").getNodeType = null as any;
export const getNodeTypeOutput: typeof import("./getNodeType").getNodeTypeOutput = null as any;
utilities.lazyLoad(exports, ["getNodeType","getNodeTypeOutput"], () => require("./getNodeType"));

export { GetServiceArgs, GetServiceResult, GetServiceOutputArgs } from "./getService";
export const getService: typeof import("./getService").getService = null as any;
export const getServiceOutput: typeof import("./getService").getServiceOutput = null as any;
utilities.lazyLoad(exports, ["getService","getServiceOutput"], () => require("./getService"));

export { ManagedClusterArgs } from "./managedCluster";
export type ManagedCluster = import("./managedCluster").ManagedCluster;
export const ManagedCluster: typeof import("./managedCluster").ManagedCluster = null as any;
utilities.lazyLoad(exports, ["ManagedCluster"], () => require("./managedCluster"));

export { NodeTypeArgs } from "./nodeType";
export type NodeType = import("./nodeType").NodeType;
export const NodeType: typeof import("./nodeType").NodeType = null as any;
utilities.lazyLoad(exports, ["NodeType"], () => require("./nodeType"));

export { ServiceArgs } from "./service";
export type Service = import("./service").Service;
export const Service: typeof import("./service").Service = null as any;
utilities.lazyLoad(exports, ["Service"], () => require("./service"));


// Export enums:
export * from "../../types/enums/servicefabric/v20210501";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:servicefabric/v20210501:Application":
                return new Application(name, <any>undefined, { urn })
            case "azure-native:servicefabric/v20210501:ApplicationType":
                return new ApplicationType(name, <any>undefined, { urn })
            case "azure-native:servicefabric/v20210501:ApplicationTypeVersion":
                return new ApplicationTypeVersion(name, <any>undefined, { urn })
            case "azure-native:servicefabric/v20210501:ManagedCluster":
                return new ManagedCluster(name, <any>undefined, { urn })
            case "azure-native:servicefabric/v20210501:NodeType":
                return new NodeType(name, <any>undefined, { urn })
            case "azure-native:servicefabric/v20210501:Service":
                return new Service(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "servicefabric/v20210501", _module)
