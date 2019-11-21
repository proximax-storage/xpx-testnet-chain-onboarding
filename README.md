# vagrant-testnet-sirius-peernode
This readme contains the official ProximaX instructions on the the Vagrant distribution for a community-available Sirius Blockchain Peer.  The deployable  Peer automatically hooks-up to ProximaX Public Testnet.

## Prework Instructions for the community


### Local Network requirements
Ensure that your local network can allow inbound/outbound comms on these ports:
```

    3000tcp
    7900tcp
    7901tcp
    7902tcp
```


### System Requirements
The Sirius Blokchain Peer is designed to work with minimum perfomance on   1 CPU core and 2GB RAM, so your Vagrant VM must have these minimum resource specs.

### Install Vagrant on your local machine
In order to run the vagrant distribution, its a prerequisite to install vagrant first.

Access the installation steps [here](https://www.vagrantup.com/intro/getting-started/install.html)

### Download the community-oriented vagrant box for Sirius Peer Node v2
This step will show how to obtain a community peer vagrantbox based on Proximax Sirius Peer v0.4.1:

```bash
mkdir -m 777 ~/proximax-sirius-testnetpeer-v0.4.3
cd ~/proximax-sirius-testnetpeer-v0.4.3
wget https://proximax-vagrant-storage.s3-ap-southeast-1.amazonaws.com/proximax-sirius-testnetpeer-v0.4.3.box
```


Once download completes, setup your vagrant environment to point to the boxfile, then initiatize it:
```bash
vagrant box add proximax-sirius-testnetpeer-v0.4.3 proximax-sirius-testnetpeer-v0.4.3.box
vagrant init proximax-sirius-testnetpeer-v0.4.3
```

At this point, you will have a new Vagrantfile in your current DIR.


Go ahead and standup this box in your local machine:
```bash
vagrant up
```

SSH into your local VM using command:
```bash
vagrant ssh
```




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

# Run the dockerized Sirius Peer inside the Vagrantbox...
Once keys are set, execute your peer to sync to Testnet Chain:
```bash

cd /opt/catapult-config/
sudo docker-compose up

```


---


### Sirius BlockChain Config Files
Inspect the path of the config files of your Sirius Peer Node (located in path /etc/sirius/chain/resources/) :
```bash

vagrant@vagrant:~$ ls -la /opt/catapult-config/resources/
total 96
drwxr--r-- 2 vagrant vagrant 4096 Nov 20 15:51 .
drwxrwxrwx 5 root    root    4096 Nov 20 16:08 ..
-rw-r--r-- 1 vagrant vagrant  655 Nov 20 15:51 config-database.properties
-rw-r--r-- 1 vagrant vagrant  212 Nov 20 15:51 config-extensions-broker.properties
-rw-r--r-- 1 vagrant vagrant  215 Nov 20 15:51 config-extensions-recovery.properties
-rw-r--r-- 1 vagrant vagrant  692 Nov 20 15:51 config-extensions-server.properties
-rw-r--r-- 1 vagrant vagrant  340 Nov 20 15:51 config-harvesting.properties
-rw-r--r-- 1 vagrant vagrant  419 Nov 20 15:51 config-immutable.properties
-rw-r--r-- 1 vagrant vagrant   69 Nov 20 15:51 config-inflation.properties
-rw-r--r-- 1 vagrant vagrant  360 Nov 20 15:51 config-logging-broker.properties
-rw-r--r-- 1 vagrant vagrant  362 Nov 20 15:51 config-logging-recovery.properties
-rw-r--r-- 1 vagrant vagrant  363 Nov 20 15:51 config-logging-server.properties
-rw-r--r-- 1 vagrant vagrant   35 Nov 20 15:51 config-messaging.properties
-rw-r--r-- 1 vagrant vagrant   29 Nov 20 15:51 config-networkheight.properties
-rw-r--r-- 1 vagrant vagrant 2640 Nov 20 15:51 config-network.properties
-rw-r--r-- 1 vagrant vagrant 1529 Nov 20 15:51 config-node.properties
-rw-r--r-- 1 vagrant vagrant   76 Nov 20 15:51 config-pt.properties
-rw-r--r-- 1 vagrant vagrant 1098 Nov 20 15:51 config-task.properties
-rw-r--r-- 1 vagrant vagrant   37 Nov 20 15:51 config-timesync.properties
-rw-r--r-- 1 vagrant vagrant  229 Nov 20 15:51 config-user.properties
-rw-r--r-- 1 vagrant vagrant 1456 Nov 20 15:51 peers-api.json
-rw-r--r-- 1 vagrant vagrant 4592 Nov 20 15:51 peers-p2p.json
-rw-r--r-- 1 vagrant vagrant 2184 Nov 20 15:51 supported-entities.json

```















### Marvel at the chain's behaviour
Once executed, you can easily witness a new chain will be created every 15 seconds by default and the log entry looks like this:
```bash


2019-08-20 10:29:31.094175 0x00007f723fa60700: <info> (chain::RemoteApiForwarder.h@66) completed 'synchronizer task' (api2uswest2 @ XXXX:7900) with result Success 
2019-08-20 10:29:35.932941 0x00007f723fa60700: <info> (src::DiagnosticsService.cpp@39) --- current counter values ---
  MEM CUR RSS : 118
  MEM MAX RSS : 118
 MEM CUR VIRT : 919
  MEM SHR RSS : 37
TOT CONF TXES : 114872
     ACNTST C : 46
 ACNTST C HVA : 31
     BLKDIF C : 526
       HASH C : 77567
 SECRETLOCK C : 0
     CONFIG C : 1
   CONTRACT C : 0
 REPUTATION C : 0
     MOSAIC C : 5
   MULTISIG C : 0
   METADATA C : 0
   HASHLOCK C : 0
    UPGRADE C : 1
         NS C : 1
      NS C AS : 6
      NS C DS : 6
   PROPERTY C : 0
     UT CACHE : 0
    B WRITERS : 0
      WRITERS : 5
 BLK ELEM TOT : 5
 BLK ELEM ACT : 1
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

2019-08-20 10:29:37.009770 0x00007f722d0ab700: <info> (disruptor::ConsumerDispatcher.cpp@43) completing processing of element 5 (400 blocks (heights 19602 - 20001) [862C1F55] from Remote_Pull), last consumer is 0 elements behind 
2019-08-20 10:29:39.896084 0x00007f723fa60700: <info> (ionet::PacketSocket.cpp@450) connected to XXXX [XXXX:7900]
```



...also observe the Blockchain's height increasing:
```

(disruptor::ConsumerDispatcher.cpp@43) completing processing of element 22 (3 blocks (heights 10725 - 10727) [FBD6794A] from Remote_Pull), last consumer is 0 elements behind 
```


Take note of the "heights" statement in the logs, which shows that the blockchain creates new blocks every 15 seconds (by default).




 


