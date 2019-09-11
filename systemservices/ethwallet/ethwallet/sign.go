package ethwallet

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/mesg-foundation/engine/protobuf/types"
	"github.com/mesg-foundation/engine/systemservices/ethwallet/x/xgo-ethereum/xaccounts"
)

type transaction struct {
	ChainID  int64          `json:"chainID"`
	Nonce    uint64         `json:"nonce"`
	To       common.Address `json:"to"`
	Value    string         `json:"value"`
	Gas      uint64         `json:"gas"`
	GasPrice string         `json:"gasPrice"`
	Data     hexutil.Bytes  `json:"data"`
}

func (s *Ethwallet) sign(inputs *types.Struct) (*types.Struct, error) {
	address := common.HexToAddress(inputs.GetStringValue("address"))
	passphrase := inputs.GetStringValue("passphrase")
	tx := inputs.GetStructValue("transaction")

	account, err := xaccounts.GetAccount(s.keystore, address)
	if err != nil {
		return nil, errAccountNotFound
	}

	value := new(big.Int)
	if _, ok := value.SetString(tx.GetStringValue("value"), 0); !ok {
		return nil, errCannotParseValue
	}

	gasPrice := new(big.Int)
	if _, ok := gasPrice.SetString(tx.GetStringValue("gasPrice"), 0); !ok {
		return nil, errCannotParseGasPrice
	}

	data, err := hex.DecodeString(tx.GetStringValue("data"))
	if err != nil {
		return nil, err
	}

	transaction := ethtypes.NewTransaction(
		uint64(tx.GetNumberValue("nonce")),
		common.HexToAddress(tx.GetStringValue("to")),
		value,
		uint64(tx.GetNumberValue("gas")),
		gasPrice,
		data,
	)

	signedTransaction, err := s.keystore.SignTxWithPassphrase(
		account,
		passphrase,
		transaction,
		big.NewInt(int64(tx.GetNumberValue("chainID"))),
	)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	signedTransaction.EncodeRLP(&buf)
	rawTx := fmt.Sprintf("0x%x", buf.Bytes())

	return types.NewStruct(map[string]*types.Value{
		"signedTransaction": types.NewValueFrom(rawTx),
	}), nil
}
