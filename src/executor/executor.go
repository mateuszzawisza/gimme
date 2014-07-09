package executor

import (
    "fmt"
    "os/exec"
    "sync"
)

func Execute(job string) {
    out,_ := exec.Command("sh", "-c", job).Output()
    fmt.Printf("%s",out)
}

func asyncExecute(job string, wg_handle *sync.WaitGroup) {
    defer wg_handle.Done()
    Execute(job)
}

func AsyncExecuteJobs(jobs_list []string) {
    var wg sync.WaitGroup
    for i:= range jobs_list {
        wg.Add(1)
        go asyncExecute(jobs_list[i], &wg)
    }
    wg.Wait()
}
