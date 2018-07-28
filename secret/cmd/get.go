package cmd

import (
	"fmt"
	"secret/crypt"
	"secret/fileHandle"

	"github.com/spf13/cobra"
)

//variable defines the help template for get subcommand.
var getTemplate = `fetch the key related to the service name

Usage:
  get [flags] [service_name]

Flags:
  -h, --help         help for secret
      --key string   encryption key (required)

Arguments:
service_name	name of the service of which the key is to be saved` + "\n\n\n"

//fetch key func is the operation carried out when get command is accessed
//with correct arguments.
func fetchKey(serviceName string) {
	hexSecretKey, err := fileHandle.GetSecret(serviceName)
	if err != nil {
		fmt.Println("Following error occured during fetching the secret key : ", err)
		return
	}
	if len(hexSecretKey) == 0 {
		fmt.Println("Secret Key for the service not found")
		return
	}
	secretKey, err := crypt.Decrypt(key, hexSecretKey)
	if err != nil {
		fmt.Println("Following error occured during fetching secret key : ", err)
		return
	}
	fmt.Println("Secret Key : ", secretKey)

}

//the variable below defines the sub-command to fetch key
//related to a service name
var getCmd = &cobra.Command{

	Use:   "get",
	Short: "fetch the secret key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 || len(key) == 0 {
			fmt.Println(" invalid input ")
			return
		}
		serviceName := args[0]
		fetchKey(serviceName)
	},
}

func init() {
	getCmd.SetUsageTemplate(getTemplate)
	MainCmd.AddCommand(getCmd)
}