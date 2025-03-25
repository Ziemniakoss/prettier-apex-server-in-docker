ARG BASE_IMAGE_VERSION="19-jdk-alpine3.16"

FROM openjdk:${BASE_IMAGE_VERSION}

ENV PLUGIN_VERSION="1.0.0"
RUN apk update
RUN apk add npm
RUN echo Instlalling version ${PLUGIN_VERSION}
RUN npm install --global prettier-plugin-apex@${PLUGIN_VERSION}
EXPOSE 2117
ENTRYPOINT ["start-apex-server"]
