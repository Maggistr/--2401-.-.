/* Задание 2: Кросс-компиляция для разных ОС
Для Windows (64-bit):
GOOS=windows GOARCH=amd64 go build -o hello.exe hello.go
Создаст hello.exe для Windows.

Для Linux (64-bit):
GOOS=linux GOARCH=amd64 go build -o hello_linux hello.go
Создаст hello_linux для Linux.

Для macOS (Intel/Apple Silicon, 64-bit):
GOOS=darwin GOARCH=amd64 go build -o hello_macos hello.go
Создаст hello_macos для macOS (Intel и Apple Silicon).

После выполнения команд в папке появятся 3 файла:
hello.exe (Windows)
hello_linux (Linux)
hello_macos (macOS)

Перенесите соответствующий файл на нужную ОС и запустите:
Windows: hello.exe
Linux: chmod +x hello_linux && ./hello_linux
macOS: chmod +x hello_macos && ./hello_macos
*/
