package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mynewgoproject/internal/core/domain"
	"mynewgoproject/internal/core/port/driving"
	"time"

	"github.com/spf13/cobra"
)

type CliServer struct {
	s       driving.Server
	rootCmd *cobra.Command
}

func NewCliServer(s driving.Server) *CliServer {
	rootCmd := &cobra.Command{Use: "goome"}
	httpServer := &CliServer{rootCmd: rootCmd, s: s}

	lightCmd := &cobra.Command{Use: "light"}
	rootCmd.AddCommand(lightCmd)

	lightCmd.AddCommand(httpServer.lightOn())
	lightCmd.AddCommand(httpServer.lightOff())
	lightCmd.AddCommand(httpServer.lightChangeColor())
	lightCmd.AddCommand(httpServer.lightChangeWhite())

	return httpServer
}

func (s *CliServer) Run() error {
	return s.rootCmd.Execute()
}

func (s *CliServer) lightOn() *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:   "on",
		Short: "Switch on the light",
		Long:  `Switch on the light.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			err := s.s.LightOn(ctxTimeout, name)
			s.handleError("switch on the ligh", err)
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the light")
	return cmd
}

func (s *CliServer) lightOff() *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:   "off",
		Short: "Switch off the light",
		Long:  `Switch off the light.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			err := s.s.LightOff(ctxTimeout, name)
			s.handleError("switch off the ligh", err)
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the light")
	return cmd
}

func (s *CliServer) lightChangeColor() *cobra.Command {
	var name string
	var color domain.Color
	cmd := &cobra.Command{
		Use:   "color",
		Short: "Change the color of the light",
		Long:  `Change the color of the light.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			err := s.s.LightChangeColor(ctxTimeout, name, &color)
			s.handleError("change color of the light", err)
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the light")
	cmd.Flags().Int32VarP(&color.Red, "red", "r", 100, "Red component")
	cmd.Flags().Int32VarP(&color.Green, "green", "g", 100, "Green component")
	cmd.Flags().Int32VarP(&color.Blue, "blue", "b", 100, "Blue component")
	cmd.Flags().Int32VarP(&color.White, "white", "w", 100, "White component")
	cmd.Flags().Int32Var(&color.Gain, "gain", 100, "Gain component")
	return cmd
}

func (s *CliServer) lightChangeWhite() *cobra.Command {
	var name string
	var white domain.White
	cmd := &cobra.Command{
		Use:   "white",
		Short: "Change the white of the light",
		Long:  `Change the white of the light.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			err := s.s.LightChangeWhite(ctxTimeout, name, &white)
			s.handleError("change white of the light", err)
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the light")
	cmd.Flags().Int32VarP(&white.Brightness, "brightness", "b", 50, "Brightness component")
	cmd.Flags().Int32VarP(&white.Temp, "temp", "t", 4750, "Temp component")
	return cmd
}

func (s *CliServer) handleError(action string, err error) {
	switch {
	case errors.Is(err, domain.ErrorEventNotValid):
		log.Println(fmt.Sprintf("❌ failed to %s: %s", action, err.Error()))
	case errors.Is(err, domain.ErrorDeviceNotFound):
		log.Println(fmt.Sprintf("🕵️‍♂️ failed to %s: %s", action, err.Error()))
	case err != nil:
		log.Println(fmt.Sprintf("🛠️ failed to %s: %s", action, err.Error()))
	default:
		log.Println(fmt.Sprintf("✅ successfully %s", action))
	}
}
