# registry-watchtower

Basic image ment for testing with [self-hosted container registry](https://github.com/distribution/distribution) and [watchtower](https://github.com/containrrr/watchtower)

Short guide in use follows.

**NOTICE: Watchtower works by pulling and restarting whenever the SHA256 digest of an image changes, it does not change between tags. Thus we push to the `latest` tag whenever we wish to update.**

## Running watchtower + registry

**NOTICE: These do not have to be run on the same machine.**

### Watchtower

We run watchtower with the flag `--interval 30` to make watchtower check every 30 seconds instead of the default 24 hours;
this is done for demonstration purposes, in real environments, you would use a higher interval.

```
docker run --detach \
    --name watchtower \
    --restart=always \
    --volume /var/run/docker.sock:/var/run/docker.sock \
    containrrr/watchtower \
    --interval 30
```

### Registry

The registry runs on port 5000 in the container, so we use port publishing to map it to port 4000 on the host machine. You can use whichever port you like, I just use 4000 for demonstration purposes.

On your self hosted registry, you could also map it to 80\* or 443, which means you don't need to include the port when pulling.

\* It will only look for port 80 automatically if you have defined your registry as an [insecure registry](#insecure-registries), otherwise the default is 443.

```
 docker run --detach \
  --publish 4000:5000 \
  --restart=always \
  --name registry \
  registry:2
```

## Pushing first version of image

### Insecure registries

If your docker registry is hosted on another machine than your own,
you will either need a https certificate for the domain pointing to it,
(since by default docker attempts to connect with https),
or include this in your /etc/docker/daemon.json file, and then restart the docker service:

```json
{
  ...
  "insecure-registries": ["REGISTRY_HOST_IP:4000"]
  ...
}
```

By default 127.0.0.1 is an insecure registry, so you don't need to do it if your development machine, registry machine and watchtower/container machine is the same.

You can verify it worked by running `docker info` and seeing `REGISTRY_HOST_IP:4000` under insecure registries.

### Build + Push

**NOTICE: this should be done on your development machine.**

Build and tag as `.../my-img:v1` and `.../my-img`.

By not including a tag such as `.../my-img`, it implicitly tags it as `.../my-img:latest`.

You are not required to tag your images as `vX`, you could also only tag `latest`, however since this is just as much about using the registry, we use version tags aswell.

You are also not required to follow the format `vX` for your tags, it could be anything as long as it is matches the regex `[a-zA-Z0-9_][a-z-A-Z0-9_.-]{0,127}`.

```
docker build -t REGISTRY_HOST_IP:4000/my-org/my-img:v1 -t REGISTRY_HOST_IP:4000/my-org/my-img .
docker push REGISTRY_HOST_IP:4000/my-org/my-img:v1
docker push REGISTRY_HOST_IP:4000/my-org/my-img
```

### Pull + Run

**NOTICE: this should be done on the machine running watchtower.**

```
docker run -d -p 80:8080 --name "test-image" --restart=always REGISTRY_HOST_IP:4000/my-org/my-img
```

Access your website on `WATCHTOWER_HOST_IP:80`, and confirm the version is 1.

## Pushing the second version of image

Now that version 1 is working, time to push version 2.

1. Edit the version number in src/index.html.

2. Repeat the steps from building and pushing the first version, but tag it as v2 this time:

```
docker build -t REGISTRY_HOST_IP:4000/my-org/my-img:v2 -t REGISTRY_HOST_IP:4000/my-org/my-img .
docker push REGISTRY_HOST_IP:4000/my-org/my-img:v2
docker push REGISTRY_HOST_IP:4000/my-org/my-img
```

Now all you have to do is wait a minute or so until watchtower notices the new update, pulls- and restarts the image.

Now access your website on `WATCHTOWER_HOST_IP:80`, and confirm the version is 2, and you are done.
