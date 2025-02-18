// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as v20220401preview from "./v20220401preview";

export {
    v20220401preview,
};

export const ManifestType = {
    /**
     * Repositories using helm
     */
    Helm: "helm",
    /**
     * Repositories using kubernetes manifests
     */
    Kube: "kube",
} as const;

/**
 * Determines the type of manifests within the repository.
 */
export type ManifestType = (typeof ManifestType)[keyof typeof ManifestType];
