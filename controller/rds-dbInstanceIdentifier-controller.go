package controller

import (
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-rds/cmd"
	"github.com/aws/aws-sdk-go/service/rds"
)

func GetRdsDbInstanceIdentifierByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) ([]*rds.DescribeDBInstancesOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetRdsDbInstanceIdentifiersByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetRdsDbInstanceIdentifierByUserCreds(region string, accesskey string, secretKey string, crossAccountRoleArn string, externalId string) ([]*rds.DescribeDBInstancesOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accesskey, secretKey, crossAccountRoleArn, externalId)
	return GetRdsDbInstanceIdentifiersByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetRdsDbInstanceIdentifiersByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) ([]*rds.DescribeDBInstancesOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := cmd.GetInstanceList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetRdsDbInstanceIdentifiers(clientAuth *client.Auth) ([]*rds.DescribeDBInstancesOutput, error) {
	response, err := cmd.GetInstanceList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
