package job

import "executor"

var Jobs = map[string]executor.Job{
    "tcp_dump": executor.Job{
        []string{ "tcpdump -i eth0 -G 60 -W 1 -w tcp.pcap" },
        0,
        0,
        false,
    },
    "connections_list": executor.Job{
        []string{ "ss -an" },
        12,
        5,
        true,
    },
    "connections_stats": executor.Job{
        []string{ "ss -s" },
        12,
        5,
        true,
    },
}
