package pedagio

// EnvioCalculo representa a caixa principal de dados para a API
type EnvioCalculo struct {
	DataHoraEmissao       string   `json:"dataHoraEmissao"`
	CodigoMunicipioOrigem int      `json:"codigoMunicipioOrigem"`
	UFMunicipioOrigem     string   `json:"ufMunicipioOrigem"`
	CST                   string   `json:"cst"`
	CClassTrib            string   `json:"cClassTrib"`
	BaseCalculo           float64  `json:"baseCalculo"`
	Trechos               []Trecho `json:"trechos"`
}

type Trecho struct {
	Numero    int     `json:"numero"`
	Municipio int     `json:"municipio"`
	UF        string  `json:"uf"`
	Extensao  float64 `json:"extensao"`
}
