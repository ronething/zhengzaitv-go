package cmd

import (
	"log"

	"github.com/modood/table"
	"github.com/spf13/cobra"
)

type AddressItem struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Province  string `json:"province"`
	City      string `json:"city"`
	County    string `json:"county"`
	Address   string `json:"address"`
	IsDefault bool   `json:"isDefault"`
	CreatedAt string `json:"createdAt"`
}

// addressCmd represents the address command
var addressCmd = &cobra.Command{
	Use:     "address",
	Short:   "获取地址列表",
	Long:    `获取地址列表`,
	PreRunE: PreRunZhengzaitv,
	Run: func(cmd *cobra.Command, args []string) {
		addrList, err := zhengzai.GetAddrList()
		if err != nil {
			log.Print(err)
			return
		}

		res := make([]AddressItem, 0)
		for _, addressItem := range addrList {
			res = append(res, AddressItem{
				Name:      addressItem.Name,
				Phone:     addressItem.Phone,
				Province:  addressItem.Province,
				City:      addressItem.City,
				County:    addressItem.County,
				Address:   addressItem.Address,
				IsDefault: addressItem.IsDefault,
				CreatedAt: addressItem.CreatedAt,
			})

		}

		table.Output(res)
	},
}

func init() {
	rootCmd.AddCommand(addressCmd)
}
