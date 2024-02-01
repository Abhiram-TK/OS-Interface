package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	gomail "gopkg.in/mail.v2"
)

func showEmail() {

	//a := app.New()

	w := myApp.NewWindow("Email Sender Client")

	// lets create UI for all fields

	title := canvas.NewText("Email Sender Client", color.White)
	title.TextSize = 22 // font size is increased
	title.Alignment = fyne.TextAlignCenter

	message_status := widget.NewLabel("Enter / Fill all fields")

	sender := widget.NewEntry()
	sender.SetPlaceHolder("Enter Sender Email")

	receiver := widget.NewEntry()
	receiver.SetPlaceHolder("Enter Receiver Email")

	subject := widget.NewEntry()
	subject.SetPlaceHolder("Enter Subject")

	message_body := widget.NewMultiLineEntry()
	message_body.SetPlaceHolder("Enter Message...")

	// create our first button

	send_btn := widget.NewButton("Send Email", func() {

		message := gomail.NewMessage()

		// message from/ sender email
		// remove all static values

		message.SetHeader("From", sender.Text)
		message.SetHeader("To", receiver.Text)
		message.SetHeader("Subject", subject.Text)
		message.SetBody("text/plain", message_body.Text)

		// put/type smtp client email host(server),port, username,password
		// you can use gmail and any other server.

		dailer := gomail.NewDialer("smtp-relay.sendinblue.com", 587, "abhiramsuresh109@gmail.com", "C74JZnNO98WL0tPv")

		err := dailer.DialAndSend(message)
		if err != nil {
			fmt.Println(err)
			// message/ error in message status
			message_status.Text = err.Error()
			message_status.Refresh()
		} else {
			fmt.Println("Mail Sent successfully...")
			// message/ error in message status
			message_status.Text = "Mail Sent successfully..."
			message_status.Refresh()
		}

	})

	w.Resize(fyne.NewSize(600, 400))

	// send all entries in container

	w.SetContent(container.NewVBox(
		title, // title should be centered
		sender,
		receiver,
		subject,
		message_body,
		send_btn,
		message_status,
	))

	w.Show()

}
