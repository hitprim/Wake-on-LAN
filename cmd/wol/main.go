package main

import (
	"fmt"
	"os"
	"wol/internal/wol"

	"github.com/spf13/cobra"
)

func main() {
	var broadcast string
	var port int

	root := &cobra.Command{
		Use:   "wol <MAC> [optional]",
		Short: "Magic packet to PC",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			mac := args[0]
			if err := wol.Send(mac, broadcast, port); err != nil {
				fmt.Println("Ошибка:", err)
				os.Exit(1)
			}
			fmt.Println("✅ Магический пакет отправлен!")
		},
	}
	root.Flags().IntVarP(&port, "port", "p", 9, "port to listen on")
	root.Flags().StringVarP(&broadcast, "broadcast", "b", "255.255.255.255", "broadcast packet")

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
