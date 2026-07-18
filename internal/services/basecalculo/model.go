package basecalculo

type IsMercadorias struct {
	AnoFatoGerador        int     `json:"anoFatoGerador"`
	ValorBem              float64 `json:"valorBem"`
	AjusteAcrescimo       float64 `json:"ajusteAcrescimo"`
	Juros                 float64 `json:"juros"`
	Multas                float64 `json:"multas"`
	Encargos              float64 `json:"encargos"`
	FreteCobrado          float64 `json:"freteCobrado"`
	OutrosTributos        float64 `json:"outrosTributos"`
	DemaisImportancias    float64 `json:"demaisImportancias"`
	Icms                  float64 `json:"icms"`
	Iss                   float64 `json:"iss"`
	Cosip                 float64 `json:"cosip"`
	Ipi                   float64 `json:"ipi"`
	DescontoIncondicional float64 `json:"descontoIncondicional"`
	Bonificacao           float64 `json:"bonificacao"`
	DevolucaoVenda        float64 `json:"devolucaoVenda"`
}
