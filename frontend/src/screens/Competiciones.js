import React, { useEffect, useState } from 'react';
import { View, Text, Button, StyleSheet } from 'react-native';

const Competiciones = ({ navigation }) => {
    const [competitions, setCompetitions] = useState([]);

    useEffect(() => {
        const fetchCompetitions = async () => {
            try {
                const response = await fetch('http://192.168.0.208:3000/categorias');
                const data = await response.json();
                setCompetitions(data);
            } catch (error) {
                console.error('Error fetching competitions:', error);
            }
        };

        fetchCompetitions();
    }, []);

    const handlePress = (competition) => {
        if (!competition) return;
        if (competition.ID==1) {
            navigation.navigate('CompeticionSumo');
        }else if (competition.ID==2) {
            navigation.navigate('CompeticionSigueLineas');
        } else if (competition.ID==3) {
            navigation.navigate('CompeticionSigueLineas');
        } 
    };

    return (
        <View style={styles.container}>
            <Text style={styles.title}>Competiciones</Text>
            {competitions.map((competition) => (
                <Button
                    key={competition.ID}
                    title={competition.Nombre}
                    onPress={() => handlePress(competition)}
                />
            ))}
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
        marginBottom: 20,
    },
});

export default Competiciones;