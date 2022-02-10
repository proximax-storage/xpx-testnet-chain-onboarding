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
- Ubuntu 20.04 ++
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

### For new peer setup

```sh
wget https://github.com/proximax-storage/xpx-testnet-chain-onboarding/releases/download/release-v0.6.9/public-testnet-peer-package-release-v0.6.9.tar.gz
# verify the SHA256 Hash Checksum is correct
wget https://github.com/proximax-storage/xpx-testnet-chain-onboarding/releases/download/release-v0.6.9/public-testnet-peer-package-release-v0.6.9.tar.gz.sha256
shasum -c public-testnet-peer-package-release-v0.6.9.tar.gz.sha256
# If ok, you have downloaded an authentic file, otherwise the file is corrupted.
tar -xvf public-testnet-peer-package-release-v0.6.9.tar.gz
cd public-testnet-peer-package
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
$ chmod +x tools/delegate_harvesting_testnet
$ tools/delegate_harvesting_testnet
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

Verify that your account has successfully activated delegated validation by checking ProximaX online [explorer](https://bctestnetexplorer.xpxsirius.io) using the generated transaction hash or using your account address.

Please note that if your account does not have any XPX or previously linked to another remote account, the transaction will be unsuccessful.

**For more info, please read our online documentations [here](https://bcdocs.xpxsirius.io/docs/protocol/validating/)**

## Generate a keypair
```
chmod +x tools/generate_key_pair
tools/generate_key_pair
```

## Insert private key in [config-user.properties](resources/config-user.properties)

In [config-user.properties](resources/config-user.properties), replace `BOOTKEY_PRIVATE_KEY` with the generated private key.  This is the account which holds the node reputation.

```
[account]

bootKey = BOOTKEY_PRIVATE_KEY 

[storage]

dataDirectory = /data
pluginsDirectory = 
```

## Start the Peer Node
```
$ chmod +x entrypoint.sh
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
