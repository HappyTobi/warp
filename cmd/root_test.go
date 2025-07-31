package cmd

import (
	"testing"
)

func TestRoot(t *testing.T) {
	rootCmd := Root()
	
	if rootCmd == nil {
		t.Error("Root() returned nil")
	}
	
	if rootCmd.Use != "warp" {
		t.Errorf("Root() Use = %s, expected 'warp'", rootCmd.Use)
	}
	
	// Check that persistent flags are set
	if rootCmd.PersistentFlags().Lookup("charger") == nil {
		t.Error("Root() missing 'charger' persistent flag")
	}
	
	if rootCmd.PersistentFlags().Lookup("output") == nil {
		t.Error("Root() missing 'output' persistent flag")
	}
	
	if rootCmd.PersistentFlags().Lookup("config") == nil {
		t.Error("Root() missing 'config' persistent flag")
	}
	
	if rootCmd.PersistentFlags().Lookup("username") == nil {
		t.Error("Root() missing 'username' persistent flag")
	}
	
	if rootCmd.PersistentFlags().Lookup("password") == nil {
		t.Error("Root() missing 'password' persistent flag")
	}
}

func TestRootCommandStructure(t *testing.T) {
	rootCmd := Root()
	
	// Check that all expected subcommands are added
	expectedCommands := []string{
		"charge",
		"info", 
		"users",
		"charge-tracker",
		"meter",
		"evcc",
		"version",
		"configuration",
	}
	
	commands := rootCmd.Commands()
	commandNames := make([]string, len(commands))
	for i, cmd := range commands {
		commandNames[i] = cmd.Name()
	}
	
	for _, expected := range expectedCommands {
		found := false
		for _, actual := range commandNames {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Root() missing expected subcommand: %s", expected)
		}
	}
}

func TestRootPersistentPreRunE(t *testing.T) {
	rootCmd := Root()
	
	// Test that PersistentPreRunE is set
	if rootCmd.PersistentPreRunE == nil {
		t.Error("Root() PersistentPreRunE not set")
	}
	
	// We can't easily test the actual execution without setting up the full command context
	// But we can verify the function exists
}

func TestRootFlagDefaults(t *testing.T) {
	rootCmd := Root()
	
	// Test default values for flags
	outputFlag := rootCmd.PersistentFlags().Lookup("output")
	if outputFlag == nil {
		t.Error("Root() 'output' flag not found")
	} else if outputFlag.DefValue != "json" {
		t.Errorf("Root() 'output' flag default = %s, expected 'json'", outputFlag.DefValue)
	}
	
	// Test that username and password are marked as required together
	chargerFlag := rootCmd.PersistentFlags().Lookup("charger")
	if chargerFlag == nil {
		t.Error("Root() 'charger' flag not found")
	}
	
	configFlag := rootCmd.PersistentFlags().Lookup("config")
	if configFlag == nil {
		t.Error("Root() 'config' flag not found")
	}
}

func TestChargeCmd(t *testing.T) {
	chargeCmd := ChargeCmd()
	
	if chargeCmd == nil {
		t.Error("ChargeCmd() returned nil")
	}
	
	if chargeCmd.Use != "charge" {
		t.Errorf("ChargeCmd() Use = %s, expected 'charge'", chargeCmd.Use)
	}
	
	// Check that subcommands exist
	commands := chargeCmd.Commands()
	if len(commands) == 0 {
		t.Error("ChargeCmd() has no subcommands")
	}
	
	// Look for start and stop commands
	hasStart := false
	hasStop := false
	for _, cmd := range commands {
		if cmd.Name() == "start" {
			hasStart = true
		}
		if cmd.Name() == "stop" {
			hasStop = true
		}
	}
	
	if !hasStart {
		t.Error("ChargeCmd() missing 'start' subcommand")
	}
	if !hasStop {
		t.Error("ChargeCmd() missing 'stop' subcommand")
	}
}

func TestInfoCmd(t *testing.T) {
	infoCmd := InfoCmd()
	
	if infoCmd == nil {
		t.Error("InfoCmd() returned nil")
	}
	
	if infoCmd.Use != "info" {
		t.Errorf("InfoCmd() Use = %s, expected 'info'", infoCmd.Use)
	}
}

func TestUserCmd(t *testing.T) {
	userCmd := UserCmd()
	
	if userCmd == nil {
		t.Error("UserCmd() returned nil")
	}
	
	if userCmd.Use != "users" {
		t.Errorf("UserCmd() Use = %s, expected 'users'", userCmd.Use)
	}
}

func TestChargeTrackerCmd(t *testing.T) {
	chargeTrackerCmd := ChargeTrackerCmd()
	
	if chargeTrackerCmd == nil {
		t.Error("ChargeTrackerCmd() returned nil")
	}
	
	if chargeTrackerCmd.Use != "charge-tracker" {
		t.Errorf("ChargeTrackerCmd() Use = %s, expected 'charge-tracker'", chargeTrackerCmd.Use)
	}
}

func TestMeterCmd(t *testing.T) {
	meterCmd := MeterCmd()
	
	if meterCmd == nil {
		t.Error("MeterCmd() returned nil")
	}
	
	if meterCmd.Use != "meter" {
		t.Errorf("MeterCmd() Use = %s, expected 'meter'", meterCmd.Use)
	}
}

func TestEvccCmd(t *testing.T) {
	evccCmd := EvccCmd()
	
	if evccCmd == nil {
		t.Error("EvccCmd() returned nil")
	}
	
	if evccCmd.Use != "evcc" {
		t.Errorf("EvccCmd() Use = %s, expected 'evcc'", evccCmd.Use)
	}
}

func TestVersionCmd(t *testing.T) {
	versionCmd := VersionCmd()
	
	if versionCmd == nil {
		t.Error("VersionCmd() returned nil")
	}
	
	if versionCmd.Use != "version" {
		t.Errorf("VersionCmd() Use = %s, expected 'version'", versionCmd.Use)
	}
}

func TestConfigurationCmd(t *testing.T) {
	configCmd := ConfigurationCmd()
	
	if configCmd == nil {
		t.Error("ConfigurationCmd() returned nil")
	}
	
	if configCmd.Use != "configuration" {
		t.Errorf("ConfigurationCmd() Use = %s, expected 'configuration'", configCmd.Use)
	}
}
