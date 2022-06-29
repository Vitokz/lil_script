package x

import (
	"encoding/hex"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func Bech32ToHex(bechAddress string) (string, error) {
	acc, err := sdkTypes.AccAddressFromBech32(bechAddress)
	if err != nil {
		return "", err
	}

	str := hex.EncodeToString(acc.Bytes())
	return fmt.Sprintf("0x%s", str), nil
}
