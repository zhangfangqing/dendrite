// generated by stringer -type=reqResWriterState; DO NOT EDIT

package tchannel

import "fmt"

const _reqResWriterState_name = "reqResWriterPreArg1reqResWriterPreArg2reqResWriterPreArg3reqResWriterComplete"

var _reqResWriterState_index = [...]uint8{0, 19, 38, 57, 77}

func (i reqResWriterState) String() string {
	if i < 0 || i+1 >= reqResWriterState(len(_reqResWriterState_index)) {
		return fmt.Sprintf("reqResWriterState(%d)", i)
	}
	return _reqResWriterState_name[_reqResWriterState_index[i]:_reqResWriterState_index[i+1]]
}
