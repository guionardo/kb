---
title: Fixing DotNet on Ubuntu
tags:
    - csharp
    - setup
    - dotnet
---
## A fatal error occurred. The folder [/usr/share/dotnet/host/fxr] does not exist

When .NET (Core) was first released for Linux, it was not yet available in the official Ubuntu repo. So instead, many of us added the Microsoft APT repo in order to install it.

Now, the packages are part of the Ubuntu repo, and they are conflicting with the Microsoft packages. This error is a result of mixed packages.

So you need to pick which one you're going to use, and ensure they don't mix. Personally, I decided to stick with the Microsoft packages because I figured they'll be better kept up-to-date.

First, remove all existing packages to get to a clean state:

``` bash
sudo apt remove dotnet*
sudo apt remove aspnetcore*
sudo apt remove netstandard*
```

Then, create a file in /etc/apt/preferences.d (I named mine 99microsoft-dotnet.pref, following the convention that files in such *.d directories are typically prefixed with a 2-digit number so that they sort and load in a predictable order) with the following contents:

```text
Package: *
Pin: origin "packages.microsoft.com"
Pin-Priority: 1001
```

Then, the regular update & install:

```bash
sudo apt update
sudo apt install dotnet-sdk-6.0
```

If you would rather use the official Ubuntu packages, remove all the existing packages as above, but instead of creating the /etc/apt/preferences.d entry, just delete the Microsoft repo:

```bash
sudo rm /etc/apt/sources.list.d/microsoft-prod.list
sudo apt update
sudo apt install dotnet-sdk-6.0
```

However, note that the Microsoft repo contains other packages such as PowerShell, SQL Server Command-Line Tools, etc., so removing it may not be desirable.

I'm sure it's possible to make the APT config more specific to just these packages, but this is working for me for now. Hopefully Microsoft and Ubuntu work together to fix this soon.

More info on the issue and various solutions is available here:

- [Microsoft Linux Package Mixup](https://learn.microsoft.com/en-us/dotnet/core/install/linux-package-mixup)
- [GitHub Dotnet core issue 7699](https://github.com/dotnet/core/issues/7699)
- [Original post](https://stackoverflow.com/questions/73753672/a-fatal-error-occurred-the-folder-usr-share-dotnet-host-fxr-does-not-exist)
