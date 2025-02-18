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
    'CustomDomainResponse',
    'EndpointsResponse',
]

@pulumi.output_type
class CustomDomainResponse(dict):
    """
    The custom domain assigned to this storage account. This can be set via Update.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "useSubDomainName":
            suggest = "use_sub_domain_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomDomainResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomDomainResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomDomainResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: Optional[str] = None,
                 use_sub_domain_name: Optional[bool] = None):
        """
        The custom domain assigned to this storage account. This can be set via Update.
        :param str name: Gets or sets the custom domain name. Name is the CNAME source.
        :param bool use_sub_domain_name: Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if use_sub_domain_name is not None:
            pulumi.set(__self__, "use_sub_domain_name", use_sub_domain_name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets or sets the custom domain name. Name is the CNAME source.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="useSubDomainName")
    def use_sub_domain_name(self) -> Optional[bool]:
        """
        Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates
        """
        return pulumi.get(self, "use_sub_domain_name")


@pulumi.output_type
class EndpointsResponse(dict):
    """
    The URIs that are used to perform a retrieval of a public blob, queue or table object.
    """
    def __init__(__self__, *,
                 blob: Optional[str] = None,
                 queue: Optional[str] = None,
                 table: Optional[str] = None):
        """
        The URIs that are used to perform a retrieval of a public blob, queue or table object.
        :param str blob: Gets the blob endpoint.
        :param str queue: Gets the queue endpoint.
        :param str table: Gets the table endpoint.
        """
        if blob is not None:
            pulumi.set(__self__, "blob", blob)
        if queue is not None:
            pulumi.set(__self__, "queue", queue)
        if table is not None:
            pulumi.set(__self__, "table", table)

    @property
    @pulumi.getter
    def blob(self) -> Optional[str]:
        """
        Gets the blob endpoint.
        """
        return pulumi.get(self, "blob")

    @property
    @pulumi.getter
    def queue(self) -> Optional[str]:
        """
        Gets the queue endpoint.
        """
        return pulumi.get(self, "queue")

    @property
    @pulumi.getter
    def table(self) -> Optional[str]:
        """
        Gets the table endpoint.
        """
        return pulumi.get(self, "table")


