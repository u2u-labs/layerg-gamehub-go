package layerggamehub

type Environment string

const (
	Development Environment = "Development"
	Production  Environment = "Production"
)

func GetBaseURL(env Environment) string {
	var baseURLs = map[Environment]string{
		Development: "https://agg-dev.layerg.xyz/api",
		Production:  "https://agg-dev.layerg.xyz/api",
	}

	return baseURLs[env]
}
