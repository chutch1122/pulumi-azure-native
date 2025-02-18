// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const BindingType = {
    ApplicationInsights: "ApplicationInsights",
    ApacheSkyWalking: "ApacheSkyWalking",
    AppDynamics: "AppDynamics",
    Dynatrace: "Dynatrace",
    NewRelic: "NewRelic",
    ElasticAPM: "ElasticAPM",
} as const;

/**
 * Buildpack Binding Type
 */
export type BindingType = (typeof BindingType)[keyof typeof BindingType];

export const ManagedIdentityType = {
    None: "None",
    SystemAssigned: "SystemAssigned",
    UserAssigned: "UserAssigned",
    SystemAssigned_UserAssigned: "SystemAssigned,UserAssigned",
} as const;

/**
 * Type of the managed identity
 */
export type ManagedIdentityType = (typeof ManagedIdentityType)[keyof typeof ManagedIdentityType];
