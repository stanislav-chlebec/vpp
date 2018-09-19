*** Settings ***
Documentation     Test suite to test basic ping, udp, tcp and dns functionality of the network plugin.
Resource          ${CURDIR}/../libraries/all_libs.robot
Suite Setup       OneNodeK8sSetup
#Suite Teardown    OneNodeK8sTeardown

*** Test Cases ***
Pod_To_Pod_Ping
    [Documentation]    Execute "ping -c 5" command between pods (both ways), require no packet loss.
#    set suite variable ${index}		1
#    set suite variable  ${VM_SSH_ALIAS_PREFIX}${index}   HelloWorld
#    SshCommons.Open_Ssh_Connection    'VM_SSH_ALIAS_PREFIX1'    ${KUBE_CLUSTER_1_VM_1_PUBLIC_IP}    ${KUBE_CLUSTER_1_VM_1_USER}    ${KUBE_CLUSTER_1_VM_1_PSWD}
#    SSHLibrary.Set_Client_Configuration    prompt=${KUBE_CLUSTER_1_VM_1_PROMPT}
#    Get_Machine_Status   'VM_SSH_ALIAS_PREFIX1' 
#
    [Setup]    Setup_Hosts_Connections
#    ${stdout} =    KubernetesEnv.Run_Finite_Command_In_Pod    apt-get update&&apt-get -y upgrade&&apt-get -y install iputils-ping    ssh_session=${client_connection}
#    ${stdout} =    KubernetesEnv.Run_Finite_Command_In_Pod    apt-get update&&apt-get -y upgrade&&apt-get -y install iputils-ping    ssh_session=${server_connection}
    ${stdout} =    KubernetesEnv.Run_Finite_Command_In_Pod    ping -c 5 ${server_ip}    ssh_session=${client_connection}
    BuiltIn.Should_Contain   ${stdout}    5 received, 0% packet loss
    ${stdout} =    KubernetesEnv.Run_Finite_Command_In_Pod    ping -c 5 ${client_ip}    ssh_session=${server_connection}
#Testsuite_Setup    BuiltIn.Should_Contain   ${stdout}    5 received, 0% packet loss
    [Teardown]    Teardown_Hosts_Connections

*** Keywords ***
OneNodeK8sSetup
    [Documentation]    Execute common setup, reinit 1node cluster, deploy client and server pods.
    setup-teardown.Testsuite_Setup
    KubernetesEnv.Reinit_One_Node_Kube_Cluster
    KubernetesEnv.Deploy_Client_And_Server_Pod_And_Verify_Running    ${testbed_connection}

OneNodeK8sTeardown
    [Documentation]    Log leftover output from pods, remove pods, execute common teardown.
    KubernetesEnv.Log_Pods_For_Debug    ${testbed_connection}    exp_nr_vswitch=1
    KubernetesEnv.Remove_Client_And_Server_Pod_And_Verify_Removed    ${testbed_connection}
    setup-teardown.Testsuite Teardown

Setup_Hosts_Connections
    [Documentation]    Open and store two more SSH connections to master host, in them open
    ...    pod shells to client and server pod, parse their IP addresses and store them.
    EnvConnections.Open_Client_Connection
    EnvConnections.Open_Server_Connection

Teardown_Hosts_Connections
    [Documentation]    Exit pod shells, close corresponding SSH connections.
    KubernetesEnv.Leave_Container_Prompt_In_Pod    ${client_connection}
    KubernetesEnv.Leave_Container_Prompt_In_Pod    ${server_connection}
    SSHLibrary.Switch_Connection    ${client_connection}
    SSHLibrary.Close_Connection
    SSHLibrary.Switch_Connection    ${server_connection}
    SSHLibrary.Close_Connection
