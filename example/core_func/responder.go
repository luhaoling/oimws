package core_func

import (
	"fmt"

	"github.com/OpenIMSDK/tools/errs"
	"github.com/openimsdk/openim-sdk-core/v3/pkg/sdkerrs"
)

type RespMessage struct {
	respMessagesChan chan *EventData
}

func NewRespMessage(respMessagesChan chan *EventData) *RespMessage {
	return &RespMessage{respMessagesChan: respMessagesChan}
}

func (r *RespMessage) sendOnSuccessResp(operationID string, data string) {
	r.respMessagesChan <- &EventData{
		Event:       Success,
		OperationID: operationID,
		Data:        data,
	}
}

func (r *RespMessage) sendOnErrorResp(operationID string, err error) {
	resp := &EventData{
		Event:       Failed,
		OperationID: operationID,
	}
	if code, ok := err.(errs.CodeError); ok {
		resp.ErrCode = int32(code.Code())
		resp.ErrMsg = code.Error()
	} else {
		resp.ErrCode = sdkerrs.UnknownCode
		resp.ErrMsg = fmt.Sprintf("error %T not implement CodeError: %s", err, err)
	}
	r.respMessagesChan <- resp
}

// sendEventFailedRespNoErr sends a failed event response without error details.
// event: Name of the event.
func (r *RespMessage) sendEventFailedRespNoErr(event string) {
	r.respMessagesChan <- &EventData{
		Event: event,
	}
}

// sendEventSuccessRespWithData sends a successful event response with associated data.
func (r *RespMessage) sendEventSuccessRespWithData(event string, data string) {
	r.respMessagesChan <- &EventData{
		Event: event,
		Data:  data,
	}
}

// sendEventSuccessRespNoData sends a successful event response without any associated data.
// This is included for completeness but not used in the above callback methods.
func (r *RespMessage) sendEventSuccessRespNoData(event string) {
	r.respMessagesChan <- &EventData{
		Event: event,
	}
}

// sendEventFailedRespNoData sends a failed event response with error code and message, without any associated data.
// This function may be used if there are any future error handling requirements in the SignalingCallback.
func (r *RespMessage) sendEventFailedRespNoData(event string, errCode int32, errMsg string) {
	r.respMessagesChan <- &EventData{
		Event:   event,
		ErrCode: errCode,
		ErrMsg:  errMsg,
	}
}
