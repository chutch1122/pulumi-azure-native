# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'StorageAccountPropertiesResponse',
]

@pulumi.output_type
class StorageAccountPropertiesResponse(dict):
    """
    The properties of a storage account for a container registry.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accessKey":
            suggest = "access_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StorageAccountPropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StorageAccountPropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StorageAccountPropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 access_key: str,
                 name: str):
        """
        The properties of a storage account for a container registry.
        :param str access_key: The access key to the storage account.
        :param str name: The name of the storage account.
        """
        pulumi.set(__self__, "access_key", access_key)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> str:
        """
        The access key to the storage account.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the storage account.
        """
        return pulumi.get(self, "name")


