# !/bin/bash

# Description: set up the Golang develop environment automatically!
# Author:      jusonalien
# Email:       jusonalien@gmail.com

wget https://storage.googleapis.com/golang/go1.5.3.linux-amd64.tar.gz

tar -zxvf go1.5.3.linux-amd64.tar.gz

golocate="/usr/local/go"
upprojects="home/go"
projects="home/go/projects"

if [ -d "$golocate" ]; then
   rm -rf "$golocate"
fi

mv go /usr/local

if [ -d "$upprojects" ]; then
   rmdir "$upprojects"
   mkdir "$upprojects"
fi

if [ -d "$projects" ]; then
   rmdir "$projects"
   mkdir "$projects"
fi


echo "export GOPATH=/home/go/goprojects" >> ~/.bashrc

echo "export PATH=\$PATH:\$GOPATH/bin" >> ~/.bashrc

echo "export GOROOT=/usr/local/go" >> ~/.bashrc

echo "export PATH=\$PATH:\$GOROOT/bin" >> ~/.bashrc

source ~/.bashrc

