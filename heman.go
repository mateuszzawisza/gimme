package main

import "executor"
import "jobs"


func main() {
    executor.AsyncExecuteJobs(jobs.Jobs)
}
