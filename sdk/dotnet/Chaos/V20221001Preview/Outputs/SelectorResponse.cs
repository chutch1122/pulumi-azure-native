// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Chaos.V20221001Preview.Outputs
{

    /// <summary>
    /// Model that represents a selector in the Experiment resource.
    /// </summary>
    [OutputType]
    public sealed class SelectorResponse
    {
        /// <summary>
        /// Model that represents available filter types that can be applied to a targets list.
        /// </summary>
        public readonly Outputs.SimpleFilterResponse? Filter;
        /// <summary>
        /// String of the selector ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of Target references.
        /// </summary>
        public readonly ImmutableArray<Outputs.TargetReferenceResponse> Targets;
        /// <summary>
        /// Enum of the selector type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SelectorResponse(
            Outputs.SimpleFilterResponse? filter,

            string id,

            ImmutableArray<Outputs.TargetReferenceResponse> targets,

            string type)
        {
            Filter = filter;
            Id = id;
            Targets = targets;
            Type = type;
        }
    }
}
