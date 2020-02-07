FROM ubuntu

RUN apt update
RUN apt install -y wget gnupg software-properties-common
RUN wget -qO - https://packages.confluent.io/deb/5.4/archive.key | apt-key add -
RUN add-apt-repository "deb [arch=amd64] https://packages.confluent.io/deb/5.4 stable main"
RUN apt update
RUN apt install librdkafka-dev -y
ADD kafnostic /usr/bin/

CMD ["/usr/bin/kafnostic"]