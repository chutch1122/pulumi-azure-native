// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CommunicationServiceArgs } from "./communicationService";
export type CommunicationService = import("./communicationService").CommunicationService;
export const CommunicationService: typeof import("./communicationService").CommunicationService = null as any;
utilities.lazyLoad(exports, ["CommunicationService"], () => require("./communicationService"));

export { DomainArgs } from "./domain";
export type Domain = import("./domain").Domain;
export const Domain: typeof import("./domain").Domain = null as any;
utilities.lazyLoad(exports, ["Domain"], () => require("./domain"));

export { EmailServiceArgs } from "./emailService";
export type EmailService = import("./emailService").EmailService;
export const EmailService: typeof import("./emailService").EmailService = null as any;
utilities.lazyLoad(exports, ["EmailService"], () => require("./emailService"));

export { GetCommunicationServiceArgs, GetCommunicationServiceResult, GetCommunicationServiceOutputArgs } from "./getCommunicationService";
export const getCommunicationService: typeof import("./getCommunicationService").getCommunicationService = null as any;
export const getCommunicationServiceOutput: typeof import("./getCommunicationService").getCommunicationServiceOutput = null as any;
utilities.lazyLoad(exports, ["getCommunicationService","getCommunicationServiceOutput"], () => require("./getCommunicationService"));

export { GetDomainArgs, GetDomainResult, GetDomainOutputArgs } from "./getDomain";
export const getDomain: typeof import("./getDomain").getDomain = null as any;
export const getDomainOutput: typeof import("./getDomain").getDomainOutput = null as any;
utilities.lazyLoad(exports, ["getDomain","getDomainOutput"], () => require("./getDomain"));

export { GetEmailServiceArgs, GetEmailServiceResult, GetEmailServiceOutputArgs } from "./getEmailService";
export const getEmailService: typeof import("./getEmailService").getEmailService = null as any;
export const getEmailServiceOutput: typeof import("./getEmailService").getEmailServiceOutput = null as any;
utilities.lazyLoad(exports, ["getEmailService","getEmailServiceOutput"], () => require("./getEmailService"));

export { ListCommunicationServiceKeysArgs, ListCommunicationServiceKeysResult, ListCommunicationServiceKeysOutputArgs } from "./listCommunicationServiceKeys";
export const listCommunicationServiceKeys: typeof import("./listCommunicationServiceKeys").listCommunicationServiceKeys = null as any;
export const listCommunicationServiceKeysOutput: typeof import("./listCommunicationServiceKeys").listCommunicationServiceKeysOutput = null as any;
utilities.lazyLoad(exports, ["listCommunicationServiceKeys","listCommunicationServiceKeysOutput"], () => require("./listCommunicationServiceKeys"));


// Export enums:
export * from "../types/enums/communication";

// Export sub-modules:
import * as v20200820 from "./v20200820";
import * as v20200820preview from "./v20200820preview";
import * as v20211001preview from "./v20211001preview";
import * as v20220701preview from "./v20220701preview";

export {
    v20200820,
    v20200820preview,
    v20211001preview,
    v20220701preview,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:communication:CommunicationService":
                return new CommunicationService(name, <any>undefined, { urn })
            case "azure-native:communication:Domain":
                return new Domain(name, <any>undefined, { urn })
            case "azure-native:communication:EmailService":
                return new EmailService(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "communication", _module)
