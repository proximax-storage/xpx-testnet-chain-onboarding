# ProximaX Sirius TESTNET Onboarding for API node

## OS Requirements
Ensure that your local network allows inbound/outbound traffic on these ports:
- 3000/tcp
- 7900/tcp
- 7901/tcp
- 7902/tcp

A note on System Requirements:
Theoretically, this dockerized Sirius API package can run on any OS running Docker version 19.03.3 and docker-compose version 1.24.0.

The API node has 3 components:  sirius chain, Mongo and rest gateway.  We were able to run the API node on a AWS EC2 T3a.small with specs of 2 vCPU and 2GiB RAM.  We recommend a min. 2GB RAM as Mongo DB can be memory intensive.

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

### For new API setup 

```sh 
wget https://github.com/proximax-storage/xpx-testnet-chain-onboarding/releases/download/release-v0.6.9/public-testnet-api-package-release-v0.6.9.tar.gz
# verify the SHA256 Hash Checksum is correct
wget https://github.com/proximax-storage/xpx-testnet-chain-onboarding/releases/download/release-v0.6.9/public-testnet-api-package-release-v0.6.9.tar.gz.sha256
shasum -c public-testnet-api-package-release-v0.6.9.tar.gz.sha256
# If ok, you have downloaded an authentic file, otherwise the file is corrupted.
tar -zxvf public-testnet-api-package-release-v0.6.9.tar.gz
cd public-testnet-api-package
```

## Assign keys to Node and Rest

To set up node bootkeys and client rest keys for the API node, run the following script:

```
$ tools/auto_install_keys.sh
```

The script will generate random keypairs and will insert the node bootkey and client rest key into the following files: `resources/config-user.properties` and `restuserconfig/rest.json`

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

## Testing REST API

Run the following to test your API node:

```sh
# check network info:
curl http://localhost:3000/network
# check node chain height:
curl http://localhost:3000/chain/height
# check api node info:
curl http://localhost:3000/node/info
```

If the above is successful, you may make REST API calls to your node.  Run the following to get the public IP address of our node

```sh
curl ifconfig.me
```

Use the ip address that you get from `curl ifconfig.me` and enter into the web-browser as follow http://<node public ip address>:3000/chain/height.

Example:
In the linux shell terminal:  `$curl ifconfig.me` outputs `202.187.132.85`

I will enter the following in my Chrome web browser:
http://202.187.132.85:3000/chain/height
and I should see something like this: `{"height":[1440873,0]}`

If the above fails, please check your node's firewall setting and that port 3000 is accessible from the Internet.

You can add your node to the web explorer `http://explorer.xpxsirius.io`.  Select `node`, `Add node`, key in `http://<ip_address>:3000`, and click `Add`.  Your node should appear in the Node section of the explorer.

Refer to [here](https://bcdocs.xpxsirius.io/endpoints/) to get the full list of available endpoints of **ProximaX Sirius Chain**.

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

## Helpdesk
We have a [telegram helpdesk](https://t.me/proximaxhelpdesk) to assist general queries.

For validation-specific queries, you may discuss it at [ProximaX Blockchain Validators Group](https://t.me/xpxtestnetvalidator)
