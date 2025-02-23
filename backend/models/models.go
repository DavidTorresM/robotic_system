package models

import (
	"gorm.io/gorm"
)

type Categoria struct {
	ID          uint   `gorm:"primaryKey"`
	Nombre      string `gorm:"size:100;not null"`
	Descripcion string `gorm:"size:255"`
}

type Equipo struct {
	ID            uint           `gorm:"primaryKey"`
	Nombre        string         `gorm:"size:100;not null"`
	Descripcion   string         `gorm:"size:255"`
	RegistroFecha string         `gorm:"type:timestamp;default:now"`
	Participantes []Participante `gorm:"foreignKey:EquipoID" json:"Participantes,omitempty"` // Relación Uno a Muchos
	Robots        []Robot        `gorm:"foreignKey:EquipoID" json:"Robots,omitempty"`        // Relación Uno a Muchos
}

type Participante struct {
	ID                uint    `gorm:"primaryKey"`
	Nombre            string  `gorm:"size:100;not null"`
	Correo            string  `gorm:"size:100;unique;not null"`
	Password          string  `gorm:"size:100;" json:"Equipo,omitempty"`
	TokenVerificacion string  `gorm:"size:32;" json:"TokenVerificacion,omitempty"`
	Verified          bool    `gorm:"default:false"  json:"Verified"`
	Telefono          string  `gorm:"size:15"`
	EquipoID          uint    `gorm:"not null"`
	Equipo            *Equipo `gorm:"foreignKey:EquipoID" json:"Equipo,omitempty"` // Relación Muchos a Uno
}

type Robot struct {
	ID          uint   `gorm:"primaryKey"`
	Nombre      string `gorm:"size:100;not null"`
	Descripcion string `gorm:"size:255"`
	EquipoID    uint
	Equipo      *Equipo    `gorm:"foreignKey:EquipoID" json:"Equipo,omitempty"`
	CategoriaID uint       `gorm:"not null"`
	Categoria   *Categoria `gorm:"foreignKey:CategoriaID" json:"Categoria,omitempty"` // Relación Muchos a Uno
}

type Puntuacion struct {
	ID             uint    `gorm:"primaryKey"`
	RobotID        uint    `gorm:"not null"`
	ArbitroID      uint    `gorm:"not null"`
	CategoriaID    uint    `gorm:"not null"`
	Puntaje        float64 `gorm:"type:decimal(10,2)"`
	Tiempo         float64 `gorm:"type:decimal(10,2)"`
	ResultadoRonda string  `gorm:"size:50"`
	Ronda          uint
	Comentarios    string `gorm:"size:255"`
	FechaHora      string `gorm:"type:timestamp;default:now"`
	Robot          *Robot `gorm:"foreignKey:RobotID"` // Relación Muchos a Uno
}

type Ronda struct {
	ID                 uint `gorm:"primaryKey"`
	CategoriaID        uint `gorm:"not null"`
	RobotAID           uint
	RobotBID           uint
	RobotGanadorID     *uint
	NumeroRonda        uint
	FechaHoraInsercion string `gorm:"type:timestamp;default:now"`
	FechaHoraCompetion string `gorm:"type:timestamp;default:null"`
	RobotA             *Robot `gorm:"foreignKey:RobotAID"`                                     // Relación Muchos a Uno
	RobotB             *Robot `gorm:"foreignKey:RobotBID"`                                     // Relación Muchos a Uno
	RobotGanador       *Robot `gorm:"foreignKey:RobotGanadorID" json:"RobotGanador,omitempty"` // Relación Muchos a Uno
}
type RondaSigueLineas struct {
	ID                 uint `gorm:"primaryKey"`
	CategoriaID        uint `gorm:"not null"`
	RobotID            uint `gorm:"not null"`
	Intento            uint
	Tiempo             string `gorm:"type:TIME(6)"`
	SumaPenalizaciones uint   `gorm:"type:decimal(10,2)"`
	FechaHoraInsercion string `gorm:"type:timestamp;default:now"`
	FechaHoraCompetion string `gorm:"type:timestamp;default:null"`
	Robot              *Robot `gorm:"foreignKey:RobotID"` // Relación Muchos a Uno
}

type Arbitro struct {
	ID          uint       `gorm:"primaryKey"`
	Nombre      string     `gorm:"size:100;not null"`
	Correo      string     `gorm:"size:100;unique;not null"`
	Password    string     `gorm:"size:255;not null"`
	CategoriaID uint       `gorm:"not null"`
	Categoria   *Categoria `gorm:"foreignKey:CategoriaID"` // Relación Muchos a Uno
}

func MigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&Categoria{},
		&Equipo{},
		&Participante{},
		&Robot{},
		&Arbitro{},
		&Puntuacion{},
		&Ronda{},
		&RondaSigueLineas{},
	)
}

func InsertCategorias(db *gorm.DB) error {
	var count int64 = 0
	db.Model(&Categoria{}).Count(&count)
	if count > 0 {
		return nil
	}
	categorias := []Categoria{
		{Nombre: "Sumo", Descripcion: "Robots de sumo"},
		{Nombre: "Peleas", Descripcion: "Robots de peleas de robots"},
		{Nombre: "Seguidor de linea", Descripcion: "Seguidores de linea"},
	}

	for _, c := range categorias {
		if err := db.Create(&c).Error; err != nil {
			return err
		}
	}
	return nil
}
