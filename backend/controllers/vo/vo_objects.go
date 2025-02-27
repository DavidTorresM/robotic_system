package vo

type RequestBody struct {
	IDRobotGanador uint `json:"idRobotGanador"`
	IDRonda        int  `json:"idRonda"`
	GanadorA       int  `json:"ganadorA"`
	PuntosRobotA   int  `json:"puntosRobotA"`
	PuntosRobotB   int  `json:"puntosRobotB"`
	Descalificado  bool `json:"descalificado"`
}
