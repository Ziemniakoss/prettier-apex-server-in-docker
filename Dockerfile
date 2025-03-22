ARG PLUGIN_VERSION="1.0.0"
ARG BASE_IMAGE_VERSION="19-jdk-alpine3.16"

FROM openjdk:$BASE_IMAGE_VERSION
RUN apk update
RUN apk add npm
RUN npm install --global prettier-plugin-apex:$PLUGIN_VERSION
EXPOSE 2117
ENTRYPOINT ["start-apex-server"]
