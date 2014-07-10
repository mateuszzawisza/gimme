package jobs

import "executor"

var Jobs = map[string]executor.Job{
    "tcp_dump": executor.Job{
        []string{ "tcpdump -i eth0 -G 60 -W 1 -w tcp.pcap" },
        0,
        0,
        false,
    },
    "connections_list": executor.Job{
        []string{ "ss -aen" },
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
    "netstat": executor.Job{
        []string{ "netstat -s || echo 0" },
        12,
        5,
        true,
    },
    "who": executor.Job{
        []string{ "who","w" },
        0,
        0,
        true,
    },
    "processlist": executor.Job{
        []string{ "ps aux" },
        12,
        5,
        true,
    },
    "iostat": executor.Job{
        []string{ "iostat -x 3 20" },
        0,
        0,
        true,
    },
    "uptime": executor.Job{
        []string{ "uptime" },
        6,
        10,
        true,
    },
}
