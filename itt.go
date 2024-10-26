package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Используйте: itt <lang>")
		return
	}

	cmd := exec.Command("xclip", "-selection", "clipboard", "-t", "image/png", "-o")
	imgBytes, err := cmd.Output()
	if err != nil {
		fmt.Println("Ошибка получения изображения из буфера обмена:", err)
		return
	}

	tempFile, err := ioutil.TempFile("", "clipboard-image-*.png")
	if err != nil {
		fmt.Println("Ошибка создания временного файла:", err)
		return
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(imgBytes); err != nil {
		fmt.Println("Ошибка записи изображения во временный файл:", err)
		return
	}
	tempFile.Close()

	client := gosseract.NewClient()
	defer client.Close()

	client.SetLanguage(os.Args[1])

	client.SetImage(tempFile.Name())

	text, err := client.Text()
	if err != nil {
		fmt.Println("Ошибка извлечения текста:", err)
		return
	}

	fmt.Println("Извлеченный текст:", text)
}
