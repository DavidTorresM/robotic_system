import React, { useState } from 'react';
import { View, Text, StyleSheet, Button, ActivityIndicator } from 'react-native';

const CompeticionSigueLineas = () => {
    const [competicion, setCompeticion] = useState(null);
    const [loading, setLoading] = useState(false);

    const handlePress = async () => {
        setLoading(true);
        try {
            const response = await fetch('http://192.168.0.208:3000/competicion?id=3');
            const data = await response.json();
            setCompeticion(data);
        } catch (error) {
            console.error('Error fetching competition:', error);
        } finally {
            setLoading(false);
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
                <Text>Ronda ID: {competicion.ronda.ID}</Text>
                <Text>Robot A: {competicion.ronda.RobotA.Nombre}</Text>
                <Text>Robot B: {competicion.ronda.RobotB.Nombre}</Text>
                {/* Agrega más detalles según sea necesario */}
            </View>
        );
    }

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Competición Sigue lineas</Text>
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
});

export default CompeticionSigueLineas;