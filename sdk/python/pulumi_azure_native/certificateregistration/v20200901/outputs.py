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
    'AppServiceCertificateResponse',
    'CertificateDetailsResponse',
    'SystemDataResponse',
]

@pulumi.output_type
class AppServiceCertificateResponse(dict):
    """
    Key Vault container for a certificate that is purchased through Azure.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "provisioningState":
            suggest = "provisioning_state"
        elif key == "keyVaultId":
            suggest = "key_vault_id"
        elif key == "keyVaultSecretName":
            suggest = "key_vault_secret_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AppServiceCertificateResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AppServiceCertificateResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AppServiceCertificateResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 provisioning_state: str,
                 key_vault_id: Optional[str] = None,
                 key_vault_secret_name: Optional[str] = None):
        """
        Key Vault container for a certificate that is purchased through Azure.
        :param str provisioning_state: Status of the Key Vault secret.
        :param str key_vault_id: Key Vault resource Id.
        :param str key_vault_secret_name: Key Vault secret name.
        """
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if key_vault_id is not None:
            pulumi.set(__self__, "key_vault_id", key_vault_id)
        if key_vault_secret_name is not None:
            pulumi.set(__self__, "key_vault_secret_name", key_vault_secret_name)

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Status of the Key Vault secret.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> Optional[str]:
        """
        Key Vault resource Id.
        """
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter(name="keyVaultSecretName")
    def key_vault_secret_name(self) -> Optional[str]:
        """
        Key Vault secret name.
        """
        return pulumi.get(self, "key_vault_secret_name")


@pulumi.output_type
class CertificateDetailsResponse(dict):
    """
    SSL certificate details.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "notAfter":
            suggest = "not_after"
        elif key == "notBefore":
            suggest = "not_before"
        elif key == "rawData":
            suggest = "raw_data"
        elif key == "serialNumber":
            suggest = "serial_number"
        elif key == "signatureAlgorithm":
            suggest = "signature_algorithm"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CertificateDetailsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CertificateDetailsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CertificateDetailsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 issuer: str,
                 not_after: str,
                 not_before: str,
                 raw_data: str,
                 serial_number: str,
                 signature_algorithm: str,
                 subject: str,
                 thumbprint: str,
                 version: int):
        """
        SSL certificate details.
        :param str issuer: Certificate Issuer.
        :param str not_after: Date Certificate is valid to.
        :param str not_before: Date Certificate is valid from.
        :param str raw_data: Raw certificate data.
        :param str serial_number: Certificate Serial Number.
        :param str signature_algorithm: Certificate Signature algorithm.
        :param str subject: Certificate Subject.
        :param str thumbprint: Certificate Thumbprint.
        :param int version: Certificate Version.
        """
        pulumi.set(__self__, "issuer", issuer)
        pulumi.set(__self__, "not_after", not_after)
        pulumi.set(__self__, "not_before", not_before)
        pulumi.set(__self__, "raw_data", raw_data)
        pulumi.set(__self__, "serial_number", serial_number)
        pulumi.set(__self__, "signature_algorithm", signature_algorithm)
        pulumi.set(__self__, "subject", subject)
        pulumi.set(__self__, "thumbprint", thumbprint)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def issuer(self) -> str:
        """
        Certificate Issuer.
        """
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter(name="notAfter")
    def not_after(self) -> str:
        """
        Date Certificate is valid to.
        """
        return pulumi.get(self, "not_after")

    @property
    @pulumi.getter(name="notBefore")
    def not_before(self) -> str:
        """
        Date Certificate is valid from.
        """
        return pulumi.get(self, "not_before")

    @property
    @pulumi.getter(name="rawData")
    def raw_data(self) -> str:
        """
        Raw certificate data.
        """
        return pulumi.get(self, "raw_data")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> str:
        """
        Certificate Serial Number.
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="signatureAlgorithm")
    def signature_algorithm(self) -> str:
        """
        Certificate Signature algorithm.
        """
        return pulumi.get(self, "signature_algorithm")

    @property
    @pulumi.getter
    def subject(self) -> str:
        """
        Certificate Subject.
        """
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter
    def thumbprint(self) -> str:
        """
        Certificate Thumbprint.
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter
    def version(self) -> int:
        """
        Certificate Version.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class SystemDataResponse(dict):
    """
    Metadata pertaining to creation and last modification of the resource.
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
                 created_at: Optional[str] = None,
                 created_by: Optional[str] = None,
                 created_by_type: Optional[str] = None,
                 last_modified_at: Optional[str] = None,
                 last_modified_by: Optional[str] = None,
                 last_modified_by_type: Optional[str] = None):
        """
        Metadata pertaining to creation and last modification of the resource.
        :param str created_at: The timestamp of resource creation (UTC).
        :param str created_by: The identity that created the resource.
        :param str created_by_type: The type of identity that created the resource.
        :param str last_modified_at: The timestamp of resource last modification (UTC)
        :param str last_modified_by: The identity that last modified the resource.
        :param str last_modified_by_type: The type of identity that last modified the resource.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if created_by_type is not None:
            pulumi.set(__self__, "created_by_type", created_by_type)
        if last_modified_at is not None:
            pulumi.set(__self__, "last_modified_at", last_modified_at)
        if last_modified_by is not None:
            pulumi.set(__self__, "last_modified_by", last_modified_by)
        if last_modified_by_type is not None:
            pulumi.set(__self__, "last_modified_by_type", last_modified_by_type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        The timestamp of resource creation (UTC).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[str]:
        """
        The identity that created the resource.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByType")
    def created_by_type(self) -> Optional[str]:
        """
        The type of identity that created the resource.
        """
        return pulumi.get(self, "created_by_type")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> Optional[str]:
        """
        The timestamp of resource last modification (UTC)
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> Optional[str]:
        """
        The identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedByType")
    def last_modified_by_type(self) -> Optional[str]:
        """
        The type of identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by_type")


