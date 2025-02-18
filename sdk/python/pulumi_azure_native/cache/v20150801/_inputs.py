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
    'SkuArgs',
]

@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 capacity: pulumi.Input[int],
                 family: pulumi.Input[Union[str, 'SkuFamily']],
                 name: pulumi.Input[Union[str, 'SkuName']]):
        """
        SKU parameters supplied to the create Redis operation.
        :param pulumi.Input[int] capacity: What size of Redis cache to deploy. Valid values: for C family (0, 1, 2, 3, 4, 5, 6), for P family (1, 2, 3, 4).
        :param pulumi.Input[Union[str, 'SkuFamily']] family: Which family to use. Valid values: (C, P).
        :param pulumi.Input[Union[str, 'SkuName']] name: What type of Redis cache to deploy. Valid values: (Basic, Standard, Premium).
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "family", family)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Input[int]:
        """
        What size of Redis cache to deploy. Valid values: for C family (0, 1, 2, 3, 4, 5, 6), for P family (1, 2, 3, 4).
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: pulumi.Input[int]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def family(self) -> pulumi.Input[Union[str, 'SkuFamily']]:
        """
        Which family to use. Valid values: (C, P).
        """
        return pulumi.get(self, "family")

    @family.setter
    def family(self, value: pulumi.Input[Union[str, 'SkuFamily']]):
        pulumi.set(self, "family", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[Union[str, 'SkuName']]:
        """
        What type of Redis cache to deploy. Valid values: (Basic, Standard, Premium).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[Union[str, 'SkuName']]):
        pulumi.set(self, "name", value)


