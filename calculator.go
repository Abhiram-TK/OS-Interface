package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalculator() {

	var historyArr []string

	displayHistory := false

	output := ""
	input := widget.NewLabel(output)

	historyStr := ""
	history := widget.NewLabel(historyStr)

	BackBtn := widget.NewButton("Back", func() {

		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}

	})

	HistoryBtn := widget.NewButton("History", func() {

		if displayHistory {
			historyStr = ""
			myWindow.Resize(fyne.NewSize(500, 280))
		} else {
			for i := len(historyArr) - 1; i >= 0; i-- {
				historyStr = historyStr + historyArr[i]
				historyStr += "\n"
			}
		}

		history.SetText(historyStr)
		displayHistory = !displayHistory

	})

	ClearBtn := widget.NewButton("Clear", func() {

		output = ""
		input.SetText(output)

	})

	OpenParamBtn := widget.NewButton("(", func() {

		output = output + "("
		input.SetText(output)

	})

	CloseParamBtn := widget.NewButton(")", func() {

		output = output + ")"
		input.SetText(output)

	})

	SevenBtn := widget.NewButton("7", func() {

		output = output + "7"
		input.SetText(output)

	})

	EightBtn := widget.NewButton("8", func() {

		output = output + "8"
		input.SetText(output)

	})

	NineBtn := widget.NewButton("9", func() {

		output = output + "9"
		input.SetText(output)

	})

	DivideBtn := widget.NewButton("/", func() {

		output = output + "/"
		input.SetText(output)

	})

	FourBtn := widget.NewButton("4", func() {

		output = output + "4"
		input.SetText(output)

	})

	FiveBtn := widget.NewButton("5", func() {

		output = output + "5"
		input.SetText(output)

	})

	SixBtn := widget.NewButton("6", func() {

		output = output + "6"
		input.SetText(output)

	})

	MultiplyBtn := widget.NewButton("*", func() {

		output = output + "*"
		input.SetText(output)
	})

	OneBtn := widget.NewButton("1", func() {

		output = output + "1"
		input.SetText(output)

	})

	TwoBtn := widget.NewButton("2", func() {

		output = output + "2"
		input.SetText(output)

	})

	ThreeBtn := widget.NewButton("3", func() {

		output = output + "3"
		input.SetText(output)

	})

	MinusBtn := widget.NewButton("-", func() {

		output = output + "-"
		input.SetText(output)

	})

	PlusBtn := widget.NewButton("+", func() {

		output = output + "+"
		input.SetText(output)

	})

	ZeroBtn := widget.NewButton("0", func() {

		output = output + "0"
		input.SetText(output)

	})

	DotBtn := widget.NewButton(".", func() {

		output = output + "."
		input.SetText(output)

	})

	EqualBtn := widget.NewButton("=", func() {

		expression, err := govaluate.NewEvaluableExpression(output)

		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + "=" + ans
				historyArr = append(historyArr, strToAppend)
				output = ans
			} else {
				output = "error"
			}
		} else {
			output = "error"
		}
		input.SetText(output)
		fmt.Println(historyArr)

	})

	EqualBtn.Importance = widget.HighImportance

	calculatorContainer := container.NewVBox(
		container.NewVBox(
			input,
			history,
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(2,
					HistoryBtn,
					BackBtn,
				),
				container.NewGridWithColumns(4,
					ClearBtn,
					OpenParamBtn,
					CloseParamBtn,
					DivideBtn),
				container.NewGridWithColumns(4,
					SevenBtn,
					EightBtn,
					NineBtn,
					MultiplyBtn),
				container.NewGridWithColumns(4,
					FourBtn,
					FiveBtn,
					SixBtn,
					MinusBtn),
				container.NewGridWithColumns(4,
					OneBtn,
					TwoBtn,
					ThreeBtn,
					PlusBtn),
				container.NewGridWithColumns(2,
					container.NewGridWithColumns(2,
						ZeroBtn,
						DotBtn),
					EqualBtn,
				))),
	)

	w := myApp.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500, 280))

	w.SetContent(
		container.NewBorder(DeskBtn, nil, nil, nil, calculatorContainer),
	)
	w.Show()

}
