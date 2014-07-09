package main

import "executor"

func main() {
    jobs := make([]string, 3)
    jobs[0] = "sleep 2 && echo 1"
    jobs[1] = "sleep 10 && echo 2"
    jobs[2] = "echo 'test'"
    executor.AsyncExecuteJobs(jobs)
}
