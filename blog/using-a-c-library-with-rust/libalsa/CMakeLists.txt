cmake_minimum_required(VERSION 3.6)
project(LibAlsa)

# Add the static library
add_library(alsa STATIC alsa.c)

# Install the static library
install(TARGETS alsa DESTINATION .)
