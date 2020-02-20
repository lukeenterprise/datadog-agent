name "aerospike-py2"
default_version "3.10.0"

dependency "aerospike-client"
dependency "pip2"

build do
  env = {
    "DOWNLOAD_C_CLIENT" => "0",
    "AEROSPIKE_C_HOME" => "#{install_dir}/embedded/lib/aerospike",
    "LDFLAGS" => "-L#{install_dir}/embedded/lib -I#{install_dir}/embedded/include",
    "CFLAGS" => "-L#{install_dir}/embedded/lib -I#{install_dir}/embedded/include"
  }

  command "#{install_dir}/embedded/bin/pip2 install --no-binary :all: aerospike==#{version}", :env => env
end