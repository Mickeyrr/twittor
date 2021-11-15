package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Mickeyrr/twittor/middlew"
	"github.com/Mickeyrr/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores set el puerto y Handler y pongo a escuchar al servidor */
func Manejadores() {
	// router -> sera para el manejo de las rutas de acceso al sistema
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.CheckBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.CheckBD(middlew.ValidateJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.CheckBD(middlew.ValidateJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckBD(middlew.ValidateJWT(routers.GrabarTweet))).Methods("POST")
	router.HandleFunc("/leertweets", middlew.CheckBD(middlew.ValidateJWT(routers.LeerTweets))).Methods("GET")
	router.HandleFunc("/eliminartweet", middlew.CheckBD(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckBD(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.CheckBD(middlew.ValidateJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/uploadBanner", middlew.CheckBD(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.CheckBD(middlew.ValidateJWT(routers.ObtenerBanner))).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.CheckBD(middlew.ValidateJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/eliminarRelacion", middlew.CheckBD(middlew.ValidateJWT(routers.DeleteRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.CheckBD(middlew.ValidateJWT(routers.ConsultarRelacion))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.CheckBD(middlew.ValidateJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leerTweetsFollow", middlew.CheckBD(middlew.ValidateJWT(routers.LeerTweetsFollow))).Methods("GET")

	// Si no existe una variable de entorno con el puerto se le asigna por defualt el puerto 8080
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	// handler -> Otorga permisos se conexion desde cualquier IP donde sea llamada la aplicacion
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
