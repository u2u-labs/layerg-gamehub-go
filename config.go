package layerggamehub

type Mode string

const (
	Sandbox    Mode = "Sandbox"
	Production Mode = "Production"
)

func GetBaseURL(mode Mode) string {
	var baseURLs = map[Mode]string{
		Sandbox:    "https://agg-dev.layerg.xyz/api",
		Production: "https://agg-dev.layerg.xyz/api",
	}

	return baseURLs[mode]
}
