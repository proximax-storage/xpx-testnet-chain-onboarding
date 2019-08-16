# vagrant-community-sirius-peernode
This is for developing the Vagrant distribution for a community-available Sirius Blockchain Peer.  This deployable  Peer automatically hooks-up to Public Testnet.

## Prework Instructions for the community

### Download the community-oriented vagrant box for Sirius Peer Node
This will download a community peer based on Proximax Sirius Peer v0.4.0:

```
mkdir -m 777 ~/proximax-peer
cd ~/proximax-peer
wget http://3.18.104.76:8000/proximax-sirius-v0.4.0.box
```


Once download completes, setup your vagrant environment to point to the boxfile, then initiatize it:
```
vagrant box add proximax-community-peer proximax-sirius-v0.4.0.box
vagrant init proximax-community-peer
```

At this point, you will have a new Vagrantfile in your current DIR.


Go ahead and standup this box in your local machine:
```
vagrant up
```

SSH into your local VM using command:
```
vagrant ssh
```



## Running the Peer node container inside the VM


Once inside the VM, navigate to the core directory of the Sirius blockchain and initialize the Peer container:
```
cd /opt/catapult-config/
sudo docker-compose -d up


```

Inspect the current running docker container of Peer node:
```
sudo docker ps
```

Expect to see the following output:
```
CONTAINER ID        IMAGE                                                                                       COMMAND                  CREATED             STATUS              PORTS                              NAMES
XXXXX        249767383774.dkr.ecr.ap-southeast-1.amazonaws.com/proximax-catapult-server:release-v0.4.0   "/catapult/bin/siriu…"   6 seconds ago       Up 4 seconds        0.0.0.0:7900-7902->7900-7902/tcp   catapult-config_catapult_1
```

As you can see above, there is a running container inside your VM, with ContainerID "XXXXX"
At this point, you can watch the "default" blockchain harvesting behaviour of our Proximax Testnet by using command:
```

sudo docker logs -f <containerID>
```



...and expect to see logs like this:
```

(disruptor::ConsumerDispatcher.cpp@43) completing processing of element 22 (3 blocks (heights 10725 - 10727) [FBD6794A] from Remote_Pull), last consumer is 0 elements behind 
```


Take note of the "heights" statement in the logs, which shows that the blockchain creates new blocks every 15 seconds (by default).




 


