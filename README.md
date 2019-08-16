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



