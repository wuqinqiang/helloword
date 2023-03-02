

# cross_compiles
make -f ./Makefile.cross-compiles
rm -rf ./release/packages
mkdir -p ./release/packages

os_all='linux windows darwin'
arch_all='amd64 arm64'

cd ./release

for os in $os_all; do
    for arch in $arch_all; do
        hello_dir_name="helloword_${os}_${arch}"
        hello_path="./packages/helloword_${os}_${arch}"

        if [ "x${os}" = x"windows" ]; then
            if [ ! -f "./helloword_${os}_${arch}.exe" ]; then
                continue
            fi
            mkdir ${hello_path}
            mv ./helloword_${os}_${arch}.exe ${hello_path}/helloword.exe
        else
            if [ ! -f "./helloword_${os}_${arch}" ]; then
                continue
            fi
            mkdir ${hello_path}
            mv ./helloword_${os}_${arch} ${hello_path}/helloword
        fi
        cp ../LICENSE ${hello_path}
        cp ../library/* ${hello_path}
        cp ../conf/conf.example.yml ${hello_path}/conf.yml

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
