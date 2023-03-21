package dbInstancecmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-rds/authenticator"
	"github.com/Appkube-awsx/awsx-rds/client"

	// "github.com/aws/aws-sdk-go/aws"
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
		env := cmd.Parent().PersistentFlags().Lookup("env").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String() 
		
		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, env, externalId)
		print(authFlag)
		// authFlag := true
		if authFlag {
			dbInstanceIdentifier, _ := cmd.Flags().GetString("dbInstanceIdentifier")
			if(dbInstanceIdentifier != "") {
				getClusterDetails(region, crossAccountRoleArn, acKey, secKey, env, externalId,dbInstanceIdentifier)
			} else {
				log.Fatalln("DB Instance Name not provided. Program exit")
			}
		}
	}, 
} 

func getClusterDetails(region string, crossAccountRoleArn string,accessKey string, secretKey string, env string, externalId string,dbInstanceIdentifier string) *rds.DescribeDBInstancesOutput{
	log.Println("Getting aws db Instance data")
	 listClusterClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &rds.DescribeDBInstancesInput{
			DBInstanceIdentifier: aws.String(dbInstanceIdentifier),
			}
	clusterDetailsResponse, err := listClusterClient.DescribeDBInstances(input)
	log.Println(clusterDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return clusterDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("dbInstanceIdentifier", "t", "", "DB instance name")

	if err := GetConfigDataCmd.MarkFlagRequired("dbInstanceIdentifier"); err != nil {
		fmt.Println(err)
	}
}