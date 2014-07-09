package main

import "executor"
import "job"


func main() {
    //jobs_list := [][]string{
    //    job.Jobs["tcp_dump"],
    //    job.Jobs["connections_list"],
    //    job.Jobs["connections_stats"],
    //}
    executor.AsyncExecuteJobs(job.Jobs)
}
