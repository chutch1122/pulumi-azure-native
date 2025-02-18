# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'ContainerServiceAgentPoolProfileResponse',
    'ContainerServiceLinuxProfileResponse',
    'ContainerServiceServicePrincipalProfileResponse',
    'ContainerServiceSshConfigurationResponse',
    'ContainerServiceSshPublicKeyResponse',
    'KeyVaultSecretRefResponse',
]

@pulumi.output_type
class ContainerServiceAgentPoolProfileResponse(dict):
    """
    Profile for the container service agent pool.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "vmSize":
            suggest = "vm_size"
        elif key == "dnsPrefix":
            suggest = "dns_prefix"
        elif key == "osDiskSizeGB":
            suggest = "os_disk_size_gb"
        elif key == "osType":
            suggest = "os_type"
        elif key == "storageProfile":
            suggest = "storage_profile"
        elif key == "vnetSubnetID":
            suggest = "vnet_subnet_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerServiceAgentPoolProfileResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerServiceAgentPoolProfileResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerServiceAgentPoolProfileResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 fqdn: str,
                 name: str,
                 vm_size: str,
                 count: Optional[int] = None,
                 dns_prefix: Optional[str] = None,
                 os_disk_size_gb: Optional[int] = None,
                 os_type: Optional[str] = None,
                 ports: Optional[Sequence[int]] = None,
                 storage_profile: Optional[str] = None,
                 vnet_subnet_id: Optional[str] = None):
        """
        Profile for the container service agent pool.
        :param str fqdn: FQDN for the agent pool.
        :param str name: Unique name of the agent pool profile in the context of the subscription and resource group.
        :param str vm_size: Size of agent VMs.
        :param int count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
        :param str dns_prefix: DNS prefix to be used to create the FQDN for the agent pool.
        :param int os_disk_size_gb: OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        :param str os_type: OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        :param Sequence[int] ports: Ports number array used to expose on this agent pool. The default opened ports are different based on your choice of orchestrator.
        :param str storage_profile: Storage profile specifies what kind of storage used. Choose from StorageAccount and ManagedDisks. Leave it empty, we will choose for you based on the orchestrator choice.
        :param str vnet_subnet_id: VNet SubnetID specifies the VNet's subnet identifier.
        """
        pulumi.set(__self__, "fqdn", fqdn)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "vm_size", vm_size)
        if count is None:
            count = 1
        if count is not None:
            pulumi.set(__self__, "count", count)
        if dns_prefix is not None:
            pulumi.set(__self__, "dns_prefix", dns_prefix)
        if os_disk_size_gb is not None:
            pulumi.set(__self__, "os_disk_size_gb", os_disk_size_gb)
        if os_type is not None:
            pulumi.set(__self__, "os_type", os_type)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if storage_profile is not None:
            pulumi.set(__self__, "storage_profile", storage_profile)
        if vnet_subnet_id is not None:
            pulumi.set(__self__, "vnet_subnet_id", vnet_subnet_id)

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        FQDN for the agent pool.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique name of the agent pool profile in the context of the subscription and resource group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> str:
        """
        Size of agent VMs.
        """
        return pulumi.get(self, "vm_size")

    @property
    @pulumi.getter
    def count(self) -> Optional[int]:
        """
        Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="dnsPrefix")
    def dns_prefix(self) -> Optional[str]:
        """
        DNS prefix to be used to create the FQDN for the agent pool.
        """
        return pulumi.get(self, "dns_prefix")

    @property
    @pulumi.getter(name="osDiskSizeGB")
    def os_disk_size_gb(self) -> Optional[int]:
        """
        OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        """
        return pulumi.get(self, "os_disk_size_gb")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> Optional[str]:
        """
        OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter
    def ports(self) -> Optional[Sequence[int]]:
        """
        Ports number array used to expose on this agent pool. The default opened ports are different based on your choice of orchestrator.
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> Optional[str]:
        """
        Storage profile specifies what kind of storage used. Choose from StorageAccount and ManagedDisks. Leave it empty, we will choose for you based on the orchestrator choice.
        """
        return pulumi.get(self, "storage_profile")

    @property
    @pulumi.getter(name="vnetSubnetID")
    def vnet_subnet_id(self) -> Optional[str]:
        """
        VNet SubnetID specifies the VNet's subnet identifier.
        """
        return pulumi.get(self, "vnet_subnet_id")


@pulumi.output_type
class ContainerServiceLinuxProfileResponse(dict):
    """
    Profile for Linux VMs in the container service cluster.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "adminUsername":
            suggest = "admin_username"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerServiceLinuxProfileResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerServiceLinuxProfileResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerServiceLinuxProfileResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 admin_username: str,
                 ssh: 'outputs.ContainerServiceSshConfigurationResponse'):
        """
        Profile for Linux VMs in the container service cluster.
        :param str admin_username: The administrator username to use for Linux VMs.
        :param 'ContainerServiceSshConfigurationResponse' ssh: SSH configuration for Linux-based VMs running on Azure.
        """
        pulumi.set(__self__, "admin_username", admin_username)
        pulumi.set(__self__, "ssh", ssh)

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> str:
        """
        The administrator username to use for Linux VMs.
        """
        return pulumi.get(self, "admin_username")

    @property
    @pulumi.getter
    def ssh(self) -> 'outputs.ContainerServiceSshConfigurationResponse':
        """
        SSH configuration for Linux-based VMs running on Azure.
        """
        return pulumi.get(self, "ssh")


@pulumi.output_type
class ContainerServiceServicePrincipalProfileResponse(dict):
    """
    Information about a service principal identity for the cluster to use for manipulating Azure APIs. Either secret or keyVaultSecretRef must be specified.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientId":
            suggest = "client_id"
        elif key == "keyVaultSecretRef":
            suggest = "key_vault_secret_ref"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerServiceServicePrincipalProfileResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerServiceServicePrincipalProfileResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerServiceServicePrincipalProfileResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_id: str,
                 key_vault_secret_ref: Optional['outputs.KeyVaultSecretRefResponse'] = None,
                 secret: Optional[str] = None):
        """
        Information about a service principal identity for the cluster to use for manipulating Azure APIs. Either secret or keyVaultSecretRef must be specified.
        :param str client_id: The ID for the service principal.
        :param 'KeyVaultSecretRefResponse' key_vault_secret_ref: Reference to a secret stored in Azure Key Vault.
        :param str secret: The secret password associated with the service principal in plain text.
        """
        pulumi.set(__self__, "client_id", client_id)
        if key_vault_secret_ref is not None:
            pulumi.set(__self__, "key_vault_secret_ref", key_vault_secret_ref)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        The ID for the service principal.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="keyVaultSecretRef")
    def key_vault_secret_ref(self) -> Optional['outputs.KeyVaultSecretRefResponse']:
        """
        Reference to a secret stored in Azure Key Vault.
        """
        return pulumi.get(self, "key_vault_secret_ref")

    @property
    @pulumi.getter
    def secret(self) -> Optional[str]:
        """
        The secret password associated with the service principal in plain text.
        """
        return pulumi.get(self, "secret")


@pulumi.output_type
class ContainerServiceSshConfigurationResponse(dict):
    """
    SSH configuration for Linux-based VMs running on Azure.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "publicKeys":
            suggest = "public_keys"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerServiceSshConfigurationResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerServiceSshConfigurationResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerServiceSshConfigurationResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 public_keys: Sequence['outputs.ContainerServiceSshPublicKeyResponse']):
        """
        SSH configuration for Linux-based VMs running on Azure.
        :param Sequence['ContainerServiceSshPublicKeyResponse'] public_keys: The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
        """
        pulumi.set(__self__, "public_keys", public_keys)

    @property
    @pulumi.getter(name="publicKeys")
    def public_keys(self) -> Sequence['outputs.ContainerServiceSshPublicKeyResponse']:
        """
        The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
        """
        return pulumi.get(self, "public_keys")


@pulumi.output_type
class ContainerServiceSshPublicKeyResponse(dict):
    """
    Contains information about SSH certificate public key data.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "keyData":
            suggest = "key_data"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerServiceSshPublicKeyResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerServiceSshPublicKeyResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerServiceSshPublicKeyResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 key_data: str):
        """
        Contains information about SSH certificate public key data.
        :param str key_data: Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.
        """
        pulumi.set(__self__, "key_data", key_data)

    @property
    @pulumi.getter(name="keyData")
    def key_data(self) -> str:
        """
        Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.
        """
        return pulumi.get(self, "key_data")


@pulumi.output_type
class KeyVaultSecretRefResponse(dict):
    """
    Reference to a secret stored in Azure Key Vault.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "secretName":
            suggest = "secret_name"
        elif key == "vaultID":
            suggest = "vault_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in KeyVaultSecretRefResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        KeyVaultSecretRefResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        KeyVaultSecretRefResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 secret_name: str,
                 vault_id: str,
                 version: Optional[str] = None):
        """
        Reference to a secret stored in Azure Key Vault.
        :param str secret_name: The secret name.
        :param str vault_id: Key vault identifier.
        :param str version: The secret version.
        """
        pulumi.set(__self__, "secret_name", secret_name)
        pulumi.set(__self__, "vault_id", vault_id)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> str:
        """
        The secret name.
        """
        return pulumi.get(self, "secret_name")

    @property
    @pulumi.getter(name="vaultID")
    def vault_id(self) -> str:
        """
        Key vault identifier.
        """
        return pulumi.get(self, "vault_id")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        The secret version.
        """
        return pulumi.get(self, "version")


