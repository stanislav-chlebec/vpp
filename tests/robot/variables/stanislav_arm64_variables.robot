*** Settings ***
Resource                          common_variables.robot

*** Variables ***
${KUBE_CLUSTER_1_NODES}            1
${KUBE_CLUSTER_1_VM_1_PUBLIC_IP}   147.75.98.202
${KUBE_CLUSTER_1_VM_1_LOCAL_IP}    147.75.98.202
${KUBE_CLUSTER_1_VM_1_HOST_NAME}   vppagent.pantheon.tech
${KUBE_CLUSTER_1_VM_1_USER}        robotusername
${KUBE_CLUSTER_1_VM_1_PSWD}        robotpassword
${KUBE_CLUSTER_1_VM_1_ROLE}        master
${KUBE_CLUSTER_1_VM_1_LABEL}       client_node

${KUBE_CLUSTER_1_NODES}            1
${KUBE_CLUSTER_1_VM_2_PUBLIC_IP}   147.75.72.194
${KUBE_CLUSTER_1_VM_2_LOCAL_IP}    147.75.72.194
${KUBE_CLUSTER_1_VM_2_HOST_NAME}   contivvpp.pantheon.tech
${KUBE_CLUSTER_1_VM_2_USER}        robotusername
${KUBE_CLUSTER_1_VM_2_PSWD}        robotpassword
${KUBE_CLUSTER_1_VM_2_ROLE}        slave
${KUBE_CLUSTER_1_VM_2_LABEL}       client_node

${KUBE_CLUSTER_1_DOCKER_COMMAND}   docker
${ARCHITECTURE}                    -arm64


