# ProximaX Sirius TESTNET Onboarding

## Pre-requisite
This onboarding method requires `docker` and `docker-compose`.  

Run the following command to install `docker`:
```
$ curl -fsSL https://get.docker.com -o get-docker.sh
$ sh get-docker.sh
```

Installation instructions for docker-compose can be found [here](https://docs.docker.com/compose/install/). 


## Download and Extract the package
```
$ wget https://github.com/proximax-storage/xpx-testnet-chain-onboarding/releases/download/release-v0.4.3-buster/$ public-testnet-peer-package-v0.4.3.tar.gz
$ tar -zxvf public-testnet-peer-package-v0.4.3
$ cd public-testnet-peer-package-v0.4.3
```

## Generate a keypair
```
$ tools/generate_key_pair
```

## Insert private key in [config-user.properties](resources/config-user.properties)

Replace `BOOTKEY_PRIVATE_KEY` with the generated private key

```
[account]

bootKey = BOOTKEY_PRIVATE_KEY 

[storage]

dataDirectory = /data
pluginsDirectory = 
```

## Assign a friendly name in  [config-node.properties](resources/config-node.properties) (OPTIONAL)

Replace `FRIENDLY_NAME` with any name you wish to assign to your node.  Replace localhost with your public IP. (You can find your public IP from https://www.whatismyip.com/)

```
[localnode]

host = localhost
friendlyName = FRIENDLY_NAME
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

## Stop the Peer Node
```
$ docker-compose down
```

## Restart the Peer Node
```
$ docker-compose restart
```

## Check logs
There are 2 ways to view the logs:
1. docker logs
```
$ docker-compose logs --tail=100
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