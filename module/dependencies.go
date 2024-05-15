package module

import (
	"collector/utils"
	"runtime"
)

func InstallJava() error {
	if runtime.GOOS == "linux" {
		err := utils.Execute("apt-get", "update")
		if err != nil {
			return err
		}

		err = utils.Execute("apt", "install", "-y", "openjdk-11-jdk")
		if err != nil {
			return err
		}
	}
	return nil
}

func UninstallJava() error {
	if runtime.GOOS == "linux" {
		err := utils.Execute("apt", "remove", "-y", "openjdk-11-jdk")
		if err != nil {
			return err
		}
		err = utils.Execute("apt", "autoremove", "-y")
		if err != nil {
			return err
		}
	}
	return nil
}
