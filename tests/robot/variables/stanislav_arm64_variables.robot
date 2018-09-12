*** Settings ***
Resource                          common_variables.robot

*** Variables ***
${KUBE_CLUSTER_1_NODES}            1
${KUBE_CLUSTER_1_VM_1_PUBLIC_IP}   147.75.98.202
${KUBE_CLUSTER_1_VM_1_LOCAL_IP}    147.75.98.202
${KUBE_CLUSTER_1_VM_1_HOST_NAME}   vppagent.pantheon.tech
${KUBE_CLUSTER_1_VM_1_USER}        robot
${KUBE_CLUSTER_1_VM_1_PSWD}        ac1153364449639ba8cb6c23cc6f4cb55ec2d1f6
${KUBE_CLUSTER_1_VM_1_ROLE}        master
${KUBE_CLUSTER_1_VM_1_LABEL}       client_node
${KUBE_CLUSTER_1_DOCKER_COMMAND}   docker


