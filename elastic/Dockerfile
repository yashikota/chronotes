ARG ES_VERSION="8.13.4"
FROM docker.elastic.co/elasticsearch/elasticsearch:${ES_VERSION}

ARG ES_VERSION="8.13.4"
ARG SUDACHI_VERSION="3.2.2"
ARG SUDACHI_DIC_NAME="sudachi-dictionary-20240716"
ARG SUDACHI_DIC_NAME_ZIP="${SUDACHI_DIC_NAME}-full.zip"

RUN bin/elasticsearch-plugin install --batch \
    https://github.com/WorksApplications/elasticsearch-sudachi/releases/download/v${SUDACHI_VERSION}/elasticsearch-${ES_VERSION}-analysis-sudachi-${SUDACHI_VERSION}.zip

RUN mkdir -p /usr/share/elasticsearch/config/sudachi

RUN curl -Lo ${SUDACHI_DIC_NAME_ZIP} http://sudachi.s3-website-ap-northeast-1.amazonaws.com/sudachidict/${SUDACHI_DIC_NAME_ZIP} && \
    unzip ${SUDACHI_DIC_NAME_ZIP} && \
    mv ${SUDACHI_DIC_NAME}/system_full.dic /usr/share/elasticsearch/config/sudachi && \
    rm -rf ${SUDACHI_DIC_NAME_ZIP} ${SUDACHI_DIC_NAME}

COPY sudachi.json /usr/share/elasticsearch/config/sudachi
