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
    'RedisAccessKeysResponse',
    'SkuResponse',
]

@pulumi.output_type
class RedisAccessKeysResponse(dict):
    """
    Redis cache access keys.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "primaryKey":
            suggest = "primary_key"
        elif key == "secondaryKey":
            suggest = "secondary_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RedisAccessKeysResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RedisAccessKeysResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RedisAccessKeysResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 primary_key: Optional[str] = None,
                 secondary_key: Optional[str] = None):
        """
        Redis cache access keys.
        :param str primary_key: The current primary key that clients can use to authenticate with Redis cache.
        :param str secondary_key: The current secondary key that clients can use to authenticate with Redis cache.
        """
        if primary_key is not None:
            pulumi.set(__self__, "primary_key", primary_key)
        if secondary_key is not None:
            pulumi.set(__self__, "secondary_key", secondary_key)

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> Optional[str]:
        """
        The current primary key that clients can use to authenticate with Redis cache.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> Optional[str]:
        """
        The current secondary key that clients can use to authenticate with Redis cache.
        """
        return pulumi.get(self, "secondary_key")


@pulumi.output_type
class SkuResponse(dict):
    """
    SKU parameters supplied to the create Redis operation.
    """
    def __init__(__self__, *,
                 capacity: int,
                 family: str,
                 name: str):
        """
        SKU parameters supplied to the create Redis operation.
        :param int capacity: What size of Redis cache to deploy. Valid values: for C family (0, 1, 2, 3, 4, 5, 6), for P family (1, 2, 3, 4).
        :param str family: Which family to use. Valid values: (C, P).
        :param str name: What type of Redis cache to deploy. Valid values: (Basic, Standard, Premium).
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "family", family)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def capacity(self) -> int:
        """
        What size of Redis cache to deploy. Valid values: for C family (0, 1, 2, 3, 4, 5, 6), for P family (1, 2, 3, 4).
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def family(self) -> str:
        """
        Which family to use. Valid values: (C, P).
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        What type of Redis cache to deploy. Valid values: (Basic, Standard, Premium).
        """
        return pulumi.get(self, "name")


