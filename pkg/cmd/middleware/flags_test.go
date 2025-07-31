package middleware

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestLoadWarpRequest_MissingFlags(t *testing.T) {
	// Create a test command without required flags
	cmd := &cobra.Command{
		Use: "test",
	}
	
	// This should return an error since required flags are missing
	_, err := LoadWarpRequest(cmd)
	if err == nil {
		t.Error("LoadWarpRequest() should return error when required flags are missing")
	}
}

func TestLoadWarpRequest_WithChargerFlag(t *testing.T) {
	// Create a test root command with flags
	rootCmd := &cobra.Command{
		Use: "root",
	}
	
	// Add the flags that the function expects on the root command
	rootCmd.Flags().StringP("charger", "c", "", "Charger URL")
	rootCmd.Flags().StringP("output", "o", "json", "Output format")
	rootCmd.Flags().StringP("username", "u", "", "Username")
	rootCmd.Flags().StringP("password", "p", "", "Password")
	
	// Create a child command
	cmd := &cobra.Command{
		Use: "test",
	}
	rootCmd.AddCommand(cmd)
	
	// Set flag values on the root command
	err := rootCmd.Flags().Set("charger", "http://test.local")
	if err != nil {
		t.Fatalf("Failed to set charger flag: %v", err)
	}
	
	err = rootCmd.Flags().Set("output", "yaml")
	if err != nil {
		t.Fatalf("Failed to set output flag: %v", err)
	}
	
	request, err := LoadWarpRequest(cmd)
	if err != nil {
		t.Errorf("LoadWarpRequest() error = %v", err)
	}
	
	if request.Warp != "http://test.local" {
		t.Errorf("LoadWarpRequest() Warp = %s, expected 'http://test.local'", request.Warp)
	}
}

func TestLoadGlobalParams(t *testing.T) {
	// Create a test root command with flags
	rootCmd := &cobra.Command{
		Use: "root",
	}
	
	// Add the flags
	rootCmd.Flags().StringP("charger", "c", "", "Charger URL")
	rootCmd.Flags().StringP("output", "o", "json", "Output format")
	rootCmd.Flags().StringP("username", "u", "", "Username")
	rootCmd.Flags().StringP("password", "p", "", "Password")
	
	// Create a child command
	cmd := &cobra.Command{
		Use: "test",
	}
	rootCmd.AddCommand(cmd)
	
	// Set values on the root command
	rootCmd.Flags().Set("charger", "http://example.com")
	rootCmd.Flags().Set("username", "testuser")
	rootCmd.Flags().Set("password", "testpass")
	rootCmd.Flags().Set("output", "yaml")
	
	// Test callback function
	called := false
	testCallback := func(charger, username, password, output string) {
		called = true
		if charger != "http://example.com" {
			t.Errorf("LoadGlobalParams() charger = %s, expected 'http://example.com'", charger)
		}
		if username != "testuser" {
			t.Errorf("LoadGlobalParams() username = %s, expected 'testuser'", username)
		}
		if password != "testpass" {
			t.Errorf("LoadGlobalParams() password = %s, expected 'testpass'", password)
		}
		if output != "yaml" {
			t.Errorf("LoadGlobalParams() output = %s, expected 'yaml'", output)
		}
	}
	
	err := LoadGlobalParams(cmd, testCallback)
	if err != nil {
		t.Errorf("LoadGlobalParams() error = %v", err)
	}
	
	if !called {
		t.Error("LoadGlobalParams() callback function was not called")
	}
}

func TestLoadGlobalParams_WithViper(t *testing.T) {
	// Reset viper state
	viper.Reset()
	
	// Set some viper values
	viper.Set("charger.url", "http://viper.local")
	viper.Set("charger.username", "viperuser")
	
	// Create a test command with minimal flags
	cmd := &cobra.Command{
		Use: "test",
	}
	
	cmd.PersistentFlags().StringP("charger", "c", "", "Charger URL")
	cmd.PersistentFlags().StringP("output", "o", "json", "Output format")
	cmd.PersistentFlags().StringP("username", "u", "", "Username")
	cmd.PersistentFlags().StringP("password", "p", "", "Password")
	
	// Only set output flag, others should come from viper
	cmd.PersistentFlags().Set("output", "json")
	
	called := false
	testCallback := func(charger, username, password, output string) {
		called = true
		// Values should come from viper when flags are not set
		if charger != "http://viper.local" {
			t.Errorf("LoadGlobalParams() charger from viper = %s, expected 'http://viper.local'", charger)
		}
		if username != "viperuser" {
			t.Errorf("LoadGlobalParams() username from viper = %s, expected 'viperuser'", username)
		}
	}
	
	err := LoadGlobalParams(cmd, testCallback)
	if err != nil {
		t.Errorf("LoadGlobalParams() with viper error = %v", err)
	}
	
	if !called {
		t.Error("LoadGlobalParams() callback function was not called")
	}
	
	// Clean up
	viper.Reset()
}
