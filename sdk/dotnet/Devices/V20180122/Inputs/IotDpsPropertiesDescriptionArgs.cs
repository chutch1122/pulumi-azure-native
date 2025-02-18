// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Devices.V20180122.Inputs
{

    /// <summary>
    /// the service specific properties of a provisioning service, including keys, linked iot hubs, current state, and system generated properties such as hostname and idScope
    /// </summary>
    public sealed class IotDpsPropertiesDescriptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allocation policy to be used by this provisioning service.
        /// </summary>
        [Input("allocationPolicy")]
        public InputUnion<string, Pulumi.AzureNative.Devices.V20180122.AllocationPolicy>? AllocationPolicy { get; set; }

        [Input("authorizationPolicies")]
        private InputList<Inputs.SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs>? _authorizationPolicies;

        /// <summary>
        /// List of authorization keys for a provisioning service.
        /// </summary>
        public InputList<Inputs.SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs> AuthorizationPolicies
        {
            get => _authorizationPolicies ?? (_authorizationPolicies = new InputList<Inputs.SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs>());
            set => _authorizationPolicies = value;
        }

        [Input("iotHubs")]
        private InputList<Inputs.IotHubDefinitionDescriptionArgs>? _iotHubs;

        /// <summary>
        /// List of IoT hubs associated with this provisioning service.
        /// </summary>
        public InputList<Inputs.IotHubDefinitionDescriptionArgs> IotHubs
        {
            get => _iotHubs ?? (_iotHubs = new InputList<Inputs.IotHubDefinitionDescriptionArgs>());
            set => _iotHubs = value;
        }

        /// <summary>
        /// The ARM provisioning state of the provisioning service.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Current state of the provisioning service.
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNative.Devices.V20180122.State>? State { get; set; }

        public IotDpsPropertiesDescriptionArgs()
        {
        }
        public static new IotDpsPropertiesDescriptionArgs Empty => new IotDpsPropertiesDescriptionArgs();
    }
}
