# Build the manager binary
FROM golang:1.10.3 as builder

# Copy in the go src
WORKDIR /go/src/github.com/detiber/cluster-api-provider-ansible
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager github.com/detiber/cluster-api-provider-ansible/cmd/manager

FROM ubuntu:latest as kubeadm
RUN apt-get update
RUN apt-get install -y curl
RUN curl -sSL https://dl.k8s.io/release/v1.11.2/bin/linux/amd64/kubeadm > /usr/bin/kubeadm
RUN chmod a+rx /usr/bin/kubeadm

# Copy the controller-manager into a thin image
FROM ubuntu:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/detiber/cluster-api-provider-ansible/manager .
COPY --from=kubeadm /usr/bin/kubeadm /usr/bin/kubeadm
ENTRYPOINT ["./manager"]
