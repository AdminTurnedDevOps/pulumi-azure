# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Account(pulumi.CustomResource):
    access_tier: pulumi.Output[str]
    """
    Defines the access tier for `BlobStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
    """
    account_encryption_source: pulumi.Output[str]
    """
    The Encryption Source for this Storage Account. Possible values are `Microsoft.Keyvault` and `Microsoft.Storage`. Defaults to `Microsoft.Storage`.
    """
    account_kind: pulumi.Output[str]
    """
    Defines the Kind of account. Valid options are `Storage`,
    `StorageV2` and `BlobStorage`. Changing this forces a new resource to be created.
    Defaults to `Storage`.
    """
    account_replication_type: pulumi.Output[str]
    """
    Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS` and `ZRS`.
    """
    account_tier: pulumi.Output[str]
    """
    Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. Changing this forces a new resource to be created
    """
    account_type: pulumi.Output[str]
    custom_domain: pulumi.Output[dict]
    """
    A `custom_domain` block as documented below.
    """
    enable_blob_encryption: pulumi.Output[bool]
    """
    Boolean flag which controls if Encryption Services are enabled for Blob storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
    """
    enable_file_encryption: pulumi.Output[bool]
    """
    Boolean flag which controls if Encryption Services are enabled for File storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
    """
    enable_https_traffic_only: pulumi.Output[bool]
    """
    Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
    for more information.
    """
    identity: pulumi.Output[dict]
    """
    A Managed Service Identity block as defined below.
    """
    is_hns_enabled: pulumi.Output[bool]
    """
    Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the
    resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The Custom Domain Name to use for the Storage Account, which will be validated by Azure.
    """
    network_rules: pulumi.Output[dict]
    """
    A `network_rules` block as documented below.
    """
    primary_access_key: pulumi.Output[str]
    """
    The primary access key for the storage account.
    """
    primary_blob_connection_string: pulumi.Output[str]
    """
    The connection string associated with the primary blob location.
    """
    primary_blob_endpoint: pulumi.Output[str]
    """
    The endpoint URL for blob storage in the primary location.
    """
    primary_blob_host: pulumi.Output[str]
    """
    The hostname with port if applicable for blob storage in the primary location.
    """
    primary_connection_string: pulumi.Output[str]
    """
    The connection string associated with the primary location.
    """
    primary_dfs_endpoint: pulumi.Output[str]
    """
    The endpoint URL for DFS storage in the primary location.
    """
    primary_dfs_host: pulumi.Output[str]
    """
    The hostname with port if applicable for DFS storage in the primary location.
    """
    primary_file_endpoint: pulumi.Output[str]
    """
    The endpoint URL for file storage in the primary location.
    """
    primary_file_host: pulumi.Output[str]
    """
    The hostname with port if applicable for file storage in the primary location.
    """
    primary_location: pulumi.Output[str]
    """
    The primary location of the storage account.
    """
    primary_queue_endpoint: pulumi.Output[str]
    """
    The endpoint URL for queue storage in the primary location.
    """
    primary_queue_host: pulumi.Output[str]
    """
    The hostname with port if applicable for queue storage in the primary location.
    """
    primary_table_endpoint: pulumi.Output[str]
    """
    The endpoint URL for table storage in the primary location.
    """
    primary_table_host: pulumi.Output[str]
    """
    The hostname with port if applicable for table storage in the primary location.
    """
    primary_web_endpoint: pulumi.Output[str]
    """
    The endpoint URL for web storage in the primary location.
    """
    primary_web_host: pulumi.Output[str]
    """
    The hostname with port if applicable for web storage in the primary location.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to
    create the storage account. Changing this forces a new resource to be created.
    """
    secondary_access_key: pulumi.Output[str]
    """
    The secondary access key for the storage account.
    """
    secondary_blob_connection_string: pulumi.Output[str]
    """
    The connection string associated with the secondary blob location.
    """
    secondary_blob_endpoint: pulumi.Output[str]
    """
    The endpoint URL for blob storage in the secondary location.
    """
    secondary_blob_host: pulumi.Output[str]
    """
    The hostname with port if applicable for blob storage in the secondary location.
    """
    secondary_connection_string: pulumi.Output[str]
    """
    The connection string associated with the secondary location.
    """
    secondary_dfs_endpoint: pulumi.Output[str]
    """
    The endpoint URL for DFS storage in the secondary location.
    """
    secondary_dfs_host: pulumi.Output[str]
    """
    The hostname with port if applicable for DFS storage in the secondary location.
    """
    secondary_file_endpoint: pulumi.Output[str]
    """
    The endpoint URL for file storage in the secondary location.
    """
    secondary_file_host: pulumi.Output[str]
    """
    The hostname with port if applicable for file storage in the secondary location.
    """
    secondary_location: pulumi.Output[str]
    """
    The secondary location of the storage account.
    """
    secondary_queue_endpoint: pulumi.Output[str]
    """
    The endpoint URL for queue storage in the secondary location.
    """
    secondary_queue_host: pulumi.Output[str]
    """
    The hostname with port if applicable for queue storage in the secondary location.
    """
    secondary_table_endpoint: pulumi.Output[str]
    """
    The endpoint URL for table storage in the secondary location.
    """
    secondary_table_host: pulumi.Output[str]
    """
    The hostname with port if applicable for table storage in the secondary location.
    """
    secondary_web_endpoint: pulumi.Output[str]
    """
    The endpoint URL for web storage in the secondary location.
    """
    secondary_web_host: pulumi.Output[str]
    """
    The hostname with port if applicable for web storage in the secondary location.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, access_tier=None, account_encryption_source=None, account_kind=None, account_replication_type=None, account_tier=None, account_type=None, custom_domain=None, enable_blob_encryption=None, enable_file_encryption=None, enable_https_traffic_only=None, identity=None, is_hns_enabled=None, location=None, name=None, network_rules=None, resource_group_name=None, tags=None, __name__=None, __opts__=None):
        """
        Manage an Azure Storage Account.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_tier: Defines the access tier for `BlobStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
        :param pulumi.Input[str] account_encryption_source: The Encryption Source for this Storage Account. Possible values are `Microsoft.Keyvault` and `Microsoft.Storage`. Defaults to `Microsoft.Storage`.
        :param pulumi.Input[str] account_kind: Defines the Kind of account. Valid options are `Storage`,
               `StorageV2` and `BlobStorage`. Changing this forces a new resource to be created.
               Defaults to `Storage`.
        :param pulumi.Input[str] account_replication_type: Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS` and `ZRS`.
        :param pulumi.Input[str] account_tier: Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. Changing this forces a new resource to be created
        :param pulumi.Input[dict] custom_domain: A `custom_domain` block as documented below.
        :param pulumi.Input[bool] enable_blob_encryption: Boolean flag which controls if Encryption Services are enabled for Blob storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
        :param pulumi.Input[bool] enable_file_encryption: Boolean flag which controls if Encryption Services are enabled for File storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
        :param pulumi.Input[bool] enable_https_traffic_only: Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
               for more information.
        :param pulumi.Input[dict] identity: A Managed Service Identity block as defined below.
        :param pulumi.Input[bool] is_hns_enabled: Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the
               resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The Custom Domain Name to use for the Storage Account, which will be validated by Azure.
        :param pulumi.Input[dict] network_rules: A `network_rules` block as documented below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the storage account. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_account.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['access_tier'] = access_tier

        __props__['account_encryption_source'] = account_encryption_source

        __props__['account_kind'] = account_kind

        if account_replication_type is None:
            raise TypeError("Missing required property 'account_replication_type'")
        __props__['account_replication_type'] = account_replication_type

        if account_tier is None:
            raise TypeError("Missing required property 'account_tier'")
        __props__['account_tier'] = account_tier

        __props__['account_type'] = account_type

        __props__['custom_domain'] = custom_domain

        __props__['enable_blob_encryption'] = enable_blob_encryption

        __props__['enable_file_encryption'] = enable_file_encryption

        __props__['enable_https_traffic_only'] = enable_https_traffic_only

        __props__['identity'] = identity

        __props__['is_hns_enabled'] = is_hns_enabled

        __props__['location'] = location

        __props__['name'] = name

        __props__['network_rules'] = network_rules

        if resource_group_name is None:
            raise TypeError("Missing required property 'resource_group_name'")
        __props__['resource_group_name'] = resource_group_name

        __props__['tags'] = tags

        __props__['primary_access_key'] = None
        __props__['primary_blob_connection_string'] = None
        __props__['primary_blob_endpoint'] = None
        __props__['primary_blob_host'] = None
        __props__['primary_connection_string'] = None
        __props__['primary_dfs_endpoint'] = None
        __props__['primary_dfs_host'] = None
        __props__['primary_file_endpoint'] = None
        __props__['primary_file_host'] = None
        __props__['primary_location'] = None
        __props__['primary_queue_endpoint'] = None
        __props__['primary_queue_host'] = None
        __props__['primary_table_endpoint'] = None
        __props__['primary_table_host'] = None
        __props__['primary_web_endpoint'] = None
        __props__['primary_web_host'] = None
        __props__['secondary_access_key'] = None
        __props__['secondary_blob_connection_string'] = None
        __props__['secondary_blob_endpoint'] = None
        __props__['secondary_blob_host'] = None
        __props__['secondary_connection_string'] = None
        __props__['secondary_dfs_endpoint'] = None
        __props__['secondary_dfs_host'] = None
        __props__['secondary_file_endpoint'] = None
        __props__['secondary_file_host'] = None
        __props__['secondary_location'] = None
        __props__['secondary_queue_endpoint'] = None
        __props__['secondary_queue_host'] = None
        __props__['secondary_table_endpoint'] = None
        __props__['secondary_table_host'] = None
        __props__['secondary_web_endpoint'] = None
        __props__['secondary_web_host'] = None

        super(Account, __self__).__init__(
            'azure:storage/account:Account',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

