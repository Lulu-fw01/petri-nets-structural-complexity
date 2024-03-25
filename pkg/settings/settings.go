package settings

type Settings struct {
	AgentsToTransitions map[string][]string `json:"agentsToTransitions"`
	SilentTransitions   []string            `json:"silentTransitions"`
}
