package executor

import (
    "fmt"
    "os/exec"
    "sync"
)

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

func AsyncExecuteJobs(jobs_list [][]string) {
    var wg sync.WaitGroup
    for i:= range jobs_list {
        wg.Add(1)
        go asyncExecute(jobs_list[i], &wg)
    }
    wg.Wait()
}
