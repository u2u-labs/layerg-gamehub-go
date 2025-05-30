package layerggamehub

type Environment string

const (
	Dev  Environment = "dev"
	Prod Environment = "prod"
)

func GetBaseURL(env Environment) string {
	var baseURLs = map[Environment]string{
		Dev:  "https://agg-dev.layerg.xyz/api",
		Prod: "https://agg-dev.layerg.xyz/api",
	}

	return baseURLs[env]
}
