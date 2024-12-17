package rpc

func (i *InstructionInfoEnvelope) GetParsed() *InstructionInfo {
	if i.asInstructionInfo != nil {
		return i.asInstructionInfo
	}
	return nil
}
