package keys

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto"
	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ Keybase = lazyKeybase{}

// NOTE: lazyKeybase will be deprecated in favor of lazyKeybaseKeyring.
type lazyKeybase struct {
	name    string
	dir     string
	config  *sdk.Config
	options []KeybaseOption
}

// New creates a new instance of a lazy keybase.
func New(name, dir string, config *sdk.Config, opts ...KeybaseOption) Keybase {
	if err := cmn.EnsureDir(dir, 0700); err != nil {
		panic(fmt.Sprintf("failed to create Keybase directory: %s", err))
	}

	return lazyKeybase{name: name, dir: dir, config: config, options: opts}
}

func (lkb lazyKeybase) List() ([]Info, error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).List()
}

func (lkb lazyKeybase) Get(name string) (Info, error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).Get(name)
}

func (lkb lazyKeybase) GetByAddress(address sdk.AccAddress) (Info, error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).GetByAddress(address)
}

func (lkb lazyKeybase) Delete(name, passphrase string, skipPass bool) error {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).Delete(name, passphrase, skipPass)
}

func (lkb lazyKeybase) Sign(name, passphrase string, msg []byte) ([]byte, crypto.PubKey, error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).Sign(name, passphrase, msg)
}

func (lkb lazyKeybase) CreateMnemonic(name string, language Language, passwd string, algo SigningAlgo) (info Info, seed string, err error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, "", err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).CreateMnemonic(name, language, passwd, algo)
}

func (lkb lazyKeybase) CreateAccount(name, mnemonic, bip39Passwd, encryptPasswd string, account uint32, index uint32) (Info, error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).CreateAccount(name, mnemonic, bip39Passwd, encryptPasswd, account, index)
}

func (lkb lazyKeybase) Derive(name, mnemonic, bip39Passwd, encryptPasswd string, params hd.BIP44Params) (Info, error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).Derive(name, mnemonic, bip39Passwd, encryptPasswd, params)
}

func (lkb lazyKeybase) CreateLedger(name string, algo SigningAlgo, hrp string, account, index uint32) (info Info, err error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).CreateLedger(name, algo, hrp, account, index)
}

func (lkb lazyKeybase) CreateOffline(name string, pubkey crypto.PubKey) (info Info, err error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).CreateOffline(name, pubkey)
}

func (lkb lazyKeybase) CreateMulti(name string, pubkey crypto.PubKey) (info Info, err error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).CreateMulti(name, pubkey)
}

func (lkb lazyKeybase) Update(name, oldpass string, getNewpass func() (string, error)) error {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).Update(name, oldpass, getNewpass)
}

func (lkb lazyKeybase) Import(name string, armor string) (err error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).Import(name, armor)
}

func (lkb lazyKeybase) ImportPrivKey(name string, armor string, passphrase string) error {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).ImportPrivKey(name, armor, passphrase)
}

func (lkb lazyKeybase) ImportPubKey(name string, armor string) (err error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).ImportPubKey(name, armor)
}

func (lkb lazyKeybase) Export(name string) (armor string, err error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return "", err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).Export(name)
}

func (lkb lazyKeybase) ExportPubKey(name string) (armor string, err error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return "", err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).ExportPubKey(name)
}

func (lkb lazyKeybase) ExportPrivateKeyObject(name string, passphrase string) (crypto.PrivKey, error) {
	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).ExportPrivateKeyObject(name, passphrase)
}

func (lkb lazyKeybase) ExportPrivKey(name string, decryptPassphrase string,
	encryptPassphrase string) (armor string, err error) {

	db, err := sdk.NewLevelDB(lkb.name, lkb.dir)
	if err != nil {
		return "", err
	}
	defer db.Close()

	return newDBKeybase(db, lkb.config, lkb.options...).ExportPrivKey(name, decryptPassphrase, encryptPassphrase)
}

func (lkb lazyKeybase) CloseDB() {}
