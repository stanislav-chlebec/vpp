# Contiv - VPP ONS 2018 DEMO

This Kubernetes network plugin uses FD.io VPP to provide network connectivity
between PODs. Currently, only Kubernetes 1.9.X versions are supported. This deployment uses SFC controller to deploy and manage VNF(s).


## Quickstart
Get started with Contiv-VPP using the Vagrant environment:
* Use the [Contiv-VPP Vagrant Installation][1] instructions to start a 
  simulated Kubernetes cluster with a couple of hosts running in VirtualBox
  VMs. This is the easiest way to bring up a cluster for exploring the 
  capabilities and features of Contiv-VPP.
* For the ONS 2018 demo, please use only a production environment (VirtualBox/VMWare Fusion)
   
### Deploying VNF(s) and/or Pods
To get started first ssh into the master node and browse into the ons2018 folder: 
```
vagrant ssh k8s-master
cd /home/vagrant/vpp/k8s/demo/ons2018
```

To run the example, you will need to apply labels to your cluster's nodes. To do so, run the setLabels script located in the ons2018 folder. 

### Scenario 1
Deploy two L2PP service chains in two nodes. The four VNFs should be deployed on master and worker nodes. 
First you need to create two K8s config Maps, common for all the VNFs. 
```
cd /home/vagrant/vpp/k8s/demo/ons2018/vnf-examples
kubectl apply -f configMaps.yaml
```

When the configuration maps are ready issue the following commands:
```
kubectl apply -f vnf1.yaml
kubectl apply -f vnf2.yaml
kubectl apply -f vnf3.yaml
kubectl apply -f vnf4.yaml
```

Verify that VNFs are running on different nodes with:
```
kubectl get po -o wide
```

To clean up: 
```
kubectl delete po vnf1
kubectl delete po vnf2
kubectl delete po vnf3
kubectl delete po vnf4
```

#### The above scenario can also be combined with Pods that won't necessarily use memif interfaces. 
You can test a combined setup by deploying an nginx server and a simple web client pod. 

Run a nginx Pod with labels app=web and expose it at port 80:
```
kubectl run web --image=nginx --labels app=web --expose --port 80
```

Verify that the nginx pod is running and get its IP address: 
```
kubectl get po -o wide
```

Run a temporary Pod and make a request to web Service by using the IP address from the previous output:
```
$ kubectl run --rm -i -t --image=alpine test-$RANDOM -- sh
/ # wget -qO- http://ipaddress
```

[1]: ../../../vagrant/README.md


