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


### Unpack the docker-compose Package!

```bash
sudo mkdir -m 777 /opt/catapult-config
cd /opt/catapult-config
wget https://proximax-vagrant-storage.s3-ap-southeast-1.amazonaws.com/public-testnet-peer-dockerpackage.tar.gz
tar -xvf public-testnet-peer-dockerpackage.tar.gz
cp -R public-testnet-peer-package/* /opt/catapult-config


```

---

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









### Run the Sirius Peer using the MAINNET configs:
```bash

sudo sirius.bc /etc/sirius/chain/mainnet

```


# Staking?
Simply put your harvestKey into configfile:
```
/etc/sirius/chain/mainnet/resources/config-harvesting.properties
```

