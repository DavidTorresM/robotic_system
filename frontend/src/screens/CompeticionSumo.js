import React, { useState } from 'react';
import { View, Text, StyleSheet, Button, ActivityIndicator, Alert, Platform } from 'react-native';
import {
    useNavigation,
  } from '@react-navigation/native';
const CompeticionSumo = () => {

    const navigation = useNavigation();
    const [competicion, setCompeticion] = useState(null);
    const [loading, setLoading] = useState(false);
    const [ganador, setGanador] = useState(null);

    const handlePress = async () => {
        setLoading(true);
        try {
            const response = await fetch('http://192.168.0.208:3000/competicion?id=1');
            if (!response.ok) {
                if (response.status === 404) {
                    alert('No hay más rondas de sumo disponibles. Posiblemente ya se hayan tomado todas las rondas.');
                } else {
                    throw new Error('Error fetching competition');
                }
            } else {
                const data = await response.json();
                data.ronda.RobotA.Puntos = 0;
                data.ronda.RobotB.Puntos = 0;
                setCompeticion(data);
            }
        } catch (error) {
            console.error('Error fetching competition:', error);
        } finally {
            setLoading(false);
        }
    };

    const handleSelectGanador = (robotId) => {
        setGanador(robotId);
    };
    const enviarGanador = async (idGanador, descalificado) => {
        fetch('http://192.168.0.208:3000/competicion/sumo/ganador', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                idRobotGanador: idGanador,
                idRonda: competicion.ronda.ID,
                ganadorA: competicion.ronda.RobotA.ID === idGanador ? 1 : 0,
                puntosRobotA: competicion.ronda.RobotA.Puntos,
                puntosRobotB: competicion.ronda.RobotB.Puntos,
                descalificado: descalificado,
            }),
        })
            .then((response) => {
                if (!response.ok) {
                    throw new Error('Error submitting winner');
                }
                return response.json();
            })
            .then((data) => {
                console.log('Ganador enviado exitosamente:', data);
                navigation.navigate('Competiciones');
            })
            .catch((error) => {
                console.error('Error submitting winner:', error);
            });
    }
    const handleSubmit = () => {
        if (Platform.OS === 'web') {
            if (window.confirm('¿Está seguro que desea enviar los datos?')) {
                // Lógica para enviar los datos
                console.log('Datos enviados');
                enviarGanador(ganador, false);
            }
        } else {
            Alert.alert(
                'Confirmación',
                '¿Está seguro que desea enviar los datos?',
                [
                    {
                        text: 'Cancelar',
                        style: 'cancel',
                    },
                    {
                        text: 'Aceptar',
                        onPress: () => {
                            // Lógica para enviar los datos
                            console.log('Datos enviados');
                            enviarGanador(ganador, false);
                        },
                    },
                ],
                { cancelable: false }
            );
        }
    };

    const handleAddPoint = (robot) => {
        setCompeticion((prevCompeticion) => {
            const updatedCompeticion = { ...prevCompeticion };
            updatedCompeticion.ronda[robot].Puntos += 1;
            return updatedCompeticion;
        });
    };

    const handleRemovePoint = (robot) => {
        setCompeticion((prevCompeticion) => {
            const updatedCompeticion = { ...prevCompeticion };
            if (updatedCompeticion.ronda[robot].Puntos > 0) {
                updatedCompeticion.ronda[robot].Puntos -= 1;
            }
            return updatedCompeticion;
        });
    };

    const handleDisqualify = (robot) => {

        if (Platform.OS === 'web') {
            if (window.confirm('¿Está seguro que desea descalificar a este competidor?')) {
                console.log('Descalificando al robot:', robot);
                enviarGanador(robot === 'RobotA' ? competicion.ronda.RobotB.ID : competicion.ronda.RobotA.ID, true);
            }
        } else {
            Alert.alert(
                'Confirmación de Descalificación',
                `¿Está seguro que desea descalificar a ${competicion.ronda[robot].Nombre}?`,
                [
                    {
                        text: 'Cancelar',
                        style: 'cancel',
                    },
                    {
                        text: 'Aceptar',
                        onPress: () => {
                            // Lógica para descalificar al robot
                            console.log(`${competicion.ronda[robot].Nombre} descalificado`);
                            enviarGanador(robot === 'RobotA' ? competicion.ronda.RobotB.ID : competicion.ronda.RobotA.ID, true);
                        },
                    },
                ],
                { cancelable: false }
            );
        }
    };

    if (loading) {
        return (
            <View style={styles.container}>
                <ActivityIndicator size="large" color="#0000ff" />
            </View>
        );
    }

    if (competicion) {
        return (
            <View style={styles.container}>
                <Text style={styles.title}>Competición en curso</Text>
                <View style={styles.row}>
                    <View style={styles.halfContainer}>
                        <Text style={styles.robotName}>{competicion.ronda.RobotA.Nombre}</Text>
                        <Text>Puntos: {competicion.ronda.RobotA.Puntos}</Text>
                        <Button title="Agregar Punto" onPress={() => handleAddPoint('RobotA')} />
                        <Button title="Quitar Punto" onPress={() => handleRemovePoint('RobotA')} />
                        <Button title="Descalificar" onPress={() => handleDisqualify('RobotA')} />
                        <Button
                            title="Seleccionar como Ganador"
                            onPress={() => handleSelectGanador(competicion.ronda.RobotA.ID)}
                            color={ganador === competicion.ronda.RobotA.ID ? 'green' : 'gray'}
                        />
                    </View>
                    <View style={styles.halfContainer}>
                        <Text style={styles.robotName}>{competicion.ronda.RobotB.Nombre}</Text>
                        <Text>Puntos: {competicion.ronda.RobotB.Puntos}</Text>
                        <Button title="Agregar Punto" onPress={() => handleAddPoint('RobotB')} />
                        <Button title="Quitar Punto" onPress={() => handleRemovePoint('RobotB')} />
                        <Button title="Descalificar" onPress={() => handleDisqualify('RobotB')} />
                        <Button
                            title="Seleccionar como Ganador"
                            onPress={() => handleSelectGanador(competicion.ronda.RobotB.ID)}
                            color={ganador === competicion.ronda.RobotB.ID ? 'green' : 'gray'}
                        />
                    </View>
                </View>
                <Button title="Enviar Datos" onPress={handleSubmit} />
            </View>
        );
    }

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Competición Sumo</Text>
            <Button title="Obtener Competición" onPress={handlePress} />
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#fff',
    },
    title: {
        fontSize: 24,
        fontWeight: 'bold',
    },
    row: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        width: '100%',
    },
    halfContainer: {
        flex: 1,
        alignItems: 'center',
    },
    robotName: {
        fontSize: 18,
        fontWeight: 'bold',
    },
});

export default CompeticionSumo;