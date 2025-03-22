# Prettier Apex Server in docker image

[![Docker image version tag](https://img.shields.io/docker/v/ziemniakoss/prettier-apex-server?label=image%20version)](https://hub.docker.com/r/ziemniakoss/prettier-apex-server)
[![badge with image size](https://img.shields.io/docker/image-size/ziemniakoss/prettier-apex-server)](https://hub.docker.com/r/ziemniakoss/prettier-apex-server)

Run apex parsing server in docker to speed up apex formatting with [prettier](https://github.com/dangmai/prettier-plugin-apex).

# How to use

## Developer machine

```shell
docker run -it -p 2117:2117 ziemniakoss/prettier-apex-server:latest
```

And after that whenever you run prettier, add these flags:

```
--apex-standalone-parser built-in
```

## Bitbucket

Make sure that you have these dependencies and scripts in package.json

```json
"devDependencies": {
	"prettier": "^2.8.1",
	"prettier-plugin-apex": "Your selected version",
	"wait-on": "^6.0.1"
},
"scripts": {
	"preprettier-check": "wait-on http://127.0.0.1:2117/api/ast/",
	"prettier-check": "prettier  --apex-standalone-parser built-in --list-different ."
}
```

After that, add prettier apex server definition to bitbucket-pipelines.yml

```yaml
image: atlassian/default-image:3
definitions:
  services:
    prettier-apex:
      image: ziemniakoss/prettier-apex-server:latest # or your selected version
pipelines:
  default:
    - step:
        name: "Validate with prettier"
        services:
          - docker
          - prettier-apex
        caches:
          - docker
        script:
          - npm ci
          - npm run prettier-check
```

## Github

TODO

Just like in bitbucket, this can be implemented using [service containers](https://docs.github.com/en/actions/using-containerized-services/about-service-containers)
