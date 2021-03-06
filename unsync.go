package minidis

// ClearCommands removes the application commands from the guild.
// If there are no guilds specified using `SyncToGuilds()`, global commands will be removed.
func (m *Minidis) ClearCommands() error {
	for _, v := range m.guilds {
		// get application commands
		guildCommands, err := m.session.ApplicationCommands(m.AppID, v)
		if err != nil {
			return err
		}

		// remove each command
		for _, cmd := range guildCommands {
			if err = m.session.ApplicationCommandDelete(m.AppID, v, cmd.ID); err != nil {
				return err
			}
		}
	}

	return nil
}
