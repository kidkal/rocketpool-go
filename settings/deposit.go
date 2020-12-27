package settings

import (
    "fmt"
    "math/big"
    "sync"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"

    "github.com/rocket-pool/rocketpool-go/rocketpool"
)


// Deposits currently enabled
func GetDepositEnabled(rp *rocketpool.RocketPool, opts *bind.CallOpts) (bool, error) {
    rocketDepositSettings, err := getRocketDepositSettings(rp)
    if err != nil {
        return false, err
    }
    depositsEnabled := new(bool)
    if err := rocketDepositSettings.Call(opts, depositsEnabled, "getDepositEnabled"); err != nil {
        return false, fmt.Errorf("Could not get deposit enabled status: %w", err)
    }
    return *depositsEnabled, nil
}


// Deposit assignments currently enabled
func GetAssignDepositsEnabled(rp *rocketpool.RocketPool, opts *bind.CallOpts) (bool, error) {
    rocketDepositSettings, err := getRocketDepositSettings(rp)
    if err != nil {
        return false, err
    }
    assignDepositsEnabled := new(bool)
    if err := rocketDepositSettings.Call(opts, assignDepositsEnabled, "getAssignDepositsEnabled"); err != nil {
        return false, fmt.Errorf("Could not get deposit assignments enabled status: %w", err)
    }
    return *assignDepositsEnabled, nil
}


// Minimum deposit size
func GetMinimumDeposit(rp *rocketpool.RocketPool, opts *bind.CallOpts) (*big.Int, error) {
    rocketDepositSettings, err := getRocketDepositSettings(rp)
    if err != nil {
        return nil, err
    }
    minimumDeposit := new(*big.Int)
    if err := rocketDepositSettings.Call(opts, minimumDeposit, "getMinimumDeposit"); err != nil {
        return nil, fmt.Errorf("Could not get minimum deposit: %w", err)
    }
    return *minimumDeposit, nil
}


// The maximum size of the deposit pool
func GetMaximumDepositPoolSize(rp *rocketpool.RocketPool, opts *bind.CallOpts) (*big.Int, error) {
    rocketDepositSettings, err := getRocketDepositSettings(rp)
    if err != nil {
        return nil, err
    }
    maximumDepositPoolSize := new(*big.Int)
    if err := rocketDepositSettings.Call(opts, maximumDepositPoolSize, "getMaximumDepositPoolSize"); err != nil {
        return nil, fmt.Errorf("Could not get maximum deposit pool size: %w", err)
    }
    return *maximumDepositPoolSize, nil
}


// Maximum deposit assignments per transaction
func GetMaximumDepositAssignments(rp *rocketpool.RocketPool, opts *bind.CallOpts) (uint64, error) {
    rocketDepositSettings, err := getRocketDepositSettings(rp)
    if err != nil {
        return 0, err
    }
    maximumDepositAssignments := new(*big.Int)
    if err := rocketDepositSettings.Call(opts, maximumDepositAssignments, "getMaximumDepositAssignments"); err != nil {
        return 0, fmt.Errorf("Could not get maximum deposit assignments: %w", err)
    }
    return (*maximumDepositAssignments).Uint64(), nil
}


// Get contracts
var rocketDepositSettingsLock sync.Mutex
func getRocketDepositSettings(rp *rocketpool.RocketPool) (*bind.BoundContract, error) {
    rocketDepositSettingsLock.Lock()
    defer rocketDepositSettingsLock.Unlock()
    return rp.GetContract("rocketDepositSettings")
}

