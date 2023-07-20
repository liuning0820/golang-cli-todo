/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/liuning0820/golang-cli-todo/utils"

	"github.com/spf13/cobra"

	"encoding/json"

	"io/ioutil"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a todo",
	Long:  `This command will create todo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		taskTitle, _ := cmd.Flags().GetString("todo")
		taskStatus, _ := cmd.Flags().GetBool("completed")
		task := utils.Task{
			Todo:       taskTitle,
			Completed: taskStatus,
			UserId: "1",
		}
		fmt.Printf("Creating task %+v\n", task)

		content, err := json.Marshal(task)

		if err != nil {
			fmt.Println(err)
		}

		err = ioutil.WriteFile("task.json", content, 0644)
		if err != nil {
			fmt.Println(err)
		}



		// Storing task in backend calling my-todos REST API
		resp := utils.WriteAPI(task)
		fmt.Println("Task created with ID:", resp)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.Flags().StringP("todo", "t", "", "specify task title")
	createCmd.Flags().BoolP("completed", "c", false, "specify task status")
	createCmd.Flags().Int32("userId",'u',"specify user id")
}
