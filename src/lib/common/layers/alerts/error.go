// Package alerts: alter defs
package alerts

import (
	"time"

	"leetui/src/lib/viewmodel"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type ErrorModalModel struct {
	viewmodel.ViewModel
	Show           bool
	CurrentMessage string
}

func MakeErrorModalModel() ErrorModalModel {
	return ErrorModalModel{
		ViewModel: viewmodel.ViewModel{
			Dims: viewmodel.ViewModelDims{
				Width:  0,
				Height: 3,
			},
		},
		Show:           false,
		CurrentMessage: "no error",
	}
}

func (m ErrorModalModel) Init() tea.Cmd {
	return tea.Batch()
}

func (m ErrorModalModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case AlertErrorMsg:
		m.Show = true
		m.CurrentMessage = msg.Message
		m.Dims.Width = len(msg.Message) + 4
		return m, DurationTimeoutCmd(msg.Duration)
	case HideAlertErrorMsg:
		m.Show = false
		return m, nil
	}

	return m, tea.Batch()
}

func (m ErrorModalModel) View() tea.View {
	modal := ""
	if m.Show {
		modal = m.NewErrorModal()
	}

	return tea.NewView(modal)
}

func (m ErrorModalModel) NewErrorModal() string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("1")).
		Height(m.Dims.Height).
		Width(m.Dims.Width). // Padding around sides
		Padding(0, 1).
		Align(lipgloss.Center).
		Render(m.CurrentMessage)
}

type AlertErrorMsg struct {
	Message  string
	Duration time.Duration
}

type HideAlertErrorMsg struct{}

func AlertErrorCmd(msg AlertErrorMsg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

func DurationTimeoutCmd(wait time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(wait)
		return HideAlertErrorMsg{}
	}
}
