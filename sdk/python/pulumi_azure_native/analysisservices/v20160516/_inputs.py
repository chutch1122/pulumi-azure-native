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
    'ResourceSkuArgs',
    'ServerAdministratorsArgs',
]

@pulumi.input_type
class ResourceSkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 capacity: Optional[pulumi.Input[int]] = None,
                 tier: Optional[pulumi.Input[Union[str, 'SkuTier']]] = None):
        """
        Represents the SKU name and Azure pricing tier for Analysis Services resource.
        :param pulumi.Input[str] name: Name of the SKU level.
        :param pulumi.Input[int] capacity: The number of instances in the read only query pool.
        :param pulumi.Input[Union[str, 'SkuTier']] tier: The name of the Azure pricing tier to which the SKU applies.
        """
        pulumi.set(__self__, "name", name)
        if capacity is None:
            capacity = 1
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the SKU level.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[int]]:
        """
        The number of instances in the read only query pool.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[Union[str, 'SkuTier']]]:
        """
        The name of the Azure pricing tier to which the SKU applies.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[Union[str, 'SkuTier']]]):
        pulumi.set(self, "tier", value)


@pulumi.input_type
class ServerAdministratorsArgs:
    def __init__(__self__, *,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        An array of administrator user identities
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: An array of administrator user identities.
        """
        if members is not None:
            pulumi.set(__self__, "members", members)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of administrator user identities.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)


