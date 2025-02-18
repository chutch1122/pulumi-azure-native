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
    'AssignmentLockSettingsArgs',
    'ManagedServiceIdentityArgs',
    'ParameterDefinitionArgs',
    'ParameterValueBaseArgs',
    'ResourceGroupDefinitionArgs',
    'ResourceGroupValueArgs',
]

@pulumi.input_type
class AssignmentLockSettingsArgs:
    def __init__(__self__, *,
                 mode: Optional[pulumi.Input[Union[str, 'AssignmentLockMode']]] = None):
        """
        Defines how Blueprint-managed resources will be locked.
        :param pulumi.Input[Union[str, 'AssignmentLockMode']] mode: Lock mode.
        """
        if mode is not None:
            pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[Union[str, 'AssignmentLockMode']]]:
        """
        Lock mode.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[Union[str, 'AssignmentLockMode']]]):
        pulumi.set(self, "mode", value)


@pulumi.input_type
class ManagedServiceIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[Union[str, 'ManagedServiceIdentityType']],
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        Managed Service Identity
        :param pulumi.Input[Union[str, 'ManagedServiceIdentityType']] type: Type of the Managed Service Identity.
        :param pulumi.Input[str] principal_id: Azure Active Directory principal ID associated with this Identity.
        :param pulumi.Input[str] tenant_id: ID of the Azure Active Directory.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[Union[str, 'ManagedServiceIdentityType']]:
        """
        Type of the Managed Service Identity.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[Union[str, 'ManagedServiceIdentityType']]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        Azure Active Directory principal ID associated with this Identity.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Azure Active Directory.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class ParameterDefinitionArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[Union[str, 'TemplateParameterType']],
                 allowed_values: Optional[pulumi.Input[Sequence[Any]]] = None,
                 default_value: Optional[Any] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 strong_type: Optional[pulumi.Input[str]] = None):
        """
        Represent a parameter with constrains and metadata.
        :param pulumi.Input[Union[str, 'TemplateParameterType']] type: Allowed data types for Azure Resource Manager template parameters.
        :param pulumi.Input[Sequence[Any]] allowed_values: Array of allowed values for this parameter.
        :param Any default_value: Default Value for this parameter.
        :param pulumi.Input[str] description: Description of this parameter/resourceGroup.
        :param pulumi.Input[str] display_name: DisplayName of this parameter/resourceGroup.
        :param pulumi.Input[str] strong_type: StrongType for UI to render rich experience during assignment time.
        """
        pulumi.set(__self__, "type", type)
        if allowed_values is not None:
            pulumi.set(__self__, "allowed_values", allowed_values)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if strong_type is not None:
            pulumi.set(__self__, "strong_type", strong_type)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[Union[str, 'TemplateParameterType']]:
        """
        Allowed data types for Azure Resource Manager template parameters.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[Union[str, 'TemplateParameterType']]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="allowedValues")
    def allowed_values(self) -> Optional[pulumi.Input[Sequence[Any]]]:
        """
        Array of allowed values for this parameter.
        """
        return pulumi.get(self, "allowed_values")

    @allowed_values.setter
    def allowed_values(self, value: Optional[pulumi.Input[Sequence[Any]]]):
        pulumi.set(self, "allowed_values", value)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[Any]:
        """
        Default Value for this parameter.
        """
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: Optional[Any]):
        pulumi.set(self, "default_value", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of this parameter/resourceGroup.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        DisplayName of this parameter/resourceGroup.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="strongType")
    def strong_type(self) -> Optional[pulumi.Input[str]]:
        """
        StrongType for UI to render rich experience during assignment time.
        """
        return pulumi.get(self, "strong_type")

    @strong_type.setter
    def strong_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "strong_type", value)


@pulumi.input_type
class ParameterValueBaseArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None):
        """
        Base class for ParameterValue.
        :param pulumi.Input[str] description: Optional property, just to establish ParameterValueBase as a BaseClass.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional property, just to establish ParameterValueBase as a BaseClass.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class ResourceGroupDefinitionArgs:
    def __init__(__self__, *,
                 depends_on: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 strong_type: Optional[pulumi.Input[str]] = None):
        """
        Represents an Azure resource group in a Blueprint definition.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] depends_on: Artifacts which need to be deployed before this resource group.
        :param pulumi.Input[str] description: Description of this parameter/resourceGroup.
        :param pulumi.Input[str] display_name: DisplayName of this parameter/resourceGroup.
        :param pulumi.Input[str] location: Location of this resourceGroup, leave empty if the resource group location will be specified during the Blueprint assignment.
        :param pulumi.Input[str] name: Name of this resourceGroup, leave empty if the resource group name will be specified during the Blueprint assignment.
        :param pulumi.Input[str] strong_type: StrongType for UI to render rich experience during assignment time.
        """
        if depends_on is not None:
            pulumi.set(__self__, "depends_on", depends_on)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if strong_type is not None:
            pulumi.set(__self__, "strong_type", strong_type)

    @property
    @pulumi.getter(name="dependsOn")
    def depends_on(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Artifacts which need to be deployed before this resource group.
        """
        return pulumi.get(self, "depends_on")

    @depends_on.setter
    def depends_on(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "depends_on", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of this parameter/resourceGroup.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        DisplayName of this parameter/resourceGroup.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location of this resourceGroup, leave empty if the resource group location will be specified during the Blueprint assignment.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of this resourceGroup, leave empty if the resource group name will be specified during the Blueprint assignment.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="strongType")
    def strong_type(self) -> Optional[pulumi.Input[str]]:
        """
        StrongType for UI to render rich experience during assignment time.
        """
        return pulumi.get(self, "strong_type")

    @strong_type.setter
    def strong_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "strong_type", value)


@pulumi.input_type
class ResourceGroupValueArgs:
    def __init__(__self__, *,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Represents an Azure resource group.
        :param pulumi.Input[str] location: Location of the resource group
        :param pulumi.Input[str] name: Name of the resource group
        """
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location of the resource group
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


