package cmd

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-rds/cmd/dbclustercmd"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/spf13/cobra"
)

var AwsxRdsInstanceCmd = &cobra.Command{
	Use:   "getRdsInstanceList",
	Short: "getRdsInstanceList command gets RDS instance list",
	Long:  `getRdsInstanceList command gets RDS instance list of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			GetInstanceList(*clientAuth)
		} else {
			cmd.Help()
			return
		}

	},
}

func GetInstanceList(auth client.Auth) ([]*rds.DescribeDBInstancesOutput, error) {
	log.Println(" aws describe rds instance list count summary")

	rdsClient := client.GetClient(auth, client.RDS_CLIENT).(*rds.RDS)

	dbRequest := rds.DescribeDBInstancesInput{}

	dbClusterResponse, err := rdsClient.DescribeDBInstances(&dbRequest)

	if err != nil {
		log.Fatalln("Error:", err)
	}

	allInstances := []*rds.DescribeDBInstancesOutput{}

	for _, dbInstanceIdentifier := range dbClusterResponse.DBInstances {
		dbInstancesDetail := dbclustercmd.GetInstance(rdsClient, *dbInstanceIdentifier.DBInstanceArn)

		allInstances = append(allInstances, dbInstancesDetail)

	}
	log.Println(allInstances)
	return allInstances, err
}

func Execute() {
	err := AwsxRdsInstanceCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	AwsxRdsInstanceCmd.AddCommand(dbclustercmd.GetConfigDataCmd)

	AwsxRdsInstanceCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxRdsInstanceCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxRdsInstanceCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxRdsInstanceCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxRdsInstanceCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxRdsInstanceCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxRdsInstanceCmd.PersistentFlags().String("env", "", "aws key Required")
	AwsxRdsInstanceCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn Required")
	AwsxRdsInstanceCmd.PersistentFlags().String("externalId", "", "aws externalId Required")

}
