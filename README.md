To run the docker image given docker daemon is running.

```
docker build -t experiment . --progress plain
docker run -it --rm --name experiment-run experiment
```