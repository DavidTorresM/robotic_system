package services

import (
	"errors"
	"fmt"
	"robotica_concursos/controllers/vo"
	"robotica_concursos/models"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

func GetRondaCompeticionSumo() (*models.Ronda, error) {
	db := GetDatabase()
	var ronda models.Ronda
	err := db.Table("rondas").Where("fecha_hora_competion is null").First(&ronda).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no hay mas rondas de sumo disponibles. Posiblemente ya se hayan tomado todas las rondas")
		} else {
			return nil, fmt.Errorf("error al obtener la ronda de sumo: %w", err)
		}
	}
	db.Model(&ronda).Update("fecha_hora_competion", time.Now())
	//TODO falta poner el id del arbitro
	fmt.Printf("Ronda tomada por arbitro id:%d ronda:[%v]\n", -1, ronda)
	ronda.RobotA = &models.Robot{}
	ronda.RobotB = &models.Robot{}
	db.First(ronda.RobotA, "ID = ?", ronda.RobotAID)
	db.First(ronda.RobotB, "ID = ?", ronda.RobotBID)
	return &ronda, nil
}

func GetRondaCompeticionSigueLineas() (*models.RondaSigueLineas, error) {
	db := GetDatabase()
	var ronda models.RondaSigueLineas
	err := db.Table("ronda_sigue_lineas").Where("fecha_hora_competion is null").First(&ronda).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no hay mas rondas de sigue lineas disponibles. Posiblemente ya se hayan tomado todas las rondas")
		} else {
			return nil, fmt.Errorf("error al obtener la ronda de sigue lineas: %w", err)
		}
	}
	db.Model(&ronda).Update("fecha_hora_competion", time.Now())
	fmt.Printf("Ronda tomada por arbitro id:%d ronda:[%v]\n", -1, ronda)
	return &ronda, nil
}

func StartCompetitionByID(id int) error {
	db := GetDatabase()
	//Obtenemos la categoria por ID
	var categoria models.Categoria
	result := db.First(&categoria, id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("failed to check last start time: %w", errors.New("Id not found"))
	}
	if id == 3 {
		err := StartCompetitionSigueLineas(categoria.ID)
		if err != nil {
			return fmt.Errorf("failed to start competition: %w", err)
		}
	} else if id == 1 {
		err := StartCompetitionSumo(categoria.ID)
		if err != nil {
			return fmt.Errorf("failed to start competition: %w", err)
		}
	} else {
		return fmt.Errorf("failed to start competition: %w", errors.New("invalid competition id"))
	}

	return nil
}

func StartCompetitionSumo(id uint) error {
	db := GetDatabase()
	// Verificar si ya hay una ronda en los últimos 3 meses
	var ronda models.Ronda
	err := db.Table("rondas").
		Select("MAX(fecha_hora_insercion) AS fecha_hora_insercion").
		Scan(&ronda).Error
	if err != nil {
		return fmt.Errorf("error al obtener la última ronda: %w", err)
	}
	if ronda.FechaHoraInsercion != "" {
		layout := time.RFC3339
		fechaTime, err := time.Parse(layout, ronda.FechaHoraInsercion)
		if err != nil {
			return fmt.Errorf("error al parsear la fecha de la última ronda: %w", err)
		}
		if time.Now().Before(fechaTime.AddDate(0, 3, 0)) {
			return fmt.Errorf("ya hay una competicion iniciada en los ultimos 3 meses")
		}
	}
	//Consultamos los robots de la competicion de sumo
	var robots []models.Robot
	db.Where("categoria_id = ?", id).Find(&robots)
	if len(robots) == 0 {
		fmt.Println("No hay robots en la competicion de sumo")
		return fmt.Errorf("error al obtener los robots de sumo: %w", errors.New("No hay robots en la competicion de sumo"))
	}
	if len(robots)%2 != 0 {
		fmt.Println("No hay robots pares en la competicion de sumo, se insertara un robot dummy")
		robot := models.Robot{}
		robot.Nombre = "Robot Dummy"
		robot.CategoriaID = id
		if err := db.Create(&robot).Error; err != nil {
			fmt.Printf("error al insertar robot dummy: %w\n", err)
		}
		robots = append(robots, robot)
	}
	//Insertamos en las rondas
	mensajeErrores := []string{}
	for i := 0; i < len(robots); i += 2 {
		rondaNueva := models.Ronda{}
		robotA := robots[i]
		robotB := robots[i+1]
		rondaNueva.RobotAID = &robotA.ID
		rondaNueva.RobotBID = &robotB.ID
		rondaNueva.CategoriaID = id
		if err := db.Create(&rondaNueva).Error; err != nil {
			msjError := fmt.Sprintf("error al insertar la ronda para los robots (%d,%s) y (%d,%s): %w", robotA.ID, robotA.Nombre, robotB.ID, robotB.Nombre, err)
			mensajeErrores = append(mensajeErrores, msjError)
		}
	}
	if len(mensajeErrores) > 0 {
		fmt.Println("Ocurrieron errores al insertar las rondas")
		for i := 0; i < len(mensajeErrores); i++ {
			fmt.Println(mensajeErrores[i])
		}
		return fmt.Errorf("ocurrieron errores al insertar las rondas: %w", errors.New(strings.Join(mensajeErrores, "\n")))
	} else {
		return nil
	}
}
func StartCompetitionSigueLineas(id uint) error {
	db := GetDatabase()
	fmt.Println("Moviendo robots a la tabla de competicion sigue lineas")
	// Verificar si ya hay una ronda en los últimos 3 meses
	var ronda models.RondaSigueLineas
	err := db.Table("ronda_sigue_lineas").
		Select("MAX(fecha_hora_insercion) AS fecha_hora_insercion").
		Scan(&ronda).Error
	if err != nil {
		return fmt.Errorf("error al obtener la última ronda: %w", err)
	}
	if ronda.FechaHoraInsercion != "" {
		layout := time.RFC3339
		fechaTime, err := time.Parse(layout, ronda.FechaHoraInsercion)
		if err != nil {
			return fmt.Errorf("error al parsear la fecha de la última ronda: %w", err)
		}
		if time.Now().Before(fechaTime.AddDate(0, 3, 0)) {
			return fmt.Errorf("ya hay una competicion iniciada en los ultimos 3 meses")
		}
	}

	// Insertar robots en la ronda
	query := "INSERT INTO ronda_sigue_lineas (robot_id, categoria_id) SELECT id, categoria_id FROM robots WHERE categoria_id = " + strconv.Itoa(int(id))
	result := db.Exec(query)
	if result.Error != nil {
		return fmt.Errorf("ocurrio un fallo insertando los robots de sigue lineas: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no se insertaron robots en la competicion sigue lineas")
	}

	return nil
}

func SetWinnerSumo(req vo.RequestBody) error {
	db := GetDatabase()
	db = db.Debug()
	var ronda models.Ronda
	err := db.First(&ronda, req.IDRonda).Error
	if err != nil {
		return fmt.Errorf("error al actualizar la ronda: %w", err)
	}
	//Seteamos el nuevo ganador
	ronda.RobotGanadorID = &req.IDRobotGanador
	ronda.FechaHoraCompetion = time.Now().Format("2006-01-02 15:04:05.000000")
	err = db.Save(&ronda).Error
	if err != nil {
		return fmt.Errorf("error al actualizar la ronda: %w", err)
	}
	//Buscamos la siguiente ronda disponible
	ronda = models.Ronda{}
	err = db.Table("rondas").Where("robot_b_id is null").First(&ronda).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//Generar el insert de la ronda
			rondaNueva := models.Ronda{}
			rondaNueva.CategoriaID = ronda.CategoriaID
			rondaNueva.FechaHoraInsercion = time.Now().Format("2006-01-02 15:04:05.000000")
			rondaNueva.RobotBID = nil
			rondaNueva.RobotAID = &req.IDRobotGanador
			err = db.Create(&rondaNueva).Error
			if err != nil {
				return fmt.Errorf("error al insertar la nueva ronda: %w", err)
			}
			return nil
		} else {
			return fmt.Errorf("error al obtener la ronda de sumo: %w", err)
		}
	}
	//Hacemos el update de la ronda
	db.Model(&ronda).Update("robot_b_id", req.IDRobotGanador)
	return nil
}
