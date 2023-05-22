package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/exec"
)

type CommandManagement struct{}

func (cm *CommandManagement) CurlRequest(ctx echo.Context) error {
	// https://www.bd-pratidin.com/assets/news_images/2023/05/18/105320_bangladesh_pratidin_MoFA_pic_with_govt_logo.jpg

	// Execute the command
	err := execCommand("test/insert.sh")
	if err != nil {
		println("Error executing command:", err)
		return err
	}

	return ctx.JSON(http.StatusOK, "Command executed successfully")
}

func (cm *CommandManagement) CreateVm(ctx echo.Context) error {
	// TODO get vm name

	err := execCommand("test/create_vm.sh")
	if err != nil {
		println("Error creating vm:", err)
		return err
	}

	return ctx.JSON(http.StatusOK, "VM created successfully")
}

// param: executable bash file path.
func execCommand(file string) error {
	cmd := exec.Command("bash", file)

	// Get the output of the command to os.Stdout
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
