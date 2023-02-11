package main

import (
	"context"
	"fmt"

	"github.com/chesiver/snappy/shadowclient"
	"github.com/chesiver/snappy/ssh"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SelectFile() string {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg)",
				Pattern:     "*.png;*.jpg",
			},
			{
				DisplayName: "Videos (*.mov;*.mp4)",
				Pattern:     "*.mov;*.mp4",
			},
			{
				DisplayName: "Text (*.txt)",
				Pattern:     "*.txt",
			},
		},
	})
	if err != nil {
		runtime.LogFatal(a.ctx, "Error when select file in dialo!")
	}
	return path
}

func (a *App) UploadFile(path string) {
	err := ssh.CopyFile(path)
	if err != nil {
		runtime.LogFatal(a.ctx, "Error when upload file!")
	}
	runtime.LogInfo(a.ctx, "Upload succeeds!")
}

func (a *App) ConnectShadowsocks(host, port, cipher string) {
	// runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
	// 	Message: "Connection started!",
	// })
	addr := fmt.Sprintf("ss://%v:wjlydntms38@%v:%v", cipher, host, port)
	err := shadowclient.Start(addr, "", "wjlydntms38", ":1080")
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:   "Warning",
			Message: "Connection start failed!",
		})
	}
}

func (a *App) DumpLogContent(truncate bool) string {
	return shadowclient.DumpLogContent(truncate)
}
