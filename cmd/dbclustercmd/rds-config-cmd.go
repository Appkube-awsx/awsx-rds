package dbclustercmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
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

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			dbInstanceIdentifier, _ := cmd.Flags().GetString("dbInstanceIdentifier")
			if dbInstanceIdentifier != "" {
				getInstanceDetails(dbInstanceIdentifier, *clientAuth)
			} else {
				log.Fatalln("cluster name not provided. program exit")
			}
		}
	},
}

func getInstanceDetails(dbInstanceIdentifier string, auth client.Auth) *rds.DescribeDBInstancesOutput {
	log.Println("Getting aws db Instance data")

	listInstanceClient := client.GetClient(auth, client.RDS_CLIENT).(*rds.RDS)

	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(dbInstanceIdentifier),
	}

	instanceDetailsResponse, err := listInstanceClient.DescribeDBInstances(input)
	log.Println(instanceDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return instanceDetailsResponse
}

func GetInstance(rdsClient *rds.RDS, dbInstanceIdentifier string) *rds.DescribeDBInstancesOutput {
	log.Println("Getting aws instance detail for instance: ", dbInstanceIdentifier)

	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(dbInstanceIdentifier),
	}
	clusterDetailsResponse, err := rdsClient.DescribeDBInstances(input)
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
