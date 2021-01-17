package main

type Configuration struct {
	Message string

	Guards map[string][]string
}

// OnConfigurationChange is invoked when configuration changes may have been made.
func (p *guard) OnConfigurationChange() error {
	var c Configuration

	if err := p.API.LoadPluginConfiguration(&c); err != nil {
		p.API.LogError(err.Error())
		return err
	}

	p.message = c.Message
	p.guards.Store(c.Guards)

	return nil

}

func (p *guard) getGuards() map[string][]string {
	return p.guards.Load().(map[string][]string)
}