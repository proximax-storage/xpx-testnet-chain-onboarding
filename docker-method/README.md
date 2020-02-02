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
```
$ wget https://github.com/proximax-storage/xpx-testnet-chain-onboarding/releases/download/release-v0.5.3-buster/public-testnet-peer-package-v0.5.3.tar.gz
$ tar -zxvf public-testnet-peer-package-v0.5.3.tar.gz
$ cd public-testnet-peer-package-v0.5.3
```

## Generate a keypair
```
$ tools/generate_key_pair
```

## Insert private key in [config-user.properties](resources/config-user.properties)

Replace `BOOTKEY_PRIVATE_KEY` with the generated private key. This is the account which holds the node reputation.

```
[account]

bootKey = BOOTKEY_PRIVATE_KEY 

[storage]

dataDirectory = /data
pluginsDirectory = 
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

## Delegated Validating
You may activate your account for delegated validating by running the following tool:
```
$ tools/delegate_validating_testnet
```

After running the above tool, add the delegated remote account private key in the [config-harvesting.properties](resources/config-harvesting.properties):
```
[harvesting]
# private keys are 64 characters
harvestKey = REMOTE_ACCOUNT_PRIVATE_KEY
beneficiary = 0000000000000000000000000000000000000000000000000000000000000000
isAutoHarvestingEnabled = true
maxUnlockedAccounts = 5
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

## Common Issue
*Failure_Core_Insufficient_Balance at block 55476*

Testnet was upgraded at block height `55476` with delta `60`. The delta is lower than default `400`.  Therefore, you will be required to stop the peer node and update the config to overcome this delta.
[config-node.properties](resources/config-node.properties#L13):

```
maxBlocksPerSyncAttempt = 40
```

Restart the peer node and check the logs whether the peer node is able to continue to sync with the block chain.  If it is able to, you may revert the config to the default value
```
maxBlocksPerSyncAttempt = 400
```

Restart the node to allow the config to take into effect.

## Helpdesk
We have a [telegram helpdesk](https://t.me/proximaxhelpdesk) to assist general queries.

For validation-specific queries, you may discuss it at [ProximaX Blockchain Validators Group](https://t.me/xpxtestnetvalidator)
