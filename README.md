# OS Interface

Welcome to OS Interface, a robust operating system interface crafted using Go Language and the Fyne Framework. This project seamlessly integrates various independent applications into one cohesive interface, enhancing your digital experience.

✨ Feel free to explore, contribute, and enhance your digital experience with OS Interface! 

## Project Overview & Demo:

#### 🌐 Weather App:
  - Stay updated with live weather data synchronized from OpenWeatherAPI.
  - Get instant access to temperature, wind speed, and rain conditions.
  - Supports popular cities for accurate weather forecasts..

![WeatherApp-ezgif com-optimize](https://github.com/Abhiram-TK/OS-Interface/assets/158244906/0eb9e95e-9566-49dd-ad0f-5d7b5fcce85a)

#### 🖼️ Gallery App:
  - Conveniently access and view your stored images using the built-in image storage feature.

![ezgif com-optimize](https://github.com/Abhiram-TK/OS-Interface/assets/158244906/16d57358-9a3e-4fca-a547-a3275b7a8d05)

#### 🎵 Audio Player App:
  - Effortlessly manage your audio files with the system's database integration.
  - Play, pause, and stop audio tracks with ease.

![AudioPlayerApp-ezgif com-optimize](https://github.com/Abhiram-TK/OS-Interface/assets/158244906/52fa7a13-99ee-4c81-bd1d-797a6bb40739)

#### 🧮 Calculator App:
  - Perform basic arithmetic operations and keep track of your calculation history.
  - Easily view and delete previous calculations.

![CalculatorApp-ezgif com-optimize](https://github.com/Abhiram-TK/OS-Interface/assets/158244906/1335a816-1bdc-4aa1-a19d-9dcf146ca894)

#### 📝 Text Editor App:
  - Seamlessly create, save, open, and edit documents with system's database support.

![TextEditorApp-ezgif com-optimize](https://github.com/Abhiram-TK/OS-Interface/assets/158244906/34ae31b1-b06e-4d02-bf3c-e4942f03ecac)

#### ✉️ SendMail App:
  - Send emails efficiently using SMTP integration with a mail host server.
  - Simply input sender and recipient email, subject, and message.

![SenMailApp-ezgif com-optimize](https://github.com/Abhiram-TK/OS-Interface/assets/158244906/83492477-68f1-49db-bd3b-53a3d26f4afc)

## Installation & Setup:

1. Download Go from [official website](https://go.dev/).
2. Install Go and verify the installation using `go --version`.
3. Download TDM-GCC Compiler from [sourceforge](https://sourceforge.net/projects/tdm-gcc/).
4. Install TDM-GCC Compiler and verify using `gcc --version`.

## Run Locally:

### Install Fyne ToolKit Library

To install the library, run the following code in your terminal:

```bash
  go get fyne.io/fyne/v2@latest
```

This is a common way to manage dependencies in Go projects using the `go get` command.

### Customize Image Path:

To set the background image path, follow these steps:

1. Open **osInterface.go**.
2. Locate the line:

```bash
  img = canvas.NewImageFromFile("D:\\OS Interface\\osbg.jpg")
```

3. Update the file path according to your local setup.

### OpenWeatherAPI Setup:

To use the OpenWeatherAPI in your project, follow these steps:

1. Get API Key:
- Sign up on [OpenWeather](https://openweathermap.org/).
- Generate and copy the API key.

2. Update **weather.go**:

- Locate the line:

```bash
  resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + selection + "&appid=example")
```

- Replace example with the copied API key.

Now your project is set up to fetch weather data using your personal OpenWeatherAPI key.

### SMTP Setup

To integrate SMTP for sending emails in your project, follow these steps:

1. Get SMTP Key: 
- Sign up on [Brevo](https://onboarding.brevo.com/account/register).
- Generate and copy the SMTP key.
   
2. Update **sendmail.go**:
   
- Locate the Line:

```bash
  dailer := gomail.NewDialer("smtp-relay.brevo.com", 587, "username", "password")
```

- Replace **username** with your login email.
- Replace **password** with your SMTP key value.

Now your project is set up to use the SMTP server for sending emails from A to B.

### Run the Project

To run the project, execute this code in your terminal:

```bash
  go run .\osInterface.go .\weather.go .\gallery.go .\audioplayer.go .\calculator.go .\texteditor.go .\sendmail.go
```

Now, your project is ready to go! 🚀

## Documentation Links:

- [Go Documentation](https://go.dev/doc/)
- [Fyne ToolKit Documentation](https://docs.fyne.io/) 

## Contact Information:

❤️ Connect with me on [LinkedIn](#insert_linkedin_profile_link_here). 

🌟 Thank you for your interest in OS Interface. Looking forward to connecting with you!
