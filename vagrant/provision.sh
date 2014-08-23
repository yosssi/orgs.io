#!/bin/bash

# Define variables.
BASH_PROFILE=/home/vagrant/.bash_profile
GO_HOME=/host/go

GO_VERSION=go1.3.1
GO_FILE_NAME=$GO_VERSION.linux-amd64.tar.gz

NODE_VERSION=v0.10.31
NODE_FILE_NAME=$NODE_VERSION-linux-x64.tar.gz

# Create .bash_profile
touch $BASH_PROFILE
chown vagrant:vagrant $BASH_PROFILE

apt-get update

# Install curl
apt-get install -y curl

# Install Git
apt-get install -y git

# Install Go
curl -o /usr/local/$GO_FILE_NAME https://storage.googleapis.com/golang/$GO_FILE_NAME
tar -C /usr/local -xzf /usr/local/$GO_FILE_NAME
rm /usr/local/$GO_FILE_NAME
echo "export GOROOT=/usr/local/go" >> $BASH_PROFILE
echo "export GOPATH=$GO_HOME" >> $BASH_PROFILE
echo "export PATH=\$PATH:\$GOPATH/bin\$GOROOT:\$GOROOT/bin" >> $BASH_PROFILE
. $BASH_PROFILE

# Install Node.js
curl -o /usr/local/lib/node-$NODE_FILE_NAME http://nodejs.org/dist/$NODE_VERSION/node-$NODE_FILE_NAME
tar -C /usr/local/lib -xzf /usr/local/lib/node-$NODE_FILE_NAME
rm /usr/local/lib/node-$NODE_FILE_NAME
echo "export NODE_HOME=/usr/local/lib/node-$NODE_VERSION-linux-x64" >> $BASH_PROFILE
echo "export PATH=\$PATH:\$NODE_HOME/bin" >> $BASH_PROFILE
. $BASH_PROFILE

# Install Bower
npm install -g bower

# Install Grunt CLI
npm install -g grunt-cli

# Install PostgreSQL
apt-get install -y postgresql postgresql-contrib
