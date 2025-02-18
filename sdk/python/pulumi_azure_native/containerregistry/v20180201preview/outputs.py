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
    'BaseImageDependencyResponse',
    'BuildArgumentResponse',
    'DockerBuildStepResponse',
    'PlatformPropertiesResponse',
    'SourceControlAuthInfoResponse',
    'SourceRepositoryPropertiesResponse',
]

@pulumi.output_type
class BaseImageDependencyResponse(dict):
    """
    Properties that describe a base image dependency.
    """
    def __init__(__self__, *,
                 digest: Optional[str] = None,
                 registry: Optional[str] = None,
                 repository: Optional[str] = None,
                 tag: Optional[str] = None,
                 type: Optional[str] = None):
        """
        Properties that describe a base image dependency.
        :param str digest: The sha256-based digest of the image manifest.
        :param str registry: The registry login server.
        :param str repository: The repository name.
        :param str tag: The tag name.
        :param str type: The type of the base image dependency.
        """
        if digest is not None:
            pulumi.set(__self__, "digest", digest)
        if registry is not None:
            pulumi.set(__self__, "registry", registry)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def digest(self) -> Optional[str]:
        """
        The sha256-based digest of the image manifest.
        """
        return pulumi.get(self, "digest")

    @property
    @pulumi.getter
    def registry(self) -> Optional[str]:
        """
        The registry login server.
        """
        return pulumi.get(self, "registry")

    @property
    @pulumi.getter
    def repository(self) -> Optional[str]:
        """
        The repository name.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        The tag name.
        """
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of the base image dependency.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class BuildArgumentResponse(dict):
    """
    Properties of a build argument.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "isSecret":
            suggest = "is_secret"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BuildArgumentResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BuildArgumentResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BuildArgumentResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 type: str,
                 value: str,
                 is_secret: Optional[bool] = None):
        """
        Properties of a build argument.
        :param str name: The name of the argument.
        :param str type: The type of the argument.
        :param str value: The value of the argument.
        :param bool is_secret: Flag to indicate whether the argument represents a secret and want to be removed from build logs.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)
        if is_secret is None:
            is_secret = False
        if is_secret is not None:
            pulumi.set(__self__, "is_secret", is_secret)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the argument.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the argument.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the argument.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="isSecret")
    def is_secret(self) -> Optional[bool]:
        """
        Flag to indicate whether the argument represents a secret and want to be removed from build logs.
        """
        return pulumi.get(self, "is_secret")


@pulumi.output_type
class DockerBuildStepResponse(dict):
    """
    The Docker build step.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "baseImageDependencies":
            suggest = "base_image_dependencies"
        elif key == "provisioningState":
            suggest = "provisioning_state"
        elif key == "baseImageTrigger":
            suggest = "base_image_trigger"
        elif key == "buildArguments":
            suggest = "build_arguments"
        elif key == "contextPath":
            suggest = "context_path"
        elif key == "dockerFilePath":
            suggest = "docker_file_path"
        elif key == "imageNames":
            suggest = "image_names"
        elif key == "isPushEnabled":
            suggest = "is_push_enabled"
        elif key == "noCache":
            suggest = "no_cache"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DockerBuildStepResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DockerBuildStepResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DockerBuildStepResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 base_image_dependencies: Sequence['outputs.BaseImageDependencyResponse'],
                 provisioning_state: str,
                 type: str,
                 base_image_trigger: Optional[str] = None,
                 branch: Optional[str] = None,
                 build_arguments: Optional[Sequence['outputs.BuildArgumentResponse']] = None,
                 context_path: Optional[str] = None,
                 docker_file_path: Optional[str] = None,
                 image_names: Optional[Sequence[str]] = None,
                 is_push_enabled: Optional[bool] = None,
                 no_cache: Optional[bool] = None):
        """
        The Docker build step.
        :param Sequence['BaseImageDependencyResponse'] base_image_dependencies: List of base image dependencies for a step.
        :param str provisioning_state: The provisioning state of the build step.
        :param str type: The type of the step.
               Expected value is 'Docker'.
        :param str base_image_trigger: The type of the auto trigger for base image dependency updates.
        :param str branch: The repository branch name.
        :param Sequence['BuildArgumentResponse'] build_arguments: The custom arguments for building this build step.
        :param str context_path: The relative context path for a docker build in the source.
        :param str docker_file_path: The Docker file path relative to the source control root.
        :param Sequence[str] image_names: The fully qualified image names including the repository and tag.
        :param bool is_push_enabled: The value of this property indicates whether the image built should be pushed to the registry or not.
        :param bool no_cache: The value of this property indicates whether the image cache is enabled or not.
        """
        pulumi.set(__self__, "base_image_dependencies", base_image_dependencies)
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        pulumi.set(__self__, "type", 'Docker')
        if base_image_trigger is not None:
            pulumi.set(__self__, "base_image_trigger", base_image_trigger)
        if branch is not None:
            pulumi.set(__self__, "branch", branch)
        if build_arguments is not None:
            pulumi.set(__self__, "build_arguments", build_arguments)
        if context_path is not None:
            pulumi.set(__self__, "context_path", context_path)
        if docker_file_path is not None:
            pulumi.set(__self__, "docker_file_path", docker_file_path)
        if image_names is not None:
            pulumi.set(__self__, "image_names", image_names)
        if is_push_enabled is None:
            is_push_enabled = True
        if is_push_enabled is not None:
            pulumi.set(__self__, "is_push_enabled", is_push_enabled)
        if no_cache is None:
            no_cache = False
        if no_cache is not None:
            pulumi.set(__self__, "no_cache", no_cache)

    @property
    @pulumi.getter(name="baseImageDependencies")
    def base_image_dependencies(self) -> Sequence['outputs.BaseImageDependencyResponse']:
        """
        List of base image dependencies for a step.
        """
        return pulumi.get(self, "base_image_dependencies")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the build step.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the step.
        Expected value is 'Docker'.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="baseImageTrigger")
    def base_image_trigger(self) -> Optional[str]:
        """
        The type of the auto trigger for base image dependency updates.
        """
        return pulumi.get(self, "base_image_trigger")

    @property
    @pulumi.getter
    def branch(self) -> Optional[str]:
        """
        The repository branch name.
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter(name="buildArguments")
    def build_arguments(self) -> Optional[Sequence['outputs.BuildArgumentResponse']]:
        """
        The custom arguments for building this build step.
        """
        return pulumi.get(self, "build_arguments")

    @property
    @pulumi.getter(name="contextPath")
    def context_path(self) -> Optional[str]:
        """
        The relative context path for a docker build in the source.
        """
        return pulumi.get(self, "context_path")

    @property
    @pulumi.getter(name="dockerFilePath")
    def docker_file_path(self) -> Optional[str]:
        """
        The Docker file path relative to the source control root.
        """
        return pulumi.get(self, "docker_file_path")

    @property
    @pulumi.getter(name="imageNames")
    def image_names(self) -> Optional[Sequence[str]]:
        """
        The fully qualified image names including the repository and tag.
        """
        return pulumi.get(self, "image_names")

    @property
    @pulumi.getter(name="isPushEnabled")
    def is_push_enabled(self) -> Optional[bool]:
        """
        The value of this property indicates whether the image built should be pushed to the registry or not.
        """
        return pulumi.get(self, "is_push_enabled")

    @property
    @pulumi.getter(name="noCache")
    def no_cache(self) -> Optional[bool]:
        """
        The value of this property indicates whether the image cache is enabled or not.
        """
        return pulumi.get(self, "no_cache")


@pulumi.output_type
class PlatformPropertiesResponse(dict):
    """
    The platform properties against which the build has to happen.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "osType":
            suggest = "os_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PlatformPropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PlatformPropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PlatformPropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 os_type: str,
                 cpu: Optional[int] = None):
        """
        The platform properties against which the build has to happen.
        :param str os_type: The operating system type required for the build.
        :param int cpu: The CPU configuration in terms of number of cores required for the build.
        """
        pulumi.set(__self__, "os_type", os_type)
        if cpu is not None:
            pulumi.set(__self__, "cpu", cpu)

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> str:
        """
        The operating system type required for the build.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter
    def cpu(self) -> Optional[int]:
        """
        The CPU configuration in terms of number of cores required for the build.
        """
        return pulumi.get(self, "cpu")


@pulumi.output_type
class SourceControlAuthInfoResponse(dict):
    """
    The authorization properties for accessing the source code repository.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expiresIn":
            suggest = "expires_in"
        elif key == "refreshToken":
            suggest = "refresh_token"
        elif key == "tokenType":
            suggest = "token_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SourceControlAuthInfoResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SourceControlAuthInfoResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SourceControlAuthInfoResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 token: str,
                 expires_in: Optional[int] = None,
                 refresh_token: Optional[str] = None,
                 scope: Optional[str] = None,
                 token_type: Optional[str] = None):
        """
        The authorization properties for accessing the source code repository.
        :param str token: The access token used to access the source control provider.
        :param int expires_in: Time in seconds that the token remains valid
        :param str refresh_token: The refresh token used to refresh the access token.
        :param str scope: The scope of the access token.
        :param str token_type: The type of Auth token.
        """
        pulumi.set(__self__, "token", token)
        if expires_in is not None:
            pulumi.set(__self__, "expires_in", expires_in)
        if refresh_token is not None:
            pulumi.set(__self__, "refresh_token", refresh_token)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if token_type is not None:
            pulumi.set(__self__, "token_type", token_type)

    @property
    @pulumi.getter
    def token(self) -> str:
        """
        The access token used to access the source control provider.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="expiresIn")
    def expires_in(self) -> Optional[int]:
        """
        Time in seconds that the token remains valid
        """
        return pulumi.get(self, "expires_in")

    @property
    @pulumi.getter(name="refreshToken")
    def refresh_token(self) -> Optional[str]:
        """
        The refresh token used to refresh the access token.
        """
        return pulumi.get(self, "refresh_token")

    @property
    @pulumi.getter
    def scope(self) -> Optional[str]:
        """
        The scope of the access token.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter(name="tokenType")
    def token_type(self) -> Optional[str]:
        """
        The type of Auth token.
        """
        return pulumi.get(self, "token_type")


@pulumi.output_type
class SourceRepositoryPropertiesResponse(dict):
    """
    The properties of the source code repository.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "repositoryUrl":
            suggest = "repository_url"
        elif key == "sourceControlType":
            suggest = "source_control_type"
        elif key == "isCommitTriggerEnabled":
            suggest = "is_commit_trigger_enabled"
        elif key == "sourceControlAuthProperties":
            suggest = "source_control_auth_properties"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SourceRepositoryPropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SourceRepositoryPropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SourceRepositoryPropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 repository_url: str,
                 source_control_type: str,
                 is_commit_trigger_enabled: Optional[bool] = None,
                 source_control_auth_properties: Optional['outputs.SourceControlAuthInfoResponse'] = None):
        """
        The properties of the source code repository.
        :param str repository_url: The full URL to the source code repository
        :param str source_control_type: The type of source control service.
        :param bool is_commit_trigger_enabled: The value of this property indicates whether the source control commit trigger is enabled or not.
        :param 'SourceControlAuthInfoResponse' source_control_auth_properties: The authorization properties for accessing the source code repository.
        """
        pulumi.set(__self__, "repository_url", repository_url)
        pulumi.set(__self__, "source_control_type", source_control_type)
        if is_commit_trigger_enabled is None:
            is_commit_trigger_enabled = False
        if is_commit_trigger_enabled is not None:
            pulumi.set(__self__, "is_commit_trigger_enabled", is_commit_trigger_enabled)
        if source_control_auth_properties is not None:
            pulumi.set(__self__, "source_control_auth_properties", source_control_auth_properties)

    @property
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> str:
        """
        The full URL to the source code repository
        """
        return pulumi.get(self, "repository_url")

    @property
    @pulumi.getter(name="sourceControlType")
    def source_control_type(self) -> str:
        """
        The type of source control service.
        """
        return pulumi.get(self, "source_control_type")

    @property
    @pulumi.getter(name="isCommitTriggerEnabled")
    def is_commit_trigger_enabled(self) -> Optional[bool]:
        """
        The value of this property indicates whether the source control commit trigger is enabled or not.
        """
        return pulumi.get(self, "is_commit_trigger_enabled")

    @property
    @pulumi.getter(name="sourceControlAuthProperties")
    def source_control_auth_properties(self) -> Optional['outputs.SourceControlAuthInfoResponse']:
        """
        The authorization properties for accessing the source code repository.
        """
        return pulumi.get(self, "source_control_auth_properties")


