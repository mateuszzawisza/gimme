package main

import(
    "fmt"
    "executor"
    "jobs"
    "os"
    "time"
)


func main() {
    prepare()
    executor.AsyncExecuteJobs(jobs.Jobs)
}

func prepare(){
    dir:="/root/heman/" + time.Now().Format("20060102150405")
    os.MkdirAll(dir,os.ModeDir)
    os.Chdir(dir)
    fmt.Println("Saving stats into ",dir)
}
