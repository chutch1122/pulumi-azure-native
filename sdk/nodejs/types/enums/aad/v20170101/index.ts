// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ExternalAccess = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * A flag to determine whether or not Secure LDAP access over the internet is enabled or disabled.
 */
export type ExternalAccess = (typeof ExternalAccess)[keyof typeof ExternalAccess];

export const FilteredSync = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Enabled or Disabled flag to turn on Group-based filtered sync
 */
export type FilteredSync = (typeof FilteredSync)[keyof typeof FilteredSync];

export const Ldaps = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * A flag to determine whether or not Secure LDAP is enabled or disabled.
 */
export type Ldaps = (typeof Ldaps)[keyof typeof Ldaps];

export const NotifyDcAdmins = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Should domain controller admins be notified
 */
export type NotifyDcAdmins = (typeof NotifyDcAdmins)[keyof typeof NotifyDcAdmins];

export const NotifyGlobalAdmins = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Should global admins be notified
 */
export type NotifyGlobalAdmins = (typeof NotifyGlobalAdmins)[keyof typeof NotifyGlobalAdmins];

export const NtlmV1 = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * A flag to determine whether or not NtlmV1 is enabled or disabled.
 */
export type NtlmV1 = (typeof NtlmV1)[keyof typeof NtlmV1];

export const SyncNtlmPasswords = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * A flag to determine whether or not SyncNtlmPasswords is enabled or disabled.
 */
export type SyncNtlmPasswords = (typeof SyncNtlmPasswords)[keyof typeof SyncNtlmPasswords];

export const TlsV1 = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * A flag to determine whether or not TlsV1 is enabled or disabled.
 */
export type TlsV1 = (typeof TlsV1)[keyof typeof TlsV1];
