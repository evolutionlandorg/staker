all    :; source .env.local && dapp --use solc:0.5.16 build
clean  :; dapp clean
test   :; dapp test
deploy :; dapp create Staker
