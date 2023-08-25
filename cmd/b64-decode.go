package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  b64Cmd.AddCommand(b64_decodeCmd)


  
}

var b64_decodeCmd = &cobra.Command{
  Use:   "decode",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
  },
}