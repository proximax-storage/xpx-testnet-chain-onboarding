# vagrant-community-sirius-peernode
This is for developing the Vagrant distribution for a community-available Sirius Blockchain Peer.  This deployable  Peer automatically hooks-up to Public Testnet.

## Prework Instructions for the community

### Download the community-oriented vagrant box for Sirius Peer Node
This will download a community peer based on Proximax Sirius Peer v0.4.0:

```
wget http://3.18.104.76:8000/proximax-sirius-v0.4.0.box
```


Once downloaded, setup your vagrant environment to point to the boxfile:
```
vagrant box add proximax-community-peer proximax-sirius-v0.4.0.box 
```



