/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	b64 "encoding/base64"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jim-at-jibba/devtools/tui"
	"github.com/spf13/cobra"
)

// encodeBase64Cmd represents the encodeBase64 command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode base64 string",
	Long:  "Decode base64 string",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetBool("url")
		p := tea.NewProgram(initialDecodeModel(url))

		if err := p.Start(); err != nil {
			fmt.Println("WHat", err)
			os.Exit(1)
		}

	},
}

func init() {
	base64Cmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	decodeCmd.PersistentFlags().BoolP("url", "u", false, "URL-compatible base64 format")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeBase64Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type decodeModel struct {
	rawString textinput.Model
	decoded   string
	err       error
	url       bool
}

type decodeStr struct {
	decoded string
}

func Decode(raw string, url bool) decodeStr {
	if url {
		sDec, _ := b64.URLEncoding.DecodeString(raw)
		return decodeStr{decoded: string(sDec)}
	} else {
		sDec, _ := b64.StdEncoding.DecodeString(raw)
		return decodeStr{decoded: string(sDec)}
	}
}

func (m decodeModel) decodeMsg() tea.Msg {
	decoded := Decode(m.rawString.Value(), m.url)
	return decoded
}

func initialDecodeModel(url bool) decodeModel {
	ti := textinput.New()
	ti.Placeholder = "String to decode"
	ti.Focus()

	return decodeModel{
		rawString: ti,
		decoded:   "",
		err:       nil,
		url:       url,
	}
}

func (m decodeModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m decodeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			return m, m.decodeMsg
		}

	case decodeStr:
		decoded := msg
		m.decoded = decoded.decoded
		return m, tea.Quit

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil

	}

	m.rawString, cmd = m.rawString.Update(msg)
	return m, cmd
}

func (m decodeModel) View() string {
	if len(m.decoded) > 0 {
		return tui.ContainerStyle.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				tui.LabelStyle.Render("Decoded string:"),
				tui.Spacer.Render(""),
				tui.ValueStyle.Render(m.decoded),
			),
		)
	} else {
		return lipgloss.JoinVertical(lipgloss.Left,
			tui.LabelStyle.Render("Enter the string you want to decode."),
			m.rawString.View(),
			tui.ValueStyle.Render("(q to quit)"),
		)
	}
}
