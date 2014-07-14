#gimme

gimme is simple tool written in golang. It collects various data from server for debugging purposes. Data is afterwards archived and sent to s3 bucket for later analysis.

##Build

Make sure you have GOPATH environment variable set properly. To build execute

```shell
go get
go build gimme.go
```

##Run

To simply gather some data and create a tgz archive run gimme without any arguments
```shell
gimme
```

gimme can upload the tgz archive to specified S3 bucket. To do so add arguments like in the example below.
```shell
gimme -upload-to-s3 -aws-access-key-id=AAA -aws-secret-access-key=BBB -aws-bucket=my-bucket-name
```

Sample output

```
2014/07/13 20:48:18 Saving stats into  /root/gimme_output/20140713204818
2014/07/13 20:48:18 Starting tcp_dump
2014/07/13 20:48:18 Starting connections_list
2014/07/13 20:48:18 Starting connections_stats
2014/07/13 20:48:18 Starting netstat
2014/07/13 20:48:18 Starting who
2014/07/13 20:48:18 Starting processlist
2014/07/13 20:48:18 Starting iostat
2014/07/13 20:48:18 Starting uptime
2014/07/13 20:48:18 Finished who
2014/07/13 20:49:15 Finished iostat
2014/07/13 20:49:18 Finished tcp_dump
2014/07/13 20:49:18 Finished uptime
2014/07/13 20:49:18 Finished connections_stats
2014/07/13 20:49:18 Finished netstat
2014/07/13 20:49:18 Finished connections_list
2014/07/13 20:49:18 Finished processlist
2014/07/13 20:49:18 Creating tgz archive: /root/gimme_output/20140713204818.tar.gz
2014/07/13 20:49:19 Uploading file to https://s3.amazonaws.com/my-test-bucket/fancy-sever-name/20140713204818.tar.gz
2014/07/13 20:49:20 Done
```
