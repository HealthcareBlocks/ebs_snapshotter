FROM healthcareblocks/alpine:latest

COPY bin/ebs_snapshotter-linux-amd64 /bin/ebs_snapshotter
ENTRYPOINT ["ebs_snapshotter"]
