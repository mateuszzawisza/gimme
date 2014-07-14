package main

import (
	"flag"
	"os"

	"github.com/mateuszzawisza/gimme/archivist"
	"github.com/mateuszzawisza/gimme/executor"
	"github.com/mateuszzawisza/gimme/jobs"

	"time"

	"log"
)

var awsBucket = flag.String(
	"aws-bucket",
	"my-bucket-name",
	"Set the bucket name for the diagnostic information to be uploaded to",
)
var awsAccessKeyId = flag.String(
	"aws-access-key-id",
	"ACCESS_KEY_ID",
	"Set the access key id for AWS API communication",
)
var awsSecretAccessKey = flag.String(
	"aws-secret-access-key",
	"SECRET_ACCESS_KEY_ID",
	"Set the secret access key for AWS API communication",
)
var uploadToS3 = flag.Bool("upload-to-s3", false, "Upload files to s3 bucket?")

func prepare() string {
	path := os.Getenv("HOME") + "/gimme_output/" + time.Now().Format("20060102150405")
	err := os.MkdirAll(path, os.ModeDir|0755)
	if err != nil {
		log.Panicf("Failed to create output directory: %s\n", path, err)
	}
	err = os.Chdir(path)
	if err != nil {
		log.Panicf("Failed to change dir to output directory: %s\n", path, err)
	}
	log.Println("Saving stats into ", path)
	return path
}

func main() {
	flag.Parse()
	path := prepare()
	executor.AsyncExecuteJobs(jobs.Jobs)
	compressed_file_path := archivist.Compress(path)
	if *uploadToS3 == true {
		archivist.S3Upload(
			*awsAccessKeyId,
			*awsSecretAccessKey,
			*awsBucket,
			compressed_file_path,
		)
	}
}
