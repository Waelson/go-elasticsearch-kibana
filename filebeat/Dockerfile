FROM docker.elastic.co/beats/filebeat:7.10.0

COPY filebeat.yml /filebeat.yml
COPY entrypoint.sh /entrypoint.sh
#RUN chmod 755 /entrypoint.sh
ENTRYPOINT [ "/entrypoint.sh" ]