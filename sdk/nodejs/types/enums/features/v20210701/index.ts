// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const SubscriptionFeatureRegistrationState = {
    NotSpecified: "NotSpecified",
    NotRegistered: "NotRegistered",
    Pending: "Pending",
    Registering: "Registering",
    Registered: "Registered",
    Unregistering: "Unregistering",
    Unregistered: "Unregistered",
} as const;

/**
 * The state.
 */
export type SubscriptionFeatureRegistrationState = (typeof SubscriptionFeatureRegistrationState)[keyof typeof SubscriptionFeatureRegistrationState];
