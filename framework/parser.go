package framework

import (
	"unicode/utf8"
)

type ParserState struct {
	stream      []byte
	pos         int
	deepest     int
	currentRune rune
	currentSize int
	eos         bool
}

func (state *ParserState) syncCurrent() {
	state.eos = state.pos >= len(state.stream)
	if !state.eos {
		state.currentRune, state.currentSize = utf8.DecodeRune(state.stream[state.pos:])
	} else {
		state.currentRune = utf8.RuneError
		state.currentSize = 0
	}
}

func (state *ParserState) Pos() int {
	return state.pos
}

func (state *ParserState) Deepest() int {
	state.syncDeepest()
	return state.deepest
}

func (state *ParserState) Current() rune {
	return state.currentRune
}

func (state *ParserState) AtEndOfStream() bool {
	return state.eos
}

func (state *ParserState) Consume() bool {
	if !state.eos && state.currentSize != 0 {
		state.pos += state.currentSize
		state.syncCurrent()
		return true
	} else {
		return false
	}
}

func (state *ParserState) ConsumeRune(value rune) bool {
	if !state.eos && state.currentSize != 0 && state.currentRune == value {
		state.pos += state.currentSize
		state.syncCurrent()
		return true
	} else {
		return false
	}
}

func (state *ParserState) ConsumeRange(min rune, max rune) bool {
	if !state.eos && state.currentSize != 0 && state.currentRune >= min && state.currentRune <= max {
		state.pos += state.currentSize
		state.syncCurrent()
		return true
	} else {
		return false
	}
}

func (state *ParserState) Slice(begin int) string {
	return string(state.stream[begin:state.pos])
}

func (state *ParserState) Checkpoint() int {
	return state.pos
}

func (state *ParserState) syncDeepest() {
	if state.pos > state.deepest {
		state.deepest = state.pos
	}
}

func (state *ParserState) Restore(pos int) {
	state.syncDeepest()
	state.pos = pos
	state.syncCurrent()
}

func (state *ParserState) Repeat(min int, body func() bool) bool {
	count := 0
	checkpoint := state.Checkpoint()
	for body() {
		checkpoint = state.Checkpoint()
		count += 1
	}
	if count >= min {
		state.Restore(checkpoint)
		return true
	} else {
		return false
	}
}

func (state *ParserState) Choose(children ...func() bool) bool {
	checkpoint := state.Checkpoint()
	for i, child := range children {
		if i != 0 {
			state.Restore(checkpoint)
		}
		ok := child()
		if ok {
			return true
		}
	}
	return false
}

func (state *ParserState) Optional(child func() bool) {
	checkpoint := state.Checkpoint()
	ok := child()
	if !ok {
		state.Restore(checkpoint)
	}
}

func MakeParserState(stream []byte) *ParserState {
	state := &ParserState{
		stream: stream,
	}
	state.syncCurrent()
	return state
}
