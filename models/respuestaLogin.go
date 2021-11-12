package models

/* RespuestaLogin contiene el token nque devuelve el login */
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
