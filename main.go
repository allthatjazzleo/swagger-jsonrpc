package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/swaggest/jsonrpc"
	"github.com/swaggest/swgui"
	"github.com/swaggest/swgui/v3cdn"
	"github.com/swaggest/usecase"
)

func main() {
	apiSchema := jsonrpc.OpenAPI{}
	apiSchema.Reflector().SpecEns().Info.Title = "Simple Ethereum type JSON-RPC Methods"
	apiSchema.Reflector().SpecEns().Info.Version = "v0.0.1"
	apiSchema.Reflector().SpecEns().Info.WithDescription("This app showcases a Ethereum type JSON-RPC API connecting to Cronos testnet node. Some methods require params, a sample request body is presented in each method correspondingly. You can copy and paste those sample requests for simplicity.")

	h := &jsonrpc.Handler{}
	h.OpenAPI = &apiSchema
	h.Validator = &jsonrpc.JSONSchemaValidator{}
	h.SkipResultValidation = true

	type empty struct{}

	type Input struct {
		JsonRpc string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Id      int           `json:"id"`
		Params  []interface{} `json:"params"`
	}

	type Output struct {
		JsonRpc string `json:"jsonrpc"`
		Result  string `json:"result"`
		Id      int    `json:"id"`
	}

	// web3_clientVersion
	web3_clientVersion := usecase.NewIOI(new(empty), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	web3_clientVersion.SetName("web3_clientVersion")
	web3_clientVersion.SetTags("Web3 Methods")
	web3_clientVersion.SetTitle("Get the web3 client version.")

	// web3_sha3
	web3_sha3 := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	web3_sha3.SetName("web3_sha3")
	web3_sha3.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"web3_sha3","params":["0x67656c6c6f20776f726c64"],"id":1}`)
	web3_sha3.SetTags("Web3 Methods")
	web3_sha3.SetTitle("Returns Keccak-256 (not the standardized SHA3-256) of the given data.")

	// net_version
	net_version := usecase.NewIOI(new(empty), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	net_version.SetName("net_version")
	net_version.SetTags("Net Methods")
	net_version.SetTitle("Returns the current network id.")

	// net_peerCount
	net_peerCount := usecase.NewIOI(new(empty), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	net_peerCount.SetName("net_peerCount")
	net_peerCount.SetTags("Net Methods")
	net_peerCount.SetTitle("Returns the number of peers currently connected to the client.")

	// net_listening
	net_listening := usecase.NewIOI(new(empty), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	net_listening.SetName("net_listening")
	net_listening.SetTags("Net Methods")
	net_listening.SetTitle("Returns if client is actively listening for network connections.")

	// eth_protocolVersion
	eth_protocolVersion := usecase.NewIOI(new(empty), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_protocolVersion.SetName("eth_protocolVersion")
	eth_protocolVersion.SetTags("ETH Methods")
	eth_protocolVersion.SetTitle("Returns the current ethereum protocol version.")

	// eth_syncing
	eth_syncing := usecase.NewIOI(new(empty), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_syncing.SetName("eth_syncing")
	eth_syncing.SetTags("ETH Methods")
	eth_syncing.SetTitle("The sync status object may need to be different depending on the details of Tendermint's sync protocol. However, the 'synced' result is simply a boolean, and can easily be derived from Tendermint's internal sync state.")

	// eth_gasPrice
	eth_gasPrice := usecase.NewIOI(new(empty), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_gasPrice.SetName("eth_gasPrice")
	eth_gasPrice.SetTags("ETH Methods")
	eth_gasPrice.SetTitle("Returns the current gas price in the default EVM denomination parameter.")

	// eth_accounts
	eth_accounts := usecase.NewIOI(new(empty), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_accounts.SetName("eth_accounts")
	eth_accounts.SetTags("ETH Methods")
	eth_accounts.SetTitle("Returns array of all accounts owned by the client.")

	// eth_blockNumber
	eth_blockNumber := usecase.NewIOI(new(empty), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_blockNumber.SetName("eth_blockNumber")
	eth_blockNumber.SetTags("ETH Methods")
	eth_blockNumber.SetTitle("Returns the number of most recent block.")

	// eth_getBalance
	eth_getBalance := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getBalance.SetName("eth_getBalance")
	eth_getBalance.SetTags("ETH Methods")
	eth_getBalance.SetTitle("Returns the balance of the account of given address.")
	eth_getBalance.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getBalance","params":["0x407d73d8a49eeb85d32cf465507dd71d507100c1", "latest"],"id":1}`)

	// eth_getStorageAt
	eth_getStorageAt := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getStorageAt.SetName("eth_getStorageAt")
	eth_getStorageAt.SetTags("ETH Methods")
	eth_getStorageAt.SetTitle("Returns the value from a storage position at a given address.")
	eth_getStorageAt.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getBalance","params":["0x407d73d8a49eeb85d32cf465507dd71d507100c1", "latest"],"id":1}`)

	// eth_getTransactionCount
	eth_getTransactionCount := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getTransactionCount.SetName("eth_getTransactionCount")
	eth_getTransactionCount.SetTags("ETH Methods")
	eth_getTransactionCount.SetTitle("Returns the value from a storage position at a given address.")
	eth_getTransactionCount.SetDescription(`Request body sample: {"jsonrpc":"2.0", "method": "eth_getStorageAt", "params": ["0x295a70b2de5e3953354a6a8344e616ed314d7251", "0x0", "latest"], "id": 1}`)

	// eth_getBlockTransactionCountByHash
	eth_getBlockTransactionCountByHash := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getBlockTransactionCountByHash.SetName("eth_getBlockTransactionCountByHash")
	eth_getBlockTransactionCountByHash.SetTags("ETH Methods")
	eth_getBlockTransactionCountByHash.SetTitle("Returns the number of transactions sent from an address.")
	eth_getBlockTransactionCountByHash.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["0x407d73d8a49eeb85d32cf465507dd71d507100c1","latest"],"id":1}`)

	// eth_getBlockTransactionCountByNumber
	eth_getBlockTransactionCountByNumber := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getBlockTransactionCountByNumber.SetName("eth_getBlockTransactionCountByNumber")
	eth_getBlockTransactionCountByNumber.SetTags("ETH Methods")
	eth_getBlockTransactionCountByNumber.SetTitle("Returns the number of transactions in a block from a block matching the given block hash.")
	eth_getBlockTransactionCountByNumber.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getBlockTransactionCountByHash","params":["0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238"],"id":1}`)

	// eth_getUncleCountByBlockHash
	eth_getUncleCountByBlockHash := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getUncleCountByBlockHash.SetName("eth_getUncleCountByBlockHash")
	eth_getUncleCountByBlockHash.SetTags("ETH Methods")
	eth_getUncleCountByBlockHash.SetTitle("Returns the number of transactions in a block matching the given block number.")
	eth_getUncleCountByBlockHash.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getBlockTransactionCountByNumber","params":["0xe8"],"id":1}`)

	// eth_getUncleCountByBlockNumber
	eth_getUncleCountByBlockNumber := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getUncleCountByBlockNumber.SetName("eth_getUncleCountByBlockNumber")
	eth_getUncleCountByBlockNumber.SetTags("ETH Methods")
	eth_getUncleCountByBlockNumber.SetTitle("Returns the number of uncles in a block from a block matching the given block number.")
	eth_getUncleCountByBlockNumber.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getUncleCountByBlockNumber","params":["0xe8"],"id":1}`)

	// eth_getCode
	eth_getCode := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getCode.SetName("eth_getCode")
	eth_getCode.SetTags("ETH Methods")
	eth_getCode.SetTitle("Returns code at a given address.")
	eth_getCode.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getCode","params":["0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b", "0x2"],"id":1}`)

	// eth_sign
	eth_sign := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_sign.SetName("eth_sign")
	eth_sign.SetTags("ETH Methods")
	eth_sign.SetTitle(`The sign method calculates an Ethereum specific signature with: sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message))).`)
	eth_sign.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_sign","params":["0x9b2055d370f73ec7d8a03e965129118dc8f5bf83", "0xdeadbeaf"],"id":1}`)

	// eth_sendTransaction
	eth_sendTransaction := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_sendTransaction.SetName("eth_sendTransaction")
	eth_sendTransaction.SetTags("ETH Methods")
	eth_sendTransaction.SetTitle("Creates new message call transaction or a contract creation, if the data field contains code.")
	eth_sendTransaction.SetDescription(`Request body sample: { "id": 1, "jsonrpc": "2.0", "method": "eth_sendTransaction", "params": [{ "from": "0xb60e8dd61c5d32be8058bb8eb970870f07233155", "data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675" }] }`)

	// eth_sendRawTransaction
	eth_sendRawTransaction := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_sendRawTransaction.SetName("eth_sendRawTransaction")
	eth_sendRawTransaction.SetTags("ETH Methods")
	eth_sendRawTransaction.SetTitle("Creates new message call transaction or a contract creation for signed transactions.")
	eth_sendRawTransaction.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"],"id":1}`)

	// eth_getBlockByHash
	eth_getBlockByHash := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getBlockByHash.SetName("eth_getBlockByHash")
	eth_getBlockByHash.SetTags("ETH Methods")
	eth_getBlockByHash.SetTitle("Returns information about a block by hash.")
	eth_getBlockByHash.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getBlockByHash","params":["0xdc0818cf78f21a8e70579cb46a43643f78291264dda342ae31049421c82d21ae", false],"id":1}`)

	// eth_getBlockByNumber
	eth_getBlockByNumber := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getBlockByNumber.SetName("eth_getBlockByNumber")
	eth_getBlockByNumber.SetTags("ETH Methods")
	eth_getBlockByNumber.SetTitle("Returns information about a block by block number.")
	eth_getBlockByNumber.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x1b4", true],"id":1}`)

	// eth_getTransactionByHash
	eth_getTransactionByHash := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getTransactionByHash.SetName("eth_getTransactionByHash")
	eth_getTransactionByHash.SetTags("ETH Methods")
	eth_getTransactionByHash.SetTitle("Returns the information about a transaction requested by transaction hash.")
	eth_getTransactionByHash.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":["0x88df016429689c079f3b2f6ad39fa052532c56795b733da78a91ebe6a713944b"],"id":1}`)

	// eth_getTransactionByBlockHashAndIndex
	eth_getTransactionByBlockHashAndIndex := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getTransactionByBlockHashAndIndex.SetName("eth_getTransactionByBlockHashAndIndex")
	eth_getTransactionByBlockHashAndIndex.SetTags("ETH Methods")
	eth_getTransactionByBlockHashAndIndex.SetTitle("Returns information about a transaction by block hash and transaction index position.")
	eth_getTransactionByBlockHashAndIndex.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getTransactionByBlockHashAndIndex","params":["0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b", "0x0"],"id":1}`)

	// eth_getTransactionByBlockNumberAndIndex
	eth_getTransactionByBlockNumberAndIndex := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getTransactionByBlockNumberAndIndex.SetName("eth_getTransactionByBlockNumberAndIndex")
	eth_getTransactionByBlockNumberAndIndex.SetTags("ETH Methods")
	eth_getTransactionByBlockNumberAndIndex.SetTitle("Returns information about a transaction by block number and transaction index position.")
	eth_getTransactionByBlockNumberAndIndex.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getTransactionByBlockNumberAndIndex","params":["0x29c", "0x0"],"id":1}`)

	// eth_getTransactionReceipt
	eth_getTransactionReceipt := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getTransactionReceipt.SetName("eth_getTransactionReceipt")
	eth_getTransactionReceipt.SetTags("ETH Methods")
	eth_getTransactionReceipt.SetTitle("Returns the receipt of a transaction by transaction hash.")
	eth_getTransactionReceipt.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238"],"id":1}`)

	// eth_getUncleByBlockHashAndIndex
	eth_getUncleByBlockHashAndIndex := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getUncleByBlockHashAndIndex.SetName("eth_getUncleByBlockHashAndIndex")
	eth_getUncleByBlockHashAndIndex.SetTags("ETH Methods")
	eth_getUncleByBlockHashAndIndex.SetTitle("Returns information about a uncle of a block by hash and uncle index position.")
	eth_getUncleByBlockHashAndIndex.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getUncleByBlockHashAndIndex","params":["0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b", "0x0"],"id":1}`)

	// eth_getUncleByBlockNumberAndIndex
	eth_getUncleByBlockNumberAndIndex := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getUncleByBlockNumberAndIndex.SetName("eth_getUncleByBlockNumberAndIndex")
	eth_getUncleByBlockNumberAndIndex.SetTags("ETH Methods")
	eth_getUncleByBlockNumberAndIndex.SetTitle("Returns information about a uncle of a block by number and uncle index position.")
	eth_getUncleByBlockNumberAndIndex.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getUncleByBlockNumberAndIndex","params":["0x29c", "0x0"],"id":1}`)

	// eth_newFilter
	eth_newFilter := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_newFilter.SetName("eth_newFilter")
	eth_newFilter.SetTags("ETH Methods")
	eth_newFilter.SetTitle("Creates a filter object, based on filter options, to notify when the state changes (logs).")
	eth_newFilter.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_newFilter","params":[{
		"fromBlock": "0x1",
		"toBlock": "0x2",
		"address": "0x8888f1f195afa192cfee860698584c030f4c9db1",
		"topics": ["0x000000000000000000000000a94f5374fce5edbc8e2a8697c15331677e6ebf0b", null, ["0x000000000000000000000000a94f5374fce5edbc8e2a8697c15331677e6ebf0b", "0x0000000000000000000000000aff3454fce5edbc8cca8697c15331677e6ebccc"]]
	  }],"id":73}`)

	// eth_newBlockFilter
	eth_newBlockFilter := usecase.NewIOI(new(empty), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_newBlockFilter.SetName("eth_newBlockFilter")
	eth_newBlockFilter.SetTags("ETH Methods")
	eth_newBlockFilter.SetTitle("Creates a filter in the node, to notify when a new block arrives.")

	// eth_newPendingTransactionFilter
	eth_newPendingTransactionFilter := usecase.NewIOI(new(empty), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_newPendingTransactionFilter.SetName("eth_newPendingTransactionFilter")
	eth_newPendingTransactionFilter.SetTags("ETH Methods")
	eth_newPendingTransactionFilter.SetTitle("Creates a filter in the node, to notify when new pending transactions arrive.")

	// eth_uninstallFilter
	eth_uninstallFilter := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_uninstallFilter.SetName("eth_uninstallFilter")
	eth_uninstallFilter.SetTags("ETH Methods")
	eth_uninstallFilter.SetTitle("Uninstalls a filter with given id. Should always be called when watch is no longer needed.")
	eth_uninstallFilter.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_uninstallFilter","params":["0xb"],"id":73}`)

	// eth_getFilterChanges
	eth_getFilterChanges := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getFilterChanges.SetName("eth_getFilterChanges")
	eth_getFilterChanges.SetTags("ETH Methods")
	eth_getFilterChanges.SetTitle("Polling method for a filter, which returns an array of logs which occurred since last poll.")
	eth_getFilterChanges.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getFilterChanges","params":["0x16"],"id":73}`)

	// eth_getFilterLogs
	eth_getFilterLogs := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getFilterLogs.SetName("eth_getFilterLogs")
	eth_getFilterLogs.SetTags("ETH Methods")
	eth_getFilterLogs.SetTitle("Returns an array of all logs matching filter with given id.")
	eth_getFilterLogs.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getFilterLogs","params":["0x16"],"id":74}`)

	// eth_getLogs
	eth_getLogs := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_getLogs.SetName("eth_getLogs")
	eth_getLogs.SetTags("ETH Methods")
	eth_getLogs.SetTitle("Returns an array of all logs matching a given filter object.")
	eth_getLogs.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_getLogs","params":[{"topics":["0x000000000000000000000000a94f5374fce5edbc8e2a8697c15331677e6ebf0b"]}],"id":74}`)

	// eth_call
	eth_call := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_call.SetName("eth_call")
	eth_call.SetTags("ETH Methods")
	eth_call.SetTitle("Executes a new message call immediately without creating a transaction on the block chain.")
	eth_call.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_call","params":[{
		"to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567"
	}, "latest"],"id":1}`)

	// eth_estimateGas
	eth_estimateGas := usecase.NewIOI(new(Input), new(Output), func(ctx context.Context, input, output interface{}) error {
		return nil
	})
	eth_estimateGas.SetName("eth_estimateGas")
	eth_estimateGas.SetTags("ETH Methods")
	eth_estimateGas.SetTitle("Generates and returns an estimate of how much gas is necessary to allow the transaction to complete. The transaction will not be added to the blockchain. Note that the estimate may be significantly more than the amount of gas actually used by the transaction, for a variety of reasons including EVM mechanics and node performance.")
	eth_estimateGas.SetDescription(`Request body sample: {"jsonrpc":"2.0","method":"eth_estimateGas","params":[{
		"to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567"
	}, "latest"],"id":1}`)

	h.Add(web3_clientVersion)
	h.Add(web3_sha3)
	h.Add(net_version)
	h.Add(net_peerCount)
	h.Add(net_listening)
	h.Add(eth_protocolVersion)
	h.Add(eth_syncing)
	h.Add(eth_gasPrice)
	h.Add(eth_accounts)
	h.Add(eth_blockNumber)
	h.Add(eth_getBalance)
	h.Add(eth_getStorageAt)
	h.Add(eth_getTransactionCount)
	h.Add(eth_getBlockTransactionCountByHash)
	h.Add(eth_getBlockTransactionCountByNumber)
	h.Add(eth_getUncleCountByBlockHash)
	h.Add(eth_getUncleCountByBlockNumber)
	h.Add(eth_getCode)
	h.Add(eth_sign)
	h.Add(eth_sendTransaction)
	h.Add(eth_sendRawTransaction)
	h.Add(eth_getBlockByHash)
	h.Add(eth_getBlockByNumber)
	h.Add(eth_getTransactionByHash)
	h.Add(eth_getTransactionByBlockHashAndIndex)
	h.Add(eth_getTransactionByBlockNumberAndIndex)
	h.Add(eth_getTransactionReceipt)
	h.Add(eth_getUncleByBlockHashAndIndex)
	h.Add(eth_getUncleByBlockNumberAndIndex)
	h.Add(eth_newFilter)
	h.Add(eth_newBlockFilter)
	h.Add(eth_newPendingTransactionFilter)
	h.Add(eth_uninstallFilter)
	h.Add(eth_getFilterChanges)
	h.Add(eth_getFilterLogs)
	h.Add(eth_getLogs)
	h.Add(eth_call)
	h.Add(eth_estimateGas)

	r := chi.NewRouter()

	r.Mount("/rpc", h)

	// Swagger UI endpoint at /docs/swagger.
	r.Method(http.MethodGet, "/docs/swagger/jsonrpc.json", h.OpenAPI)

	r.Mount("/docs/swagger", v3cdn.NewHandlerWithConfig(swgui.Config{
		Title:       apiSchema.Reflector().SpecEns().Info.Title,
		SwaggerJSON: "/docs/swagger/jsonrpc.json",
		BasePath:    "/docs/swagger",
		SettingsUI:  SwguiSettings(nil),
	}))

	// Start server.
	log.Println("http://localhost:443/docs/swagger")

	if err := http.ListenAndServe(":443", r); err != nil {
		log.Fatal(err)
	}
}

func SwguiSettings(settingsUI map[string]string) map[string]string {
	if settingsUI == nil {
		settingsUI = make(map[string]string)
	}

	settingsUI["requestInterceptor"] = `function(request) {
				if (request.loadSpec) {
					return request;
				}
				var networkUrl = "https://cronos-testnet-3.crypto.org:8545/";
				var url = window.location.protocol + '//'+ window.location.host;
				var method = request.url.substring(url.length);
				var params = '{"jsonrpc": "2.0", "method": "' + method + '", "id": 1, "params": []}';
				if (request.body) {
					params = request.body;
				}
				if (window.location.search.split('=')[1] === 'mainnet') {
					networkUrl = "https://evm-cronos.crypto.org/";
				}

				request.url = networkUrl;
				request.headers = {"Content-Type": "application/json"}
				request.body = params;
				return request;
			}`

	settingsUI["responseInterceptor"] = `function(response) {
		if (response.loadSpec) {
			return response;
		}
		response.headers = {"Content-Type": "application/json"}
		return response;
	}`

	return settingsUI
}
