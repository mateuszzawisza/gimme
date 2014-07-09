package main

import "executor"
import "job"


func main() {
    jobs_list := [][]string{job.Jobs["tcp_dump"], job.Jobs["list"]}
    executor.AsyncExecuteJobs(jobs_list)
}
