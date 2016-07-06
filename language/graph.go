package language

/* Generated from regenerate_cacao, do not edit by hand. */

type Program struct {
}

func CreateProgram() *Program {
	o := &Program{}
	return o
}

type FunctionGraph struct {
	controlRegionsHead *ControlRegion
	controlRegionsTail *ControlRegion
}

func (region *Program) CreateFunctionGraph() *FunctionGraph {
	o := &FunctionGraph{}
	return o
}

type functionGraphControlRegionsIterator struct {
	current *ControlRegion
}

func (src *FunctionGraph) appendToControlRegions(dst *ControlRegion) {
	if dst.controlRegionsPrev != nil || dst.controlRegionsNext != nil {
		panic(dst)
	}
	if src.controlRegionsHead == nil {
		src.controlRegionsHead = dst
		src.controlRegionsTail = dst
	} else {
		tail := src.controlRegionsTail
		tail.controlRegionsNext = dst
		dst.controlRegionsPrev = tail
		src.controlRegionsTail = dst
	}
}

func (src *FunctionGraph) IterControlRegions() functionGraphControlRegionsIterator {
	return functionGraphControlRegionsIterator{current: src.controlRegionsHead}
}

func (iter *functionGraphControlRegionsIterator) HasNext() bool {
	return iter.current != nil
}

func (iter *functionGraphControlRegionsIterator) GetNext() *ControlRegion {
	temp := iter.current
	iter.current = temp.controlRegionsNext
	return temp
}

type ControlRegion struct {
	controlRegionsNext *ControlRegion
	controlRegionsPrev *ControlRegion
	opsHead            *Operation
	opsTail            *Operation
	transfers          []*ControlRegion
}

func (region *FunctionGraph) CreateControlRegion(function *FunctionGraph) *ControlRegion {
	o := &ControlRegion{}
	function.appendToControlRegions(o)
	return o
}

type controlRegionOpsIterator struct {
	current *Operation
}

func (src *ControlRegion) appendToOps(dst *Operation) {
	if dst.opsPrev != nil || dst.opsNext != nil {
		panic(dst)
	}
	if src.opsHead == nil {
		src.opsHead = dst
		src.opsTail = dst
	} else {
		tail := src.opsTail
		tail.opsNext = dst
		dst.opsPrev = tail
		src.opsTail = dst
	}
}

func (src *ControlRegion) IterOps() controlRegionOpsIterator {
	return controlRegionOpsIterator{current: src.opsHead}
}

func (iter *controlRegionOpsIterator) HasNext() bool {
	return iter.current != nil
}

func (iter *controlRegionOpsIterator) GetNext() *Operation {
	temp := iter.current
	iter.current = temp.opsNext
	return temp
}

type controlRegionTransfersCreator struct {
	src   *ControlRegion
	index int
}

type controlRegionTransfersIterator struct {
	src   *ControlRegion
	index int
}

func (src *ControlRegion) CreateTransfers(count int) controlRegionTransfersCreator {
	src.transfers = make([]*ControlRegion, count)
	return controlRegionTransfersCreator{src: src}
}

func (iter *controlRegionTransfersCreator) SetNext(dst *ControlRegion) {
	iter.src.transfers[iter.index] = dst
	iter.index++
}

func (src *ControlRegion) IterTransfers() controlRegionTransfersIterator {
	return controlRegionTransfersIterator{src: src}
}

func (iter *controlRegionTransfersIterator) HasNext() bool {
	return iter.index < len(iter.src.transfers)
}

func (iter *controlRegionTransfersIterator) GetNext() *ControlRegion {
	temp := iter.src.transfers[iter.index]
	iter.index++
	return temp
}

type Operation struct {
	opsNext *Operation
	opsPrev *Operation
	inputs  []*Value
	outputs []*Value
}

func (region *FunctionGraph) CreateOperation(control *ControlRegion) *Operation {
	o := &Operation{}
	control.appendToOps(o)
	return o
}

type operationInputsCreator struct {
	src   *Operation
	index int
}

type operationInputsIterator struct {
	src   *Operation
	index int
}

func (src *Operation) CreateInputs(count int) operationInputsCreator {
	src.inputs = make([]*Value, count)
	return operationInputsCreator{src: src}
}

func (iter *operationInputsCreator) SetNext(dst *Value) {
	iter.src.inputs[iter.index] = dst
	iter.index++
}

func (src *Operation) IterInputs() operationInputsIterator {
	return operationInputsIterator{src: src}
}

func (iter *operationInputsIterator) HasNext() bool {
	return iter.index < len(iter.src.inputs)
}

func (iter *operationInputsIterator) GetNext() *Value {
	temp := iter.src.inputs[iter.index]
	iter.index++
	return temp
}

type operationOutputsCreator struct {
	src   *Operation
	index int
}

type operationOutputsIterator struct {
	src   *Operation
	index int
}

func (src *Operation) CreateOutputs(count int) operationOutputsCreator {
	src.outputs = make([]*Value, count)
	return operationOutputsCreator{src: src}
}

func (iter *operationOutputsCreator) SetNext(dst *Value) {
	iter.src.outputs[iter.index] = dst
	iter.index++
}

func (src *Operation) IterOutputs() operationOutputsIterator {
	return operationOutputsIterator{src: src}
}

func (iter *operationOutputsIterator) HasNext() bool {
	return iter.index < len(iter.src.outputs)
}

func (iter *operationOutputsIterator) GetNext() *Value {
	temp := iter.src.outputs[iter.index]
	iter.index++
	return temp
}

type Value struct {
}

func (region *FunctionGraph) CreateValue() *Value {
	o := &Value{}
	return o
}
