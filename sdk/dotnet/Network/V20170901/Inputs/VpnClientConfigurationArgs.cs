// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20170901.Inputs
{

    /// <summary>
    /// VpnClientConfiguration for P2S client.
    /// </summary>
    public sealed class VpnClientConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The radius server address property of the VirtualNetworkGateway resource for vpn client connection.
        /// </summary>
        [Input("radiusServerAddress")]
        public Input<string>? RadiusServerAddress { get; set; }

        /// <summary>
        /// The radius secret property of the VirtualNetworkGateway resource for vpn client connection.
        /// </summary>
        [Input("radiusServerSecret")]
        public Input<string>? RadiusServerSecret { get; set; }

        /// <summary>
        /// The reference of the address space resource which represents Address space for P2S VpnClient.
        /// </summary>
        [Input("vpnClientAddressPool")]
        public Input<Inputs.AddressSpaceArgs>? VpnClientAddressPool { get; set; }

        [Input("vpnClientProtocols")]
        private InputList<Union<string, Pulumi.AzureNative.Network.V20170901.VpnClientProtocol>>? _vpnClientProtocols;

        /// <summary>
        /// VpnClientProtocols for Virtual network gateway.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Network.V20170901.VpnClientProtocol>> VpnClientProtocols
        {
            get => _vpnClientProtocols ?? (_vpnClientProtocols = new InputList<Union<string, Pulumi.AzureNative.Network.V20170901.VpnClientProtocol>>());
            set => _vpnClientProtocols = value;
        }

        [Input("vpnClientRevokedCertificates")]
        private InputList<Inputs.VpnClientRevokedCertificateArgs>? _vpnClientRevokedCertificates;

        /// <summary>
        /// VpnClientRevokedCertificate for Virtual network gateway.
        /// </summary>
        public InputList<Inputs.VpnClientRevokedCertificateArgs> VpnClientRevokedCertificates
        {
            get => _vpnClientRevokedCertificates ?? (_vpnClientRevokedCertificates = new InputList<Inputs.VpnClientRevokedCertificateArgs>());
            set => _vpnClientRevokedCertificates = value;
        }

        [Input("vpnClientRootCertificates")]
        private InputList<Inputs.VpnClientRootCertificateArgs>? _vpnClientRootCertificates;

        /// <summary>
        /// VpnClientRootCertificate for virtual network gateway.
        /// </summary>
        public InputList<Inputs.VpnClientRootCertificateArgs> VpnClientRootCertificates
        {
            get => _vpnClientRootCertificates ?? (_vpnClientRootCertificates = new InputList<Inputs.VpnClientRootCertificateArgs>());
            set => _vpnClientRootCertificates = value;
        }

        public VpnClientConfigurationArgs()
        {
        }
        public static new VpnClientConfigurationArgs Empty => new VpnClientConfigurationArgs();
    }
}
