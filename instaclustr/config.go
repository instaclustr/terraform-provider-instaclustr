package instaclustr

const (
    //DefaultApiHostname string = "https://api.instaclustr.com"
    DefaultApiHostname string = "http://localhost:8090"

)

type Config struct {
    Username string
    ApiKey string
    ApiServerEnv string
    apiServerHostname string
    Client *APIClient
}

func (c *Config) Init() {
    c.Client = new(APIClient)
    c.Client.InitClient(c.apiServerHostname, c.Username, c.ApiKey)
}

func (c *Config) GetHostname() string {
    return c.apiServerHostname
}
