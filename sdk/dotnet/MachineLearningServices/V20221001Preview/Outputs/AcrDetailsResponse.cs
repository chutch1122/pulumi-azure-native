// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20221001Preview.Outputs
{

    /// <summary>
    /// Details of ACR account to be used for the Registry
    /// </summary>
    [OutputType]
    public sealed class AcrDetailsResponse
    {
        public readonly Outputs.SystemCreatedAcrAccountResponse? SystemCreatedAcrAccount;
        public readonly Outputs.UserCreatedAcrAccountResponse? UserCreatedAcrAccount;

        [OutputConstructor]
        private AcrDetailsResponse(
            Outputs.SystemCreatedAcrAccountResponse? systemCreatedAcrAccount,

            Outputs.UserCreatedAcrAccountResponse? userCreatedAcrAccount)
        {
            SystemCreatedAcrAccount = systemCreatedAcrAccount;
            UserCreatedAcrAccount = userCreatedAcrAccount;
        }
    }
}
