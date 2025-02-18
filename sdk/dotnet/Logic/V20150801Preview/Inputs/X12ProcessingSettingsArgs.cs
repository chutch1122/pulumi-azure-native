// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logic.V20150801Preview.Inputs
{

    public sealed class X12ProcessingSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value indicating whether to convert numerical type to implied decimal.
        /// </summary>
        [Input("convertImpliedDecimal")]
        public Input<bool>? ConvertImpliedDecimal { get; set; }

        /// <summary>
        /// The value indicating whether to create empty xml tags for trailing separators.
        /// </summary>
        [Input("createEmptyXmlTagsForTrailingSeparators")]
        public Input<bool>? CreateEmptyXmlTagsForTrailingSeparators { get; set; }

        /// <summary>
        /// The value indicating whether to mask security information.
        /// </summary>
        [Input("maskSecurityInfo")]
        public Input<bool>? MaskSecurityInfo { get; set; }

        /// <summary>
        /// The value indicating whether to preserve interchange.
        /// </summary>
        [Input("preserveInterchange")]
        public Input<bool>? PreserveInterchange { get; set; }

        /// <summary>
        /// The value indicating whether to suspend interchange on error.
        /// </summary>
        [Input("suspendInterchangeOnError")]
        public Input<bool>? SuspendInterchangeOnError { get; set; }

        /// <summary>
        /// The value indicating whether to use dot as decimal separator.
        /// </summary>
        [Input("useDotAsDecimalSeparator")]
        public Input<bool>? UseDotAsDecimalSeparator { get; set; }

        public X12ProcessingSettingsArgs()
        {
        }
        public static new X12ProcessingSettingsArgs Empty => new X12ProcessingSettingsArgs();
    }
}
