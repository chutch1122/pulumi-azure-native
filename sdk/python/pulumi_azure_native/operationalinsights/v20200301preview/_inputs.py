# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'ClusterSkuArgs',
    'IdentityArgs',
    'KeyVaultPropertiesArgs',
    'StorageAccountArgs',
    'TagArgs',
    'WorkspaceCappingArgs',
    'WorkspaceSkuArgs',
]

@pulumi.input_type
class ClusterSkuArgs:
    def __init__(__self__, *,
                 capacity: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[Union[str, 'ClusterSkuNameEnum']]] = None):
        """
        The cluster sku definition.
        :param pulumi.Input[float] capacity: The capacity value
        :param pulumi.Input[Union[str, 'ClusterSkuNameEnum']] name: The name of the SKU.
        """
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[float]]:
        """
        The capacity value
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[Union[str, 'ClusterSkuNameEnum']]]:
        """
        The name of the SKU.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[Union[str, 'ClusterSkuNameEnum']]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class IdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input['IdentityType']):
        """
        Identity for the resource.
        :param pulumi.Input['IdentityType'] type: The identity type.
        """
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['IdentityType']:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['IdentityType']):
        pulumi.set(self, "type", value)


@pulumi.input_type
class KeyVaultPropertiesArgs:
    def __init__(__self__, *,
                 key_name: Optional[pulumi.Input[str]] = None,
                 key_vault_uri: Optional[pulumi.Input[str]] = None,
                 key_version: Optional[pulumi.Input[str]] = None):
        """
        The key vault properties.
        :param pulumi.Input[str] key_name: The name of the key associated with the Log Analytics cluster.
        :param pulumi.Input[str] key_vault_uri: The Key Vault uri which holds they key associated with the Log Analytics cluster.
        :param pulumi.Input[str] key_version: The version of the key associated with the Log Analytics cluster.
        """
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if key_vault_uri is not None:
            pulumi.set(__self__, "key_vault_uri", key_vault_uri)
        if key_version is not None:
            pulumi.set(__self__, "key_version", key_version)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the key associated with the Log Analytics cluster.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="keyVaultUri")
    def key_vault_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The Key Vault uri which holds they key associated with the Log Analytics cluster.
        """
        return pulumi.get(self, "key_vault_uri")

    @key_vault_uri.setter
    def key_vault_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_vault_uri", value)

    @property
    @pulumi.getter(name="keyVersion")
    def key_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the key associated with the Log Analytics cluster.
        """
        return pulumi.get(self, "key_version")

    @key_version.setter
    def key_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_version", value)


@pulumi.input_type
class StorageAccountArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 key: pulumi.Input[str]):
        """
        Describes a storage account connection.
        :param pulumi.Input[str] id: The Azure Resource Manager ID of the storage account resource.
        :param pulumi.Input[str] key: The storage account key.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The Azure Resource Manager ID of the storage account resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The storage account key.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)


@pulumi.input_type
class TagArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A tag of a saved search.
        :param pulumi.Input[str] name: The tag name.
        :param pulumi.Input[str] value: The tag value.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The tag name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The tag value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class WorkspaceCappingArgs:
    def __init__(__self__, *,
                 daily_quota_gb: Optional[pulumi.Input[float]] = None):
        """
        The daily volume cap for ingestion.
        :param pulumi.Input[float] daily_quota_gb: The workspace daily quota for ingestion. -1 means unlimited.
        """
        if daily_quota_gb is not None:
            pulumi.set(__self__, "daily_quota_gb", daily_quota_gb)

    @property
    @pulumi.getter(name="dailyQuotaGb")
    def daily_quota_gb(self) -> Optional[pulumi.Input[float]]:
        """
        The workspace daily quota for ingestion. -1 means unlimited.
        """
        return pulumi.get(self, "daily_quota_gb")

    @daily_quota_gb.setter
    def daily_quota_gb(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "daily_quota_gb", value)


@pulumi.input_type
class WorkspaceSkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[Union[str, 'WorkspaceSkuNameEnum']],
                 capacity_reservation_level: Optional[pulumi.Input[int]] = None):
        """
        The SKU (tier) of a workspace.
        :param pulumi.Input[Union[str, 'WorkspaceSkuNameEnum']] name: The name of the SKU.
        :param pulumi.Input[int] capacity_reservation_level: The capacity reservation level for this workspace, when CapacityReservation sku is selected.
        """
        pulumi.set(__self__, "name", name)
        if capacity_reservation_level is not None:
            pulumi.set(__self__, "capacity_reservation_level", capacity_reservation_level)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[Union[str, 'WorkspaceSkuNameEnum']]:
        """
        The name of the SKU.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[Union[str, 'WorkspaceSkuNameEnum']]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="capacityReservationLevel")
    def capacity_reservation_level(self) -> Optional[pulumi.Input[int]]:
        """
        The capacity reservation level for this workspace, when CapacityReservation sku is selected.
        """
        return pulumi.get(self, "capacity_reservation_level")

    @capacity_reservation_level.setter
    def capacity_reservation_level(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity_reservation_level", value)


