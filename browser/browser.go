package browser

import (
	"log"
	"os/exec"
	"runtime"
)

// Open открывает указанный URL в браузере
func Open(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		cmd = "xdg-open"
		args = []string{url}
	}

	if err := exec.Command(cmd, args...).Start(); err != nil {
		log.Printf("Ошибка открытия браузера: %v\n", err)
	}
}