package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var outputFile string

var compressCmd = &cobra.Command{
	Use:   "compress [input-file]",
	Short: "compresses a file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inputFile := args[0]

		start := time.Now()

		data, err := os.ReadFile(inputFile)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", inputFile, err)
		}

		err = os.WriteFile(outputFile, data, 0644)
		if err != nil {
			return fmt.Errorf("failed to write %s: %w", outputFile, err)
		}

		elapsed := time.Since(start)

		fmt.Printf("Original size:    %d bytes\n", len(data))
		fmt.Printf("Compressed size:  %d bytes\n", len(data))
		fmt.Printf("Ratio:            100.00%%\n")
		fmt.Printf("Time elapsed:     %v\n", elapsed)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
	compressCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file (required)")
	compressCmd.MarkFlagRequired("output")
}
