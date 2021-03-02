FROM micro/go-micro
FROM Augustu/micro

WORKDIR /
COPY entrypoint.sh /
RUN chmod 755 entrypoint.sh
ENTRYPOINT ["bash", "/entrypoint.sh"]
