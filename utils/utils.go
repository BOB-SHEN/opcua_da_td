package utils

import (
	"database/sql"
	"log"
)

func CreateDatabase(taos *sql.DB) {
	_, err := taos.Exec("CREATE DATABASE IF NOT EXISTS DEC21")
	if err != nil {
		log.Fatalln("failed to create database, err:", err)
	}
}

func CreateStableT(taos *sql.DB) {
	_, err := taos.Exec("CREATE STABLE IF NOT EXISTS DEC21.Temperature (ts TIMESTAMP, varname BINARY(64), temperature FLOAT) TAGS (location BINARY(64), groupId INT)")
	if err != nil {
		log.Fatalln("failed to create stable, err:", err)
	}
}

func CreateStableP(taos *sql.DB) {
	_, err := taos.Exec("CREATE STABLE IF NOT EXISTS DEC21.Pressure (ts TIMESTAMP, varname BINARY(64), pressure FLOAT) TAGS (location BINARY(64), groupId INT)")
	if err != nil {
		log.Fatalln("failed to create stable, err:", err)
	}
}

func CreateStableL(taos *sql.DB) {
	_, err := taos.Exec("CREATE STABLE IF NOT EXISTS DEC21.Level (ts TIMESTAMP, varname BINARY(64), level FLOAT) TAGS (location BINARY(64), groupId INT)")
	if err != nil {
		log.Fatalln("failed to create stable, err:", err)
	}
}

func CreateStableF(taos *sql.DB) {
	_, err := taos.Exec("CREATE STABLE IF NOT EXISTS DEC21.Flow (ts TIMESTAMP, varname BINARY(64), flow FLOAT) TAGS (location BINARY(64), groupId INT)")
	if err != nil {
		log.Fatalln("failed to create stable, err:", err)
	}
}

func CreateStableVFD(taos *sql.DB) {
	_, err := taos.Exec("CREATE STABLE IF NOT EXISTS DEC21.VFD (ts TIMESTAMP, varname BINARY(64), speed FLOAT, torque FLOAT, power FLOAT, current FLOAT) TAGS (location BINARY(64), groupId INT)")
	if err != nil {
		log.Fatalln("failed to create stable, err:", err)
	}
}

func CreateStableTH(taos *sql.DB) {
	_, err := taos.Exec("CREATE STABLE IF NOT EXISTS DEC21.TemperatureH (ts TIMESTAMP, varname BINARY(64), temperature FLOAT) TAGS (location BINARY(64), groupId INT)")
	if err != nil {
		log.Fatalln("failed to create stable, err:", err)
	}
}
