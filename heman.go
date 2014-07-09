package main

import "executor"
import "job"


func main() {
    executor.AsyncExecuteJobs(job.Jobs)
}
