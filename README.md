# RDS CLi's

## To list all the RDS dbinstances,run the following command:

```bash
awsx-rds --zone <zone> --acccessKey <acccessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --externalId <externalId> --env <env>
```

## To retrieve the configuration details of a specific dbInstancecmd, run the following command:

```bash
awsx-rds getConfigData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env> --dbInstanceIdentifier <dbInstanceIdentifier>
```

## To retrieve the cost details of a specific RDS dbInstancecmd, run the following command:

```bash
awsx-rds getCostData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env>
```

## To retrieve the cost Spikes of a specific RDS dbInstancecmd, run the following command:

```bash
awsx-rds GetCostSpike -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env> --granularity <granularity> --startDate <startDate> --endDate <endDate> 
```
