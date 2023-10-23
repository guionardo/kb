---
title: Network tools
tags:
    - network
    - speed
    - connectivity
---
# Network tools

My LAN was working at 100Mbps. To improve data transfer here, I bought a new 1GBit switch.

In the first moment, nothing changed. All connections was still transfering at 100Mbps. I have machines running Linux and Windows. Checking the NIC of all computers, I see that all of them are 1Gbps compatible.

So, let´s check the NIC´s.

## Windows

Run this command at the console (powershell or cmd).

```powershell
wmic nic where netEnabled=true get name,speed
```

My result:

``` bash
Name
Realtek PCIe GBe Family Controller  100000000
```

Ok, the windows running machine has a Gigabit NIC, but is using only 100Mbps. Should be the lan cable. I switched to a new CAT-5e, and it worked like a charm.

New results:

``` bash
Name
Realtek PCIe GBe Family Controller  1000000000
```

In another machine, a server running Debian, I have to check the same thing.

Listing the /sys/class/net folder we can get all the NIC´s (physical and logical). We need to identify the physical.

```bash
guionardo@furlan-server:~$ ls /sys/class/net
br-118e120c0ed5  br-8006552144f0  br-9baa19a039ab  br-c031f2fd9034  enp2s0  veth0e1f9f5  veth3edbe57  veth763a5d9  vethff997b4
br-6f4f54be3aa6  br-8b05e808f775  br-be4b1c4c6490  docker0          lo      veth0fe6d08  veth5ba2e03  veth7c4fb5c
```

In my case, is the **enp2s0**.

Now, we need to check the settings of the NIC.

```bash
cat /sys/class/net/enp2s0/speed
```

The result was:

```
100
```

In linux, we can also use the *ethtool* command. If it isn´t installed, run the command `sudo apt install ethtool` or use the package manager of your linux distro.

```bash
guionardo@furlan-server:~$ ethtool enp2s0
 for enp2s0:
        Supported ports: [ TP    MII ]
        Supported link modes:   10baseT/Half 10baseT/Full
                                100baseT/Half 100baseT/Full
                                1000baseT/Half 1000baseT/Full
        Supported pause frame use: Symmetric Receive-only
        Supports auto-negotiation: Yes
        Supported FEC modes: Not reported
        Advertised link modes:  10baseT/Half 10baseT/Full
                                100baseT/Half 100baseT/Full
                                1000baseT/Half 1000baseT/Full
        Advertised pause frame use: Symmetric Receive-only
        Advertised auto-negotiation: Yes
        Advertised FEC modes: Not reported
        Link partner advertised link modes:  10baseT/Half 10baseT/Full
                                             100baseT/Half 100baseT/Full
                                             1000baseT/Full
        Link partner advertised pause frame use: Symmetric Receive-only
        Link partner advertised auto-negotiation: Yes
        Link partner advertised FEC modes: Not reported
        Speed: 100Mb/s
        Duplex: Full
        Auto-negotiation: on
        master-slave cfg: preferred slave
        master-slave status: slave
        Port: Twisted Pair
        PHYAD: 1
        Transceiver: external
        MDI-X: Unknown
 error: Operation not permitted
        Current message level: 0x000000ff (255)
                               drv probe link timer ifdown ifup rx_err tx_err
        Link detected: yes
```

So, I had to replace the ETH cable like previous machine.

```bash
cat /sys/class/net/enp2s0/speed
```

The new result was:

```
1000
```

Ok, we have a Gigabit connection enabled. Let´s test it.

I´ll use [iperf](https://iperf.fr/). On debian, if it isn´t installed, just run a `sudo apt install iperf3` on both machines.

On one machine, start a iperf instance, as a server:

``` bash
iperf3 -s
`` 

And on another one, start a iperf client. Check the IP orof the server machine.

``` bash
iperf3 -c furlan-server
```

The results:

```
                                                │guionardo@furlan-server:~$ iperf3 -s
/mnt/c/dev                                                                                 │-----------------------------------------------------------
❯ iperf3 -c furlan-server                                                                  │Server listening on 5201
Connecting to host furlan-server, port 5201                                                │-----------------------------------------------------------
[  5] local 172.20.185.164 port 51290 connected to 192.168.88.35 port 5201                 │Accepted connection from 192.168.88.100, port 58042
[ ID] Interval           Transfer     Bitrate         Retr  Cwnd                           │[  5] local 192.168.88.35 port 5201 connected to 192.168.88.100 port 58044
[  5]   0.00-1.00   sec   107 MBytes   896 Mbits/sec    0   3.11 MBytes                    │[ ID] Interval           Transfer     Bitrate
[  5]   1.00-2.00   sec   104 MBytes   870 Mbits/sec    0   3.11 MBytes                    │[  5]   0.00-1.00   sec  98.2 MBytes   823 Mbits/sec
[  5]   2.00-3.00   sec   101 MBytes   850 Mbits/sec    0   3.11 MBytes                    │[  5]   1.00-2.00   sec   105 MBytes   877 Mbits/sec
[  5]   3.00-4.00   sec   105 MBytes   881 Mbits/sec    0   3.11 MBytes                    │[  5]   2.00-3.00   sec   101 MBytes   851 Mbits/sec
[  5]   4.00-5.00   sec   106 MBytes   891 Mbits/sec    0   3.11 MBytes                    │[  5]   3.00-4.00   sec   104 MBytes   876 Mbits/sec
[  5]   5.00-6.00   sec   108 MBytes   902 Mbits/sec    1   2.24 MBytes                    │[  5]   4.00-5.00   sec   106 MBytes   888 Mbits/sec
[  5]   6.00-7.00   sec   110 MBytes   923 Mbits/sec    0   2.44 MBytes                    │[  5]   5.00-6.00   sec   109 MBytes   912 Mbits/sec
[  5]   7.00-8.00   sec   110 MBytes   923 Mbits/sec    0   2.62 MBytes                    │[  5]   6.00-7.00   sec   109 MBytes   917 Mbits/sec
[  5]   8.00-9.00   sec   110 MBytes   923 Mbits/sec    0   2.75 MBytes                    │[  5]   7.00-8.00   sec   110 MBytes   919 Mbits/sec
[  5]   9.00-10.00  sec   110 MBytes   923 Mbits/sec    0   2.86 MBytes                    │[  5]   8.00-9.00   sec   111 MBytes   928 Mbits/sec
- - - - - - - - - - - - - - - - - - - - - - - - -                                          │[  5]   9.00-10.00  sec   110 MBytes   919 Mbits/sec
[ ID] Interval           Transfer     Bitrate         Retr                                 │[  5]  10.00-10.07  sec  7.32 MBytes   910 Mbits/sec
[  5]   0.00-10.00  sec  1.05 GBytes   898 Mbits/sec    1             sender               │- - - - - - - - - - - - - - - - - - - - - - - - -
[  5]   0.00-10.07  sec  1.04 GBytes   891 Mbits/sec                  receiver             │[ ID] Interval           Transfer     Bitrate
                                                │[  5]   0.00-10.07  sec  1.04 GBytes   891 Mbits/sec                  receiver
iperf Done.                                                                                │-----------------------------------------------------------
                                                │Server listening on 5201
/mnt/c/dev took 10s                                                                                       │-----------------------------------------------------------
❯                                                                                                         │
```

## Links

* [How to Verify the Speed of My NIC?](https://www.baeldung.com/linux/nic-speed)
* [Como saber a velocidade da placa de rede](https://dtnetwork.com.br/como-saber-a-velocidade-da-placa-de-rede/)
* [How to test the network speed/throughput between two Linux servers](https://www.cyberciti.biz/faq/how-to-test-the-network-speedthroughput-between-two-linux-servers/)
