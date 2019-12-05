# Docker-Compose method for Testnet Sirius
---
This method leverages the public Sirius docker image, embedded in the release section's tar.gz file. [link](https://github.com/proximax-storage/vagrant-testnet-sirius-peernode/releases/tag/release-v0.4.3)



### OS Network requirements
Ensure that your local network can allow inbound/outbound comms on these ports:
```

    3000tcp
    7900tcp
    7901tcp
    7902tcp
```


### A note on System Requirements
Theoretically, this dockerized Sirius Peer package can perform on any Nix OS flavor running Docker version 19.03.3 and docker-compose version version 1.24.0.

But if you really need a minimum benchmark, we have seen the  Sirius Blockchain Peer to work with a minimum Hardware perfomance of   1 CPU core and 2GB RAM.

This doc was prepared by testing the package on:
```
- Debian 10 ++
- Ubuntu 19.04 ++
- CentOS 7 ++

```
---

### Prework! Obtain the Dependencies
Setup your node dependencies to properly run the dockerized Sirius Peer.

Example for an Ubuntu box, run below commands either manually or in script:
```bash

#! /bin/bash
sudo apt update -y
sudo apt upgrade -y
sudo curl -L "https://github.com/docker/compose/releases/download/1.24.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo apt remove docker docker-engine docker.io containerd runc -y
sudo apt update -y ; sudo apt upgrade -y
sudo apt install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - && sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository -y "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt update -y
sudo apt install docker-ce docker-ce-cli containerd.io -y

```

Example for a CentOS box, run below commands either manually or in script:
```bash

#! /bin/bash
sudo yum update -y
sudo curl -L "https://github.com/docker/compose/releases/download/1.24.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo yum remove docker docker-common docker-selinux docker-engine-selinux docker-engine docker-ce
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install -y docker-ce docker-ce-cli containerd.io
sudo systemctl enable docker.service
sudo systemctl start docker.service
sudo systemctl status docker.service

```

### Unpack the docker-compose Package!

```bash
sudo mkdir -m 777 /opt/catapult-config
cd /opt/catapult-config
wget https://proximax-vagrant-storage.s3-ap-southeast-1.amazonaws.com/public-testnet-peer-dockerpackage-v043.tar.gz
tar -xvf public-testnet-peer-dockerpackage-v043.tar.gz
rm -rf public-testnet-peer-dockerpackage-v043.tar.gz
cp -R public-testnet-peer-package-v0.4.3/* /opt/catapult-config
rm -rf public-testnet-peer-package-v0.4.3/

```

---


### Unpack the SEED DATA:
```bash
cd /opt/catapult-config
tar -xvf seed-public-test-v0.4.3.tar.gz

```


### Load the Public ProximaX Peer Docker image into local:
```bash

cd /opt/catapult-config
sudo docker load -i proximax-sirius-peer-toolless.tar

```

---





### Tag your fresh  Docker image with the proper version release:
```bash

cd /opt/catapult-config
sudo docker tag a23e56f572d1  proximax-sirius-peer-toolless:release-v0.4.3


#check the tagged existence of your Public ProximaX Peer image:
sudo docker images
REPOSITORY                      TAG                 IMAGE ID            
proximax-sirius-peer-toolless   release-v0.4.3      a23e56f572d1        

```
---








# Before running Sirius Peer...
You need to create a Testnet Keypair to serve as your HarvestKey.

There are two ways to create a TESTNET harvestKey:
---

1. Using BCTESTNET WALLET - [testnetwallet-link](https://bctestnetwallet.xpxsirius.io/)

2. Create an "account" using the ProximaX Sirius SDK - [sdk-link](https://bcdocs.xpxsirius.io/docs/guides/account/creating-and-opening-an-account/)

---

Once created, simply put your privateKey harvestKey into configfile:
```
/opt/catapult-config/resources/config-harvesting.properties
```

---

You can change the name of the node by changing host and friendlyName parameters:
host 	        string 	Node host (leave empty to auto-detect IP). 	 
friendlyName 	string 	Node friendly name (leave empty to use address).
```
/opt/catapult-config/resources/config-node.properties
```

---

# Run the dockerized Sirius Peer...
Once keys are set, execute your peer to sync to Testnet Chain:
```bash

cd /opt/catapult-config/
sudo /usr/local/bin/docker-compose up

```


---







### Marvel at the chain's behaviour
Once executed, you can easily witness a new chain will be created every 15 seconds by default and the SYNC log entry looks like this:
```bash


2019-11-13 11:24:53.662099 0x00007f8828188700: <info> (chain::ChainSynchronizer.cpp@207) peer returned 400 blocks (heights 1202 - 1601) 
2019-11-13 11:24:53.662605 0x00007f8828188700: <info> (chain::RemoteApiForwarder.h@66) completed 'synchronizer task' (tier1B-mainnet-chain-peer8 @ proioxis.brimstone.xpxsirius.io:7900) with result Success 
2019-11-13 11:24:53.934271 0x00007f881617a700: <info> (disruptor::ConsumerDispatcher.cpp@43) completing processing of element 3 (400 blocks (heights 1202 - 1601) [CE595B87] from Remote_Pull), last consumer is 0 elements behind 
2019-11-13 11:24:57.152197 0x00007f8828188700: <info> (chain::ChainSynchronizer.cpp@207) peer returned 400 blocks (heights 1602 - 2001) 
2019-11-13 11:24:57.152635 0x00007f8828188700: <info> (chain::RemoteApiForwarder.h@66) completed 'synchronizer task' (tier1a-mainnet-chain-peer3 @ omicronlyrae.brimstone.xpxsirius.io:7900) with result Success 
2019-11-13 11:24:57.392612 0x00007f881617a700: <info> (disruptor::ConsumerDispatcher.cpp@43) completing processing of element 4 (400 blocks (heights 1602 - 2001) [DC6021C9] from Remote_Pull), last consumer is 0 elements behind 
2019-11-13 11:25:00.643511 0x00007f8828188700: <info> (chain::ChainSynchronizer.cpp@207) peer returned 400 blocks (heights 2002 - 2401) 
2019-11-13 11:25:00.644119 0x00007f8828188700: <info> (chain::RemoteApiForwarder.h@66) completed 'synchronizer task' (tier1a-mainnet-chain-peer3 @ omicronlyrae.brimstone.xpxsirius.io:7900) with result Success 
2019-11-13 11:25:00.883634 0x00007f881617a700: <info> (disruptor::ConsumerDispatcher.cpp@43) completing processing of element 5 (400 blocks (heights 2002 - 2401) [4462893B] from Remote_Pull), last consumer is 0 elements behind 
2019-11-13 11:25:04.139147 0x00007f8828188700: <info> (chain::ChainSynchronizer.cpp@207) peer returned 400 blocks (heights 2402 - 2801) 
2019-11-13 11:25:04.139515 0x00007f8828188700: <info> (chain::RemoteApiForwarder.h@66) completed 'synchronizer task' (tier1B-mainnet-chain-peer10 @ acallaris.brimstone.xpxsirius.io:7900) with result Success 
2019-11-13 11:25:04.405145 0x00007f881617a700: <info> (disruptor::ConsumerDispatcher.cpp@43) completing processing of element 6 (400 blocks (heights 2402 - 2801) [DEE28CC8] from Remote_Pull), last consumer is 0 elements behind 
2019-11-13 11:25:07.778694 0x00007f8828188700: <info> (chain::ChainSynchronizer.cpp@207) peer returned 400 blocks (heights 2802 - 3201) 
2019-11-13 11:25:07.779163 0x00007f8828188700: <info> (chain::RemoteApiForwarder.h@66) completed 'synchronizer task' (tier1B-mainnet-chain-peer8 @ proioxis.brimstone.xpxsirius.io:7900) with result Success 
2019-11-13 11:27:17.727823 0x00007f4fd8731700: <info> (src::DiagnosticsService.cpp@39) --- current counter values ---
  MEM CUR RSS : 61
  MEM MAX RSS : 61
 MEM CUR VIRT : 864
  MEM SHR RSS : 37
TOT CONF TXES : 40
     ACNTST C : 15
 ACNTST C HVA : 7
     BLKDIF C : 566
       HASH C : 0
 SECRETLOCK C : 0
     CONFIG C : 1
   CONTRACT C : 0
 REPUTATION C : 0
     MOSAIC C : 6
   MULTISIG C : 0
   METADATA C : 0
   HASHLOCK C : 0
    UPGRADE C : 1
         NS C : 2
      NS C AS : 8
      NS C DS : 8
   PROPERTY C : 0
     UT CACHE : 0
    B WRITERS : 0
      WRITERS : 10
 BLK ELEM TOT : 15
 BLK ELEM ACT : 0
  TX ELEM TOT : 0
  TX ELEM ACT : 0
RB COMMIT ALL : 0
RB COMMIT RCT : 0
RB IGNORE ALL : 0
RB IGNORE RCT : 0
 UNLKED ACCTS : 0
TS OFFSET ABS : 0
TS OFFSET DIR : 0
  TS NODE AGE : 0
 TS TOTAL REQ : 0
 ACTIVE PINGS : 0
  TOTAL PINGS : 0
SUCCESS PINGS : 0
      READERS : 0
        TASKS : 12 

2019-11-13 11:27:18.299695 0x00007f4fd8731700: <info> (chain::ChainSynchronizer.cpp@207) peer returned 400 blocks (heights 16402 - 16801) 
2019-11-13 11:27:18.300154 0x00007f4fd8731700: <info> (chain::RemoteApiForwarder.h@66) completed 'synchronizer task' (tier1B-mainnet-chain-peer8 @ proioxis.brimstone.xpxsirius.io:7900) with result Success 




```



...also observe the Blockchain's height increasing:
```

(disruptor::ConsumerDispatcher.cpp@43) completing processing of element 22 (3 blocks (heights 10725 - 10727) [FBD6794A] from Remote_Pull), last consumer is 0 elements behind 
```


Take note of the "heights" statement in the logs, which shows that the blockchain creates new blocks every 15 seconds (by default).


---


