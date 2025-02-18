// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { DataExportArgs } from "./dataExport";
export type DataExport = import("./dataExport").DataExport;
export const DataExport: typeof import("./dataExport").DataExport = null as any;
utilities.lazyLoad(exports, ["DataExport"], () => require("./dataExport"));

export { GetDataExportArgs, GetDataExportResult, GetDataExportOutputArgs } from "./getDataExport";
export const getDataExport: typeof import("./getDataExport").getDataExport = null as any;
export const getDataExportOutput: typeof import("./getDataExport").getDataExportOutput = null as any;
utilities.lazyLoad(exports, ["getDataExport","getDataExportOutput"], () => require("./getDataExport"));

export { GetLinkedServiceArgs, GetLinkedServiceResult, GetLinkedServiceOutputArgs } from "./getLinkedService";
export const getLinkedService: typeof import("./getLinkedService").getLinkedService = null as any;
export const getLinkedServiceOutput: typeof import("./getLinkedService").getLinkedServiceOutput = null as any;
utilities.lazyLoad(exports, ["getLinkedService","getLinkedServiceOutput"], () => require("./getLinkedService"));

export { GetLinkedStorageAccountArgs, GetLinkedStorageAccountResult, GetLinkedStorageAccountOutputArgs } from "./getLinkedStorageAccount";
export const getLinkedStorageAccount: typeof import("./getLinkedStorageAccount").getLinkedStorageAccount = null as any;
export const getLinkedStorageAccountOutput: typeof import("./getLinkedStorageAccount").getLinkedStorageAccountOutput = null as any;
utilities.lazyLoad(exports, ["getLinkedStorageAccount","getLinkedStorageAccountOutput"], () => require("./getLinkedStorageAccount"));

export { LinkedServiceArgs } from "./linkedService";
export type LinkedService = import("./linkedService").LinkedService;
export const LinkedService: typeof import("./linkedService").LinkedService = null as any;
utilities.lazyLoad(exports, ["LinkedService"], () => require("./linkedService"));

export { LinkedStorageAccountArgs } from "./linkedStorageAccount";
export type LinkedStorageAccount = import("./linkedStorageAccount").LinkedStorageAccount;
export const LinkedStorageAccount: typeof import("./linkedStorageAccount").LinkedStorageAccount = null as any;
utilities.lazyLoad(exports, ["LinkedStorageAccount"], () => require("./linkedStorageAccount"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:operationalinsights/v20190801preview:DataExport":
                return new DataExport(name, <any>undefined, { urn })
            case "azure-native:operationalinsights/v20190801preview:LinkedService":
                return new LinkedService(name, <any>undefined, { urn })
            case "azure-native:operationalinsights/v20190801preview:LinkedStorageAccount":
                return new LinkedStorageAccount(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "operationalinsights/v20190801preview", _module)
