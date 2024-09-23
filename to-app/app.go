package main

import (
	"errors"
	"fmt"
)

type Task struct {
	id int
	description string
	completed bool
}

var TaskList []Task;
var taskIDCounter int = 1;

func enterTask(description string, completed bool) {
	newTask := Task{
		id: taskIDCounter,
		description: description,
		completed: completed,
	}
	TaskList = append(TaskList, newTask);
	taskIDCounter++;
}

func findTask(numOfTask int) (Task, error) {
	if numOfTask < 0 || numOfTask > len(TaskList) {
		return Task{}, errors.New("The number of task is invalid");
	}

	for _, task := range TaskList {
		if task.id == numOfTask {
			return task, nil;
		}
	}
	return Task{}, errors.New("Task not found");
}


func main() {
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("This is the %d task", i+1);
		enterTask(str, false);
	}

	task, err := findTask(1);
	if err != nil {
		fmt.Println(err);
	} else {
		fmt.Println(task);
	}

	fmt.Println(TaskList);
}