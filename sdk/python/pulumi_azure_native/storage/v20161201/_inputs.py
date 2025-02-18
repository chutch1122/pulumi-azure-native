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
    'CustomDomainArgs',
    'EncryptionServicesArgs',
    'EncryptionServiceArgs',
    'EncryptionArgs',
    'SkuArgs',
]

@pulumi.input_type
class CustomDomainArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 use_sub_domain_name: Optional[pulumi.Input[bool]] = None):
        """
        The custom domain assigned to this storage account. This can be set via Update.
        :param pulumi.Input[str] name: Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
        :param pulumi.Input[bool] use_sub_domain_name: Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
        """
        pulumi.set(__self__, "name", name)
        if use_sub_domain_name is not None:
            pulumi.set(__self__, "use_sub_domain_name", use_sub_domain_name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="useSubDomainName")
    def use_sub_domain_name(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
        """
        return pulumi.get(self, "use_sub_domain_name")

    @use_sub_domain_name.setter
    def use_sub_domain_name(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_sub_domain_name", value)


@pulumi.input_type
class EncryptionServicesArgs:
    def __init__(__self__, *,
                 blob: Optional[pulumi.Input['EncryptionServiceArgs']] = None,
                 file: Optional[pulumi.Input['EncryptionServiceArgs']] = None):
        """
        A list of services that support encryption.
        :param pulumi.Input['EncryptionServiceArgs'] blob: The encryption function of the blob storage service.
        :param pulumi.Input['EncryptionServiceArgs'] file: The encryption function of the file storage service.
        """
        if blob is not None:
            pulumi.set(__self__, "blob", blob)
        if file is not None:
            pulumi.set(__self__, "file", file)

    @property
    @pulumi.getter
    def blob(self) -> Optional[pulumi.Input['EncryptionServiceArgs']]:
        """
        The encryption function of the blob storage service.
        """
        return pulumi.get(self, "blob")

    @blob.setter
    def blob(self, value: Optional[pulumi.Input['EncryptionServiceArgs']]):
        pulumi.set(self, "blob", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input['EncryptionServiceArgs']]:
        """
        The encryption function of the file storage service.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input['EncryptionServiceArgs']]):
        pulumi.set(self, "file", value)


@pulumi.input_type
class EncryptionServiceArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        A service that allows server-side encryption to be used.
        :param pulumi.Input[bool] enabled: A boolean indicating whether or not the service encrypts the data as it is stored.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean indicating whether or not the service encrypts the data as it is stored.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class EncryptionArgs:
    def __init__(__self__, *,
                 key_source: pulumi.Input[str],
                 services: Optional[pulumi.Input['EncryptionServicesArgs']] = None):
        """
        The encryption settings on the storage account.
        :param pulumi.Input[str] key_source: The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage
        :param pulumi.Input['EncryptionServicesArgs'] services: List of services which support encryption.
        """
        pulumi.set(__self__, "key_source", key_source)
        if services is not None:
            pulumi.set(__self__, "services", services)

    @property
    @pulumi.getter(name="keySource")
    def key_source(self) -> pulumi.Input[str]:
        """
        The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage
        """
        return pulumi.get(self, "key_source")

    @key_source.setter
    def key_source(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_source", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input['EncryptionServicesArgs']]:
        """
        List of services which support encryption.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input['EncryptionServicesArgs']]):
        pulumi.set(self, "services", value)


@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input['SkuName']):
        """
        The SKU of the storage account.
        :param pulumi.Input['SkuName'] name: Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input['SkuName']:
        """
        Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input['SkuName']):
        pulumi.set(self, "name", value)


