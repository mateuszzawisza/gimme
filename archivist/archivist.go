package archivist

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

const s3UploadPartSize int64 = 5 * 1048576

func S3Upload(aws_access_key_id, aws_secret_access_key, aws_bucket, file_path string) {
	file, _ := os.Open(file_path)
	defer file.Close()
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
	multi, err := diag_bucket.InitMulti(s3_path, "application/x-compressed", s3.BucketOwnerFull)
	parts, multi_err := multi.PutAll(file, s3UploadPartSize)
	if multi_err != nil {
		log.Panic("Failed to upload file", err)
	}
	complete_err := multi.Complete(parts)
	if complete_err != nil {
		log.Panic("Failed to upload file", err)
	}
	log.Println("Done")
}

func Compress(path string) string {
	dir, archive_name := dir_and_file_name(path)
	err := os.Chdir(dir)
	if err != nil {
		log.Panicf("Failed to change dir to output directory: %s\n", dir, err)
	}
	log.Printf("Creating tgz archive: %s/%s.tar.gz\n", dir, archive_name)
	tar_command := fmt.Sprintf("tar cvzf %s.tar.gz %s", archive_name, archive_name)
	_, err = exec.Command("sh", "-c", tar_command).Output()
	if err != nil {
		log.Panicf("Failed to execute: %s\n", tar_command, err)
	}
	return fmt.Sprintf("%s/%s.tar.gz", dir, archive_name)
}

func dir_and_file_name(path string) (string, string) {
	split := strings.Split(path, "/")
	dir := strings.Join(split[0:len(split)-1], "/")
	file_name := split[len(split)-1:][0]
	return dir, file_name
}
