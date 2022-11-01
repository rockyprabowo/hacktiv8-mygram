package config

const DevDatabaseEngine = "sqlite"
const DevDatabaseDSN = "file::memory:?cache=shared"

func DevelopmentEnvironments() []string {
	return []string{"dev", "development"}
}

func ProductionEnvironments() []string {
	return []string{"prod", "production"}
}
