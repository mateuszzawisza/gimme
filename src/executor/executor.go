package executor

import (
    "fmt"
    "os/exec"
    "sync"
)
type Job struct{
    Commands []string
}

type JobsList struct{
    Jobs []Job
}

func Execute(job string) {
    out,err := exec.Command("sh", "-c", job).Output()
    if err != nil {
        fmt.Printf("[%s] failed\n",job)
    }else {
        fmt.Printf("[%s] %s\n",job, out)
    }
}

func asyncExecute(commands []string, wg_handle *sync.WaitGroup) {
    defer wg_handle.Done()
    for i:= range commands{
        Execute(commands[i])
    }
}

func AsyncExecuteJobs(jobs_list JobsList) {
    var wg sync.WaitGroup
    jobs := jobs_list.Jobs
    for i:= range jobs {
        wg.Add(1)
        go asyncExecute(jobs[i].Commands, &wg)
    }
    wg.Wait()
}
