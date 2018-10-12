package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/urfave/cli"
)

func main() {
	// Set up our CLI utilities
	app := cli.NewApp()
	app.Name = "go-stack-my-cloud"
	app.Usage = "scan a directory for CloudFormation templates and then do them"
	app.Action = func(c *cli.Context) error {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}

		// Create a session to be used for AWS activities
		sess := session.Must(session.NewSession())

		// Create a client to be used with CloudFormation
		cfClient := cloudformation.New(sess)
		// Perform the action on all elements of the list at the same time
		for _, file := range files {
			go func(file os.FileInfo) {
				ext := filepath.Ext(file.Name())
				stackName := file.Name()[0 : len(file.Name())-len(ext)]

				// Don't just print the extension. Do something with it
				// If the file extension is yaml, yml, or json, upload it to CF

				// Check to see if the file name (sans extension) is already a stack
				describeStacksInput := &cloudformation.DescribeStacksInput{
					StackName: aws.String(stackName),
				}
				stackDescription, err := cfClient.DescribeStacks(describeStacksInput)
				if err, ok := err.(*cloudformation.AmazonCloudFormationException);

				// We want an AmazonCloudformationException
				ok {
					// Stack does not exist. Go ahead and create that stack
					createStackInput := &cloudformation.CreateStackInput{}
					cfClient.CreateStack()
				}
				fmt.Println(ext)
			}(file)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
