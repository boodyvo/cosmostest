## Installation

Build

```bash
make install
```

Init config
```bash
ctd init boodyvo --chain-id testchain
ctcli config chain-id testchain
ctcli config output json
ctcli config trust-node true
ctcli config keyring-backend test

ctcli keys add jack
ctcli keys add alice

ctd add-genesis-account $(ctcli keys show jack -a) 1000nametoken,100000000stake
ctd add-genesis-account $(ctcli keys show alice -a) 1000nametoken,100000000stake

ctd gentx --name jack --keyring-backend test
```

## Usage

Start the chain
```bash
ctd start
```

Start the rest-server
```bash
ctcli rest-server --chain-id testchain --trust-node
```

Login jack to the account
```
curl -s http://localhost:1317/auth/accounts/$(ctcli keys show jack -a)
```

Buy the name (need to set correct sequence (incremental) and account-number from login)
```
curl -XPOST -s http://localhost:1317/cosmostest/names --data-binary '{"base_req":{"from":"'$(ctcli keys show jack -a)'","chain_id":"testchain"},"name":"simplename","amount":"5nametoken","buyer":"'$(ctcli keys show jack -a)'"}' > unsignedTx.json
ctcli tx sign unsignedTx.json --from jack --offline --chain-id testchain --sequence 1 --account-number 3 > signedTx.json
ctcli tx broadcast signedTx.json
```

Set some value to the name
```
curl -XPUT -s http://localhost:1317/cosmostest/names --data-binary '{"base_req":{"from":"'$(ctcli keys show jack -a)'","chain_id":"testchain"},"name":"simplename","value":"1512","owner":"'$(ctcli keys show jack -a)'"}' > unsignedTx.json
ctcli tx sign unsignedTx.json --from jack --offline --chain-id testchain --sequence 2 --account-number 3 > signedTx.json
ctcli tx broadcast signedTx.json
```

Get the value for the name
```
curl -s http://localhost:1317/cosmostest/names/simplename
```

List all names (without owners)
```
curl -s http://localhost:1317/cosmostest/names
```
