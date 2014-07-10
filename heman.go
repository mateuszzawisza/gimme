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
    dir:=os.Getenv("HOME") + "/heman_output/" + time.Now().Format("20060102150405")
    err:=os.MkdirAll(dir,os.ModeDir)
    if err !=nil {panic(err)}
    err=os.Chdir(dir)
    if err !=nil {panic(err)}
    fmt.Println("Saving stats into ",dir)
}
