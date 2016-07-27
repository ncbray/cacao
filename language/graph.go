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

func (src *FunctionGraph) appendToFunctionGraphControlRegions(dst *ControlRegion) {
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

func (region *Program) IterFunctionGraphControlRegions(src *FunctionGraph) functionGraphControlRegionsIterator {
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

func (region *FunctionGraph) CreateControlRegion() *ControlRegion {
	o := &ControlRegion{}
	region.appendToFunctionGraphControlRegions(o)
	return o
}

type controlRegionOpsIterator struct {
	current *Operation
}

func (region *FunctionGraph) appendToControlRegionOps(src *ControlRegion, dst *Operation) {
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

func (region *FunctionGraph) IterControlRegionOps(src *ControlRegion) controlRegionOpsIterator {
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

func (region *FunctionGraph) CreateControlRegionTransfers(src *ControlRegion, count int) controlRegionTransfersCreator {
	if src.transfers != nil {
		panic(src)
	}
	src.transfers = make([]*ControlRegion, count)
	return controlRegionTransfersCreator{src: src}
}

func (iter *controlRegionTransfersCreator) SetNext(dst *ControlRegion) {
	if iter.index >= len(iter.src.transfers) {
		panic(iter.src)
	}
	iter.src.transfers[iter.index] = dst
	iter.index++
}

func (region *FunctionGraph) IterControlRegionTransfers(src *ControlRegion) controlRegionTransfersIterator {
	return controlRegionTransfersIterator{src: src}
}

func (iter *controlRegionTransfersIterator) HasNext() bool {
	return iter.index < len(iter.src.transfers)
}

func (iter *controlRegionTransfersIterator) GetNext() *ControlRegion {
	if !iter.HasNext() {
		panic(iter.src)
	}
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
	region.appendToControlRegionOps(control, o)
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

func (region *FunctionGraph) CreateOperationInputs(src *Operation, count int) operationInputsCreator {
	if src.inputs != nil {
		panic(src)
	}
	src.inputs = make([]*Value, count)
	return operationInputsCreator{src: src}
}

func (iter *operationInputsCreator) SetNext(dst *Value) {
	if iter.index >= len(iter.src.inputs) {
		panic(iter.src)
	}
	iter.src.inputs[iter.index] = dst
	iter.index++
}

func (region *FunctionGraph) IterOperationInputs(src *Operation) operationInputsIterator {
	return operationInputsIterator{src: src}
}

func (iter *operationInputsIterator) HasNext() bool {
	return iter.index < len(iter.src.inputs)
}

func (iter *operationInputsIterator) GetNext() *Value {
	if !iter.HasNext() {
		panic(iter.src)
	}
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

func (region *FunctionGraph) CreateOperationOutputs(src *Operation, count int) operationOutputsCreator {
	if src.outputs != nil {
		panic(src)
	}
	src.outputs = make([]*Value, count)
	return operationOutputsCreator{src: src}
}

func (iter *operationOutputsCreator) SetNext(dst *Value) {
	if iter.index >= len(iter.src.outputs) {
		panic(iter.src)
	}
	iter.src.outputs[iter.index] = dst
	iter.index++
}

func (region *FunctionGraph) IterOperationOutputs(src *Operation) operationOutputsIterator {
	return operationOutputsIterator{src: src}
}

func (iter *operationOutputsIterator) HasNext() bool {
	return iter.index < len(iter.src.outputs)
}

func (iter *operationOutputsIterator) GetNext() *Value {
	if !iter.HasNext() {
		panic(iter.src)
	}
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
