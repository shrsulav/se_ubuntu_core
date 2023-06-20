# Download base image ubuntu 22.04
FROM ubuntu:22.04

# LABEL about the custom image
LABEL maintainer="ssulavshr@gmail.com"
LABEL version="0.1"
LABEL description="Docker Image to build and test linux image with qemu"

# Disable Prompt During Packages Installation
ARG DEBIAN_FRONTEND=noninteractive

ADD src /home/bootable_linux
WORKDIR "/home/bootable_linux"
