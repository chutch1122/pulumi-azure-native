// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorSimple.V20161001
{
    public static class GetManagerEncryptionKey
    {
        /// <summary>
        /// This class can be used as the Type for any secret entity represented as Value, ValueCertificateThumbprint, EncryptionAlgorithm. In this case, "Value" is a secret and the "valueThumbprint" represents the certificate thumbprint of the value. The algorithm field is mainly for future usage to potentially allow different entities encrypted using different algorithms.
        /// </summary>
        public static Task<GetManagerEncryptionKeyResult> InvokeAsync(GetManagerEncryptionKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetManagerEncryptionKeyResult>("azure-native:storsimple/v20161001:getManagerEncryptionKey", args ?? new GetManagerEncryptionKeyArgs(), options.WithDefaults());

        /// <summary>
        /// This class can be used as the Type for any secret entity represented as Value, ValueCertificateThumbprint, EncryptionAlgorithm. In this case, "Value" is a secret and the "valueThumbprint" represents the certificate thumbprint of the value. The algorithm field is mainly for future usage to potentially allow different entities encrypted using different algorithms.
        /// </summary>
        public static Output<GetManagerEncryptionKeyResult> Invoke(GetManagerEncryptionKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagerEncryptionKeyResult>("azure-native:storsimple/v20161001:getManagerEncryptionKey", args ?? new GetManagerEncryptionKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetManagerEncryptionKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName", required: true)]
        public string ManagerName { get; set; } = null!;

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetManagerEncryptionKeyArgs()
        {
        }
        public static new GetManagerEncryptionKeyArgs Empty => new GetManagerEncryptionKeyArgs();
    }

    public sealed class GetManagerEncryptionKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName", required: true)]
        public Input<string> ManagerName { get; set; } = null!;

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetManagerEncryptionKeyInvokeArgs()
        {
        }
        public static new GetManagerEncryptionKeyInvokeArgs Empty => new GetManagerEncryptionKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetManagerEncryptionKeyResult
    {
        /// <summary>
        /// Algorithm used to encrypt "Value"
        /// </summary>
        public readonly string EncryptionAlgorithm;
        /// <summary>
        /// The value of the secret itself. If the secret is in plaintext or null then EncryptionAlgorithm will be none
        /// </summary>
        public readonly string Value;
        /// <summary>
        /// Thumbprint cert that was used to encrypt "Value"
        /// </summary>
        public readonly string? ValueCertificateThumbprint;

        [OutputConstructor]
        private GetManagerEncryptionKeyResult(
            string encryptionAlgorithm,

            string value,

            string? valueCertificateThumbprint)
        {
            EncryptionAlgorithm = encryptionAlgorithm;
            Value = value;
            ValueCertificateThumbprint = valueCertificateThumbprint;
        }
    }
}
