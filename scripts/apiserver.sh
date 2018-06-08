#!/bin/bash

set -e

echo "##############################################"
echo "# WARNING: not intended for production use"
echo "##############################################"
echo

sudo apt-get update
sudo apt-get install -y curl software-properties-common
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
sudo apt-get update
sudo apt-get install kubeadm kubectl kubelet
sudo kubeadm init
mkdir -p /home/vagrant/.kube
sudo cp -i /etc/kubernetes/admin.conf /home/vagrant/.kube/config
sudo chown -R vagrant:vagrant /home/vagrant/.kube
kubectl --kubeconfig /etc/kubernetes/admin.conf apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"
kubectl --kubeconfig /etc/kubernetes/admin.conf taint nodes --all node-role.kubernetes.io/master-

# install cfssl tools
curl https://pkg.cfssl.org/R1.2/cfssl_linux-amd64 -o /usr/local/bin/cfssl
curl https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64 -o /usr/local/bin/cfssljson
chmod +x /usr/local/bin/cfssl /usr/local/bin/cfssljson

mkdir cluster-api-ca
pushd cluster-api-ca
cat << EOF > ca-csr.json
{
    "CN": "cluster-api-ca",
    "key": {
        "algo": "rsa",
        "size": 4096
    },
    "ca": {
        "expiry": "26280h"
    }
}
EOF
cfssl genkey -initca ca-csr.json | cfssljson -bare ca
cfssl print-defaults config > ca-config.json
cat << EOF > cluster-api-csr.json
{
    "hosts": [
        "cluster-api.svc.cluster.local",
        "cluster-api.svc"
    ],
    "cn": "cluster-api.cluster-api.svc",
    "key": {
        "algo": "rsa",
        "size": 2048
    }
}
EOF
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem --config=ca-config.json -profile=www cluster-api-csr.json | cfssljson -bare cluster-api
popd

# load the cert and key as secrets
kubectl --kubeconfig /etc/kubernetes/admin.conf create secret tls cluster-apiserver-certs --cert=cluster-api-ca/cluster-api.pem --key=cluster-api-ca/cluster-api-key.pem
kubectl --kubeconfig /etc/kubernetes/admin.conf apply -f /vagrant/manifests/clusterapi.yaml