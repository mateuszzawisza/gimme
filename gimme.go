package main

import (
	"flag"
	"os"

	"time"

	"github.com/mateuszzawisza/gimme/archivist"
	"github.com/mateuszzawisza/gimme/executor"
	"github.com/mateuszzawisza/gimme/jobs"

	"log"
)

type Options struct {
	AwsBucket, AwsAccessKeyId, AwsSecretAccessKey string
}

func gather_flags() Options {
	var aws_bucket = flag.String("aws-bucket", "my-bucket-name", "Set the bucket name for the diagnostic information to be uploaded to")
	var aws_access_key_id = flag.String("aws-access-key-id", "ACCESS_KEY_ID", "Set the access key id for AWS API communication")
	var aws_secret_access_key = flag.String("aws-secret-access-key", "SECRET_ACCESS_KEY_ID", "Set the secret access key for AWS API communication")
	flag.Parse()
	options := Options{
		AwsBucket:          *aws_bucket,
		AwsAccessKeyId:     *aws_access_key_id,
		AwsSecretAccessKey: *aws_secret_access_key,
	}
	return options
}

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
	options := gather_flags()
	path := prepare()
	executor.AsyncExecuteJobs(jobs.Jobs)
	compressed_file_path := archivist.Compress(path)
	archivist.S3Upload(
		options.AwsAccessKeyId,
		options.AwsSecretAccessKey,
		options.AwsBucket,
		compressed_file_path,
	)
}
