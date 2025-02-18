# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = ['DeviceGroupArgs', 'DeviceGroup']

@pulumi.input_type
class DeviceGroupArgs:
    def __init__(__self__, *,
                 iot_defender_location: pulumi.Input[str],
                 device_group_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DeviceGroup resource.
        :param pulumi.Input[str] iot_defender_location: Defender for IoT location
        :param pulumi.Input[str] device_group_name: Device group name
        """
        pulumi.set(__self__, "iot_defender_location", iot_defender_location)
        if device_group_name is not None:
            pulumi.set(__self__, "device_group_name", device_group_name)

    @property
    @pulumi.getter(name="iotDefenderLocation")
    def iot_defender_location(self) -> pulumi.Input[str]:
        """
        Defender for IoT location
        """
        return pulumi.get(self, "iot_defender_location")

    @iot_defender_location.setter
    def iot_defender_location(self, value: pulumi.Input[str]):
        pulumi.set(self, "iot_defender_location", value)

    @property
    @pulumi.getter(name="deviceGroupName")
    def device_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Device group name
        """
        return pulumi.get(self, "device_group_name")

    @device_group_name.setter
    def device_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_group_name", value)


class DeviceGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_group_name: Optional[pulumi.Input[str]] = None,
                 iot_defender_location: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Device group
        API Version: 2021-02-01-preview.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_group_name: Device group name
        :param pulumi.Input[str] iot_defender_location: Defender for IoT location
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeviceGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Device group
        API Version: 2021-02-01-preview.

        :param str resource_name: The name of the resource.
        :param DeviceGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_group_name: Optional[pulumi.Input[str]] = None,
                 iot_defender_location: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceGroupArgs.__new__(DeviceGroupArgs)

            __props__.__dict__["device_group_name"] = device_group_name
            if iot_defender_location is None and not opts.urn:
                raise TypeError("Missing required property 'iot_defender_location'")
            __props__.__dict__["iot_defender_location"] = iot_defender_location
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:iotsecurity/v20210201preview:DeviceGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DeviceGroup, __self__).__init__(
            'azure-native:iotsecurity:DeviceGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DeviceGroup':
        """
        Get an existing DeviceGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DeviceGroupArgs.__new__(DeviceGroupArgs)

        __props__.__dict__["name"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return DeviceGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

