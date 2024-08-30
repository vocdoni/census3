package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/vocdoni/census3/api"
	"github.com/vocdoni/census3/apiclient"
)

var client *apiclient.HTTPclient

func main() {
	var apiUrl string
	var authToken string

	// Root command
	rootCmd := &cobra.Command{
		Use:   "census3-cli",
		Short: "Census3 API Interactive Client",
		Long:  "An interactive client to interact with the Census3 API using Go.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Setup client before running commands
			u, err := url.Parse(apiUrl)
			if err != nil {
				log.Fatalf("Invalid API URL: %v", err)
			}
			client, err = apiclient.NewHTTPclient(u, nil)
			if err != nil {
				log.Fatalf("Failed to create API client: %v", err)
			}

			if authToken != "" {
				t := uuid.MustParse(authToken)
				client.SetAuthToken(&t)
			}
		},
	}

	rootCmd.PersistentFlags().StringVar(&apiUrl, "api-url", "https://census3.vocdoni.net/api", "Census3 API URL")
	rootCmd.PersistentFlags().StringVar(&authToken, "auth-token", "", "Bearer authentication token")

	// Info command
	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "Get API information",
		Run: func(cmd *cobra.Command, args []string) {
			info, err := client.Info()
			if err != nil {
				color.Red("Error fetching API info: %v", err)
				return
			}
			color.Green("API Information:")
			fmt.Printf("  Supported Chains: %v\n", info.SupportedChains)
		},
	}
	rootCmd.AddCommand(infoCmd)

	// Tokens command
	tokensCmd := &cobra.Command{
		Use:   "tokens",
		Short: "List all tokens",
		Run: func(cmd *cobra.Command, args []string) {
			tokens, err := client.Tokens(100, "", "")
			if err != nil {
				color.Red("Error fetching tokens: %v", err)
				return
			}
			for _, token := range tokens {
				color.Cyan("Token ID: %s\n", token.ID)
				fmt.Printf("  Type: %s\n", token.Type)
				fmt.Printf("  Chain ID: %d\n", token.ChainID)
				fmt.Printf("  Tags: %v\n", token.Tags)
				fmt.Println()
			}
		},
	}
	rootCmd.AddCommand(tokensCmd)

	// Token command
	tokenCmd := &cobra.Command{
		Use:   "token [tokenID] [chainID]",
		Short: "Get information about a specific token",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			tokenID := args[0]
			chainID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				color.Red("Invalid chain ID: %v", err)
				return
			}

			token, err := client.Token(tokenID, chainID, "")
			if err != nil {
				color.Red("Error fetching token: %v", err)
				return
			}
			color.Green("Token Details:")
			fmt.Printf("  ID: %s\n", token.ID)
			fmt.Printf("  Type: %s\n", token.Type)
			fmt.Printf("  Chain ID: %d\n", token.ChainID)
			fmt.Printf("  Start Block: %d\n", token.StartBlock)
			fmt.Printf("  Tags: %v\n", token.Tags)
		},
	}
	rootCmd.AddCommand(tokenCmd)

	// Add Token command
	addTokenCmd := &cobra.Command{
		Use:   "add_token",
		Short: "Add a new token",
		Run: func(cmd *cobra.Command, args []string) {
			var tokenID, tokenType, chainID, startBlock, tags string

			fmt.Print("Enter Token ID: ")
			if _, err := fmt.Scanln(&tokenID); err != nil {
				color.Red("Invalid token ID: %v", err)
				return
			}
			fmt.Print("Enter Token Type (e.g., erc20): ")
			if _, err := fmt.Scanln(&tokenType); err != nil {
				color.Red("Invalid token type: %v", err)
				return
			}
			fmt.Print("Enter Chain ID: ")
			if _, err := fmt.Scanln(&chainID); err != nil {
				color.Red("Invalid chain ID: %v", err)
				return
			}
			fmt.Print("Enter Start Block: ")
			fmt.Scanln(&startBlock)
			fmt.Print("Enter Tags (comma-separated): ")
			fmt.Scanln(&tags)

			chainIDUint, err := strconv.ParseUint(chainID, 10, 64)
			if err != nil {
				color.Red("Invalid chain ID: %v", err)
				return
			}
			startBlockUint, err := strconv.ParseUint(startBlock, 10, 64)
			if err != nil {
				startBlockUint = 0
			}

			token := &api.Token{
				ID:         tokenID,
				Type:       tokenType,
				ChainID:    chainIDUint,
				StartBlock: startBlockUint,
				Tags:       tags,
			}

			if err := client.CreateToken(token); err != nil {
				color.Red("Error creating token: %v", err)
				return
			}
			color.Green("Token created successfully.")
		},
	}
	rootCmd.AddCommand(addTokenCmd)

	// Token command
	holdersCmd := &cobra.Command{
		Use:   "holders <strategyID>",
		Short: "Get the list of token holders for an strategy",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			strategyID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				color.Red("Invalid strategy ID: %v", err)
				return
			}
			holders, err := client.AllHoldersByStrategy(strategyID, true)
			if err != nil {
				color.Red("Error fetching holders: %v", err)
				return
			}
			color.Green("Holders:")
			for addr, amount := range holders {
				fmt.Printf("  %s %s\n", addr.String(), amount.String())
			}
		},
	}
	rootCmd.AddCommand(holdersCmd)

	// Strategies command
	strategiesCmd := &cobra.Command{
		Use:   "strategies [tokenID]",
		Short: "List all strategies (filter by token if provided)",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			strategies, err := client.Strategies(-1, "", "")
			if err != nil {
				color.Red("Error fetching strategies: %v", err)
				return
			}

			for _, strategy := range strategies {
				if args[0] != "" && !strings.Contains(strategy.Predicate, args[0]) {
					continue
				}
				color.Cyan("Strategy ID: %d\n", strategy.ID)
				fmt.Printf("  Alias: %s\n", strategy.Alias)
				fmt.Printf("  Predicate: %s\n", strategy.Predicate)
				tokens := ""
				for _, token := range strategy.Tokens {
					tokens += token.ID + ", "
				}
				fmt.Printf("  URI: %s\n", strategy.URI)
				fmt.Println()
			}
		},
	}
	rootCmd.AddCommand(strategiesCmd)

	// Run the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(1)
	}
}
