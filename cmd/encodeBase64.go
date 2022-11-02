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
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode base64 string",
	Long:  "Encode base64 string",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initialEncodeModel())

		if err := p.Start(); err != nil {
			fmt.Println("WHat", err)
			os.Exit(1)
		}

	},
}

func init() {
	base64Cmd.AddCommand(encodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeBase64Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeBase64Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type (
	errMsg error
)

type encodeModel struct {
	rawString textinput.Model
	encoded   string
	err       error
}

type encodeStr struct {
	encoded string
}

func Encode(raw string) encodeStr {
	sEnc := b64.StdEncoding.EncodeToString([]byte(raw))
	return encodeStr{encoded: sEnc}
}

func (m encodeModel) encodeMsg() tea.Msg {
	encoded := Encode(m.rawString.Value())
	return encoded
}

func initialEncodeModel() encodeModel {
	ti := textinput.New()
	ti.Placeholder = "String to encode"
	ti.Focus()

	return encodeModel{
		rawString: ti,
		encoded:   "",
		err:       nil,
	}
}

func (m encodeModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m encodeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			return m, m.encodeMsg
		}

	case encodeStr:
		encoded := msg
		m.encoded = encoded.encoded
		return m, tea.Quit

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil

	}

	m.rawString, cmd = m.rawString.Update(msg)
	return m, cmd
}

func (m encodeModel) View() string {
	if len(m.encoded) > 0 {
		return tui.ContainerStyle.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				tui.LabelStyle.Render("Encoded string:"),
				tui.Spacer.Render(""),
				tui.ValueStyle.Render(m.encoded),
			),
		)
	} else {
		return lipgloss.JoinVertical(lipgloss.Left,
			tui.LabelStyle.Render("Enter the string you want to encode."),
			m.rawString.View(),
			tui.ValueStyle.Render("(q to quit)"),
		)
	}
}
