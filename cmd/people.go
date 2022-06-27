package cmd

import (
	"log"
	"zhengzaitv-go/wap"

	"github.com/modood/table"
	"github.com/spf13/cobra"
)

var zhengzai *wap.Zhengzaitv

type PeopleItem struct {
	EntersID  string `json:"entersId"`
	Name      string `json:"name"`
	IDCard    string `json:"idCard"`
	IsDefault bool   `json:"isDefault"`
	CreatedAt string `json:"createdAt"`
}

var peopleCmd = &cobra.Command{
	Use:     "people",
	Short:   "查询观演人列表",
	Long:    `查询观演人列表`,
	PreRunE: PreRunZhengzaitv, // 注意这里需要是 return error 的类型，不然直接 return 无法阻止程序运行 Run func
	Run: func(cmd *cobra.Command, args []string) {
		peoples, err := zhengzai.GetEntersPeople()
		if err != nil {
			log.Printf("get peoples err: %v", err)
			return
		}

		res := make([]PeopleItem, 0)
		for _, peopleItem := range peoples {
			res = append(res, PeopleItem{
				EntersID:  peopleItem.EntersID,
				Name:      peopleItem.Name,
				IDCard:    peopleItem.IDCard,
				IsDefault: peopleItem.IsDefault,
				CreatedAt: peopleItem.CreatedAt,
			})
		}

		table.Output(res)
	},
}

func init() {
	rootCmd.AddCommand(peopleCmd)
}
