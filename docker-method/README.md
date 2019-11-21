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


### System Requirements
The Sirius Blokchain Peer is designed to work with minimum Hardware perfomance of   1 CPU core and 2GB RAM, so your Vagrant VM must have these minimum resource specs.

As for OS Version requirements, the debian installer will run with these minimums:
```
- Debian 10 ++
- Ubuntu 19.04 ++

```
---

### Prework! Obtain the Dependencies
Setup your node dependencies to properly run the dockerized Sirius Peer.

Feel free to run below commands=.
```bash

#! /bin/bash
sudo apt update -y
sudo apt upgrade -y
sudo curl -L "https://github.com/docker/compose/releases/download/1.24.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo apt remove docker docker-engine docker.io containerd runc -y
sudo apt update && sudo apt upgrade -y
sudo apt install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - && sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository -y "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt update -y
sudo apt install docker-ce docker-ce-cli containerd.io -y

```


### Unpack the Docker-compose Package!

```bash
sudo mkdir -m 777 /opt/catapult-config
cd /opt/catapult-config
wget <blah>


```

---


### Update your OS:
```bash

sudo apt update -y && sudo apt upgrade -y

```

---




### Install the Package with APT:
```bash

sudo apt install ./sirius-chain-0.4.3-2.deb

```

Marvel at how the process installs the dependencies.

---



### Run the Sirius Peer using the MAINNET configs:
```bash

sudo sirius.bc /etc/sirius/chain/mainnet

```


# Staking?
Simply put your harvestKey into configfile:
```
/etc/sirius/chain/mainnet/resources/config-harvesting.properties
```

