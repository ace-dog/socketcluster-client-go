package scclient

import (
	"github.com/ace-dog/socketcluster-client-go/scclient/models"
	"github.com/ace-dog/socketcluster-client-go/scclient/utils"
)

func (client *Client) Transmit(eventName string, data interface{}) {
	transmitObject := models.GetTransmitEventObject(eventName, data)
	transmitData := utils.SerializeDataIntoString(transmitObject)
	client.socket.SendText(transmitData)
}