---
title: Commands and Aliases
tags:
    - bash
    - shell
---

## Commands and aliases for shell

### Listing local IPs

```bash
alias ips='ip -c -br a'
```

```bash
✦ ❯ ips
lo               UNKNOWN        127.0.0.1/8 ::1/128 
enp1s0           UP             192.168.88.72/24 fe80::14:f207:b13:562f/64 
wlp2s0           UP             192.168.88.21/24 fe80::356b:b5b5:1fae:4a22/64 
br-aa643bd30424  DOWN           172.19.0.1/16 
br-5b8d6bec92cf  DOWN           172.22.0.1/16 
br-91c18c436867  DOWN           172.20.0.1/16 
br-98b6cfd1fb53  DOWN           172.23.0.1/16 
br-d26beecfc65c  DOWN           172.24.0.1/16 
br-e0165e66c384  DOWN           172.25.0.1/16 
br-fbbdabe1e428  DOWN           172.18.0.1/16 
docker0          DOWN           172.17.0.1/16 
br-58d103d2fc5a  DOWN           172.21.0.1/16 
br-981cb645b58a  UP             172.26.0.1/16 fe80::42:89ff:fe6b:bc92/64 
vethcd4aaa2@if19 UP             fe80::4845:40ff:fe18:865c/64 
vpn              UNKNOWN        10.186.112.113/32 fe80::fe07:f1d8:daf6:bd47/64 
```

### Listing open ports

```bash
alias ports='sudo netstat -tulanp'
```

```bash
✦ ❯ ports
Conexões Internet Ativas (servidores e estabelecidas)
Proto Recv-Q Send-Q Endereço Local          Endereço Remoto         Estado       PID/Program name    
tcp        0      0 0.0.0.0:27017           0.0.0.0:*               OUÇA       922978/docker-proxy 
tcp        0      0 127.0.0.53:53           0.0.0.0:*               OUÇA       1229870/systemd-res 
tcp        0      0 127.0.0.1:8000          0.0.0.0:*               OUÇA       1229353/python3.11  
tcp        0      0 127.0.0.1:39451         0.0.0.0:*               OUÇA       867801/Rider.Backen 
tcp        0      0 127.0.0.1:39485         0.0.0.0:*               OUÇA       867801/Rider.Backen 
tcp        0      0 0.0.0.0:22              0.0.0.0:*               OUÇA       1007/sshd: /usr/sbi 
tcp        0      0 0.0.0.0:111             0.0.0.0:*               OUÇA       1/init              
tcp        0      0 127.0.0.1:46727         0.0.0.0:*               OUÇA       894/confighandler   
tcp        0      0 127.0.0.1:46595         0.0.0.0:*               OUÇA       867801/Rider.Backen 
tcp        0      0 127.0.0.1:37085         0.0.0.0:*               OUÇA       867801/Rider.Backen 
tcp        0      0 127.0.0.1:37545         0.0.0.0:*               OUÇA       925/containerd      
tcp        0      0 127.0.0.1:52783         0.0.0.0:*               OUÇA       6087/kbfsfuse       
tcp        0      0 127.0.0.1:44959         0.0.0.0:*               OUÇA       867801/Rider.Backen 
tcp        0      0 127.0.0.1:17915         0.0.0.0:*               OUÇA       5995/keybase        
tcp        0      0 127.0.0.1:25            0.0.0.0:*               OUÇA       3069/master         
tcp        0      0 127.0.0.1:33117         0.0.0.0:*               OUÇA       867801/Rider.Backen 
tcp        0      0 127.0.0.1:41653         0.0.0.0:*               OUÇA       1225749/Code --stan 
tcp        0      0 127.0.0.1:631           0.0.0.0:*               OUÇA       473882/cupsd        
tcp        0      0 192.168.88.72:35030     104.18.29.31:443        ESTABELECIDA 38367/chrome --type 
tcp        0      0 192.168.88.72:35076     52.113.207.5:443        ESTABELECIDA 6478/teams --type=u 
tcp        0      0 192.168.88.72:34128     157.240.226.60:443      ESTABELECIDA 38367/chrome --type 
tcp        0      0 192.168.88.72:52924     54.205.255.92:443       ESTABELECIDA 6087/kbfsfuse       
tcp        0      0 127.0.0.1:27017         127.0.0.1:55954         ESTABELECIDA 922978/docker-proxy 
tcp        0      0 10.186.112.113:60616    172.22.181.9:443        ESTABELECIDA 13601/firefox       
tcp        0      0 192.168.88.72:51870     45.60.38.211:443        ESTABELECIDA 38367/chrome --type 
tcp        0      0 127.0.0.1:33117         127.0.0.1:42350         ESTABELECIDA 867801/Rider.Backen 
```

### Searching commands in history

```bash
alias gh='history|grep'
```

call ```gh``` with the grep expression you are searching for.

```bash
 ❯ gh sudo
  106  sudo gpg --no-default-keyring --keyring /usr/share/keyrings/k6-archive-keyring.gpg --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
  107  echo "deb [signed-by=/usr/share/keyrings/k6-archive-keyring.gpg] https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list
  108  sudo apt update
  109  sudo apt install k6
  210  sudo cat /proc/931/mem
  211  sudo cat /proc/931/cmdline
  506  sudo dpkg -i libssl1.1_1.1.1n-0+deb11u3_amd64.deb 
  736  sudo ./hd_health.sh
```

### Disk usage

```bash
sudo apt install duf
```

```bash
✦ ❯ duf
╭─────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│ 3 local devices                                                                                         │
├───────────────────────────┬────────┬───────┬────────┬───────────────────────────────┬──────┬────────────┤
│ MOUNTED ON                │   SIZE │  USED │  AVAIL │              USE%             │ TYPE │ FILESYSTEM │
├───────────────────────────┼────────┼───────┼────────┼───────────────────────────────┼──────┼────────────┤
│ /                         │ 204.0G │ 81.1G │ 112.4G │ [#######.............]  39.8% │ ext4 │ /dev/sda1  │
│ /boot/efi                 │  96.0M │ 44.9M │  51.1M │ [#########...........]  46.8% │ vfat │ /dev/sdb2  │
│ /var/snap/firefox/common/ │ 204.0G │ 81.1G │ 112.4G │ [#######.............]  39.8% │ ext4 │ /dev/sda1  │
│ host-hunspell             │        │       │        │                               │      │            │
╰───────────────────────────┴────────┴───────┴────────┴───────────────────────────────┴──────┴────────────╯
╭────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│ 1 fuse device                                                                                          │
├───────────────────────────┬────────┬──────┬────────┬───────────────────────────────┬──────┬────────────┤
│ MOUNTED ON                │   SIZE │ USED │  AVAIL │              USE%             │ TYPE │ FILESYSTEM │
├───────────────────────────┼────────┼──────┼────────┼───────────────────────────────┼──────┼────────────┤
│ /run/user/1000/keybase/kb │ 250.0G │ 1.6G │ 248.4G │ [....................]   0.6% │ fuse │ /dev/fuse  │
│ fs                        │        │      │        │                               │      │            │
╰───────────────────────────┴────────┴──────┴────────┴───────────────────────────────┴──────┴────────────╯
╭────────────────────────────────────────────────────────────────────────────────────────────────╮
│ 6 special devices                                                                              │
├────────────────┬──────┬────────┬───────┬───────────────────────────────┬──────────┬────────────┤
│ MOUNTED ON     │ SIZE │   USED │ AVAIL │              USE%             │ TYPE     │ FILESYSTEM │
├────────────────┼──────┼────────┼───────┼───────────────────────────────┼──────────┼────────────┤
│ /dev           │ 7.7G │     0B │  7.7G │                               │ devtmpfs │ udev       │
│ /dev/shm       │ 7.8G │ 637.2M │  7.1G │ [#...................]   8.0% │ tmpfs    │ tmpfs      │
│ /run           │ 1.6G │   2.6M │  1.5G │ [....................]   0.2% │ tmpfs    │ tmpfs      │
│ /run/lock      │ 5.0M │   4.0K │  5.0M │ [....................]   0.1% │ tmpfs    │ tmpfs      │
│ /run/snapd/ns  │ 1.6G │   2.6M │  1.5G │ [....................]   0.2% │ tmpfs    │ tmpfs      │
│ /run/user/1000 │ 1.6G │ 236.0K │  1.6G │ [....................]   0.0% │ tmpfs    │ tmpfs      │
╰────────────────┴──────┴────────┴───────┴───────────────────────────────┴──────────┴────────────╯
```

### DNS informations

Old command: dnslookup

```bash
dig google.com
```

```bash
✦ ❯ dig google.com

; <<>> DiG 9.18.1-1ubuntu1.3-Ubuntu <<>> google.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 61639
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;google.com.   IN A

;; ANSWER SECTION:
google.com.  211 IN A 142.250.79.206

;; Query time: 7 msec
;; SERVER: 127.0.0.53#53(127.0.0.53) (UDP)
;; WHEN: Wed Feb 01 18:53:52 -03 2023
;; MSG SIZE  rcvd: 55
```

### Finding files

Old command: find

```bash
sudo apt install fd-find
```

```bash
✦ ❯ fdfind pdf
FORBE CAPITAL (1).pdf
QuebraCabeca.pdf
ManualTecnico.pdf
RabbitMQ-2023-JAN-26.pdf
```

### Process monitoring

Old command: top, htop, etc

```bash
sudo apt install bashtop
```

- [Seu Terminal Linux não será mais o mesmo depois disso! - Shell ALIAS no Linux](https://www.youtube.com/watch?v=k8SycT32-yA)
- [Os comandos que VOCÊ USA estão OBSOLETOS! 10 alternativas mais eficientes para o terminal Linux](https://www.youtube.com/watch?v=uE4wiVA4YBw)
