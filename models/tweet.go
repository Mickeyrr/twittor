package models

/* Tweet captura del Bodu, el  mensaje que nos llega*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
