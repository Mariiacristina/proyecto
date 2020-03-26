package schema

type Signin struct {
	Id_cliente int    `json:"id_cliente, omitempty"`
	Nombre     string `json:"nombre, omitempty"`
	Password   string `json: "password, omitempty"`
}

type Cliente struct {
	Id_cliente int    `json:"id_cliente, omitempty"`
	Nombre     string `json:"nombre, omitempty"`
	Password   string `json:"password, omitempty"`
	Tipo       string `json:"tipo, omitempty"`
	Nacimiento string `json:"nacimiento, omitempty"`
}

type Reptil struct {
	Tipo          string `json:"tipo, omitempty"`
	Sol_max       string `json:"sol_max, omitempty"`
	Sol_min       string `json:"sol_min, omitempty"`
	Temp_max      string `json:"temp_max, omitempty"`
	Temp_min      string `json:"temp_min, omitempty"`
	Humedad_min   string `json:"humedad_min, omitempty"`
	Config_chosen string `json:"config_chosen, omitempty"`
	Uv_inicio     string `json:"uv_inicio, omitempty"`
	Uv_tiempo     string `json:"uv_tiempo, omitempty"`
	Catarata_on   string `json:"catarata_on, omitempty"`
	Catarata_off  string `json:"catarata_off, omitempty"`
}

type Sensores struct {
	Sol					 int 	  `json:"sol, omitempty"`
	Terrario 		 int    `json:"terrario, omitempty"`
	Humedad  		 int    `json:"humedad, omitempty"`
}

type Objetos struct {
	Uv 							int		 `json:"uv, omitempty"`
	FocoTermico 	  int    `json:"focotermico, omitempty"`
	PlacaTermica  	int    `json:"placatermica, omitempty"`
	Catarata  			int    `json:"catarata, omitempty"`
}
