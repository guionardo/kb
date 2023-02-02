---
title: Cronjob to clean all dotnet projects
tags:
    - cron
    - csharp
    - dotnet
    - bash
---

My codebase is backuped every 6 hours. Some projects aren't touched by weeks, and the build files, binaries and objects are useless to backup.

So I created a script to run every week, cleaning the projects and reducing the data sent to backup.

## clean_csharp_all.sh

```bash
#!/bin/bash
# Script to clean all C# projects in a folder

no_clean=0
clean=0
bytes_saved=0

do_clean() {
    echo "Cleaning $1"
    folder=$(dirname "$1")
    folder_size_before=$(du -b -c "$folder" | tail -n 1 | awk '{print $1;}')
    cd $folder
    dotnet clean
    folder_size_after=$(du -b -c "$folder" | tail -n 1 | awk '{print $1;}')
    if [ $folder_size_after -eq $folder_size_before ]; then
        ((no_clean++))
    else
        echo "Folder size before: $folder_size_before -> after: $folder_size_after"
        ((clean++))
        bytes_saved=$((bytes_saved + $folder_size_before - $folder_size_after))
    fi
}

# This command searches for all solution files in my $HOME/dev folder   
for s in $(find ~/dev -type f -name "*.sln"); do do_clean "$s"; done

echo "Cleaned $clean projects, $no_clean projects were already clean"
echo "Bytes saved: $bytes_saved"
```

## crontab

```bash
# Runs the clean dotnet projects every monday @ 23:50
50 23 * * 1 /home/guionardo/dev/scripts/clean_csharp_all.sh
```
