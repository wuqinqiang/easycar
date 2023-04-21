

# cross_compiles
make -f ./Makefile.cross-compiles
rm -rf ./release/packages
mkdir -p ./release/packages

os_all='linux windows darwin'
arch_all='amd64 arm64'

cd ./release

for os in $os_all; do
    for arch in $arch_all; do
        hello_dir_name="easycar_${os}_${arch}"
        hello_path="./packages/easycar_${os}_${arch}"

        if [ "x${os}" = x"windows" ]; then
            if [ ! -f "./easycar_${os}_${arch}.exe" ]; then
                continue
            fi
            mkdir ${hello_path}
            mv ./easycar_${os}_${arch}.exe ${hello_path}/easycar.exe
        else
            if [ ! -f "./easycar_${os}_${arch}" ]; then
                continue
            fi
            mkdir ${hello_path}
            mv ./easycar_${os}_${arch} ${hello_path}/easycar
        fi
        cp ../LICENSE ${hello_path}
        cp ../conf.example.yml ${hello_path}/conf.yml

        # packages
        cd ./packages
        if [ "x${os}" = x"windows" ]; then
            zip -rq ${hello_dir_name}.zip ${hello_dir_name}
        else
            tar -zcf ${hello_dir_name}.tar.gz ${hello_dir_name}
        fi
        cd ..
        rm -rf ${hello_path}
    done
done
