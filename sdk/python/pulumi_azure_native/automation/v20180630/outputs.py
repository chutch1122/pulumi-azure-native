# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'ContentHashResponse',
    'ContentLinkResponse',
    'ModuleErrorInfoResponse',
    'RunbookDraftResponse',
    'RunbookParameterResponse',
]

@pulumi.output_type
class ContentHashResponse(dict):
    """
    Definition of the runbook property type.
    """
    def __init__(__self__, *,
                 algorithm: str,
                 value: str):
        """
        Definition of the runbook property type.
        :param str algorithm: Gets or sets the content hash algorithm used to hash the content.
        :param str value: Gets or sets expected hash value of the content.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        Gets or sets the content hash algorithm used to hash the content.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Gets or sets expected hash value of the content.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class ContentLinkResponse(dict):
    """
    Definition of the content link.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "contentHash":
            suggest = "content_hash"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContentLinkResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContentLinkResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContentLinkResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 content_hash: Optional['outputs.ContentHashResponse'] = None,
                 uri: Optional[str] = None,
                 version: Optional[str] = None):
        """
        Definition of the content link.
        :param 'ContentHashResponse' content_hash: Gets or sets the hash.
        :param str uri: Gets or sets the uri of the runbook content.
        :param str version: Gets or sets the version of the content.
        """
        if content_hash is not None:
            pulumi.set(__self__, "content_hash", content_hash)
        if uri is not None:
            pulumi.set(__self__, "uri", uri)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="contentHash")
    def content_hash(self) -> Optional['outputs.ContentHashResponse']:
        """
        Gets or sets the hash.
        """
        return pulumi.get(self, "content_hash")

    @property
    @pulumi.getter
    def uri(self) -> Optional[str]:
        """
        Gets or sets the uri of the runbook content.
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        Gets or sets the version of the content.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class ModuleErrorInfoResponse(dict):
    """
    Definition of the module error info type.
    """
    def __init__(__self__, *,
                 code: Optional[str] = None,
                 message: Optional[str] = None):
        """
        Definition of the module error info type.
        :param str code: Gets or sets the error code.
        :param str message: Gets or sets the error message.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)
        if message is not None:
            pulumi.set(__self__, "message", message)

    @property
    @pulumi.getter
    def code(self) -> Optional[str]:
        """
        Gets or sets the error code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        """
        Gets or sets the error message.
        """
        return pulumi.get(self, "message")


@pulumi.output_type
class RunbookDraftResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "creationTime":
            suggest = "creation_time"
        elif key == "draftContentLink":
            suggest = "draft_content_link"
        elif key == "inEdit":
            suggest = "in_edit"
        elif key == "lastModifiedTime":
            suggest = "last_modified_time"
        elif key == "outputTypes":
            suggest = "output_types"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RunbookDraftResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RunbookDraftResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RunbookDraftResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 creation_time: Optional[str] = None,
                 draft_content_link: Optional['outputs.ContentLinkResponse'] = None,
                 in_edit: Optional[bool] = None,
                 last_modified_time: Optional[str] = None,
                 output_types: Optional[Sequence[str]] = None,
                 parameters: Optional[Mapping[str, 'outputs.RunbookParameterResponse']] = None):
        """
        :param str creation_time: Gets or sets the creation time of the runbook draft.
        :param 'ContentLinkResponse' draft_content_link: Gets or sets the draft runbook content link.
        :param bool in_edit: Gets or sets whether runbook is in edit mode.
        :param str last_modified_time: Gets or sets the last modified time of the runbook draft.
        :param Sequence[str] output_types: Gets or sets the runbook output types.
        :param Mapping[str, 'RunbookParameterResponse'] parameters: Gets or sets the runbook draft parameters.
        """
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if draft_content_link is not None:
            pulumi.set(__self__, "draft_content_link", draft_content_link)
        if in_edit is not None:
            pulumi.set(__self__, "in_edit", in_edit)
        if last_modified_time is not None:
            pulumi.set(__self__, "last_modified_time", last_modified_time)
        if output_types is not None:
            pulumi.set(__self__, "output_types", output_types)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[str]:
        """
        Gets or sets the creation time of the runbook draft.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="draftContentLink")
    def draft_content_link(self) -> Optional['outputs.ContentLinkResponse']:
        """
        Gets or sets the draft runbook content link.
        """
        return pulumi.get(self, "draft_content_link")

    @property
    @pulumi.getter(name="inEdit")
    def in_edit(self) -> Optional[bool]:
        """
        Gets or sets whether runbook is in edit mode.
        """
        return pulumi.get(self, "in_edit")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> Optional[str]:
        """
        Gets or sets the last modified time of the runbook draft.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="outputTypes")
    def output_types(self) -> Optional[Sequence[str]]:
        """
        Gets or sets the runbook output types.
        """
        return pulumi.get(self, "output_types")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, 'outputs.RunbookParameterResponse']]:
        """
        Gets or sets the runbook draft parameters.
        """
        return pulumi.get(self, "parameters")


@pulumi.output_type
class RunbookParameterResponse(dict):
    """
    Definition of the runbook parameter type.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "defaultValue":
            suggest = "default_value"
        elif key == "isMandatory":
            suggest = "is_mandatory"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RunbookParameterResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RunbookParameterResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RunbookParameterResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 default_value: Optional[str] = None,
                 is_mandatory: Optional[bool] = None,
                 position: Optional[int] = None,
                 type: Optional[str] = None):
        """
        Definition of the runbook parameter type.
        :param str default_value: Gets or sets the default value of parameter.
        :param bool is_mandatory: Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
        :param int position: Get or sets the position of the parameter.
        :param str type: Gets or sets the type of the parameter.
        """
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)
        if is_mandatory is not None:
            pulumi.set(__self__, "is_mandatory", is_mandatory)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[str]:
        """
        Gets or sets the default value of parameter.
        """
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter(name="isMandatory")
    def is_mandatory(self) -> Optional[bool]:
        """
        Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
        """
        return pulumi.get(self, "is_mandatory")

    @property
    @pulumi.getter
    def position(self) -> Optional[int]:
        """
        Get or sets the position of the parameter.
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Gets or sets the type of the parameter.
        """
        return pulumi.get(self, "type")


