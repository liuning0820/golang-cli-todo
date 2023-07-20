package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

)

const TASK_SERVER = "https://dummyjson.com/todos"



type TaskCreateResp struct {
	Todo       string `json:"todo,omitempty"`
	Completed bool `json:"completed,omitempty"`
	UserId       string `json:"userId,omitempty"`
	Id string `json:"id,omitempty"`
}

type Task struct {
	Todo       string `json:"todo,omitempty"`
	Completed bool `json:"completed,omitempty"`
	UserId       string `json:"userId,omitempty"`
}



func ReadAPI() []Task {
	response, err := http.Get(TASK_SERVER + "/")

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	tasks := []Task{}
	_ = json.Unmarshal(responseData, &tasks)

	println(len(tasks))

	return tasks
}

func WriteAPI(task Task) TaskCreateResp {
	url := TASK_SERVER + "/add"
	taskJson, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}


	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(taskJson)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	responseData, _ := ioutil.ReadAll(resp.Body)

	println(url)
	print(resp.Status)


	taskCreateResp := TaskCreateResp{}
	_ = json.Unmarshal(responseData, &taskCreateResp)

	return taskCreateResp
}



