// // x/btc/client/cli/tx.go
package cli

// import (
// 	"github.com/cosmos/cosmos-sdk/client"
// 	"github.com/cosmos/cosmos-sdk/client/flags"
// 	"github.com/cosmos/cosmos-sdk/client/tx"
// 	"github.com/spf13/cobra"
// )

// func GetCmdConfirmBTCTx() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "confirm-btc-tx [tx-hash] [block-height]",
// 		Short: "Confirm a Bitcoin transaction",
// 		Args:  cobra.ExactArgs(2),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx := client.GetClientContextFromCmd(cmd)

// 			msg := types.Msg

// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }
