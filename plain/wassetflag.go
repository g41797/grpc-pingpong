
package plain

type WasSetFlag {
	wasset bool
}

func (ws *WasSetFlag) SetFlag() {
	if ws != nil {
		ws.wasset = true
	}
}

func (ws *WasSetFlag) IsSet() bool {
	if ws == nil {return false}
	return ws.wasset
}
