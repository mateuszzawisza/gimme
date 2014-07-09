package job

var Jobs = map[string][]string{
    "tcp_dump": []string{
        "sleep 2 && echo 1",
        "echo 'test'",
        "tcpdump -i eth0 -G 10 -W 1 -w /vagrant/test.pcap",
    },
    "list": []string{
        "ls -la",
        "sleep 10",
        "ls -latr",
    },
}
