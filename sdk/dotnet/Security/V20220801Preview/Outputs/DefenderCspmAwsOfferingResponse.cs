// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.V20220801Preview.Outputs
{

    /// <summary>
    /// The CSPM P1 for Aws offering
    /// </summary>
    [OutputType]
    public sealed class DefenderCspmAwsOfferingResponse
    {
        /// <summary>
        /// The offering description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The type of the security offering.
        /// Expected value is 'DefenderCspmAws'.
        /// </summary>
        public readonly string OfferingType;
        /// <summary>
        /// The Microsoft Defender for Server VM scanning configuration
        /// </summary>
        public readonly Outputs.DefenderCspmAwsOfferingResponseVmScanners? VmScanners;

        [OutputConstructor]
        private DefenderCspmAwsOfferingResponse(
            string description,

            string offeringType,

            Outputs.DefenderCspmAwsOfferingResponseVmScanners? vmScanners)
        {
            Description = description;
            OfferingType = offeringType;
            VmScanners = vmScanners;
        }
    }
}
