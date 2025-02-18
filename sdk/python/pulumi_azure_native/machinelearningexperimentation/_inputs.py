# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'StorageAccountPropertiesArgs',
]

@pulumi.input_type
class StorageAccountPropertiesArgs:
    def __init__(__self__, *,
                 access_key: pulumi.Input[str],
                 storage_account_id: pulumi.Input[str]):
        """
        The properties of a storage account for a machine learning team account.
        :param pulumi.Input[str] access_key: The access key to the storage account.
        :param pulumi.Input[str] storage_account_id: The fully qualified arm Id of the storage account.
        """
        pulumi.set(__self__, "access_key", access_key)
        pulumi.set(__self__, "storage_account_id", storage_account_id)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Input[str]:
        """
        The access key to the storage account.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Input[str]:
        """
        The fully qualified arm Id of the storage account.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_account_id", value)


