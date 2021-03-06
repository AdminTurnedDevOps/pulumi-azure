# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Database(pulumi.CustomResource):
    auto_pause_delay_in_minutes: pulumi.Output[float]
    """
    Time in minutes after which database is automatically paused. A value of `-1` means that automatic pause is disabled. This property is only settable for General Purpose Serverless databases.
    """
    collation: pulumi.Output[str]
    """
    Specifies the collation of the database. Changing this forces a new resource to be created.
    """
    create_mode: pulumi.Output[str]
    """
    The create mode of the database. Possible values are `Copy`, `Default`, `OnlineSecondary`, `PointInTimeRestore`, `Restore`, `RestoreExternalBackup`, `RestoreExternalBackupSecondary`, `RestoreLongTermRetentionBackup` and `Secondary`. 
    """
    creation_source_database_id: pulumi.Output[str]
    """
    The id of the source database to be referred to create the new database. This should only be used for databases with `create_mode` values that use another database as reference. Changing this forces a new resource to be created.
    """
    elastic_pool_id: pulumi.Output[str]
    """
    Specifies the ID of the elastic pool containing this database. Changing this forces a new resource to be created.
    """
    license_type: pulumi.Output[str]
    """
    Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
    """
    max_size_gb: pulumi.Output[float]
    """
    The max size of the database in gigabytes. 
    """
    min_capacity: pulumi.Output[float]
    """
    Minimal capacity that database will always have allocated, if not paused. This property is only settable for General Purpose Serverless databases.
    """
    name: pulumi.Output[str]
    """
    The name of the Ms SQL Database. Changing this forces a new resource to be created.
    """
    read_replica_count: pulumi.Output[float]
    """
    The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
    """
    read_scale: pulumi.Output[bool]
    """
    If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
    """
    restore_point_in_time: pulumi.Output[str]
    """
    Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. This property is only settable for `create_mode`= `PointInTimeRestore`  databases.
    """
    sample_name: pulumi.Output[str]
    """
    Specifies the name of the sample schema to apply when creating this database. Possible value is `AdventureWorksLT`.
    """
    server_id: pulumi.Output[str]
    """
    The id of the Ms SQL Server on which to create the database. Changing this forces a new resource to be created.
    """
    sku_name: pulumi.Output[str]
    """
    Specifies the name of the sku used by the database. Changing this forces a new resource to be created. For example, `GP_S_Gen5_2`,`HS_Gen4_1`,`BC_Gen5_2`, `ElasticPool`, `Basic`,`S0`, `P2` ,`DW100c`, `DS100`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    zone_redundant: pulumi.Output[bool]
    """
    Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones. This property is only settable for Premium and Business Critical databases.
    """
    def __init__(__self__, resource_name, opts=None, auto_pause_delay_in_minutes=None, collation=None, create_mode=None, creation_source_database_id=None, elastic_pool_id=None, license_type=None, max_size_gb=None, min_capacity=None, name=None, read_replica_count=None, read_scale=None, restore_point_in_time=None, sample_name=None, server_id=None, sku_name=None, tags=None, zone_redundant=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a MS SQL Database.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] auto_pause_delay_in_minutes: Time in minutes after which database is automatically paused. A value of `-1` means that automatic pause is disabled. This property is only settable for General Purpose Serverless databases.
        :param pulumi.Input[str] collation: Specifies the collation of the database. Changing this forces a new resource to be created.
        :param pulumi.Input[str] create_mode: The create mode of the database. Possible values are `Copy`, `Default`, `OnlineSecondary`, `PointInTimeRestore`, `Restore`, `RestoreExternalBackup`, `RestoreExternalBackupSecondary`, `RestoreLongTermRetentionBackup` and `Secondary`. 
        :param pulumi.Input[str] creation_source_database_id: The id of the source database to be referred to create the new database. This should only be used for databases with `create_mode` values that use another database as reference. Changing this forces a new resource to be created.
        :param pulumi.Input[str] elastic_pool_id: Specifies the ID of the elastic pool containing this database. Changing this forces a new resource to be created.
        :param pulumi.Input[str] license_type: Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
        :param pulumi.Input[float] max_size_gb: The max size of the database in gigabytes. 
        :param pulumi.Input[float] min_capacity: Minimal capacity that database will always have allocated, if not paused. This property is only settable for General Purpose Serverless databases.
        :param pulumi.Input[str] name: The name of the Ms SQL Database. Changing this forces a new resource to be created.
        :param pulumi.Input[float] read_replica_count: The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
        :param pulumi.Input[bool] read_scale: If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
        :param pulumi.Input[str] restore_point_in_time: Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. This property is only settable for `create_mode`= `PointInTimeRestore`  databases.
        :param pulumi.Input[str] sample_name: Specifies the name of the sample schema to apply when creating this database. Possible value is `AdventureWorksLT`.
        :param pulumi.Input[str] server_id: The id of the Ms SQL Server on which to create the database. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: Specifies the name of the sku used by the database. Changing this forces a new resource to be created. For example, `GP_S_Gen5_2`,`HS_Gen4_1`,`BC_Gen5_2`, `ElasticPool`, `Basic`,`S0`, `P2` ,`DW100c`, `DS100`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] zone_redundant: Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones. This property is only settable for Premium and Business Critical databases.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['auto_pause_delay_in_minutes'] = auto_pause_delay_in_minutes
            __props__['collation'] = collation
            __props__['create_mode'] = create_mode
            __props__['creation_source_database_id'] = creation_source_database_id
            __props__['elastic_pool_id'] = elastic_pool_id
            __props__['license_type'] = license_type
            __props__['max_size_gb'] = max_size_gb
            __props__['min_capacity'] = min_capacity
            __props__['name'] = name
            __props__['read_replica_count'] = read_replica_count
            __props__['read_scale'] = read_scale
            __props__['restore_point_in_time'] = restore_point_in_time
            __props__['sample_name'] = sample_name
            if server_id is None:
                raise TypeError("Missing required property 'server_id'")
            __props__['server_id'] = server_id
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            __props__['zone_redundant'] = zone_redundant
        super(Database, __self__).__init__(
            'azure:mssql/database:Database',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auto_pause_delay_in_minutes=None, collation=None, create_mode=None, creation_source_database_id=None, elastic_pool_id=None, license_type=None, max_size_gb=None, min_capacity=None, name=None, read_replica_count=None, read_scale=None, restore_point_in_time=None, sample_name=None, server_id=None, sku_name=None, tags=None, zone_redundant=None):
        """
        Get an existing Database resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] auto_pause_delay_in_minutes: Time in minutes after which database is automatically paused. A value of `-1` means that automatic pause is disabled. This property is only settable for General Purpose Serverless databases.
        :param pulumi.Input[str] collation: Specifies the collation of the database. Changing this forces a new resource to be created.
        :param pulumi.Input[str] create_mode: The create mode of the database. Possible values are `Copy`, `Default`, `OnlineSecondary`, `PointInTimeRestore`, `Restore`, `RestoreExternalBackup`, `RestoreExternalBackupSecondary`, `RestoreLongTermRetentionBackup` and `Secondary`. 
        :param pulumi.Input[str] creation_source_database_id: The id of the source database to be referred to create the new database. This should only be used for databases with `create_mode` values that use another database as reference. Changing this forces a new resource to be created.
        :param pulumi.Input[str] elastic_pool_id: Specifies the ID of the elastic pool containing this database. Changing this forces a new resource to be created.
        :param pulumi.Input[str] license_type: Specifies the license type applied to this database. Possible values are `LicenseIncluded` and `BasePrice`.
        :param pulumi.Input[float] max_size_gb: The max size of the database in gigabytes. 
        :param pulumi.Input[float] min_capacity: Minimal capacity that database will always have allocated, if not paused. This property is only settable for General Purpose Serverless databases.
        :param pulumi.Input[str] name: The name of the Ms SQL Database. Changing this forces a new resource to be created.
        :param pulumi.Input[float] read_replica_count: The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
        :param pulumi.Input[bool] read_scale: If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
        :param pulumi.Input[str] restore_point_in_time: Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. This property is only settable for `create_mode`= `PointInTimeRestore`  databases.
        :param pulumi.Input[str] sample_name: Specifies the name of the sample schema to apply when creating this database. Possible value is `AdventureWorksLT`.
        :param pulumi.Input[str] server_id: The id of the Ms SQL Server on which to create the database. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sku_name: Specifies the name of the sku used by the database. Changing this forces a new resource to be created. For example, `GP_S_Gen5_2`,`HS_Gen4_1`,`BC_Gen5_2`, `ElasticPool`, `Basic`,`S0`, `P2` ,`DW100c`, `DS100`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] zone_redundant: Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones. This property is only settable for Premium and Business Critical databases.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_pause_delay_in_minutes"] = auto_pause_delay_in_minutes
        __props__["collation"] = collation
        __props__["create_mode"] = create_mode
        __props__["creation_source_database_id"] = creation_source_database_id
        __props__["elastic_pool_id"] = elastic_pool_id
        __props__["license_type"] = license_type
        __props__["max_size_gb"] = max_size_gb
        __props__["min_capacity"] = min_capacity
        __props__["name"] = name
        __props__["read_replica_count"] = read_replica_count
        __props__["read_scale"] = read_scale
        __props__["restore_point_in_time"] = restore_point_in_time
        __props__["sample_name"] = sample_name
        __props__["server_id"] = server_id
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        __props__["zone_redundant"] = zone_redundant
        return Database(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

