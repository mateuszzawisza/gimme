package executor

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	Command    []string
	Repeat     int
	Sleep      int
	Log_output bool
}

func execute(job string) []byte {
	output, err := exec.Command("sh", "-c", job).Output()
	if err != nil {
		log.Printf("Failed executing command: \"%s\"\n", job)
	}
	return output
}

func executeTasks(name string, task Job) {
	for i := range task.Command {
		output := execute(task.Command[i])
		if task.Log_output == true {
			saveOutput(name, output)
		}
	}
}

func asyncExecute(name string, task Job, wg_handle *sync.WaitGroup) {
	defer wg_handle.Done()
	log.Printf("Starting %s\n", name)
	if task.Repeat > 0 {
		for i := 0; i < task.Repeat; i++ {
			task_name := name + "." + strconv.Itoa(i)
			executeTasks(task_name, task)
			time.Sleep(time.Duration(task.Sleep) * time.Second)
		}
	} else {
		executeTasks(name, task)
	}
	log.Printf("Finished %s\n", name)
}

func AsyncExecuteJobs(jobs_list map[string]Job) {
	var wg sync.WaitGroup
	for name, jobs := range jobs_list {
		wg.Add(1)
		go asyncExecute(name, jobs, &wg)
	}
	wg.Wait()
}

func saveOutput(name string, content []byte) {
	file, err := os.OpenFile(name+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write([]byte("===== Content inserted on " + time.Now().String() + "=====\n"))
	file.Write(content)
}
