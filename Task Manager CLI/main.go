package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	TO_DO = iota
	IN_PROGRESS
	DONE
)

var progressState = map[int]string{
	TO_DO:       "Todo",
	IN_PROGRESS: "In progress",
	DONE:        "Done",
}

const (
	LOW = iota
	MEDIUM
	HIGH
)

var priorityState = map[int]string{
	LOW:    "Low",
	MEDIUM: "Medium",
	HIGH:   "High",
}

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	DueDate     int64  `json:"due_date"`
	IsCompleted bool   `json:"is_completed"`
	Metadata    Metadata
}

type Metadata struct {
	TaskID        int   `json:"task_id"`
	CreatedAt     int64 `json:"created_at"`
	LastUpdatedAt int64 `json:"last_updated_at"`
	CompletedAt   int64 `json:"completed_at"`
	IsDeleted     bool  `json:"is_deleted"`
}

type tasks []Task

var jsonTasks tasks

func ReadInputFromUser(userInput string) (string, error) {
	fmt.Print()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(userInput)
	userString, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	userString = strings.TrimSpace(userString)
	return userString, nil
}

func (t *tasks) AddTask() error {
	fmt.Println("<-New Task Information->")
	titleString, err := ReadInputFromUser("Enter the title: ")
	if err != nil {
		return err
	}
	descriptionString, err := ReadInputFromUser("Enter the description: ")
	if err != nil {
		return err
	}
	priorityString, err := ReadInputFromUser("Enter the priority type\n1.LOW\n2.MEDIUM\n3.HIGH\nEnter the option: ")
	if err != nil {
		return err
	}
	var priorityValue string
	if priorityInt, err := strconv.Atoi(priorityString); err != nil {
		return err
	} else {
		priorityValue = priorityState[priorityInt-1]
	}
	dateFormat := "02-01-2006"
	dateString, err := ReadInputFromUser("Enter the Due Date(DD-MM-YYYY): ")
	if err != nil {
		return err
	}
	dueDateTime, err := time.Parse(dateFormat, dateString)
	if err != nil {
		return err
	}
	metadata := Metadata{
		TaskID:        rand.IntN(1000),
		CreatedAt:     time.Now().Unix(),
		LastUpdatedAt: time.Now().Unix(),
		CompletedAt:   time.Now().Unix(),
		IsDeleted:     false,
	}
	newTask := Task{
		Title:       titleString,
		Description: descriptionString,
		Status:      progressState[0],
		Priority:    priorityValue,
		DueDate:     dueDateTime.Unix(),
		IsCompleted: false,
		Metadata:    metadata,
	}
	*t = append(*t, newTask)
	fmt.Print("Added Task successfully \n\n")
	return nil
}

func (t *tasks) UpdateTask(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	taskInd := -1
	for ind, task := range *t {
		if task.Metadata.TaskID == intId {
			taskInd = ind
		}
	}
	if taskInd == -1 {
		return errors.New("No Task with given Id")
	}
	titleString, err := ReadInputFromUser("Enter new title(Enter to skip):")
	if err != nil {
		return err
	}
	duedateString, err := ReadInputFromUser("Enter new due date(DD-MM-YYYY)(Enter to skip):")
	if err != nil {
		return err
	}
	var dueDateTime time.Time
	dateFormat := "02-01-2006"
	if duedateString != "" {
		dueDateTime, err = time.Parse(dateFormat, duedateString)
		if err != nil {
			return err
		}
	}
	if duedateString == "" && titleString == "" {
		return errors.New("No new update to make.")
	}
	task := (*t)[taskInd]
	if titleString != "" {
		task.Title = titleString
	}
	if duedateString != "" {
		task.DueDate = dueDateTime.Unix()
	}
	//update metadata
	task.Metadata.LastUpdatedAt = time.Now().Unix()
	(*t)[taskInd] = task
	fmt.Print("Updated Successfully!!\n\n")
	return nil
}

func (t *tasks) UpdateTaskStatus(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	taskInd := -1
	for ind, task := range *t {
		if task.Metadata.TaskID == intId {
			taskInd = ind
		}
	}
	if taskInd == -1 {
		return errors.New("No Task with given Id")
	}
	statusString, err := ReadInputFromUser("Enter the current status\n1.TODO\n2.IN_PROGRESS\n3.DONE\nEnter the option(Enter to skip): ")
	if err != nil {
		return err
	}
	var statusValue string
	if statusInt, err := strconv.Atoi(statusString); err != nil {
		return err
	} else {
		statusValue = progressState[statusInt-1]
	}
	if strings.ToUpper(statusValue) == "DONE" {
		(*t)[taskInd].IsCompleted = true
		(*t)[taskInd].Metadata.CompletedAt = time.Now().Unix()
	}
	(*t)[taskInd].Status = statusValue
	(*t)[taskInd].Metadata.LastUpdatedAt = time.Now().Unix()
	fmt.Print("Updated task status successfully!! \n\n")
	return nil
}

func (t *tasks) DeleteTask(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	taskInd := -1
	for ind, task := range *t {
		if task.Metadata.TaskID == intId {
			taskInd = ind
		}
	}
	if taskInd == -1 {
		return errors.New("No Task with given Id")
	}
	(*t)[taskInd].Metadata.IsDeleted = true
	(*t)[taskInd].Metadata.LastUpdatedAt = time.Now().Unix()
	return nil
}

func (t *tasks) ShowDueTasks() error {
	taskCount := 0
	for _, task := range *t {
		elaspedTime := time.Since(time.Unix(task.DueDate, 0))
		if elaspedTime.Seconds() > 0 {
			fmt.Printf("---Task %d---\n", taskCount+1)
			fmt.Printf("Title: %s\n", task.Title)
			fmt.Printf("Description: %s\n", task.Description)
			fmt.Printf("Status: %s\n", task.Status)
			fmt.Printf("Priority: %s\n", task.Priority)
			fmt.Printf("DueDate: %v\n\n", time.Unix(task.DueDate, 0))
			taskCount++
		}
	}
	return nil
}

func (t *tasks) ShowAllTasks() {
	fmt.Printf("\n")
	for i, task := range *t {
		if !task.Metadata.IsDeleted {
			fmt.Printf("---Task %d---\n", i+1)
			fmt.Printf("Title: %s\n", task.Title)
			fmt.Printf("Description: %s\n", task.Description)
			fmt.Printf("Status: %s\n", task.Status)
			fmt.Printf("Priority: %s\n", task.Priority)
			fmt.Printf("DueDate: %v\n\n", time.Unix(task.DueDate, 0))
		}
	}
}

func processUserOpt(opt string) error {
	switch opt {
	case "1":
		if err := jsonTasks.AddTask(); err != nil {
			return err
		}
	case "2":
		idString, err := ReadInputFromUser("Enter the task id to update: ")
		if err != nil {
			return err
		}
		jsonTasks.UpdateTask(idString)
	case "3":
		idString, err := ReadInputFromUser("Enter the task id to update status: ")
		if err != nil {
			return err
		}
		jsonTasks.UpdateTaskStatus(idString)
	case "4":
		idString, err := ReadInputFromUser("Enter the task id to delete: ")
		if err != nil {
			return err
		}
		jsonTasks.DeleteTask(idString)
	case "5":
		jsonTasks.ShowAllTasks()
	}
	return nil
}

func main() {
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("      				     	TASK MANAGER CLI										")
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	jsonFile, err := os.Open("./tasks.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(jsonFile)
	if err = decoder.Decode(&jsonTasks); err != nil {
		// Handle empty file or invalid JSON
		if err.Error() != "EOF" {
			panic(err)
		}
	}
	jsonFile.Close()

	for {
		fmt.Println("Operations available:")
		fmt.Print("\n1. Add Task\n2. Update Task Info\n3. Update Task Status\n4. Delete Task\n5. View All Tasks\n6. Exit\n\n")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a option: ")
		userOpt, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		userOptClean := strings.TrimSpace(userOpt)
		if userOptClean == "6" {
			break
		}
		if err = processUserOpt(userOptClean); err != nil {
			panic(err)
		}
		jsonData, err := json.MarshalIndent(jsonTasks, "", " ")
		if err != nil {
			panic(err)
		}
		if err := os.WriteFile("./tasks.json", jsonData, 0644); err != nil {
			panic(err)
		}
	}
}
