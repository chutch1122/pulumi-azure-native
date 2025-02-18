// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.DevHub
{
    /// <summary>
    /// Determines the type of manifests within the repository.
    /// </summary>
    [EnumType]
    public readonly struct ManifestType : IEquatable<ManifestType>
    {
        private readonly string _value;

        private ManifestType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Repositories using helm
        /// </summary>
        public static ManifestType Helm { get; } = new ManifestType("helm");
        /// <summary>
        /// Repositories using kubernetes manifests
        /// </summary>
        public static ManifestType Kube { get; } = new ManifestType("kube");

        public static bool operator ==(ManifestType left, ManifestType right) => left.Equals(right);
        public static bool operator !=(ManifestType left, ManifestType right) => !left.Equals(right);

        public static explicit operator string(ManifestType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ManifestType other && Equals(other);
        public bool Equals(ManifestType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
