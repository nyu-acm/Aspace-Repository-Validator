package main

import (
	"bufio"
	"fmt"
	"github.com/nyudlts/go-aspace"
	"os"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	repoId := 2
	client, err := aspace.NewClient("dev", 20)
	handleErr(err)

	resourceIDs, err := client.GetResourceIDsByRepository(repoId)
	handleErr(err)

	outputFile, err := os.Create("failures.txt")
	handleErr(err)
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)

	for _, resourceId := range resourceIDs {
		uri := fmt.Sprintf("/repositories/%d/resources/%d", repoId, resourceId)
		fmt.Println("Validating: ", uri)
		ead, err := client.SerializeEAD(repoId, resourceId, true, false, false, false, false)
		handleErr(err)
		err = aspace.ValidateEAD(ead); if err != nil {
			writer.WriteString(uri + "\n")
			writer.Flush()
		}
	}

}

