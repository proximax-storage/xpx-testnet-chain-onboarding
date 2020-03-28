# ProximaX Sirius TESTNET Onboarding for API node

## OS Requirements
Ensure that your local network allows inbound/outbound traffic on these ports:
- 3000/tcp
- 7900/tcp
- 7901/tcp
- 7902/tcp

A note on System Requirements:
Theoretically, this dockerized Sirius API package can run on any OS running Docker version 19.03.3 and docker-compose version 1.24.0.

But if you really need a minimum benchmark, we have seen the Sirius Blockchain API to work with a minimum Hardware of 1 CPU core and 2GB RAM.

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

A good practice is to use Docker as a non-root.  As highlighted by the post-installation message, you can use Docker as a non-root user with something like:
```
sudo usermod -aG docker $USER
```
where `$USER` is your username.

Installation instructions for docker-compose can be found [here](https://docs.docker.com/compose/install/). 

Enable and start Docker:
```
$ sudo systemctl enable docker.service
$ sudo systemctl start docker.service
$ sudo systemctl status docker.service
```

## Download and Extract the package
```
$ wget https://github.com/proximax-storage/xpx-testnet-chain-onboarding/releases/download/release-v0.5.3-buster/public-testnet-api-package-v0.6.0.tar.gz
$ tar -zxvf public-testnet-api-package-v0.6.0.tar.gz
$ cd public-testnet-api-package-v0.6.0
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

## Update rest gateway configuration file in [rest.json](restuserconfig/rest.json)

Replace `REST_PRIVATE_KEY` and `BOOTKEY_PUBLIC_KEY` with the generated keys from the above mentioned script.

```
{
    "network": {
        "name": "publicTest",
        "description": "ProximaX Sirius Chain Brimstone Test Network"
    },

    "port": 3000,
    "crossDomainHttpMethods": ["GET", "POST", "PUT", "OPTIONS"],
    "cors": "*",
    "clientPrivateKey": "REST_PRIVATE_KEY",
    "extensions": ["accountLink", "accountProperties", "aggregate", "exchange", "config", "lock", "metadata", "mosaic", "multisig", "namespace", "operation", "receipts", "richlist", "service", "supercontract", "transfer", "upgrade"],
    "db": {
        "url": "mongodb://db:27017/",
        "name": "catapult",
        "pageSizeMin": 10,
        "pageSizeMax": 100,
        "pageSizeStep": 25,
        "maxConnectionAttempts": 5,
        "baseRetryDelay": 500
    },

    "apiNode": {
        "host": "catapult-api-node",
        "port": 7900,
        "publicKey": "BOOTKEY_PUBLIC_KEY",
        "timeout": 1000
    },
...
```

## Assign a friendly name in  [config-node.properties](resources/config-node.properties) (OPTIONAL)

Add the domain name or public IP address to the `host` parameter, leave empty to auto-detect. Add a `friendlyName` to assign a name to your node (like: testnet-donaldduck-01).

```
[localnode]

host =
friendlyName =
version = 0
roles = API
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

## Start the API Node
```
$ docker-compose up -d
```

## Check if containers are running
```
$ docker-compose ps
```
You should see something similar like the following:
```
Name                                  Command                        State       Ports
-------------------------------------------------------------------------------------------------
catapult-config_catapult-api-node_1   /bin/bash -c rm -f /data/s ...   Up       0.0.0.0:7900->7900/tcp, 0.0.0.0:7901->7901/tcp, 0.0.0.0:7902->7902/tcp
catapult-config_db_1                  docker-entrypoint.sh bash  ...   Up       27017/tcp
catapult-config_init-db_1             docker-entrypoint.sh bash  ...   Exit 0
catapult-config_rest-gateway_1        docker-entrypoint.sh ash - ...   Up       0.0.0.0:3000->3000/tcp
```

Note:  init-db initializes the db and will exit after DB has been initialize.  If the DB initialization is successful, it should exit with status 0 (i.e. `Exit 0`).

## Stop the API Node
```
$ docker-compose down
```

## Restart the API Node
```
$ docker-compose restart
```

## Reset the API Node
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
*Remote_Pull due to Failure_Chain_Block_Inconsistent_State_Hash*

The error occur when the node fails to update exchange cache after rollback.  This bug has been reported to ProximaX core developers and the developers are looking into resolving this issue.  Meantime, the node will need to be [reset](#Reset-the-API-Node) as mentioned above.

## Helpdesk
We have a [telegram helpdesk](https://t.me/proximaxhelpdesk) to assist general queries.

For validation-specific queries, you may discuss it at [ProximaX Blockchain Validators Group](https://t.me/xpxtestnetvalidator)
