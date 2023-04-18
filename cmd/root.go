package cmd

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-rds/authenticator"
	"github.com/Appkube-awsx/awsx-rds/client"
	"github.com/Appkube-awsx/awsx-rds/cmd/dbInstancecmd"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/spf13/cobra"
)

var AwsxdbClusterMetadataCmd = &cobra.Command{
	Use:   "DescribeRDSMetdataDetails",
	Short: "DescribeRDSMetdataDetails command gets resource counts",
	Long:  `DescribeRDSMetdataDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command getElementDetails started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		env := cmd.PersistentFlags().Lookup("env").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, env, externalId)

		if authFlag {
			getRDSResourceList(region, crossAccountRoleArn, acKey, secKey, env, externalId)
		}

	},
}

func getRDSResourceList(region string, crossAccountRoleArn string, accessKey string, secretKey string, env string, externalId string) (*rds.DescribeDBInstancesOutput, error) {
	log.Println(" aws describe rds instance metadata count summary")
	dbclient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	dbRequest := rds.DescribeDBInstancesInput{}
	dbclusterResponse, err := dbclient.DescribeDBInstances(&dbRequest)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	log.Println(dbclusterResponse)
	return dbclusterResponse, err
}

func Execute() {
	err := AwsxdbClusterMetadataCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	AwsxdbClusterMetadataCmd.AddCommand(dbInstancecmd.GetConfigDataCmd)
	AwsxdbClusterMetadataCmd.AddCommand(dbInstancecmd.GetCostDataCmd)
	AwsxdbClusterMetadataCmd.AddCommand(dbInstancecmd.GetCostSpikeCmd)

	AwsxdbClusterMetadataCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxdbClusterMetadataCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxdbClusterMetadataCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxdbClusterMetadataCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxdbClusterMetadataCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxdbClusterMetadataCmd.PersistentFlags().String("env", "", "aws key Required")
	AwsxdbClusterMetadataCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn Required")
	AwsxdbClusterMetadataCmd.PersistentFlags().String("externalId", "", "aws externalId Required")

}
