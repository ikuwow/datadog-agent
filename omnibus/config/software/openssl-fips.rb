# Unless explicitly stated otherwise all files in this repository are licensed
# under the Apache License Version 2.0.
# This product includes software developed at Datadog (https:#www.datadoghq.com/).
# Copyright 2016-present Datadog, Inc.

name "openssl-fips"
default_version "0.0.1"

resources_path="#{Omnibus::Config.project_root}/resources/fips"

OPENSSL_VERSION="3.0.13"
OPENSSL_SHA256_SUM="88525753f79d3bec27d2fa7c66aa0b92b3aa9498dafd93d7cfa4b3780cdae313"
OPENSSL_FILENAME="openssl-#{OPENSSL_VER}.tar.gz"

DIST_DIR="#{install_dir}/embedded"

build do
    dependency "openssl-fips-provider"

    source url: "https://www.openssl.org/source/#{OPENSSL_FILENAME}",
           sha256: "#{OPENSSL_SHA256_SUM}",
           extract: :seven_zip
    
    command "tar xvzf #{project_dir}/#{OPENSSL_FILENAME}"

    command "./Configure --prefix=\"#{DIST_DIR}\" \
                --libdir=lib \
                -Wl,-rpath=\"#{DIST_DIR}/lib\" \
                no-asm no-comp no-ssl2 no-ssl3 \
                shared zlib"
    
    command "make depend -j"
    command "make -j"
    command "make install_sw -j"
    command "openssl version -v"

    copy "/usr/local/ssl/fipsmodule.cnf", "#{install_dir}/ssl/fipsmodule.cnf"
    copy "/usr/local/lib/ossl-modules/fips.so", "#{install_dir}/lib/ossl-modules/fips.so"
    copy '#{resources_path}/openssl.cnf', "#{install_dir}/embedded/ssl/openssl.cnf.tmp"
    copy '#{resources_path}/fipsinstall.sh', "#{install_dir}/embedded/bin/fipsinstall.sh"
end 