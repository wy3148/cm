package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"log"
)

var clientListCmd = &cobra.Command{
	Use:     "client",
	Short:   "see all lists of a given client",
	Long:    `all lists of a given client`,
	Example: "client --api_key your_key --client_id your_client_id",
	Run:     clientListCmdF,
}

func init() {
	clientListCmd.Flags().String("api_key", "", "your api key")
	clientListCmd.Flags().String("client_id", "", "your client id")
	rootCmd.AddCommand(clientListCmd)
}

func clientListCmdF(cmd *cobra.Command, args []string) {

	var res gorequest.Response
	var body string

	api, _ := cmd.Flags().GetString("api_key")

	if len(api) == 0 {
		log.Println("api key is needed")
	}

	clientId, _ := cmd.Flags().GetString("client_id")

	if len(api) == 0 {
		log.Println("client_id is needed")
	}

	request := gorequest.New()
	super := request.SetBasicAuth(api, "x")

	res, body, _ = super.Get(cmUrl + "/clients/" + clientId + ".json").End()

	log.Println("response code:", res.StatusCode)

	if len(body) > 0 {
		var prettyJSON bytes.Buffer
		error := json.Indent(&prettyJSON, []byte(body), "", " ")
		if error != nil {
			log.Println("JSON parse error: ", error)
			return
		}

		log.Println("body response:\n", string(prettyJSON.Bytes()))
	}

	return
}
