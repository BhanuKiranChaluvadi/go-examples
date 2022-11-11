package main

import "path"

type Model struct {
	name  string
	value string
}

func (m *Model) GetAbsoluteInstallationPath() string {
	return path.Join(INSTALLATION_PATH, m.name, m.value)
}
