# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "heptio/quickstart-ubuntu"
  config.vm.provider "libvirt" do |libvirt|
     libvirt.cpus = 2
     libvirt.memory = 4096
  end

  config.vm.provision "shell", path: "scripts/apiserver.sh", privileged: true
end
