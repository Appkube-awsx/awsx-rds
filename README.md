# RDS CLi's

## To list all the RDS dbinstances,run the following command:

```bash
awsx-rds --zone <zone> --acccessKey <acccessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --externalId <externalId> 
```

## To retrieve the configuration details of a specific dbInstancecmd, run the following command:

```bash
awsx-rds getConfigData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --dbInstanceIdentifier <dbInstanceIdentifier>
```


