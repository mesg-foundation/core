package cosmos

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authutils "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// LCD is a simple cosmos LCD client.
type LCD struct {
	endpoint     string
	cdc          *codec.Codec
	kb           keys.Keybase
	chainID      string
	accName      string
	accPassword  string
	minGasPrices sdktypes.DecCoins
	gasLimit     uint64

	// local state
	acc             *auth.BaseAccount
	getAccountMutex sync.Mutex
}

// NewLCD initializes a cosmos LCD client.
func NewLCD(endpoint string, cdc *codec.Codec, kb keys.Keybase, chainID, accName, accPassword, minGasPrices string, gasLimit uint64) (*LCD, error) {
	minGasPricesDecoded, err := sdktypes.ParseDecCoins(minGasPrices)
	if err != nil {
		return nil, err
	}
	return &LCD{
		endpoint:     endpoint,
		cdc:          cdc,
		kb:           kb,
		chainID:      chainID,
		accName:      accName,
		accPassword:  accPassword,
		minGasPrices: minGasPricesDecoded,
		gasLimit:     gasLimit,
	}, nil
}

// Get executes a get request on the LCD.
func (lcd *LCD) Get(path string, ptr interface{}) error {
	resp, err := http.Get(lcd.endpoint + path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("request status code is not 2XX, got %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	cosResp := rest.ResponseWithHeight{}
	if err := lcd.cdc.UnmarshalJSON(body, &cosResp); err != nil {
		return err
	}
	if len(cosResp.Result) > 0 {
		if err := lcd.cdc.UnmarshalJSON(cosResp.Result, ptr); err != nil {
			return err
		}
	}
	return nil
}

// Post executes a post request on the LCD.
// The response is expected to be a ResponseWithHeight that contains the expected result.
func (lcd *LCD) Post(path string, req interface{}, ptr interface{}) error {
	cosResp := rest.ResponseWithHeight{}
	if err := lcd.PostBare(path, req, &cosResp); err != nil {
		return err
	}
	if len(cosResp.Result) > 0 {
		if err := lcd.cdc.UnmarshalJSON(cosResp.Result, ptr); err != nil {
			return err
		}
	}
	return nil
}

// PostBare executes a post request on the LCD.
// There is no expectation on the type of response.
func (lcd *LCD) PostBare(path string, req interface{}, ptr interface{}) error {
	reqBody, err := lcd.cdc.MarshalJSON(req)
	if err != nil {
		return err
	}
	resp, err := http.Post(lcd.endpoint+path, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("request status code is not 2XX, got %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := lcd.cdc.UnmarshalJSON(body, ptr); err != nil {
		return err
	}
	return nil
}

//BroadcastMsg sign and broadcast a transaction from a message.
func (lcd *LCD) BroadcastMsg(msg sdk.Msg) ([]byte, error) {
	return lcd.BroadcastMsgs([]sdk.Msg{msg})
}

//BroadcastMsgs sign and broadcast a transaction from multiple messages.
func (lcd *LCD) BroadcastMsgs(msgs []sdk.Msg) ([]byte, error) {
	tx, err := lcd.createAndSignTx(msgs)
	if err != nil {
		return nil, err
	}
	req := authrest.BroadcastReq{
		Tx:   tx,
		Mode: "block", // TODO: should be changed to "sync" and wait for the tx event
	}
	var res sdk.TxResponse
	if err := lcd.PostBare("txs", req, &res); err != nil {
		return nil, err
	}
	if abci.CodeTypeOK != res.Code {
		return nil, fmt.Errorf("transaction returned with invalid code %d: %s", res.Code, res)
	}
	result, err := hex.DecodeString(res.Data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (lcd *LCD) getAccount() (*auth.BaseAccount, error) {
	lcd.getAccountMutex.Lock()
	defer lcd.getAccountMutex.Unlock()
	if lcd.acc == nil {
		accKb, err := lcd.kb.Get(lcd.accName)
		if err != nil {
			return nil, err
		}
		lcd.acc = auth.NewBaseAccount(accKb.GetAddress(), nil, accKb.GetPubKey(), 0, 0)
	}
	localSeq := lcd.acc.GetSequence()
	if err := lcd.Get("auth/accounts/"+lcd.acc.GetAddress().String(), &lcd.acc); err != nil {
		return nil, err
	}
	// replace seq if sup
	if localSeq > lcd.acc.GetSequence() {
		lcd.acc.SetSequence(localSeq)
	}
	return lcd.acc, nil
}

func (lcd *LCD) createAndSignTx(msgs []sdk.Msg) (authtypes.StdTx, error) {
	// retrieve account
	accR, err := lcd.getAccount()
	if err != nil {
		return authtypes.StdTx{}, err
	}
	sequence := accR.GetSequence()
	accR.SetSequence(accR.GetSequence() + 1)

	// Create TxBuilder
	txBuilder := authtypes.NewTxBuilder(
		authutils.GetTxEncoder(lcd.cdc),
		accR.GetAccountNumber(),
		sequence,
		lcd.gasLimit,
		flags.DefaultGasAdjustment,
		true,
		lcd.chainID,
		"",
		nil,
		lcd.minGasPrices,
	).WithKeybase(lcd.kb)

	// create StdSignMsg
	stdSignMsg, err := txBuilder.BuildSignMsg(msgs)
	if err != nil {
		return authtypes.StdTx{}, err
	}

	// create StdTx
	stdTx := authtypes.NewStdTx(stdSignMsg.Msgs, stdSignMsg.Fee, nil, stdSignMsg.Memo)

	// sign StdTx
	signedTx, err := txBuilder.SignStdTx(lcd.accName, lcd.accPassword, stdTx, false)
	if err != nil {
		return authtypes.StdTx{}, err
	}

	return signedTx, nil
}