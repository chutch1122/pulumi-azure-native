// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20150801.Inputs
{

    /// <summary>
    /// Note: properties are serialized in JSON format and stored in DB. 
    ///             if new properties are added they might not be in the previous data rows 
    ///             so please handle nulls
    /// </summary>
    public sealed class DatabaseBackupSetting : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Contains a connection string to a database which is being backed up/restored. If the restore should happen to a new database, the database name inside is the new one.
        /// </summary>
        [Input("connectionString")]
        public string? ConnectionString { get; set; }

        /// <summary>
        /// Contains a connection string name that is linked to the SiteConfig.ConnectionStrings.
        ///             This is used during restore with overwrite connection strings options.
        /// </summary>
        [Input("connectionStringName")]
        public string? ConnectionStringName { get; set; }

        /// <summary>
        /// SqlAzure / MySql
        /// </summary>
        [Input("databaseType")]
        public string? DatabaseType { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public DatabaseBackupSetting()
        {
        }
        public static new DatabaseBackupSetting Empty => new DatabaseBackupSetting();
    }
}
