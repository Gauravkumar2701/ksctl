package cmd

/*
Kubesimplify
@maintainer: 	Dipankar Das <dipankardas0115@gmail.com>
				Anurag Kumar <contact.anurag7@gmail.com>
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	util "github.com/kubesimplify/ksctl/api/utils"
	"github.com/spf13/cobra"
)

type printer struct {
	ClusterName string `json:"cluster_name"`
	Region      string `json:"region"`
	Provider    string `json:"provider"`
}

const (
	ALL    = int(0)
	CIVOC  = int(1)
	LOCALC = int(2)
)

func printUtil(cargo []byte) {
	//TODO: Added Table type display
	fmt.Println(string(cargo))
}

// Printer TODO: Move the contents of get clusters to their respective api files
func Printer(i int) {
	var toBePrinted []printer

	folders, err := os.ReadDir(util.GetPath(util.CLUSTER_PATH, "civo", "managed"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range folders {
		if file.IsDir() {
			info := strings.Split(file.Name(), " ")
			toBePrinted = append(toBePrinted, printer{ClusterName: info[0], Region: info[1], Provider: "CIVO (MANAGED)"})
		}
	}

	folders, err = os.ReadDir(util.GetPath(util.CLUSTER_PATH, "civo", "ha"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range folders {
		if file.IsDir() {
			info := strings.Split(file.Name(), " ")
			toBePrinted = append(toBePrinted, printer{ClusterName: info[0], Region: info[1], Provider: "CIVO (HA)"})
		}
	}

	folders, err = os.ReadDir(util.GetPath(util.CLUSTER_PATH, "local"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range folders {
		if file.IsDir() {
			toBePrinted = append(toBePrinted, printer{ClusterName: file.Name(), Region: "N/A", Provider: "local"})
		}
	}

	arr, err := json.MarshalIndent(toBePrinted, "", "  ")
	if err != nil {
		panic(fmt.Errorf("JSON Convertion failed"))
	}
	printUtil(arr)
}

// viewClusterCmd represents the viewCluster command
var getClusterCmd = &cobra.Command{
	Use:     "get-clusters",
	Aliases: []string{"get"},
	Short:   "Use to get clusters",
	Long: `It is used to view clusters. For example:

ksctl get-clusters `,
	Run: func(cmd *cobra.Command, args []string) {
		Printer(ALL)
		//Printer(CIVOC)
		//Printer(LOCALC)
	},
}

func init() {
	rootCmd.AddCommand(getClusterCmd)
}
