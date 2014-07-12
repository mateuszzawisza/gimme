package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/mateuszzawisza/gimme/executor"
	"github.com/mateuszzawisza/gimme/jobs"

	"log"

	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

type Options struct {
	AwsBucket, AwsAccessKeyId, AwsSecretAccessKey string
}

func main() {
	options := gather_flags()
	path := prepare()
	dir, file_name := dir_and_file_name(path)
	executor.AsyncExecuteJobs(jobs.Jobs)
	compressed_file_path := compress(dir, file_name)
	s3Uploader(
		options.AwsAccessKeyId,
		options.AwsSecretAccessKey,
		options.AwsBucket,
		compressed_file_path,
	)
}

func s3Uploader(aws_access_key_id, aws_secret_access_key, aws_bucket, file_path string) {
	file, _ := os.Open(file_path)
	file_stat, _ := file.Stat()
	file_size := file_stat.Size()
	content := make([]byte, file_size)
	file.Read(content)

	auth := aws.Auth{
		AccessKey: aws_access_key_id,
		SecretKey: aws_secret_access_key,
	}
	useast := aws.USEast
	bucket_name := aws_bucket

	connection := s3.New(auth, useast)
	diag_bucket := connection.Bucket(bucket_name)
	hostname, _ := os.Hostname()
	_, file_name := dir_and_file_name(file_path)
	s3_path := fmt.Sprintf("%s/%s", hostname, file_name)
	log.Printf("Uploading file to https://s3.amazonaws.com/%s/%s\n", bucket_name, s3_path)
	err := diag_bucket.Put(s3_path, content, "application/x-compressed", s3.BucketOwnerFull)
	if err != nil {
		log.Panic("Failed to upload file", err)
	}
	log.Println("Done")
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

func compress(dir, output_name string) string {
	err := os.Chdir(dir)
	if err != nil {
		log.Panicf("Failed to change dir to output directory: %s\n", dir, err)
	}

	log.Printf("Creating tgz archive:  %s/%s.tar.gz\n", dir, output_name)
	tar_command := fmt.Sprintf("tar cvzf %s.tar.gz %s", output_name, output_name)
	_, err = exec.Command("sh", "-c", tar_command).Output()
	if err != nil {
		log.Panicf("Failed to execute: %s\n", tar_command, err)
	}
	return fmt.Sprintf("%s/%s.tar.gz", dir, output_name)
}

func dir_and_file_name(path string) (string, string) {
	split := strings.Split(path, "/")
	dir := strings.Join(split[0:len(split)-1], "/")
	file_name := split[len(split)-1:][0]
	return dir, file_name
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
