#!/bin/sh

apt-get install -y \
    build-essential \
    cmake \
    deutex \
    libcurl4-gnutls-dev \
    libasound2-dev \
    libegl1-mesa-dev \
    libcairo2-dev \
    libgl1-mesa-dev \
    libglu1-mesa-dev \
    libjsoncpp-dev \
    libpango1.0-dev \
    libpng-dev \
    libportmidi-dev \
    libsdl2-dev \
    libsdl2-mixer-dev \
    libwayland-dev \
    libx11-dev \
    libxft-dev \
    libzstd-dev \
    zlib1g-dev

mkdir -p build/
cd build/

cmake .. \
    -DBUILD_SERVER=0 \
    -DBUILD_CLIENT=1 \
    -DBUILD_LAUNCHER=0 \
    -DBUILD_MASTER=0 \
    -DBUILD_OR_FAIL=1 \
    -DCMAKE_BUILD_TYPE=MinSizeRel \
    -DCMAKE_C_COMPILER=gcc \
    -DCMAKE_C_FLAGS="-O2 -w" \
    -DCMAKE_CXX_COMPILER=g++ \
    -DCMAKE_CXX_FLAGS="-O2 -w"

make -j$(nproc) odamex install
