Elastos Hive IPFS
==============
[![](https://img.shields.io/badge/made%20by-Elastos%20org-blue.svg?style=flat-square)](http://elastos.org)
[![](https://img.shields.io/badge/project-Hive-blue.svg?style=flat-square)](http://elastos.org/)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Build Status](https://travis-ci.org/elastos/Elastos.NET.Hive.IPFS.svg?branch=master)](https://travis-ci.org/elastos/Elastos.NET.Hive.IPFS)

## What is Hive Project
Elastos Hive is a basic service infrastructure that provide storage capabilities to dApps with decentralized characterstics, which leveraged standard IPFS/Cluster implementation but would be totally  a new storage network seperated from standard IPFS network.

## What is IPFS

IPFS is a global, versioned, peer-to-peer filesystem. It combines good ideas from Git, BitTorrent, Kademlia, SFS, and the Web. It is like a single bittorrent swarm, exchanging git objects. IPFS provides an interface as simple as the HTTP web, but with permanence built in. You can also mount the world at /ipfs.

For more info see: https://github.com/ipfs/ipfs.

Please put all issues regarding:

  - IPFS node in the [ipfs repo issues](https://github.com/elastos/Elastos.NET.Hive.Hive)
  - IPFS cluster in the [cluster repo issues](htttps://github.com/elastos/Elastos.NET.Hive.Cluster)

## Table of Contents

- [Install](#install)
  - [System Requirements](#system-requirements)
  - [Build from Source](#2.-Build from Source)
    - [Install Go](#Install-Go)
    - [Download Source](#Download-Source)
    - [Build IPFS](#Build-IPFS)
- [Usage](#usage)
- [Getting Started](#getting-started)
  - [Initialize config files](#1.Initialize-config-files)
  - [Swarm Key](#2.Swarm-Key)
  - [Run as daemon](#3.Run-as-Daemon)
  - [Some things to try](#4.Some-things-to-try)
- [Contributions](#contributions)
- [Acknowledgments](#Acknowledgments)
- [License](#license)

## Install

The canonical download instructions for IPFS are over at: http://ipfs.io/docs/install/. It is **highly suggested** you follow those instructions if you are not interested in working on IPFS development.

### 1. System Requirements

IPFS can run on most Linux, macOS, and Windows systems. We recommend running it on a machine with at least 2 GB of RAM (it’ll do fine with only one CPU core), but it should run fine with as little as 1 GB of RAM. On systems with less memory, it may not be completely stable.

### 2. Build from Source

#### Install Go

The build process for ipfs requires Go 1.10 or higher. Download propriate binary archive from [golang.org](https://golang.org/dl),  and install it specific directory:

```shell
$ curl go1.11.4.linux-amd64.tar.gz -o golang.tar.gz
$ tar -xzvf golang.tar.gz -C YOUR-PATH
```

Then add path `go/bin` of **YOUR-INSTALL-PATH** to user environment variable **PATH**.

```shell
$ export PATH="YOUR-INSTALL-PATH/go/bin:$PATH"
```

You also would be required to  setup build environment for `go` projects:

```shell
$ export GOPATH="YOUR-GO-PATH"
$ export PATH="$GOPATH/bin:$PATH"
```

In convinience way,  just add these lines to your environment profile `$HOME/.profile`, then validate it with command:

```shell
$ . $HOME/.profile
$ export 
```

Then use **export** command to check if it validated or not.

**Notice** : If you run into trouble, see the [Go install instructions](https://golang.org/doc/install).

#### Download Source

There are two ways to setup your source code. One is directly to download source code under `$GOPATH` environment as below:

```shell
$ cd $GOPATH/src/github.com/elastos
$ git clone https://github.com/elastos/Elastos.NET.Hive.IPFS.git
```

The other way is to download source to specific location, and then create linkage to that source directory under `$GOPATH` environment:

```shell
$ cd YOUR-PATH
$ git clone https://github.com/elastos/Elastos.NET.Hive.IPFS.git 
$ cd $GOPATH/src/github.com/elastos
$ link -s YOUR-PATH/Elastos.NET.Hive.IPFS Elastos.NET.Hive.IPFS
```

#### Build IPFS

To build IPFS with the following commands:

```shell
$ cd $GOPATH/src/github.com/elastos/Elastos.NET.Hive.IPFS
$ make
$ make install
```

After installation, **ipfs** execute binary would be generated under directory `$GOPATH/bin`.

## Usage

Run **ipfs** with option `help` to understand **HOWTO** run ipfs.

```
$ ipfs help
USAGE
  ipfs - Global p2p merkle-dag filesystem.

  ipfs [<flags>] <command> [<arg>] ...

SUBCOMMANDS
  BASIC COMMANDS
    init          Initialize ipfs local configuration
    add <path>    Add a file to ipfs
    cat <ref>     Show ipfs object data
    get <ref>     Download ipfs objects
    ls <ref>      List links from an object
    refs <ref>    List hashes of links from an object

  DATA STRUCTURE COMMANDS
    block         Interact with raw blocks in the datastore
    object        Interact with raw dag nodes
    files         Interact with objects as if they were a unix filesystem

  ADVANCED COMMANDS
    daemon        Start a long-running daemon process
    mount         Mount an ipfs read-only mountpoint
    resolve       Resolve any type of name
    name          Publish or resolve IPNS names
    dns           Resolve DNS links
    pin           Pin objects to local storage
    repo          Manipulate an IPFS repository

  NETWORK COMMANDS
    id            Show info about ipfs peers
    bootstrap     Add or remove bootstrap peers
    swarm         Manage connections to the p2p network
    dht           Query the DHT for values or peers
    ping          Measure the latency of a connection
    diag          Print diagnostics

  TOOL COMMANDS
    config        Manage configuration
    version       Show ipfs version information
    update        Download and apply go-ipfs updates
    commands      List all available commands

  Use 'ipfs <command> --help' to learn more about each command.

  ipfs uses a repository in the local file system. By default, the repo is located
  at ~/.ipfs. To change the repo location, set the $IPFS_PATH environment variable:

    export IPFS_PATH=/path/to/ipfsrepo
```

## Getting Started

### 1. Initialize config files

To start using IPFS, you must first initialize IPFS's config files on your system, which would be done with the command below:

```shell
$ ipfs init
$ ls $HOME/.ipfs
blocks		config		datastore	datastore_spec	keystore	version
```

See `ipfs init --help` for more information on the optional arguments it takes.

As to Hive project, you need to connect to private IPFS network by reconfiguring bootstrap list with the following commands:

```shell
$ ipfs bootstrap rm --all
$ ipfs bootstrap add /ip4/YOUR-BOOTSTRAP-IP/tcp/4001/YOUR-BOOTSTRAP-NODEID
```

### 2.  Swarm Key

In general IPFS,  IPFS swarm is the network of peers connection. There is a swarm key(swarm.key file) that could be shared by the network. If IPFS daemon does not locate the swarm.key file, it will connect to global IPFS network. Otherwise, with swarm key,  the IPFS daemon will try to use it to connect a private secret network.

```
$ get -u github.com/Kubuxu/go-ipfs-swarm-key-gen/ipfs-swarm-key-gen
$ ipfs-swarm-key-gen > swarm.key
$ cp swarm.key $HOME/.ipfs
```

### 3. Run as Daemon

After initialization and swarm key configured, now it's time to run IPFS daemon with backgroud suggested:

```shell
$ ipfs daemon &
go-ipfs version: 0.4.18-743181038
Repo version: 7
System version: amd64/darwin
Golang version: go1.11.4
Successfully raised file descriptor limit to 2048.

dev-mbp:html stiartsly$ Swarm listening on /ip4/10.8.189.103/tcp/4001
Swarm listening on /ip4/127.0.0.1/tcp/4001
Swarm listening on /ip4/172.16.30.1/tcp/4001
Swarm listening on /ip4/192.168.1.112/tcp/4001
Swarm listening on /ip4/192.168.9.1/tcp/4001
Swarm listening on /ip6/::1/tcp/4001
Swarm listening on /p2p-circuit
Swarm announcing /ip4/10.8.189.103/tcp/4001
Swarm announcing /ip4/127.0.0.1/tcp/4001
Swarm announcing /ip4/172.16.30.1/tcp/4001
Swarm announcing /ip4/192.168.1.112/tcp/4001
Swarm announcing /ip4/192.168.9.1/tcp/4001
Swarm announcing /ip6/::1/tcp/4001
API server listening on /ip4/127.0.0.1/tcp/5001
Gateway (readonly) server listening on /ip4/127.0.0.1/tcp/8080
Daemon is ready
```

### 4. Some things to try

Basic proof of 'ipfs working' locally:

	$ echo "hello world" > hello
	$ ipfs add hello
	# This should output a hash string that looks something like:
	# QmT78zSuBmuS4z925WZfrqQ1qHaJ56DQaTfyMUF7F8ff5o
	$ ipfs cat <that hash>

Then you can try `ipfs` command with various options to check IPFS daemon is working.

### Docker usage

`TODO`

### Testing

```
make test
```

## Contributions

We welcome any contribution to Elastos Hive projects.

## Acknowledgments

A sincere thank you to all teams and projects that we rely on directly or indirectly, especially`go-ipfs` Project.

## License

This project is licensed under the terms of the [MIT license](https://github.com/elastos/Elastos.NET.Carrier.Native.SDK/blob/master/LICENSE).