import React from 'react';
import { View, Text, StyleSheet, Button, ScrollView } from 'react-native';
import { Image } from 'react-native';
import {
    createStaticNavigation,
    useNavigation,
  } from '@react-navigation/native';
const Home = () => {
    const navigation = useNavigation();
    return (
        <ScrollView contentContainerStyle={styles.container}>
            <Text style={styles.legend}>
                asdasBienvenido a nuestra app de gestión de equipos robóticos. Aquí puedes registrar tu equipo, ver marcadores y acceder a tu cuenta.
            </Text>
            <View style={styles.menu}>
                <Button
                    title="Marcadores"
                    onPress={() => navigation.navigate('Marcadores')}
                />
                <Button
                    title="Registrar equipo"
                    onPress={() => navigation.navigate('RegistrarEquipo')}
                />
                <Button
                    title="Login"
                    onPress={() => navigation.navigate('Login')}
                />
            </View>
        </ScrollView>
    );
};

const styles = StyleSheet.create({
    container: {
        flexGrow: 1,
        justifyContent: 'center',
        alignItems: 'center',
        padding: 20,
    },
    legend: {
        fontSize: 18,
        textAlign: 'center',
        marginBottom: 20,
    },
    menu: {
        width: '100%',
        justifyContent: 'space-around',
        alignItems: 'center',
    },
});

export default Home;