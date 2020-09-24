# ProximaX Sirius TESTNET Onboarding

## OS Requirements
Ensure that your local network allows inbound/outbound traffic on these ports:
- 3000/tcp
- 7900/tcp
- 7901/tcp
- 7902/tcp

A note on System Requirements:
Theoretically, this dockerized Sirius Peer package can run on any OS running Docker version 19.03.3 and docker-compose version 1.24.0.

But if you really need a minimum benchmark, we have seen the Sirius Blockchain Peer to work with a minimum Hardware of 1 CPU core and 2GB RAM.

This README was prepared by testing the package on:
- Debian 10 ++
- Ubuntu 19.04 ++
- CentOS 7 ++

## Pre-requisite
This onboarding method requires `docker` and `docker-compose`.  

Run the following command to install `docker`:
```
$ curl -fsSL https://get.docker.com -o get-docker.sh
$ sh get-docker.sh
```

Installation instructions for docker-compose can be found [here](https://docs.docker.com/compose/install/). 

Enable and start Docker:
```
$ sudo systemctl enable docker.service
$ sudo systemctl start docker.service
$ sudo systemctl status docker.service
```

## Download and Extract the package

### For new peer setup

**If you are upgrading from a previous version, please skip this section and go to next section below**

```sh
wget https://files.proximax.io/public-testnet-peer-package-latest.tar.gz
# verify the SHA256 Hash Checksum is correct
wget https://files.proximax.io/public-testnet-peer-package-latest.tar.gz.sha256
shasum -c public-testnet-peer-package-latest.tar.gz.sha256
# If ok, you have downloaded an authentic file, otherwise the file is corrupted.
tar -xvf public-testnet-peer-package-latest.tar.gz
# rename folder
mv public-testnet-peer-package-v0.6.5 public-testnet-peer-package
cd public-testnet-peer-package
# download latest snapshot
wget https://proximax-files-dist.s3-ap-southeast-1.amazonaws.com/testnet-data-snapshot.tar.xz
# extract data folder from snapshot
tar -xvf testnet-data-snapshot.tar.xz data
# delete snapshot 
rm testnet-data-snapshot.tar.xz
```

### Upgrading

The following instruction is assuming that existing node installation is located in `~/public-testnet-peer-package`.  If it is different, please change the path accordingly.

Make sure you have `rsync` installed. if not, follow either of the commands below.

```
yum install rsync // using yum 
```
or
```
apt-get install rsync // using advance package tool (apt)
```

After installing `rsync`, run the following commands to pull the latest package.

```sh
# stop docker
cd ~/public-testnet-peer-package
docker-compose down

# download new files in tmp folder
cd /tmp
wget https://files.proximax.io/public-testnet-peer-package-latest.tar.gz
tar -xvf public-testnet-peer-package-latest.tar.gz
# verify the SHA256 Hash Checksum is correct
wget https://files.proximax.io/public-testnet-peer-package-latest.tar.gz.sha256
shasum -c public-testnet-peer-package-latest.tar.gz.sha256
# If ok, you have downloaded an authentic file, otherwise the file is corrupted.
rsync -av --progress \
    --exclude 'data' \
    --exclude 'resources/config-user.properties' \
    --exclude 'resources/config-node.properties' \
    --exclude 'resources/config-harvesting.properties' 
    public-mainnet-peer-package-v0.6.5/ ~/public-mainnet-peer-package

# resume docker
cd ~/public-mainnet-peer-package
docker-compose up -d
```

## Delegated Validating

You may activate your account for delegated validating by running the following tool:
```
$ tools/delegate_validating_testnet
```

After running the above tool, it will generate a private key.  This is the delegated remote account private key.

Insert this private key to [config-harvesting.properties](resources/config-harvesting.properties) and [config-user.properties](resources/config-user.properties)

In [config-harvesting.properties](resources/config-harvesting.properties), replace `REMOTE_ACCOUNT_PRIVATE_KEY` with the generated private key.

```
[harvesting]
# private keys are 64 characters
harvestKey = REMOTE_ACCOUNT_PRIVATE_KEY
beneficiary = 0000000000000000000000000000000000000000000000000000000000000000
isAutoHarvestingEnabled = true
maxUnlockedAccounts = 5
```

## Assign a friendly name in  [config-node.properties](resources/config-node.properties) (OPTIONAL)

Add the domain name or public IP address to the `host` parameter, leave empty to auto-detect. Add a `friendlyName` to assign a name to your node (like: testnet-donaldduck-01).

```
[localnode]

host =
friendlyName =
version = 0
roles = Peer
```

In [config-user.properties](resources/config-user.properties), replace `BOOTKEY_PRIVATE_KEY` with the generated private key.  This is the account which holds the node reputation.

```
[account]

bootKey = BOOTKEY_PRIVATE_KEY 

[storage]

dataDirectory = /data
pluginsDirectory = 
```

Verify that your account has successfully activated delegated validation by checking the explorer using the generated transaction hash or using your account address.

**For more info, please read our online documentations [here](https://bcdocs.xpxsirius.io/docs/protocol/validating/)**

## Start the Peer Node
```
$ docker-compose up -d
```

## Check if container is running
```
$ docker container ls
```

## Stop the Peer Node
```
$ docker-compose down
```

## Restart the Peer Node
```
$ docker-compose restart
```

## Reset the Peer Node
```
$ docker-compose down
$ ./reset.sh
$ docker-compose up -d
```

## Check logs
There are 2 ways to view the logs:
1. docker logs
```
$ docker-compose logs --tail=100 -f
```

2. log files in `logs` directory

## Helpdesk
We have a [telegram helpdesk](https://t.me/proximaxhelpdesk) to assist general queries.

For validation-specific queries, you may discuss it at [ProximaX Blockchain Validators Group](https://t.me/xpxtestnetvalidator)
