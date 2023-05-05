package dbInstancecmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-rds/authenticator"
	"github.com/Appkube-awsx/awsx-rds/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
		// print(authFlag)
		// authFlag := true
		if authFlag {
			dbInstanceIdentifier, _ := cmd.Flags().GetString("dbInstanceIdentifier")
			if dbInstanceIdentifier != "" {
				getInstanceDetails(region, crossAccountRoleArn, acKey, secKey, externalId, dbInstanceIdentifier)
			} else {
				log.Fatalln("DB Instance Name not provided. Program exit")
			}
		}
	},
}

func getInstanceDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string, dbInstanceIdentifier string) *rds.DescribeDBInstancesOutput {
	log.Println("Getting aws db Instance data")
	listInstanceClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(dbInstanceIdentifier),
	}
	InstanceDetailsResponse, err := listInstanceClient.DescribeDBInstances(input)
	log.Println(InstanceDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return InstanceDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("dbInstanceIdentifier", "t", "", "DB instance name")

	if err := GetConfigDataCmd.MarkFlagRequired("dbInstanceIdentifier"); err != nil {
		fmt.Println(err)
	}
}
