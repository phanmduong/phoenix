package registry

import (
	"nimbus/controller"
)

func GetControllerRegistry() map[string]interface{} {
	controllerMap := map[string]interface{}{
		"UserController": controller.UserControler{},
	}

	return controllerMap
}
