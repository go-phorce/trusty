FROM ghcr.io/ekspand/trusty-docker-base:latest
LABEL org.opencontainers.image.source https://github.com/ekspand/trusty

ENV TRUSTY_DIR=/opt/trusty
ENV PATH=$PATH:$TRUSTY_DIR/bin

RUN mkdir -p $TRUSTY_DIR/bin $TRUSTY_DIR/scripts/sql
COPY ./bin/trusty* $TRUSTY_DIR/bin/
COPY ./scripts/build/* $TRUSTY_DIR/bin/
COPY ./scripts/sql/ $TRUSTY_DIR/scripts/sql/

VOLUME ["/var/trusty/certs", "/var/trusty/logs", "/var/trusty/audit"]

EXPOSE 7880 7891 7892 7893

# Define default command.
CMD ["$TRUSTY_DIR/bin/trusty"]