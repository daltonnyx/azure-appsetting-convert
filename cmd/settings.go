package cmd

type Setting struct {
	Name        string `json:"name"`
	Value       any    `json:"value"`
	SlotSetting bool   `json:"slotSetting"`
}
