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
    'ComplianceStatusResponse',
    'HelmOperatorPropertiesResponse',
    'SystemDataResponse',
]

@pulumi.output_type
class ComplianceStatusResponse(dict):
    """
    Compliance Status details
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "complianceState":
            suggest = "compliance_state"
        elif key == "lastConfigApplied":
            suggest = "last_config_applied"
        elif key == "messageLevel":
            suggest = "message_level"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ComplianceStatusResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ComplianceStatusResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ComplianceStatusResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 compliance_state: str,
                 last_config_applied: Optional[str] = None,
                 message: Optional[str] = None,
                 message_level: Optional[str] = None):
        """
        Compliance Status details
        :param str compliance_state: The compliance state of the configuration.
        :param str last_config_applied: Datetime the configuration was last applied.
        :param str message: Message from when the configuration was applied.
        :param str message_level: Level of the message.
        """
        pulumi.set(__self__, "compliance_state", compliance_state)
        if last_config_applied is not None:
            pulumi.set(__self__, "last_config_applied", last_config_applied)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if message_level is not None:
            pulumi.set(__self__, "message_level", message_level)

    @property
    @pulumi.getter(name="complianceState")
    def compliance_state(self) -> str:
        """
        The compliance state of the configuration.
        """
        return pulumi.get(self, "compliance_state")

    @property
    @pulumi.getter(name="lastConfigApplied")
    def last_config_applied(self) -> Optional[str]:
        """
        Datetime the configuration was last applied.
        """
        return pulumi.get(self, "last_config_applied")

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        """
        Message from when the configuration was applied.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter(name="messageLevel")
    def message_level(self) -> Optional[str]:
        """
        Level of the message.
        """
        return pulumi.get(self, "message_level")


@pulumi.output_type
class HelmOperatorPropertiesResponse(dict):
    """
    Properties for Helm operator.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "chartValues":
            suggest = "chart_values"
        elif key == "chartVersion":
            suggest = "chart_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in HelmOperatorPropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        HelmOperatorPropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        HelmOperatorPropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 chart_values: Optional[str] = None,
                 chart_version: Optional[str] = None):
        """
        Properties for Helm operator.
        :param str chart_values: Values override for the operator Helm chart.
        :param str chart_version: Version of the operator Helm chart.
        """
        if chart_values is not None:
            pulumi.set(__self__, "chart_values", chart_values)
        if chart_version is not None:
            pulumi.set(__self__, "chart_version", chart_version)

    @property
    @pulumi.getter(name="chartValues")
    def chart_values(self) -> Optional[str]:
        """
        Values override for the operator Helm chart.
        """
        return pulumi.get(self, "chart_values")

    @property
    @pulumi.getter(name="chartVersion")
    def chart_version(self) -> Optional[str]:
        """
        Version of the operator Helm chart.
        """
        return pulumi.get(self, "chart_version")


@pulumi.output_type
class SystemDataResponse(dict):
    """
    Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "createdBy":
            suggest = "created_by"
        elif key == "createdByType":
            suggest = "created_by_type"
        elif key == "lastModifiedAt":
            suggest = "last_modified_at"
        elif key == "lastModifiedBy":
            suggest = "last_modified_by"
        elif key == "lastModifiedByType":
            suggest = "last_modified_by_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SystemDataResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: str,
                 created_by: str,
                 created_by_type: str,
                 last_modified_at: str,
                 last_modified_by: str,
                 last_modified_by_type: str):
        """
        Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
        :param str created_at: The timestamp of resource creation (UTC)
        :param str created_by: A string identifier for the identity that created the resource
        :param str created_by_type: The type of identity that created the resource: user, application, managedIdentity, key
        :param str last_modified_at: The timestamp of resource last modification (UTC)
        :param str last_modified_by: A string identifier for the identity that last modified the resource
        :param str last_modified_by_type: The type of identity that last modified the resource: user, application, managedIdentity, key
        """
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "created_by", created_by)
        pulumi.set(__self__, "created_by_type", created_by_type)
        pulumi.set(__self__, "last_modified_at", last_modified_at)
        pulumi.set(__self__, "last_modified_by", last_modified_by)
        pulumi.set(__self__, "last_modified_by_type", last_modified_by_type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The timestamp of resource creation (UTC)
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> str:
        """
        A string identifier for the identity that created the resource
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByType")
    def created_by_type(self) -> str:
        """
        The type of identity that created the resource: user, application, managedIdentity, key
        """
        return pulumi.get(self, "created_by_type")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> str:
        """
        The timestamp of resource last modification (UTC)
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> str:
        """
        A string identifier for the identity that last modified the resource
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedByType")
    def last_modified_by_type(self) -> str:
        """
        The type of identity that last modified the resource: user, application, managedIdentity, key
        """
        return pulumi.get(self, "last_modified_by_type")


