package cmd

import (
	"log"

	"github.com/modood/table"
	"github.com/spf13/cobra"
)

type TicketItem struct {
	TicketId string  `json:"ticketId"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
}

var ticketsCmd = &cobra.Command{
	Use:     "tickets",
	Short:   "列出指定场次 ticketId 列表",
	Long:    `根据场次 Id(activityId) 列出指定场次 ticketId 列表`,
	PreRunE: PreRunZhengzaitv,
	Run: func(cmd *cobra.Command, args []string) {

		performanceId := getFlag(cmd, "activityId")
		if performanceId == "" { // 为空则 return
			log.Println("activityId 场次 id 为空，请检查")
			return
		}

		tickets, err := zhengzai.GetTicketIds(performanceId)
		if err != nil {
			log.Printf("获取票种 id 失败, err: %v", err)
			return
		}

		res := make([]TicketItem, 0)
		for _, ticketItem := range tickets {
			res = append(res, TicketItem{
				TicketId: ticketItem.TicketsID,
				Title:    ticketItem.Title,
				Price:    ticketItem.Price,
			})
		}

		table.Output(res)

	},
}

func init() {
	rootCmd.AddCommand(ticketsCmd)

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ticketsCmd.Flags().StringP("activityId", "a", "", "场次 id")

}
