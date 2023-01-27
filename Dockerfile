FROM golang:1.18
RUN \
    apt-get update && apt-get install -y \
    curl jq vim unzip less \
    && rm -rf /var/lib/apt/lists/* 


RUN useradd -ms /bin/bash developer
RUN chmod 777 -R /go
#USER developer
RUN  go install github.com/spf13/cobra-cli@latest
WORKDIR /home/developer/app
ADD . /home/developer/app/
RUN go build . 
RUN chmod +x cluster-tools
#ENTRYPOINT /home/developer/app/cluster-tools