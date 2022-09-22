# [macos-golink-wrapper](https://github.com/eisenxp/macos-golink-wrapper)

solution to "syscall.Mprotect panic: permission denied" in Golang on macOS Catalina 10.15.x when using gomonkey or gohook

## Usage

Assuming Go is installed in /opt/go

Navigate to Path: /opt/go/pkg/tool/darwin_amd64
Rename file link to original_link
Download link wrapper for replacement
Make sure link wrapper is executable: chmod +x link
Enjoy it