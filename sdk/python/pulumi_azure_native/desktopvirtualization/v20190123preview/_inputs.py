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
    'RegistrationInfoArgs',
]

@pulumi.input_type
class RegistrationInfoArgs:
    def __init__(__self__, *,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 reset_token: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        Represents a RegistrationInfo definition.
        :param pulumi.Input[str] expiration_time: Expiration time of registration token.
        :param pulumi.Input[bool] reset_token: Update registration token.
        :param pulumi.Input[str] token: The registration token base64 encoded string.
        """
        if expiration_time is not None:
            pulumi.set(__self__, "expiration_time", expiration_time)
        if reset_token is not None:
            pulumi.set(__self__, "reset_token", reset_token)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration time of registration token.
        """
        return pulumi.get(self, "expiration_time")

    @expiration_time.setter
    def expiration_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_time", value)

    @property
    @pulumi.getter(name="resetToken")
    def reset_token(self) -> Optional[pulumi.Input[bool]]:
        """
        Update registration token.
        """
        return pulumi.get(self, "reset_token")

    @reset_token.setter
    def reset_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reset_token", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The registration token base64 encoded string.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


