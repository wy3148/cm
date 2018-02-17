package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list command",
}

var listCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "create a list",
	Long:    `this is used to create a list for a client`,
	Example: "list create ",
	Run:     listCreateCmdF,
}

func init() {
	listCreateCmd.Flags().String("api_key", "", "your api key")
	listCreateCmd.Flags().String("client_id", "", "your client id")
	listCreateCmd.Flags().String("json_file", "", "the file include the list details, /directory/list.json")
	listCmd.AddCommand(listCreateCmd)

	rootCmd.AddCommand(listCmd)
}

func listCreateCmdF(cmd *cobra.Command, args []string) {

	var res gorequest.Response
	var body string

	api, _ := cmd.Flags().GetString("api_key")

	if len(api) == 0 {
		log.Println("api key is needed")
		return
	}

	clientId, _ := cmd.Flags().GetString("client_id")

	if len(api) == 0 {
		log.Println("client_id is needed")
		return
	}

	jsonFile, _ := cmd.Flags().GetString("json_file")

	if len(jsonFile) == 0 {
		log.Println("json_file is needed")
		return
	}

	data, err := ioutil.ReadFile(jsonFile)

	if err != nil {
		log.Println("failed to read file:", jsonFile)
		return
	}

	if len(data) == 0 {
		log.Println("json file is empty:", jsonFile)
		return
	}

	log.Println("will create a new list:", string(data))

	request := gorequest.New()
	request.SetBasicAuth(api, "x")
	request.Header.Set("Content-Type", "application/json")

	res, body, _ = request.Post(cmUrl + "/lists/" + clientId + ".json").Send(string(data)).End()

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
