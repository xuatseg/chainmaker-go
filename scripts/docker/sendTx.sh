contract_name=fact01

cd ../../tools/cmc
echo -e "\nbuild cmc..."
go build

echo -e "\nprepare cert..."
rm -rf testdata/crypto-config
cp -rf ../../config/crypto-config/ testdata/

echo -e "\nsend transaction install wasmer contract..."
./cmc client contract user create \
--contract-name=$contract_name \
--runtime-type=WASMER \
--byte-code-path=./testdata/claim-wasm-demo/rust-fact-2.0.0.wasm \
--version=1.0 \
--sdk-conf-path=./testdata/sdk_config.yml \
--sync-result=true \
--params="{}"

echo -e "\nsend transaction invoke contract..."
./cmc client contract user invoke \
--contract-name=$contract_name \
--method=save \
--sdk-conf-path=./testdata/sdk_config.yml \
--params="{\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"6543234\"}" \
--sync-result=true \
--result-to-string=true


echo -e "\nsend transaction query contract..."
./cmc client contract user get \
--contract-name=$contract_name \
--method=find_by_file_hash \
--sdk-conf-path=./testdata/sdk_config.yml \
--params="{\"file_hash\":\"ab3456df5799b87c77e7f88\"}" \
--result-to-string=true
