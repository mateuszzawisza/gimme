package executor

import (
    "fmt"
    "os/exec"
    "sync"
    "job"
)

func Execute(job string) {
    _,err := exec.Command("sh", "-c", job).Output()
    if err != nil {
        fmt.Printf("  Failed executing command: \"%s\"\n",job)
    }
}

func asyncExecute(name string, task job.Job, wg_handle *sync.WaitGroup) {
    defer wg_handle.Done()
    fmt.Printf("Starting %s\n", name)
    for i:= range task.Command{
        Execute(task.Command[i])
    }
    fmt.Printf("Finished %s\n", name)
}

func AsyncExecuteJobs(jobs_list map[string]job.Job) {
    var wg sync.WaitGroup
    for name,jobs:= range jobs_list {
        wg.Add(1)
        go asyncExecute(name,jobs, &wg)
    }
    wg.Wait()
}
