// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetStorageAccountArgs, GetStorageAccountResult, GetStorageAccountOutputArgs } from "./getStorageAccount";
export const getStorageAccount: typeof import("./getStorageAccount").getStorageAccount = null as any;
export const getStorageAccountOutput: typeof import("./getStorageAccount").getStorageAccountOutput = null as any;
utilities.lazyLoad(exports, ["getStorageAccount","getStorageAccountOutput"], () => require("./getStorageAccount"));

export { ListStorageAccountKeysArgs, ListStorageAccountKeysResult, ListStorageAccountKeysOutputArgs } from "./listStorageAccountKeys";
export const listStorageAccountKeys: typeof import("./listStorageAccountKeys").listStorageAccountKeys = null as any;
export const listStorageAccountKeysOutput: typeof import("./listStorageAccountKeys").listStorageAccountKeysOutput = null as any;
utilities.lazyLoad(exports, ["listStorageAccountKeys","listStorageAccountKeysOutput"], () => require("./listStorageAccountKeys"));

export { ListStorageAccountSASArgs, ListStorageAccountSASResult, ListStorageAccountSASOutputArgs } from "./listStorageAccountSAS";
export const listStorageAccountSAS: typeof import("./listStorageAccountSAS").listStorageAccountSAS = null as any;
export const listStorageAccountSASOutput: typeof import("./listStorageAccountSAS").listStorageAccountSASOutput = null as any;
utilities.lazyLoad(exports, ["listStorageAccountSAS","listStorageAccountSASOutput"], () => require("./listStorageAccountSAS"));

export { ListStorageAccountServiceSASArgs, ListStorageAccountServiceSASResult, ListStorageAccountServiceSASOutputArgs } from "./listStorageAccountServiceSAS";
export const listStorageAccountServiceSAS: typeof import("./listStorageAccountServiceSAS").listStorageAccountServiceSAS = null as any;
export const listStorageAccountServiceSASOutput: typeof import("./listStorageAccountServiceSAS").listStorageAccountServiceSASOutput = null as any;
utilities.lazyLoad(exports, ["listStorageAccountServiceSAS","listStorageAccountServiceSASOutput"], () => require("./listStorageAccountServiceSAS"));

export { StorageAccountArgs } from "./storageAccount";
export type StorageAccount = import("./storageAccount").StorageAccount;
export const StorageAccount: typeof import("./storageAccount").StorageAccount = null as any;
utilities.lazyLoad(exports, ["StorageAccount"], () => require("./storageAccount"));


// Export enums:
export * from "../../types/enums/storage/v20170601";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:storage/v20170601:StorageAccount":
                return new StorageAccount(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "storage/v20170601", _module)
