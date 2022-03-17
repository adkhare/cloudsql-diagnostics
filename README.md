# CloudSQL Diagnostics
CloudSQL Diagnostics for MySQL.

This tool is meant to collect diagnostics data from cloudsql similar to what pt-stalk collects from mysql.

With CloudSQL for MySQL, we have an ability to connect to mysql over public IP or via private IP. This tool relies on [CloudSQL GO Connector](https://github.com/googlecloudplatform/cloud-sql-go-connector) to establish connection with CloudSQL over public/private IP

## Using the library
1. Initialize the diagnostic tool by providing the driver name. This registers the protocol to be used for connecting CloudSQL
```
mysql := mysql.Mysql{
    Driver: "cloudsql-mysql",
}
err := mysql.Init(
    cloudsqlconn.WithDefaultDialOptions(
        cloudsqlconn.WithPrivateIP(),
    ),
)
```

2. Following is an example of the code to trigger diagnostics
```
dbConfig:= models.DbConnConfig{
				InstanceConnection: 	"my-project:my-region:my:instance",
				User:      			    "mysql-user",
				Password:  			    "mysql-password",
				DBName:				    "mysql-db",
				Driver: 			    "cloudsql-mysql",
			}

d := diagnostic.Diagnosis {
		DbConfig: &dbConfig,
	}

err := d.Diagnose(ctx)
if err != nil {
    return nil, err
}
```
