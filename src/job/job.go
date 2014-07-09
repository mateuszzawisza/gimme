package job

type Job struct{
    Command []string
    Repeat int
    Sleep int
    Log_output bool
}
var Jobs = map[string]Job{
    "tcp_dump": Job{
        []string{ "tcpdump -i eth0 -G 60 -W 1 -w tcp.pcap" },
        0,
        1,
        false,
    },
    "connections_list": Job{
        []string{ "ss -an" },
        5,
        1,
        true,
    },
    "connections_stats": Job{
        []string{ "ss -s" },
        5,
        1,
        true,
    },
}
