FROM ubuntu:14.04
RUN apt-get update \
  && apt-get install -y python3-pip python3-dev \
  && cd /usr/local/bin \
  && ln -s /usr/bin/python3 python \
  && pip3 install --upgrade pip
RUN sudo apt-get -y upgrade  
RUN sudo apt-get -y install curl
RUN sudo curl -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
RUN sudo tar -xvf go1.6.linux-amd64.tar.gz
RUN sudo mv go /usr/local
RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
RUN source ~/.profile