# Attempt to operate Contiv - VPP SFC-CONTROLLER DEMO on ARM64

This document is describing the attempt to setup the example of ligato/sfc-controller and 4 VNFs which is present in contiv/vpp repository: https://github.com/contiv/vpp/tree/master/k8s/examples/sfc-controller.

## Prerequisites

Basically the setup which is described in contiv/vpp repository is applicable:
https://github.com/contiv/vpp/tree/master/docs/arm64

* Two bare metal ARM64 machines with Ubuntu 18.04 LTS
```
stanislav@vppagent:~$ lsb_release -a
No LSB modules are available.
Distributor ID:	Ubuntu
Description:	Ubuntu 18.04.2 LTS
Release:	18.04
Codename:	bionic
stanislav@vppagent:~$ 

```
* Kubernetes master is `vppagent` and kubernetes worker is `contivvpp`. 
* There are installed kubernetes v1.13.4 and docker 18.09.3
```
kubeadm version
kubectl version
docker version
```
* swap is disabled in `/etc/fstab` 
```
cat /etc/fstab | grep swap
```
* It is loaded vfio_pci in `/etc/modules` 
```
cat /etc/modules | grep vfio-pci
```
* prepared configuration file for VPP - /etc/vpp/contiv-vswitch 
```
unix {
    nodaemon
    cli-listen /run/vpp/cli.sock
    cli-no-pager
    coredump-size unlimited
    full-coredump
    log /tmp/vpp.log
}
nat {
    endpoint-dependent
    translation hash buckets 1048576
    translation hash memory 268435456
    user hash buckets 1024
    max translations per user 10000
}
api-trace {
   on
   nitems 500
}
dpdk {
    dev 0002:01:00.2
    uio-driver vfio-pci
}
acl-plugin {
     use tuple merge 0
}  
socksvr {
   default
}

```
* interface enP2p1s0f2 is DOWN 
```
stanislav@contivvpp:~$ ip a
...
...
2: enP2p1s0f1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether 30:0e:d5:ff:0d:fd brd ff:ff:ff:ff:ff:ff
    inet 147.75.72.194/30 brd 147.75.72.195 scope global enP2p1s0f1
       valid_lft forever preferred_lft forever
...
...
990: enP2p1s0f2: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 30:0e:d5:ff:0d:fe brd ff:ff:ff:ff:ff:ff
stanislav@contivvpp:~$ 

```
* contivvpp images built present on both servers 
```
stanislav@vppagent:~/work/contiv/vpp/k8s/examples/sfc-controller$ docker images | grep v2.1.3-98-g1e8734b51
contivvpp/ui-arm64                                   v2.1.3-98-g1e8734b51                                            bac615985870        8 days ago          12.3MB
contivvpp/stn-arm64                                  v2.1.3-98-g1e8734b51                                            fe799c60aa11        8 days ago          9.78MB
contivvpp/ksr-arm64                                  v2.1.3-98-g1e8734b51                                            d420224383d4        8 days ago          28.7MB
contivvpp/crd-arm64                                  v2.1.3-98-g1e8734b51                                            b02c2d58ad0b        8 days ago          74.6MB
contivvpp/cni-arm64                                  v2.1.3-98-g1e8734b51                                            9077375bc113        8 days ago          26.4MB
contivvpp/vswitch-arm64                              v2.1.3-98-g1e8734b51                                            00f6a55b4dc8        8 days ago          281MB
contivvpp/dev-vswitch-arm64                          v2.1.3-98-g1e8734b51                                            d1daf5491985        8 days ago          2.72GB
contivvpp/dev-vswitch-arm64                          v2.1.3-98-g1e8734b51-13f5dcf9152287e06b9b5d67774b9f4b576ebaa7   d1daf5491985        8 days ago          2.72GB
stanislav@vppagent:~/work/contiv/vpp/k8s/examples/sfc-controller$ 

```
* helm installed 
```
helm version --tiller-connection-timeout 0
```
* suppose contiv/vpp config without STN - so we need two NICs 
* 1Gb hugepages configured 
```
grep Huge /proc/meminfo
AnonHugePages:         0 kB
ShmemHugePages:        0 kB
HugePages_Total:      16
HugePages_Free:       16
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:    1048576 kB
```
## Prepare NIC enP2p1s0f2: for use with VPP
```
git clone http://dpdk.org/git/dpdk
sudo ./dpdk/usertools/dpdk-devbind.py --bind=vfio-pci 0002:01:00.2
./dpdk/usertools/dpdk-devbind.py --status | grep 0002:01:00.2
```
You should see:
```
0002:01:00.2 'THUNDERX Network Interface Controller virtual function a034' drv=vfio-pci unused=nicvf
```
## Start Kubernetes
```
sudo kubeadm init --token-ttl 0 --pod-network-cidr=10.1.0.0/16
```
At the end of the process you should see:
```
....
....
Your Kubernetes master has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

You can now join any number of machines by running the following on each node
as root:

  kubeadm join 147.75.98.202:6443 --token e7lmfu.mv2x8rwfpp9cj5uf --discovery-token-ca-cert-hash sha256:549fd70c9587aa0b6c1abee2334947f18c108b2dbab2f6530617c38abddd4460

```
Follow instructions about using your cluster as a regular user ....

## Prepare configuration file for contiv/vpp
```
git clone https://github.com/contiv/vpp.git
cd k8s/contiv-vpp/
sed -i 's/tag: latest/tag: v2.1.3-98-g1e8734b51/g' ./values-latest.yaml
sed -i 's/useNodeIP: true/useNodeIP: false/g' ./values.yaml
helm template --name my-release ../contiv-vpp -f ./values-latest.yaml,./values-arm64.yaml,./values.yaml --set vswitch.defineMemoryLimits=true --set vswitch.hugePages1giLimit=8Gi --set vswitch.memoryLimit=8Gi --set etcd.secureTransport=True  > manifest-arm64.yaml
```
You can compare the generates file with prepared ones:
```
diff contiv-vpp.yaml manifest-arm64.yaml
diff contiv-vpp-arm64.yaml manifest-arm64.yaml
```
Copy the prepared configuration file in-place of current:
```
cp manifest-arm64.yaml contiv-vpp-arm64.yaml
```

Note: Prepared contiv-vpp-arm64.yaml file have to have set useNodeIP = false to allow to contivvpp/vswitch on worker node to connect to the container contiv-etcd which is on master node.
 In case useNodeIP = true contivvpp/vswitch oo worker node was unable to connect to etcd



## Deploy kubernetes network plugin contiv/vpp
```
kubectl apply -f ./contiv-vpp-arm64.yaml
```
wait for completelly start contiv containers
You should see
```
kubectl get pods -n kube-system -o wide
NAME                               READY   STATUS    RESTARTS   AGE   IP              NODE       NOMINATED NODE   READINESS GATES
contiv-crd-57wlw                   1/1     Running   0          61s   147.75.98.202   vppagent   <none>           <none>
contiv-etcd-0                      1/1     Running   0          63s   147.75.98.202   vppagent   <none>           <none>
contiv-ksr-kts62                   1/1     Running   0          61s   147.75.98.202   vppagent   <none>           <none>
contiv-vswitch-zw8s4               1/1     Running   0          63s   147.75.98.202   vppagent   <none>           <none>
coredns-86c58d9df4-9254j           1/1     Running   0          69s   10.1.1.3        vppagent   <none>           <none>
coredns-86c58d9df4-kl97c           1/1     Running   0          69s   10.1.1.2        vppagent   <none>           <none>
etcd-vppagent                      1/1     Running   0          44s   147.75.98.202   vppagent   <none>           <none>
kube-apiserver-vppagent            1/1     Running   0          51s   147.75.98.202   vppagent   <none>           <none>
kube-controller-manager-vppagent   1/1     Running   0          56s   147.75.98.202   vppagent   <none>           <none>
kube-proxy-z54m4                   1/1     Running   0          69s   147.75.98.202   vppagent   <none>           <none>
kube-scheduler-vppagent            1/1     Running   0          33s   147.75.98.202   vppagent   <none>           <none>

```

## Start second kubernetes node 
```
  kubeadm join 147.75.98.202:6443 --token e7lmfu.mv2x8rwfpp9cj5uf --discovery-token-ca-cert-hash sha256:549fd70c9587aa0b6c1abee2334947f18c108b2dbab2f6530617c38abddd4460
```

# Contiv VPP SFC-CONTROLLER example
See https://github.com/stanislav-chlebec/vpp/tree/mydev/k8s/examples/sfc-controller
Here are latest changes related to this example.

## Set node labels
https://github.com/stanislav-chlebec/vpp/blob/mydev/k8s/examples/sfc-controller/set-node-labels

Comment on changes to adjusting to your premises:
- you need to replace the name of servers, `vppagent` is in my case kubernetes master, and `contivvpp` is on my case kubernetes worker node.

Then start the bash script:
```
./set-node-labels  
```
The output should be like this:
```
node/vppagent untainted
error: taint "node-role.kubernetes.io/master:" not found
node/vppagent labeled
node/contivvpp labeled

```
You can check if the nodes were labelled:
```
kubectl describe node -n kube-system vppagent | grep aff -C10
kubectl describe node -n kube-system contivvpp | grep aff -C10
```
Note: run this script only after deploying contiv/vpp network

## Deploy sfc-controller
See https://github.com/stanislav-chlebec/vpp/blob/mydev/k8s/examples/sfc-controller/sfc-controller.yaml

Comment on changes to adjusting to your premises:
- this container is using hostworking (hostNetwork: true) and it runs on master. For connecting to contiv-etcd it is correct to specify etcd endpoint like localhost
```
 endpoints:
      - "127.0.0.1:32379"
```
- in case you use for etcd `insecure-transport: false` you need to specify for container volumeMount and uses volume.
```
      containers:
...
... 
          volumeMounts:
...
...
            - name: etcd-secrets
              mountPath: /var/contiv/etcd-secrets
              readOnly: true
...
...
      volumes:
...
...
        - name: etcd-secrets
          hostPath:
            path: /var/contiv/etcd-secrets

```
- specify the used docker image
```
          image: ligato/sfc-controller-arm64:dev
```
- Check the configMap etcd.conf . The settings for etcd should be the same as in contiv-vpp-arm64.yaml. You can try both possibilities for insecure-transport (true / false)
```
  etcd.conf: |
    insecure-transport: false
    ca-file: /var/contiv/etcd-secrets/ca.pem
    cert-file: /var/contiv/etcd-secrets/client.pem
    key-file: /var/contiv/etcd-secrets/client-key.pem
    dial-timeout: 10000000000
    allow-delayed-start: true
    endpoints:
      - "127.0.0.1:32379"

```

- Check the configMap sfc.conf . You need to update the data about network nodes
```
    network_nodes:
      - metadata:
          name: vppagent 
        spec:
          node_type: host
          interfaces:
            - name: VirtualFunctionEthernet1/0/2
              bypass_renderer: true
              if_type: ethernet
              ip_addresses:
                - "192.168.16.1/24"
      - metadata:
          name: contivvpp
        spec:
          node_type: host
          interfaces:
            - name: VirtualFunctionEthernet1/0/2
              bypass_renderer: true
              if_type: ethernet
              ip_addresses:
                - "192.168.16.2/24"

```
 names of netwrok nodes were changed (vppagent and contivvpp) 
 names of interfaces were changed (VirtualFunctionEthernet1/0/2)

You can find out the names of interfaces like this:
```
# CHECK THE MASTER NODE - VPP
masterIP=`kubectl get nodes -o wide --no-headers=true | grep master | awk '{print $6}'`
contivvswitchcontainer=`kubectl get pods -o wide --no-headers=true --include-uninitialized -n kube-system | grep $masterIP | awk '{print $1}' | grep 'contiv-vswitch'`
kubectl exec -n kube-system $contivvswitchcontainer -- /usr/bin/vppctl show interface addr

```
You will see output like this
```
VirtualFunctionEthernet1/0/2 (up):
  L3 192.168.16.1/24
local0 (dn):
loop0 (up):
  L3 10.1.1.1/24 ip4 table-id 1 fib-idx 1
loop1 (up):
  L2 bridge bd-id 1 idx 1 shg 1 bvi
  L3 192.168.30.1/24 ip4 table-id 1 fib-idx 1
tap0 (up):
  L3 172.30.1.1/24
tap1 (up): 
  unnumbered, use loop0
  L3 10.1.1.1/24 ip4 table-id 1 fib-idx 1
tap2 (up): 
  unnumbered, use loop0
  L3 10.1.1.1/24 ip4 table-id 1 fib-idx 1

```
```
# CHECK THE  WORKER NODE - VPP
notmasterIP=`kubectl get nodes -o wide --no-headers=true | grep -v  master | awk '{print $6}'`
contivvswitchcontainer=`kubectl get pods -o wide --no-headers=true --include-uninitialized -n kube-system | grep $notmasterIP | awk '{print $1}' | grep 'contiv-vswitch'`
kubectl exec -n kube-system $contivvswitchcontainer -- /usr/bin/vppctl show interface addr

```
You will see output like this
```
VirtualFunctionEthernet1/0/2 (up):
  L3 192.168.16.2/24
local0 (dn):
loop0 (up):
  L3 10.1.2.1/24 ip4 table-id 1 fib-idx 1
loop1 (up):
  L2 bridge bd-id 1 idx 1 shg 1 bvi
  L3 192.168.30.2/24 ip4 table-id 1 fib-idx 1
tap0 (up):
  L3 172.30.2.1/24
vxlan_tunnel0 (up):
  L2 bridge bd-id 1 idx 1 shg 1  

```

## Deploy sfc-controller
```
kubectl apply -f sfc-controller.yaml

```

## Check configmaps prepared for VNFs
See https://github.com/stanislav-chlebec/vpp/blob/mydev/k8s/examples/sfc-controller/vnf-pods/configMaps.yaml
You need to prepare configMap etcd.conf - accordingly of the value of insecure-transport
Here you need to use IP address of kubernetes master node for etcd endpoint because VNFs do not use host networking. Also you need to copy the files for securing connection from kubernetes master to kubernetes worker. (in case you use insecure-transport:false)
```
  etcd.conf: |
    insecure-transport: false
    ca-file: /var/contiv/etcd-secrets/ca.pem
    cert-file: /var/contiv/etcd-secrets/client.pem
    key-file: /var/contiv/etcd-secrets/client-key.pem
    dial-timeout: 10000000000
    allow-delayed-start: true
    endpoints:
      - "147.75.98.202:32379"

```

You need to prepare configMap for vpp.conf. This is the configuration file used by VNFs` VPPs. This file was adjusted according the file used by contivvpp/vswitch containers. Because version 2 of ligato/vpp-agent uses by default sockets you need to have in this configuration file the clause 
```
    socksvr {
      default
    }

Note: The old way of communication was via shared memmory – you can set GOVPPMUX_NOSOCK = 1 to revert to the old way – not recommended. Currently: default govpp adapter to socketclient.

```
## Deploy config maps used by VNFs
```
kubectl apply -f vnf-pods/configMaps.yaml

```

## Check and adjust VNFs
See https://github.com/stanislav-chlebec/vpp/tree/mydev/k8s/examples/sfc-controller/vnf-pods
Comment on changes:
- change the docker image used
```
      image: ligato/vpp-agent-arm64:dev

```
- add the environment variable ETCD_CONFIG to indicate to VNF where it should search for the configuration file (the content of which is defined by configmap highlighted above)
```
      env:
        - name: MICROSERVICE_LABEL
          value: vnf1
#       - name: RETAIN_SUPERVISOR
#         value: "y"
        - name: ETCD_CONFIG
          value: "/opt/vpp-agent/dev/etcd.conf"

```
Note: the variable RETAIN_SUPERVISOR can be used for debugging purposes in case that VNF pods are restarting

- in case you use for etcd `insecure-transport: false` you need to specify for container volumeMount and uses volume.
```
      containers:
...
... 
          volumeMounts:
...
...
            - name: etcd-secrets
              mountPath: /var/contiv/etcd-secrets
              readOnly: true
...
...
      volumes:
...
...
        - name: etcd-secrets
          hostPath:
            path: /var/contiv/etcd-secrets

```
- define resources (hugepages and memory)
```
      resources:
        limits:
          hugepages-1Gi: 2Gi
          memory: 2Gi
        requests:
          cpu: 250m

```


## Start VNFs
```
kubectl apply -f vnf-pods/vnf1.yaml
kubectl apply -f vnf-pods/vnf2.yaml
kubectl apply -f vnf-pods/vnf3.yaml
kubectl apply -f vnf-pods/vnf4.yaml
```

## Check started VNFs
```
kubectl get pods
```

## How to access etcd database
```
kubectl exec -i --tty -n kube-system contiv-etcd-0 -- etcdctl get --endpoints=127.0.0.1:12379 --key /var/contiv/etcd-secrets/client-key.pem  --cert /var/contiv/etcd-secrets/client.pem --cacert=/var/contiv/etcd-secrets/ca.pem --prefix="true" ""
```
