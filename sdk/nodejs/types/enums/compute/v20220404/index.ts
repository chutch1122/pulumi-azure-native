// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const CloudServiceSlotType = {
    Production: "Production",
    Staging: "Staging",
} as const;

/**
 * Slot type for the cloud service.
 * Possible values are <br /><br />**Production**<br /><br />**Staging**<br /><br />
 * If not specified, the default value is Production.
 */
export type CloudServiceSlotType = (typeof CloudServiceSlotType)[keyof typeof CloudServiceSlotType];

export const CloudServiceUpgradeMode = {
    Auto: "Auto",
    Manual: "Manual",
    Simultaneous: "Simultaneous",
} as const;

/**
 * Update mode for the cloud service. Role instances are allocated to update domains when the service is deployed. Updates can be initiated manually in each update domain or initiated automatically in all update domains.
 * Possible Values are <br /><br />**Auto**<br /><br />**Manual** <br /><br />**Simultaneous**<br /><br />
 * If not specified, the default value is Auto. If set to Manual, PUT UpdateDomain must be called to apply the update. If set to Auto, the update is automatically applied to each update domain in sequence.
 */
export type CloudServiceUpgradeMode = (typeof CloudServiceUpgradeMode)[keyof typeof CloudServiceUpgradeMode];
