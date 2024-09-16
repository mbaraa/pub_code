cmake_minimum_required(VERSION 3.6)
project(LibCairo)

# Add the static library
add_library(cairo STATIC cairo.c)

# Install the static library
install(TARGETS cairo DESTINATION .)

