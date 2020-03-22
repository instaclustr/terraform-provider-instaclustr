package instaclustr

const (
    ApiHostname string = "https://api.dev.instaclustr.com"
)

type Config struct {
    Username string
    ApiKey string
    ApiServerEnv string
    apiServerHostname string
    Client *APIClient
}

func (c *Config) Init() {
    c.apiServerHostname = ApiHostname
    c.Client = new(APIClient)
    c.Client.InitClient(c.apiServerHostname, c.Username, c.ApiKey)
}

func (c *Config) GetHostname() string {
    return c.apiServerHostname
}
