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
    'CreationDataArgs',
    'DiskSkuArgs',
    'EncryptionSetIdentityArgs',
    'EncryptionSettingsCollectionArgs',
    'EncryptionSettingsElementArgs',
    'EncryptionArgs',
    'ImageDiskReferenceArgs',
    'KeyVaultAndKeyReferenceArgs',
    'KeyVaultAndSecretReferenceArgs',
    'SnapshotSkuArgs',
    'SourceVaultArgs',
]

@pulumi.input_type
class CreationDataArgs:
    def __init__(__self__, *,
                 create_option: pulumi.Input[Union[str, 'DiskCreateOption']],
                 gallery_image_reference: Optional[pulumi.Input['ImageDiskReferenceArgs']] = None,
                 image_reference: Optional[pulumi.Input['ImageDiskReferenceArgs']] = None,
                 source_resource_id: Optional[pulumi.Input[str]] = None,
                 source_uri: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 upload_size_bytes: Optional[pulumi.Input[float]] = None):
        """
        Data used when creating a disk.
        :param pulumi.Input[Union[str, 'DiskCreateOption']] create_option: This enumerates the possible sources of a disk's creation.
        :param pulumi.Input['ImageDiskReferenceArgs'] gallery_image_reference: Required if creating from a Gallery Image. The id of the ImageDiskReference will be the ARM id of the shared galley image version from which to create a disk.
        :param pulumi.Input['ImageDiskReferenceArgs'] image_reference: Disk source information.
        :param pulumi.Input[str] source_resource_id: If createOption is Copy, this is the ARM id of the source snapshot or disk.
        :param pulumi.Input[str] source_uri: If createOption is Import, this is the URI of a blob to be imported into a managed disk.
        :param pulumi.Input[str] storage_account_id: Required if createOption is Import. The Azure Resource Manager identifier of the storage account containing the blob to import as a disk.
        :param pulumi.Input[float] upload_size_bytes: If createOption is Upload, this is the size of the contents of the upload including the VHD footer. This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512 bytes for the VHD footer).
        """
        pulumi.set(__self__, "create_option", create_option)
        if gallery_image_reference is not None:
            pulumi.set(__self__, "gallery_image_reference", gallery_image_reference)
        if image_reference is not None:
            pulumi.set(__self__, "image_reference", image_reference)
        if source_resource_id is not None:
            pulumi.set(__self__, "source_resource_id", source_resource_id)
        if source_uri is not None:
            pulumi.set(__self__, "source_uri", source_uri)
        if storage_account_id is not None:
            pulumi.set(__self__, "storage_account_id", storage_account_id)
        if upload_size_bytes is not None:
            pulumi.set(__self__, "upload_size_bytes", upload_size_bytes)

    @property
    @pulumi.getter(name="createOption")
    def create_option(self) -> pulumi.Input[Union[str, 'DiskCreateOption']]:
        """
        This enumerates the possible sources of a disk's creation.
        """
        return pulumi.get(self, "create_option")

    @create_option.setter
    def create_option(self, value: pulumi.Input[Union[str, 'DiskCreateOption']]):
        pulumi.set(self, "create_option", value)

    @property
    @pulumi.getter(name="galleryImageReference")
    def gallery_image_reference(self) -> Optional[pulumi.Input['ImageDiskReferenceArgs']]:
        """
        Required if creating from a Gallery Image. The id of the ImageDiskReference will be the ARM id of the shared galley image version from which to create a disk.
        """
        return pulumi.get(self, "gallery_image_reference")

    @gallery_image_reference.setter
    def gallery_image_reference(self, value: Optional[pulumi.Input['ImageDiskReferenceArgs']]):
        pulumi.set(self, "gallery_image_reference", value)

    @property
    @pulumi.getter(name="imageReference")
    def image_reference(self) -> Optional[pulumi.Input['ImageDiskReferenceArgs']]:
        """
        Disk source information.
        """
        return pulumi.get(self, "image_reference")

    @image_reference.setter
    def image_reference(self, value: Optional[pulumi.Input['ImageDiskReferenceArgs']]):
        pulumi.set(self, "image_reference", value)

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        If createOption is Copy, this is the ARM id of the source snapshot or disk.
        """
        return pulumi.get(self, "source_resource_id")

    @source_resource_id.setter
    def source_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_resource_id", value)

    @property
    @pulumi.getter(name="sourceUri")
    def source_uri(self) -> Optional[pulumi.Input[str]]:
        """
        If createOption is Import, this is the URI of a blob to be imported into a managed disk.
        """
        return pulumi.get(self, "source_uri")

    @source_uri.setter
    def source_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_uri", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required if createOption is Import. The Azure Resource Manager identifier of the storage account containing the blob to import as a disk.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_id", value)

    @property
    @pulumi.getter(name="uploadSizeBytes")
    def upload_size_bytes(self) -> Optional[pulumi.Input[float]]:
        """
        If createOption is Upload, this is the size of the contents of the upload including the VHD footer. This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512 bytes for the VHD footer).
        """
        return pulumi.get(self, "upload_size_bytes")

    @upload_size_bytes.setter
    def upload_size_bytes(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "upload_size_bytes", value)


@pulumi.input_type
class DiskSkuArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[Union[str, 'DiskStorageAccountTypes']]] = None):
        """
        The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
        :param pulumi.Input[Union[str, 'DiskStorageAccountTypes']] name: The sku name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[Union[str, 'DiskStorageAccountTypes']]]:
        """
        The sku name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[Union[str, 'DiskStorageAccountTypes']]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class EncryptionSetIdentityArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[Union[str, 'DiskEncryptionSetIdentityType']]] = None):
        """
        The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
        :param pulumi.Input[Union[str, 'DiskEncryptionSetIdentityType']] type: The type of Managed Identity used by the DiskEncryptionSet. Only SystemAssigned is supported.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[Union[str, 'DiskEncryptionSetIdentityType']]]:
        """
        The type of Managed Identity used by the DiskEncryptionSet. Only SystemAssigned is supported.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[Union[str, 'DiskEncryptionSetIdentityType']]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class EncryptionSettingsCollectionArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 encryption_settings: Optional[pulumi.Input[Sequence[pulumi.Input['EncryptionSettingsElementArgs']]]] = None,
                 encryption_settings_version: Optional[pulumi.Input[str]] = None):
        """
        Encryption settings for disk or snapshot
        :param pulumi.Input[bool] enabled: Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is null in the request object, the existing settings remain unchanged.
        :param pulumi.Input[Sequence[pulumi.Input['EncryptionSettingsElementArgs']]] encryption_settings: A collection of encryption settings, one for each disk volume.
        :param pulumi.Input[str] encryption_settings_version: Describes what type of encryption is used for the disks. Once this field is set, it cannot be overwritten. '1.0' corresponds to Azure Disk Encryption with AAD app.'1.1' corresponds to Azure Disk Encryption.
        """
        pulumi.set(__self__, "enabled", enabled)
        if encryption_settings is not None:
            pulumi.set(__self__, "encryption_settings", encryption_settings)
        if encryption_settings_version is not None:
            pulumi.set(__self__, "encryption_settings_version", encryption_settings_version)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is null in the request object, the existing settings remain unchanged.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="encryptionSettings")
    def encryption_settings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EncryptionSettingsElementArgs']]]]:
        """
        A collection of encryption settings, one for each disk volume.
        """
        return pulumi.get(self, "encryption_settings")

    @encryption_settings.setter
    def encryption_settings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EncryptionSettingsElementArgs']]]]):
        pulumi.set(self, "encryption_settings", value)

    @property
    @pulumi.getter(name="encryptionSettingsVersion")
    def encryption_settings_version(self) -> Optional[pulumi.Input[str]]:
        """
        Describes what type of encryption is used for the disks. Once this field is set, it cannot be overwritten. '1.0' corresponds to Azure Disk Encryption with AAD app.'1.1' corresponds to Azure Disk Encryption.
        """
        return pulumi.get(self, "encryption_settings_version")

    @encryption_settings_version.setter
    def encryption_settings_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_settings_version", value)


@pulumi.input_type
class EncryptionSettingsElementArgs:
    def __init__(__self__, *,
                 disk_encryption_key: Optional[pulumi.Input['KeyVaultAndSecretReferenceArgs']] = None,
                 key_encryption_key: Optional[pulumi.Input['KeyVaultAndKeyReferenceArgs']] = None):
        """
        Encryption settings for one disk volume.
        :param pulumi.Input['KeyVaultAndSecretReferenceArgs'] disk_encryption_key: Key Vault Secret Url and vault id of the disk encryption key
        :param pulumi.Input['KeyVaultAndKeyReferenceArgs'] key_encryption_key: Key Vault Key Url and vault id of the key encryption key. KeyEncryptionKey is optional and when provided is used to unwrap the disk encryption key.
        """
        if disk_encryption_key is not None:
            pulumi.set(__self__, "disk_encryption_key", disk_encryption_key)
        if key_encryption_key is not None:
            pulumi.set(__self__, "key_encryption_key", key_encryption_key)

    @property
    @pulumi.getter(name="diskEncryptionKey")
    def disk_encryption_key(self) -> Optional[pulumi.Input['KeyVaultAndSecretReferenceArgs']]:
        """
        Key Vault Secret Url and vault id of the disk encryption key
        """
        return pulumi.get(self, "disk_encryption_key")

    @disk_encryption_key.setter
    def disk_encryption_key(self, value: Optional[pulumi.Input['KeyVaultAndSecretReferenceArgs']]):
        pulumi.set(self, "disk_encryption_key", value)

    @property
    @pulumi.getter(name="keyEncryptionKey")
    def key_encryption_key(self) -> Optional[pulumi.Input['KeyVaultAndKeyReferenceArgs']]:
        """
        Key Vault Key Url and vault id of the key encryption key. KeyEncryptionKey is optional and when provided is used to unwrap the disk encryption key.
        """
        return pulumi.get(self, "key_encryption_key")

    @key_encryption_key.setter
    def key_encryption_key(self, value: Optional[pulumi.Input['KeyVaultAndKeyReferenceArgs']]):
        pulumi.set(self, "key_encryption_key", value)


@pulumi.input_type
class EncryptionArgs:
    def __init__(__self__, *,
                 disk_encryption_set_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[Union[str, 'EncryptionType']]] = None):
        """
        Encryption at rest settings for disk or snapshot
        :param pulumi.Input[str] disk_encryption_set_id: ResourceId of the disk encryption set to use for enabling encryption at rest.
        :param pulumi.Input[Union[str, 'EncryptionType']] type: The type of key used to encrypt the data of the disk.
        """
        if disk_encryption_set_id is not None:
            pulumi.set(__self__, "disk_encryption_set_id", disk_encryption_set_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="diskEncryptionSetId")
    def disk_encryption_set_id(self) -> Optional[pulumi.Input[str]]:
        """
        ResourceId of the disk encryption set to use for enabling encryption at rest.
        """
        return pulumi.get(self, "disk_encryption_set_id")

    @disk_encryption_set_id.setter
    def disk_encryption_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_encryption_set_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[Union[str, 'EncryptionType']]]:
        """
        The type of key used to encrypt the data of the disk.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[Union[str, 'EncryptionType']]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ImageDiskReferenceArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 lun: Optional[pulumi.Input[int]] = None):
        """
        The source image used for creating the disk.
        :param pulumi.Input[str] id: A relative uri containing either a Platform Image Repository or user image reference.
        :param pulumi.Input[int] lun: If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use. For OS disks, this field is null.
        """
        pulumi.set(__self__, "id", id)
        if lun is not None:
            pulumi.set(__self__, "lun", lun)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        A relative uri containing either a Platform Image Repository or user image reference.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def lun(self) -> Optional[pulumi.Input[int]]:
        """
        If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use. For OS disks, this field is null.
        """
        return pulumi.get(self, "lun")

    @lun.setter
    def lun(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "lun", value)


@pulumi.input_type
class KeyVaultAndKeyReferenceArgs:
    def __init__(__self__, *,
                 key_url: pulumi.Input[str],
                 source_vault: pulumi.Input['SourceVaultArgs']):
        """
        Key Vault Key Url and vault id of KeK, KeK is optional and when provided is used to unwrap the encryptionKey
        :param pulumi.Input[str] key_url: Url pointing to a key or secret in KeyVault
        :param pulumi.Input['SourceVaultArgs'] source_vault: Resource id of the KeyVault containing the key or secret
        """
        pulumi.set(__self__, "key_url", key_url)
        pulumi.set(__self__, "source_vault", source_vault)

    @property
    @pulumi.getter(name="keyUrl")
    def key_url(self) -> pulumi.Input[str]:
        """
        Url pointing to a key or secret in KeyVault
        """
        return pulumi.get(self, "key_url")

    @key_url.setter
    def key_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_url", value)

    @property
    @pulumi.getter(name="sourceVault")
    def source_vault(self) -> pulumi.Input['SourceVaultArgs']:
        """
        Resource id of the KeyVault containing the key or secret
        """
        return pulumi.get(self, "source_vault")

    @source_vault.setter
    def source_vault(self, value: pulumi.Input['SourceVaultArgs']):
        pulumi.set(self, "source_vault", value)


@pulumi.input_type
class KeyVaultAndSecretReferenceArgs:
    def __init__(__self__, *,
                 secret_url: pulumi.Input[str],
                 source_vault: pulumi.Input['SourceVaultArgs']):
        """
        Key Vault Secret Url and vault id of the encryption key 
        :param pulumi.Input[str] secret_url: Url pointing to a key or secret in KeyVault
        :param pulumi.Input['SourceVaultArgs'] source_vault: Resource id of the KeyVault containing the key or secret
        """
        pulumi.set(__self__, "secret_url", secret_url)
        pulumi.set(__self__, "source_vault", source_vault)

    @property
    @pulumi.getter(name="secretUrl")
    def secret_url(self) -> pulumi.Input[str]:
        """
        Url pointing to a key or secret in KeyVault
        """
        return pulumi.get(self, "secret_url")

    @secret_url.setter
    def secret_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_url", value)

    @property
    @pulumi.getter(name="sourceVault")
    def source_vault(self) -> pulumi.Input['SourceVaultArgs']:
        """
        Resource id of the KeyVault containing the key or secret
        """
        return pulumi.get(self, "source_vault")

    @source_vault.setter
    def source_vault(self, value: pulumi.Input['SourceVaultArgs']):
        pulumi.set(self, "source_vault", value)


@pulumi.input_type
class SnapshotSkuArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[Union[str, 'SnapshotStorageAccountTypes']]] = None):
        """
        The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS.
        :param pulumi.Input[Union[str, 'SnapshotStorageAccountTypes']] name: The sku name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[Union[str, 'SnapshotStorageAccountTypes']]]:
        """
        The sku name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[Union[str, 'SnapshotStorageAccountTypes']]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class SourceVaultArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None):
        """
        The vault id is an Azure Resource Manager Resource id in the form /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}
        :param pulumi.Input[str] id: Resource Id
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource Id
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


