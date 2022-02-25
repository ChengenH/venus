// Code generated by github.com/filecoin-project/venus/venus-devtool/api-gen. DO NOT EDIT.
package wallet

import (
	"context"

	address "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/venus/venus-shared/types"
	"github.com/filecoin-project/venus/venus-shared/types/wallet"
)

type IStrategyVerifyStruct struct {
	Internal struct {
		ContainWallet func(ctx context.Context, address address.Address) bool                                             `perm:"admin"`
		ScopeWallet   func(ctx context.Context) (*wallet.AddressScope, error)                                             `perm:"admin"`
		Verify        func(ctx context.Context, address address.Address, msgType types.MsgType, msg *types.Message) error `perm:"admin"`
	}
}

func (s *IStrategyVerifyStruct) ContainWallet(p0 context.Context, p1 address.Address) bool {
	return s.Internal.ContainWallet(p0, p1)
}
func (s *IStrategyVerifyStruct) ScopeWallet(p0 context.Context) (*wallet.AddressScope, error) {
	return s.Internal.ScopeWallet(p0)
}
func (s *IStrategyVerifyStruct) Verify(p0 context.Context, p1 address.Address, p2 types.MsgType, p3 *types.Message) error {
	return s.Internal.Verify(p0, p1, p2, p3)
}

type IStrategyStruct struct {
	Internal struct {
		AddMethodIntoKeyBind     func(ctx context.Context, name string, methods []string) (*wallet.KeyBind, error)                               `perm:"admin"`
		AddMsgTypeIntoKeyBind    func(ctx context.Context, name string, codes []int) (*wallet.KeyBind, error)                                    `perm:"admin"`
		GetGroupByName           func(ctx context.Context, name string) (*wallet.Group, error)                                                   `perm:"admin"`
		GetKeyBindByName         func(ctx context.Context, name string) (*wallet.KeyBind, error)                                                 `perm:"admin"`
		GetKeyBinds              func(ctx context.Context, address address.Address) ([]*wallet.KeyBind, error)                                   `perm:"admin"`
		GetMethodTemplateByName  func(ctx context.Context, name string) (*wallet.MethodTemplate, error)                                          `perm:"admin"`
		GetMsgTypeTemplate       func(ctx context.Context, name string) (*wallet.MsgTypeTemplate, error)                                         `perm:"admin"`
		GetWalletTokenInfo       func(ctx context.Context, token string) (*wallet.GroupAuth, error)                                              `perm:"admin"`
		GetWalletTokensByGroup   func(ctx context.Context, groupName string) ([]string, error)                                                   `perm:"admin"`
		ListGroups               func(ctx context.Context, fromIndex, toIndex int) ([]*wallet.Group, error)                                      `perm:"admin"`
		ListKeyBinds             func(ctx context.Context, fromIndex, toIndex int) ([]*wallet.KeyBind, error)                                    `perm:"admin"`
		ListMethodTemplates      func(ctx context.Context, fromIndex, toIndex int) ([]*wallet.MethodTemplate, error)                             `perm:"admin"`
		ListMsgTypeTemplates     func(ctx context.Context, fromIndex, toIndex int) ([]*wallet.MsgTypeTemplate, error)                            `perm:"admin"`
		NewGroup                 func(ctx context.Context, name string, keyBindNames []string) error                                             `perm:"admin"`
		NewKeyBindCustom         func(ctx context.Context, name string, address address.Address, codes []int, methods []wallet.MethodName) error `perm:"admin"`
		NewKeyBindFromTemplate   func(ctx context.Context, name string, address address.Address, mttName, mtName string) error                   `perm:"admin"`
		NewMethodTemplate        func(ctx context.Context, name string, methods []string) error                                                  `perm:"admin"`
		NewMsgTypeTemplate       func(ctx context.Context, name string, codes []int) error                                                       `perm:"admin"`
		NewStToken               func(ctx context.Context, groupName string) (token string, err error)                                           `perm:"admin"`
		RemoveGroup              func(ctx context.Context, name string) error                                                                    `perm:"admin"`
		RemoveKeyBind            func(ctx context.Context, name string) error                                                                    `perm:"admin"`
		RemoveKeyBindByAddress   func(ctx context.Context, address address.Address) (int64, error)                                               `perm:"admin"`
		RemoveMethodFromKeyBind  func(ctx context.Context, name string, methods []string) (*wallet.KeyBind, error)                               `perm:"admin"`
		RemoveMethodTemplate     func(ctx context.Context, name string) error                                                                    `perm:"admin"`
		RemoveMsgTypeFromKeyBind func(ctx context.Context, name string, codes []int) (*wallet.KeyBind, error)                                    `perm:"admin"`
		RemoveMsgTypeTemplate    func(ctx context.Context, name string) error                                                                    `perm:"admin"`
		RemoveStToken            func(ctx context.Context, token string) error                                                                   `perm:"admin"`
	}
}

func (s *IStrategyStruct) AddMethodIntoKeyBind(p0 context.Context, p1 string, p2 []string) (*wallet.KeyBind, error) {
	return s.Internal.AddMethodIntoKeyBind(p0, p1, p2)
}
func (s *IStrategyStruct) AddMsgTypeIntoKeyBind(p0 context.Context, p1 string, p2 []int) (*wallet.KeyBind, error) {
	return s.Internal.AddMsgTypeIntoKeyBind(p0, p1, p2)
}
func (s *IStrategyStruct) GetGroupByName(p0 context.Context, p1 string) (*wallet.Group, error) {
	return s.Internal.GetGroupByName(p0, p1)
}
func (s *IStrategyStruct) GetKeyBindByName(p0 context.Context, p1 string) (*wallet.KeyBind, error) {
	return s.Internal.GetKeyBindByName(p0, p1)
}
func (s *IStrategyStruct) GetKeyBinds(p0 context.Context, p1 address.Address) ([]*wallet.KeyBind, error) {
	return s.Internal.GetKeyBinds(p0, p1)
}
func (s *IStrategyStruct) GetMethodTemplateByName(p0 context.Context, p1 string) (*wallet.MethodTemplate, error) {
	return s.Internal.GetMethodTemplateByName(p0, p1)
}
func (s *IStrategyStruct) GetMsgTypeTemplate(p0 context.Context, p1 string) (*wallet.MsgTypeTemplate, error) {
	return s.Internal.GetMsgTypeTemplate(p0, p1)
}
func (s *IStrategyStruct) GetWalletTokenInfo(p0 context.Context, p1 string) (*wallet.GroupAuth, error) {
	return s.Internal.GetWalletTokenInfo(p0, p1)
}
func (s *IStrategyStruct) GetWalletTokensByGroup(p0 context.Context, p1 string) ([]string, error) {
	return s.Internal.GetWalletTokensByGroup(p0, p1)
}
func (s *IStrategyStruct) ListGroups(p0 context.Context, p1, p2 int) ([]*wallet.Group, error) {
	return s.Internal.ListGroups(p0, p1, p2)
}
func (s *IStrategyStruct) ListKeyBinds(p0 context.Context, p1, p2 int) ([]*wallet.KeyBind, error) {
	return s.Internal.ListKeyBinds(p0, p1, p2)
}
func (s *IStrategyStruct) ListMethodTemplates(p0 context.Context, p1, p2 int) ([]*wallet.MethodTemplate, error) {
	return s.Internal.ListMethodTemplates(p0, p1, p2)
}
func (s *IStrategyStruct) ListMsgTypeTemplates(p0 context.Context, p1, p2 int) ([]*wallet.MsgTypeTemplate, error) {
	return s.Internal.ListMsgTypeTemplates(p0, p1, p2)
}
func (s *IStrategyStruct) NewGroup(p0 context.Context, p1 string, p2 []string) error {
	return s.Internal.NewGroup(p0, p1, p2)
}
func (s *IStrategyStruct) NewKeyBindCustom(p0 context.Context, p1 string, p2 address.Address, p3 []int, p4 []wallet.MethodName) error {
	return s.Internal.NewKeyBindCustom(p0, p1, p2, p3, p4)
}
func (s *IStrategyStruct) NewKeyBindFromTemplate(p0 context.Context, p1 string, p2 address.Address, p3, p4 string) error {
	return s.Internal.NewKeyBindFromTemplate(p0, p1, p2, p3, p4)
}
func (s *IStrategyStruct) NewMethodTemplate(p0 context.Context, p1 string, p2 []string) error {
	return s.Internal.NewMethodTemplate(p0, p1, p2)
}
func (s *IStrategyStruct) NewMsgTypeTemplate(p0 context.Context, p1 string, p2 []int) error {
	return s.Internal.NewMsgTypeTemplate(p0, p1, p2)
}
func (s *IStrategyStruct) NewStToken(p0 context.Context, p1 string) (string, error) {
	return s.Internal.NewStToken(p0, p1)
}
func (s *IStrategyStruct) RemoveGroup(p0 context.Context, p1 string) error {
	return s.Internal.RemoveGroup(p0, p1)
}
func (s *IStrategyStruct) RemoveKeyBind(p0 context.Context, p1 string) error {
	return s.Internal.RemoveKeyBind(p0, p1)
}
func (s *IStrategyStruct) RemoveKeyBindByAddress(p0 context.Context, p1 address.Address) (int64, error) {
	return s.Internal.RemoveKeyBindByAddress(p0, p1)
}
func (s *IStrategyStruct) RemoveMethodFromKeyBind(p0 context.Context, p1 string, p2 []string) (*wallet.KeyBind, error) {
	return s.Internal.RemoveMethodFromKeyBind(p0, p1, p2)
}
func (s *IStrategyStruct) RemoveMethodTemplate(p0 context.Context, p1 string) error {
	return s.Internal.RemoveMethodTemplate(p0, p1)
}
func (s *IStrategyStruct) RemoveMsgTypeFromKeyBind(p0 context.Context, p1 string, p2 []int) (*wallet.KeyBind, error) {
	return s.Internal.RemoveMsgTypeFromKeyBind(p0, p1, p2)
}
func (s *IStrategyStruct) RemoveMsgTypeTemplate(p0 context.Context, p1 string) error {
	return s.Internal.RemoveMsgTypeTemplate(p0, p1)
}
func (s *IStrategyStruct) RemoveStToken(p0 context.Context, p1 string) error {
	return s.Internal.RemoveStToken(p0, p1)
}

type ILocalStrategyStruct struct {
	IStrategyVerifyStruct
	IStrategyStruct
}

type IWalletStruct struct {
	Internal struct {
		WalletDelete func(ctx context.Context, addr address.Address) error                                                           `perm:"admin"`
		WalletExport func(ctx context.Context, addr address.Address) (*types.KeyInfo, error)                                         `perm:"admin"`
		WalletHas    func(ctx context.Context, address address.Address) (bool, error)                                                `perm:"read"`
		WalletImport func(ctx context.Context, ki *types.KeyInfo) (address.Address, error)                                           `perm:"admin"`
		WalletList   func(ctx context.Context) ([]address.Address, error)                                                            `perm:"read"`
		WalletNew    func(ctx context.Context, kt types.KeyType) (address.Address, error)                                            `perm:"admin"`
		WalletSign   func(ctx context.Context, signer address.Address, toSign []byte, meta types.MsgMeta) (*crypto.Signature, error) `perm:"sign"`
	}
}

func (s *IWalletStruct) WalletDelete(p0 context.Context, p1 address.Address) error {
	return s.Internal.WalletDelete(p0, p1)
}
func (s *IWalletStruct) WalletExport(p0 context.Context, p1 address.Address) (*types.KeyInfo, error) {
	return s.Internal.WalletExport(p0, p1)
}
func (s *IWalletStruct) WalletHas(p0 context.Context, p1 address.Address) (bool, error) {
	return s.Internal.WalletHas(p0, p1)
}
func (s *IWalletStruct) WalletImport(p0 context.Context, p1 *types.KeyInfo) (address.Address, error) {
	return s.Internal.WalletImport(p0, p1)
}
func (s *IWalletStruct) WalletList(p0 context.Context) ([]address.Address, error) {
	return s.Internal.WalletList(p0)
}
func (s *IWalletStruct) WalletNew(p0 context.Context, p1 types.KeyType) (address.Address, error) {
	return s.Internal.WalletNew(p0, p1)
}
func (s *IWalletStruct) WalletSign(p0 context.Context, p1 address.Address, p2 []byte, p3 types.MsgMeta) (*crypto.Signature, error) {
	return s.Internal.WalletSign(p0, p1, p2, p3)
}

type IWalletLockStruct struct {
	Internal struct {
		Lock           func(ctx context.Context, password string) error `perm:"admin"`
		LockState      func(ctx context.Context) bool                   `perm:"admin"`
		SetPassword    func(ctx context.Context, password string) error `perm:"admin"`
		Unlock         func(ctx context.Context, password string) error `perm:"admin"`
		VerifyPassword func(ctx context.Context, password string) error `perm:"admin"`
	}
}

func (s *IWalletLockStruct) Lock(p0 context.Context, p1 string) error { return s.Internal.Lock(p0, p1) }
func (s *IWalletLockStruct) LockState(p0 context.Context) bool        { return s.Internal.LockState(p0) }
func (s *IWalletLockStruct) SetPassword(p0 context.Context, p1 string) error {
	return s.Internal.SetPassword(p0, p1)
}
func (s *IWalletLockStruct) Unlock(p0 context.Context, p1 string) error {
	return s.Internal.Unlock(p0, p1)
}
func (s *IWalletLockStruct) VerifyPassword(p0 context.Context, p1 string) error {
	return s.Internal.VerifyPassword(p0, p1)
}

type ILocalWalletStruct struct {
	IWalletStruct
	IWalletLockStruct
}

type ICommonStruct struct {
	Internal struct {
		AuthNew     func(ctx context.Context, perms []auth.Permission) ([]byte, error) `perm:"admin"`
		AuthVerify  func(ctx context.Context, token string) ([]auth.Permission, error) `perm:"read"`
		LogList     func(context.Context) ([]string, error)                            `perm:"read"`
		LogSetLevel func(context.Context, string, string) error                        `perm:"write"`
		Version     func(context.Context) (types.Version, error)                       `perm:"read"`
	}
}

func (s *ICommonStruct) AuthNew(p0 context.Context, p1 []auth.Permission) ([]byte, error) {
	return s.Internal.AuthNew(p0, p1)
}
func (s *ICommonStruct) AuthVerify(p0 context.Context, p1 string) ([]auth.Permission, error) {
	return s.Internal.AuthVerify(p0, p1)
}
func (s *ICommonStruct) LogList(p0 context.Context) ([]string, error) { return s.Internal.LogList(p0) }
func (s *ICommonStruct) LogSetLevel(p0 context.Context, p1 string, p2 string) error {
	return s.Internal.LogSetLevel(p0, p1, p2)
}
func (s *ICommonStruct) Version(p0 context.Context) (types.Version, error) {
	return s.Internal.Version(p0)
}

type IWalletEventStruct struct {
	Internal struct {
		AddNewAddress     func(ctx context.Context, newAddrs []address.Address) error `perm:"admin"`
		AddSupportAccount func(ctx context.Context, supportAccount string) error      `perm:"admin"`
	}
}

func (s *IWalletEventStruct) AddNewAddress(p0 context.Context, p1 []address.Address) error {
	return s.Internal.AddNewAddress(p0, p1)
}
func (s *IWalletEventStruct) AddSupportAccount(p0 context.Context, p1 string) error {
	return s.Internal.AddSupportAccount(p0, p1)
}

type IFullAPIStruct struct {
	ILocalStrategyStruct
	ILocalWalletStruct
	ICommonStruct
	IWalletEventStruct
}
