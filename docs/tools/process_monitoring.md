---
title: Process Monitoring Graphics
tags: [process,monitoring]
---

# Generating graphics from processes

[Procpath](https://pypi.org/project/Procpath/) is a python CLI tool to record and generate graphics of processes.

Example for run a process and get RSS and CPU data

```bash
# Run the process and continue
my_process_binary & 

# Save the process id from last command
pid=$!

# record data from process for 60 seconds
procpath record -i 1 -r 60 -d process.sqlite '$..children[?(@.stat.pid == $pid)]'

# plots the data into a SVG graphic
procpath plot -d process.sqlite -q cpu -q rss cpu_rss.svg

# Remove data file (optional)
rm process.sqlite

```

[Source](https://unix.stackexchange.com/questions/554/how-to-monitor-cpu-memory-usage-of-a-single-process)