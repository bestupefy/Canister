Vagrant.configure('2') do |config|
    # grab Ubuntu 14.04 official image
    config.vm.box = "ubuntu/trusty64" # Ubuntu 14.04

    # install Build Dependencies (GOLANG)
    config.vm.provision :shell, :privileged => false, :path => "scripts/vagrant/install-go.sh"

end
